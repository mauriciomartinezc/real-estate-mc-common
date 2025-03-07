package storage

import "io"

type StorageProvider interface {
	// Init inicializa el proveedor.
	Init() error
	// CreateBucket crea un bucket si no existe.
	CreateBucket(bucketName string) error
	// Upload sube un objeto al bucket.
	Upload(bucketName, objectName, filePath, contentType string) error
	// Download descarga un objeto del bucket.
	Download(bucketName, objectName string) (io.ReadCloser, error)
	// DeleteObject elimina un objeto del bucket.
	DeleteObject(bucketName, objectName string) error
}
