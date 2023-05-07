package api

import (
	"context"
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

	reservations := []*domain.Reservation{}

	reservationsWithStatusWAITING, err := handler.getReservationsWithStatusWAITING()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, res := range reservationsWithStatusWAITING.GetReservations() {
		// create a new reservation with the details from the current reservation
		reservation := domain.Reservation{
			Start:    res.StartDate,
			End:      res.EndDate,
			GuestNum: res.GuestNum,
		}

		reservations = append(reservations, &reservation)

		user := handler.getUserForReservation()
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *ReservationHandler) getReservationsWithStatusWAITING() (*accommodation_reserve.GetAllForConfirmationResponse, error) {
	reservationClient := services.NewReservationClient(handler.accommodationReserveClientAddress)
	return reservationClient.GetAllForConfirmation(context.TODO(), &reservation.GetAllForConfirmationRequest{})
}

func (handler *ReservationHandler) getUserForReservation(id string) (*user_info.GetUserByIDResponse, error) {
	reservationClient := services.NewReservationClient(handler.accommodationReserveClientAddress)
	return reservationClient.GetAllForConfirmation(context.TODO(), &reservation.GetAllForConfirmationRequest{})
}

func (handler *ReservationHandler) getAccommodationByForReservation(id string) (*accommodation.GetAccommodationByIdResponse, error) {
	userClient := services.NewUserClient(handler.userInfoClientAddress)
	return userClient.GetUserByID()
}
