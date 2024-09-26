package cache

import "github.com/go-redis/redis/v8"

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(address string, password string, db int) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	return &RedisCache{client: client}
}
