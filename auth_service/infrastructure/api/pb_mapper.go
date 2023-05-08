package api

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/domain/model"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/auth_service"
)

func mapUser(userPb *pb.NewUser) (*model.User, error) {
	return &model.User{Name: userPb.Name, Surname: userPb.Surname, Email: userPb.Email}, nil
}

func mapUserPb(user *model.User) *pb.User {
	userPb := &pb.User{Id: user.Id.String(), Name: user.Name, Surname: user.Surname, Email: user.Email}
	return userPb
}
