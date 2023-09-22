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

func (c *Controller) Update(ctx *gin.Context) {
	var input UpdateUser
	if ctx.ShouldBindJSON(&input) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	pwd := ""
	if input.OldPassword != "" && input.NewPassword != "" {
		u, err := c.srv.User().Get(ctx, &user.User{ID: input.ID}, nil)
		if err != nil {
			zaplog.L().Error("用户不存在", zap.Error(err))
			core.WriteResponse(ctx, myErrors.ApiRecordNotFound, nil)
			return
		}
		if !crypto.CheckPassword(input.OldPassword, u.Password) {
			zaplog.L().Error("密码错误")
			core.WriteResponse(ctx, myErrors.ApiErrPassword, nil)
			return
		}
		pwd, err = crypto.EncryptPassword(input.NewPassword)
		if err != nil {
			zaplog.L().Error("用户密码加密错误", zap.Error(err))
			core.WriteResponse(ctx, myErrors.ApiErrSystemErr, nil)
			return
		}
	}

	err := c.srv.User().Update(ctx,
		&user.User{
			ID:       input.ID,
			NickName: input.NickName,
			Password: pwd,
			Email:    input.Email,
			Avatar:   input.Avatar,
			Intro:    input.Intro,
		}, nil)
	if err != nil {
		zaplog.L().Error("更新用户信息失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiNoUpdateRows, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
