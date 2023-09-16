package router

import (
	"go-blog/internal/controller/list"
	"go-blog/internal/db/store"

	"github.com/gin-gonic/gin"
)

func installList(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	listController := list.NewController(storeFactory)
	linksR := e.Group("/list")
	{
		linksR.GET("", listController.List)
		linksR.GET("/:id", listController.Get)
	}
}

func installListIam(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	listController := list.NewController(storeFactory)
	linksR := e.Group("/list")
	{
		linksR.POST("", listController.Create)
		linksR.PATCH("", listController.Update)
		linksR.DELETE("", listController.Del)
	}
}

func init() {
	installs = append(installs, installList)
	installsIam = append(installsIam, installListIam)
}
