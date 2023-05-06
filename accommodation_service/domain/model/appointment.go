package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Appointment struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	From     time.Time          `bson:"from"`
	To       time.Time          `bson:"to"`
	Status   AppointmentStatus  `bson:"status"`
	Price    float64            `bson:"price"`
	PerGuest bool               `bson:"per_guest"`
}
