package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port                string
	AccommodationDBHost string
	AccommodationDBPort string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	return &Config{
		Port:                os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDBHost: os.Getenv("ACCOMMODATION_DB_HOST"),
		AccommodationDBPort: os.Getenv("ACCOMMODATION_DB_PORT"),
	}
}
