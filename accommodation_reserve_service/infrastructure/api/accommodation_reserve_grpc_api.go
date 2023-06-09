package api

import (
	"context"
	"fmt"
	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"

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

	automaticConfirmation, parsingErr := strconv.ParseBool(request.NewReservation.AutomaticConfirmation)
	if parsingErr != nil {
		return nil, err
	}

	err = handler.service.Create(reservation, automaticConfirmation)

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

func (handler *AccommodationReserveHandler) GetReservedAccommodationIds(ctx context.Context, request *pb.GetReservedAccommodationIdsRequest) (*pb.GetReservedAccommodationIdsResponse, error) {

	ids, err := handler.service.GetReservedAccommodationsIds(request.SearchParams.From, request.SearchParams.To)
	if err != nil {
		return nil, err
	}

	accommodationIds := make([]string, len(ids))
	for i, objectId := range ids {
		accommodationIds[i] = objectId.Hex()
	}

	response := &pb.GetReservedAccommodationIdsResponse{
		Id: accommodationIds,
	}

	return response, nil
}

func (handler *AccommodationReserveHandler) GetReservationsByUserID(ctx context.Context, request *pb.IdMessageRequest) (*pb.GetReservationsByUserIDResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
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

func (handler *AccommodationReserveHandler) ChangeReservationStatus(ctx context.Context, req *pb.ChangeReservationStatusRequest) (*pb.ChangeReservationStatusResponse, error) {
	id := req.ReservationWithIdAndStatus.Id

	objectId, err := primitive.ObjectIDFromHex(id)
	statusNum, err := strconv.Atoi(req.ReservationWithIdAndStatus.Status)

	if err != nil {
		fmt.Println("Error:", err)
	}

	err = handler.service.ChangeReservationStatus(objectId, model.ReservationStatus(statusNum))
	if err != nil {
		return nil, err
	}

	response := &pb.ChangeReservationStatusResponse{}

	return response, nil
}

func (handler *AccommodationReserveHandler) GetAllReservationsThatPassed(ctx context.Context, request *pb.IdMessageRequest) (*pb.GetAllReservationsThatPassedResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	reservations, err := handler.service.GetAllReservationsThatPassed(objectId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetAllReservationsThatPassedResponse{
		Reservations: []*pb.Reservation{},
	}
	for _, reservation := range reservations {
		current := mapReservationPb(reservation)
		response.Reservations = append(response.Reservations, current)
	}
	return response, nil
}

func (handler *AccommodationReserveHandler) DeleteAllUserReservations(ctx context.Context, request *pb.DeleteAllUserReservationsRequest) (*pb.DeleteAllUserReservationsResponse, error) {
	id := request.Id
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result, _ := handler.service.DeleteAllUserReservations(objectId)

	return &pb.DeleteAllUserReservationsResponse{Status: result}, nil
}
