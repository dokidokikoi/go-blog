package model

import "time"

type Model struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"delete_at,omitempty"`
}
