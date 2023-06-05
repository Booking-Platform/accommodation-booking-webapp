package domain

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingStore interface {
	CreateRatingForHost(user *model.HostRating) error
	GetRatingByUserAndGuest(userID, guestID primitive.ObjectID) (*model.HostRating, error)
	UpdateHostRatingByID(ratingID primitive.ObjectID, newRating int) error
	UpdateAccommodationRatingByID(ratingID primitive.ObjectID, newRating int) error
	CreateRatingForAccommodation(hostRating *model.AccommodationRating) error
	GetRatingByUserAndAccommodationName(accommodationName string, guestID primitive.ObjectID) (*model.AccommodationRating, error)
}
