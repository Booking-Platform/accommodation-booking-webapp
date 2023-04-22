package config

import "os"

type Config struct {
	Address                            string
	AccommodationReserveServiceAddress string
}

func GetConfig() Config {
	return Config{
		AccommodationReserveServiceAddress: os.Getenv("ACCOMMODATION_RESERVE_SERVICE_ADDRESS"),
		Address:                            os.Getenv("GATEWAY_ADDRESS"),
	}
}
