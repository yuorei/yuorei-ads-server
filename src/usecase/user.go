package usecase

import (
	"context"

	"github.com/yuorei/yuorei-ads/src/domain"
	"github.com/yuorei/yuorei-ads/src/usecase/port"
)

type UserUseCase struct {
	adsRepository port.UserRepository
}

func NewUserUseCase(adsRepository port.UserRepository) *UserUseCase {
	return &UserUseCase{
		adsRepository: adsRepository,
	}
}

func (a *UserUseCase) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	result, err := a.adsRepository.DBCreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}
