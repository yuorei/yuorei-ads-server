package usecase

import (
	"context"

	"github.com/yuorei/yuorei-ads/src/domain"
	"github.com/yuorei/yuorei-ads/src/usecase/port"
)

type AdsUseCase struct {
	adsRepository port.AdsRepository
}

func NewAdsUseCase(adsRepository port.AdsRepository) *AdsUseCase {
	return &AdsUseCase{
		adsRepository: adsRepository,
	}
}

func (a *AdsUseCase) CreateCampaign(ctx context.Context, campaign *domain.Campaign) (*domain.Campaign, error) {
	result, err := a.adsRepository.DBCreateCampaign(ctx, campaign)
	if err != nil {
		return nil, err
	}

	return result, nil
}
