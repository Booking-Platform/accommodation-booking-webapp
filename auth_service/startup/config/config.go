package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port       string
	UserDBPort string
	UserDBHost string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	return &Config{
		Port:       os.Getenv("AUTH_SERVICE_PORT"),
		UserDBHost: os.Getenv("AUTH_DB_HOST"),
		UserDBPort: os.Getenv("AUTH_DB_PORT"),
	}
}
