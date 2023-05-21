package api

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/domain/model"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/auth_service"
)

func mapUser(userPb *pb.NewUser) (*model.User, error) {
	return &model.User{Email: userPb.Email, Password: userPb.Password, Role: userPb.Role}, nil
}

func mapUserPb(user *model.User) *pb.User {
	userPb := &pb.User{Id: user.Id.Hex(), Email: user.Email}
	return userPb
}
