package user

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/user"
	myErrors "go-blog/internal/errors"

	"github.com/dokidokikoi/go-common/crypto"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Register(ctx *gin.Context) {
	var createUser CreateUser
	if ctx.ShouldBindJSON(&createUser) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	pwd, err := crypto.EncryptPassword(createUser.Password)
	if err != nil {
		zaplog.L().Error("用户密码加密错误", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, nil)
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
		Password: pwd,
		RoleID:   targetRole.ID,
	}
	if err := c.srv.User().Create(ctx, targetUser, nil); err != nil {
		zaplog.L().Error("创建用户失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrDatabase, nil)
		return
	}

	core.WriteResponse(ctx, myErrors.Success("创建成功"), nil)
}
