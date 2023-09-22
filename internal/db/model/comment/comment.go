package comment

import (
	"time"

	"gorm.io/gorm"
)

// 评论
type Comment struct {
	ID         uint           `json:"id" grom:"primaryKey"`
	ArticleID  uint           `json:"article_id" gorm:"default:0;"`
	PID        uint           `json:"pid" gorm:"column:pid;default:0;"`
	Content    string         `json:"content"`
	Nickname   string         `json:"nickname"`
	ToNickname string         `json:"to_nickname"`
	Children   []*Comment     `json:"children" gorm:"-"`
	Weight     int8           `json:"weight" gorm:"default:1;"`
	CreatedAt  time.Time      `json:"created_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
