package article

import "gorm.io/datatypes"

type ArticleBody struct {
	Id        uint              `json:"id" gorm:"primarykey"`
	Content   string            `json:"content" gorm:"type:text"`
	ArticleId uint              `json:"articleId" gorm:"Foreignkey:"`
	Catalog   datatypes.JSONMap `json:"catalog"`
}
