package article

import (
	"context"
	"fmt"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/store"
	meta "go-blog/pkg/meta/option"
)

type ArticleSrv interface {
	CreateArticle(ctx context.Context, a *article.Article) error
	CreateBody(ctx context.Context, a *article.ArticleBody) error
	List(ctx context.Context, keyword string, option *meta.ListOption) ([]*article.Article, int64, error)
}

type articleSrv struct {
	store store.Factory
}

func (srv articleSrv) CreateArticle(ctx context.Context, a *article.Article) error {
	a.ID = 0
	return srv.store.Article().Create(ctx, a, nil)
}

func (srv articleSrv) CreateBody(ctx context.Context, a *article.ArticleBody) error {
	a.ID = 0
	return srv.store.ArticleBody().Create(ctx, a, nil)
}

func (srv articleSrv) List(ctx context.Context, keyword string, option *meta.ListOption) ([]*article.Article, int64, error) {
	if keyword == "" {
		count, _ := srv.store.Article().Count(ctx, &article.Article{}, &option.GetOption)
		result, err := srv.store.Article().List(ctx, &article.Article{}, option)
		return result, count, err
	}

	root := &meta.WhereNode{}
	root.Conditions = append(root.Conditions,
		&meta.Condition{Field: "title", Operator: meta.LIKE, Value: fmt.Sprintf("%s%s%s", "%", keyword, "%")})
	count, _ := srv.store.Article().CountComplex(ctx, &article.Article{}, root, &option.GetOption)
	result, err := srv.store.Article().ListComplex(ctx, &article.Article{}, root, option)
	return result, count, err
}

func NewArticleSrv(store store.Factory) ArticleSrv {
	return &articleSrv{store: store}
}
