package tag

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/tag"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Del(ctx *gin.Context) {
	var input DelTag
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	ids := []*tag.Tag{}
	for _, id := range input.IDs {
		ids = append(ids, &tag.Tag{ID: id})
	}

	errs := c.srv.Tag().DeleteCollection(ctx, ids, nil)
	if errs != nil {
		zaplog.L().Error("批量删除标签失败", zap.Errors("errs", errs))
		core.WriteResponse(ctx, myErrors.ApiErrDatabasDel, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
