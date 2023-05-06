package application

import (
	"accommodation_service/domain"
	"accommodation_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationService struct {
	store domain.AccommodationStore
}

func NewAccommodationService(store domain.AccommodationStore) *AccommodationService {
	return &AccommodationService{
		store: store,
	}
}

func (service *AccommodationService) GetAccommodationByID(id primitive.ObjectID) (*model.Accommodation, error) {
	return service.store.GetAccomodationByID(id)
}

func (service *AccommodationService) GetAllAccommodations() ([]*model.Accommodation, error) {
	return service.store.GetAllAccommodations()
}

func (service *AccommodationService) Create(accommodation *model.Accommodation) error {
	accommodation.Appointments = []*model.Appointment{}
	return service.store.Insert(accommodation)
}
