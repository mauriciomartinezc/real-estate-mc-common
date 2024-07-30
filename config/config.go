package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetDSN() string {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")
	sslCert := os.Getenv("DB_SSL_CERT")
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	sslCert = filepath.Join(cwd, sslCert)
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s sslrootcert=%s",
		user,
		pass,
		host,
		port,
		dbName,
		sslMode,
		sslCert,
	)
}
