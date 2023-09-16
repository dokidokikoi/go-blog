package postgres

import (
	"go-blog/internal/db/model/site"

	db "github.com/dokidokikoi/go-common/db/base"
)

type siteTags struct {
	db.PgModel[site.SiteTag]
}

func newSiteTags(ds *Store) *siteTags {
	return &siteTags{PgModel: db.PgModel[site.SiteTag]{DB: ds.db}}
}
