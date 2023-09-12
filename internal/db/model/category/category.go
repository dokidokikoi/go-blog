package category

const (
	CATE_TYPE_ARTICLE = 1
	CATE_TYPE_SITE    = 2
)

// 文章分类
type Category struct {
	ID           uint   `json:"id" gorm:"primarykey" form:"id"`
	CategoryName string `json:"category_name" gorm:"uniqueIndex:uni_category_name" binding:"required"`
	Summary      string `json:"summary"`
	Type         int    `json:"-" gorm:"default:1"`
}

func (c *Category) TableName() string {
	return "categories"
}
