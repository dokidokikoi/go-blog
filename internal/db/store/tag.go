package store

import (
	"go-blog/internal/db/model/tag"

	"github.com/dokidokikoi/go-common/db/base"
)

type Tag interface {
	base.BasicCURD[tag.Tag]
}
