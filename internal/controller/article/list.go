package article

import (
	"fmt"
	"go-blog/internal/core"
	"go-blog/internal/db/model/article"
	myErrors "go-blog/internal/errors"
	"strconv"

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
	listOption.Preload = []string{"Category", "Tags", "Series"}
	var (
		items   []*article.Article
		total   int64
		err     error
		example = &article.Article{CategoryID: input.CategoryID, SeriesID: input.SeriesID}
	)
	if input.Keyword != "" {
		node := &meta.WhereNode{
			Conditions: []*meta.Condition{
				{
					Field:    "title",
					Operator: meta.LIKE,
					Value:    fmt.Sprintf("%%%s%%", input.Keyword),
				},
			},
		}
		items, total, err = c.srv.Article().ListByWhereNode(ctx, example, node, listOption)
	} else {
		items, total, err = c.srv.Article().List(ctx, example, listOption)
	}
	if err != nil {
		zaplog.L().Error("获取文章列表失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiRecordNotFound, "")
		return
	}

	core.WriteListResponse(ctx, nil, total, items)
}

func (c *Controller) ListTagArticle(ctx *gin.Context) {
	tagID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		zaplog.L().Error("参数校验失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}
	var pageQuery query.PageQuery
	if ctx.ShouldBindQuery(&pageQuery) != nil {
		zaplog.L().Error("参数校验失败")
		core.WriteResponse(ctx, myErrors.ApiErrValidation, "")
		return
	}
	listOption := pageQuery.GetListOption()
	listOption.Preload = []string{"Category", "Tags", "Series"}
	articles, total, err := c.srv.Article().ListTagArticle(ctx, uint(tagID), listOption)
	if err != nil {
		zaplog.L().Error("获取文章列表失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiRecordNotFound, "")
		return
	}
	core.WriteListResponse(ctx, nil, total, articles)
}
