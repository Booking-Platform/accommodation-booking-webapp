package persistence

import (
	"context"
	"fmt"
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/rating_service/domain/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	COLLECTION = "rating"
	DATABASE   = "rating"
)

type RatingMongoDBStore struct {
	ratings *mongo.Collection
}

func (store *RatingMongoDBStore) CreateRating(rating *model.Rating) error {
	rating.Id = primitive.NewObjectID()

	_, err := store.ratings.InsertOne(context.TODO(), rating)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

func NewRatingMongoDBStore(client *mongo.Client) domain.RatingStore {
	ratings := client.Database(DATABASE).Collection(COLLECTION)
	return &RatingMongoDBStore{
		ratings: ratings,
	}
}

func (store *RatingMongoDBStore) filter(filter interface{}) ([]*model.Rating, error) {
	cursor, err := store.ratings.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (products []*model.Rating, err error) {
	for cursor.Next(context.TODO()) {
		var rating model.Rating
		err = cursor.Decode(&rating)
		if err != nil {
			return
		}
		products = append(products, &rating)
	}
	err = cursor.Err()
	return
}
