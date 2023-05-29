package config

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/common/utils"
	"os"
)

type Config struct {
	Port                string
	AccommodationDBHost string
	AccommodationDBPort string
}

func NewConfig() *Config {
	utils.LoadEnv()

	return &Config{
		Port:                os.Getenv("ACCOMMODATION_SERVICE_PORT"),
		AccommodationDBHost: os.Getenv("ACCOMMODATION_DB_HOST"),
		AccommodationDBPort: os.Getenv("ACCOMMODATION_DB_PORT"),
	}
}
