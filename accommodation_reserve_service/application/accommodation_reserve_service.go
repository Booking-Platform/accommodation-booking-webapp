package application

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/domain"
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

func (service *AccommodationReserveService) Create(reservation *domain.Reservation) error {
	return service.store.Insert(reservation)
}
