package postgres

import (
	db "github.com/dokidokikoi/go-common/db/base"
	"go-blog/internal/db/model/article"
)

type articleTags struct {
	db.PgModel[article.ArticleTag]
}

func newArticleTags(ds *Store) *articleTags {
	return &articleTags{PgModel: db.PgModel[article.ArticleTag]{DB: ds.db}}
}
