package tag

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/tag"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Update(ctx *gin.Context) {
	var input UpdateTag
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	err := c.srv.Tag().Update(ctx, &tag.Tag{ID: input.ID, TagName: input.TagName}, nil)
	if err != nil {
		zaplog.L().Error("更新标签失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
