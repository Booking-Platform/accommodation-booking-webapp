package api

import (
	"context"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/application"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/domain/model"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/auth_service"
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
		"exp":    time.Now().Add(time.Minute * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("yE54RkqqgahuNbMCPmlrqbcoDeUXadi4ibXdsKjssQTwSl3FZaJoNyvc05553OGA"))

	return &pb.LoginUserResponse{TokenString: tokenString}, nil
}
