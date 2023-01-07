package article

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

	if err := c.srv.Article().DeleteArticleByID(ctx, uint(id)); err != nil {
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}

	core.WriteResponse(ctx, myErrors.Success("删除成功"), nil)
}
