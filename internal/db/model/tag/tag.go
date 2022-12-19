package tag

const (
	Article int = iota + 1
	Site
)

// 文章标签
type Tag struct {
	Id      uint   `json:"id" gorm:"primarykey"`
	TagName string `json:"tag_name"`
	Type    int    `json:"-"`
}
