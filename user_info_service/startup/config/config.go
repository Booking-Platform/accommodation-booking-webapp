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
		Port:       os.Getenv("USER_INFO_SERVICE_PORT"),
		UserDBHost: os.Getenv("USER_INFO_DB_HOST"),
		UserDBPort: os.Getenv("USER_INFO_DB_PORT"),
	}
}
