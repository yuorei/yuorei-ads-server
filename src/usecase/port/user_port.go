package port

import (
	"context"

	"github.com/yuorei/yuorei-ads/src/domain"
)

type UserInputPort interface {
	CreateUser(context.Context, string, string) (*domain.User, error)
}

type UserRepository interface {
	DBCreateUser(context.Context, *domain.User) (*domain.User, error)
}
