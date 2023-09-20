package service

import (
	"context"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/store"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type ArticleSrv interface {
	Create(ctx context.Context, example *article.Article, option *meta.CreateOption) error
	Get(ctx context.Context, example *article.Article, option *meta.GetOption) (*article.Article, error)
	Update(ctx context.Context, example *article.Article, option *meta.UpdateOption) error
	Del(ctx context.Context, example *article.Article, option *meta.DeleteOption) error
	DeleteCollection(ctx context.Context, examples []*article.Article, option *meta.DeleteCollectionOption) []error
	List(ctx context.Context, example *article.Article, option *meta.ListOption) ([]*article.Article, int64, error)
	ListByWhereNode(ctx context.Context, example *article.Article, node *meta.WhereNode, option *meta.ListOption) ([]*article.Article, int64, error)

	DeleteArticleAllTags(ctx context.Context, articleID uint) error
	CreateArticleTagsCollection(ctx context.Context, ats []*article.ArticleTag, option *meta.CreateCollectionOption) []error
	ListTagArticle(ctx context.Context, tagID uint, option *meta.ListOption) ([]*article.Article, int64, error)
}

type articleSrv struct {
	store store.Factory
}

func (as *articleSrv) Create(ctx context.Context, example *article.Article, option *meta.CreateOption) error {
	return as.store.Article().Create(ctx, example, option)
}

func (as *articleSrv) Get(ctx context.Context, example *article.Article, option *meta.GetOption) (*article.Article, error) {
	return as.store.Article().Get(ctx, example, option)
}

func (as *articleSrv) Update(ctx context.Context, example *article.Article, option *meta.UpdateOption) error {
	return as.store.Article().Update(ctx, example, option)
}

func (as *articleSrv) Del(ctx context.Context, example *article.Article, option *meta.DeleteOption) error {
	return as.store.Article().Delete(ctx, example, option)
}

func (as *articleSrv) DeleteCollection(ctx context.Context, examples []*article.Article, option *meta.DeleteCollectionOption) []error {
	return as.store.Article().DeleteCollection(ctx, examples, option)
}

func (as *articleSrv) List(ctx context.Context, example *article.Article, option *meta.ListOption) ([]*article.Article, int64, error) {
	total, err := as.store.Article().Count(ctx, example, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}

	list, err := as.store.Article().List(ctx, example, option)
	return list, total, err
}

func (as *articleSrv) ListByWhereNode(ctx context.Context, example *article.Article, node *meta.WhereNode, option *meta.ListOption) ([]*article.Article, int64, error) {
	total, err := as.store.Article().CountComplex(ctx, example, node, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}

	list, err := as.store.Article().ListComplex(ctx, example, node, option)
	return list, total, err
}

func (as *articleSrv) DeleteArticleAllTags(ctx context.Context, articleID uint) error {
	return as.store.ArticleTag().Delete(ctx, &article.ArticleTag{ArticleID: articleID}, nil)
}

func (as *articleSrv) ListTagArticle(ctx context.Context, tagID uint, option *meta.ListOption) ([]*article.Article, int64, error) {
	articleTags, err := as.store.ArticleTag().List(ctx, &article.ArticleTag{TagID: tagID}, &meta.ListOption{PageSize: 1000})
	if err != nil {
		return nil, 0, err
	}
	articleIDs := []uint{}
	for _, at := range articleTags {
		articleIDs = append(articleIDs, at.ArticleID)
	}
	node := &meta.WhereNode{
		Conditions: []*meta.Condition{
			{
				Field:    "id",
				Operator: meta.IN,
				Value:    articleIDs,
			},
		},
	}
	articles, err := as.store.Article().ListComplex(ctx, nil, node, option)
	return articles, int64(len(articleTags)), err
}

func (as *articleSrv) CreateArticleTagsCollection(ctx context.Context, ats []*article.ArticleTag, option *meta.CreateCollectionOption) []error {
	return as.store.ArticleTag().CreateCollection(ctx, ats, option)
}

func newArticleSrv(store store.Factory) ArticleSrv {
	return &articleSrv{store: store}
}
