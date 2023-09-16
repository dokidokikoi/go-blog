package store

import (
	"go-blog/internal/db/model/category"

	"github.com/dokidokikoi/go-common/db/base"
)

type Category interface {
	base.BasicCURD[category.Category]
}
