package data

import (
	"context"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/store"
	"go-blog/internal/db/store/data/postgres"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type articleTag struct {
	pg *postgres.Store
}

func (a articleTag) Create(ctx context.Context, t *article.ArticleTag, option *meta.CreateOption) error {
	return a.pg.ArticleTag().Create(ctx, t, option)
}

func (a articleTag) CreateCollection(ctx context.Context, t []*article.ArticleTag, option *meta.CreateCollectionOption) []error {
	return a.pg.ArticleTag().CreateCollection(ctx, t, option)
}

func (a articleTag) Update(ctx context.Context, t *article.ArticleTag, option *meta.UpdateOption) error {
	return a.pg.ArticleTag().Update(ctx, t, option)
}

func (a articleTag) UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *article.ArticleTag, option *meta.UpdateOption) error {
	return a.pg.ArticleTag().UpdateByWhere(ctx, node, example, option)
}

func (a articleTag) UpdateCollection(ctx context.Context, t []*article.ArticleTag, option *meta.UpdateCollectionOption) []error {
	return a.pg.ArticleTag().UpdateCollection(ctx, t, option)
}

func (a articleTag) Save(ctx context.Context, t *article.ArticleTag, option *meta.UpdateOption) error {
	return a.pg.ArticleTag().Save(ctx, t, option)
}

func (a articleTag) Get(ctx context.Context, t *article.ArticleTag, option *meta.GetOption) (*article.ArticleTag, error) {
	return a.pg.ArticleTag().Get(ctx, t, option)
}

func (a articleTag) Count(ctx context.Context, t *article.ArticleTag, option *meta.GetOption) (int64, error) {
	return a.pg.ArticleTag().Count(ctx, t, option)
}

func (a articleTag) CountComplex(ctx context.Context, example *article.ArticleTag, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	return a.pg.ArticleTag().CountComplex(ctx, example, condition, option)
}

func (a articleTag) List(ctx context.Context, t *article.ArticleTag, option *meta.ListOption) ([]*article.ArticleTag, error) {
	return a.pg.ArticleTag().List(ctx, t, option)
}

func (a articleTag) ListComplex(ctx context.Context, example *article.ArticleTag, condition *meta.WhereNode, option *meta.ListOption) ([]*article.ArticleTag, error) {
	return a.pg.ArticleTag().ListComplex(ctx, example, condition, option)
}

func (a articleTag) Delete(ctx context.Context, t *article.ArticleTag, option *meta.DeleteOption) error {
	return a.pg.ArticleTag().Delete(ctx, t, option)
}

func (a articleTag) DeleteCollection(ctx context.Context, t []*article.ArticleTag, option *meta.DeleteCollectionOption) []error {
	return a.pg.ArticleTag().DeleteCollection(ctx, t, option)
}

func (a articleTag) DeleteByIds(ctx context.Context, ids []uint) error {
	return a.pg.ArticleTag().DeleteByIds(ctx, ids)
}

func newArticleTag(d *dataCenter) store.ArticleTag {
	return &articleTag{pg: d.pg}
}
