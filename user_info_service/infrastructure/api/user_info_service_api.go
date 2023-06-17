package api

import (
	"context"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/user_info_service"
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	user, _ := handler.service.GetByID(id)

	response := &pb.GetUserByIDResponse{
		Surname: user.Surname,
		Name:    user.Name,
		//		AvgRating: 4.8,
	}

	return response, nil
}

func (handler *UserInfoHandler) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	user, err := mapUser(request.NewUser)
	if err != nil {
		return nil, err
	}

	err = handler.service.Create(user)

	if err != nil {
		return nil, err
	}

	response := &pb.CreateUserResponse{User: mapUserPb(user)}

	return response, nil
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
