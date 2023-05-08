package api

import (
	"context"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/application"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/domain/model"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/auth_service"
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
