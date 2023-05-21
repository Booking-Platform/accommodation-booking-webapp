package api

import (
	"fmt"
	"github.com/Booking-Platform/accommodation-booking-webapp/accommodation_reserve_service/domain/model"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_reserve_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/genproto/googleapis/type/date"
	"strconv"
	"time"
)

func mapReservation(reservationPb *pb.NewReservation) (*model.Reservation, error) {

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

	accommodationID, err := primitive.ObjectIDFromHex(reservationPb.AccommodationID)
	if err != nil {
		return nil, err
	}

	reservationID, err := primitive.ObjectIDFromHex(reservationPb.UserID)
	if err != nil {
		return nil, err
	}

	guestNum, err := strconv.ParseUint(reservationPb.GuestNum, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	reservation := &model.Reservation{
		AccommodationID: accommodationID,
		Start:           startDate,
		End:             endDate,
		GuestNum:        uint(guestNum),
		UserID:          reservationID,
	}

	return reservation, nil
}

func mapReservationPb(reservation *model.Reservation) *pb.Reservation {
	reservationPb := &pb.Reservation{
		EndDate:         getDateStringForm(reservation.End),
		StartDate:       getDateStringForm(reservation.Start),
		UserID:          reservation.UserID.Hex(), // convert ObjectID to string
		GuestNum:        strconv.FormatUint(uint64(reservation.GuestNum), 10),
		AccommodationID: reservation.AccommodationID.Hex(), // convert ObjectID to string
		Status:          mapStatus(reservation.ReservationStatus),
		Id:              reservation.Id.Hex(),
	}
	return reservationPb
}

func getDateStringForm(date date.Date) string {
	return fmt.Sprintf("%d-%02d-%02d", date.Year, date.Month, date.Day)
}

func mapStatus(status model.ReservationStatus) string {
	if status == model.WAITING {
		return "WAITING"
	}
	if status == model.CONFIRMED {
		return "CONFIRMED"
	}
	if status == model.UNCONFIRMED {
		return "UNCONFIRMED"
	}
	if status == model.CANCELED {
		return "CANCELED"
	}

	return ""
}
