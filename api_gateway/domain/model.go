package domain

import (
	"google.golang.org/genproto/googleapis/type/date"
)

type Reservations struct {
	Reservations []Reservation
}

type Reservation struct {
	Start         date.Date
	End           date.Date
	Accommodation Accommodation
	GuestNum      int
	User          User
}

type User struct {
	Name    string
	Surname string
	Email   string
}

type Accommodation struct {
	Name    string
	Id      string
	Address string
}
