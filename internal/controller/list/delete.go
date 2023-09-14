package list

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/list"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Del(ctx *gin.Context) {
	var input DelItem
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	ids := []*list.Item{}
	for _, id := range input.IDs {
		ids = append(ids, &list.Item{ID: id})
	}
	errs := c.srv.Items().DeleteCollection(ctx, ids, nil)
	if errs != nil {
		zaplog.L().Error("批量删除项目失败", zap.Errors("errs", errs))
		core.WriteResponse(ctx, myErrors.ApiErrDatabasDel, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
