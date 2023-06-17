package api

import (
	"context"
	"fmt"
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/infrastructure/services"
	rating "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/rating_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type RatingHandler struct {
	accommodationReserveClientAddress string
	userInfoClientAddress             string
	accommodationClientAddress        string
	ratingClientAddress               string
}

func NewRatingHandler(ratingClientAddress, accommodationReserveClientAddress, userInfoClientAddress, accommodationClientAddress string) *RatingHandler {
	return &RatingHandler{
		accommodationReserveClientAddress: accommodationReserveClientAddress,
		userInfoClientAddress:             userInfoClientAddress,
		accommodationClientAddress:        accommodationClientAddress,
		ratingClientAddress:               ratingClientAddress,
	}
}

func (handler *RatingHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/api/rating/getAllRatingsForAccommodation/{id}", handler.GetAllRatingsForAccommodation)

	if err != nil {
		panic(err)
	}

}

func (handler *RatingHandler) GetAllRatingsForAccommodation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]

	handler.getRatingsByAccommodationID(id)
	if id == "0" {
		fmt.Println("ss")
	}
}

func (handler *RatingHandler) getRatingsByAccommodationID(id string) (*rating.RatingsResponse, error) {
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)

	return ratingClient.GetAccommodationRatingsByAccommodationID(context.TODO(), &rating.IdMessageRequest{Id: id})
}
