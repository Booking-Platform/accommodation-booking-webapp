package api

import (
	"accommodation_service/domain/model"
	"fmt"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapNewAccommodation(accommodationPb *pb.NewAccommodation) (*model.Accommodation, error) {
	accommodation := &model.Accommodation{
		Name:        accommodationPb.Name,
		MinGuestNum: int(accommodationPb.MinGuestNum),
		MaxGuestNum: int(accommodationPb.MaxGuestNum),
		Address: model.Address{
			ID:      primitive.NewObjectID(),
			Country: accommodationPb.Address.Country,
			City:    accommodationPb.Address.City,
			Street:  accommodationPb.Address.Street,
			Number:  accommodationPb.Address.Number,
		},
		AutomaticConfirmation: accommodationPb.AutomaticConfirmation,
		Photo:                 accommodationPb.Photo,
		Benefits:              make([]*model.Benefit, 0),
	}

	for _, benefitPb := range accommodationPb.Benefits {
		benefit := &model.Benefit{
			Name: benefitPb.Name,
			ID:   primitive.NewObjectID(),
		}
		accommodation.Benefits = append(accommodation.Benefits, benefit)
	}

	return accommodation, nil
}

func mapAccommodationPb(accommodation *model.Accommodation) *pb.Accommodation {
	var pbBenefits []*pb.Benefit
	for _, benefit := range accommodation.Benefits {
		pbBenefits = append(pbBenefits, mapBenefitPb(benefit))
	}

	var pbAppointments []*pb.Appointment
	for _, appointment := range accommodation.Appointments {
		pbAppointments = append(pbAppointments, mapAppointmentPb(appointment))
	}

	accommodationPb := &pb.Accommodation{
		Id:                    accommodation.ID.Hex(),
		Name:                  accommodation.Name,
		MinGuestNum:           int32(accommodation.MinGuestNum),
		MaxGuestNum:           int32(accommodation.MaxGuestNum),
		Address:               mapAddressPb(&accommodation.Address),
		AutomaticConfirmation: accommodation.AutomaticConfirmation,
		Photo:                 accommodation.Photo,
		Benefits:              pbBenefits,
		Appointments:          pbAppointments,
	}
	return accommodationPb
}

func mapBenefitPb(benefit *model.Benefit) *pb.Benefit {
	return &pb.Benefit{
		Id:   benefit.ID.Hex(),
		Name: benefit.Name,
	}
}

func mapAppointmentPb(appointment *model.Appointment) *pb.Appointment {
	return &pb.Appointment{
		Id:       appointment.ID.Hex(),
		From:     timestamppb.New(appointment.From),
		To:       timestamppb.New(appointment.To),
		Status:   pb.AppointmentStatus(appointment.Status),
		Price:    appointment.Price,
		PerGuest: appointment.PerGuest,
	}
}

func mapAddressPb(address *model.Address) *pb.Address {
	return &pb.Address{
		Id:      address.ID.Hex(),
		Country: address.Country,
		City:    address.City,
		Street:  address.Street,
		Number:  address.Number,
	}
}

func getDateStringForm(date date.Date) string {
	return fmt.Sprintf("%d-%02d-%02d", date.Year, date.Month, date.Day)
}
