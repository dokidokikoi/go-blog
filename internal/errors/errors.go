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
	ErrSystemErr   = errors.New("系统错误")
	ErrPassword    = errors.New("密码错误")
	ErrCaptcha     = errors.New("验证码错误")
	ErrDatabasDel  = errors.New("删除失败")
)

var (
	ApiErrValidation  = ClientFailed(ErrValidation.Error(), code.ErrBind)
	ApiErrDatabase    = ServerFailed(ErrDatabaseOp.Error(), code.ErrDatabase)
	ApiRecordNotFound = ServerFailed(RecordNotFound.Error(), code.ErrRecordNotFound)
	ApiErrSystemErr   = ServerFailed(ErrSystemErr.Error(), code.ErrSystemErr)
	ApiErrPassword    = ClientFailed(ErrPassword.Error(), code.ErrPassword)
	ApiErrCaptcha     = ClientFailed(ErrCaptcha.Error(), code.ErrCaptcha)
	ApiErrDatabasDel  = ServerFailed(ErrDatabasDel.Error(), code.ErrDatabasDel)
	ApiNoUpdateRows   = ServerFailed(NoUpdateRows.Error(), code.NoUpdateRows)
)
