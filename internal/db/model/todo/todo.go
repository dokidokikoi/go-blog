package todo

import "time"

// 待办
type Todo struct {
	ID         uint      `json:"id" gorm:"primarykey" form:"id"`
	Content    string    `json:"content"`
	StatusCode int       `josn:"status_code" gorm:"default:0"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
}
