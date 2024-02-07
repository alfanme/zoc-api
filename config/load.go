package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if os.Getenv("APP_ENV") != "production" {
		dotEnvErr := godotenv.Load()
		if dotEnvErr != nil {
			log.Fatal("Error loading .env file")
		}
	}
}