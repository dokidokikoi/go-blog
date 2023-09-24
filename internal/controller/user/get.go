package user

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/user"
	myErrors "go-blog/internal/errors"
	"strconv"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Get(ctx *gin.Context) {
	userID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		zaplog.L().Error("参数校验失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	u, err := c.srv.User().Get(ctx, &user.User{ID: uint(userID)}, nil)
	if err != nil {
		zaplog.L().Error("获取用户信息失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiRecordNotFound, nil)
		return
	}
	core.WriteResponse(ctx, nil, u)
}

func (c *Controller) Host(ctx *gin.Context) {
	u, err := c.srv.User().Get(ctx, &user.User{ID: uint(1)}, nil)
	if err != nil {
		zaplog.L().Error("获取站长信息失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiRecordNotFound, nil)
		return
	}
	core.WriteResponse(ctx, nil, u)
}
