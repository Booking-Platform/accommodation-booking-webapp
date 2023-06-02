package api

import (
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/rating_service"
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapRating(ratingPb *pb.NewRating) (*model.Rating, error) {
	guestID, err := primitive.ObjectIDFromHex(ratingPb.GuestID)
	if err != nil {
		return nil, err
	}

	hostID, err := primitive.ObjectIDFromHex(ratingPb.HostID)

	if err != nil {
		return nil, err
	}

	rating := &model.Rating{

		Rating:  1,
		HostId:  hostID,
		GuestId: guestID,
	}

	return rating, nil
}
