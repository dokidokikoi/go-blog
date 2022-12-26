package errors

import (
	"errors"
	"go-blog/internal/code"
)

var (
	RecordNotFound = errors.New("该记录未找到")
	NoUpdateRows   = errors.New("无更新记录")
	NameDuplicate  = errors.New("名称重复")
	ErrValidation  = errors.New("参数格式错误")
	ErrDatabaseOp  = errors.New("录入失败")
)

var (
	ApiErrValidation = ClientFailed(ErrValidation.Error(), code.ErrBind)
	ApiErrDatabase   = ServerFailed(ErrDatabaseOp.Error(), code.ErrDatabase)
)
