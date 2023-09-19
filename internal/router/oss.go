package router

import (
	"go-blog/internal/controller/oss"
	"go-blog/internal/db/store"

	"github.com/gin-gonic/gin"
)

func installOssIam(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	e.POST("/upload", oss.NewController(storeFactory).Upload)
}

func init() {
	installs = append(installs, installOssIam)
}
