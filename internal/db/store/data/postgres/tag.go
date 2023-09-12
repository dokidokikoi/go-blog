package postgres

import (
	db "github.com/dokidokikoi/go-common/db/base"
	"go-blog/internal/db/model/tag"
)

type tags struct {
	db.PgModel[tag.Tag]
}

func newTags(ds *Store) *tags {
	return &tags{PgModel: db.PgModel[tag.Tag]{DB: ds.db}}
}
