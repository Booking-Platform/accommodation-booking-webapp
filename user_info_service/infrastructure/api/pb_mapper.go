package api

import (
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/user_info_service"
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapUser(userPb *pb.NewUser) (*model.User, error) {
	id, err := primitive.ObjectIDFromHex(userPb.Id)
	if err != nil {
		return nil, err
	}
	return &model.User{Id: id, Name: userPb.Name, Surname: userPb.Surname}, nil
}

func mapUserPb(user *model.User) *pb.User {
	userPb := &pb.User{Id: user.Id.String(), Name: user.Name, Surname: user.Surname}
	return userPb
}
