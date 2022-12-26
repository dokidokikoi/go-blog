package menu

// 菜单
type Menu struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	MenuName string `json:"menu_name" gorm:"uniqueIndex:uni_menu_name"`
	MenuPath string `json:"path" gorm:"uniqueIndex:uni_menu_path"`
	ParentId uint   `json:"parent_id"`
}
