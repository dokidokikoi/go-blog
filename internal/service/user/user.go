package user

import (
	"context"
	"fmt"
	"go-blog/internal/db/model/user"
	"go-blog/internal/db/store"
	meta "go-blog/pkg/meta/option"
)

type UserSrv interface {
	Create(ctx context.Context, u *user.User, option *meta.CreateOption) error
	List(ctx context.Context, keyword string, option *meta.ListOption) ([]*user.User, int64, error)
}

type userSrv struct {
	store store.Factory
}

func (srv userSrv) Create(ctx context.Context, u *user.User, option *meta.CreateOption) error {
	u.ID = 0
	return srv.store.Users().Create(ctx, u, option)
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

func NewUserSrv(store store.Factory) UserSrv {
	return &userSrv{store: store}
}
