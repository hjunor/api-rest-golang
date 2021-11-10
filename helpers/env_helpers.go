package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GevToEnv(key string) string {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}
