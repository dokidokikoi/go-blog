package router

import (
	"go-blog/internal/controller/user"
	"go-blog/internal/db/store"

	"github.com/gin-gonic/gin"
)

func installUser(r *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	usersR := r.Group("/users")
	userController := user.NewController(storeFactory)
	{
		usersR.POST("/register", userController.Register)
		usersR.POST("/login", userController.Login)
		usersR.POST("/captcha", userController.GetCaptha)
		usersR.GET("/host", userController.Host)
	}
}

func installUserIam(r *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	usersR := r.Group("/users")
	{
		userController := user.NewController(storeFactory)
		usersR.GET("/:id", userController.Get)
		usersR.PATCH("", userController.Update)
		usersR.POST("/logout", userController.Logout)
	}
}

func init() {
	installs = append(installs, installUser)
	installsIam = append(installsIam, installUserIam)
}
