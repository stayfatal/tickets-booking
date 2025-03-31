package testhelpers

import (
	"testing"
	"ticketsbooking/libs/config"

	"github.com/redis/go-redis/v9"
)

func PrepareRedis(t *testing.T) (*redis.Client, error) {
	db, err := config.NewRedisDB()
	if err != nil {
		return nil, err
	}

	t.Cleanup(func() {
		db.Close()
	})

	return db, nil
}
