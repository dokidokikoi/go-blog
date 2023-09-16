package router

import (
	"go-blog/internal/controller/link"
	"go-blog/internal/db/store"

	"github.com/gin-gonic/gin"
)

func installLink(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	linkController := link.NewController(storeFactory)
	linksR := e.Group("/links")
	{
		linksR.GET("", linkController.List)
		linksR.GET("/:id", linkController.Get)
	}
}

func installLinkIam(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	linkController := link.NewController(storeFactory)
	linksR := e.Group("/links")
	{
		linksR.POST("", linkController.Create)
		linksR.PATCH("", linkController.Update)
		linksR.DELETE("", linkController.Del)
	}
}

func init() {
	installs = append(installs, installLink)
	installsIam = append(installsIam, installLinkIam)
}
