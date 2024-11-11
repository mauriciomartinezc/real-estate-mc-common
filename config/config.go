package config

import (
	"fmt"
	"github.com/mauriciomartinezc/real-estate-mc-common/cache"
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

func NewCacheClient() cache.Cache {
	var cacheClient cache.Cache

	if os.Getenv("CACHE_TYPE") == "redis" {
		cacheClient = cache.NewRedisCache(
			os.Getenv("CACHE_HOST")+":"+os.Getenv("CACHE_PORT"),
			os.Getenv("CACHE_PASSWORD"),
			0,
		)
	}

	if cacheClient == nil || os.Getenv("CACHE_TYPE") == "memory" {
		cacheClient = cache.NewInMemoryCache()
	}

	return cacheClient
}
