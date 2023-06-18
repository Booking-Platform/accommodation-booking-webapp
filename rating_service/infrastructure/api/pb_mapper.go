package api

import (
	"fmt"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/rating_service"
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

func MapNewRatingToHostRating(ratingPb *pb.NewRating) (*model.HostRating, error) {
	guestID, err := primitive.ObjectIDFromHex(ratingPb.GuestID)
	if err != nil {
		return nil, err
	}

	hostID, err := primitive.ObjectIDFromHex(ratingPb.HostID)
	if err != nil {
		return nil, err
	}

	ratingNumber, err := strconv.Atoi(ratingPb.Rating)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	rating := &model.HostRating{
		Rating:  ratingNumber,
		GuestId: guestID,
		HostId:  hostID,
	}

	return rating, nil
}

func MapNewAccommodationRatingToAccommodationRating(ratingPb *pb.NewAccommodationRating) (*model.AccommodationRating, error) {
	guestID, err := primitive.ObjectIDFromHex(ratingPb.GuestID)
	if err != nil {
		return nil, err
	}

	ratingNumber, err := strconv.Atoi(ratingPb.Rating)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	rating := &model.AccommodationRating{
		Rating:            ratingNumber,
		GuestId:           guestID,
		AccommodationName: ratingPb.GetAccommodationName(),
	}

	return rating, nil
}

func MapAccommodationRatingToPb(accommodationRating *model.AccommodationRating) (*pb.AccommodationRating, error) {
	return &pb.AccommodationRating{
		GuestID:           accommodationRating.GuestId.Hex(),
		AccommodationName: accommodationRating.AccommodationName,
		Rating:            strconv.Itoa(accommodationRating.Rating),
		Date:              accommodationRating.Date.String(),
	}, nil
}

func MapHostRatingToPb(hostRating *model.HostRating) (*pb.HostRating, error) {
	return &pb.HostRating{
		Rating:  strconv.Itoa(hostRating.Rating),
		Date:    hostRating.Date.String(),
		GuestID: hostRating.GuestId.Hex(),
		HostID:  hostRating.HostId.Hex(),
	}, nil
}
