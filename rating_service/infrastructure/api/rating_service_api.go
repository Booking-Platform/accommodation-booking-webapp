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

	rating, err := mapRating(request.NewRating)
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

	rating, err := mapAccommodationRating(request.NewAccommodationRating)
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

func (handler *RatingHandler) GetAccommodationRatingsByAccommodationID(ctx context.Context, request *pb.IdMessageRequest) (*pb.RatingsResponse, error) {

	_, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
