package category

import (
	"go-blog/internal/code"
	"go-blog/internal/core"
	myErrors "go-blog/internal/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Get(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	cate, err := c.srv.Category().GetById(ctx, uint(id), nil)
	if err != nil {
		core.WriteResponse(ctx, myErrors.ClientFailed("分类未找到", code.ErrCategoryNotFound), nil)
		return
	}

	core.WriteResponse(ctx, nil, cate)
}
