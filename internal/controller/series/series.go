package series

import (
	"go-blog/internal/db/store"
	"go-blog/internal/service"
)

type CreateSeries struct {
	SeriesName string `json:"series_name" binding:"required"`
	Summary    string `json:"summary" binding:"required"`
}

type UpdateSeries struct {
	ID         uint   `json:"id" binding:"required"`
	SeriesName string `json:"series_name"`
	Summary    string `json:"summary"`
}

type DelSeries struct {
	IDs []uint `json:"ids" binding:"required"`
}

type Query struct {
	Keyword string `form:"keyword"`
}

type Controller struct {
	srv service.Service
}

func NewController(store store.Factory) *Controller {
	return &Controller{service.NewService(store)}
}
