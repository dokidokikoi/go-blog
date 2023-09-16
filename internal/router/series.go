package router

import (
	"go-blog/internal/controller/series"
	"go-blog/internal/db/store"

	"github.com/gin-gonic/gin"
)

func installSeries(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	seriesController := series.NewController(storeFactory)
	seriesR := e.Group("/series")
	{
		seriesR.GET("", seriesController.List)
		seriesR.GET("/:id", seriesController.Get)
	}
}

func installSeriesIam(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	seriesController := series.NewController(storeFactory)
	seriesR := e.Group("/series")
	{
		seriesR.POST("", seriesController.Create)
		seriesR.PATCH("", seriesController.Update)
		seriesR.DELETE("", seriesController.Del)
	}
}

func init() {
	installs = append(installs, installSeries)
	installsIam = append(installsIam, installSeriesIam)
}
