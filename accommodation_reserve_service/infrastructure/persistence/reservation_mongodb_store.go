package persistence

import (
	"context"
	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/genproto/googleapis/type/date"
	"time"
)

const (
	DATABASE   = "reservation"
	COLLECTION = "reservation"
)

type ReservationMongoDBStore struct {
	reservations *mongo.Collection
}

func (store *ReservationMongoDBStore) ChangeReservationStatus(id primitive.ObjectID, status model.ReservationStatus) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"reservation_status": status}}

	_, err := store.reservations.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func NewReservationMongoDBStore(client *mongo.Client) domain.ReservationStore {
	reservations := client.Database(DATABASE).Collection(COLLECTION)
	return &ReservationMongoDBStore{
		reservations: reservations,
	}
}

func (store *ReservationMongoDBStore) GetAllByUserID(id primitive.ObjectID) ([]*model.Reservation, error) {
	filter := bson.M{"user_id": id}
	return store.filter(filter)
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

func (store *ReservationMongoDBStore) GetReservedAccommodationsIds(from string, to string) ([]*primitive.ObjectID, error) {
	fromDate, err := time.Parse("2006-01-02", from)
	if err != nil {
		return nil, err
	}
	toDate, err := time.Parse("2006-01-02", to)
	if err != nil {
		return nil, err
	}

	fromParsedDate := date.Date{Year: int32(fromDate.Year()), Month: int32(fromDate.Month()), Day: int32(fromDate.Day())}
	toParsedDate := date.Date{Year: int32(toDate.Year()), Month: int32(toDate.Month()), Day: int32(toDate.Day())}

	filter := bson.M{
		"$or": []bson.M{
			bson.M{
				"$and": []bson.M{
					bson.M{"start": bson.M{"$lte": fromParsedDate}},
					bson.M{"end": bson.M{"$gte": fromParsedDate}},
				},
			},
			bson.M{
				"$and": []bson.M{
					bson.M{"start": bson.M{"$lte": toParsedDate}},
					bson.M{"end": bson.M{"$gte": toParsedDate}},
				},
			},
			bson.M{
				"$and": []bson.M{
					bson.M{"start": bson.M{"$gte": fromParsedDate}},
					bson.M{"end": bson.M{"$lte": toParsedDate}},
				},
			},
		},
		"reservation_status": bson.M{
			"$eq": 0,
		},
	}
	// definiramo projekciju koja će vratiti samo accommodation_id polje
	projection := bson.M{"accommodation_id": 1}

	// filtriramo rezervacije koristeći filter i projekciju
	cursor, err := store.reservations.Find(context.Background(), filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// iteriramo kroz sve pronađene dokumente i izvlačimo accommodation_id vrijednosti
	var accommodationIds []*primitive.ObjectID
	for cursor.Next(context.Background()) {
		var reservation model.Reservation
		if err := cursor.Decode(&reservation); err != nil {
			return nil, err
		}

		accommodationIds = append(accommodationIds, &reservation.AccommodationID)
	}
	return accommodationIds, nil
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
