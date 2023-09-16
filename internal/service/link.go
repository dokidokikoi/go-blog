package service

import (
	"context"
	"go-blog/internal/db/model/link"
	"go-blog/internal/db/store"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type LinkSrv interface {
	Create(ctx context.Context, example *link.Link, option *meta.CreateOption) error
	Get(ctx context.Context, example *link.Link, option *meta.GetOption) (*link.Link, error)
	Update(ctx context.Context, example *link.Link, option *meta.UpdateOption) error
	Del(ctx context.Context, example *link.Link, option *meta.DeleteOption) error
	DeleteCollection(ctx context.Context, examples []*link.Link, option *meta.DeleteCollectionOption) []error
	List(ctx context.Context, example *link.Link, option *meta.ListOption) ([]*link.Link, int64, error)
	ListByWhereNode(ctx context.Context, example *link.Link, node *meta.WhereNode, option *meta.ListOption) ([]*link.Link, int64, error)
}

type linkSrv struct {
	store store.Factory
}

func (ls linkSrv) Create(ctx context.Context, example *link.Link, option *meta.CreateOption) error {
	return ls.store.Link().Create(ctx, example, option)
}

func (ls linkSrv) Get(ctx context.Context, example *link.Link, option *meta.GetOption) (*link.Link, error) {
	return ls.store.Link().Get(ctx, example, option)
}

func (ls linkSrv) Update(ctx context.Context, example *link.Link, option *meta.UpdateOption) error {
	return ls.store.Link().Update(ctx, example, option)
}

func (ls linkSrv) Del(ctx context.Context, example *link.Link, option *meta.DeleteOption) error {
	return ls.store.Link().Delete(ctx, example, option)
}

func (ls linkSrv) DeleteCollection(ctx context.Context, examples []*link.Link, option *meta.DeleteCollectionOption) []error {
	return ls.store.Link().DeleteCollection(ctx, examples, option)
}

func (ls linkSrv) List(ctx context.Context, example *link.Link, option *meta.ListOption) ([]*link.Link, int64, error) {
	total, err := ls.store.Link().Count(ctx, example, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}
	list, err := ls.store.Link().List(ctx, example, option)
	return list, total, err
}

func (ls linkSrv) ListByWhereNode(ctx context.Context, example *link.Link, node *meta.WhereNode, option *meta.ListOption) ([]*link.Link, int64, error) {
	total, err := ls.store.Link().CountComplex(ctx, example, node, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}
	list, err := ls.store.Link().ListComplex(ctx, example, node, option)
	return list, total, err
}

func newLinkSrv(store store.Factory) LinkSrv {
	return &linkSrv{store: store}
}
