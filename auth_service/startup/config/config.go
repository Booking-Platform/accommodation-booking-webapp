package config

type Config struct {
	Port       string
	UserDBPort string
	UserDBHost string
}

func NewConfig() *Config {
	return &Config{
		Port:       "8000",
		UserDBHost: "auth_db",
		UserDBPort: "27017",
	}
}
