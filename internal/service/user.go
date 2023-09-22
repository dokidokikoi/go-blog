package service

import (
	"context"
	"fmt"
	"go-blog/internal/db/model/user"
	"go-blog/internal/db/store"
	"time"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type UserSrv interface {
	Create(ctx context.Context, u *user.User, option *meta.CreateOption) error
	Get(ctx context.Context, u *user.User, option *meta.GetOption) (*user.User, error)
	List(ctx context.Context, keyword string, option *meta.ListOption) ([]*user.User, int64, error)
	Update(ctx context.Context, u *user.User, option *meta.UpdateOption) error

	SetRedisKvExpire(ctx context.Context, key, code string, expire time.Duration) error
	GetRedisKv(ctx context.Context, key string) (text string, err error)
	DelRedisKv(ctx context.Context, key string) error
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

func (srv userSrv) SetRedisKvExpire(ctx context.Context, key, code string, expire time.Duration) error {
	return srv.store.Users().SetRedisKvExpire(ctx, key, code, expire)
}

func (srv userSrv) GetRedisKv(ctx context.Context, key string) (text string, err error) {
	return srv.store.Users().GetRedisKv(ctx, key)
}
func (srv userSrv) DelRedisKv(ctx context.Context, key string) error {
	return srv.store.Users().DelRedisKv(ctx, key)
}

func (srv userSrv) Update(ctx context.Context, u *user.User, option *meta.UpdateOption) error {
	return srv.store.Users().Update(ctx, u, option)
}

func newUserSrv(store store.Factory) UserSrv {
	return &userSrv{store: store}
}
