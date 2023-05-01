package persistence

import (
	"context"
	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/domain/model"
	"go.mongodb.org/mongo-driver/bson"
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

func (store ReservationMongoDBStore) Get(id primitive.ObjectID) (*model.Reservation, error) {
	//TODO implement me
	panic("implement me")
}

func (store *ReservationMongoDBStore) Insert(reservation *model.Reservation) error {
	if reservation.Id.IsZero() {
		reservation.Id = primitive.NewObjectID()
	}

	result, err := store.reservations.InsertOne(context.TODO(), reservation)
	if err != nil {
		return err
	}
	reservation.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *ReservationMongoDBStore) GetByStatus(status model.ReservationStatus) ([]*model.Reservation, error) {
	filter := bson.M{"reservation_status": status}
	return store.filter(filter)
}

func NewReservationMongoDBStore(client *mongo.Client) domain.ReservationStore {
	reservations := client.Database(DATABASE).Collection(COLLECTION)
	return &ReservationMongoDBStore{
		reservations: reservations,
	}
}

func (store *ReservationMongoDBStore) filter(filter interface{}) ([]*model.Reservation, error) {
	cursor, err := store.reservations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (products []*model.Reservation, err error) {
	for cursor.Next(context.TODO()) {
		var product model.Reservation
		err = cursor.Decode(&product)
		if err != nil {
			return
		}
		products = append(products, &product)
	}
	err = cursor.Err()
	return
}
