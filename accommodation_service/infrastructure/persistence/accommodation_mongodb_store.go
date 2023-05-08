package persistence

import (
	"accommodation_service/domain"
	"accommodation_service/domain/model"
	"context"
	"fmt"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

const (
	DATABASE   = "accomodation"
	COLLECTION = "accomodations"
)

type AccommodationMongoDBStore struct {
	accommodations *mongo.Collection
}

func NewAccommodationMongoDBStore(client *mongo.Client) domain.AccommodationStore {
	accommodations := client.Database(DATABASE).Collection(COLLECTION)
	return &AccommodationMongoDBStore{
		accommodations: accommodations,
	}
}

func (store *AccommodationMongoDBStore) GetAllAccommodations() ([]*model.Accommodation, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *AccommodationMongoDBStore) GetAccomodationByID(id primitive.ObjectID) (*model.Accommodation, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *AccommodationMongoDBStore) Insert(accommodation *model.Accommodation) error {
	if accommodation.ID.IsZero() {
		accommodation.ID = primitive.NewObjectID()
	}

	result, err := store.accommodations.InsertOne(context.TODO(), accommodation)
	if err != nil {
		return err
	}
	accommodation.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *AccommodationMongoDBStore) AddAppointment(accommodationID primitive.ObjectID, appointment *model.Appointment) error {
	// Check if there is an overlap with an existing appointment
	filter := bson.M{
		"_id": accommodationID,
		"appointments.from": bson.M{
			"$nin": []time.Time{appointment.To},
			"$lt":  appointment.To,
		},
		"appointments.to": bson.M{
			"$nin": []time.Time{appointment.From},
			"$gt":  appointment.From,
		},
	}
	count, err := store.accommodations.CountDocuments(context.Background(), filter)
	if err != nil {
		return status.Errorf(
			codes.Internal,
			"Error counting documents: %v",
			err,
		)
	}
	if count > 0 {
		return status.Errorf(
			codes.FailedPrecondition,
			"Appointment overlaps with existing appointment",
		)
	}

	// Add the new appointment to the appointments array
	update := bson.M{
		"$push": bson.M{
			"appointments": appointment,
		},
	}
	filter = bson.M{"_id": accommodationID}
	_, err = store.accommodations.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return status.Errorf(
			codes.Internal,
			"Error updating document: %v",
			err,
		)
	}
	return nil
}

func (store *AccommodationMongoDBStore) GetAllAccommodationsByParams(searchParams *pb.SearchParams, accommodationIds []primitive.ObjectID) ([]*model.Accommodation, error) {
	if searchParams.NumOfGuests < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Number of guests must be greater than 0")
	}

	regexPattern := fmt.Sprintf(".*%s.*", searchParams.City)
	filter := bson.M{
		"_id": bson.M{
			"$nin": accommodationIds,
		},
		"address.city": bson.M{
			"$regex":   regexPattern,
			"$options": "i",
		},
	}

	if searchParams.NumOfGuests > 0 {
		filter["max_guest_num"] = bson.M{
			"$gte": searchParams.NumOfGuests,
		}
	}

	return store.filter(filter)
}

func (store *AccommodationMongoDBStore) filter(filter interface{}) ([]*model.Accommodation, error) {
	cursor, err := store.accommodations.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *AccommodationMongoDBStore) filterOne(filter interface{}) (accommodation *model.Accommodation, err error) {
	result := store.accommodations.FindOne(context.TODO(), filter)
	err = result.Decode(&accommodation)
	return
}

func decode(cursor *mongo.Cursor) (products []*model.Accommodation, err error) {
	for cursor.Next(context.TODO()) {
		var product model.Accommodation
		err = cursor.Decode(&product)
		if err != nil {
			return
		}
		products = append(products, &product)
	}
	err = cursor.Err()
	return
}
