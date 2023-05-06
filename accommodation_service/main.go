package main

import (
	"accommodation_service/startup"
	cfg "accommodation_service/startup/config"
)

func main() {

	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()

}
