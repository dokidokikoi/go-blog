package data

import (
	"context"
	"go-blog/internal/db/model/category"
	"go-blog/internal/db/store"
	"go-blog/internal/db/store/data/postgres"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type articleCategories struct {
	pg *postgres.Store
}

func (a articleCategories) Create(ctx context.Context, t *category.Category, option *meta.CreateOption) error {
	return a.pg.Categories().Create(ctx, t, option)
}

func (a articleCategories) CreateCollection(ctx context.Context, t []*category.Category, option *meta.CreateCollectionOption) []error {
	return a.pg.Categories().CreateCollection(ctx, t, option)
}

func (a articleCategories) Update(ctx context.Context, t *category.Category, option *meta.UpdateOption) error {
	return a.pg.Categories().Update(ctx, t, option)
}

func (a articleCategories) UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *category.Category, option *meta.UpdateOption) error {
	return a.pg.Categories().UpdateByWhere(ctx, node, example, option)
}

func (a articleCategories) UpdateCollection(ctx context.Context, t []*category.Category, option *meta.UpdateCollectionOption) []error {
	return a.pg.Categories().UpdateCollection(ctx, t, option)
}

func (a articleCategories) Save(ctx context.Context, t *category.Category, option *meta.UpdateOption) error {
	return a.pg.Categories().Save(ctx, t, option)
}

func (a articleCategories) Get(ctx context.Context, t *category.Category, option *meta.GetOption) (*category.Category, error) {
	return a.pg.Categories().Get(ctx, t, option)
}

func (a articleCategories) Count(ctx context.Context, t *category.Category, option *meta.GetOption) (int64, error) {
	return a.pg.Categories().Count(ctx, t, option)
}

func (a articleCategories) CountComplex(ctx context.Context, example *category.Category, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	return a.pg.Categories().CountComplex(ctx, example, condition, option)
}

func (a articleCategories) List(ctx context.Context, t *category.Category, option *meta.ListOption) ([]*category.Category, error) {
	return a.pg.Categories().List(ctx, t, option)
}

func (a articleCategories) ListComplex(ctx context.Context, example *category.Category, condition *meta.WhereNode, option *meta.ListOption) ([]*category.Category, error) {
	return a.pg.Categories().ListComplex(ctx, example, condition, option)
}

func (a articleCategories) Delete(ctx context.Context, t *category.Category, option *meta.DeleteOption) error {
	return a.pg.Categories().Delete(ctx, t, option)
}

func (a articleCategories) DeleteCollection(ctx context.Context, t []*category.Category, option *meta.DeleteCollectionOption) []error {
	return a.pg.Categories().DeleteCollection(ctx, t, option)
}

func (a articleCategories) DeleteByIds(ctx context.Context, ids []uint) error {
	return a.pg.Categories().DeleteByIds(ctx, ids)
}

func newCategories(d *dataCenter) store.Category {
	return &articleCategories{pg: d.pg}
}
