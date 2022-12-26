package data

import (
	"context"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/store"
	"go-blog/internal/db/store/data/postgres"
	meta "go-blog/pkg/meta/option"
)

type articleCategories struct {
	pg *postgres.Store
}

func (a articleCategories) Create(ctx context.Context, t *article.Category, option *meta.CreateOption) error {
	return a.pg.ArticleCategories().Create(ctx, t, option)
}

func (a articleCategories) CreateCollection(ctx context.Context, t []*article.Category, option *meta.CreateCollectionOption) []error {
	return a.pg.ArticleCategories().CreateCollection(ctx, t, option)
}

func (a articleCategories) Update(ctx context.Context, t *article.Category, option *meta.UpdateOption) error {
	return a.pg.ArticleCategories().Update(ctx, t, option)
}

func (a articleCategories) UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *article.Category, option *meta.UpdateOption) error {
	return a.pg.ArticleCategories().UpdateByWhere(ctx, node, example, option)
}

func (a articleCategories) UpdateCollection(ctx context.Context, t []*article.Category, option *meta.UpdateCollectionOption) []error {
	return a.pg.ArticleCategories().UpdateCollection(ctx, t, option)
}

func (a articleCategories) Save(ctx context.Context, t *article.Category, option *meta.UpdateOption) error {
	return a.pg.ArticleCategories().Save(ctx, t, option)
}

func (a articleCategories) Get(ctx context.Context, t *article.Category, option *meta.GetOption) (*article.Category, error) {
	return a.pg.ArticleCategories().Get(ctx, t, option)
}

func (a articleCategories) Count(ctx context.Context, t *article.Category, option *meta.GetOption) (int64, error) {
	return a.pg.ArticleCategories().Count(ctx, t, option)
}

func (a articleCategories) CountComplex(ctx context.Context, example *article.Category, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	return a.pg.ArticleCategories().CountComplex(ctx, example, condition, option)
}

func (a articleCategories) List(ctx context.Context, t *article.Category, option *meta.ListOption) ([]*article.Category, error) {
	return a.pg.ArticleCategories().List(ctx, t, option)
}

func (a articleCategories) ListComplex(ctx context.Context, example *article.Category, condition *meta.WhereNode, option *meta.ListOption) ([]*article.Category, error) {
	return a.pg.ArticleCategories().ListComplex(ctx, example, condition, option)
}

func (a articleCategories) Delete(ctx context.Context, t *article.Category, option *meta.DeleteOption) error {
	return a.pg.ArticleCategories().Delete(ctx, t, option)
}

func (a articleCategories) DeleteCollection(ctx context.Context, t []*article.Category, option *meta.DeleteCollectionOption) []error {
	return a.pg.ArticleCategories().DeleteCollection(ctx, t, option)
}

func (a articleCategories) DeleteByIds(ctx context.Context, ids []uint) error {
	return a.pg.ArticleCategories().DeleteByIds(ctx, ids)
}

func newArticleCategories(d *dataCenter) store.ArticleCategory {
	return &articleCategories{pg: d.pg}
}
