package api

import (
	"context"
	"encoding/json"
	"errors"
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
	err = mux.HandlePath("GET", "/reservation/getReservationsByUserID/{id}", handler.GetReservationsByUserID)
	if err != nil {
		panic(err)
	}

}

func (handler *ReservationHandler) GetAllForConfirmation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reservations, err := handler.getReservationsWithStatusWAITING()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	response := make([]*domain.Reservation, len(reservations.GetReservations()))

	for i, res := range reservations.GetReservations() {
		acc, err := handler.getAccommodationForReservation(res.AccommodationID)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		user, err := handler.getUserForReservation(res.UserID)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		response[i] = &domain.Reservation{
			Status:   res.Status,
			Start:    res.StartDate,
			End:      res.EndDate,
			GuestNum: res.GuestNum,
			Accommodation: domain.Accommodation{
				Id:      acc.GetAccommodation().Id,
				Name:    acc.GetAccommodation().Name,
				Address: acc.GetAccommodation().Address.GetCity() + "," + acc.GetAccommodation().Address.GetCountry(),
			},
			User: domain.User{
				Name:    user.Name,
				Surname: user.Surname,
			},
		}
	}

	writeJsonResponse(w, http.StatusOK, response)
}

func (handler *ReservationHandler) GetReservationsByUserID(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	if id == "" {
		writeErrorResponse(w, http.StatusBadRequest, errors.New("invalid user id"))
		return
	}

	reservations, err := handler.getReservationsByUserID(id)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	response := make([]*domain.Reservation, len(reservations.GetReservations()))

	for i, res := range reservations.GetReservations() {
		acc, err := handler.getAccommodationForReservation(res.AccommodationID)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		user, err := handler.getUserForReservation(res.UserID)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		response[i] = &domain.Reservation{
			Status:   res.Status,
			Start:    res.StartDate,
			End:      res.EndDate,
			GuestNum: res.GuestNum,
			Accommodation: domain.Accommodation{
				Id:      acc.GetAccommodation().Id,
				Name:    acc.GetAccommodation().Name,
				Address: acc.GetAccommodation().Address.GetCity() + "," + acc.GetAccommodation().Address.GetCountry(),
			},
			User: domain.User{
				Name:    user.Name,
				Surname: user.Surname,
			},
		}
	}

	writeJsonResponse(w, http.StatusOK, response)
}

func writeErrorResponse(w http.ResponseWriter, status int, err error) {
	http.Error(w, err.Error(), status)
}

func writeJsonResponse(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
	}
}
func (handler *ReservationHandler) getReservationsWithStatusWAITING() (*accommodation_reserve.GetAllForConfirmationResponse, error) {
	reservationClient := services.NewReservationClient(handler.accommodationReserveClientAddress)
	return reservationClient.GetAllForConfirmation(context.TODO(), &reservation.GetAllForConfirmationRequest{})
}

func (handler *ReservationHandler) getReservationsByUserID(id string) (*accommodation_reserve.GetReservationsByUserIDResponse, error) {
	reservationClient := services.NewReservationClient(handler.accommodationReserveClientAddress)
	return reservationClient.GetReservationsByUserID(context.TODO(), &reservation.IdMessageRequest{Id: id})
}

func (handler *ReservationHandler) getAccommodationForReservation(id string) (*accommodation.GetAccommodationByIdResponse, error) {
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	return accommodationClient.GetById(context.TODO(), &accommodation.GetAccommodationByIdRequest{Id: id})
}

func (handler *ReservationHandler) getUserForReservation(id string) (*user_info.GetUserByIDResponse, error) {
	userClient := services.NewUserClient(handler.userInfoClientAddress)
	return userClient.GetUserByID(context.TODO(), &user_info.GetUserByIDRequest{Id: id})
}
