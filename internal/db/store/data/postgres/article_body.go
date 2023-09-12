package postgres

import (
	db "github.com/dokidokikoi/go-common/db/base"
	"go-blog/internal/db/model/article"
)

type articleBodys struct {
	db.PgModel[article.ArticleBody]
}

func newArticleBodys(ds *Store) *articleBodys {
	return &articleBodys{PgModel: db.PgModel[article.ArticleBody]{DB: ds.db}}
}
