package article

import (
	"context"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/store"
	meta "go-blog/pkg/meta/option"
)

type CategorySrv interface {
	Create(ctx context.Context, a *article.Category) error
	IsExist(ctx context.Context, a *article.Category) (bool, error)
}

type categorySrv struct {
	store store.Factory
}

func (c categorySrv) Create(ctx context.Context, a *article.Category) error {
	a.ID = 0
	return c.store.ArticleCategory().Create(ctx, a, nil)
}

func (c categorySrv) IsExist(ctx context.Context, a *article.Category) (bool, error) {
	if a.ID == 0 {
		return false, nil
	}
	cate, err := c.store.ArticleCategory().Get(ctx, a, &meta.GetOption{Include: []string{"id"}})
	if err != nil {
		return false, err
	}
	if cate == nil {
		return false, err
	}
	return true, nil
}

func NewCategorySrv(store store.Factory) CategorySrv {
	return &categorySrv{store: store}
}
