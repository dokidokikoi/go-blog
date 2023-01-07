package store

import (
	"go-blog/internal/db/model/user"
	"go-blog/pkg/db/base"
)

type Roles interface {
	base.BasicCURD[user.Role]
}
