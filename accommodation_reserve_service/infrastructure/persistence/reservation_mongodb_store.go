package persistence

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "reservation"
	COLLECTION = "reservation"
)

type ReservationMongoDBStore struct {
	reservations *mongo.Collection
}

func (r ReservationMongoDBStore) Get(id primitive.ObjectID) (*domain.Reservation, error) {
	//TODO implement me
	panic("implement me")
}

func (r ReservationMongoDBStore) Insert(product *domain.Reservation) error {
	//TODO implement me
	panic("implement me")
}

func NewReservationMongoDBStore(client *mongo.Client) domain.ReservationStore {
	reservations := client.Database(DATABASE).Collection(COLLECTION)
	return &ReservationMongoDBStore{
		reservations: reservations,
	}
}
