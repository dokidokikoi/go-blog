package article

type Series struct {
	ID         uint   `json:"id" gorm:"primarykey" form:"id"`
	SeriesName string `json:"series_name" gorm:"uniqueIndex:uni_series_name" binding:"required"`
	Summary    string `json:"summary"`
}
