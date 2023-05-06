package startup

import (
	"context"
	"fmt"
	"log"
	"net/http"

	accommodation_reserve_Gw "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
	accommodation_Gw "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"

	cfg "github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/startup/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *cfg.Config
	mux    *runtime.ServeMux
}

func NewServer(config *cfg.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	server.initHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	accommodationReserveEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationReserveHost, server.config.AccommodationReservePort)
	err := accommodation_reserve_Gw.RegisterAccommodationReserveServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationReserveEndpoint, opts)
	if err != nil {
		panic(err)
	}

	//accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	//accom_err := accommodation_Gw.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationEndpoint, opts)
	//if accom_err != nil {
	//	panic(accom_err)
	//}

	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	err2 := accommodation_Gw.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationEndpoint, opts)
	if err2 != nil {
		panic(err2)
	}

}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}
