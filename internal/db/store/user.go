package store

import (
	"context"
	"go-blog/internal/db/model/user"

	"github.com/dokidokikoi/go-common/db/base"
)

type Users interface {
	base.BasicCURD[user.User]
	SetCaptchCode(ctx context.Context, key, code string) error
	GetCaptchCode(ctx context.Context, key string) (text string, err error)
	DelCaptchCode(ctx context.Context, key string) error
}
