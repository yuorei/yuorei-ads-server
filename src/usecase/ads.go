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
