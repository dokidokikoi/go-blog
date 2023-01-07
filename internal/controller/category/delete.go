package category

import (
	"go-blog/internal/core"
	myErrors "go-blog/internal/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	if c.srv.Category().DeleteById(ctx, uint(id)) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}

	core.WriteResponse(ctx, myErrors.Success("删除成功"), nil)
}

func (c *Controller) DeleteByIds(ctx *gin.Context) {
	var input DeleteIds
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	if c.srv.Category().DeleteByIds(ctx, input.Ids) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}

	core.WriteResponse(ctx, myErrors.Success("删除成功"), nil)
}
