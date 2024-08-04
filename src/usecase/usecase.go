package usecase

import (
	"github.com/yuorei/yuorei-ads/src/adapter/infrastructure"
	"github.com/yuorei/yuorei-ads/src/usecase/port"
)

type UseCase struct {
	port.UserInputPort
	port.AdsInputPort
}

func NewUseCase(infra *infrastructure.Infrastructure) *UseCase {
	user := NewUserUseCase(infra)
	ads := NewAdsUseCase(infra)
	return &UseCase{
		UserInputPort: user,
		AdsInputPort:  ads,
	}
}
