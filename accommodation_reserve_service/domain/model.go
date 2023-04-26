package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/genproto/googleapis/type/date"
)

type Reservation struct {
	Id              primitive.ObjectID `bson:"_id"`
	AccommodationID primitive.ObjectID `bson:"accommodation_id"`
	GuestNum        string             `bson:"guest_num"`
	Start           date.Date          `bson:"from"`
	End             date.Date          `bson:"to"`
}
