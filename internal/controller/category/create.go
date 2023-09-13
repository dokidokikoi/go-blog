package category

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/category"
	myErrors "go-blog/internal/errors"

	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Create(ctx *gin.Context) {
	var input CreateCategory
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	_, err := c.srv.Category().Get(ctx, &category.Category{CategoryName: input.CategoryName, Type: input.Type}, nil)
	if err == nil {
		core.WriteResponse(ctx, nil, nil)
		return
	}

	err = c.srv.Category().Create(
		ctx,
		&category.Category{
			CategoryName: input.CategoryName,
			Summary:      input.Summary,
			Type:         input.Type,
		},
		nil,
	)
	if err != nil {
		zaplog.L().Error("创建分类失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
