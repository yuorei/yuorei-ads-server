package presentation

import (
	"context"
	"fmt"
	"time"

	"connectrpc.com/connect"

	adsv1 "github.com/yuorei/yuorei-ads/gen/rpc/ads/v1"
	"github.com/yuorei/yuorei-ads/src/domain"
	"github.com/yuorei/yuorei-ads/src/usecase"
)

type AdsServer struct {
	usecase *usecase.UseCase
}

func NewAdsServer(repository *usecase.Repository) *AdsServer {
	return &AdsServer{
		usecase: usecase.NewUseCase(repository),
	}
}

func (s *AdsServer) CreateCampaign(ctx context.Context, req *connect.Request[adsv1.CreateCampaignRequest]) (*connect.Response[adsv1.CreateCampaignResponse], error) {
	// TODO: 認証後のユーザIDを取得
	userID := "1"
	// TODO:変換	req.Msg.StartDate, req.Msg.EndDate を time.Time に変換して代入する
	startDate := time.Now()
	endDate := time.Now()
	campaign := domain.NewCampaign("id", userID, req.Msg.Name, int(req.Msg.Budget), startDate, endDate, false, time.Now(), time.Now(), nil)
	result, err := s.usecase.CreateCampaign(ctx, campaign)
	if err != nil {
		return nil, fmt.Errorf("failed to create campaign: %w", err)
	}

	res := connect.NewResponse(&adsv1.CreateCampaignResponse{
		CampaignId: result.CampaignID,
	})

	return res, nil
}
