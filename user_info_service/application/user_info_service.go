package application

type UserInfoService struct{}

func NewUserInfoService() *UserInfoService {
	return &UserInfoService{}
}

func (service *UserInfoService) Get() error {
	return nil
}