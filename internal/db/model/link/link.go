package link

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	ID        uint   `gorm:"primarykey"`
	LinkName  string `json:"link_name"`
	Avatar    string `json:"avatar"`
	Summary   string `json:"summary"`
	Url       string `json:"url"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
