package service

import (
	"context"
	"go-blog/internal/db/model/category"
	"go-blog/internal/db/store"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type CategorySrv interface {
	Create(ctx context.Context, example *category.Category, option *meta.CreateOption) error
	Get(ctx context.Context, example *category.Category, option *meta.GetOption) (*category.Category, error)
	Update(ctx context.Context, example *category.Category, option *meta.UpdateOption) error
	Del(ctx context.Context, example *category.Category, option *meta.DeleteOption) error
	DeleteCollection(ctx context.Context, examples []*category.Category, option *meta.DeleteCollectionOption) []error
	List(ctx context.Context, example *category.Category, option *meta.ListOption) ([]*category.Category, int64, error)
	ListByWhereNode(ctx context.Context, example *category.Category, node *meta.WhereNode, option *meta.ListOption) ([]*category.Category, int64, error)
}

type categorySrv struct {
	store store.Factory
}

func (cs *categorySrv) Create(ctx context.Context, example *category.Category, option *meta.CreateOption) error {
	return cs.store.Category().Create(ctx, example, option)
}

func (cs *categorySrv) Get(ctx context.Context, example *category.Category, option *meta.GetOption) (*category.Category, error) {
	return cs.store.Category().Get(ctx, example, option)
}

func (cs *categorySrv) Update(ctx context.Context, example *category.Category, option *meta.UpdateOption) error {
	return cs.store.Category().Update(ctx, example, option)
}

func (cs *categorySrv) Del(ctx context.Context, example *category.Category, option *meta.DeleteOption) error {
	return cs.store.Category().Delete(ctx, example, option)
}

func (cs *categorySrv) DeleteCollection(ctx context.Context, examples []*category.Category, option *meta.DeleteCollectionOption) []error {
	return cs.store.Category().DeleteCollection(ctx, examples, option)
}

func (cs *categorySrv) List(ctx context.Context, example *category.Category, option *meta.ListOption) ([]*category.Category, int64, error) {
	total, err := cs.store.Category().Count(ctx, example, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}
	categorys, err := cs.store.Category().List(ctx, example, option)
	return categorys, total, err
}

func (cs *categorySrv) ListByWhereNode(ctx context.Context, example *category.Category, node *meta.WhereNode, option *meta.ListOption) ([]*category.Category, int64, error) {
	total, err := cs.store.Category().CountComplex(ctx, example, node, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}
	categorys, err := cs.store.Category().ListComplex(ctx, example, node, option)
	return categorys, total, err
}

func newCategorySrv(store store.Factory) CategorySrv {
	return &categorySrv{store: store}
}
