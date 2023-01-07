package list

import (
	"go-blog/internal/db/model/category"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ItemName       string            `json:"item_name"`
	Total          uint32            `json:"total"`
	Progress       uint32            `json:"progress"`
	Summary        string            `json:"summary"`
	Company        string            `json:"company"`
	Author         string            `json:"author"`
	SerialNumber   string            `json:"serial_number"`
	ProductionDate time.Time         `json:"production_date"`
	CategoryID     uint              `json:"category_id"`
	Category       category.Category `json:"category" gorm:"foreignKey:CategoryID"`
}
