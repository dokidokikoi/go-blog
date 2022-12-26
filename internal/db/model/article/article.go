package article

import (
	"gorm.io/gorm"
)

// 文章
type Article struct {
	gorm.Model
	Title         string   `json:"title"`
	Summary       string   `json:"summary"`
	Cover         string   `json:"cover"`
	ViewCounts    uint     `json:"view_counts" gorm:"default:0"`
	CommentCounts uint     `json:"comment_counts" gorm:"default:0"`
	Weight        int      `json:"weight,omitempty" gorm:"default:0"`
	ArticleBodyID uint     `json:"article_body_id" gorm:"uniqueIndex:fk_article_body_id"`
	StatusCode    int      `json:"status_code,omitempty" gorm:"default:0"`
	CategoryID    uint     `json:"category_id" gorm:"uniqueIndex:fk_category_id"`
	Category      Category `json:"article" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tags          []Tag    `json:"tags" gorm:"many2many:article_article_tag"` //多对多关系.
	SeriesID      uint     `json:"series_id" gorm:"uniqueIndex:fk_cseries_id"`
	Series        Series   `json:"series" gorm:"foreignKey:SeriesID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ArticleBody struct {
	ID      uint    `json:"id" gorm:"primarykey" form:"id"`
	Content string  `json:"content" gorm:"type:text"`
	Article Article `json:"article" gorm:"foreignKey:ArticleBodyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
