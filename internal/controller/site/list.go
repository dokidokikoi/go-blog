package site

import (
	"fmt"
	"go-blog/internal/core"
	"go-blog/internal/db/model/site"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	meta "github.com/dokidokikoi/go-common/meta/option"
	"github.com/dokidokikoi/go-common/query"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) List(ctx *gin.Context) {
	var input Query
	if ctx.ShouldBindQuery(&input) != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}
	var pageQuery query.PageQuery
	if ctx.ShouldBindQuery(&pageQuery) != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}

	listOption := pageQuery.GetListOption()
	var (
		items   []*site.Site
		total   int64
		err     error
		example = &site.Site{}
	)
	if input.Keyword != "" {
		node := &meta.WhereNode{
			Conditions: []*meta.Condition{
				{
					Field:    "site_name",
					Operator: meta.LIKE,
					Value:    fmt.Sprintf("%%%s%%", input.Keyword),
				},
			},
		}
		items, total, err = c.srv.Site().ListByWhereNode(ctx, example, node, listOption)
	} else {
		items, total, err = c.srv.Site().List(ctx, example, listOption)
	}
	if err != nil {
		zaplog.L().Error("获取网站列表失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiRecordNotFound, "")
		return
	}

	core.WriteListResponse(ctx, nil, total, items)
}
