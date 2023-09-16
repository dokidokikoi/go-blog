package article

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/article"
	myErrors "go-blog/internal/errors"
	"strconv"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	meta "github.com/dokidokikoi/go-common/meta/option"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Get(ctx *gin.Context) {
	articleID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		zaplog.L().Error("参数校验失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	option := &meta.GetOption{Preload: []string{"Category", "Tags", "Series", "ArticleBody"}}
	s, err := c.srv.Article().Get(ctx, &article.Article{ID: uint(articleID)}, option)
	if err != nil {
		zaplog.L().Error("获取文章信息失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiRecordNotFound, nil)
		return
	}
	core.WriteResponse(ctx, nil, s)
}
