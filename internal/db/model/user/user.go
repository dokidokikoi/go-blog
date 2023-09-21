package user

import (
	"time"

	"gorm.io/gorm"
)

// 用户
type User struct {
	ID         uint           `json:"id" gorm:"primarykey"`
	Account    string         `json:"account" gorm:"unique"`
	Avatar     string         `json:"avatar"`
	Email      string         `json:"email" gorm:"unique"`
	NickName   string         `json:"nick_name"`
	Password   string         `json:"-"`
	StatusCode int            `json:"status_code" gorm:"fefault:0"`
	LastLogin  time.Time      `json:"last_login"`
	RoleID     uint           `json:"role_id"`
	Role       Role           `json:"role" gorm:"foreignKey:RoleID"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

// 角色
type Role struct {
	gorm.Model
	RoleName string `json:"role_name" gorm:"uniqueIndex:uni_role_name"`
}
