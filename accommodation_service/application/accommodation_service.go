package application

type AccommodationService struct {
}

func NewAccommodationService() *AccommodationService {
	return &AccommodationService{}

}

func (service *AccommodationService) Get() error {
	return nil
}
