package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/infrastructure/services"
	reservation "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
	accommodation "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"
	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"strconv"
	"strings"
	"time"
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
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := handler.validateToken(tokenString)
	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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

func (handler *AccommodationHandler) validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte("yE54RkqqgahuNbMCPmlrqbcoDeUXadi4ibXdsKjssQTwSl3FZaJoNyvc05553OGA"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
		if time.Now().UTC().After(expirationTime) {
			return nil, fmt.Errorf("token has expired")
		}
	}

	return token, nil
}
