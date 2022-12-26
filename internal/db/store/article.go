package store

import (
	"go-blog/internal/db/model/article"
	"go-blog/pkg/db/base"
)

type Article interface {
	base.BasicCURD[article.Article]
}
