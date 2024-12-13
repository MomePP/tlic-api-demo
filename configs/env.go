package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load() // Load from a `.env` file
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}
