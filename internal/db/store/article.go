package store

import (
	"github.com/dokidokikoi/go-common/db/base"
	"go-blog/internal/db/model/article"
)

type Article interface {
	base.BasicCURD[article.Article]
}
