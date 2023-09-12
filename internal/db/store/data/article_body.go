package data

import (
	"context"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/store"
	"go-blog/internal/db/store/data/postgres"

	meta "github.com/dokidokikoi/go-common/meta/option"
	"gorm.io/gorm"
)

type articleBodys struct {
	pg *postgres.Store
}

func (a articleBodys) Create(ctx context.Context, t *article.ArticleBody, option *meta.CreateOption) error {
	return a.pg.ArticleBodys().Create(ctx, t, option)
}

func (a articleBodys) CreateCollection(ctx context.Context, t []*article.ArticleBody, option *meta.CreateCollectionOption) []error {
	return a.pg.ArticleBodys().CreateCollection(ctx, t, option)
}

func (a articleBodys) Update(ctx context.Context, t *article.ArticleBody, option *meta.UpdateOption) error {
	return a.pg.ArticleBodys().Update(ctx, t, option)
}

func (a articleBodys) UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *article.ArticleBody, option *meta.UpdateOption) error {
	return a.pg.ArticleBodys().UpdateByWhere(ctx, node, example, option)
}

func (a articleBodys) UpdateCollection(ctx context.Context, t []*article.ArticleBody, option *meta.UpdateCollectionOption) []error {
	return a.pg.ArticleBodys().UpdateCollection(ctx, t, option)
}

func (a articleBodys) Save(ctx context.Context, t *article.ArticleBody, option *meta.UpdateOption) error {
	return a.pg.ArticleBodys().Save(ctx, t, option)
}

func (a articleBodys) Get(ctx context.Context, t *article.ArticleBody, option *meta.GetOption) (*article.ArticleBody, error) {
	return a.pg.ArticleBodys().Get(ctx, t, option)
}

func (a articleBodys) Count(ctx context.Context, t *article.ArticleBody, option *meta.GetOption) (int64, error) {
	return a.pg.ArticleBodys().Count(ctx, t, option)
}

func (a articleBodys) CountDB(ctx context.Context, t *article.ArticleBody, option *meta.GetOption) *gorm.DB {
	return a.pg.ArticleBodys().CountDB(ctx, t, option)
}

func (a articleBodys) CountComplex(ctx context.Context, example *article.ArticleBody, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	return a.pg.ArticleBodys().CountComplex(ctx, example, condition, option)
}

func (a articleBodys) CountComplexDB(ctx context.Context, example *article.ArticleBody, condition *meta.WhereNode, option *meta.GetOption) *gorm.DB {
	return a.pg.ArticleBodys().CountComplexDB(ctx, example, condition, option)
}

func (a articleBodys) CreateMany2Many(ctx context.Context, example *article.ArticleBody, ids interface{}, option *meta.CreateOption) error {
	return a.pg.ArticleBodys().CreateMany2Many(ctx, example, ids, option)
}

func (a articleBodys) List(ctx context.Context, t *article.ArticleBody, option *meta.ListOption) ([]*article.ArticleBody, error) {
	return a.pg.ArticleBodys().List(ctx, t, option)
}

func (a articleBodys) ListComplex(ctx context.Context, example *article.ArticleBody, condition *meta.WhereNode, option *meta.ListOption) ([]*article.ArticleBody, error) {
	return a.pg.ArticleBodys().ListComplex(ctx, example, condition, option)
}

func (a articleBodys) Delete(ctx context.Context, t *article.ArticleBody, option *meta.DeleteOption) error {
	return a.pg.ArticleBodys().Delete(ctx, t, option)
}

func (a articleBodys) DeleteCollection(ctx context.Context, t []*article.ArticleBody, option *meta.DeleteCollectionOption) []error {
	return a.pg.ArticleBodys().DeleteCollection(ctx, t, option)
}

func (a articleBodys) DeleteByIds(ctx context.Context, ids []uint) error {
	return a.pg.ArticleBodys().DeleteByIds(ctx, ids)
}

func newArticleBodys(d *dataCenter) store.ArticleBody {
	return &articleBodys{pg: d.pg}
}
