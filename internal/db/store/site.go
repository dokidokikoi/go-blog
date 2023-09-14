package store

import (
	"go-blog/internal/db/model/site"

	"github.com/dokidokikoi/go-common/db/base"
)

type Sites interface {
	base.BasicCURD[site.Site]
}
