package store

import (
	"github.com/dokidokikoi/go-common/db/base"
	"go-blog/internal/db/model/category"
)

type Category interface {
	base.BasicCURD[category.Category]
}
