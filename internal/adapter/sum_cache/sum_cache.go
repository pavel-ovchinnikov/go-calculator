package sum_cache

import (
	"context"

	redis "github.com/redis/go-redis/v9"
)

type SumCache struct {
	rdb *redis.Client
}

func (cache *SumCache) Set(ctx context.Context, key string, value interface{}) error {
	return cache.rdb.Set(ctx, key, value, 0).Err()
}

func (cache *SumCache) Get(ctx context.Context, key string) (interface{}, error) {
	return cache.rdb.Get(ctx, key).Result()
}

func NewSumCache() (*SumCache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       11,
	})

	return &SumCache{
		rdb: rdb,
	}, nil
}
