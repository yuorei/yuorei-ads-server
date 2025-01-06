package presentation

import (
	"context"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/timestamppb"

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

func (s *AdsServer) ListCampaignByOrganizationID(ctx context.Context, req *connect.Request[adsv1.ListCampaignByOrganizationIDRequest]) (*connect.Response[adsv1.ListCampaignByOrganizationIDResponse], error) {
	userID, ok := ctx.Value("uid").(string)
	if !ok || userID == "" {
		return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("failed to get userID"))
	}
	err := s.usecase.CheckOrganizationID(ctx, req.Msg.OrganizationId, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to check organizationID: %w", err)
	}

	campaigns, err := s.usecase.ListCampaignByOrganizationID(ctx, req.Msg.OrganizationId, int(req.Msg.Offset), int(req.Msg.Limit))
	if err != nil {
		return nil, fmt.Errorf("failed to list campaign by organizationID: %w", err)
	}

	var campaignList []*adsv1.Campaign
	for _, campaign := range campaigns {
		campaignList = append(campaignList, &adsv1.Campaign{
			CampaignId: campaign.CampaignID,
			Name:       campaign.Name,
			Budget:     int32(campaign.Budget),
			StartDate:  campaign.StartDate.Format("2006-01-02"),
			EndDate:    campaign.EndDate.Format("2006-01-02"),
		})
	}

	res := connect.NewResponse(&adsv1.ListCampaignByOrganizationIDResponse{
		Campaigns: campaignList,
	})
	return res, nil
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

func (s *AdsServer) GetAd(ctx context.Context, req *connect.Request[adsv1.GetAdRequest]) (*connect.Response[adsv1.GetAdResponse], error) {
	ad, err := s.usecase.GetAd(ctx, req.Msg.AdId)
	if err != nil {
		return nil, fmt.Errorf("failed to create ad video: %w", err)
	}

	var deleteAt *timestamppb.Timestamp
	if ad.DeleteAt == nil {
		deleteAt = nil
	} else {
		deleteAt = timestamppb.New(*ad.DeleteAt)
	}

	res := connect.NewResponse(&adsv1.GetAdResponse{
		Ad: &adsv1.Ad{
			AdId:       ad.AdID,
			CampaignId: ad.CampaignID,
			AdType:     ad.AdType,
			CreatedAt:  timestamppb.New(ad.CreatedAt),
			UpdatedAt:  timestamppb.New(ad.UpdatedAt),
			DeletedAt:  deleteAt,
			IsApproval: ad.IsApproval,
			IsOpen:     ad.IsOpen,
			AdLink:     ad.AdLink,
		},
	})
	return res, nil

}

func (s *AdsServer) ListAdminAds(ctx context.Context, req *connect.Request[adsv1.ListAdminAdsRequest]) (*connect.Response[adsv1.ListAdminAdsResponse], error) {
	userID, ok := ctx.Value("uid").(string)
	if !ok || userID == "" {
		return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("failed to get userID"))
	}

	ads, err := s.usecase.ListAdminAds(ctx, userID, int(req.Msg.Offset), int(req.Msg.Limit))
	if err != nil {
		return nil, fmt.Errorf("failed to list ads: %w", err)
	}

	var adList []*adsv1.Ad
	for _, ad := range ads {
		var deleteAt *timestamppb.Timestamp
		if ad.DeleteAt == nil {
			deleteAt = nil
		} else {
			deleteAt = timestamppb.New(*ad.DeleteAt)
		}
		adList = append(adList, &adsv1.Ad{
			AdId:       ad.AdID,
			CampaignId: ad.CampaignID,
			AdType:     ad.AdType,
			CreatedAt:  timestamppb.New(ad.CreatedAt),
			UpdatedAt:  timestamppb.New(ad.UpdatedAt),
			DeletedAt:  deleteAt,
			IsApproval: ad.IsApproval,
			IsOpen:     ad.IsOpen,
			AdLink:     ad.AdLink,
		})
	}

	res := connect.NewResponse(&adsv1.ListAdminAdsResponse{
		Ads: adList,
	})
	return res, nil
}

func (s *AdsServer) ListAdsByCampaignID(ctx context.Context, req *connect.Request[adsv1.ListAdsByCampaignIDRequest]) (*connect.Response[adsv1.ListAdsByCampaignIDResponse], error) {
	userID, ok := ctx.Value("uid").(string)
	if !ok || userID == "" {
		return nil, connect.NewError(connect.CodePermissionDenied, fmt.Errorf("failed to get userID"))
	}

	ads, err := s.usecase.ListAdsByCampaignID(ctx, req.Msg.CampaignId, int(req.Msg.Offset), int(req.Msg.Limit))
	if err != nil {
		return nil, fmt.Errorf("failed to list ads by campaignID: %w", err)
	}

	var adList []*adsv1.Ad
	for _, ad := range ads {
		var deleteAt *timestamppb.Timestamp
		if ad.DeleteAt == nil {
			deleteAt = nil
		} else {
			deleteAt = timestamppb.New(*ad.DeleteAt)
		}
		adList = append(adList, &adsv1.Ad{
			AdId:       ad.AdID,
			CampaignId: ad.CampaignID,
			AdType:     ad.AdType,
			CreatedAt:  timestamppb.New(ad.CreatedAt),
			UpdatedAt:  timestamppb.New(ad.UpdatedAt),
			DeletedAt:  deleteAt,
			IsApproval: ad.IsApproval,
			IsOpen:     ad.IsOpen,
			AdLink:     ad.AdLink,
		})
	}

	res := connect.NewResponse(&adsv1.ListAdsByCampaignIDResponse{
		Ads: adList,
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
