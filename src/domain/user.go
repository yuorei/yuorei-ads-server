package domain

import "time"

type User struct {
	ID        string
	Username  string
	Password  string
	Email     string
	Roles     []string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  *time.Time
}

func NewUser(id, username, password, email string, roles []string, createdAt time.Time, updatedAt time.Time, deletedAt *time.Time) *User {
	return &User{
		ID:        id,
		Username:  username,
		Password:  password,
		Email:     email,
		Roles:     roles,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeleteAt:  deletedAt,
	}
}
