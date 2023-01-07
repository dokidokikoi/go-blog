package store

import (
	"go-blog/internal/db/model/category"
	"go-blog/pkg/db/base"
)

type Category interface {
	base.BasicCURD[category.Category]
}
