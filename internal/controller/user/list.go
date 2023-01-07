package user

import (
	"go-blog/internal/code"
	"go-blog/internal/core"
	myErrors "go-blog/internal/errors"
	"go-blog/pkg/log/zap"
	"go-blog/pkg/query"

	"github.com/gin-gonic/gin"
)

type Query struct {
	Keyword string `form:"keyword"`
}

func (c *Controller) List(ctx *gin.Context) {
	var q Query
	var err error
	if err = ctx.ShouldBindQuery(&q); err != nil {
		zap.Suger().Errorf("err: %v", err)
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	var pageQuery query.PageQuery
	if err = ctx.ShouldBindQuery(&pageQuery); err != nil {
		zap.Suger().Errorf("err: %v", err)
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}
	option := pageQuery.GetListOption()
	option.GetOption.Preload = append(option.GetOption.Preload, []interface{}{"Role"})

	res, total, err := c.srv.User().List(ctx, q.Keyword, option)
	if err != nil {
		core.WriteResponse(ctx, myErrors.ClientFailed("用户未找到", code.ErrUserNotFound), nil)
		return
	}

	core.WriteListResponse(ctx, nil, total, res)
}
