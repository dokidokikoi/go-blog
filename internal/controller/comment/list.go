package comment

import (
	"fmt"
	"go-blog/internal/core"
	"go-blog/internal/db/model/comment"
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
		items   []*comment.Comment
		total   int64
		err     error
		example = &comment.Comment{ArticleID: input.ArticleID, PID: 0}
	)
	root := &meta.WhereNode{
		Conditions: []*meta.Condition{
			{
				Field:    "article_id",
				Operator: meta.EQUAL,
				Value:    input.ArticleID,
			},
		},
	}
	root.Next = &meta.WhereNode{
		Conditions: []*meta.Condition{
			{
				Field:    "pid",
				Operator: meta.EQUAL,
				Value:    0,
			},
		},
	}
	if input.Keyword != "" {
		node := &meta.WhereNode{
			Conditions: []*meta.Condition{
				{
					Field:    "content",
					Operator: meta.LIKE,
					Value:    fmt.Sprintf("%%%s%%", input.Keyword),
				},
			},
		}
		root.Next.Next = node
	}
	items, total, err = c.srv.Comment().ListByWhereNode(ctx, example, root, listOption)
	if err != nil {
		zaplog.L().Error("获取评论列表失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiRecordNotFound, "")
		return
	}

	errs := c.srv.Comment().SetCommentChildren(ctx, items)
	if errs != nil {
		zaplog.L().Error("获取评论列表子评论失败", zap.Errors("errs", errs))
	}

	core.WriteListResponse(ctx, nil, total, items)
}
