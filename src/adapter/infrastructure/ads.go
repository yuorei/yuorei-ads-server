package infrastructure

import (
	"context"
	"database/sql"

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

func (i *Infrastructure) DBCreateAd(ctx context.Context, ad *domain.Ad) (*domain.Ad, error) {
	_, err := i.db.Database.CreateAd(ctx,
		sqlc.CreateAdParams{
			AdID:       ad.AdID,
			CampaignID: ad.CampaignID,
			AdType:     ad.AdType,
			IsApproval: sql.NullBool{
				Bool:  ad.IsApproval,
				Valid: true,
			},
			IsOpen: ad.IsOpen,
			AdLink: sql.NullString{
				String: ad.AdLink,
				Valid:  true,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return ad, nil
}

func (i *Infrastructure) DBCreateAdVideo(ctx context.Context, adVideo *domain.AdVideo) (*domain.AdVideo, error) {
	_, err := i.db.Database.CreateAdVideo(ctx,
		sqlc.CreateAdVideoParams{
			AdID:         adVideo.AdID,
			Title:        adVideo.Title,
			Description:  adVideo.Description,
			VideoUrl:     adVideo.VideoUrl,
			ThumbnailUrl: adVideo.ThumbnailUrl,
		},
	)
	if err != nil {
		return nil, err
	}
	return adVideo, nil
}
