package article

import (
	"go-blog/internal/code"
	"go-blog/internal/core"
	myErrors "go-blog/internal/errors"
	"go-blog/pkg/log/zap"
	"go-blog/pkg/query"

	"github.com/gin-gonic/gin"
)

type Query struct {
	Keyword string `form:"keyword"`
}

func (c Controller) List(ctx *gin.Context) {
	var q Query
	var err error
	if err = ctx.ShouldBindQuery(&q); err != nil {
		zap.Suger().Errorf("err: %v", err)
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	var pageQuery query.PageQuery
	if err = ctx.ShouldBindQuery(&pageQuery); err != nil {
		zap.Suger().Errorf("err: %v", err)
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}
	option := pageQuery.GetListOption()
	// category := &article.Category{}
	// tag := &article.Tag{}
	// series := &article.Series{}
	option.GetOption.Preload = append(option.GetOption.Preload, "Category", "Tags", "Series")

	res, total, err := c.srv.Article().List(ctx, q.Keyword, option)
	if err != nil {
		core.WriteResponse(ctx, myErrors.ClientFailed("文章未找到", code.ErrArticleNotFound), nil)
		return
	}

	core.WriteListResponse(ctx, nil, total, res)
}
