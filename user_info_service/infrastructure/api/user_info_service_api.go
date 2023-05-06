package api

import (
	"context"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/user_info_service"
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/application"
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/domain/model"
)

type UserInfoHandler struct {
	service *application.UserService
	pb.UnimplementedUserInfoServiceServer
}

func NewUserInfoHandler(service *application.UserService) *UserInfoHandler {
	return &UserInfoHandler{
		service: service,
	}
}

func (handler *UserInfoHandler) GetUserByID(ctx context.Context, request *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {

	user := model.User{

		Email:   "email@com.com",
		Name:    "bojan",
		Surname: "radovic",
	}

	response := &pb.GetUserByIDResponse{
		Surname: user.Surname,
		Name:    user.Name,
	}

	return response, nil
}
