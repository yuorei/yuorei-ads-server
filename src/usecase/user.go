package usecase

import (
	"context"

	"github.com/yuorei/yuorei-ads/src/domain"
	"github.com/yuorei/yuorei-ads/src/usecase/port"
)

type UserUseCase struct {
	userRepository port.UserRepository
}

func NewUserRepository(repository port.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: repository,
	}
}

func (r *Repository) CreateUser(ctx context.Context, userID, role string) (*domain.User, error) {
	user := &domain.User{
		ID:   userID,
		Role: role,
	}
	result, err := r.userRepository.userRepository.DBCreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}
