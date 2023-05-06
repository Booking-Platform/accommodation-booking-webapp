package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Benefit struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}
