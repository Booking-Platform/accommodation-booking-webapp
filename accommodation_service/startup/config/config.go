package config

type Config struct {
	Port                string
	AccommodationDBHost string
	AccommodationDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:                "8000",
		AccommodationDBHost: "accommodation_reserve_db",
		AccommodationDBPort: "27017",
	}
}
