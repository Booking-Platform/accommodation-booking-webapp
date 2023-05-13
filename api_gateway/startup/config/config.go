package config

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
	return &Config{
		Port:                     "8000",
		AccommodationReserveHost: "accommodation_reserve_service",
		AccommodationReservePort: "8000",
		AccommodationHost:        "accommodation_service",
		AccommodationPort:        "8000",
		UserInfoHost:             "user_info_service",
		UserInfoPort:             "8000",
		AuthHost:                 "auth_service",
		AuthPort:                 "8000",

		AllowedCorsOrigin: "http://localhost:4200",
	}
}
