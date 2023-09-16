package router

import (
	"go-blog/internal/middleware"

	"github.com/gin-gonic/gin"
)

type installFunc func(r *gin.Engine)

var installs []installFunc
var installsIam []installFunc

func InstallAll(r *gin.Engine) {
	for _, f := range installs {
		f(r)
	}

	r.Use(middleware.Auth())
	for _, f := range installsIam {
		f(r)
	}
}
