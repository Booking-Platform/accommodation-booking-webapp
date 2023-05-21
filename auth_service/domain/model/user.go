package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
}
