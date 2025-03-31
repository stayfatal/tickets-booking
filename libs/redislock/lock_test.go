package redislock

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestLockUnlock(t *testing.T) {
	db, err := NewRedisDb()
	if err != nil {
		t.Fatal(err)
	}

	locker := NewLocker(db)

	expected := []string{"first", "second"}
	got := []string{}
	key := "key"

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		ident := uuid.NewString()

		time.Sleep(time.Second)
		err := locker.Lock(key, ident)
		if err != nil {
			t.Error(err)
		}

		got = append(got, "second")

		err = locker.Unlock(key, ident)
		if err != nil {
			t.Error(err)
		}
	}()

	ident := uuid.NewString()

	err = locker.Lock(key, ident)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second * 2)

	got = append(got, "first")

	err = locker.Unlock(key, ident)
	if err != nil {
		t.Fatal(err)
	}

	wg.Wait()

	assert.Equal(t, expected, got)
}

func NewRedisDb() (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "mypass",
		DB:       6,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := db.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return db, nil
}
