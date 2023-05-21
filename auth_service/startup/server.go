package startup

import (
	"fmt"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/application"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/infrastructure/api"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/infrastructure/persistence"
	cfg "github.com/Booking-Platform/accommodation-booking-webapp/auth_service/startup/config"
	user_info "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/auth_service"
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
	client, err := persistence.GetClient(server.config.UserDBHost, server.config.UserDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	userStore := server.initUserStore(mongoClient)

	authService := server.initUserService(userStore)
	authHandler := server.initUserHandler(authService)

	server.startGrpcServer(authHandler)
}

func (server *Server) initUserStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	return store
}

func (server *Server) initUserService(store domain.UserStore) *application.UserService {
	return application.NewUserService(store)
}

func (server *Server) initUserHandler(service *application.UserService) *api.UserInfoHandler {
	return api.NewUserHandler(service)
}

func (server *Server) startGrpcServer(accommodationReserveHandler *api.UserInfoHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	user_info.RegisterAuthServiceServer(grpcServer, accommodationReserveHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
