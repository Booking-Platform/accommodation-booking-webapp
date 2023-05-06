package main

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/startup"
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
