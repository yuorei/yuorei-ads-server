package port

import (
	"context"
	"time"

	"github.com/yuorei/yuorei-ads/src/domain"
)

type AdsInputPort interface {
	GetCampaign(context.Context, string) (*domain.Campaign, error)
	CheckOrganizationID(context.Context, string, string) error
	ListCampaignByOrganizationID(context.Context, string, int, int) ([]*domain.Campaign, error)
	CreateCampaign(context.Context, *domain.Campaign) (*domain.Campaign, error)
	ListAdminAds(context.Context, string, int, int) ([]*domain.Ad, error)
	ListAdsByCampaignID(context.Context, string, int, int) ([]*domain.Ad, error)

	GetAd(context.Context, string) (*domain.Ad, error)
	CreateAdVideo(context.Context, *domain.Ad, *domain.AdVideo, string, string, string, string) (*domain.Ad, error)
	GetAdVideos(context.Context, *domain.GetAdVideoRequest) ([]*domain.AdVideoResponse, error)
	WatchCountAdVideo(context.Context, *domain.WatchCountAdVideo) error
	GetDailyWatchCountAdVideo(context.Context, string, time.Time, time.Time) (*domain.AdsViewedPerDays, error)
}

type AdsRepository interface {
	DBGetCampaign(context.Context, string) (*domain.Campaign, error)
	DBGetAd(context.Context, string) (*domain.Ad, error)
	DBCheckOrganizationID(context.Context, string, string) error
	DBListCampaignByOrganizationID(context.Context, string, int, int) ([]*domain.Campaign, error)
	DBListAdminAds(context.Context, string, int, int) ([]*domain.Ad, error)
	DBCreateCampaign(context.Context, *domain.Campaign) (*domain.Campaign, error)
	DBListAdsByCampaignID(context.Context, string, int, int) ([]*domain.Ad, error)

	DBCreateAd(context.Context, *domain.Ad) (*domain.Ad, error)
	DBCreateAdVideo(context.Context, *domain.AdVideo) (*domain.AdVideo, error)

	UploadVideoForYuoVision(context.Context, *domain.UploadVideo, string, string, string) error
	UploadThumbnailForYuoVision(context.Context, domain.ThumbnailImage, string) error

	DBGetAdVideos(context.Context, *domain.GetAdVideoRequest) ([]*domain.AdVideoResponse, error)
	BigQueryWatchCountAdVideoInsert(context.Context, *domain.WatchCountAdVideo) error
	BigQueryGetDailyWatchCountAdVideo(context.Context, string, time.Time, time.Time) (*domain.AdsViewedPerDays, error)
}
