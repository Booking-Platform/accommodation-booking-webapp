package application

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/domain/model"
	"google.golang.org/genproto/googleapis/type/date"
	"time"
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

	currentDate := time.Now().UTC()

	if fetchedRating == nil {
		rating.Date = date.Date{
			Year:  int32(currentDate.Year()),
			Month: int32(int(currentDate.Month())),
			Day:   int32(currentDate.Day()),
		}
		return service.store.CreateRatingForAccommodation(rating)
	} else {
		fetchedRating.Date = date.Date{
			Year:  int32(currentDate.Year()),
			Month: int32(int(currentDate.Month())),
			Day:   int32(currentDate.Day()),
		}
		return service.store.UpdateAccommodationRatingByID(fetchedRating.Id, rating.Rating)
	}
}
