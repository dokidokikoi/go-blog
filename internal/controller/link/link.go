package link

import (
	"go-blog/internal/db/store"
	"go-blog/internal/service"
)

type CreateLink struct {
	LinkName string `json:"link_name" binding:"required"`
	Avatar   string `json:"avatar"`
	Summary  string `json:"summary"`
	Url      string `json:"url" binding:"required"`
}

type UpdateLink struct {
	ID       uint   `json:"id" binding:"required"`
	LinkName string `json:"link_name"`
	Avatar   string `json:"avatar"`
	Summary  string `json:"summary"`
	Url      string `json:"url"`
}

type DelLink struct {
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
