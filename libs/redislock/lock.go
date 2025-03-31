package redislock

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

var (
	retriesAmount = 5
	retrySleep    = time.Second
	lockDuration  = time.Second * 5
)

type Locker interface {
	Lock(key string, lockId string) error
	Unlock(key string, lockId string) error
}

type locker struct {
	db *redis.Client
}

func NewLocker(db *redis.Client) Locker {
	return &locker{db}
}

func (l *locker) Lock(key string, lockId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var ok bool
	for i := 0; i < retriesAmount; i++ {
		var err error
		ok, err = l.db.SetNX(ctx, key, lockId, lockDuration).Result()
		if err != nil {
			return errors.Wrap(err, "locking via setnx")
		}
		if ok {
			break
		}
		time.Sleep(retrySleep)
	}

	if !ok {
		return errors.New("run out of retries")
	}

	return nil
}

func (l *locker) Unlock(key string, lockId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result, err := l.db.Get(ctx, key).Result()
	if err != nil {
		return errors.Wrap(err, "getting lock val from redis")
	}

	if result != lockId {
		return errors.New("identifiers dont match")
	}

	err = l.db.Del(ctx, key).Err()
	if err != nil {
		return errors.Wrap(err, "deleting lock from redis")
	}

	return nil
}
