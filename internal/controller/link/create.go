package link

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/link"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Create(ctx *gin.Context) {
	var input CreateLink
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	err := c.srv.Link().Create(
		ctx,
		&link.Link{
			LinkName: input.LinkName,
			Summary:  input.Summary,
			Avatar:   input.Avatar,
			Url:      input.Url,
		},
		nil,
	)
	if err != nil {
		zaplog.L().Error("创建友链失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
