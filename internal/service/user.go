package service

import (
	"context"
	"fmt"
	"go-blog/internal/db/model/user"
	"go-blog/internal/db/store"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type UserSrv interface {
	Create(ctx context.Context, u *user.User, option *meta.CreateOption) error
	Get(ctx context.Context, u *user.User, option *meta.GetOption) (*user.User, error)
	List(ctx context.Context, keyword string, option *meta.ListOption) ([]*user.User, int64, error)

	SetCaptchaCode(ctx context.Context, uuid, code string) error
	GetCaptchCode(ctx context.Context, key string) (text string, code error)
	DelCaptchCode(ctx context.Context, key string) error
}

type userSrv struct {
	store store.Factory
}

func (srv userSrv) Create(ctx context.Context, u *user.User, option *meta.CreateOption) error {
	u.ID = 0
	return srv.store.Users().Create(ctx, u, option)
}

func (srv userSrv) Get(ctx context.Context, u *user.User, option *meta.GetOption) (*user.User, error) {
	return srv.store.Users().Get(ctx, u, option)
}

func (srv userSrv) List(ctx context.Context, keyword string, option *meta.ListOption) ([]*user.User, int64, error) {
	if keyword == "" {
		count, _ := srv.store.Users().Count(ctx, &user.User{}, &option.GetOption)
		result, err := srv.store.Users().List(ctx, &user.User{}, option)
		return result, count, err
	}

	root := &meta.WhereNode{}
	root.Conditions = append(root.Conditions,
		&meta.Condition{Field: "nick_name", Operator: meta.LIKE, Value: fmt.Sprintf("%s%s%s", "%", keyword, "%")},
		&meta.Condition{Field: "account", Operator: meta.LIKE, Value: fmt.Sprintf("%s%s%s", "%", keyword, "%")})
	count, _ := srv.store.Users().CountComplex(ctx, &user.User{}, root, &option.GetOption)
	result, err := srv.store.Users().ListComplex(ctx, &user.User{}, root, option)
	return result, count, err
}

func (srv userSrv) SetCaptchaCode(ctx context.Context, uuid, code string) error {
	return srv.store.Users().SetCaptchCode(ctx, uuid, code)
}

func (srv userSrv) GetCaptchCode(ctx context.Context, key string) (text string, code error) {
	return srv.store.Users().GetCaptchCode(ctx, key)
}

func (srv userSrv) DelCaptchCode(ctx context.Context, key string) error {
	return srv.store.Users().DelCaptchCode(ctx, key)
}

func newUserSrv(store store.Factory) UserSrv {
	return &userSrv{store: store}
}
