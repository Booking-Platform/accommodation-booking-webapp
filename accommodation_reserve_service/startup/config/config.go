package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port              string
	ReservationDBHost string
	ReservationDBPort string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	return &Config{
		Port:              os.Getenv("ACCOMMODATION_RESERVE_SERVICE_PORT"),
		ReservationDBHost: os.Getenv("ACCOMMODATION_RESERVE_DB_HOST"),
		ReservationDBPort: os.Getenv("ACCOMMODATION_RESERVE_DB_PORT"),
	}
}
