package role

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/user"
	myErrors "go-blog/internal/errors"

	"github.com/gin-gonic/gin"
)

type CreateRole struct {
	RoleName string `json:"role_name" binding:"required"`
}

func (c *Controller) Create(ctx *gin.Context) {
	var createRole CreateRole
	if ctx.ShouldBindJSON(&createRole) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	targetRole := &user.Role{
		RoleName: createRole.RoleName,
	}
	if c.srv.Role().Create(ctx, targetRole, nil) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}

	core.WriteResponse(ctx, myErrors.Success("创建成功"), nil)
}
