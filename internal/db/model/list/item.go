package list

import (
	"time"

	"gorm.io/gorm"
)

const (
	TYPE_ANIME = iota + 1
	TYPE_GAMES
	TYPE_MOVIE
	TYPE_BOOK
)

type Item struct {
	ID             uint           `json:"id" gorm:"primarykey"`
	ItemName       string         `json:"item_name"`
	Cover          string         `json:"cover"`
	Total          uint32         `json:"total"`
	Progress       uint32         `json:"progress"`
	Summary        string         `json:"summary"`
	Company        string         `json:"company"`
	Author         string         `json:"author"`
	SerialNumber   string         `json:"serial_number"`
	ProductionDate time.Time      `json:"production_date"`
	Type           int8           `json:"type"`
	Rate           float32        `json:"rate"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
