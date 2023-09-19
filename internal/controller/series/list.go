package series

import (
	"fmt"
	"go-blog/internal/core"
	"go-blog/internal/db/model/series"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	meta "github.com/dokidokikoi/go-common/meta/option"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) List(ctx *gin.Context) {
	var query Query
	if ctx.ShouldBindQuery(&query) != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	listOption := &meta.ListOption{PageSize: 1000}
	var (
		list    []*series.Series
		total   int64
		err     error
		example = &series.Series{}
	)
	if query.Keyword != "" {
		node := &meta.WhereNode{
			Conditions: []*meta.Condition{
				{
					Field:    "series_name",
					Operator: meta.LIKE,
					Value:    fmt.Sprintf("%%%s%%", query.Keyword),
				},
			},
		}
		list, total, err = c.srv.Series().ListByWhereNode(ctx, example, node, listOption)
	} else {
		list, total, err = c.srv.Series().List(ctx, example, listOption)
	}
	if err != nil {
		zaplog.L().Error("获取系列列表失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiRecordNotFound, "")
		return
	}

	core.WriteListResponse(ctx, nil, total, list)
}
