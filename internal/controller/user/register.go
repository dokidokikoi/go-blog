package user

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/user"
	myErrors "go-blog/internal/errors"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Register(ctx *gin.Context) {
	var createUser CreateUser
	if ctx.ShouldBindJSON(&createUser) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	targetRole := &user.Role{
		RoleName: "none",
	}
	if ok, err := c.srv.Role().IsExist(ctx, targetRole); !ok || err != nil {
		c.srv.Role().Create(ctx, targetRole, nil)
	}

	targetUser := &user.User{
		Account:  createUser.Account,
		Avatar:   createUser.Avatar,
		Email:    createUser.Email,
		NickName: createUser.NickName,
		Password: createUser.Password,
		RoleID:   targetRole.ID,
	}
	if c.srv.User().Create(ctx, targetUser, nil) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}

	core.WriteResponse(ctx, myErrors.Success("创建成功"), nil)
}
