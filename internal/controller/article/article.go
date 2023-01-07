package article

import (
	"go-blog/internal/db/model/article"
	"go-blog/internal/db/model/category"
	"go-blog/internal/db/store"
	"go-blog/internal/service"
)

type ArticleCreate struct {
	Title       string            `json:"title" binding:"required,min=2"`
	Summary     string            `json:"summary"`
	Cover       string            `json:"cover" binding:"required"`
	Category    category.Category `json:"category" binding:"required"`
	Tags        []article.Tag     `json:"tags"`
	Series      article.Series    `json:"series"`
	ArticleBody ArticleBody       `json:"article_body"`
}

type ArticleBody struct {
	Content string `json:"content" binding:"required"`
}

type ArticleUpdate struct {
	ID          uint              `json:"id" binding:"required"`
	Title       string            `json:"title" binding:"required,min=2"`
	Summary     string            `json:"summary"`
	Cover       string            `json:"cover" binding:"required"`
	Category    category.Category `json:"category" binding:"required"`
	Tags        []article.Tag     `json:"tags"`
	Series      article.Series    `json:"series"`
	ArticleBody ArticleBodyUpdate `json:"article_body"`
}

type ArticleBodyUpdate struct {
	ID      uint   `json:"id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type Controller struct {
	srv service.Service
}

func NewController(store store.Factory) *Controller {
	return &Controller{service.NewService(store)}
}
