package store

import (
	"go-blog/internal/db/model/user"
	"go-blog/pkg/db/base"
)

type Users interface {
	base.BasicCURD[user.User]
}
