package article

import (
	"context"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/store"
	meta "go-blog/pkg/meta/option"
)

type TagSrv interface {
	Create(ctx context.Context, a *article.Tag) error
	IsExist(ctx context.Context, a *article.Tag) (bool, error)

	CreateArticleTag(ctx context.Context, a *article.ArticleTag) error
	CreateArticleTagCollection(ctx context.Context, as []*article.ArticleTag) []error
}

type tagSrv struct {
	store store.Factory
}

func (t tagSrv) Create(ctx context.Context, a *article.Tag) error {
	a.ID = 0
	return t.store.ArticleTag().Create(ctx, a, nil)
}

func (t tagSrv) IsExist(ctx context.Context, a *article.Tag) (bool, error) {
	if a.ID == 0 {
		return false, nil
	}
	tag, err := t.store.ArticleTag().Get(ctx, a, &meta.GetOption{Include: []string{"id"}})
	if err != nil {
		return false, err
	}
	if tag == nil {
		return false, err
	}
	return true, nil
}

func (t tagSrv) CreateArticleTag(ctx context.Context, a *article.ArticleTag) error {
	return t.store.ArticleArticleTag().Create(ctx, a, nil)
}

func (t tagSrv) CreateArticleTagCollection(ctx context.Context, as []*article.ArticleTag) []error {
	return t.store.ArticleArticleTag().CreateCollection(ctx, as, nil)
}

func NewTagSrv(store store.Factory) TagSrv {
	return &tagSrv{store: store}
}
