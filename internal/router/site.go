package router

import (
	"go-blog/internal/controller/site"
	"go-blog/internal/db/store"

	"github.com/gin-gonic/gin"
)

// func installSite(e *gin.Engine) {
// 	storeFactory, _ := store.GetStoreFactory()
// 	siteController := site.NewController(storeFactory)
// 	sitesR := e.Group("/sites")
// 	{
// 		sitesR.GET("", siteController.List)
// 		sitesR.GET("/:id", siteController.Get)
// 		sitesR.GET("/tag/:id", siteController.ListTagSite)
// 	}
// }

func installSiteIam(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	siteController := site.NewController(storeFactory)
	sitesR := e.Group("/sites")
	{
		sitesR.GET("", siteController.List)
		sitesR.GET("/:id", siteController.Get)
		sitesR.GET("/tag/:id", siteController.ListTagSite)
		sitesR.POST("", siteController.Create)
		sitesR.PATCH("", siteController.Update)
		sitesR.DELETE("", siteController.Del)
	}
}

func init() {
	// installs = append(installs, installSite)
	installsIam = append(installsIam, installSiteIam)
}
