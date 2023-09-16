package store

import (
	"go-blog/internal/db/model/comment"

	"github.com/dokidokikoi/go-common/db/base"
)

type Comments interface {
	base.BasicCURD[comment.Comment]
}
