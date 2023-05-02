package api

import (
	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/domain/model"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/genproto/googleapis/type/date"
	"time"
)

func mapReservation(reservationPb *pb.NewReservation) (*model.Reservation, error) {
	accommodationID, err := primitive.ObjectIDFromHex(reservationPb.AccommodationID)
	userID, err := primitive.ObjectIDFromHex(reservationPb.UserID)

	start, err := time.Parse("2006-01-02", reservationPb.StartDate)
	if err != nil {
		return nil, err
	}
	startDate := date.Date{Year: int32(start.Year()), Month: int32(start.Month()), Day: int32(start.Day())}

	end, err := time.Parse("2006-01-02", reservationPb.EndDate)
	if err != nil {
		return nil, err
	}

	endDate := date.Date{Year: int32(end.Year()), Month: int32(end.Month()), Day: int32(end.Day())}

	reservation := &model.Reservation{
		AccommodationID: accommodationID,
		Start:           startDate,
		End:             endDate,
		UserID:          userID,
	}

	return reservation, nil
}

func mapReservationPb(reservation *model.Reservation) *pb.Reservation {
	reservationPb := &pb.Reservation{
		EndDate:         reservation.End.String(),
		StartDate:       reservation.Start.String(),
		GuestNum:        string(reservation.GuestNum),
		AccommodationID: reservation.AccommodationID.String(),
		Id:              reservation.Id.String(),
	}
	return reservationPb
}
