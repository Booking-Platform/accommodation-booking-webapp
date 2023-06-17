package api

import (
	"context"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/application"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/domain/model"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/auth_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserInfoHandler struct {
	service *application.UserService
	pb.UnimplementedAuthServiceServer
}

func NewUserHandler(service *application.UserService) *UserInfoHandler {
	return &UserInfoHandler{
		service: service,
	}
}

func (handler *UserInfoHandler) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	user, err := mapUser(request.NewUser)
	if err != nil {
		return nil, err
	}

	err = handler.service.Create(user)

	if err != nil {
		response := &pb.CreateUserResponse{User: mapUserPb(&model.User{
			Email: "error",
		})}
		return response, nil
	}

	response := &pb.CreateUserResponse{User: mapUserPb(user)}

	return response, nil
}

func (handler *UserInfoHandler) LoginUser(ctx context.Context, request *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	user, err := mapUser(request.NewUser)
	if err != nil {
		return nil, err
	}

	user, err = handler.service.Login(user)
	if err != nil {
		return &pb.LoginUserResponse{TokenString: "no match"}, nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.Id.Hex(),
		"role":   user.Role,
		"exp":    time.Now().Add(time.Minute * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_SECRET")))

	return &pb.LoginUserResponse{TokenString: tokenString}, nil
}

func (handler *UserInfoHandler) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result, _ := handler.service.Delete(objectId)
	return &pb.DeleteUserResponse{Status: result}, nil
}
