package comment

import (
	"go-blog/internal/db/store"
	"go-blog/internal/service"
)

type CreateComment struct {
	ArticleID uint   `json:"article_id"`
	PID       uint   `json:"pid"`
	Content   string `json:"content" binding:"required"`
	Name      string `json:"name" binding:"required"`
}

type DelComment struct {
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
