package domain

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationStore interface {
	Insert(reservation *model.Reservation) error
	GetByStatus(status model.ReservationStatus) ([]*model.Reservation, error)
	GetAllByUserID(id primitive.ObjectID) ([]*model.Reservation, error)
	ChangeReservationStatus(id primitive.ObjectID, canceled model.ReservationStatus) error
	GetReservedAccommodationsIds(from string, to string) ([]*primitive.ObjectID, error)
	GetAllReservationsThatPassed(id primitive.ObjectID) ([]*model.Reservation, error)
}
