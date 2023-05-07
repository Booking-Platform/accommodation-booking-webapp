package api

import (
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/user_info_service"
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/domain/model"
)

func mapUser(userPb *pb.NewUser) (*model.User, error) {
	return &model.User{Name: userPb.Name, Surname: userPb.Surname, Email: userPb.Email, Password: userPb.Password}, nil
}

func mapUserPb(user *model.User) *pb.User {
	userPb := &pb.User{Id: user.Id.String(), Name: user.Name, Surname: user.Surname, Email: user.Email, Password: user.Password}
	return userPb
}
