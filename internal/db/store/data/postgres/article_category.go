package postgres

import (
	"go-blog/internal/db/model/article"
	db "go-blog/pkg/db/base"
)

type articleCategories struct {
	db.PgModel[article.Category]
}

func newArticleCategories(ds *Store) *articleCategories {
	return &articleCategories{PgModel: db.PgModel[article.Category]{DB: ds.db}}
}
