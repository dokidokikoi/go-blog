package service

import (
	"go-blog/internal/db/store"
	"go-blog/internal/service/article"
	"sync"
)

type Service interface {
	Article() article.ArticleSrv
	ArticleTag() article.TagSrv
	ArticleCategory() article.CategorySrv
	ArticleSeries() article.SeriesSrv
}

type service struct {
	store store.Factory
}

var (
	serviceClient Service
	once          sync.Once
)

func (s service) Article() article.ArticleSrv {
	return article.NewArticleSrv(s.store)
}

func (s service) ArticleTag() article.TagSrv {
	return article.NewTagSrv(s.store)
}

func (s service) ArticleCategory() article.CategorySrv {
	return article.NewCategorySrv(s.store)
}

func (s service) ArticleSeries() article.SeriesSrv {
	return article.NewSeriesSrv(s.store)
}

func NewService(store store.Factory) Service {
	once.Do(func() {
		serviceClient = &service{
			store: store,
		}
	})

	return serviceClient
}
