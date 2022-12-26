package postgres

import (
	"go-blog/internal/db/model/article"
	db "go-blog/pkg/db/base"
)

type articleSeries struct {
	db.PgModel[article.Series]
}

func newArticleSeries(ds *Store) *articleSeries {
	return &articleSeries{PgModel: db.PgModel[article.Series]{DB: ds.db}}
}
