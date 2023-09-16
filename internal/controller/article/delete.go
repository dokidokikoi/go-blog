package article

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/article"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Del(ctx *gin.Context) {
	var input DelArticle
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	ids := []*article.Article{}
	for _, id := range input.IDs {
		ids = append(ids, &article.Article{ID: id})
	}
	errs := c.srv.Article().DeleteCollection(ctx, ids, nil)
	if errs != nil {
		zaplog.L().Error("批量删除文章失败", zap.Errors("errs", errs))
		core.WriteResponse(ctx, myErrors.ApiErrDatabasDel, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
