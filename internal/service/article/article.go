package article

import (
	"context"
	"fmt"
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/store"
	meta "go-blog/pkg/meta/option"

	"gorm.io/gorm"
)

type ArticleSrv interface {
	CreateArticle(ctx context.Context, a *article.Article) error
	CreateBody(ctx context.Context, a *article.ArticleBody) error
	List(ctx context.Context, keyword string, option *meta.ListOption) ([]*article.Article, int64, error)
	UpdateArticle(ctx context.Context, a *article.Article) error
	UpdateBody(ctx context.Context, a *article.ArticleBody) error
	GetArticleByID(ctx context.Context, id uint, option *meta.GetOption) (*article.Article, error)
	UpdateViewCntByID(ctx context.Context, id uint)
	UpdateCommentCntByID(ctx context.Context, id uint)
	DeleteArticleByID(ctx context.Context, id uint) error
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

func (srv articleSrv) UpdateArticle(ctx context.Context, a *article.Article) error {
	return srv.store.Article().Update(ctx, a, nil)
}

func (srv articleSrv) UpdateBody(ctx context.Context, a *article.ArticleBody) error {
	return srv.store.ArticleBody().Update(ctx, a, nil)
}

func (srv articleSrv) GetArticleByID(ctx context.Context, id uint, option *meta.GetOption) (*article.Article, error) {
	return srv.store.Article().Get(ctx, &article.Article{Model: gorm.Model{ID: id}}, option)
}

func (srv articleSrv) UpdateViewCntByID(ctx context.Context, id uint) {
	go func() {
		for i := 0; i < 10; i++ {
			option := meta.GetOption{}
			option.Select = append(option.Select, "view_counts")
			res, _ := srv.store.Article().Get(ctx, &article.Article{Model: gorm.Model{ID: id}}, &option)
			if res == nil {
				return
			}

			updateOption := meta.UpdateOption{}
			updateOption.Select = append(updateOption.Select, "view_counts")
			node := meta.WhereNode{}
			node.Conditions = append(node.Conditions, &meta.Condition{Field: "view_counts", Operator: meta.EQUAL, Value: res.ViewCounts})
			if srv.store.Article().UpdateByWhere(ctx, &node, &article.Article{Model: gorm.Model{ID: id}, ViewCounts: res.ViewCounts + 1}, &updateOption) == nil {
				return
			}
		}
	}()
}

func (srv articleSrv) UpdateCommentCntByID(ctx context.Context, id uint) {
	go func() {
		for i := 0; i < 10; i++ {
			option := meta.GetOption{}
			option.Select = append(option.Select, "comment_counts")
			res, _ := srv.store.Article().Get(ctx, &article.Article{Model: gorm.Model{ID: id}}, &option)
			if res == nil {
				return
			}

			updateOption := meta.UpdateOption{}
			updateOption.Select = append(updateOption.Select, "comment_counts")
			node := meta.WhereNode{}
			node.Conditions = append(
				node.Conditions,
				&meta.Condition{
					Field:    "comment_counts",
					Operator: meta.EQUAL,
					Value:    res.CommentCounts})
			if srv.store.Article().UpdateByWhere(ctx, &node, &article.Article{Model: gorm.Model{ID: id}, CommentCounts: res.CommentCounts + 1}, &updateOption) == nil {
				return
			}
		}
	}()
}

func (srv articleSrv) DeleteArticleByID(ctx context.Context, id uint) error {
	return srv.store.Article().Delete(ctx, &article.Article{Model: gorm.Model{ID: id}}, nil)
}

func NewArticleSrv(store store.Factory) ArticleSrv {
	return &articleSrv{store: store}
}
