package usecase

import (
	"github.com/yuorei/yuorei-ads/src/adapter/infrastructure"
	"github.com/yuorei/yuorei-ads/src/usecase/port"
)

type UseCase struct {
	port.OrganizationInputPort
	port.UserInputPort
	port.AdsInputPort
}

type Repository struct {
	organizationRepository *OrganizationUseCase
	userRepository         *UserUseCase
	adsRepository          *AdsUseCase
}

func NewUseCase(repository *Repository) *UseCase {
	return &UseCase{
		OrganizationInputPort: repository,
		UserInputPort:         repository,
		AdsInputPort:          repository,
	}
}

func NewRepository(infra *infrastructure.Infrastructure) *Repository {
	organization := NewOrganizationRepository(infra)
	user := NewUserRepository(infra)
	ads := NewAdsRepository(infra)
	return &Repository{
		organizationRepository: organization,
		userRepository:         user,
		adsRepository:          ads,
	}
}
