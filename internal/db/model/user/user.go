package user

import (
	"go-blog/internal/db/model"
	"time"
)

// 用户
type User struct {
	Id         uint      `json:"id" gorm:"primarykey"`
	Account    string    `json:"account" gorm:"unique"`
	Avatar     string    `json:"avatar"`
	Email      string    `json:"email" gorm:"unique"`
	NickName   string    `json:"nick_name"`
	Password   string    `json:"password,omitempty"`
	StatusCode int       `json:"status_code"`
	LastLogin  time.Time `json:"last_login"`
	model.Model
}
