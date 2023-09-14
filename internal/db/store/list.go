package store

import (
	"go-blog/internal/db/model/list"

	"github.com/dokidokikoi/go-common/db/base"
)

type Items interface {
	base.BasicCURD[list.Item]
}
