package site

import (
	"go-blog/internal/db/model/category"
	"go-blog/internal/db/model/tag"
	"time"

	"gorm.io/gorm"
)

// 网站
type Site struct {
	gorm.Model
	SiteName   string            `json:"site_name"`
	Logo       string            `json:"logo"`
	Summary    string            `json:"summary"`
	Addr       string            `json:"addr"`
	Url        string            `json:"url"`
	CategoryID uint              `json:"category_id"`
	Category   category.Category `json:"category" gorm:"foreignKey:CategoryID"`
	Tags       []tag.Tag         `json:"tags" gorm:"many2many:site_tag"` //多对多关系.
}

// 网站_标签中间表
type SiteTag struct {
	SiteID    uint      `json:"site_id" gorm:"uniqueIndex:uni_site_tag"`
	TagID     uint      `json:"tag_id" gorm:"uniqueIndex:uni_site_tag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (siteTag *SiteTag) TableName() string {
	return "site_tag"
}
