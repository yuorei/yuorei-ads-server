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

type AdVideoResponse struct {
	AdID         string
	Title        string
	Description  string
	VideoUrl     string
	ThumbnailUrl string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeleteAt     *time.Time
	AdLink       string
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

type GetAdVideoRequest struct {
	// ブラウザ情報
	UserAgent            string `json:"user_agent,omitempty"`
	Platform             string `json:"platform,omitempty"`
	Language             string `json:"language,omitempty"`
	URL                  string `json:"url,omitempty"`
	PageTitle            string `json:"page_title,omitempty"`
	Referrer             string `json:"referrer,omitempty"`
	NetworkDownlink      string `json:"network_downlink,omitempty"`
	NetworkEffectiveType string `json:"network_effective_type,omitempty"`
	IPAddress            string `json:"ip_address,omitempty"`
	Location             string `json:"location,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	City                 string `json:"city,omitempty"`
	Region               string `json:"region,omitempty"`
	Country              string `json:"country,omitempty"`
	Org                  string `json:"org,omitempty"`
	Postal               string `json:"postal,omitempty"`
	Timezone             string `json:"timezone,omitempty"`
	// ビデオ情報
	VideoID     string   `json:"video_id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description *string  `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	// ユーザー情報
	UserID   string `json:"user_id,omitempty"`
	ClientID string `json:"client_id,omitempty"`
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

func NewAdVideoRequest(userAgent, platform, language, url, pageTitle, referrer, networkDownlink, networkEffectiveType, ipAddress, location, hostname, city, region, country, org, postal, timezone, videoID, title, userID, clientID string, description *string, tags []string) *GetAdVideoRequest {
	return &GetAdVideoRequest{
		UserAgent:            userAgent,
		Platform:             platform,
		Language:             language,
		URL:                  url,
		PageTitle:            pageTitle,
		Referrer:             referrer,
		NetworkDownlink:      networkDownlink,
		NetworkEffectiveType: networkEffectiveType,
		IPAddress:            ipAddress,
		Location:             location,
		Hostname:             hostname,
		City:                 city,
		Region:               region,
		Country:              country,
		Org:                  org,
		Postal:               postal,
		Timezone:             timezone,
		VideoID:              videoID,
		Title:                title,
		Description:          description,
		Tags:                 tags,
		UserID:               userID,
		ClientID:             clientID,
	}
}
