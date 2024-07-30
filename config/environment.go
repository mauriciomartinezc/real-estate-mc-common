package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func LoadEnv() error {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	envPath := filepath.Join(cwd, "/cmd/.env")
	err = godotenv.Load(envPath)
	if err != nil {
		return err
	}

	return nil
}
