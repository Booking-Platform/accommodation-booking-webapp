package startup

import (
	"accommodation_service/application"
	"accommodation_service/infrastructure/api"
	"accommodation_service/startup/config"
	"fmt"
	accommodation "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {

	accommodationService := server.initAccommodationService()

	server.initAccommodationHandler(accommodationService)

	accommodationHandler := server.initAccommodationHandler(accommodationService)

	server.startGrpcServer(accommodationHandler)

}

func (server *Server) initAccommodationService() *application.AccommodationService {
	return application.NewAccommodationService()
}

func (server *Server) initAccommodationHandler(service *application.AccommodationService) *api.AccommodationHandler {
	return api.NewAccommodationHandler(service)
}

func (server *Server) startGrpcServer(accommodationHandler *api.AccommodationHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	accommodation.RegisterAccommodationServiceServer(grpcServer, accommodationHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}