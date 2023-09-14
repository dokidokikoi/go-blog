package service

import (
	"go-blog/internal/db/store"
	"sync"
)

type Service interface {
	// Article() article.ArticleSrv
	// ArticleTag() article.TagSrv
	Category() CategorySrv
	Tag() TagSrv
	Series() SeriesSrv

	Items() ItemSrv

	User() UserSrv
	Role() RoleSrv
}

type service struct {
	store store.Factory
}

var (
	serviceClient Service
	once          sync.Once
)

// func (s service) Article() article.ArticleSrv {
// 	return article.NewArticleSrv(s.store)
// }

// func (s service) ArticleTag() article.TagSrv {
// 	return article.NewTagSrv(s.store)
// }

func (s service) Category() CategorySrv {
	return newCategorySrv(s.store)
}

func (s service) Tag() TagSrv {
	return newTagSrv(s.store)
}

func (s service) Series() SeriesSrv {
	return newSeriesSrv(s.store)
}

func (s service) User() UserSrv {
	return newUserSrv(s.store)
}

func (s service) Role() RoleSrv {
	return newRoleSrv(s.store)
}

func (s service) Items() ItemSrv {
	return newItemSrv(s.store)
}

func NewService(store store.Factory) Service {
	once.Do(func() {
		serviceClient = &service{
			store: store,
		}
	})

	return serviceClient
}
