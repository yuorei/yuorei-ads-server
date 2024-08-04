package usecase

import (
	"github.com/yuorei/yuorei-ads/src/adapter/infrastructure"
	"github.com/yuorei/yuorei-ads/src/usecase/port"
)

type UseCase struct {
	port.OrganizationInputPort
	port.UserInputPort
	port.AdsInputPort
}

func NewUseCase(infra *infrastructure.Infrastructure) *UseCase {
	organization := NewOrganizationUseCase(infra)
	user := NewUserUseCase(infra)
	ads := NewAdsUseCase(infra)
	return &UseCase{
		OrganizationInputPort: organization,
		UserInputPort:         user,
		AdsInputPort:          ads,
	}
}
