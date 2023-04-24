package main

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/startup"
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
