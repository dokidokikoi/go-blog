package postgres

import (
	"go-blog/internal/db/model/user"
	db "go-blog/pkg/db/base"
)

type users struct {
	db.PgModel[user.User]
}

func newUsers(ds *Store) *users {
	return &users{db.PgModel[user.User]{DB: ds.db}}
}
