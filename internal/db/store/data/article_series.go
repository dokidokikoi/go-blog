package data

import (
	"context"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/store"
	"go-blog/internal/db/store/data/postgres"
	meta "go-blog/pkg/meta/option"
)

type articleSeries struct {
	pg *postgres.Store
}

func (a articleSeries) Create(ctx context.Context, t *article.Series, option *meta.CreateOption) error {
	return a.pg.ArticleSeries().Create(ctx, t, option)
}

func (a articleSeries) CreateCollection(ctx context.Context, t []*article.Series, option *meta.CreateCollectionOption) []error {
	return a.pg.ArticleSeries().CreateCollection(ctx, t, option)
}

func (a articleSeries) Update(ctx context.Context, t *article.Series, option *meta.UpdateOption) error {
	return a.pg.ArticleSeries().Update(ctx, t, option)
}

func (a articleSeries) UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *article.Series, option *meta.UpdateOption) error {
	return a.pg.ArticleSeries().UpdateByWhere(ctx, node, example, option)
}

func (a articleSeries) UpdateCollection(ctx context.Context, t []*article.Series, option *meta.UpdateCollectionOption) []error {
	return a.pg.ArticleSeries().UpdateCollection(ctx, t, option)
}

func (a articleSeries) Save(ctx context.Context, t *article.Series, option *meta.UpdateOption) error {
	return a.pg.ArticleSeries().Save(ctx, t, option)
}

func (a articleSeries) Get(ctx context.Context, t *article.Series, option *meta.GetOption) (*article.Series, error) {
	return a.pg.ArticleSeries().Get(ctx, t, option)
}

func (a articleSeries) Count(ctx context.Context, t *article.Series, option *meta.GetOption) (int64, error) {
	return a.pg.ArticleSeries().Count(ctx, t, option)
}

func (a articleSeries) CountComplex(ctx context.Context, example *article.Series, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	return a.pg.ArticleSeries().CountComplex(ctx, example, condition, option)
}

func (a articleSeries) List(ctx context.Context, t *article.Series, option *meta.ListOption) ([]*article.Series, error) {
	return a.pg.ArticleSeries().List(ctx, t, option)
}

func (a articleSeries) ListComplex(ctx context.Context, example *article.Series, condition *meta.WhereNode, option *meta.ListOption) ([]*article.Series, error) {
	return a.pg.ArticleSeries().ListComplex(ctx, example, condition, option)
}

func (a articleSeries) Delete(ctx context.Context, t *article.Series, option *meta.DeleteOption) error {
	return a.pg.ArticleSeries().Delete(ctx, t, option)
}

func (a articleSeries) DeleteCollection(ctx context.Context, t []*article.Series, option *meta.DeleteCollectionOption) []error {
	return a.pg.ArticleSeries().DeleteCollection(ctx, t, option)
}

func (a articleSeries) DeleteByIds(ctx context.Context, ids []uint) error {
	return a.pg.ArticleSeries().DeleteByIds(ctx, ids)
}

func newArticleSeries(d *dataCenter) store.ArticleSeries {
	return &articleSeries{pg: d.pg}
}
