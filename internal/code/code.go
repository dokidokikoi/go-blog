package code

// 基本错误码 1xxxxx
const (
	ErrBind int = iota + 100001
)

// 业务错误码 1xx1xx
const (
	ErrArticleNotFound int = iota + 100101
)

// 数据库错误码  1xx2xx
const (
	ErrDatabase int = iota + 100201
)
