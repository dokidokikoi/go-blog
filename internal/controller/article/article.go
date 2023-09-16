package article

import (
	"go-blog/internal/db/model/series"
	"go-blog/internal/db/store"
	"go-blog/internal/service"
)

type Category struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	Summary      string `json:"summary"`
}

type Tag struct {
	ID      uint   `json:"id"`
	TagName string `json:"tag_name"`
}

type CreateArticle struct {
	Title         string        `json:"title"`
	Summary       string        `json:"summary"`
	Cover         string        `json:"cover"`
	Weight        int           `json:"weight"`
	ArticleBodyID uint          `json:"article_body_id"`
	Category      Category      `json:"category"`
	Tags          []Tag         `json:"tags"`
	Series        series.Series `json:"series"`
	AuthorID      uint          `json:"author_id"`
	ArticleBody   string        `json:"article_body"`
}

type UpdateArticle struct {
	ID            uint          `json:"id" binding:"required"`
	Title         string        `json:"title"`
	Summary       string        `json:"summary"`
	Cover         string        `json:"cover"`
	Weight        int           `json:"weight"`
	ArticleBodyID uint          `json:"article_body_id"`
	Category      Category      `json:"category"`
	Tags          []Tag         `json:"tags"`
	Series        series.Series `json:"series"`
	ArticleBody   string        `json:"article_body"`
}

type DelArticle struct {
	IDs []uint `json:"ids" binding:"required"`
}

type Query struct {
	Keyword string `json:"keyword"`
}

type Controller struct {
	srv service.Service
}

func NewController(store store.Factory) *Controller {
	return &Controller{service.NewService(store)}
}
