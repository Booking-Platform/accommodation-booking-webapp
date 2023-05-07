package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/infrastructure/services"
	reservation "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
	accommodation "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"strconv"
)

type AccommodationHandler struct {
	accommodationReserveClientAddress string
	accommodationClientAddress        string
}

func NewAccommodationHandler(accommodationReserveClientAddress, accommodationClientAddress string) Handler {
	return &AccommodationHandler{
		accommodationReserveClientAddress: accommodationReserveClientAddress,
		accommodationClientAddress:        accommodationClientAddress,
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
	fmt.Println(numOfGuests)
	fmt.Println(city)
	accommodationIds, err := handler.getReservedAccommodationIds(from, to)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(accommodationIds)
	stringSlice := []string{"6457a727e794f9761fcb1645", "6457a9b1eb960318f9d133a4"}
	//OVDE TREBA IZMENITI STRINGSLICE DA BUDE accommodationIds kad se popravi funkcija getReservedAccommodationIds
	accommodations, err := handler.getAccommodations(numOfGuests, city, stringSlice)

	// Marshal the slice into JSON
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
