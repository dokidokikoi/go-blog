package list

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/list"
	myErrors "go-blog/internal/errors"
	"strconv"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Get(ctx *gin.Context) {
	itemID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		zaplog.L().Error("参数校验失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	l, err := c.srv.Items().Get(ctx, &list.Item{ID: uint(itemID)}, nil)
	if err != nil {
		zaplog.L().Error("获取项目信息失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiRecordNotFound, nil)
		return
	}
	core.WriteResponse(ctx, nil, l)
}
