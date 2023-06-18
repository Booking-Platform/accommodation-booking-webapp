package api

import (
	"accommodation_service/domain/model"
	pb "github.com/Booking-Platform/accommodation-booking-webapp/common/proto/accommodation_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

func mapNewAccommodation(accommodationPb *pb.NewAccommodation) (*model.Accommodation, error) {
	hostID, err := primitive.ObjectIDFromHex(accommodationPb.HostId)
	if err != nil {
		return nil, err
	}

	accommodation := &model.Accommodation{
		Name:        accommodationPb.Name,
		HostID:      hostID,
		MinGuestNum: int(accommodationPb.MinGuestNum),
		MaxGuestNum: int(accommodationPb.MaxGuestNum),
		Address: model.Address{
			ID:      primitive.NewObjectID(),
			Country: accommodationPb.Address.Country,
			City:    accommodationPb.Address.City,
			Street:  accommodationPb.Address.Street,
			Number:  int(accommodationPb.Address.Number),
		},
		AutomaticConfirmation: accommodationPb.AutomaticConfirmation,
		Photos:                accommodationPb.Photos,
		Benefits:              make([]*model.Benefit, 0),
		//Appointments:          []*model.Appointment{},
	}

	for _, benefitName := range accommodationPb.Benefits {
		benefit := &model.Benefit{
			Name: benefitName,
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
		Photo:                 accommodation.Photos,
		Benefits:              pbBenefits,
		HostId:                accommodation.HostID.Hex(),
		Appointments:          pbAppointments,
	}
	return accommodationPb
}

func mapAccommodationDTOPb(accommodation *model.Accommodation) *pb.AccommodationDTO {
	var pbBenefits []*pb.Benefit
	for _, benefit := range accommodation.Benefits {
		pbBenefits = append(pbBenefits, mapBenefitPb(benefit))
	}

	var pbAppointments []*pb.Appointment
	for _, appointment := range accommodation.Appointments {
		pbAppointments = append(pbAppointments, mapAppointmentPb(appointment))
	}

	accommodationPb := &pb.AccommodationDTO{
		Id:                    accommodation.ID.Hex(),
		Name:                  accommodation.Name,
		MinGuestNum:           strconv.Itoa(accommodation.MinGuestNum),
		MaxGuestNum:           strconv.Itoa(accommodation.MaxGuestNum),
		Address:               mapAddressPb(&accommodation.Address),
		AutomaticConfirmation: accommodation.AutomaticConfirmation,
		Photos:                accommodation.Photos,
		Benefits:              pbBenefits,
		Appointments:          pbAppointments,
		IsFeaturedHost:        false,
		HostId:                accommodation.HostID.Hex(),
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
		From:     appointment.From.Format("2006-01-02"),
		To:       appointment.To.Format("2006-01-02"),
		Status:   pb.AppointmentStatus(appointment.Status),
		Price:    appointment.Price,
		PerGuest: strconv.FormatBool(appointment.PerGuest),
	}
}

func mapPbAppointment(appointmentPb *pb.Appointment) *model.Appointment {
	from, _ := time.Parse("2006-01-02", appointmentPb.From)
	to, _ := time.Parse("2006-01-02", appointmentPb.To)

	perGuest, _ := strconv.ParseBool(appointmentPb.PerGuest)
	return &model.Appointment{
		ID:       primitive.NewObjectID(),
		From:     from,
		To:       to,
		Status:   model.AppointmentStatus(appointmentPb.Status),
		Price:    appointmentPb.Price,
		PerGuest: perGuest,
	}
}

func mapAddressPb(address *model.Address) *pb.Address {
	return &pb.Address{
		Id:      address.ID.Hex(),
		Country: address.Country,
		City:    address.City,
		Street:  address.Street,
		Number:  strconv.Itoa(address.Number),
	}
}
