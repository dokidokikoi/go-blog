package data

import (
	"go-blog/internal/db/store"
	"go-blog/internal/db/store/data/mongo"
	"go-blog/internal/db/store/data/postgres"
	"go-blog/internal/db/store/data/redis"
	"sync"
)

type dataCenter struct {
	pg    *postgres.Store
	redis *redis.Store
	mongo *mongo.Store
}

var (
	datacFactory store.Factory
	once         sync.Once
)

func GetStoreDBFactory() (store.Factory, error) {
	once.Do(func() {
		pg, err := postgres.GetPGFactory()
		if err != nil {
			panic(err)
		}

		redisCli, err := redis.GetRedisFactory()
		if err != nil {
			panic(err)
		}

		datacFactory = &dataCenter{
			pg:    pg,
			redis: redisCli,
		}
	})

	return datacFactory, nil
}

func SetStoreDBFactory() {
	factory, err := GetStoreDBFactory()
	if err != nil {
		return
	}

	store.SetStoreFactory(factory)
}
