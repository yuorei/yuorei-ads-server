package infrastructure

import (
	"github.com/yuorei/yuorei-ads/src/driver/db"
)

type Infrastructure struct {
	db *db.DB
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{
		db: db.NewMySQLDB(),
	}
}
