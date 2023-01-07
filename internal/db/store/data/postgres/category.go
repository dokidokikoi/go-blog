package postgres

import (
	"go-blog/internal/db/model/category"
	db "go-blog/pkg/db/base"
)

type categories struct {
	db.PgModel[category.Category]
}

func newArticleCategories(ds *Store) *categories {
	return &categories{PgModel: db.PgModel[category.Category]{DB: ds.db}}
}
