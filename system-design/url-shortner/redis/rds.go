package rds

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	Client.Ping(context.Background())
}

func CloseRedis() {
	if Client != nil {
		Client.Close()
	}
}

func Set(ctx context.Context, key, val string, ttl time.Duration) error {
	return Client.Set(ctx, key, val, ttl).Err()
}

func Get(ctx context.Context, key string) (string, error) {
	return Client.Get(ctx, key).Result()
}
