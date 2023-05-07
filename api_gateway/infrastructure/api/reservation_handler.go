package api

import (
	"context"
	"encoding/json"
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/infrastructure/services"
	"github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
	reservation "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
	accommodation "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"
	user_info "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/user_info_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type ReservationHandler struct {
	accommodationReserveClientAddress string
	userInfoClientAddress             string
	accommodationClientAddress        string
}

func NewReservationHandler(accommodationReserveClientAddress, userInfoClientAddress, accommodationClientAddress string) Handler {
	return &ReservationHandler{
		accommodationReserveClientAddress: accommodationReserveClientAddress,
		userInfoClientAddress:             userInfoClientAddress,
		accommodationClientAddress:        accommodationClientAddress,
	}
}

func (handler *ReservationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/reservation/getAllForConfirmation", handler.GetAllForConfirmation)
	if err != nil {
		panic(err)
	}

}

func (handler *ReservationHandler) GetAllForConfirmation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	// Create a slice to hold all the reservations
	reservations := []*domain.Reservation{}

	// Retrieve all reservations with status WAITING
	reservationsWithStatusWAITING, err := handler.getReservationsWithStatusWAITING()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Loop through each reservation and add it to the slice
	for _, res := range reservationsWithStatusWAITING.GetReservations() {
		// Retrieve the accommodation for the reservation
		fetchedAccommodation, err := handler.getAccommodationForReservation(res.AccommodationID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Retrieve the user for the reservation
		fetchedUser, err := handler.getUserForReservation(res.UserID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Create a new reservation with the details from the current reservation
		reservation := &domain.Reservation{
			Start:    res.StartDate,
			End:      res.EndDate,
			GuestNum: res.GuestNum,
			Accommodation: domain.Accommodation{
				Id:      fetchedAccommodation.GetAccommodation().Id,
				Name:    fetchedAccommodation.GetAccommodation().Name,
				Address: fetchedAccommodation.GetAccommodation().Address.String(),
			},
			User: domain.User{
				Name:    fetchedUser.Name,
				Surname: fetchedUser.Surname,
			},
		}

		// Add the reservation to the slice
		reservations = append(reservations, reservation)
	}

	// Marshal the slice into JSON
	response, err := json.Marshal(reservations)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *ReservationHandler) getReservationsWithStatusWAITING() (*accommodation_reserve.GetAllForConfirmationResponse, error) {
	reservationClient := services.NewReservationClient(handler.accommodationReserveClientAddress)
	return reservationClient.GetAllForConfirmation(context.TODO(), &reservation.GetAllForConfirmationRequest{})
}

func (handler *ReservationHandler) getAccommodationForReservation(id string) (*accommodation.GetAccommodationByIdResponse, error) {
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	return accommodationClient.GetById(context.TODO(), &accommodation.GetAccommodationByIdRequest{Id: id})
}

func (handler *ReservationHandler) getUserForReservation(id string) (*user_info.GetUserByIDResponse, error) {
	userClient := services.NewUserClient(handler.userInfoClientAddress)
	return userClient.GetUserByID(context.TODO(), &user_info.GetUserByIDRequest{Id: id})
}
