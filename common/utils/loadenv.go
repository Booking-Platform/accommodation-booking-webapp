package utils

import (
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
	"runtime"
)

func LoadEnv() {
	var envPath string
	osName := runtime.GOOS

	if osName == "windows" {
		// Windows OS
		envPath = filepath.FromSlash("..\\.env")
	} else {
		// Linux or other OS
		envPath = filepath.FromSlash("./.env")
	}

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}
