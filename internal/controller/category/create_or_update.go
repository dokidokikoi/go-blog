package category

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/category"
	myErrors "go-blog/internal/errors"

	"github.com/gin-gonic/gin"
)

func (c Controller) CreateOrUpdate(ctx *gin.Context) {
	var input CategoryCreateUpdate
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	targetCategory := category.Category{
		ID:           input.ID,
		CategoryName: input.CategoryName,
		Summary:      input.Summary,
		Type:         input.Type,
	}

	if input.ID == 0 {
		if c.srv.Category().Create(ctx, &targetCategory) != nil {
			core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
			return
		}
	} else {
		if c.srv.Category().Update(ctx, &targetCategory) != nil {
			core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
			return
		}
	}

	core.WriteResponse(ctx, myErrors.Success("操作成功"), nil)
}
