// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"database/sql"
	"time"
)

type Ad struct {
	AdID       string
	CampaignID string
	AdType     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime
	IsApproval sql.NullBool
	IsOpen     bool
	AdLink     sql.NullString
}

type AdImage struct {
	AdID        string
	Title       string
	Description string
	ImageUrl    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

type AdVideo struct {
	AdID         string
	Title        string
	Description  string
	ThumbnailUrl string
	VideoUrl     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
}

type Campaign struct {
	CampaignID string
	UserID     string
	Name       string
	Budget     int32
	StartDate  time.Time
	EndDate    time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime
	IsApproval sql.NullBool
}

type Impression struct {
	ImpressionID string
	AdID         string
	Date         time.Time
	Impressions  int32
	Clicks       int32
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
}

type Organization struct {
	OrganizationID       string
	OrganizationName     string
	RepresentativeUserID string
	Purpose              string
	Category             string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            sql.NullTime
}

type OrganizationsUser struct {
	OrganizationID string
	UserID         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime
}

type Targeting struct {
	TargetingID string
	AdID        string
	Type        string
	Value       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

type User struct {
	UserID    string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
