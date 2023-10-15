package list

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/list"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Update(ctx *gin.Context) {
	var input UpdateItem
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	err := c.srv.Items().Update(ctx, &list.Item{
		ID:             input.ID,
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
	}, nil)
	if err != nil {
		zaplog.L().Error("更新项目失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
