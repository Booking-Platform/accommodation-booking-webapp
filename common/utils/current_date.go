package utils

import (
	"google.golang.org/genproto/googleapis/type/date"
	"time"
)

func CurrentDate() date.Date {
	current := time.Now().UTC()
	return date.Date{
		Year:  int32(current.Year()),
		Month: int32(current.Month()),
		Day:   int32(current.Day()),
	}
}
