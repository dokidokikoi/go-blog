package postgres

import (
	"go-blog/internal/db/model/link"

	db "github.com/dokidokikoi/go-common/db/base"
)

type links struct {
	db.PgModel[link.Link]
}

func newLinks(ds *Store) *links {
	return &links{PgModel: db.PgModel[link.Link]{DB: ds.db}}
}
