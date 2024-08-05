package usecase

import (
	"context"

	"github.com/yuorei/yuorei-ads/src/domain"
	"github.com/yuorei/yuorei-ads/src/usecase/port"
)

type AdsUseCase struct {
	adsRepository port.AdsRepository
}

func NewAdsRepository(adsRepository port.AdsRepository) *AdsUseCase {
	return &AdsUseCase{
		adsRepository: adsRepository,
	}
}

func (r *Repository) CreateCampaign(ctx context.Context, campaign *domain.Campaign) (*domain.Campaign, error) {
	result, err := r.adsRepository.adsRepository.DBCreateCampaign(ctx, campaign)
	if err != nil {
		return nil, err
	}

	return result, nil
}
