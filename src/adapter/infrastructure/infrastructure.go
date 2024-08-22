package infrastructure

import (
	"cloud.google.com/go/bigquery"
	"github.com/redis/go-redis/v9"
	"github.com/yuorei/yuorei-ads/src/driver/client"
	"github.com/yuorei/yuorei-ads/src/driver/db"
	"github.com/yuorei/yuorei-ads/src/driver/googlecloud"
	r "github.com/yuorei/yuorei-ads/src/driver/redis"
)

type Infrastructure struct {
	db        *db.DB
	redis     *redis.Client
	yuovision *client.ClientYuoVision
	bigquery  *bigquery.Client
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{
		db:        db.NewMySQLDB(),
		redis:     r.ConnectRedis(),
		yuovision: client.NewClientYuoVision(),
		bigquery:  googlecloud.NewBigQuery(),
	}
}
