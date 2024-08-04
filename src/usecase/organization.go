package usecase

import (
	"context"

	"github.com/yuorei/yuorei-ads/src/domain"
	"github.com/yuorei/yuorei-ads/src/usecase/port"
)

type OrganizationUseCase struct {
	organizationRepository port.OrganizationRepository
}

func NewOrganizationUseCase(organizationRepository port.OrganizationRepository) *OrganizationUseCase {
	return &OrganizationUseCase{
		organizationRepository: organizationRepository,
	}
}

func (a *OrganizationUseCase) CreateOrganization(ctx context.Context, organizationID, clientID, ClientSecret string) (*domain.Organization, error) {
	result, err := a.organizationRepository.DBCreateOrganization(ctx, organizationID, clientID, ClientSecret)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *OrganizationUseCase) CreateTmpOrganization(ctx context.Context, organization *domain.Organization, clientID, ClientSecret string) (*domain.Organization, error) {
	result, err := a.organizationRepository.TmpSaveRedisCreateOrganization(ctx, organization, clientID, ClientSecret)
	if err != nil {
		return nil, err
	}

	return result, nil
}
