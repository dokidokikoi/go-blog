package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
)

type users struct {
	cli *redis.Client
}

func (u users) SetRedisKvExpire(ctx context.Context, key, code string, expire time.Duration) error {
	return u.cli.Set(ctx, key, code, expire).Err()
}

func (u users) GetRedisKv(ctx context.Context, key string) (text string, err error) {
	return u.cli.Get(ctx, key).Result()
}
func (u users) DelRedisKv(ctx context.Context, key string) error {
	return u.cli.Del(ctx, key).Err()
}

func newUsers(cli *redis.Client) *users {
	return &users{cli: cli}
}
