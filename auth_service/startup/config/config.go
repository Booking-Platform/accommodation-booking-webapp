package config

type Config struct {
	Port       string
	UserDBPort string
	UserDBHost string
}

func NewConfig() *Config {
	return &Config{
		Port:       "8004",
		UserDBHost: "localhost",
		UserDBPort: "27017",
	}
}
