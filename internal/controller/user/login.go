package user

import (
	"fmt"
	"go-blog/internal/core"
	"go-blog/internal/db/model/user"
	myErrors "go-blog/internal/errors"
	"go-blog/pkg/captcha"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dokidokikoi/go-common/crypto"
	myJwt "github.com/dokidokikoi/go-common/jwt"
	zaplog "github.com/dokidokikoi/go-common/log/zap"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Login(ctx *gin.Context) {
	var loginParam LoginParam
	if ctx.ShouldBindJSON(&loginParam) != nil {
		core.WriteResponse(ctx, myErrors.ApiErrValidation, nil)
		return
	}

	if !verifyCaptcha(loginParam.UUID, loginParam.Code) {
		zaplog.L().Error("验证码验证失败")
		core.WriteResponse(ctx, myErrors.ApiErrCaptcha, nil)
		return
	}

	u, err := c.srv.User().Get(ctx, &user.User{Email: loginParam.Email}, nil)
	if err != nil {
		zaplog.L().Error("用户不存在", zap.Error(err))
		core.WriteResponse(ctx, myErrors.ApiRecordNotFound, nil)
		return
	}

	if !crypto.CheckPassword(loginParam.Password, u.Password) {
		zaplog.L().Error("密码错误")
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

func (c *Controller) GetCaptha(ctx *gin.Context) {
	text, captcha := captcha.GetStandCaptcha()
	uuid := uuid.New().String()
	fmt.Println(text, uuid)
	// TODO: 将text存入redis
	core.WriteResponse(ctx, nil, gin.H{"captcha": captcha})
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

func verifyCaptcha(uuid, text string) bool {
	// TODO: 根据uuid将验证码从redis拿出与用户对比
	return true
}