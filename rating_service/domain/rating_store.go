package domain

import "github.com/Booking-Platform/accommodation-booking-webapp/rating_service/domain/model"

type RatingStore interface {
	CreateRating(user *model.Rating) error
}
