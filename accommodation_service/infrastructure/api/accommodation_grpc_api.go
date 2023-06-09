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
	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	accommodation, err := handler.service.GetAccommodationByID(objectId)
	if err != nil {
		return nil, err
	}

	accommodationPb := mapAccommodationPb(accommodation)
	response := &pb.GetAccommodationByIdResponse{
		Accommodation: accommodationPb,
	}
	return response, nil
}

func (handler *AccommodationHandler) GetAll(ctx context.Context, request *pb.GetAllAccommodationsRequest) (*pb.GetAllAccommodationsResponse, error) {
	accommodations, err := handler.service.GetAllAccommodations()
	if err != nil {
		return nil, err
	}

	response := &pb.GetAllAccommodationsResponse{
		Accommodations: []*pb.AccommodationDTO{},
	}
	for _, accommodation := range accommodations {
		current := mapAccommodationDTOPb(accommodation)
		response.Accommodations = append(response.Accommodations, current)
	}
	return response, nil
}

func (handler *AccommodationHandler) CreateAccommodation(ctx context.Context, request *pb.CreateAccommodationRequest) (*pb.CreateAccommodationResponse, error) {
	accommodation, err := mapNewAccommodation(request.NewAccommodation)
	if err != nil {
		return nil, err
	}

	err = handler.service.Create(accommodation)

	if err != nil {
		return nil, err
	}

	response := &pb.CreateAccommodationResponse{
		Accommodation: mapAccommodationPb(accommodation),
	}
	return response, nil
}

func (handler *AccommodationHandler) CreateAppointment(ctx context.Context, request *pb.CreateAppointmentRequest) (*pb.CreateAppointmentResponse, error) {
	accommodationId, err := primitive.ObjectIDFromHex(request.Appointment.Id)
	if err != nil {
		return nil, err
	}

	accommodationForUpdate, err := handler.service.GetAccommodationByID(accommodationId)
	if err != nil {
		return nil, err
	}

	err = handler.service.AddAppointment(accommodationForUpdate, mapPbAppointment(request.Appointment))
	if err != nil {
		return nil, err
	}

	updated, err := handler.service.GetAccommodationByID(accommodationId)
	if err != nil {
		return nil, err
	}

	response := &pb.CreateAppointmentResponse{
		Accommodation: mapAccommodationPb(updated),
	}

	return response, nil
}

func (handler *AccommodationHandler) Search(ctx context.Context, request *pb.GetAccommodationsByParamsRequest) (*pb.GetAccommodationsByParamsResponse, error) {
	accommodations, err := handler.service.GetAllAccommodationsByParams(request.SearchParams, request.SearchParams.Id)
	if err != nil {
		return nil, err
	}

	response := &pb.GetAccommodationsByParamsResponse{
		Accommodations: []*pb.AccommodationDTO{},
	}
	for _, accommodation := range accommodations {
		current := mapAccommodationDTOPb(accommodation)
		response.Accommodations = append(response.Accommodations, current)
	}
	return response, nil
}

func (handler *AccommodationHandler) ChangeAutomaticConfirmation(ctx context.Context, request *pb.ChangeAutomaticConfirmationRequest) (*pb.ChangeAutomaticConfirmationResponse, error) {
	objectId, err := primitive.ObjectIDFromHex(request.AccommodationID)

	err = handler.service.ChangeAutomaticConfirmationStatus(objectId)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (handler *AccommodationHandler) GetAllAccommodationsByHostID(ctx context.Context, request *pb.GetAllAccommodationsByHostIDRequest) (*pb.GetAllAccommodationsByHostIDResponse, error) {

	objectId, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}

	accommodations, err := handler.service.GetAllAccommodationsByHostID(objectId)
	if err != nil {
		return nil, err
	}

	response := &pb.GetAllAccommodationsByHostIDResponse{
		Accommodations: []*pb.AccommodationDTO{},
	}
	for _, accommodation := range accommodations {
		current := mapAccommodationDTOPb(accommodation)
		response.Accommodations = append(response.Accommodations, current)
	}
	return response, nil

}
