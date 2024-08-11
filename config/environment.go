package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"strings"
)

func LoadEnv() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working directory: %w", err)
	}
	envPath := filepath.Join(cwd, "/cmd/.env")
	err = godotenv.Load(envPath)
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	return nil
}

func ValidateEnvironments() error {
	requiredEnvs := []string{
		"SERVER_PORT",
		"ALLOWED_ORIGINS",
		"ALLOWED_METHODS",
		"JWT_SECRET_KEY",
		"DB_HOST",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_SSL_MODE",
		"DB_SSL_CERT",
	}

	for _, env := range requiredEnvs {
		if getEnvironment(env) == "" {
			return getErrorSetEnvironment(env)
		}
	}

	return nil
}

func getEnvironment(environmentName string) string {
	return strings.TrimSpace(os.Getenv(environmentName))
}

func getErrorSetEnvironment(environmentName string) error {
	return fmt.Errorf("the environment variable %s is not set", environmentName)
}
