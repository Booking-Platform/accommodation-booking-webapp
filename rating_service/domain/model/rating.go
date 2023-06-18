package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/genproto/googleapis/type/date"
)

type HostRating struct {
	Id      primitive.ObjectID `bson:"_id"`
	GuestId primitive.ObjectID `bson:"guest_id"`
	HostId  primitive.ObjectID `bson:"host_id"`
	Rating  int                `bson:"rating"`
	Date    date.Date          `bson:"date"`
}

type AccommodationRating struct {
	Id                primitive.ObjectID `bson:"_id"`
	GuestId           primitive.ObjectID `bson:"guest_id"`
	AccommodationName string             `bson:"accommodation_name"`
	Rating            int                `bson:"rating"`
	Date              date.Date          `bson:"date"`
}
