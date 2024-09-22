package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	"github.com/yuorei/yuorei-ads/db/sqlc"
	"github.com/yuorei/yuorei-ads/src/domain"
	"google.golang.org/api/iterator"
)

func (i *Infrastructure) DBCreateCampaign(ctx context.Context, campaign *domain.Campaign) (*domain.Campaign, error) {
	_, err := i.db.Database.CreateCampaign(ctx,
		sqlc.CreateCampaignParams{
			CampaignID: campaign.CampaignID,
			UserID:     campaign.UserID,
			Name:       campaign.Name,
			Budget:     int32(campaign.Budget),
			StartDate:  campaign.StartDate,
			EndDate:    campaign.EndDate,
		},
	)
	if err != nil {
		return nil, err
	}
	return campaign, nil
}

func (i *Infrastructure) DBCreateAd(ctx context.Context, ad *domain.Ad) (*domain.Ad, error) {
	_, err := i.db.Database.CreateAd(ctx,
		sqlc.CreateAdParams{
			AdID:       ad.AdID,
			CampaignID: ad.CampaignID,
			AdType:     ad.AdType,
			IsApproval: sql.NullBool{
				Bool:  ad.IsApproval,
				Valid: true,
			},
			IsOpen: ad.IsOpen,
			AdLink: sql.NullString{
				String: ad.AdLink,
				Valid:  true,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return ad, nil
}

func (i *Infrastructure) DBCreateAdVideo(ctx context.Context, adVideo *domain.AdVideo) (*domain.AdVideo, error) {
	_, err := i.db.Database.CreateAdVideo(ctx,
		sqlc.CreateAdVideoParams{
			AdID:         adVideo.AdID,
			Title:        adVideo.Title,
			Description:  adVideo.Description,
			VideoUrl:     adVideo.VideoUrl,
			ThumbnailUrl: adVideo.ThumbnailUrl,
		},
	)
	if err != nil {
		return nil, err
	}
	return adVideo, nil
}

func (i *Infrastructure) DBGetAdVideos(ctx context.Context, request *domain.GetAdVideoRequest) ([]*domain.AdVideoResponse, error) {
	ads, err := i.db.Database.GetAdVideos(ctx)
	if err != nil {
		return nil, err
	}

	adVideos := make([]*domain.AdVideoResponse, 0)
	for _, ad := range ads {
		adVideos = append(adVideos, &domain.AdVideoResponse{
			AdID:         ad.AdID,
			Title:        ad.Title,
			Description:  ad.Description,
			VideoUrl:     ad.VideoUrl,
			ThumbnailUrl: ad.ThumbnailUrl,
			AdLink:       ad.AdLink.String,
		})
	}

	return adVideos, nil
}

func (i *Infrastructure) BigQueryWatchCountAdVideoInsert(ctx context.Context, request *domain.WatchCountAdVideo) error {
	datasetID := "ads_views"
	tableID := "ads_video_views"
	inserter := i.bigquery.Dataset(datasetID).Table(tableID).Inserter()

	if err := inserter.Put(ctx, request); err != nil {
		return fmt.Errorf("inserter.Put: %w", err)
	}

	return nil
}

// BigQueryGetDailyWatchCountAdVideo: 日ごとの視聴者数をBigQueryから取得
func (i *Infrastructure) BigQueryGetDailyWatchCountAdVideo(ctx context.Context, adID string) (*domain.AdsViewedPerDays, error) {
	datasetID := "ads_views"
	tableID := "ads_video_views"

	// 日ごとの視聴者数を取得するクエリ
	queryString := fmt.Sprintf(`
	SELECT
		DATE(watched_at) AS date,
		COUNT(DISTINCT user_id) AS view_count
	FROM
		%s.%s.%s
	WHERE
		ad_id = '%s'
	GROUP BY
		date
	ORDER BY
		date`, os.Getenv("GC_BQ_PROJECT_ID"), datasetID, tableID, adID)

	query := i.bigquery.Query(queryString)
	query.Parameters = []bigquery.QueryParameter{
		{
			Name:  "adID",
			Value: adID,
		},
	}

	// クエリを実行
	job, err := query.Run(ctx)
	if err != nil {
		return nil, fmt.Errorf("query.Run: %w", err)
	}

	// クエリの完了を待機
	status, err := job.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("job.Wait: %w", err)
	}
	if err := status.Err(); err != nil {
		return nil, fmt.Errorf("job.Status.Err: %w", err)
	}

	// 結果を読み込む
	it, err := job.Read(ctx)
	if err != nil {
		return nil, fmt.Errorf("job.Read: %w", err)
	}

	adsViewedPerDays := &domain.AdsViewedPerDays{
		AdID: adID,
	}

	for {
		var row struct {
			Date      civil.Date `bigquery:"date"`
			ViewCount int        `bigquery:"view_count"`
		}
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("iterator.Next: %w", err)
		}

		adsViewedPerDays.AdsViewedPerDay = append(adsViewedPerDays.AdsViewedPerDay, domain.AdsViewedPerDay{
			Day:   row.Date.String(),
			Count: row.ViewCount,
		})
	}

	return adsViewedPerDays, nil
}
