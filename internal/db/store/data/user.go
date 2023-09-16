package data

import (
	"context"
	"go-blog/internal/db/model/user"
	"go-blog/internal/db/store"
	"go-blog/internal/db/store/data/postgres"
	"go-blog/internal/db/store/data/redis"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type users struct {
	pg       *postgres.Store
	redisCli redis.Store
}

func (a users) Create(ctx context.Context, t *user.User, option *meta.CreateOption) error {
	return a.pg.Users().Create(ctx, t, option)
}

func (a users) CreateCollection(ctx context.Context, t []*user.User, option *meta.CreateCollectionOption) []error {
	return a.pg.Users().CreateCollection(ctx, t, option)
}

func (a users) Update(ctx context.Context, t *user.User, option *meta.UpdateOption) error {
	return a.pg.Users().Update(ctx, t, option)
}

func (a users) UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *user.User, option *meta.UpdateOption) error {
	return a.pg.Users().UpdateByWhere(ctx, node, example, option)
}

func (a users) UpdateCollection(ctx context.Context, t []*user.User, option *meta.UpdateCollectionOption) []error {
	return a.pg.Users().UpdateCollection(ctx, t, option)
}

func (a users) Save(ctx context.Context, t *user.User, option *meta.UpdateOption) error {
	return a.pg.Users().Save(ctx, t, option)
}

func (a users) Get(ctx context.Context, t *user.User, option *meta.GetOption) (*user.User, error) {
	return a.pg.Users().Get(ctx, t, option)
}

func (a users) Count(ctx context.Context, t *user.User, option *meta.GetOption) (int64, error) {
	return a.pg.Users().Count(ctx, t, option)
}

func (a users) CountComplex(ctx context.Context, example *user.User, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	return a.pg.Users().CountComplex(ctx, example, condition, option)
}

func (a users) List(ctx context.Context, t *user.User, option *meta.ListOption) ([]*user.User, error) {
	return a.pg.Users().List(ctx, t, option)
}

func (a users) ListComplex(ctx context.Context, example *user.User, condition *meta.WhereNode, option *meta.ListOption) ([]*user.User, error) {
	return a.pg.Users().ListComplex(ctx, example, condition, option)
}

func (a users) Delete(ctx context.Context, t *user.User, option *meta.DeleteOption) error {
	return a.pg.Users().Delete(ctx, t, option)
}

func (a users) DeleteCollection(ctx context.Context, t []*user.User, option *meta.DeleteCollectionOption) []error {
	return a.pg.Users().DeleteCollection(ctx, t, option)
}

func (a users) DeleteByIds(ctx context.Context, ids []uint) error {
	return a.pg.Users().DeleteByIds(ctx, ids)
}

func (a users) SetCaptchCode(ctx context.Context, key, code string) error {
	return a.redisCli.User().SetCaptchCode(ctx, key, code)
}

func (a users) GetCaptchCode(ctx context.Context, key string) (text string, err error) {
	return a.redisCli.User().GetCaptchCode(ctx, key)
}

func (a users) DelCaptchCode(ctx context.Context, key string) error {
	return a.DelCaptchCode(ctx, key)
}

func newUsers(d *dataCenter) store.Users {
	return &users{pg: d.pg, redisCli: *d.redisCli}
}
