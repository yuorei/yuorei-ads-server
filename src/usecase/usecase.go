package usecase

import (
	"github.com/yuorei/yuorei-ads/src/adapter/infrastructure"
	"github.com/yuorei/yuorei-ads/src/usecase/port"
)

type UseCase struct {
	port.AdsInputPort
}

func NewUseCase(infra *infrastructure.Infrastructure) *UseCase {
	ads := NewAdsUseCase(infra)
	return &UseCase{
		AdsInputPort: ads,
	}
}
