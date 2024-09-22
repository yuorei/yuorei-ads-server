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
	VideoID          string   `json:"video_id,omitempty"`
	VideoTitle       string   `json:"video_title,omitempty"`
	VideoDescription *string  `json:"video_description,omitempty"`
	VideoTags        []string `json:"video_tags,omitempty"`
	// ユーザー情報
	UserID   string `json:"user_id,omitempty"`
	ClientID string `json:"client_id,omitempty"`
}

type WatchCountAdVideo struct {
	// ブラウザ情報
	UserAgent            string `json:"user_agent,omitempty" bigquery:"user_agent"`
	Platform             string `json:"platform,omitempty" bigquery:"platform"`
	Language             string `json:"language,omitempty" bigquery:"language"`
	Url                  string `json:"url,omitempty" bigquery:"url"`
	PageTitle            string `json:"page_title,omitempty" bigquery:"page_title"`
	Referrer             string `json:"referrer,omitempty" bigquery:"referrer"`
	NetworkDownlink      string `json:"network_downlink,omitempty" bigquery:"network_downlink"`
	NetworkEffectiveType string `json:"network_effective_type,omitempty" bigquery:"network_effective_type"`
	IpAddress            string `json:"ip_address,omitempty" bigquery:"ip_address"`
	Location             string `json:"location,omitempty" bigquery:"location"`
	Hostname             string `json:"hostname,omitempty" bigquery:"hostname"`
	City                 string `json:"city,omitempty" bigquery:"city"`
	Region               string `json:"region,omitempty" bigquery:"region"`
	Country              string `json:"country,omitempty" bigquery:"country"`
	Org                  string `json:"org,omitempty" bigquery:"org"`
	Postal               string `json:"postal,omitempty" bigquery:"postal"`
	Timezone             string `json:"timezone,omitempty" bigquery:"timezone"`
	// ビデオ情報
	VideoId          string   `json:"video_id,omitempty" bigquery:"video_id"`
	VideoTitle       string   `json:"video_title,omitempty" bigquery:"video_title"`
	VideoDescription string   `json:"video_description,omitempty" bigquery:"video_description"`
	VideoTags        []string `json:"video_tags,omitempty" bigquery:"video_tags"`
	// ユーザー情報
	UserId   string `json:"user_id,omitempty" bigquery:"user_id"`
	ClientId string `json:"client_id,omitempty" bigquery:"client_id"`
	// 広告情報
	AdId string `json:"ad_id,omitempty" bigquery:"ad_id"`
	// 時間情報
	WatchedAt time.Time `json:"watched_at,omitempty" bigquery:"watched_at"`
}

type AdsViewedPerDays struct {
	AdID            string            `json:"ad_id"`
	AdsViewedPerDay []AdsViewedPerDay `json:"ads_viewed_per_day,omitempty"`
}

type AdsViewedPerDay struct {
	Day   string `json:"day"`
	Count int    `json:"count"`
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
		VideoTitle:           title,
		VideoDescription:     description,
		VideoTags:            tags,
		UserID:               userID,
		ClientID:             clientID,
	}
}
