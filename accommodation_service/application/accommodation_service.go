package application

import (
	"accommodation_service/domain"
	"accommodation_service/domain/model"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"
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
	accommodation, err := service.store.GetAccomodationByID(id)
	if err != nil {
		return nil, err
	}
	return accommodation, nil
}

func (service *AccommodationService) GetAllAccommodations() ([]*model.Accommodation, error) {
	accommodations, err := service.store.GetAllAccommodations()
	if err != nil {
		return nil, err
	}
	return accommodations, nil
}

func (service *AccommodationService) Create(accommodation *model.Accommodation) error {
	err := service.store.Insert(accommodation)
	if err != nil {
		return err
	}
	return nil
}

func (service *AccommodationService) AddAppointment(accommodation *model.Accommodation, appointment *model.Appointment) error {
	err := service.store.AddAppointment(accommodation.ID, appointment)
	if err != nil {
		return err
	}
	return nil
}

func (service *AccommodationService) GetAllAccommodationsByParams(searchParams *pb.SearchParams, reservedAccommodationIds []string) ([]*model.Accommodation, error) {
	accommodationIds := make([]primitive.ObjectID, len(reservedAccommodationIds))
	for i, s := range reservedAccommodationIds {
		objectID, err := primitive.ObjectIDFromHex(s)
		if err != nil {
			return nil, err
		}
		accommodationIds[i] = objectID
	}
	accommodations, err := service.store.GetAllAccommodationsByParams(searchParams, accommodationIds)
	if err != nil {
		return nil, err
	}
	return accommodations, nil
}

func (service *AccommodationService) ChangeAutomaticConfirmationStatus(id primitive.ObjectID) error {
	return service.store.ChangeAutomaticConfirmationStatusByAccommodationID(id)
}
