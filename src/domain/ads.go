package domain

import "time"

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
