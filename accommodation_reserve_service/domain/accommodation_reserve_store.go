package domain

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReservationStore interface {
	Get(id primitive.ObjectID) (*model.Reservation, error)
	Insert(reservation *model.Reservation) error
	GetByStatus(status model.ReservationStatus) ([]*model.Reservation, error)
}
