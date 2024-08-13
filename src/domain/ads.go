package domain

import (
	"time"
)

type Campaign struct {
	CampaignID string
	UserID     string
	Name       string
	Budget     int
	StartDate  time.Time
	EndDate    time.Time
	IsApproval bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteAt   *time.Time
}

type Ad struct {
	AdID       string
	CampaignID string
	AdType     string
	IsApproval bool
	IsOpen     bool
	AdLink     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteAt   *time.Time
	Tags       []string
}

type AdVideo struct {
	AdID         string
	Title        string
	Description  string
	VideoUrl     string
	ThumbnailUrl string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeleteAt     *time.Time
}

type Video struct {
	ID                string
	Title             string
	Description       *string
	ThumbnailImageUrl string
	UserId            string
	Tags              []string
	Private           bool
	Adult             bool
	ExternalCutout    bool
	IsAd              bool
}

type UploadVideo struct {
	ID               string
	VideoContentType string

	ImageContentType string
	Tags             []string
	Title            string
	Description      *string
	IsPrivate        bool
	IsAdult          bool
	IsExternalCutout bool
	IsAd             bool
}

type ThumbnailImage struct {
	ID          string
	ContentType string
}

func NewCampaign(campaignID, userID, name string, budget int, startDate, endDate time.Time, isApproval bool, createdAt time.Time, updatedAt time.Time, deletedAt *time.Time) *Campaign {
	return &Campaign{
		CampaignID: campaignID,
		UserID:     userID,
		Name:       name,
		Budget:     budget,
		StartDate:  startDate,
		EndDate:    endDate,
		IsApproval: isApproval,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		DeleteAt:   deletedAt,
	}
}

func NewAd(adiD, campaignID, adType string, isApproval, isOpen bool, adLink string, tags []string, createdAt, updatedAt time.Time, deletedAt *time.Time) *Ad {
	return &Ad{
		AdID:       adiD,
		CampaignID: campaignID,
		AdType:     adType,
		IsApproval: isApproval,
		IsOpen:     isOpen,
		AdLink:     adLink,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		DeleteAt:   deletedAt,
		Tags:       tags,
	}
}

func NewAdVideo(adID, title, description, videoUrl, thumbnailUrl string, createdAt, updatedAt time.Time, deletedAt *time.Time) *AdVideo {
	return &AdVideo{
		AdID:         adID,
		Title:        title,
		Description:  description,
		VideoUrl:     videoUrl,
		ThumbnailUrl: thumbnailUrl,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		DeleteAt:     deletedAt,
	}
}
