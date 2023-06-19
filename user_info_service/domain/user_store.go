package domain

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	GetUserByID(id primitive.ObjectID) (*model.User, error)
	CreateUser(user *model.User) error
	Delete(id primitive.ObjectID) (bool, error)
	SetFeaturedHost(id primitive.ObjectID, isFeatured bool) bool
}
