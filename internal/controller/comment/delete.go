package comment

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/comment"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Del(ctx *gin.Context) {
	var input DelComment
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	ids := []*comment.Comment{}
	for _, id := range input.IDs {
		ids = append(ids, &comment.Comment{ID: id})
	}
	errs := c.srv.Comment().DeleteCollection(ctx, ids, nil)
	if errs != nil {
		zaplog.L().Error("批量删除评论失败", zap.Errors("errs", errs))
		core.WriteResponse(ctx, myErrors.ApiErrDatabasDel, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
