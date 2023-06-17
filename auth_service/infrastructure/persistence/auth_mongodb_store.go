package persistence

import (
	"context"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/domain"
	"github.com/Booking-Platform/accommodation-booking-webapp/auth_service/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE   = "userCred"
	COLLECTION = "user"
)

type UserMongoDBStore struct {
	users *mongo.Collection
}

func (u *UserMongoDBStore) Delete(id primitive.ObjectID) (bool, error) {
	filter := bson.M{"_id": id}

	_, err := u.users.DeleteMany(context.Background(), filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserMongoDBStore) CreateUser(user *model.User) error {
	user.Id = primitive.NewObjectID()
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	if _, err := u.users.Indexes().CreateOne(context.Background(), indexModel); err != nil {
		return err
	}

	_, err := u.users.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserMongoDBStore) Login(user *model.User) (*model.User, error) {
	filter := bson.M{
		"email":    user.Email,
		"password": user.Password,
	}

	var result model.User
	err := u.users.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
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
