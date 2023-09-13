package series

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/series"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Update(ctx *gin.Context) {
	var input UpdateSeries
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	err := c.srv.Series().Update(ctx, &series.Series{ID: input.ID, SeriesName: input.SeriesName, Summary: input.Summary}, nil)
	if err != nil {
		zaplog.L().Error("更新系列失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
