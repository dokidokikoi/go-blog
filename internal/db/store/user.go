package store

import (
	"context"
	"go-blog/internal/db/model/user"
	"time"

	"github.com/dokidokikoi/go-common/db/base"
)

type Users interface {
	base.BasicCURD[user.User]
	SetRedisKvExpire(ctx context.Context, key, code string, expire time.Duration) error
	GetRedisKv(ctx context.Context, key string) (text string, err error)
	DelRedisKv(ctx context.Context, key string) error
}
