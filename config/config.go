package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetDSN() (string, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")
	sslCert := os.Getenv("DB_SSL_CERT")

	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory: %w", err)
	}
	sslCert = filepath.Join(cwd, sslCert)

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s sslrootcert=%s",
		user,
		pass,
		host,
		port,
		dbName,
		sslMode,
		sslCert,
	)

	return dsn, nil
}
