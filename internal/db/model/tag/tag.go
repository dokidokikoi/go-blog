package tag

const (
	TAG_TYPE_ARTICLE = 1 + iota
	TAG_TYPE_SITE
)

type Tag struct {
	ID      uint   `json:"id" gorm:"primarykey" form:"id"`
	TagName string `json:"tag_name" gorm:"uniqueIndex:uni_tag_name" binding:"required"`
	Type    int8   `json:"type" gorm:"default:0"`
}

func (tag *Tag) TableName() string {
	return "tags"
}
