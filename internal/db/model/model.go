package model

import "time"

type Model struct {
	Id        uint      `gorm:"primarykey" form:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"delete_at,omitempty"`
}
