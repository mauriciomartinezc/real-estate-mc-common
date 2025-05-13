package storage

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// AWSProvider implementa la interfaz StorageProvider usando AWS S3.
type AWSProvider struct {
	Client *s3.Client
	Region string
}

// NewAWSProvider crea e inicializa una instancia de AWSProvider leyendo la configuración desde variables de entorno.
// Se requieren las siguientes variables:
//   - AWS_REGION: Región de AWS (por ejemplo, "us-west-2")
//
// Además, AWS SDK buscará las credenciales en el entorno o archivos de configuración estándar.
func NewAWSProvider() (*AWSProvider, error) {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		return nil, fmt.Errorf("AWS_REGION no está configurada")
	}

	// Carga la configuración por defecto del SDK, la cual respeta variables de entorno, profiles, etc.
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("no se pudo cargar la configuración de AWS: %v", err)
	}

	client := s3.NewFromConfig(cfg)

	return &AWSProvider{
		Client: client,
		Region: region,
	}, nil
}

// Init inicializa el proveedor. En este caso, ya se hizo en NewAWSProvider, así que no hay acción adicional.
func (p *AWSProvider) Init() error {
	if p.Client == nil {
		return fmt.Errorf("el cliente de AWS S3 no está inicializado")
	}
	return nil
}

// CreateBucket crea un bucket en S3 si no existe.
func (p *AWSProvider) CreateBucket(bucketName string) error {
	ctx := context.Background()

	// Verifica si el bucket ya existe.
	_, err := p.Client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: &bucketName,
	})
	if err == nil {
		// El bucket existe, no se necesita crear.
		return nil
	}

	// Si el error indica que el bucket no existe, intentamos crearlo.
	_, err = p.Client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: &bucketName,
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(p.Region),
		},
	})
	if err != nil {
		return fmt.Errorf("error creando el bucket %s: %v", bucketName, err)
	}

	policy := fmt.Sprintf(`{
      "Version":"2012-10-17",
      "Statement":[{
          "Effect":"Allow",
          "Principal":"*",
          "Action":["s3:GetObject"],
          "Resource":["arn:aws:s3:::%s/*"]
      }]
    }`, bucketName)
	if _, err := p.Client.PutBucketPolicy(ctx, &s3.PutBucketPolicyInput{
		Bucket: &bucketName,
		Policy: &policy,
	}); err != nil {
		return fmt.Errorf("error aplicando bucket policy pública: %v", err)
	}

	return nil
}

// Upload sube un archivo a S3.
func (p *AWSProvider) Upload(bucketName, objectName, filePath, contentType string) error {
	ctx := context.Background()

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error abriendo el archivo %s: %v", filePath, err)
	}
	defer file.Close()

	_, err = p.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      &bucketName,
		Key:         &objectName,
		Body:        file,
		ContentType: &contentType,
		ACL:         types.ObjectCannedACLPublicRead,
	})
	if err != nil {
		return fmt.Errorf("error subiendo el objeto %s: %v", objectName, err)
	}
	return nil
}

// Download descarga un objeto de S3 y retorna un io.ReadCloser.
func (p *AWSProvider) Download(bucketName, objectName string) (io.ReadCloser, error) {
	ctx := context.Background()
	output, err := p.Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &objectName,
	})
	if err != nil {
		return nil, fmt.Errorf("error descargando el objeto %s: %v", objectName, err)
	}
	return output.Body, nil
}

// DeleteObject elimina un objeto de S3.
func (p *AWSProvider) DeleteObject(bucketName, objectName string) error {
	ctx := context.Background()
	_, err := p.Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &bucketName,
		Key:    &objectName,
	})
	if err != nil {
		return fmt.Errorf("error eliminando el objeto %s: %v", objectName, err)
	}
	return nil
}

func (p *AWSProvider) MoveObject(bucketName, srcKey, dstKey string) error {
	ctx := context.Background()

	// Construye la fuente en el formato esperado por S3
	copySource := fmt.Sprintf("%s/%s", bucketName, srcKey)

	// 1) Copiar
	_, err := p.Client.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     &bucketName,
		CopySource: &copySource,
		Key:        &dstKey,
		// (opcional) MetadataDirective: types.MetadataDirectiveCopy,
	})
	if err != nil {
		return fmt.Errorf("error copiando desde %s a %s: %v", srcKey, dstKey, err)
	}

	// 2) Esperar a que el objeto exista (opcional)
	_, err = p.Client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &bucketName,
		Key:    &dstKey,
	})
	if err != nil {
		return fmt.Errorf("error verificando copia a %s: %v", dstKey, err)
	}

	// 3) Borrar el original
	if _, err := p.Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &bucketName,
		Key:    &srcKey,
	}); err != nil {
		return fmt.Errorf("error borrando objeto original %s: %v", srcKey, err)
	}

	return nil
}
