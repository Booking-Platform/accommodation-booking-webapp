package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id            primitive.ObjectID `bson:"_id"`
	Name          string             `bson:"name"`
	Surname       string             `bson:"surname"`
	TimesCanceled int                `bson:"timesCanceled"`
	Is            int                `bson:"timesCanceled"`
}
