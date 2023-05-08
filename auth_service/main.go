package main

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/startup"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
