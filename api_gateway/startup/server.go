package startup

import (
	"context"
	"fmt"
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/infrastructure/api"
	"log"
	"net/http"

	cfg "github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/startup/config"
	accommodation_reserve_Gw "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
	accommodation_Gw "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"
	user_info_Gw "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/user_info_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
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
	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	accommodationReserveEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationReserveHost, server.config.AccommodationReservePort)
	err := accommodation_reserve_Gw.RegisterAccommodationReserveServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationReserveEndpoint, opts)
	if err != nil {
		panic(err)
	}

	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)
	err = accommodation_Gw.RegisterAccommodationServiceHandlerFromEndpoint(context.TODO(), server.mux, accommodationEndpoint, opts)
	if err != nil {
		panic(err)
	}

	userInfoEndpoint := fmt.Sprintf("%s:%s", server.config.UserInfoHost, server.config.UserInfoPort)
	err = user_info_Gw.RegisterUserInfoServiceHandlerFromEndpoint(context.TODO(), server.mux, userInfoEndpoint, opts)
	if err != nil {
		panic(err)
	}

}

func (server *Server) Start() {
	handler := server.getHandlerCORSWrapped()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), handler))
}

func (server *Server) initCustomHandlers() {
	accommodationReserveEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationReserveHost, server.config.AccommodationReservePort)
	userInfoEndpoint := fmt.Sprintf("%s:%s", server.config.UserInfoHost, server.config.UserInfoPort)
	accommodationEndpoint := fmt.Sprintf("%s:%s", server.config.AccommodationHost, server.config.AccommodationPort)

	reservationHandler := api.NewReservationHandler(accommodationReserveEndpoint, userInfoEndpoint, accommodationEndpoint)
	reservationHandler.Init(server.mux)

	accommodationHandler := api.NewAccommodationHandler(accommodationReserveEndpoint, accommodationEndpoint)
	accommodationHandler.Init(server.mux)
}

func (server *Server) getHandlerCORSWrapped() http.Handler {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{server.config.AllowedCorsOrigin},
	})
	handler := corsMiddleware.Handler(server.mux)
	return handler
}
