package tag

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/tag"

	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Create(ctx *gin.Context) {
	var input CreateTag
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}
	_, err := c.srv.Tag().Get(ctx, &tag.Tag{TagName: input.TagName, Type: int8(input.Type)}, nil)
	if err != nil {
		core.WriteResponse(ctx, nil, nil)
		return
	}

	err = c.srv.Tag().Create(ctx, &tag.Tag{TagName: input.TagName, Type: int8(input.Type)}, nil)
	if err != nil {
		zaplog.L().Error("创建分类失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
