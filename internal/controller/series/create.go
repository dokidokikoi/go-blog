package series

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/series"

	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Create(ctx *gin.Context) {
	var input CreateSeries
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	_, err := c.srv.Series().Get(ctx, &series.Series{SeriesName: input.SeriesName}, nil)
	if err == nil {
		core.WriteResponse(ctx, nil, nil)
		return
	}

	err = c.srv.Series().Create(
		ctx,
		&series.Series{
			SeriesName: input.SeriesName,
			Summary:    input.Summary,
		},
		nil,
	)
	if err != nil {
		zaplog.L().Error("创建系列失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
