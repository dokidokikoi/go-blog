package article

// 文章
type Article struct {
	Id            uint   `json:"id"`
	Title         string `json:"title"`
	Summary       string `json:"summary"`
	Cover         string `json:"cover"`
	ViewCounts    uint   `json:"view_counts"`
	CommentCounts uint   `json:"comment_counts"`
	Weight        int    `json:"weight,omitempty"`
	StatusCode    int    `json:"status_code,omitempty"`
}
