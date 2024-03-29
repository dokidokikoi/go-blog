package postgres

import (
	"go-blog/internal/db/model/article"

	db "github.com/dokidokikoi/go-common/db/base"
)

type articles struct {
	db.PgModel[article.Article]
}

func newArticles(ds *Store) *articles {
	return &articles{PgModel: db.PgModel[article.Article]{DB: ds.db}}
}
