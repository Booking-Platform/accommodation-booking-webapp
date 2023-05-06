package api

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

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

func (handler *AccommodationReserveHandler) GetAllForConfirmation(ctx context.Context, request *pb.GetAllForConfirmationRequest) (*pb.GetAllForConfirmationResponse, error) {
	reservations, err := handler.service.GetAllForConfirmation()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllForConfirmationResponse{
		Reservations: []*pb.Reservation{},
	}
	for _, reservation := range reservations {
		current := mapReservationPb(reservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}

func (handler *AccommodationReserveHandler) GetReservationsByUserID(ctx context.Context, request *pb.GetReservationsByUserIDRequest) (*pb.GetReservationsByUserIDResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	reservations, err := handler.service.GetAllByUserID(objectId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetReservationsByUserIDResponse{
		Reservations: []*pb.Reservation{},
	}
	for _, reservation := range reservations {
		current := mapReservationPb(reservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}
