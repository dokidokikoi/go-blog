package store

import (
	"go-blog/internal/db/model/article"
	"go-blog/pkg/db/base"
)

type ArticleBody interface {
	base.BasicCURD[article.ArticleBody]
}
