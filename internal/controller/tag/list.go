package tag

import (
	"fmt"
	"go-blog/internal/core"
	"go-blog/internal/db/model/tag"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	meta "github.com/dokidokikoi/go-common/meta/option"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) List(ctx *gin.Context) {
	var query Query
	if ctx.ShouldBindJSON(&query) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	listOption := &meta.ListOption{}
	var (
		list    []*tag.Tag
		total   int64
		err     error
		example = &tag.Tag{Type: query.Type}
	)
	if query.Keyword != "" {
		node := &meta.WhereNode{
			Conditions: []*meta.Condition{
				{
					Field:    "tag_name",
					Operator: meta.LIKE,
					Value:    fmt.Sprintf("%%%s%%", query.Keyword),
				},
			},
		}
		list, total, err = c.srv.Tag().ListByWhereNode(ctx, example, node, listOption)
	} else {
		list, total, err = c.srv.Tag().List(ctx, example, listOption)
	}
	if err != nil {
		zaplog.L().Error("获取标签列表失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiRecordNotFound, "")
		return
	}

	core.WriteListResponse(ctx, nil, total, list)
}
