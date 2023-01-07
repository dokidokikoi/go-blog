package user

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/user"
	myErrors "go-blog/internal/errors"

	"github.com/gin-gonic/gin"
)

type CreateUser struct {
	Account  string `json:"account" binding:"required"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email" binding:"required"`
	NickName string `json:"nick_name" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func (c *Controller) Create(ctx *gin.Context) {
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
