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
	UpdateArticleTag(ctx context.Context, a *article.Article) []error
	DeleteArticleTag(ctx context.Context, a *article.ArticleTag) error
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
		tag, err := t.store.ArticleTag().Get(ctx, a, nil)
		if err != nil {
			return false, err
		}
		if tag == nil {
			return false, nil
		}
		a.ID = tag.ID
		return true, nil
	}
	tag, err := t.store.ArticleTag().Get(ctx, &article.Tag{ID: a.ID}, &meta.GetOption{Include: []string{"id"}})
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

func (t tagSrv) UpdateArticleTag(ctx context.Context, a *article.Article) []error {
	store := t.store.Transaction().TransactionBegin()
	var errs []error
	var targetArticleTag []*article.ArticleTag
	for _, tag := range a.Tags {
		targetTag := &article.Tag{
			ID:      tag.ID,
			TagName: tag.TagName,
		}
		ok, err := t.IsExist(ctx, targetTag)
		if err != nil || !ok {
			t.Create(ctx, targetTag)
		}
		targetArticleTag = append(targetArticleTag, &article.ArticleTag{ArticleID: a.ID, TagId: targetTag.ID})
	}

	if err := t.DeleteArticleTag(ctx, &article.ArticleTag{ArticleID: a.ID}); err != nil {
		store.Transaction().TransactionRollback()
		errs = append(errs, err)
		return errs
	}

	if errs = t.CreateArticleTagCollection(ctx, targetArticleTag); errs != nil {
		for _, err := range errs {
			if err != nil {
				store.Transaction().TransactionRollback()
				return errs
			}
		}
	}

	store.Transaction().TransactionCommit()

	return errs
}

func (t tagSrv) DeleteArticleTag(ctx context.Context, a *article.ArticleTag) error {
	return t.store.ArticleArticleTag().Delete(ctx, a, nil)
}

func NewTagSrv(store store.Factory) TagSrv {
	return &tagSrv{store: store}
}
