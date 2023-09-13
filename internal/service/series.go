package service

import (
	"context"
	"go-blog/internal/db/model/series"
	"go-blog/internal/db/store"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type SeriesSrv interface {
	Create(ctx context.Context, example *series.Series, option *meta.CreateOption) error
	Get(ctx context.Context, example *series.Series, option *meta.GetOption) (*series.Series, error)
	Update(ctx context.Context, example *series.Series, option *meta.UpdateOption) error
	Del(ctx context.Context, example *series.Series, option *meta.DeleteOption) error
	DeleteCollection(ctx context.Context, examples []*series.Series, option *meta.DeleteCollectionOption) []error
	List(ctx context.Context, example *series.Series, option *meta.ListOption) ([]*series.Series, int64, error)
	ListByWhereNode(ctx context.Context, example *series.Series, node *meta.WhereNode, option *meta.ListOption) ([]*series.Series, int64, error)
}

type seriesSrv struct {
	store store.Factory
}

func (ss *seriesSrv) Create(ctx context.Context, example *series.Series, option *meta.CreateOption) error {
	return ss.store.Series().Create(ctx, example, option)
}

func (ss *seriesSrv) Get(ctx context.Context, example *series.Series, option *meta.GetOption) (*series.Series, error) {
	return ss.store.Series().Get(ctx, example, option)
}

func (ss *seriesSrv) Update(ctx context.Context, example *series.Series, option *meta.UpdateOption) error {
	return ss.store.Series().Update(ctx, example, option)
}

func (ss *seriesSrv) Del(ctx context.Context, example *series.Series, option *meta.DeleteOption) error {
	return ss.store.Series().Delete(ctx, example, option)
}

func (ss *seriesSrv) DeleteCollection(ctx context.Context, examples []*series.Series, option *meta.DeleteCollectionOption) []error {
	return ss.store.Series().DeleteCollection(ctx, examples, option)
}

func (ss *seriesSrv) List(ctx context.Context, example *series.Series, option *meta.ListOption) ([]*series.Series, int64, error) {
	total, err := ss.store.Series().Count(ctx, example, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}
	list, err := ss.store.Series().List(ctx, example, option)
	return list, total, err
}

func (ss *seriesSrv) ListByWhereNode(ctx context.Context, example *series.Series, node *meta.WhereNode, option *meta.ListOption) ([]*series.Series, int64, error) {
	total, err := ss.store.Series().CountComplex(ctx, example, node, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}
	list, err := ss.store.Series().ListComplex(ctx, example, node, option)
	return list, total, err
}

func newSeriesSrv(store store.Factory) SeriesSrv {
	return &seriesSrv{
		store: store,
	}
}
