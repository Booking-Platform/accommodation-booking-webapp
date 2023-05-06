package api

import (
	"accommodation_service/application"
	"context"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationHandler struct {
	pb.UnimplementedAccommodationServiceServer
	service *application.AccommodationService
}

func NewAccommodationHandler(service *application.AccommodationService) *AccommodationHandler {
	return &AccommodationHandler{
		service: service,
	}
}

func (handler *AccommodationHandler) GetById(ctx context.Context, request *pb.GetAccommodationByIdRequest) (*pb.GetAccommodationByIdResponse, error) {
	id := request.Id
	_, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	product := handler.service.Get()
	if product != nil {
		return nil, err
	}
	response := &pb.GetAccommodationByIdResponse{}
	return response, err
}
