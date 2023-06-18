package persistence

import (
	"context"
	"github.com/Booking-Platform/accommodation-booking-webapp/common/utils"
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
	hostRatings          *mongo.Collection
	accommodationRatings *mongo.Collection
}

func (store *RatingMongoDBStore) GetHostRatingsByHostID(id primitive.ObjectID) ([]*model.HostRating, error) {
	filter := bson.M{"host_id": id}
	return store.hostRatingsFilter(filter)

}

func (store *RatingMongoDBStore) UpdateAccommodationRatingByID(ratingID primitive.ObjectID, newRating int) error {
	filter := bson.M{"_id": ratingID}
	update := bson.M{"$set": bson.M{"rating": newRating}}

	_, err := store.accommodationRatings.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (store *RatingMongoDBStore) UpdateHostRatingByID(ratingID primitive.ObjectID, newRating int) error {
	filter := bson.M{"_id": ratingID}
	update := bson.M{"$set": bson.M{"rating": newRating}}

	_, err := store.hostRatings.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (store *RatingMongoDBStore) CreateRatingForAccommodation(accommodationRating *model.AccommodationRating) error {
	accommodationRating.Id = primitive.NewObjectID()
	accommodationRating.Date = utils.CurrentDate()

	_, err := store.accommodationRatings.InsertOne(context.TODO(), accommodationRating)
	if err != nil {
		return err
	}

	return nil
}

func (store *RatingMongoDBStore) CreateRatingForHost(rating *model.HostRating) error {
	rating.Id = primitive.NewObjectID()
	rating.Date = utils.CurrentDate()

	_, err := store.hostRatings.InsertOne(context.TODO(), rating)
	if err != nil {
		return err
	}

	return nil
}

func (store *RatingMongoDBStore) GetRatingByUserAndAccommodationName(accommodationName string, guestID primitive.ObjectID) (*model.AccommodationRating, error) {
	filter := bson.M{
		"guest_id":           guestID,
		"accommodation_name": accommodationName,
	}

	ratings, err := store.accommodationRatingsFilter(filter)
	if err != nil {
		return nil, err
	}

	if len(ratings) > 0 {
		return ratings[0], nil
	}

	return nil, nil
}

func NewRatingMongoDBStore(client *mongo.Client) domain.RatingStore {
	hostRatings := client.Database(DATABASE).Collection(HOST_RATINGS_COLLECTION)
	accommodationRatings := client.Database(DATABASE).Collection(ACCOMMODATION_RATINGS_COLLECTION)

	return &RatingMongoDBStore{
		hostRatings:          hostRatings,
		accommodationRatings: accommodationRatings,
	}
}

func (store *RatingMongoDBStore) hostRatingsFilter(filter interface{}) ([]*model.HostRating, error) {
	cursor, err := store.hostRatings.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	return decodeHostRatings(cursor)
}

func (store *RatingMongoDBStore) accommodationRatingsFilter(filter interface{}) ([]*model.AccommodationRating, error) {
	cursor, err := store.accommodationRatings.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	return decodeAccommodationRatings(cursor)
}

func decodeAccommodationRatings(cursor *mongo.Cursor) ([]*model.AccommodationRating, error) {
	var ratings []*model.AccommodationRating
	for cursor.Next(context.TODO()) {
		var rating model.AccommodationRating
		err := cursor.Decode(&rating)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, &rating)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return ratings, nil
}

func decodeHostRatings(cursor *mongo.Cursor) ([]*model.HostRating, error) {
	var ratings []*model.HostRating
	for cursor.Next(context.TODO()) {
		var rating model.HostRating
		err := cursor.Decode(&rating)
		if err != nil {
			return nil, err
		}
		ratings = append(ratings, &rating)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return ratings, nil
}

func (store *RatingMongoDBStore) GetRatingByUserAndGuest(guestID, hostID primitive.ObjectID) (*model.HostRating, error) {
	filter := bson.M{
		"guest_id": guestID,
		"host_id":  hostID,
	}

	ratings, err := store.hostRatingsFilter(filter)
	if err != nil {
		return nil, err
	}

	if len(ratings) > 0 {
		return ratings[0], nil
	}

	return nil, nil
}

func (store *RatingMongoDBStore) GetRatingsByAccommodationName(accommodationName string) ([]*model.AccommodationRating, error) {
	filter := bson.M{
		"accommodation_name": accommodationName,
	}

	ratings, err := store.accommodationRatingsFilter(filter)
	if err != nil {
		return nil, err
	}

	return ratings, nil
}
