package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Accommodation struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	Name                  string             `bson:"name"`
	MinGuestNum           int                `bson:"min_guest_num" validate:"gte=1,lte=10"`
	MaxGuestNum           int                `bson:"max_guest_num" validate:"gte=1,lte=10"`
	Address               Address            `bson:"address"`
	AutomaticConfirmation bool               `bson:"automatic_confirmation"`
	Photos                []string           `bson:"photos"`
	Benefits              []*Benefit         `bson:"benefits,omitempty"`
	Appointments          []*Appointment     `bson:"appointments,omitempty"`
}
