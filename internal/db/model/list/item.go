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
	gorm.Model
	ItemName       string    `json:"item_name"`
	Total          uint32    `json:"total"`
	Progress       uint32    `json:"progress"`
	Summary        string    `json:"summary"`
	Company        string    `json:"company"`
	Author         string    `json:"author"`
	SerialNumber   string    `json:"serial_number"`
	ProductionDate time.Time `json:"production_date"`
	Type           string    `json:"type"`
}
