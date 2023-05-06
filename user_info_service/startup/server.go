package startup

import (
	"fmt"
	user_info "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/user_info_service"
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/application"
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/infrastructure/api"
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/infrastructure/persistence"
	cfg "github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/startup/config"
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
	reservationStore := server.initUserStore(mongoClient)

	accommodationReserveService := server.initUserReserveService(reservationStore)

	userService := accommodationReserveService
	accommodationReserveHandler := server.initUserInfoHandler(userService)

	userInfoHandler := accommodationReserveHandler
	server.startGrpcServer(userInfoHandler)
}

func (server *Server) initUserStore(client *mongo.Client) domain.UserStore {
	store := persistence.NewUserMongoDBStore(client)
	return store
}

func (server *Server) initUserReserveService(store domain.UserStore) *application.UserService {
	return application.NewUserService(store)
}

func (server *Server) initUserInfoHandler(service *application.UserService) *api.UserInfoHandler {
	return api.NewUserInfoHandler(service)
}

func (server *Server) startGrpcServer(accommodationReserveHandler *api.UserInfoHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	user_info.RegisterUserInfoServiceServer(grpcServer, accommodationReserveHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
