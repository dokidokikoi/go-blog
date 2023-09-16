package postgres

import (
	"go-blog/internal/db/model/comment"

	db "github.com/dokidokikoi/go-common/db/base"
)

type comments struct {
	db.PgModel[comment.Comment]
}

func newComments(ds *Store) *comments {
	return &comments{PgModel: db.PgModel[comment.Comment]{DB: ds.db}}
}
