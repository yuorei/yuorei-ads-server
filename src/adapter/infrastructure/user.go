package infrastructure

import (
	"context"

	"github.com/yuorei/yuorei-ads/db/sqlc"
	"github.com/yuorei/yuorei-ads/src/domain"
)

func (i *Infrastructure) DBCreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	_, err := i.db.Database.CreateUser(ctx,
		sqlc.CreateUserParams{
			UserID: user.ID,
			Role:   user.Role,
		})
	if err != nil {
		return nil, err
	}
	return user, nil
}
