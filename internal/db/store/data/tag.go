package data

import (
	"context"
	"go-blog/internal/db/model/tag"
	"go-blog/internal/db/store"
	"go-blog/internal/db/store/data/postgres"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type tags struct {
	pg *postgres.Store
}

func (tag tags) Create(ctx context.Context, t *tag.Tag, option *meta.CreateOption) error {
	return tag.pg.Tags().Create(ctx, t, option)
}

func (tag tags) CreateCollection(ctx context.Context, t []*tag.Tag, option *meta.CreateCollectionOption) []error {
	return tag.pg.Tags().CreateCollection(ctx, t, option)
}

func (tag tags) Update(ctx context.Context, t *tag.Tag, option *meta.UpdateOption) error {
	return tag.pg.Tags().Update(ctx, t, option)
}

func (tag tags) UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *tag.Tag, option *meta.UpdateOption) error {
	return tag.pg.Tags().UpdateByWhere(ctx, node, example, option)
}

func (tag tags) UpdateCollection(ctx context.Context, t []*tag.Tag, option *meta.UpdateCollectionOption) []error {
	return tag.pg.Tags().UpdateCollection(ctx, t, option)
}

func (tag tags) Save(ctx context.Context, t *tag.Tag, option *meta.UpdateOption) error {
	return tag.pg.Tags().Save(ctx, t, option)
}

func (tag tags) Get(ctx context.Context, t *tag.Tag, option *meta.GetOption) (*tag.Tag, error) {
	return tag.pg.Tags().Get(ctx, t, option)
}

func (tag tags) Count(ctx context.Context, t *tag.Tag, option *meta.GetOption) (int64, error) {
	return tag.pg.Tags().Count(ctx, t, option)
}

func (tag tags) CountComplex(ctx context.Context, example *tag.Tag, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	return tag.pg.Tags().CountComplex(ctx, example, condition, option)
}

func (tag tags) List(ctx context.Context, t *tag.Tag, option *meta.ListOption) ([]*tag.Tag, error) {
	return tag.pg.Tags().List(ctx, t, option)
}

func (tag tags) ListComplex(ctx context.Context, example *tag.Tag, condition *meta.WhereNode, option *meta.ListOption) ([]*tag.Tag, error) {
	return tag.pg.Tags().ListComplex(ctx, example, condition, option)
}

func (tag tags) Delete(ctx context.Context, t *tag.Tag, option *meta.DeleteOption) error {
	return tag.pg.Tags().Delete(ctx, t, option)
}

func (tag tags) DeleteCollection(ctx context.Context, t []*tag.Tag, option *meta.DeleteCollectionOption) []error {
	return tag.pg.Tags().DeleteCollection(ctx, t, option)
}

func (tag tags) DeleteByIds(ctx context.Context, ids []uint) error {
	return tag.pg.Tags().DeleteByIds(ctx, ids)
}

func newTags(d *dataCenter) store.Tag {
	return &tags{pg: d.pg}
}
