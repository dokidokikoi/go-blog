package router

import (
	"github.com/gin-gonic/gin"
)

func InstallAll(r *gin.Engine) {
	installArticle(r)
	installUser(r)
	installCategory(r)

	installArticleIam(r)
	installUserIam(r)
	installcategoryIam(r)
}
