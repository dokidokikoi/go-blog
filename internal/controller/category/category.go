package category

import (
	"go-blog/internal/db/store"
	"go-blog/internal/service"
)

type CategoryCreateUpdate struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name" binding:"required"`
	Summary      string `json:"summary"`
	Type         int    `json:"type" binding:"required"`
}

type DeleteIds struct {
	Ids []uint `json:"ids" binding:"required"`
}

type Controller struct {
	srv service.Service
}

func NewController(store store.Factory) *Controller {
	return &Controller{service.NewService(store)}
}
