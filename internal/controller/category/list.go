package category

import (
	"go-blog/internal/code"
	"go-blog/internal/core"
	myErrors "go-blog/internal/errors"
	"go-blog/pkg/query"

	"github.com/gin-gonic/gin"
)

type Query struct {
	Keyword string `form:"keyword"`
	Type    int    `form:"type"`
}

func (c *Controller) List(ctx *gin.Context) {
	var input query.PageQuery
	if ctx.ShouldBindQuery(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	var keyword Query
	if ctx.ShouldBindQuery(&keyword) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	categorys, total, err := c.srv.Category().List(ctx, keyword.Keyword, keyword.Type, input.GetListOption())
	if err != nil {
		core.WriteResponse(ctx, myErrors.ClientFailed("分类未找到", code.ErrCategoryNotFound), nil)
		return
	}

	core.WriteListResponse(ctx, nil, total, categorys)
}
