package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinioProvider implementa StorageProvider usando MinIO.
type MinioProvider struct {
	Client *minio.Client
}

// NewMinioProvider crea una nueva instancia de MinioProvider.
func NewMinioProvider() (*MinioProvider, error) {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	if endpoint == "" {
		return nil, fmt.Errorf("MINIO_ENDPOINT no está configurada")
	}
	accessKeyID := os.Getenv("MINIO_ROOT_USER")
	if accessKeyID == "" {
		return nil, fmt.Errorf("MINIO_ROOT_USER no está configurada")
	}
	secretAccessKey := os.Getenv("MINIO_ROOT_PASSWORD")
	if secretAccessKey == "" {
		return nil, fmt.Errorf("MINIO_ROOT_PASSWORD no está configurada")
	}

	useSSL := false
	if sslStr := os.Getenv("MINIO_USE_SSL"); sslStr != "" {
		b, err := strconv.ParseBool(sslStr)
		if err == nil {
			useSSL = b
		}
	}
	region := os.Getenv("MINIO_REGION")

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
		Region: region,
	})
	if err != nil {
		return nil, fmt.Errorf("error al inicializar el cliente de MinIO: %v", err)
	}
	return &MinioProvider{Client: client}, nil
}

// Init en este caso no requiere acciones adicionales.
func (m *MinioProvider) Init() error {
	return nil
}

// CreateBucket crea un bucket si no existe.
func (m *MinioProvider) CreateBucket(bucketName string) error {
	ctx := context.Background()
	exists, err := m.Client.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("error verificando la existencia del bucket: %v", err)
	}
	if exists {
		return nil
	}

	if err := m.Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{
		Region: os.Getenv("MINIO_REGION"),
	}); err != nil {
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

	if err := m.Client.SetBucketPolicy(ctx, bucketName, policy); err != nil {
		return fmt.Errorf("error aplicando política pública al bucket %s: %v", bucketName, err)
	}

	return nil
}

// Upload sube un archivo al bucket.
func (m *MinioProvider) Upload(bucketName, objectName, filePath, contentType string) error {
	ctx := context.Background()
	_, err := m.Client.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return fmt.Errorf("error subiendo el archivo: %v", err)
	}
	return nil
}

// Download descarga un objeto del bucket.
func (m *MinioProvider) Download(bucketName, objectName string) (io.ReadCloser, error) {
	ctx := context.Background()
	obj, err := m.Client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("error descargando el objeto: %v", err)
	}
	return obj, nil
}

// DeleteObject elimina un objeto del bucket.
func (m *MinioProvider) DeleteObject(bucketName, objectName string) error {
	ctx := context.Background()
	if err := m.Client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{}); err != nil {
		return fmt.Errorf("error eliminando el objeto: %v", err)
	}
	return nil
}

func (m *MinioProvider) MoveObject(bucket, srcObject, dstObject string) error {
	ctx := context.Background()

	// 1) Copiar
	src := minio.CopySrcOptions{
		Bucket: bucket,
		Object: srcObject,
	}
	dst := minio.CopyDestOptions{
		Bucket: bucket,
		Object: dstObject,
	}
	if _, err := m.Client.CopyObject(ctx, dst, src); err != nil {
		return fmt.Errorf("error copiando objeto de %s a %s: %v", srcObject, dstObject, err)
	}

	// 2) Borrar el original
	if err := m.Client.RemoveObject(ctx, bucket, srcObject, minio.RemoveObjectOptions{}); err != nil {
		return fmt.Errorf("error borrando objeto original %s: %v", srcObject, err)
	}

	return nil
}
