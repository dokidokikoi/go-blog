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

func (d *dataCenter) Transaction() store.Transaction {
	return newTransaction(d)
}

func (d *dataCenter) Article() store.Article {
	return newArticles(d)
}

func (d *dataCenter) ArticleBody() store.ArticleBody {
	return newArticleBodys(d)
}

func (d *dataCenter) Category() store.Category {
	return newCategories(d)
}

func (d *dataCenter) Tag() store.Tag {
	return newTags(d)
}

func (d *dataCenter) Series() store.Series {
	return newSeries(d)
}

func (d *dataCenter) ArticleTag() store.ArticleTag {
	return newArticleTag(d)
}

func (d *dataCenter) Users() store.Users {
	return newUsers(d)
}

func (d *dataCenter) Roles() store.Roles {
	return newRoles(d)
}

func (d *dataCenter) Items() store.Items {
	return d.pg.Items()
}

func (d *dataCenter) Sites() store.Sites {
	return d.pg.Sites()
}

func (d *dataCenter) SiteTags() store.SiteTag {
	return d.pg.SiteTags()
}

func (d *dataCenter) Link() store.Link {
	return d.pg.Links()
}

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
