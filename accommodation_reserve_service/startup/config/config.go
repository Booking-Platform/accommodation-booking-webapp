package config

type Config struct {
	Port              string
	ReservationDBHost string
	ReservationDBPort string
}

func NewConfig() *Config {
	return &Config{
		Port:              "8001",
		ReservationDBHost: "localhost",
		ReservationDBPort: "27017",
	}
}
