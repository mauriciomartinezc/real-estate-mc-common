package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mauriciomartinezc/real-estate-mc-common/security"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func LoadEnv() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working directory: %w", err)
	}
	envPath := filepath.Join(cwd, "/cmd/.env")

	// Load .env file
	envMap, err := godotenv.Read(envPath)
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	// Process and expand environment variables
	for key, value := range envMap {
		expandedValue := expandEnvVars(value)
		os.Setenv(key, expandedValue)
	}

	return nil
}

// expandEnvVars expands ${VAR} or ${VAR:-default} patterns in the string
func expandEnvVars(s string) string {
	// Pattern to match ${VAR} or ${VAR:-default}
	re := regexp.MustCompile(`\$\{([^}:]+)(?::-([^}]*))?\}`)

	return re.ReplaceAllStringFunc(s, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match
		}

		varName := submatches[1]
		defaultValue := ""
		if len(submatches) >= 3 {
			defaultValue = submatches[2]
		}

		if value := os.Getenv(varName); value != "" {
			return value
		}
		return defaultValue
	})
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

	// Validate JWT secret strength using security package
	jwtSecret := getEnvironment("JWT_SECRET_KEY")
	isValid, errors := security.ValidateJWTSecret(jwtSecret)
	if !isValid {
		return fmt.Errorf("JWT secret validation failed: %s. Generate a strong secret with: openssl rand -base64 64", strings.Join(errors, ", "))
	}

	return nil
}

func getEnvironment(environmentName string) string {
	return strings.TrimSpace(os.Getenv(environmentName))
}

func getErrorSetEnvironment(environmentName string) error {
	return fmt.Errorf("the environment variable %s is not set", environmentName)
}
