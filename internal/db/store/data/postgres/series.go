package postgres

import (
	db "github.com/dokidokikoi/go-common/db/base"
	"go-blog/internal/db/model/series"
)

type articleSeries struct {
	db.PgModel[series.Series]
}

func newSeries(ds *Store) *articleSeries {
	return &articleSeries{PgModel: db.PgModel[series.Series]{DB: ds.db}}
}
