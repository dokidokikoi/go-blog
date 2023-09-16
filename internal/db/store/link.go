package store

import (
	"go-blog/internal/db/model/link"

	"github.com/dokidokikoi/go-common/db/base"
)

type Link interface {
	base.BasicCURD[link.Link]
}
