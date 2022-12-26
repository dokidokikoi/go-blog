package postgres

import (
	"go-blog/internal/db/model/article"
	db "go-blog/pkg/db/base"
)

type articleBodys struct {
	db.PgModel[article.ArticleBody]
}

func newArticleBodys(ds *Store) *articleBodys {
	return &articleBodys{PgModel: db.PgModel[article.ArticleBody]{DB: ds.db}}
}
