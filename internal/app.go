package application

import (
	"go-blog/internal/db/store/data"
	"go-blog/internal/router"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func initTask() {
	data.SetStoreDBFactory()
}

func Run() {
	initTask()
	e := gin.Default()
	e.Use(gin.Recovery())
	router.InstallAll(e)

	endless.ListenAndServe(":9080", e)
}
