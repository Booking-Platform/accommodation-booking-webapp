package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rating struct {
	Id      primitive.ObjectID `bson:"_id"`
	GuestId primitive.ObjectID `bson:"guest_id"`
	HostId  primitive.ObjectID `bson:"host_id"`
	Rating  int                `bson:"rating"`
}
