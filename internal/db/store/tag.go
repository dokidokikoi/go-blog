package store

import (
	"github.com/dokidokikoi/go-common/db/base"
	"go-blog/internal/db/model/tag"
)

type Tag interface {
	base.BasicCURD[tag.Tag]
}
