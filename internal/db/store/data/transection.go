package data

import (
	"go-blog/internal/db/store"
	"go-blog/internal/db/store/data/mongo"
	"go-blog/internal/db/store/data/postgres"
	"go-blog/internal/db/store/data/redis"
)

type transaction struct {
	mongo    *mongo.Store
	pg       *postgres.Store
	redisCli *redis.Store
}

func (t *transaction) TransactionBegin() store.Factory {
	return &dataCenter{mongo: t.mongo, redisCli: t.redisCli, pg: t.pg.Transaction().TransactionBegin()}
}

func (t *transaction) TransactionRollback() {
	t.pg.Transaction().TransactionRollback()
}

func (t *transaction) TransactionCommit() {
	t.pg.Transaction().TransactionCommit()
}

var _ store.Transaction = (*transaction)(nil)

func newTransaction(center *dataCenter) *transaction {
	return &transaction{
		pg:       center.pg,
		redisCli: center.redisCli,
		mongo:    center.mongo,
	}
}
