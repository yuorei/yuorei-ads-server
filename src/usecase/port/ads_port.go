package port

import (
	"context"

	"github.com/yuorei/yuorei-ads/src/domain"
)

type AdsInputPort interface {
	CreateCampaign(context.Context, *domain.Campaign) (*domain.Campaign, error)

	CreateAdVideo(context.Context, *domain.Ad, *domain.AdVideo, string, string, string, string) (*domain.Ad, error)
}

type AdsRepository interface {
	DBCreateCampaign(context.Context, *domain.Campaign) (*domain.Campaign, error)

	DBCreateAd(context.Context, *domain.Ad) (*domain.Ad, error)
	DBCreateAdVideo(context.Context, *domain.AdVideo) (*domain.AdVideo, error)

	UploadVideoForYuoVision(context.Context, *domain.UploadVideo, string, string, string) error
	UploadThumbnailForYuoVision(context.Context, domain.ThumbnailImage, string) error
}
