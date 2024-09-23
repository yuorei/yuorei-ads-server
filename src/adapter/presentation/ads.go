package presentation

import (
	"context"
	"fmt"
	"time"

	"connectrpc.com/connect"

	adsv1 "github.com/yuorei/yuorei-ads-proto/gen/rpc/ads/v1"
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
	userID, ok := ctx.Value("uid").(string)
	if !ok || userID == "" {
		return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("failed to get userID"))
	}

	// Goのデフォルトの形式
	const LAYOUT = "2006-01-02"
	startDate, err := time.Parse(LAYOUT, req.Msg.StartDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start date: %w", err)
	}

	endDate, err := time.Parse(LAYOUT, req.Msg.EndDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end date: %w", err)
	}

	campaign := domain.NewCampaign("", userID, req.Msg.Name, int(req.Msg.Budget), startDate, endDate, false, time.Now(), time.Now(), nil)
	result, err := s.usecase.CreateCampaign(ctx, campaign)
	if err != nil {
		return nil, fmt.Errorf("failed to create campaign: %w", err)
	}

	res := connect.NewResponse(&adsv1.CreateCampaignResponse{
		CampaignId: result.CampaignID,
	})

	return res, nil
}

func (s *AdsServer) GetAdVideo(ctx context.Context, req *connect.Request[adsv1.GetAdVideoRequest]) (*connect.Response[adsv1.GetAdVideoResponseList], error) {
	adVideoRequest := domain.NewAdVideoRequest(req.Msg.UserAgent, req.Msg.Platform, req.Msg.Language, req.Msg.Url, req.Msg.PageTitle, req.Msg.Referrer, req.Msg.NetworkDownlink, req.Msg.NetworkEffectiveType, req.Msg.IpAddress, req.Msg.Location, req.Msg.Hostname, req.Msg.City, req.Msg.Region, req.Msg.Country, req.Msg.Org, req.Msg.Postal, req.Msg.Timezone, req.Msg.VideoId, req.Msg.VideoTitle, req.Msg.UserId, req.Msg.ClientId, &req.Msg.VideoDescription, req.Msg.VideoTags)
	adVideos, err := s.usecase.AdsInputPort.GetAdVideos(ctx, adVideoRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to get ad video: %w", err)
	}

	getAdVideoResponse := make([]*adsv1.GetAdVideoResponse, 0)
	for _, ad := range adVideos {
		getAdVideoResponse = append(getAdVideoResponse, &adsv1.GetAdVideoResponse{
			AdId:         ad.AdID,
			Title:        ad.Title,
			Description:  ad.Description,
			AdUrl:        ad.AdLink,
			ThumbnailUrl: ad.ThumbnailUrl,
			VideoUrl:     ad.VideoUrl,
		})
	}

	res := connect.NewResponse(&adsv1.GetAdVideoResponseList{
		Responses: getAdVideoResponse,
	})
	return res, nil
}

func (s *AdsServer) WatchCountAdVideo(ctx context.Context, req *connect.Request[adsv1.WatchCountAdVideoRequest]) (*connect.Response[adsv1.WatchCountAdVideoResponse], error) {
	watchCountAdVideoRequest := &domain.WatchCountAdVideo{
		UserAgent:            req.Msg.UserAgent,
		Platform:             req.Msg.Platform,
		Language:             req.Msg.Language,
		Url:                  req.Msg.Url,
		PageTitle:            req.Msg.PageTitle,
		Referrer:             req.Msg.Referrer,
		NetworkDownlink:      req.Msg.NetworkDownlink,
		NetworkEffectiveType: req.Msg.NetworkEffectiveType,
		IpAddress:            req.Msg.IpAddress,
		Location:             req.Msg.Location,
		Hostname:             req.Msg.Hostname,
		City:                 req.Msg.City,
		Region:               req.Msg.Region,
		Country:              req.Msg.Country,
		Org:                  req.Msg.Org,
		Postal:               req.Msg.Postal,
		Timezone:             req.Msg.Timezone,
		VideoId:              req.Msg.VideoId,
		VideoTitle:           req.Msg.VideoTitle,
		VideoDescription:     req.Msg.VideoDescription,
		VideoTags:            req.Msg.VideoTags,
		UserId:               req.Msg.UserId,
		ClientId:             req.Msg.ClientId,
		AdId:                 req.Msg.AdId,
		WatchedAt:            domain.NowJST(),
	}
	err := s.usecase.WatchCountAdVideo(ctx, watchCountAdVideoRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to watch count ad video: %w", err)
	}

	res := connect.NewResponse(&adsv1.WatchCountAdVideoResponse{})
	return res, nil
}

func (s *AdsServer) GetDailyWatchCountAdVideo(ctx context.Context, req *connect.Request[adsv1.AdsViewedPerDaysRequest]) (*connect.Response[adsv1.AdsViewedPerDaysResponse], error) {
	userID, ok := ctx.Value("uid").(string)
	if !ok || userID == "" {
		return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("failed to get userID"))
	}
	// TODO: このユーザーが組織にAdIDを持っているかどうかのチェックが必要

	start := req.Msg.Start.AsTime()
	end := req.Msg.End.AsTime()
	result, err := s.usecase.GetDailyWatchCountAdVideo(ctx, req.Msg.AdId, start, end)
	if err != nil {
		return nil, fmt.Errorf("failed to get daily watch count ad video: %w", err)
	}

	var adsViewedPerDay []*adsv1.AdsViewedPerDay
	for _, ad := range result.AdsViewedPerDay {
		adsViewedPerDay = append(adsViewedPerDay, &adsv1.AdsViewedPerDay{
			Day:   ad.Day,
			Count: int32(ad.Count),
		})
	}

	res := connect.NewResponse(&adsv1.AdsViewedPerDaysResponse{
		AdId:            result.AdID,
		AdsViewedPerDay: adsViewedPerDay,
	})
	return res, nil
}
