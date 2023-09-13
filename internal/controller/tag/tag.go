package tag

import (
	"go-blog/internal/db/store"
	"go-blog/internal/service"
)

type CreateTag struct {
	TagName string `json:"tag_name" binding:"required"`
	Type    int8   `json:"type" binding:"required"`
}

type UpdateTag struct {
	ID      uint   `json:"id" binding:"required"`
	TagName string `json:"tag_name"`
}

type DelTag struct {
	IDs []uint `json:"ids" binding:"required"`
}

type Query struct {
	Keyword string `json:"keyword"`
	Type    int8   `json:"type" binding:"required"`
}

type Controller struct {
	srv service.Service
}

func NewController(store store.Factory) *Controller {
	return &Controller{service.NewService(store)}
}
