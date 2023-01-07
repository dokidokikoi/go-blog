package router

import (
	"go-blog/internal/controller/role"
	"go-blog/internal/controller/user"
	"go-blog/internal/db/store"

	"github.com/gin-gonic/gin"
)

func installUser(r *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	usersR := r.Group("/users")
	{
		userController := user.NewController(storeFactory)
		usersR.GET("", userController.List)
	}

	rolesR := r.Group("/roles")
	{
		roleController := role.NewController(storeFactory)
		rolesR.GET("", roleController.List)
	}
}

func installUserIam(r *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	usersR := r.Group("/users")
	{
		userController := user.NewController(storeFactory)
		usersR.POST("", userController.Create)
	}

	rolesR := r.Group("/roles")
	{
		roleController := role.NewController(storeFactory)
		rolesR.POST("", roleController.Create)
	}
}
