package persistence

import (
	"context"
	"fmt"
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	HOST_RATINGS_COLLECTION          = "host_ratings"
	ACCOMMODATION_RATINGS_COLLECTION = "accommodation_ratings"
	DATABASE                         = "rating"
)

type RatingMongoDBStore struct {
	host_ratings          *mongo.Collection
	accommodation_ratings *mongo.Collection
}

func (store *RatingMongoDBStore) UpdateAccommodationRatingByID(ratingID primitive.ObjectID, newRating int) error {
	filter := bson.M{"_id": ratingID}
	update := bson.M{"$set": bson.M{"rating": newRating}}

	_, err := store.accommodation_ratings.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (store *RatingMongoDBStore) CreateRatingForAccommodation(accommodationRating *model.AccommodationRating) error {
	accommodationRating.Id = primitive.NewObjectID()

	_, err := store.accommodation_ratings.InsertOne(context.TODO(), accommodationRating)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (store *RatingMongoDBStore) GetRatingByUserAndAccommodationName(accommodationName string, guestID primitive.ObjectID) (*model.AccommodationRating, error) {
	filter := bson.M{
		"guest_id":           guestID,
		"accommodation_name": accommodationName,
	}

	ratings, err := store.AccommodationsRatingsFilter(filter)

	if err != nil {
		return nil, err
	}

	if len(ratings) > 0 {
		return ratings[0], nil
	}

	return nil, nil
}

func (store *RatingMongoDBStore) CreateRatingForHost(rating *model.HostRating) error {
	rating.Id = primitive.NewObjectID()

	_, err := store.host_ratings.InsertOne(context.TODO(), rating)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

func NewRatingMongoDBStore(client *mongo.Client) domain.RatingStore {
	accommodation_ratings := client.Database(DATABASE).Collection(ACCOMMODATION_RATINGS_COLLECTION)
	host_ratings := client.Database(DATABASE).Collection(HOST_RATINGS_COLLECTION)

	return &RatingMongoDBStore{
		host_ratings:          host_ratings,
		accommodation_ratings: accommodation_ratings,
	}
}

func (store *RatingMongoDBStore) HostRatingsFilter(filter interface{}) ([]*model.HostRating, error) {
	cursor, err := store.host_ratings.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *RatingMongoDBStore) AccommodationsRatingsFilter(filter interface{}) ([]*model.AccommodationRating, error) {
	cursor, err := store.accommodation_ratings.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decodeAccommodation(cursor)
}

func decodeAccommodation(cursor *mongo.Cursor) (products []*model.AccommodationRating, err error) {
	for cursor.Next(context.TODO()) {
		var rating model.AccommodationRating
		err = cursor.Decode(&rating)
		if err != nil {
			return
		}
		products = append(products, &rating)
	}
	err = cursor.Err()
	return
}

func decode(cursor *mongo.Cursor) (products []*model.HostRating, err error) {
	for cursor.Next(context.TODO()) {
		var rating model.HostRating
		err = cursor.Decode(&rating)
		if err != nil {
			return
		}
		products = append(products, &rating)
	}
	err = cursor.Err()
	return
}

func (store *RatingMongoDBStore) GetRatingByUserAndGuest(guest_id, host_id primitive.ObjectID) (*model.HostRating, error) {
	filter := bson.M{
		"guest_id": guest_id,
		"host_id":  host_id,
	}

	ratings, err := store.HostRatingsFilter(filter)

	if err != nil {
		return nil, err
	}

	if len(ratings) > 0 {
		return ratings[0], nil
	}

	return nil, nil
}

func (store *RatingMongoDBStore) UpdateHostRatingByID(ratingID primitive.ObjectID, newRating int) error {
	filter := bson.M{"_id": ratingID}
	update := bson.M{"$set": bson.M{"rating": newRating}}

	_, err := store.host_ratings.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
