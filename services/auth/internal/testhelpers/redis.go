package testhelpers

import (
	"booking/services/auth/config"
	"testing"

	"github.com/redis/go-redis/v9"
)

func PrepareRedis(t *testing.T) (*redis.Client, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	db, err := config.NewRedisDb(cfg)
	if err != nil {
		return nil, err
	}

	t.Cleanup(func() {
		db.Close()
	})

	return db, nil
}
