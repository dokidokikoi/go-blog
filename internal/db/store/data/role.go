package data

import (
	"context"
	"go-blog/internal/db/model/user"
	"go-blog/internal/db/store"
	"go-blog/internal/db/store/data/postgres"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type roles struct {
	pg *postgres.Store
}

func (a roles) Create(ctx context.Context, t *user.Role, option *meta.CreateOption) error {
	return a.pg.Roles().Create(ctx, t, option)
}

func (a roles) CreateCollection(ctx context.Context, t []*user.Role, option *meta.CreateCollectionOption) []error {
	return a.pg.Roles().CreateCollection(ctx, t, option)
}

func (a roles) Update(ctx context.Context, t *user.Role, option *meta.UpdateOption) error {
	return a.pg.Roles().Update(ctx, t, option)
}

func (a roles) UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *user.Role, option *meta.UpdateOption) error {
	return a.pg.Roles().UpdateByWhere(ctx, node, example, option)
}

func (a roles) UpdateCollection(ctx context.Context, t []*user.Role, option *meta.UpdateCollectionOption) []error {
	return a.pg.Roles().UpdateCollection(ctx, t, option)
}

func (a roles) Save(ctx context.Context, t *user.Role, option *meta.UpdateOption) error {
	return a.pg.Roles().Save(ctx, t, option)
}

func (a roles) Get(ctx context.Context, t *user.Role, option *meta.GetOption) (*user.Role, error) {
	return a.pg.Roles().Get(ctx, t, option)
}

func (a roles) Count(ctx context.Context, t *user.Role, option *meta.GetOption) (int64, error) {
	return a.pg.Roles().Count(ctx, t, option)
}

func (a roles) CountComplex(ctx context.Context, example *user.Role, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	return a.pg.Roles().CountComplex(ctx, example, condition, option)
}

func (a roles) List(ctx context.Context, t *user.Role, option *meta.ListOption) ([]*user.Role, error) {
	return a.pg.Roles().List(ctx, t, option)
}

func (a roles) ListComplex(ctx context.Context, example *user.Role, condition *meta.WhereNode, option *meta.ListOption) ([]*user.Role, error) {
	return a.pg.Roles().ListComplex(ctx, example, condition, option)
}

func (a roles) Delete(ctx context.Context, t *user.Role, option *meta.DeleteOption) error {
	return a.pg.Roles().Delete(ctx, t, option)
}

func (a roles) DeleteCollection(ctx context.Context, t []*user.Role, option *meta.DeleteCollectionOption) []error {
	return a.pg.Roles().DeleteCollection(ctx, t, option)
}

func (a roles) DeleteByIds(ctx context.Context, ids []uint) error {
	return a.pg.Roles().DeleteByIds(ctx, ids)
}

func newRoles(d *dataCenter) store.Roles {
	return &roles{pg: d.pg}
}
