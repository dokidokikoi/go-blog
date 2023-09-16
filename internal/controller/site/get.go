package site

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/site"
	myErrors "go-blog/internal/errors"
	"strconv"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	meta "github.com/dokidokikoi/go-common/meta/option"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Get(ctx *gin.Context) {
	siteID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		zaplog.L().Error("参数校验失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	option := &meta.GetOption{Preload: []string{"Category", "Tags"}}
	s, err := c.srv.Site().Get(ctx, &site.Site{ID: uint(siteID)}, option)
	if err != nil {
		zaplog.L().Error("获取网站信息失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiRecordNotFound, nil)
		return
	}
	core.WriteResponse(ctx, nil, s)
}
