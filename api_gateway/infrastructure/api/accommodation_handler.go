package api

import (
	"context"
	"encoding/json"
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/infrastructure/services"
	reservation "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
	accommodation "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"
	user_info "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/user_info_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"strconv"
)

type AccommodationHandler struct {
	accommodationReserveClientAddress string
	accommodationClientAddress        string
	userInfoClientAddress             string
}

func NewAccommodationHandler(accommodationReserveClientAddress, userInfoClientAddress, accommodationClientAddress string) Handler {
	return &AccommodationHandler{
		accommodationReserveClientAddress: accommodationReserveClientAddress,
		accommodationClientAddress:        accommodationClientAddress,
		userInfoClientAddress:             userInfoClientAddress,
	}
}

func (handler *AccommodationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/reservation/getAllByParams", handler.GetAllByParams)
	if err != nil {
		panic(err)
	}
}

func (handler *AccommodationHandler) GetAllByParams(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")
	numOfGuests := r.URL.Query().Get("numOfGuests")
	city := r.URL.Query().Get("city")
	accommodationIds, err := handler.getReservedAccommodationIds(from, to)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	accommodations, err := handler.getAccommodations(numOfGuests, city, accommodationIds.Id)

	for _, accommodation := range accommodations.Accommodations {
		user, _ := handler.getUserById(accommodation.HostId)
		if user.AvgRating >= 4.7 {
			accommodation.IsFeaturedHost = true
		}
	}
	response, err := json.Marshal(accommodations)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *AccommodationHandler) getReservedAccommodationIds(from string, to string) (*reservation.GetReservedAccommodationIdsResponse, error) {
	searchParams := &reservation.SearchParams{
		From: from,
		To:   to,
	}
	reservationClient := services.NewReservationClient(handler.accommodationReserveClientAddress)
	return reservationClient.GetReservedAccommodationIds(context.TODO(), &reservation.GetReservedAccommodationIdsRequest{SearchParams: searchParams})
}

func (handler *AccommodationHandler) getAccommodations(numOfGuests string, city string, ids []string) (*accommodation.GetAccommodationsByParamsResponse, error) {
	guests, err := strconv.ParseInt(numOfGuests, 10, 64)
	if err != nil {
		return nil, err
	}
	searchParams := &accommodation.SearchParams{
		NumOfGuests: guests,
		City:        city,
		Id:          ids,
	}
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	return accommodationClient.Search(context.TODO(), &accommodation.GetAccommodationsByParamsRequest{SearchParams: searchParams})
}

func (handler *AccommodationHandler) getUserById(id string) (*user_info.GetUserByIDResponse, error) {
	userClient := services.NewUserClient(handler.userInfoClientAddress)
	return userClient.GetUserByID(context.TODO(), &user_info.GetUserByIDRequest{Id: id})
}
