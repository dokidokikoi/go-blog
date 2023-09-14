package postgres

import (
	"go-blog/internal/db/model/site"

	db "github.com/dokidokikoi/go-common/db/base"
)

type sites struct {
	db.PgModel[site.Site]
}

func newSites(ds *Store) *sites {
	return &sites{PgModel: db.PgModel[site.Site]{DB: ds.db}}
}
