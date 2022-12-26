package article

import "time"

// 文章标签
type Tag struct {
	ID      uint   `json:"id" gorm:"primarykey" form:"id"`
	TagName string `json:"tag_name" gorm:"uniqueIndex:uni_tag_name" binding:"required"`
}

func (tag *Tag) TableName() string {
	return "article_tags"
}

// 文章_标签中间表
type ArticleTag struct {
	ArticleId uint      `json:"article_id" gorm:"uniqueIndex:uni_article_tag"`
	TagId     uint      `json:"tag_id" gorm:"uniqueIndex:uni_article_tag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (articleTag *ArticleTag) TableName() string {
	return "article_article_tag"
}
