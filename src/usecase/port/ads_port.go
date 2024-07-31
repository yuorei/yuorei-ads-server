package port

import (
	"context"

	"github.com/yuorei/yuorei-ads/src/domain"
)

type AdsInputPort interface {
	CreateCampaign(context.Context, *domain.Campaign) (*domain.Campaign, error)
}

type AdsRepository interface {
	DBCreateCampaign(context.Context, *domain.Campaign) (*domain.Campaign, error)
}
