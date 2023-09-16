package postgres

import (
	"go-blog/internal/db/model/category"

	db "github.com/dokidokikoi/go-common/db/base"
)

type categories struct {
	db.PgModel[category.Category]
}

func newCategories(ds *Store) *categories {
	return &categories{PgModel: db.PgModel[category.Category]{DB: ds.db}}
}
