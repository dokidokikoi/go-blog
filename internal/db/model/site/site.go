package site

import (
	"go-blog/internal/db/model/category"

	"gorm.io/gorm"
)

// 网站
type Site struct {
	gorm.Model
	SiteName   string            `json:"site_name"`
	Logo       string            `json:"logo"`
	Summary    string            `json:"summary"`
	Addr       string            `json:"addr"`
	CategoryID uint              `json:"category_id"`
	Category   category.Category `json:"category" gorm:"foreignKey:CategoryID"`
}
