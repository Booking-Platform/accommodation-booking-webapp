package api

import (
	"context"
	"errors"
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/infrastructure/services"
	"github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
	accommodation "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"
	user_info "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/user_info_service"

	reservation "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type UserInfoHandler struct {
	accommodationReserveClientAddress string
	userInfoClientAddress             string
	accommodationClientAddress        string
}

func (handler UserInfoHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/host/getHostsForRating/{id}", handler.GetHostsForRatingByUserID)
	if err != nil {
		panic(err)
	}

}

func NewUserInfoHandler(accommodationReserveClientAddress, userInfoClientAddress, accommodationClientAddress string) Handler {
	return &UserInfoHandler{
		accommodationReserveClientAddress: accommodationReserveClientAddress,
		userInfoClientAddress:             userInfoClientAddress,
		accommodationClientAddress:        accommodationClientAddress,
	}
}

func (handler *UserInfoHandler) GetHostsForRatingByUserID(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	if id == "" {
		writeErrorResponse(w, http.StatusBadRequest, errors.New("invalid user id"))
		return
	}

	reservations, err := handler.getAllReservationsThatPassed(id)

	for _, reservation := range reservations.Reservations {
		accommodation, _ := handler.getAccommodationForReservation(reservation.AccommodationID)
		host, _ := handler.getHostInfo(accommodation.Accommodation.HostId)
		
	}

	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

}

func (handler *UserInfoHandler) getAllReservationsThatPassed(id string) (*accommodation_reserve.GetAllReservationsThatPassedResponse, error) {
	reservationClient := services.NewReservationClient(handler.accommodationReserveClientAddress)
	return reservationClient.GetAllReservationsThatPassed(context.TODO(), &reservation.IdMessageRequest{Id: id})
}

func (handler *UserInfoHandler) getHostInfo(id string) (*user_info.GetUserByIDResponse, error) {
	userClient := services.NewUserClient(handler.userInfoClientAddress)
	return userClient.GetUserByID(context.TODO(), &user_info.GetUserByIDRequest{Id: id})
}

func (handler *UserInfoHandler) getAccommodationForReservation(id string) (*accommodation.GetAccommodationByIdResponse, error) {
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	return accommodationClient.GetById(context.TODO(), &accommodation.GetAccommodationByIdRequest{Id: id})
}
