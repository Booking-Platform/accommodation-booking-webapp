package domain

import (
	"accommodation_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationStore interface {
	GetAccomodationByID(id primitive.ObjectID) (*model.Accommodation, error)
	GetAllAccommodations() ([]*model.Accommodation, error)
	Insert(accommodation *model.Accommodation) error
}
