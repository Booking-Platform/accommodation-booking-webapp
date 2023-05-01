package model

type ReservationStatus int64

const (
	CONFIRMED   ReservationStatus = 0
	WAITING     ReservationStatus = 1
	UNCONFIRMED ReservationStatus = 2
)
