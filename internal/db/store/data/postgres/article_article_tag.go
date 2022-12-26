package postgres

import (
	"go-blog/internal/db/model/article"
	db "go-blog/pkg/db/base"
)

type articleArticleTags struct {
	db.PgModel[article.ArticleTag]
}

func newArticleArticleTags(ds *Store) *articleArticleTags {
	return &articleArticleTags{PgModel: db.PgModel[article.ArticleTag]{DB: ds.db}}
}
