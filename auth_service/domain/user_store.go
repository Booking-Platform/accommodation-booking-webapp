package domain

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	CreateUser(user *model.User) error
	Login(user *model.User) (*model.User, error)
	Delete(id primitive.ObjectID) (bool, error)
}
