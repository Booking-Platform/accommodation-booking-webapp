package startup

import (
	"fmt"
	rating "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/rating_service"
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/application"
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/infrastructure/api"
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/infrastructure/persistence"
	cfg "github.com/Booking-Platform/accommodation-booking-webapp/rating_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *cfg.Config
}

func NewServer(config *cfg.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.RatingDBPHost, server.config.RatingDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	ratingStore := server.initReservationStore(mongoClient)

	ratingService := server.initReservationsService(ratingStore)
	ratingHandler := server.initReservationHandler(ratingService)

	server.startGrpcServer(ratingHandler)
}

func (server *Server) initReservationStore(client *mongo.Client) domain.RatingStore {
	store := persistence.NewRatingMongoDBStore(client)
	return store
}

func (server *Server) initReservationsService(store domain.RatingStore) *application.RatingService {
	return application.NewRatingService(store)
}

func (server *Server) initReservationHandler(service *application.RatingService) *api.RatingHandler {
	return api.NewRatingHandler(service)
}

func (server *Server) startGrpcServer(ratingHandler *api.RatingHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	rating.RegisterRatingServiceServer(grpcServer, ratingHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
