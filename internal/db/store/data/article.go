package data

import (
	"context"
	meta "github.com/dokidokikoi/go-common/meta/option"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/store"
	"go-blog/internal/db/store/data/postgres"
)

type articles struct {
	pg *postgres.Store
}

func (a articles) Create(ctx context.Context, t *article.Article, option *meta.CreateOption) error {
	return a.pg.Articles().Create(ctx, t, option)
}

func (a articles) CreateCollection(ctx context.Context, t []*article.Article, option *meta.CreateCollectionOption) []error {
	return a.pg.Articles().CreateCollection(ctx, t, option)
}

func (a articles) Update(ctx context.Context, t *article.Article, option *meta.UpdateOption) error {
	return a.pg.Articles().Update(ctx, t, option)
}

func (a articles) UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *article.Article, option *meta.UpdateOption) error {
	return a.pg.Articles().UpdateByWhere(ctx, node, example, option)
}

func (a articles) UpdateCollection(ctx context.Context, t []*article.Article, option *meta.UpdateCollectionOption) []error {
	return a.pg.Articles().UpdateCollection(ctx, t, option)
}

func (a articles) Save(ctx context.Context, t *article.Article, option *meta.UpdateOption) error {
	return a.pg.Articles().Save(ctx, t, option)
}

func (a articles) Get(ctx context.Context, t *article.Article, option *meta.GetOption) (*article.Article, error) {
	return a.pg.Articles().Get(ctx, t, option)
}

func (a articles) Count(ctx context.Context, t *article.Article, option *meta.GetOption) (int64, error) {
	return a.pg.Articles().Count(ctx, t, option)
}

func (a articles) CountComplex(ctx context.Context, example *article.Article, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	return a.pg.Articles().CountComplex(ctx, example, condition, option)
}

func (a articles) List(ctx context.Context, t *article.Article, option *meta.ListOption) ([]*article.Article, error) {
	return a.pg.Articles().List(ctx, t, option)
}

func (a articles) ListComplex(ctx context.Context, example *article.Article, condition *meta.WhereNode, option *meta.ListOption) ([]*article.Article, error) {
	return a.pg.Articles().ListComplex(ctx, example, condition, option)
}

func (a articles) Delete(ctx context.Context, t *article.Article, option *meta.DeleteOption) error {
	return a.pg.Articles().Delete(ctx, t, option)
}

func (a articles) DeleteCollection(ctx context.Context, t []*article.Article, option *meta.DeleteCollectionOption) []error {
	return a.pg.Articles().DeleteCollection(ctx, t, option)
}

func (a articles) DeleteByIds(ctx context.Context, ids []uint) error {
	return a.pg.Articles().DeleteByIds(ctx, ids)
}

func newArticles(d *dataCenter) store.Article {
	return &articles{pg: d.pg}
}
