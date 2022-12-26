package article

import (
	"context"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/store"
	meta "go-blog/pkg/meta/option"
)

type SeriesSrv interface {
	Create(ctx context.Context, a *article.Series) error
	IsExist(ctx context.Context, a *article.Series) (bool, error)
}

type seriesSrv struct {
	store store.Factory
}

func (s seriesSrv) Create(ctx context.Context, a *article.Series) error {
	a.ID = 0
	return s.store.ArticleSeries().Create(ctx, a, nil)
}

func (s seriesSrv) IsExist(ctx context.Context, a *article.Series) (bool, error) {
	if a.ID == 0 {
		return false, nil
	}
	series, err := s.store.ArticleSeries().Get(ctx, a, &meta.GetOption{Include: []string{"id"}})
	if err != nil {
		return false, err
	}
	if series == nil {
		return false, err
	}
	return true, nil
}

func NewSeriesSrv(store store.Factory) SeriesSrv {
	return &seriesSrv{store: store}
}
