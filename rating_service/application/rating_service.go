package application

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/domain/model"
)

type RatingService struct {
	store domain.RatingStore
}

func NewRatingService(store domain.RatingStore) *RatingService {
	return &RatingService{
		store: store,
	}
}

func (service *RatingService) RateHost(rating *model.HostRating) error {
	fetchedRating, _ := service.store.GetRatingByUserAndGuest(rating.GuestId, rating.HostId)
	if fetchedRating == nil {
		return service.store.CreateRatingForHost(rating)
	} else {
		return service.store.UpdateHostRatingByID(fetchedRating.Id, rating.Rating)
	}

}

func (service *RatingService) RateAccommodation(rating *model.AccommodationRating) error {
	fetchedRating, _ := service.store.GetRatingByUserAndAccommodationName(rating.AccommodationName, rating.GuestId)

	if fetchedRating == nil {
		return service.store.CreateRatingForAccommodation(rating)
	} else {
		return service.store.UpdateAccommodationRatingByID(fetchedRating.Id, rating.Rating)
	}
}
