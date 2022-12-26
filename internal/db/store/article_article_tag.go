package store

import (
	"go-blog/internal/db/model/article"
	"go-blog/pkg/db/base"
)

type ArticleArticleTag interface {
	base.BasicCURD[article.ArticleTag]
}
