package postgres

import (
	db "github.com/dokidokikoi/go-common/db/base"
	"go-blog/internal/db/model/user"
)

type users struct {
	db.PgModel[user.User]
}

func newUsers(ds *Store) *users {
	return &users{db.PgModel[user.User]{DB: ds.db}}
}
