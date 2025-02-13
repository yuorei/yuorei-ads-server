package usecase

import (
	"context"
	"time"

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

func (r *Repository) GetCampaign(ctx context.Context, campaignID string) (*domain.Campaign, error) {
	result, err := r.adsRepository.adsRepository.DBGetCampaign(ctx, campaignID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) CheckOrganizationID(ctx context.Context, organizationID, userID string) error {
	err := r.adsRepository.adsRepository.DBCheckOrganizationID(ctx, organizationID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ListCampaignByOrganizationID(ctx context.Context, organizationID string, offset, limit int) ([]*domain.Campaign, error) {
	result, err := r.adsRepository.adsRepository.DBListCampaignByOrganizationID(ctx, organizationID, offset, limit)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) ListAdminAds(ctx context.Context, userID string, offset, limit int) ([]*domain.Ad, error) {
	result, err := r.adsRepository.adsRepository.DBListAdminAds(ctx, userID, offset, limit)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) ListAdsByCampaignID(ctx context.Context, campaignID string, offset, limit int) ([]*domain.Ad, error) {
	result, err := r.adsRepository.adsRepository.DBListAdsByCampaignID(ctx, campaignID, offset, limit)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) CreateCampaign(ctx context.Context, campaign *domain.Campaign) (*domain.Campaign, error) {
	result, err := r.adsRepository.adsRepository.DBCreateCampaign(ctx, campaign)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) CreateAdVideo(ctx context.Context, ad *domain.Ad, adVideo *domain.AdVideo, userID, uploadID, videoType, imageType string) (*domain.Ad, error) {
	adResult, err := r.adsRepository.adsRepository.DBCreateAd(ctx, ad)
	if err != nil {
		return nil, err
	}

	uploadVideo := &domain.UploadVideo{
		ID:               ad.AdID,
		VideoContentType: videoType,
		Title:            adVideo.Title,
		Description:      &adVideo.Description,
		Tags:             ad.Tags,
		IsPrivate:        false,
		IsAdult:          false,
		IsExternalCutout: false,
		IsAd:             true,
		ImageContentType: imageType,
	}

	err = r.adsRepository.adsRepository.UploadVideoForYuoVision(ctx, uploadVideo, userID, uploadID, videoType)
	if err != nil {
		return nil, err
	}

	image := domain.ThumbnailImage{
		ID:          ad.AdID,
		ContentType: imageType,
	}
	err = r.adsRepository.adsRepository.UploadThumbnailForYuoVision(ctx, image, uploadID)
	if err != nil {
		return nil, err
	}

	_, err = r.adsRepository.adsRepository.DBCreateAdVideo(ctx, adVideo)
	if err != nil {
		return nil, err
	}

	return adResult, nil
}

func (r *Repository) GetAd(ctx context.Context, adID string) (*domain.Ad, error) {
	result, err := r.adsRepository.adsRepository.DBGetAd(ctx, adID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) GetAdVideos(ctx context.Context, req *domain.GetAdVideoRequest) ([]*domain.AdVideoResponse, error) {
	result, err := r.adsRepository.adsRepository.DBGetAdVideos(ctx, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) WatchCountAdVideo(ctx context.Context, req *domain.WatchCountAdVideo) error {
	err := r.adsRepository.adsRepository.BigQueryWatchCountAdVideoInsert(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetDailyWatchCountAdVideo(ctx context.Context, adID string, start, end time.Time) (*domain.AdsViewedPerDays, error) {
	result, err := r.adsRepository.adsRepository.BigQueryGetDailyWatchCountAdVideo(ctx, adID, start, end)
	if err != nil {
		return nil, err
	}

	return result, nil
}
