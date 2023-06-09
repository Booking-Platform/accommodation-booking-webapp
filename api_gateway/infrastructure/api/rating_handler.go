package api

import (
	"context"
	"encoding/json"
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/api_gateway/infrastructure/services"
	rating "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/rating_service"
	user_info "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/user_info_service"
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
	err := mux.HandlePath("GET", "/api/rating/getAllRatingsForAccommodation/{accommodationName}", handler.GetAllRatingsForAccommodation)
	err = mux.HandlePath("GET", "/api/rating/getRatingsForHost/{hostID}", handler.GetRatingsForHost)
	err = mux.HandlePath("POST", "/api/rating/host", handler.RateHost)

	if err != nil {
		panic(err)
	}

}

func (handler *RatingHandler) RateHost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	var newRating domain.NewRating
	err := json.NewDecoder(r.Body).Decode(&newRating)
	if err != nil {
		panic(err)
	}

	hostAvgRating, numOfRatings := handler.createHostRating(newRating)

	if hostAvgRating >= 4.7 && numOfRatings > 4 {
		handler.setFeaturedHost(newRating.HostID, true)
	} else {
		handler.setFeaturedHost(newRating.HostID, false)

	}

}

func (handler *RatingHandler) GetAllRatingsForAccommodation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	accommodationName := pathParams["accommodationName"]

	ratings, err := handler.getRatingsByAccommodationName(accommodationName)

	if err != nil {
		panic(err)
	}

	response := make([]*domain.AccommodationRating, len(ratings.GetRatings()))

	for i, res := range ratings.GetRatings() {

		user, err := handler.getUserById(res.GuestID)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		response[i] = &domain.AccommodationRating{
			Rating:            res.Rating,
			Date:              res.Date,
			AccommodationName: res.AccommodationName,
			User: domain.User{
				Name:    user.Name,
				Surname: user.Surname,
			},
		}
	}

	writeJsonResponse(w, http.StatusOK, response)

}

func (handler *RatingHandler) GetRatingsForHost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	hostID := pathParams["hostID"]

	ratings, err := handler.getRatingsByHostID(hostID)

	if err != nil {
		panic(err)
	}

	response := make([]*domain.HostRating, len(ratings.GetRatings()))

	for i, res := range ratings.GetRatings() {

		guest, err := handler.getUserById(res.GuestID)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		host, err := handler.getUserById(res.HostID)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		response[i] = &domain.HostRating{
			Rating: res.Rating,
			Date:   res.Date,
			Guest: domain.User{
				Name:    guest.Name,
				Surname: guest.Surname,
			},
			Host: domain.User{
				Name:    host.Name,
				Surname: host.Surname,
			},
		}
	}

	writeJsonResponse(w, http.StatusOK, response)

}

func (handler *RatingHandler) createHostRating(newRating domain.NewRating) (float64, uint32) {
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)

	createRatingRequest := &rating.CreateRatingRequest{
		NewRating: &rating.NewRating{
			Rating:  newRating.Rating,
			HostID:  newRating.HostID,
			GuestID: newRating.GuestID,
		},
	}

	response, err := ratingClient.CreateRating(context.TODO(), createRatingRequest)
	if err != nil {
		panic(err)
	}

	hostAvgRating := response.AvgRating
	numOfRatings := response.NumOfRatings

	return hostAvgRating, numOfRatings
}

func (handler *RatingHandler) getRatingsByAccommodationName(accommodationName string) (*rating.AccommodationRatingsResponse, error) {
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)

	return ratingClient.GetAccommodationRatingsByAccommodationID(context.TODO(), &rating.IdMessageRequest{Id: accommodationName})
}

func (handler *RatingHandler) getRatingsByHostID(hostID string) (*rating.HostRatingsResponse, error) {
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)

	return ratingClient.GetHostRatingsByHostID(context.TODO(), &rating.IdMessageRequest{Id: hostID})
}

func (handler *RatingHandler) getUserById(id string) (*user_info.GetUserByIDResponse, error) {
	userClient := services.NewUserClient(handler.userInfoClientAddress)
	return userClient.GetUserByID(context.TODO(), &user_info.GetUserByIDRequest{Id: id})
}

func (handler *RatingHandler) setFeaturedHost(id string, isFeatured bool) (*user_info.BlankResponse, error) {
	userClient := services.NewUserClient(handler.userInfoClientAddress)
	return userClient.SetFeaturedHost(context.TODO(), &user_info.SetFeaturedHostRequest{Id: id, Featured: isFeatured})
}
