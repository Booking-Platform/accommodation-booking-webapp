package config

type Config struct {
	Port                     string
	AccommodationReserveHost string
	AccommodationReservePort string
	AccommodationHost        string
	AccommodationPort        string
	UserInfoHost             string
	UserInfoPort             string
	AllowedCorsOrigin        string
}

func NewConfig() *Config {
	return &Config{
		Port:                     "8000",
		AccommodationReserveHost: "localhost",
		AccommodationReservePort: "8001",
		AccommodationHost:        "localhost",
		AccommodationPort:        "8002",
		UserInfoHost:             "localhost",
		UserInfoPort:             "8003",

		AllowedCorsOrigin: "http://localhost:4200",
	}
}
