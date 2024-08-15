package infrastructure

import (
	"github.com/redis/go-redis/v9"
	"github.com/yuorei/yuorei-ads/src/driver/client"
	"github.com/yuorei/yuorei-ads/src/driver/db"
	r "github.com/yuorei/yuorei-ads/src/driver/redis"
)

type Infrastructure struct {
	db        *db.DB
	redis     *redis.Client
	yuovision *client.ClientYuoVision
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{
		db:        db.NewMySQLDB(),
		redis:     r.ConnectRedis(),
		yuovision: client.NewClientYuoVision(),
	}
}
