package data

import (
	"context"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/store"
	"go-blog/internal/db/store/data/postgres"
	meta "go-blog/pkg/meta/option"
)

type articleArticleTag struct {
	pg *postgres.Store
}

func (a articleArticleTag) Create(ctx context.Context, t *article.ArticleTag, option *meta.CreateOption) error {
	return a.pg.ArticleArticleTag().Create(ctx, t, option)
}

func (a articleArticleTag) CreateCollection(ctx context.Context, t []*article.ArticleTag, option *meta.CreateCollectionOption) []error {
	return a.pg.ArticleArticleTag().CreateCollection(ctx, t, option)
}

func (a articleArticleTag) Update(ctx context.Context, t *article.ArticleTag, option *meta.UpdateOption) error {
	return a.pg.ArticleArticleTag().Update(ctx, t, option)
}

func (a articleArticleTag) UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *article.ArticleTag, option *meta.UpdateOption) error {
	return a.pg.ArticleArticleTag().UpdateByWhere(ctx, node, example, option)
}

func (a articleArticleTag) UpdateCollection(ctx context.Context, t []*article.ArticleTag, option *meta.UpdateCollectionOption) []error {
	return a.pg.ArticleArticleTag().UpdateCollection(ctx, t, option)
}

func (a articleArticleTag) Save(ctx context.Context, t *article.ArticleTag, option *meta.UpdateOption) error {
	return a.pg.ArticleArticleTag().Save(ctx, t, option)
}

func (a articleArticleTag) Get(ctx context.Context, t *article.ArticleTag, option *meta.GetOption) (*article.ArticleTag, error) {
	return a.pg.ArticleArticleTag().Get(ctx, t, option)
}

func (a articleArticleTag) Count(ctx context.Context, t *article.ArticleTag, option *meta.GetOption) (int64, error) {
	return a.pg.ArticleArticleTag().Count(ctx, t, option)
}

func (a articleArticleTag) CountComplex(ctx context.Context, example *article.ArticleTag, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	return a.pg.ArticleArticleTag().CountComplex(ctx, example, condition, option)
}

func (a articleArticleTag) List(ctx context.Context, t *article.ArticleTag, option *meta.ListOption) ([]*article.ArticleTag, error) {
	return a.pg.ArticleArticleTag().List(ctx, t, option)
}

func (a articleArticleTag) ListComplex(ctx context.Context, example *article.ArticleTag, condition *meta.WhereNode, option *meta.ListOption) ([]*article.ArticleTag, error) {
	return a.pg.ArticleArticleTag().ListComplex(ctx, example, condition, option)
}

func (a articleArticleTag) Delete(ctx context.Context, t *article.ArticleTag, option *meta.DeleteOption) error {
	return a.pg.ArticleArticleTag().Delete(ctx, t, option)
}

func (a articleArticleTag) DeleteCollection(ctx context.Context, t []*article.ArticleTag, option *meta.DeleteCollectionOption) []error {
	return a.pg.ArticleArticleTag().DeleteCollection(ctx, t, option)
}

func (a articleArticleTag) DeleteByIds(ctx context.Context, ids []uint) error {
	return a.pg.ArticleArticleTag().DeleteByIds(ctx, ids)
}

func newArticleArticleTag(d *dataCenter) store.ArticleArticleTag {
	return &articleArticleTag{pg: d.pg}
}
