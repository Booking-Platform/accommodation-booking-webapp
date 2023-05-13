package application

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/domain/model"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Create(user *model.User) error {
	return service.store.CreateUser(user)
}

func (service *UserService) Login(user *model.User) (*model.User, error) {
	return service.store.Login(user)
}
