package postgres

import (
	"go-blog/internal/db/model/list"

	db "github.com/dokidokikoi/go-common/db/base"
)

type items struct {
	db.PgModel[list.Item]
}

func newItems(ds *Store) *items {
	return &items{PgModel: db.PgModel[list.Item]{DB: ds.db}}
}
