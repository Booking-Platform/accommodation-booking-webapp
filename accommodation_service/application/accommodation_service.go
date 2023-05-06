package application

import (
	"fmt"
)

type AccommodationService struct {
}

func NewAccommodationService() *AccommodationService {
	fmt.Println("#####################################   4")

	return &AccommodationService{}

}

func (service *AccommodationService) Get() error {
	fmt.Println("#####################################")
	return nil
}
