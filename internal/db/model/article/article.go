package article

import (
	"go-blog/internal/db/model/category"
	"go-blog/internal/db/model/user"

	"gorm.io/gorm"
)

// 文章
type Article struct {
	gorm.Model
	Title         string            `json:"title"`
	Summary       string            `json:"summary"`
	Cover         string            `json:"cover"`
	ViewCounts    uint              `json:"view_counts" gorm:"default:0"`
	CommentCounts uint              `json:"comment_counts" gorm:"default:0"`
	Weight        int               `json:"weight,omitempty" gorm:"default:0"`
	ArticleBodyID uint              `json:"article_body_id"`
	StatusCode    int               `json:"status_code,omitempty" gorm:"default:0"`
	CategoryID    uint              `json:"category_id"`
	Category      category.Category `json:"category" gorm:"foreignKey:CategoryID"`
	Tags          []Tag             `json:"tags" gorm:"many2many:article_article_tag"` //多对多关系.
	SeriesID      uint              `json:"series_id"`
	Series        Series            `json:"series" gorm:"foreignKey:SeriesID"`
	AuthorID      uint              `json:"author_id"`
	Author        user.User         `json:"author" gorm:"foreignKey:AuthorID"`
	ArticleBody   ArticleBody       `json:"article_body" gorm:"foreignKey:ArticleBodyID"`
}

type ArticleBody struct {
	ID      uint   `json:"id" gorm:"primarykey" form:"id"`
	Content string `json:"content" gorm:"type:text"`
}
