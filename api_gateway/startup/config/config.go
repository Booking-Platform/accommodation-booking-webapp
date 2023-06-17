package config

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/common/utils"
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
	RatingHost               string
	RatingPort               string
	AllowedCorsOrigin        []string
	AllowedMethods           []string
	AllowedHeaders           []string
	AllowCredentials         bool
}

func NewConfig() *Config {

	utils.LoadEnv()

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
		RatingHost:               os.Getenv("RATING_SERVICE_HOST"),
		RatingPort:               os.Getenv("RATING_SERVICE_PORT"),

		AllowedCorsOrigin: []string{os.Getenv("CLIENT_PORT")},
		AllowedMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:    []string{"Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		AllowCredentials:  true,
	}
}
