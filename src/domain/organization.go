package domain

import "time"

type Organization struct {
	ID                   string
	OrganizationName     string
	RepresentativeUserID string
	Purpose              string
	Category             string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeleteAt             *time.Time
}

func NewOrganization(id, organizationName, representativeUserID, purpose, category string, createdAt time.Time, updatedAt time.Time, deletedAt *time.Time) *Organization {
	return &Organization{
		ID:                   id,
		OrganizationName:     organizationName,
		RepresentativeUserID: representativeUserID,
		Purpose:              purpose,
		Category:             category,
		CreatedAt:            createdAt,
		UpdatedAt:            updatedAt,
		DeleteAt:             deletedAt,
	}
}
