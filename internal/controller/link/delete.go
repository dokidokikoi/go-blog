package link

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/link"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Del(ctx *gin.Context) {
	var input DelLink
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	ids := []*link.Link{}
	for _, id := range input.IDs {
		ids = append(ids, &link.Link{ID: id})
	}
	errs := c.srv.Link().DeleteCollection(ctx, ids, nil)
	if errs != nil {
		zaplog.L().Error("批量删除友链失败", zap.Errors("errs", errs))
		core.WriteResponse(ctx, myErrors.ApiErrDatabasDel, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
