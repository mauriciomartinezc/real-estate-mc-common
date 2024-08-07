package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"strings"
)

func LoadEnv() error {
	cwd, err := os.Getwd()
	if err != nil {
		return errors.New("error getting current working directory LoadEnv()")
	}
	envPath := filepath.Join(cwd, "/cmd/.env")
	err = godotenv.Load(envPath)
	if err != nil {
		return err
	}

	return nil
}

func ValidateEnvironments() error {
	// Server
	if getEnvironment("SERVER_PORT") == "" {
		return getErrorSetEnvironment("SERVER_PORT")
	}
	if getEnvironment("ALLOWED_ORIGINS") == "" {
		return getErrorSetEnvironment("ALLOWED_ORIGINS")
	}
	if getEnvironment("SERVER_PORT") == "" {
		return getErrorSetEnvironment("SERVER_PORT")
	}
	if getEnvironment("ALLOWED_METHODS") == "" {
		return getErrorSetEnvironment("ALLOWED_METHODS")
	}
	if getEnvironment("JWT_SECRET_KEY") == "" {
		return getErrorSetEnvironment("JWT_SECRET_KEY")
	}
	// Database
	if getEnvironment("DB_HOST") == "" {
		return getErrorSetEnvironment("DB_HOST")
	}
	if getEnvironment("DB_PORT") == "" {
		return getErrorSetEnvironment("DB_PORT")
	}
	if getEnvironment("DB_USER") == "" {
		return getErrorSetEnvironment("DB_USER")
	}
	if getEnvironment("DB_PASSWORD") == "" {
		return getErrorSetEnvironment("DB_PASSWORD")
	}
	if getEnvironment("DB_NAME") == "" {
		return getErrorSetEnvironment("DB_NAME")
	}
	if getEnvironment("DB_SSL_MODE") == "" {
		return getErrorSetEnvironment("DB_SSL_MODE")
	}
	if getEnvironment("DB_SSL_CERT") == "" {
		return getErrorSetEnvironment("DB_SSL_CERT")
	}

	return nil
}

func getEnvironment(environmentName string) string {
	return strings.TrimSpace(os.Getenv(environmentName))
}

func getErrorSetEnvironment(environmentName string) error {
	return errors.New("the environment " + environmentName + " is not set")
}
