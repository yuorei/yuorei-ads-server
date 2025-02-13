package port

import (
	"context"

	"github.com/yuorei/yuorei-ads/src/domain"
)

type OrganizationInputPort interface {
	GetOrganization(ctx context.Context, organizationID string) (*domain.Organization, error)
	GetOrganizationByUserID(ctx context.Context, userID string) (*domain.Organization, error)
	CreateOrganization(ctx context.Context, organizationID, clientID, ClientSecret, userID string) (*domain.Organization, error)
	CreateTmpOrganization(ctx context.Context, organization *domain.Organization, clientID, ClientSecret string) (*domain.Organization, error)
}

type OrganizationRepository interface {
	DBGetOrganization(ctx context.Context, organizationID string) (*domain.Organization, error)
	DBGetOrganizationByUserID(ctx context.Context, userID string) (*domain.Organization, error)
	DBCreateOrganization(ctx context.Context, organizationID, clientID, ClientSecret, userID string) (*domain.Organization, error)
	TmpSaveRedisCreateOrganization(ctx context.Context, organization *domain.Organization, clientID, ClientSecret string) (*domain.Organization, error)
	DBCreateOrganizationUser(ctx context.Context, organizationID, userID string) error
}
