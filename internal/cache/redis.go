package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
	TTL    time.Duration
}

func RedisNew(redisAddr string, redisPassword string, redisDB int, ttl time.Duration) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})
	return &RedisCache{
		client: rdb,
		ctx:    context.Background(),
		TTL:    ttl,
	}
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

func (r *RedisCache) Set(key string, value string, expiration time.Duration) error {
	return r.client.Set(r.ctx, key, value, expiration).Err()
}
