package store

import (
	"github.com/dokidokikoi/go-common/db/base"
	"go-blog/internal/db/model/user"
)

type Users interface {
	base.BasicCURD[user.User]
}
