package link

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	LinkName  string         `json:"link_name"`
	Avatar    string         `json:"avatar"`
	Summary   string         `json:"summary"`
	Url       string         `json:"url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
