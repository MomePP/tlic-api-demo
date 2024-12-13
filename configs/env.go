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

// INFO: example .env
// DB_DSN=host=localhost user=your_user password=your_password dbname=your_db port=5432 sslmode=disable
// APP_PORT=8080
// JWT_SECRET=supersecretkey
