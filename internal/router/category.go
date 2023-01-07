package router

import (
	"go-blog/internal/controller/category"
	"go-blog/internal/db/store"

	"github.com/gin-gonic/gin"
)

func installCategory(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	categoriesController := category.NewController(storeFactory)
	articlesR := e.Group("/categories")
	{
		articlesR.GET("", categoriesController.List)
		articlesR.GET("/:id", categoriesController.Get)
	}
}

func installcategoryIam(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	categoriesController := category.NewController(storeFactory)
	articlesR := e.Group("/categories")
	{
		articlesR.POST("", categoriesController.CreateOrUpdate)
		articlesR.DELETE("/:id", categoriesController.Delete)
		articlesR.DELETE("", categoriesController.DeleteByIds)
	}
}
