package store

import (
	"go-blog/internal/db/model/article"
	"go-blog/pkg/db/base"
)

type ArticleCategory interface {
	base.BasicCURD[article.Category]
}
