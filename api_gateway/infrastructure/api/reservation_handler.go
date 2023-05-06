package api

import (
	"context"
	"encoding/json"
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/infrastructure/services"
	"github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
	reservation "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
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

	response, err := json.Marshal("orderDetails")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	reservations, err := handler.getReservationsWithStatusWAITING()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *ReservationHandler) getReservationsWithStatusWAITING() (*accommodation_reserve.GetAllForConfirmationResponse, error) {
	reservationClient := services.NewReservationClient(handler.accommodationReserveClientAddress)
	return reservationClient.GetAllForConfirmation(context.TODO(), &reservation.GetAllForConfirmationRequest{})
}
