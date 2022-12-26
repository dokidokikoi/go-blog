package postgres

import (
	"go-blog/internal/db/model/article"
	db "go-blog/pkg/db/base"
)

type articleTags struct {
	db.PgModel[article.Tag]
}

func newArticleTags(ds *Store) *articleTags {
	return &articleTags{PgModel: db.PgModel[article.Tag]{DB: ds.db}}
}
