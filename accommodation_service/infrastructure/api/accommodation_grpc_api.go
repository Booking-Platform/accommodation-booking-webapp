package api

import (
	"accommodation_service/application"
	"context"
	"fmt"
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

func (handler *AccommodationHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	fmt.Println("Hello from accommodation handler")
	id := request.Id
	_, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	product := handler.service.Get()
	if product != nil {
		return nil, err
	}
	response := &pb.GetResponse{
		Hello: "caoooo",
	}
	return response, err
}
