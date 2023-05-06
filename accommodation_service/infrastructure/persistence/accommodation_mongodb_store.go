package persistence

import (
	"context"
	"accommodation_service/domain"
	"accommodation_service/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "accommodation"
	COLLECTION = "reservation"
)

type AccommodationMongoDBStore struct {
	accommodations *mongo.Collection
}

func NewAccommodationMongoDBStore(client *mongo.Client) domain. {
	reservations := client.Database(DATABASE).Collection(COLLECTION)
	return &AccommodationMongoDBStore{
		accommodations: accommodations,
	}
}

func (store AccommodationMongoDBStore) GetAllAccommodationsID(id primitive.ObjectID) ([]*model.Accommodation, error) {
	filter := bson.M{"user_id": id}
	return store.filter(filter)
}

//func (store *ReservationMongoDBStore) Insert(reservation *model.Reservation) error {
//	if reservation.Id.IsZero() {
//		reservation.Id = primitive.NewObjectID()
//	}
//
//	result, err := store.reservations.InsertOne(context.TODO(), reservation)
//	if err != nil {
//		return err
//	}
//	reservation.Id = result.InsertedID.(primitive.ObjectID)
//	return nil
//}
//
//func (store *ReservationMongoDBStore) GetByStatus(status model.ReservationStatus) ([]*model.Reservation, error) {
//	filter := bson.M{"reservation_status": status}
//	return store.filter(filter)
//}
//
//func (store *ReservationMongoDBStore) filter(filter interface{}) ([]*model.Reservation, error) {
//	cursor, err := store.reservations.Find(context.TODO(), filter)
//	defer cursor.Close(context.TODO())
//
//	if err != nil {
//		return nil, err
//	}
//	return decode(cursor)
//}
//
//func decode(cursor *mongo.Cursor) (products []*model.Reservation, err error) {
//	for cursor.Next(context.TODO()) {
//		var product model.Reservation
//		err = cursor.Decode(&product)
//		if err != nil {
//			return
//		}
//		products = append(products, &product)
//	}
//	err = cursor.Err()
//	return
//}
