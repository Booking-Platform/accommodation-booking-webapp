package config

type Config struct {
	Port              string
	ReservationDBHost string
	ReservationDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:              "8000",
		ReservationDBHost: "accommodation_reserve_db",
		ReservationDBPort: "27017",
	}
}
