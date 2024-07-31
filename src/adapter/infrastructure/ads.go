package infrastructure

import (
	"context"

	"github.com/yuorei/yuorei-ads/db/sqlc"
	"github.com/yuorei/yuorei-ads/src/domain"
)

func (i *Infrastructure) DBCreateCampaign(ctx context.Context, campaign *domain.Campaign) (*domain.Campaign, error) {
	_, err := i.db.Database.CreateCampaign(ctx,
		sqlc.CreateCampaignParams{
			CampaignID: campaign.CampaignID,
			UserID:     campaign.UserID,
			Name:       campaign.Name,
			Budget:     int32(campaign.Budget),
			StartDate:  campaign.StartDate,
			EndDate:    campaign.EndDate,
		},
	)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}
