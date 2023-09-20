package oss

import (
	"go-blog/internal/config"
	"go-blog/internal/core"
	"go-blog/pkg/oss"
	"time"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Upload(ctx *gin.Context) {
	dir := time.Now().Format("2006-01-02")

	core.WriteResponse(ctx, nil,
		oss.GetPolicyToken(
			config.OssConfig.ID,
			config.OssConfig.Secret,
			config.OssConfig.Host,
			"",
			dir+"/",
			config.OssConfig.ExpireTime))
}
