package domain

import "time"

type User struct {
	ID        string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *time.Time
}

func NewUser(id string, role string, createdAt time.Time, updatedAt time.Time, deletedAt *time.Time) *User {
	return &User{
		ID:        id,
		Role:      role,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeleteAt:  deletedAt,
	}
}
