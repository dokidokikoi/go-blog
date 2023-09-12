package store

import (
	"github.com/dokidokikoi/go-common/db/base"
	"go-blog/internal/db/model/article"
)

type ArticleBody interface {
	base.BasicCURD[article.ArticleBody]
}
