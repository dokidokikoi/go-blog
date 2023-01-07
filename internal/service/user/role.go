package user

import (
	"context"
	"fmt"
	"go-blog/internal/db/model/user"
	"go-blog/internal/db/store"
	meta "go-blog/pkg/meta/option"

	"gorm.io/gorm"
)

type RoleSrv interface {
	Create(ctx context.Context, u *user.Role, option *meta.CreateOption) error
	IsExist(ctx context.Context, a *user.Role) (bool, error)
	List(ctx context.Context, keyword string, option *meta.ListOption) ([]*user.Role, int64, error)
}

type roleSrv struct {
	store store.Factory
}

func (srv roleSrv) Create(ctx context.Context, u *user.Role, option *meta.CreateOption) error {
	u.ID = 0
	return srv.store.Roles().Create(ctx, u, option)
}

func (srv roleSrv) IsExist(ctx context.Context, r *user.Role) (bool, error) {
	if r.ID == 0 {
		cate, err := srv.store.Roles().Get(ctx, r, nil)
		if err != nil {
			return false, err
		}
		if cate == nil {
			return false, nil
		}
		r.ID = cate.ID
		return true, nil
	}
	cate, err := srv.store.Roles().Get(ctx, &user.Role{Model: gorm.Model{ID: r.ID}}, &meta.GetOption{Include: []string{"id"}})
	if err != nil {
		return false, err
	}
	if cate == nil {
		return false, err
	}
	return true, nil
}

func (srv roleSrv) List(ctx context.Context, keyword string, option *meta.ListOption) ([]*user.Role, int64, error) {
	if keyword == "" {
		count, _ := srv.store.Roles().Count(ctx, &user.Role{}, &option.GetOption)
		result, err := srv.store.Roles().List(ctx, &user.Role{}, option)
		return result, count, err
	}

	root := &meta.WhereNode{}
	root.Conditions = append(root.Conditions,
		&meta.Condition{Field: "role_name", Operator: meta.LIKE, Value: fmt.Sprintf("%s%s%s", "%", keyword, "%")})
	count, _ := srv.store.Roles().CountComplex(ctx, &user.Role{}, root, &option.GetOption)
	result, err := srv.store.Roles().ListComplex(ctx, &user.Role{}, root, option)
	return result, count, err
}

func NewRoleSrv(store store.Factory) RoleSrv {
	return &roleSrv{store: store}
}
