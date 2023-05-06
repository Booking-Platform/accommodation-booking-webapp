package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Country string             `bson:"country"`
	City    string             `bson:"city"`
	Street  string             `bson:"street"`
	Number  string             `bson:"number"`
}
