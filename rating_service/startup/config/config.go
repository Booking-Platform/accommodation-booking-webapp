package config

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/common/utils"
	"os"
)

type Config struct {
	Port          string
	RatingDBPHost string
	RatingDBPort  string
}

func NewConfig() *Config {
	utils.LoadEnv()

	return &Config{
		Port:          os.Getenv("RATING_SERVICE_PORT"),
		RatingDBPHost: os.Getenv("RATING_DB_HOST"),
		RatingDBPort:  os.Getenv("RATING_DB_PORT"),
	}
}
