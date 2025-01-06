package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/yuorei/yuorei-ads/db/sqlc"
	"github.com/yuorei/yuorei-ads/src/domain"
)

func (i *Infrastructure) DBCreateOrganization(ctx context.Context, organizationID, clientID, ClientSecret, userID string) (*domain.Organization, error) {
	organization := &domain.Organization{}
	hit, err := i.getFromRedis(ctx, clientID+"_"+ClientSecret, organization)
	if err != nil {
		return nil, err
	} else if !hit {
		return nil, fmt.Errorf("organization not found")
	} else if organizationID != organization.ID {
		return nil, fmt.Errorf("organization not found")
	}

	organization.RepresentativeUserID = userID
	_, err = i.db.Database.CreateOrganization(ctx,
		sqlc.CreateOrganizationParams{
			OrganizationID:       organization.ID,
			OrganizationName:     organization.OrganizationName,
			RepresentativeUserID: organization.RepresentativeUserID,
			Purpose:              organization.Purpose,
			Category:             organization.Category,
		},
	)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (i *Infrastructure) TmpSaveRedisCreateOrganization(ctx context.Context, organization *domain.Organization, clientID, ClientSecret string) (*domain.Organization, error) {
	err := i.setToRedis(ctx, clientID+"_"+ClientSecret, 24*time.Hour, organization)
	if err != nil {
		return nil, err
	}
	return organization, nil
}

func (i *Infrastructure) DBCreateOrganizationUser(ctx context.Context, organizationID, userID string) error {
	_, err := i.db.Database.CreateOrganizationUser(ctx,
		sqlc.CreateOrganizationUserParams{
			OrganizationID: organizationID,
			UserID:         userID,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
