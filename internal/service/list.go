package service

import (
	"context"
	"go-blog/internal/db/model/list"
	"go-blog/internal/db/store"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type ItemSrv interface {
	Create(ctx context.Context, example *list.Item, option *meta.CreateOption) error
	Get(ctx context.Context, example *list.Item, option *meta.GetOption) (*list.Item, error)
	Update(ctx context.Context, example *list.Item, option *meta.UpdateOption) error
	Del(ctx context.Context, example *list.Item, option *meta.DeleteOption) error
	DeleteCollection(ctx context.Context, examples []*list.Item, option *meta.DeleteCollectionOption) []error
	List(ctx context.Context, example *list.Item, option *meta.ListOption) ([]*list.Item, int64, error)
	ListByWhereNode(ctx context.Context, example *list.Item, node *meta.WhereNode, option *meta.ListOption) ([]*list.Item, int64, error)
}

type itemSrv struct {
	store store.Factory
}

func (is *itemSrv) Create(ctx context.Context, example *list.Item, option *meta.CreateOption) error {
	return is.store.Items().Create(ctx, example, option)
}

func (is *itemSrv) Get(ctx context.Context, example *list.Item, option *meta.GetOption) (*list.Item, error) {
	return is.store.Items().Get(ctx, example, option)
}

func (is *itemSrv) Update(ctx context.Context, example *list.Item, option *meta.UpdateOption) error {
	return is.store.Items().Update(ctx, example, option)
}

func (is *itemSrv) Del(ctx context.Context, example *list.Item, option *meta.DeleteOption) error {
	return is.store.Items().Delete(ctx, example, option)
}

func (is *itemSrv) DeleteCollection(ctx context.Context, examples []*list.Item, option *meta.DeleteCollectionOption) []error {
	return is.store.Items().DeleteCollection(ctx, examples, option)
}

func (is *itemSrv) List(ctx context.Context, example *list.Item, option *meta.ListOption) ([]*list.Item, int64, error) {
	total, err := is.store.Items().Count(ctx, example, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}

	list, err := is.store.Items().List(ctx, example, option)
	return list, total, err
}

func (is *itemSrv) ListByWhereNode(ctx context.Context, example *list.Item, node *meta.WhereNode, option *meta.ListOption) ([]*list.Item, int64, error) {
	total, err := is.store.Items().CountComplex(ctx, example, node, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}

	list, err := is.store.Items().ListComplex(ctx, example, node, option)
	return list, total, err
}

func newItemSrv(store store.Factory) ItemSrv {
	return &itemSrv{store: store}
}
