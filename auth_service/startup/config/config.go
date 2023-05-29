package config

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/common/utils"
	"os"
)

type Config struct {
	Port       string
	UserDBPort string
	UserDBHost string
}

func NewConfig() *Config {
	utils.LoadEnv()

	return &Config{
		Port:       os.Getenv("AUTH_SERVICE_PORT"),
		UserDBHost: os.Getenv("AUTH_DB_HOST"),
		UserDBPort: os.Getenv("AUTH_DB_PORT"),
	}
}
