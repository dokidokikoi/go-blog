package router

import (
	"go-blog/internal/controller/tag"
	"go-blog/internal/db/store"

	"github.com/gin-gonic/gin"
)

func installTag(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	tagController := tag.NewController(storeFactory)
	tagsR := e.Group("/tags")
	{
		tagsR.GET("", tagController.List)
		tagsR.GET("/:id", tagController.Get)
	}
}

func installTagIam(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	tagController := tag.NewController(storeFactory)
	tagsR := e.Group("/tags")
	{
		tagsR.POST("", tagController.Create)
		tagsR.PATCH("", tagController.Update)
		tagsR.DELETE("", tagController.Del)
	}
}

func init() {
	installs = append(installs, installTag)
	installsIam = append(installsIam, installTagIam)
}
