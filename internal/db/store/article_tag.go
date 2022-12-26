package store

import (
	"go-blog/internal/db/model/article"
	"go-blog/pkg/db/base"
)

type ArticleTag interface {
	base.BasicCURD[article.Tag]
}
