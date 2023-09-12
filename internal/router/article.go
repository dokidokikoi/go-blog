package router

import (
	"github.com/gin-gonic/gin"
)

func installArticle(e *gin.Engine) {
	// storeFactory, _ := store.GetStoreFactory()
	// articlesController := article.NewController(storeFactory)
	// articlesR := e.Group("/articles")
	// {
	// 	articlesR.GET("", articlesController.List)
	// 	articlesR.GET("/:id", articlesController.Get)
	// }
}

func installArticleIam(e *gin.Engine) {
	// storeFactory, _ := store.GetStoreFactory()
	// articlesController := article.NewController(storeFactory)
	// articlesR := e.Group("/articles")
	// {
	// 	articlesR.POST("", articlesController.Create)
	// 	articlesR.PATCH("", articlesController.Update)
	// 	articlesR.DELETE("/:id", articlesController.Delete)
	// }
}
