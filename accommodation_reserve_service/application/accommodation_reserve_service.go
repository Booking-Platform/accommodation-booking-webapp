package application

import (
	"fmt"
)

type AccommodationReserveService struct {
}

func NewAccommodationReserveService() *AccommodationReserveService {
	fmt.Println("#####################################   4")

	return &AccommodationReserveService{}

}

func (service *AccommodationReserveService) Get() error {
	fmt.Println("#####################################")
	return nil
}
