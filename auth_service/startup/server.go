package startup

import (
	"context"
	"fmt"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/application"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/infrastructure/api"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/infrastructure/persistence"
	cfg "github.com/Booking-Platform/accommodation-booking-webapp/auth_service/startup/config"
	user_info "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/auth_service"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

func (server *Server) tokenValidationInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// Perform token validation here
	// Extract the token from the request context, validate it, and authorize the request
	// Implement your token validation logic based on your specific requirements
	// Example validation logic: check if the token is present in the request metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	token := authHeader[0]
	fmt.Printf(token)
	// Validate the token and perform authorization checks

	// If the token is valid and the request is authorized, proceed with the handler
	return handler(ctx, req)
}
