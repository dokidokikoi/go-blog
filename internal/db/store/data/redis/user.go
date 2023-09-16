package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
)

type users struct {
	cli *redis.Client
}

func (u users) SetCaptchCode(ctx context.Context, key, code string) error {
	return u.cli.Set(ctx, key, code, time.Minute).Err()
}

func (u users) GetCaptchCode(ctx context.Context, key string) (text string, err error) {
	return u.cli.Get(ctx, key).Result()
}

func (u users) DelCaptchCode(ctx context.Context, key string) error {
	return u.cli.Del(ctx, key).Err()
}

func newUsers(cli *redis.Client) *users {
	return &users{cli: cli}
}
