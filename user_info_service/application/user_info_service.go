package application

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/domain"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}
