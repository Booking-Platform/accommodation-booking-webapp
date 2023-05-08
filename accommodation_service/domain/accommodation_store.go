package domain

import (
	"accommodation_service/domain/model"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationStore interface {
	GetAccomodationByID(id primitive.ObjectID) (*model.Accommodation, error)
	GetAllAccommodations() ([]*model.Accommodation, error)
	Insert(accommodation *model.Accommodation) error
	AddAppointment(accommodationID primitive.ObjectID, appointment *model.Appointment) error
	GetAllAccommodationsByParams(searchParams *pb.SearchParams, accommodationIds []primitive.ObjectID) ([]*model.Accommodation, error)
}
