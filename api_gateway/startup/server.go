package startup

import (
	"context"
	"fmt"
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/infrastructure/api"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"

	cfg "github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/startup/config"
	accommodation_reserve_Gw "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
	accommodation_Gw "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"
	auth_Gw "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/auth_service"
	rating_Gw "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/rating_service"
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

	authEndpoint := fmt.Sprintf("%s:%s", server.config.AuthHost, server.config.AuthPort)
	err = auth_Gw.RegisterAuthServiceHandlerFromEndpoint(context.TODO(), server.mux, authEndpoint, opts)
	if err != nil {
		panic(err)
	}

	ratingEndpoint := fmt.Sprintf("%s:%s", server.config.RatingHost, server.config.RatingPort)
	err = rating_Gw.RegisterRatingServiceHandlerFromEndpoint(context.TODO(), server.mux, ratingEndpoint, opts)
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
	authEndpoint := fmt.Sprintf("%s:%s", server.config.AuthHost, server.config.AuthPort)

	reservationHandler := api.NewReservationHandler(accommodationReserveEndpoint, userInfoEndpoint, accommodationEndpoint)
	reservationHandler.Init(server.mux)

	accommodationHandler := api.NewAccommodationHandler(accommodationReserveEndpoint, accommodationEndpoint)
	accommodationHandler.Init(server.mux)

	authHandler := api.NewAuthHandler(userInfoEndpoint, authEndpoint)
	authHandler.Init(server.mux)

	userInfoHandler := api.NewUserInfoHandler(accommodationReserveEndpoint, userInfoEndpoint, accommodationEndpoint)
	userInfoHandler.Init(server.mux)

}

func (server *Server) getHandlerCORSWrapped() http.Handler {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{server.config.AllowedCorsOrigin},
	})
	handler := corsMiddleware.Handler(server.mux)
	handler = validateToken(handler)
	return handler
}

func validateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractTokenFromRequest(r)
		if tokenString == "" {
			// No token provided, return an error response
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Define the secret key used for signing the token
			secretKey := []byte("your-secret-key")
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			// Token is invalid, return an error response
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Token is valid, continue to the next handler
		next.ServeHTTP(w, r)
	})
}

func extractTokenFromRequest(r *http.Request) string {
	// Extract the token from the request headers, query parameters, cookies, or any other location where you expect to receive the token.
	// Implement the logic to retrieve the token based on your specific requirements.
	token := r.Header.Get("Authorization")
	// Additional parsing or manipulation of the token string if needed

	return token
}
