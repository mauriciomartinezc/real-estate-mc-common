package storage

import "errors"

// NewStorageProvider devuelve una implementación de StorageProvider.
func NewStorageProvider(storage string) (StorageProvider, error) {
	if storage == "minio" {
		return NewMinioProvider()
	}
	if storage == "aws" {
		return NewAWSProvider()
	}
	return nil, errors.New("storage provider not supported")
}
