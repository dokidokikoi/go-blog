package site

import (
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

type CreateSite struct {
	SiteName string   `json:"site_name" binding:"required"`
	Logo     string   `json:"logo"`
	Summary  string   `json:"summary" binding:"required"`
	Addr     string   `json:"addr"`
	Url      string   `json:"url" binding:"required"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
	Tags     []Tag    `json:"tags" gorm:"many2many:site_tag"`
}

type UpdateSite struct {
	ID       uint     `json:"id" binding:"required"`
	SiteName string   `json:"site_name"`
	Logo     string   `json:"logo"`
	Summary  string   `json:"summary"`
	Addr     string   `json:"addr"`
	Url      string   `json:"url"`
	Category Category `json:"category" gorm:"foreignKey:CategoryID"`
	Tags     []Tag    `json:"tags" gorm:"many2many:site_tag"`
}

type DelSite struct {
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
