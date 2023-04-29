package main

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/startup"
	cfg "github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/startup/config"
)

func main() {

	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
