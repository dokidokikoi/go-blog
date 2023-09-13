package category

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/category"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Update(ctx *gin.Context) {
	var input UpdateCategory
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	err := c.srv.Category().Update(ctx, &category.Category{ID: input.ID, CategoryName: input.CategoryName, Summary: input.Summary}, nil)
	if err != nil {
		zaplog.L().Error("更新分类失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
