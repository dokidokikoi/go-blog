package list

import (
	"go-blog/internal/db/store"
	"go-blog/internal/service"
	"time"
)

type CreateItem struct {
	ItemName       string    `json:"item_name" binding:"required"`
	Total          uint32    `json:"total" binding:"required"`
	Progress       uint32    `json:"progress" binding:"required"`
	Summary        string    `json:"summary" binding:"required"`
	Company        string    `json:"company"`
	Author         string    `json:"author"`
	SerialNumber   string    `json:"serial_number"`
	ProductionDate time.Time `json:"production_date"`
	Type           int8      `json:"type" binding:"required"`
}

type UpdateItem struct {
	ID             uint      `json:"id" binding:"required"`
	ItemName       string    `json:"item_name"`
	Total          uint32    `json:"total"`
	Progress       uint32    `json:"progress"`
	Summary        string    `json:"summary"`
	Company        string    `json:"company"`
	Author         string    `json:"author"`
	SerialNumber   string    `json:"serial_number"`
	ProductionDate time.Time `json:"production_date"`
}

type DelItem struct {
	IDs []uint `json:"ids" binding:"required"`
}

type Query struct {
	Keyword string `form:"keyword"`
	Type    int8   `form:"type"`
}

type Controller struct {
	srv service.Service
}

func NewController(store store.Factory) *Controller {
	return &Controller{service.NewService(store)}
}
