package startup

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/config"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

