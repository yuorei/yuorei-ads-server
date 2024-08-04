package domain

import "time"

type Organization struct {
	ID                  string
	OrganizationName    string
	RepresentativeName  string
	RepresentativeEmail string
	Purpose             string
	Category            string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeleteAt            *time.Time
}

func NewOrganization(id, organizationName, representativeName, representativeEmail, purpose, category string, createdAt time.Time, updatedAt time.Time, deletedAt *time.Time) *Organization {
	return &Organization{
		ID:                  id,
		OrganizationName:    organizationName,
		RepresentativeName:  representativeName,
		RepresentativeEmail: representativeEmail,
		Purpose:             purpose,
		Category:            category,
		CreatedAt:           createdAt,
		UpdatedAt:           updatedAt,
		DeleteAt:            deletedAt,
	}
}
