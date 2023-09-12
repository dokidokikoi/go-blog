package data

import (
	"context"
	"go-blog/internal/db/model/series"
	"go-blog/internal/db/store"
	"go-blog/internal/db/store/data/postgres"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type seriess struct {
	pg *postgres.Store
}

func (s seriess) Create(ctx context.Context, t *series.Series, option *meta.CreateOption) error {
	return s.pg.Series().Create(ctx, t, option)
}

func (s seriess) CreateCollection(ctx context.Context, t []*series.Series, option *meta.CreateCollectionOption) []error {
	return s.pg.Series().CreateCollection(ctx, t, option)
}

func (s seriess) Update(ctx context.Context, t *series.Series, option *meta.UpdateOption) error {
	return s.pg.Series().Update(ctx, t, option)
}

func (s seriess) UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *series.Series, option *meta.UpdateOption) error {
	return s.pg.Series().UpdateByWhere(ctx, node, example, option)
}

func (s seriess) UpdateCollection(ctx context.Context, t []*series.Series, option *meta.UpdateCollectionOption) []error {
	return s.pg.Series().UpdateCollection(ctx, t, option)
}

func (s seriess) Save(ctx context.Context, t *series.Series, option *meta.UpdateOption) error {
	return s.pg.Series().Save(ctx, t, option)
}

func (s seriess) Get(ctx context.Context, t *series.Series, option *meta.GetOption) (*series.Series, error) {
	return s.pg.Series().Get(ctx, t, option)
}

func (s seriess) Count(ctx context.Context, t *series.Series, option *meta.GetOption) (int64, error) {
	return s.pg.Series().Count(ctx, t, option)
}

func (s seriess) CountComplex(ctx context.Context, example *series.Series, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	return s.pg.Series().CountComplex(ctx, example, condition, option)
}

func (s seriess) List(ctx context.Context, t *series.Series, option *meta.ListOption) ([]*series.Series, error) {
	return s.pg.Series().List(ctx, t, option)
}

func (s seriess) ListComplex(ctx context.Context, example *series.Series, condition *meta.WhereNode, option *meta.ListOption) ([]*series.Series, error) {
	return s.pg.Series().ListComplex(ctx, example, condition, option)
}

func (s seriess) Delete(ctx context.Context, t *series.Series, option *meta.DeleteOption) error {
	return s.pg.Series().Delete(ctx, t, option)
}

func (s seriess) DeleteCollection(ctx context.Context, t []*series.Series, option *meta.DeleteCollectionOption) []error {
	return s.pg.Series().DeleteCollection(ctx, t, option)
}

func (s seriess) DeleteByIds(ctx context.Context, ids []uint) error {
	return s.pg.Series().DeleteByIds(ctx, ids)
}

func newSeries(d *dataCenter) store.Series {
	return &seriess{pg: d.pg}
}
