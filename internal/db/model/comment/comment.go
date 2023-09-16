package comment

import (
	"time"

	"gorm.io/gorm"
)

// 评论
type Comment struct {
	ID        uint   `json:"id" grom:"primaryKey"`
	ArticleID uint   `json:"article_id" gorm:"default:0;"`
	PID       uint   `json:"pid" gorm:"default:0;"`
	Content   string `json:"content"`
	Name      string `json:"name"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
