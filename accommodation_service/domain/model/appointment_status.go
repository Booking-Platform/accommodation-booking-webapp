package model

type AppointmentStatus int64

const (
	AVAILABLE AppointmentStatus = 0
	RESERVED  AppointmentStatus = 1
)
