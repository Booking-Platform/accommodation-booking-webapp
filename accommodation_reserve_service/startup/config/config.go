package config

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/common/utils"
	"os"
)

type Config struct {
	Port              string
	ReservationDBHost string
	ReservationDBPort string
}

func NewConfig() *Config {
	utils.LoadEnv()

	return &Config{
		Port:              os.Getenv("ACCOMMODATION_RESERVE_SERVICE_PORT"),
		ReservationDBHost: os.Getenv("ACCOMMODATION_RESERVE_DB_HOST"),
		ReservationDBPort: os.Getenv("ACCOMMODATION_RESERVE_DB_PORT"),
	}
}
