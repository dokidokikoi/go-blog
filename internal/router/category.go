package router

import (
	"go-blog/internal/controller/category"
	"go-blog/internal/db/store"

	"github.com/gin-gonic/gin"
)

func installCategory(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	categoriesController := category.NewController(storeFactory)
	categoriesR := e.Group("/categories")
	{
		categoriesR.GET("", categoriesController.List)
		categoriesR.GET("/:id", categoriesController.Get)
	}
}

func installCategoryIam(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	categoriesController := category.NewController(storeFactory)
	categoriesR := e.Group("/categories")
	{
		categoriesR.POST("", categoriesController.Create)
		categoriesR.PATCH("", categoriesController.Update)
		categoriesR.DELETE("", categoriesController.Del)
	}
}

func init() {
	installs = append(installs, installCategory)
	installsIam = append(installsIam, installCategoryIam)
}
