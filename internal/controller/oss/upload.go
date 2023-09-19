package oss

import (
	"fmt"
	"go-blog/internal/config"
	"go-blog/internal/core"
	"go-blog/pkg/oss"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Upload(ctx *gin.Context) {
	cnf := config.OssConfig
	fmt.Println(cnf)
	core.WriteResponse(ctx, nil,
		oss.GetPolicyToken(
			config.OssConfig.ID,
			config.OssConfig.Secret,
			config.OssConfig.Host,
			"",
			config.OssConfig.UploadDir,
			config.OssConfig.ExpireTime))
}
