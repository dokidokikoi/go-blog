package store

import (
	"github.com/dokidokikoi/go-common/db/base"
	"go-blog/internal/db/model/user"
)

type Roles interface {
	base.BasicCURD[user.Role]
}
