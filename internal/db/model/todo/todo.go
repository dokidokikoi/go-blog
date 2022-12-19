package todo

import "time"

// 待办
type Todo struct {
	Id         uint      `json:"id" gorm:"primarykey"`
	Content    string    `json:"content"`
	StatusCode int       `josn:"status_code"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
}
