package startup

import (
	"fmt"
	"log"
	"net"

	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/startup/config"
	accommodation_reserve "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"

	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/application"
	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/infrastructure/api"
	"google.golang.org/grpc"
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

	accommodationReserveService := server.initAccommodationReserveService()

	server.initAccommodationReserveHandler(accommodationReserveService)

	accommodationReserveHandler := server.initAccommodationReserveHandler(accommodationReserveService)

	server.startGrpcServer(accommodationReserveHandler)

}

func (server *Server) initAccommodationReserveService() *application.AccommodationReserveService {
	return application.NewAccommodationReserveService()
}

func (server *Server) initAccommodationReserveHandler(service *application.AccommodationReserveService) *api.AccommodationReserveHandler {
	return api.NewAccommodationReserveHandler(service)
}

func (server *Server) startGrpcServer(accommodationReserveHandler *api.AccommodationReserveHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	accommodation_reserve.RegisterAccommodationReserveServiceServer(grpcServer, accommodationReserveHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
