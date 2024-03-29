package list

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/list"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Create(ctx *gin.Context) {
	var input CreateItem
	if err := ctx.ShouldBindJSON(&input); err != nil {
		zaplog.L().Error("", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	err := c.srv.Items().Create(
		ctx,
		&list.Item{
			ItemName:       input.ItemName,
			Cover:          input.Cover,
			Summary:        input.Summary,
			Total:          input.Total,
			Progress:       input.Progress,
			Company:        input.Company,
			Author:         input.Author,
			Rate:           input.Rate,
			SerialNumber:   input.SerialNumber,
			ProductionDate: input.ProductionDate,
			Type:           input.Type,
		},
		nil,
	)
	if err != nil {
		zaplog.L().Error("创建项目失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
