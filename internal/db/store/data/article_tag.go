package data

import (
	"context"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/store"
	"go-blog/internal/db/store/data/postgres"
	meta "go-blog/pkg/meta/option"
)

type articleTags struct {
	pg *postgres.Store
}

func (a articleTags) Create(ctx context.Context, t *article.Tag, option *meta.CreateOption) error {
	return a.pg.ArticleTags().Create(ctx, t, option)
}

func (a articleTags) CreateCollection(ctx context.Context, t []*article.Tag, option *meta.CreateCollectionOption) []error {
	return a.pg.ArticleTags().CreateCollection(ctx, t, option)
}

func (a articleTags) Update(ctx context.Context, t *article.Tag, option *meta.UpdateOption) error {
	return a.pg.ArticleTags().Update(ctx, t, option)
}

func (a articleTags) UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *article.Tag, option *meta.UpdateOption) error {
	return a.pg.ArticleTags().UpdateByWhere(ctx, node, example, option)
}

func (a articleTags) UpdateCollection(ctx context.Context, t []*article.Tag, option *meta.UpdateCollectionOption) []error {
	return a.pg.ArticleTags().UpdateCollection(ctx, t, option)
}

func (a articleTags) Save(ctx context.Context, t *article.Tag, option *meta.UpdateOption) error {
	return a.pg.ArticleTags().Save(ctx, t, option)
}

func (a articleTags) Get(ctx context.Context, t *article.Tag, option *meta.GetOption) (*article.Tag, error) {
	return a.pg.ArticleTags().Get(ctx, t, option)
}

func (a articleTags) Count(ctx context.Context, t *article.Tag, option *meta.GetOption) (int64, error) {
	return a.pg.ArticleTags().Count(ctx, t, option)
}

func (a articleTags) CountComplex(ctx context.Context, example *article.Tag, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	return a.pg.ArticleTags().CountComplex(ctx, example, condition, option)
}

func (a articleTags) List(ctx context.Context, t *article.Tag, option *meta.ListOption) ([]*article.Tag, error) {
	return a.pg.ArticleTags().List(ctx, t, option)
}

func (a articleTags) ListComplex(ctx context.Context, example *article.Tag, condition *meta.WhereNode, option *meta.ListOption) ([]*article.Tag, error) {
	return a.pg.ArticleTags().ListComplex(ctx, example, condition, option)
}

func (a articleTags) Delete(ctx context.Context, t *article.Tag, option *meta.DeleteOption) error {
	return a.pg.ArticleTags().Delete(ctx, t, option)
}

func (a articleTags) DeleteCollection(ctx context.Context, t []*article.Tag, option *meta.DeleteCollectionOption) []error {
	return a.pg.ArticleTags().DeleteCollection(ctx, t, option)
}

func (a articleTags) DeleteByIds(ctx context.Context, ids []uint) error {
	return a.pg.ArticleTags().DeleteByIds(ctx, ids)
}

func newArticleTags(d *dataCenter) store.ArticleTag {
	return &articleTags{pg: d.pg}
}
