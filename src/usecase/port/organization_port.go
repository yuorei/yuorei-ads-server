package port

import (
	"context"

	"github.com/yuorei/yuorei-ads/src/domain"
)

type OrganizationInputPort interface {
	CreateOrganization(ctx context.Context, organizationID, clientID, ClientSecret, userID string) (*domain.Organization, error)
	CreateTmpOrganization(ctx context.Context, organization *domain.Organization, clientID, ClientSecret string) (*domain.Organization, error)
}

type OrganizationRepository interface {
	DBCreateOrganization(ctx context.Context, organizationID, clientID, ClientSecret, userID string) (*domain.Organization, error)
	TmpSaveRedisCreateOrganization(ctx context.Context, organization *domain.Organization, clientID, ClientSecret string) (*domain.Organization, error)
}
