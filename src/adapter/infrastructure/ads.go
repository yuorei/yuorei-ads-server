package infrastructure

import (
	"context"
	"database/sql"
	"fmt"

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

func (i *Infrastructure) DBGetAdVideos(ctx context.Context, request *domain.GetAdVideoRequest) ([]*domain.AdVideoResponse, error) {
	ads, err := i.db.Database.GetAdVideos(ctx)
	if err != nil {
		return nil, err
	}

	adVideos := make([]*domain.AdVideoResponse, 0)
	for _, ad := range ads {
		adVideos = append(adVideos, &domain.AdVideoResponse{
			AdID:         ad.AdID,
			Title:        ad.Title,
			Description:  ad.Description,
			VideoUrl:     ad.VideoUrl,
			ThumbnailUrl: ad.ThumbnailUrl,
			AdLink:       ad.AdLink.String,
		})
	}

	return adVideos, nil
}

func (i *Infrastructure) BigQueryWatchCountAdVideo(ctx context.Context, request *domain.WatchCountAdVideo) error {
	datasetID := "ads_views"
	tableID := "ads_video_views"
	inserter := i.bigquery.Dataset(datasetID).Table(tableID).Inserter()

	if err := inserter.Put(ctx, request); err != nil {
		return fmt.Errorf("inserter.Put: %w", err)
	}

	return nil
}
