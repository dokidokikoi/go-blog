package service

import (
	"go-blog/internal/db/store"
	"go-blog/internal/service/article"
	"go-blog/internal/service/category"
	"go-blog/internal/service/user"
	"sync"
)

type Service interface {
	Article() article.ArticleSrv
	ArticleTag() article.TagSrv
	Category() category.CategorySrv
	ArticleSeries() article.SeriesSrv

	User() user.UserSrv
	Role() user.RoleSrv
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

func (s service) Category() category.CategorySrv {
	return category.NewCategorySrv(s.store)
}

func (s service) ArticleSeries() article.SeriesSrv {
	return article.NewSeriesSrv(s.store)
}

func (s service) User() user.UserSrv {
	return user.NewUserSrv(s.store)
}

func (s service) Role() user.RoleSrv {
	return user.NewRoleSrv(s.store)
}

func NewService(store store.Factory) Service {
	once.Do(func() {
		serviceClient = &service{
			store: store,
		}
	})

	return serviceClient
}
