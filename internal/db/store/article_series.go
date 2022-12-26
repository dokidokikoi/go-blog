package store

import (
	"go-blog/internal/db/model/article"
	"go-blog/pkg/db/base"
)

type ArticleSeries interface {
	base.BasicCURD[article.Series]
}
