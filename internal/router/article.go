package router

import (
	"go-blog/internal/controller/article"
	"go-blog/internal/db/store"

	"github.com/gin-gonic/gin"
)

func InstallArticle(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	articlesController := article.NewController(storeFactory)
	articlesR := e.Group("/articles")
	{
		articlesR.GET("", articlesController.List)
	}
}

func InstallArticleIam(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	articlesController := article.NewController(storeFactory)
	articlesR := e.Group("/articles")
	{
		articlesR.POST("", articlesController.Create)
	}
}
