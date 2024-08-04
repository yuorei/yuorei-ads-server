package presentation

import (
	"context"
	"fmt"
	"time"

	"connectrpc.com/connect"

	organizationv1 "github.com/yuorei/yuorei-ads/gen/rpc/organization/v1"
	"github.com/yuorei/yuorei-ads/src/adapter/infrastructure"
	"github.com/yuorei/yuorei-ads/src/domain"
	"github.com/yuorei/yuorei-ads/src/usecase"
)

type OrganizationServer struct {
	usecase *usecase.UseCase
}

func NewOrganizationServer(infra *infrastructure.Infrastructure) *OrganizationServer {
	return &OrganizationServer{
		usecase: usecase.NewUseCase(infra),
	}
}

func (s *OrganizationServer) CreateOrganization(ctx context.Context, req *connect.Request[organizationv1.CreateOrganizationRequest]) (*connect.Response[organizationv1.CreateOrganizationResponse], error) {
	result, err := s.usecase.CreateOrganization(ctx, req.Msg.OrganizationId, req.Msg.ClientId, req.Msg.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create organization: %w", err)
	}

	res := connect.NewResponse(&organizationv1.CreateOrganizationResponse{
		OrganizationId: result.ID,
	})

	return res, nil
}

func (s *OrganizationServer) CreateTmpOrganization(ctx context.Context, req *connect.Request[organizationv1.CreateTmpSaveOrganizationRequest]) (*connect.Response[organizationv1.CreateOrganizationResponse], error) {
	organization := domain.NewOrganization(req.Msg.OrganizationId, req.Msg.OrganizationName, req.Msg.RepresentativeName, req.Msg.RepresentativeEmail, req.Msg.Purpose, req.Msg.Category, time.Now(), time.Now(), nil)
	result, err := s.usecase.CreateTmpOrganization(ctx, organization, req.Msg.ClientId, req.Msg.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create organization: %w", err)
	}

	res := connect.NewResponse(&organizationv1.CreateOrganizationResponse{
		OrganizationId: result.ID,
	})

	return res, nil
}
