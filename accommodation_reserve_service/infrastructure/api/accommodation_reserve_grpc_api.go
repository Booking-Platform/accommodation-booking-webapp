package api

import (
	"context"
	"fmt"

	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/application"

	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationReserveHandler struct {
	pb.UnimplementedAccommodationReserveServiceServer
	service *application.AccommodationReserveService
}

func NewAccommodationReserveHandler(service *application.AccommodationReserveService) *AccommodationReserveHandler {
	return &AccommodationReserveHandler{
		service: service,
	}
}

func (handler *AccommodationReserveHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	fmt.Println("Hello from accommodation reserve handler")
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
