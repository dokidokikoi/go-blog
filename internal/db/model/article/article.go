package article

import (
	"go-blog/internal/db/model/category"
	"go-blog/internal/db/model/series"
	"go-blog/internal/db/model/tag"
	"go-blog/internal/db/model/user"
	"time"

	"gorm.io/gorm"
)

// 文章
type Article struct {
	ID            uint              `json:"id" gorm:"primarykey"`
	Title         string            `json:"title"`
	Summary       string            `json:"summary"`
	Cover         string            `json:"cover"`
	ViewCounts    uint              `json:"view_counts" gorm:"default:0"`
	CommentCounts uint              `json:"comment_counts" gorm:"default:0"`
	Weight        int               `json:"weight" gorm:"default:1"`
	ArticleBodyID uint              `json:"article_body_id"`
	StatusCode    int               `json:"status_code,omitempty" gorm:"default:1"`
	CategoryID    uint              `json:"category_id"`
	Category      category.Category `json:"category" gorm:"foreignKey:CategoryID"`
	Tags          []tag.Tag         `json:"tags" gorm:"many2many:article_tag"` //多对多关系.
	SeriesID      uint              `json:"series_id" gorm:"default:null"`
	Series        series.Series     `json:"series" gorm:"foreignKey:SeriesID"`
	AuthorID      uint              `json:"author_id"`
	Author        user.User         `json:"author" gorm:"foreignKey:AuthorID"`
	ArticleBody   ArticleBody       `json:"article_body" gorm:"foreignKey:ArticleBodyID"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	DeletedAt     gorm.DeletedAt    `json:"deleted_at" gorm:"index"`
}

type ArticleBody struct {
	ID      uint   `json:"id" gorm:"primarykey" form:"id"`
	Content string `json:"content" gorm:"type:text"`
}

// 文章_标签中间表
type ArticleTag struct {
	ArticleID uint `json:"article_id" gorm:"uniqueIndex:uni_article_tag"`
	TagID     uint `json:"tag_id" gorm:"uniqueIndex:uni_article_tag"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}

func (articleTag *ArticleTag) TableName() string {
	return "article_tag"
}
