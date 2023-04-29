package api

import (
	"context"

	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/application"

	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
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

func (handler *AccommodationReserveHandler) CreateReservation(ctx context.Context, request *pb.CreateReservationRequest) (*pb.CreateReservationResponse, error) {

	reservation, err := mapReservation(request.NewReservation)
	if err != nil {
		return nil, err
	}

	err = handler.service.Create(reservation)

	if err != nil {
		return nil, err
	}

	response := &pb.CreateReservationResponse{
		Reservation: mapReservationPb(reservation),
	}

	return response, nil
}
