package usecase

import (
	"context"

	"github.com/yuorei/yuorei-ads/src/domain"
	"github.com/yuorei/yuorei-ads/src/usecase/port"
)

type OrganizationUseCase struct {
	organizationRepository port.OrganizationRepository
}

func NewOrganizationRepository(organizationRepository port.OrganizationRepository) *OrganizationUseCase {
	return &OrganizationUseCase{
		organizationRepository: organizationRepository,
	}
}

func (u *Repository) GetOrganization(ctx context.Context, organizationID string) (*domain.Organization, error) {
	result, err := u.organizationRepository.organizationRepository.DBGetOrganization(ctx, organizationID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *Repository) GetOrganizationByUserID(ctx context.Context, userID string) (*domain.Organization, error) {
	result, err := u.organizationRepository.organizationRepository.DBGetOrganizationByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) CreateOrganization(ctx context.Context, organizationID, clientID, ClientSecret, userID string) (*domain.Organization, error) {
	result, err := r.organizationRepository.organizationRepository.DBCreateOrganization(ctx, organizationID, clientID, ClientSecret, userID)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		ID:   result.RepresentativeUserID,
		Role: "admin",
	}
	_, err = r.userRepository.userRepository.DBCreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	err = r.organizationRepository.organizationRepository.DBCreateOrganizationUser(ctx, result.ID, result.RepresentativeUserID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) CreateTmpOrganization(ctx context.Context, organization *domain.Organization, clientID, ClientSecret string) (*domain.Organization, error) {
	result, err := r.organizationRepository.organizationRepository.TmpSaveRedisCreateOrganization(ctx, organization, clientID, ClientSecret)
	if err != nil {
		return nil, err
	}

	return result, nil
}
