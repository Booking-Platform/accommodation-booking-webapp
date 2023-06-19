package application

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	fetchedRating, err := service.store.GetRatingByUserAndGuest(rating.GuestId, rating.HostId)

	if err != nil {
		return err
	}

	if fetchedRating == nil {
		return service.store.CreateRatingForHost(rating)
	} else {
		return service.store.UpdateHostRatingByID(fetchedRating.Id, rating.Rating)
	}
}

func (service *RatingService) RateAccommodation(rating *model.AccommodationRating) error {
	fetchedRating, err := service.store.GetRatingByUserAndAccommodationName(rating.AccommodationName, rating.GuestId)
	if err != nil {
		return err
	}

	currentDate := time.Now().UTC()

	if fetchedRating == nil {
		rating.Date = date.Date{
			Year:  int32(currentDate.Year()),
			Month: int32(currentDate.Month()),
			Day:   int32(currentDate.Day()),
		}
		return service.store.CreateRatingForAccommodation(rating)
	} else {
		fetchedRating.Date = date.Date{
			Year:  int32(currentDate.Year()),
			Month: int32(currentDate.Month()),
			Day:   int32(currentDate.Day()),
		}
		return service.store.UpdateAccommodationRatingByID(fetchedRating.Id, rating.Rating)
	}

}

func (service *RatingService) GetAccommodationRatingsByAccommodationID(accommodationName string) ([]*model.AccommodationRating, error) {
	return service.store.GetRatingsByAccommodationName(accommodationName)
}

func (service *RatingService) GetHostRatingsByHostID(id primitive.ObjectID) ([]*model.HostRating, error) {
	return service.store.GetHostRatingsByHostID(id)
}

func (service *RatingService) GetAvgRatingForHostID(hostid primitive.ObjectID) (float64, int) {

	ratings, err := service.store.GetHostRatingsByHostID(hostid)
	if err != nil {
		panic(err)
	}

	totalRatings := len(ratings)
	if totalRatings == 0 {
		return 0.0, totalRatings
	}

	var sumRatings int
	for _, rating := range ratings {
		sumRatings += rating.Rating
	}

	avgRating := float64(sumRatings) / float64(totalRatings)
	return avgRating, totalRatings
}
