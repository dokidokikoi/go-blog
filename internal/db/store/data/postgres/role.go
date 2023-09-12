package postgres

import (
	db "github.com/dokidokikoi/go-common/db/base"
	"go-blog/internal/db/model/user"
)

type roles struct {
	db.PgModel[user.Role]
}

func newRoles(ds *Store) *roles {
	return &roles{PgModel: db.PgModel[user.Role]{DB: ds.db}}
}
