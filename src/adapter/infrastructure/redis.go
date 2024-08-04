package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/yuorei/yuorei-ads/src/domain"
)

func (i *Infrastructure) getFromRedis(ctx context.Context, key string, data any) (bool, error) {
	bytes, err := i.redis.Get(ctx, key).Bytes()
	if err != nil {
		if err.Error() == "redis: nil" {
			return false, nil
		}

		return false, err
	}

	switch v := data.(type) {
	case *domain.Organization:
		err = json.Unmarshal(bytes, v)
		if err != nil {
			return false, err
		}

	default:
		return false, fmt.Errorf("invalid type")
	}

	return true, nil
}

func (i *Infrastructure) setToRedis(ctx context.Context, key string, expiration time.Duration, value any) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	switch v := value.(type) {
	case *domain.Organization:
		err = json.Unmarshal(bytes, &v)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("invalid type")
	}

	return i.redis.Set(ctx, key, bytes, expiration).Err()
}
