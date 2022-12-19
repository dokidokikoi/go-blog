package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InstallAll(e *gin.Engine) {
	e.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, struct{ Msg string }{Msg: "hello"})
	})
}
