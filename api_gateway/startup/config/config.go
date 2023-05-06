package config

type Config struct {
	Port                     string
	AccommodationReserveHost string
	AccommodationReservePort string
	AccommodationHost        string
	AccommodationPort        string
}

func NewConfig() *Config {
	return &Config{
		Port:                     "8000",
		AccommodationReserveHost: "localhost",
		AccommodationReservePort: "8001",
		AccommodationHost:        "localhost",
		AccommodationPort:        "8002",
	}
}
