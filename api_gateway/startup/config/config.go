package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port                     string
	AccommodationReserveHost string
	AccommodationReservePort string
	AccommodationHost        string
	AccommodationPort        string
	UserInfoHost             string
	UserInfoPort             string
	AuthHost                 string
	AuthPort                 string
	AllowedCorsOrigin        string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	return &Config{
		Port:                     os.Getenv("GATEWAY_PORT"),
		AccommodationReserveHost: os.Getenv("ACCOMMODATION_RESERVE_SERVICE_HOST"),
		AccommodationReservePort: os.Getenv("ACCOMMODATION_RESERVE_SERVICE_PORT"),
		AccommodationHost:        os.Getenv("ACCOMMODATION_SERVICE_HOST"),
		AccommodationPort:        os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		UserInfoHost:             os.Getenv("USER_INFO_SERVICE_HOST"),
		UserInfoPort:             os.Getenv("USER_INFO_SERVICE_PORT"),
		AuthHost:                 os.Getenv("AUTH_SERVICE_HOST"),
		AuthPort:                 os.Getenv("AUTH_SERVICE_PORT"),
	}
}
