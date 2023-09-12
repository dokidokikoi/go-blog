package user

import (
	"go-blog/internal/core"
	"go-blog/internal/db/model/user"
	myErrors "go-blog/internal/errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dokidokikoi/go-common/crypto"
	myJwt "github.com/dokidokikoi/go-common/jwt"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Login(ctx *gin.Context) {
	var loginParam LoginParam
	if ctx.ShouldBindJSON(&loginParam) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	u, err := c.srv.User().Get(ctx, &user.User{Email: loginParam.Email}, nil)
	if err != nil {
		core.WriteResponse(ctx, myErrors.ApiRecordNotFound, nil)
		return
	}

	if !crypto.CheckPassword(loginParam.Password, u.Password) {
		core.WriteResponse(ctx, myErrors.ApiErrPassword, nil)
		return
	}
	token, err := GenerateToken(u)
	if err != nil {
		zaplog.L().Error("获取jwt token失败", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiErrSystemErr, nil)
		return
	}

	core.WriteResponse(ctx, nil, gin.H{"token": token})
}

func GenerateToken(u *user.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(364 * 60 * 60 * 24 * time.Second)
	issuer := "harukaze"
	claims := myJwt.CustomClaims{
		ID:    u.ID,
		Emial: u.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  nowTime.Unix(),
			NotBefore: nowTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := myJwt.GenerateToken(claims, "test")
	return token, err
}
