package application

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationReserveService struct {
	store domain.ReservationStore
}

func NewAccommodationReserveService(store domain.ReservationStore) *AccommodationReserveService {
	return &AccommodationReserveService{
		store: store,
	}
}

func (service *AccommodationReserveService) Get() error {
	panic("implement me")
}

func (service *AccommodationReserveService) Create(reservation *model.Reservation, automaticConfirmation bool) error {

	if automaticConfirmation == true {
		reservation.ReservationStatus = model.CONFIRMED
	} else {
		reservation.ReservationStatus = model.WAITING
	}

	return service.store.Insert(reservation)
}

func (service *AccommodationReserveService) GetAllForConfirmation() ([]*model.Reservation, error) {
	return service.store.GetByStatus(model.WAITING)
}

func (service *AccommodationReserveService) GetAllByUserID(id primitive.ObjectID) ([]*model.Reservation, error) {
	return service.store.GetAllByUserID(id)
}

func (service *AccommodationReserveService) ChangeReservationStatus(reservationID primitive.ObjectID, status model.ReservationStatus) error {
	return service.store.ChangeReservationStatus(reservationID, status)
}

func (service *AccommodationReserveService) GetReservedAccommodationsIds(from string, to string) ([]*primitive.ObjectID, error) {
	return service.store.GetReservedAccommodationsIds(from, to)
}

func (service *AccommodationReserveService) GetAllReservationsThatPassed(id primitive.ObjectID) ([]*model.Reservation, error) {
	return service.store.GetAllReservationsThatPassed(id)
}

func (service *AccommodationReserveService) DeleteAllUserReservations(id primitive.ObjectID) (bool, error) {
	return service.store.DeleteAllUserReservations(id)
}
