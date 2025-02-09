package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadENV() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
