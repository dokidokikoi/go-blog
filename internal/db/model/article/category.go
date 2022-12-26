package article

// 文章分类
type Category struct {
	ID           uint   `json:"id" gorm:"primarykey" form:"id"`
	CategoryName string `json:"category_name" gorm:"uniqueIndex:uni_category_name" binding:"required"`
	Summary      string `json:"summary"`
}

func (c *Category) TableName() string {
	return "article_categories"
}