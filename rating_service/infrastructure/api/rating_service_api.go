package api

import (
	"context"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/rating_service"
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingHandler struct {
	service *application.RatingService

	pb.UnimplementedRatingServiceServer
}

func NewRatingHandler(service *application.RatingService) *RatingHandler {
	return &RatingHandler{
		service: service,
	}
}

func (handler *RatingHandler) CreateRating(ctx context.Context, request *pb.CreateRatingRequest) (*pb.BlankResponse, error) {

	rating, err := MapNewRatingToHostRating(request.NewRating)
	if err != nil {
		return nil, err
	}

	err = handler.service.RateHost(rating)

	if err != nil {
		return nil, err
	}

	response := &pb.BlankResponse{}

	return response, nil
}

func (handler *RatingHandler) CreateAccommodationRating(ctx context.Context, request *pb.CreateAccommodationRatingRequest) (*pb.BlankResponse, error) {

	rating, err := MapNewAccommodationRatingToAccommodationRating(request.NewAccommodationRating)
	if err != nil {
		return nil, err
	}

	err = handler.service.RateAccommodation(rating)

	if err != nil {
		return nil, err
	}

	response := &pb.BlankResponse{}

	return response, nil
}

func (handler *RatingHandler) GetAccommodationRatingsByAccommodationID(ctx context.Context, request *pb.IdMessageRequest) (*pb.AccommodationRatingsResponse, error) {

	accommodationName := request.Id

	accommodationratings, err := handler.service.GetAccommodationRatingsByAccommodationID(accommodationName)

	if err != nil {
		return nil, err
	}

	response := &pb.AccommodationRatingsResponse{
		Ratings: []*pb.AccommodationRating{},
	}
	for _, accommodationrating := range accommodationratings {
		current, _ := MapAccommodationRatingToPb(accommodationrating)

		response.Ratings = append(response.Ratings, current)
	}
	return response, nil
}
func (handler *RatingHandler) GetHostRatingsByHostID(ctx context.Context, request *pb.IdMessageRequest) (*pb.HostRatingsResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	hostRatings, err := handler.service.GetHostRatingsByHostID(objectId)

	if err != nil {
		return nil, err
	}

	response := &pb.HostRatingsResponse{
		Ratings: []*pb.HostRating{},
	}
	for _, rating := range hostRatings {
		current, _ := MapHostRatingToPb(rating)

		response.Ratings = append(response.Ratings, current)
	}
	return response, nil
}
