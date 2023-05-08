package application

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) GetByID(id primitive.ObjectID) (*model.User, error) {
	return service.store.GetUserByID(id)
}
