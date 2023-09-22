package comment

import (
	"go-blog/internal/db/store"
	"go-blog/internal/service"
)

type CreateComment struct {
	ArticleID  uint   `json:"article_id"`
	PID        uint   `json:"pid"`
	Content    string `json:"content" binding:"required"`
	Nickname   string `json:"nickname" binding:"required"`
	ToNickname string `json:"to_nickname"`
}

type UpdateComment struct {
	ID     uint `json:"id" binding:"required"`
	Weight int8 `json:"weight"`
}

type DelComment struct {
	IDs []uint `json:"ids" binding:"required"`
}

type Query struct {
	Keyword   string `form:"keyword"`
	ArticleID uint   `form:"article_id"`
}

type Controller struct {
	srv service.Service
}

func NewController(store store.Factory) *Controller {
	return &Controller{service.NewService(store)}
}
