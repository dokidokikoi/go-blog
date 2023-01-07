package postgres

import (
	"go-blog/internal/db/model/user"
	db "go-blog/pkg/db/base"
)

type roles struct {
	db.PgModel[user.Role]
}

func newRoles(ds *Store) *roles {
	return &roles{PgModel: db.PgModel[user.Role]{DB: ds.db}}
}
