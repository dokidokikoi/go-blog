package user

import (
	"time"

	"gorm.io/gorm"
)

// 用户
type User struct {
	gorm.Model
	Account    string    `json:"account" gorm:"unique"`
	Avatar     string    `json:"avatar"`
	Email      string    `json:"email" gorm:"unique"`
	NickName   string    `json:"nick_name"`
	Password   string    `json:"password,omitempty"`
	StatusCode int       `json:"status_code"`
	LastLogin  time.Time `json:"last_login"`
	RoleID     uint      `json:"role_id"`
	Role       uint      `json:"role" gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// 角色
type Role struct {
	gorm.Model
	RoleName string `json:"role_name" gorm:"uniqueIndex:uni_role_name"`
}
