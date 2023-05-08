package domain

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/domain/model"
)

type UserStore interface {
	CreateUser(user *model.User) error
}
