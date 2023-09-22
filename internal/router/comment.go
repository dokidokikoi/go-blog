package router

import (
	"go-blog/internal/controller/comment"
	"go-blog/internal/db/store"

	"github.com/gin-gonic/gin"
)

func installComment(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	commentController := comment.NewController(storeFactory)
	commentsR := e.Group("/comments")
	{
		commentsR.GET("", commentController.List)
		commentsR.GET("/:id", commentController.Get)
		commentsR.POST("", commentController.Create)
	}
}

func installCommentIam(e *gin.Engine) {
	storeFactory, _ := store.GetStoreFactory()
	commentController := comment.NewController(storeFactory)
	commentsR := e.Group("/comments")
	{
		commentsR.PATCH("", commentController.Update)
		commentsR.DELETE("", commentController.Del)
	}
}

func init() {
	installs = append(installs, installComment)
	installsIam = append(installsIam, installCommentIam)
}
