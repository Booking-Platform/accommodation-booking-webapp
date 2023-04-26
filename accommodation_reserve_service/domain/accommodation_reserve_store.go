package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReservationStore interface {
	Get(id primitive.ObjectID) (*Reservation, error)
	Insert(product *Reservation) error
}
