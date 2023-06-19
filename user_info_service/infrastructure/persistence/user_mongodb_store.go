package persistence

import (
	"context"
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/user_info_service/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "user"
	COLLECTION = "user"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func (u *UserMongoDBStore) SetFeaturedHost(id primitive.ObjectID, featured bool) bool {

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isFeatured": featured}}

	_, err := u.users.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return false
	}
	return true

}

func (u *UserMongoDBStore) Delete(id primitive.ObjectID) (bool, error) {
	filter := bson.M{"_id": id}

	_, err := u.users.DeleteMany(context.Background(), filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u UserMongoDBStore) CreateUser(user *model.User) error {
	user.TimesCanceled = 0

	_, err := u.users.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (u UserMongoDBStore) GetUserByID(id primitive.ObjectID) (*model.User, error) {
	filter := bson.M{"_id": id}

	usr := model.User{}
	err := u.users.FindOne(context.Background(), filter).Decode(&usr)
	if err != nil {
		return nil, err
	}
	return &usr, nil
}

func NewUserMongoDBStore(client *mongo.Client) domain.UserStore {
	users := client.Database(DATABASE).Collection(COLLECTION)
	return &UserMongoDBStore{
		users: users,
	}
}

func (store *UserMongoDBStore) filter(filter interface{}) ([]*model.User, error) {
	cursor, err := store.users.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func decode(cursor *mongo.Cursor) (products []*model.User, err error) {
	for cursor.Next(context.TODO()) {
		var product model.User
		err = cursor.Decode(&product)
		if err != nil {
			return
		}
		products = append(products, &product)
	}
	err = cursor.Err()
	return
}
