package category

import (
	"go-blog/internal/db/store"
	"go-blog/internal/service"
)

type CreateCategory struct {
	CategoryName string `json:"category_name" binding:"required"`
	Summary      string `json:"summary" binding:"required"`
	Type         int8   `json:"type" binding:"required"`
}

type UpdateCategory struct {
	ID           uint   `json:"id" binding:"required"`
	CategoryName string `json:"category_name"`
	Summary      string `json:"summary"`
}

type DelCategory struct {
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
