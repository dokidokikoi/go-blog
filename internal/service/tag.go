package service

import (
	"context"
	"go-blog/internal/db/model/tag"
	"go-blog/internal/db/store"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type TagSrv interface {
	Create(ctx context.Context, u *tag.Tag, option *meta.CreateOption) error
	Get(ctx context.Context, u *tag.Tag, option *meta.GetOption) (*tag.Tag, error)
	Update(ctx context.Context, example *tag.Tag, option *meta.UpdateOption) error
	Del(ctx context.Context, example *tag.Tag, option *meta.DeleteOption) error
	DeleteCollection(ctx context.Context, examples []*tag.Tag, option *meta.DeleteCollectionOption) []error
	List(ctx context.Context, example *tag.Tag, option *meta.ListOption) ([]*tag.Tag, int64, error)
	ListByWhereNode(ctx context.Context, example *tag.Tag, node *meta.WhereNode, option *meta.ListOption) ([]*tag.Tag, int64, error)
}

type tagSrv struct {
	store store.Factory
}

func (ts *tagSrv) Create(ctx context.Context, example *tag.Tag, option *meta.CreateOption) error {
	return ts.store.Tag().Create(ctx, example, option)
}

func (ts *tagSrv) Get(ctx context.Context, example *tag.Tag, option *meta.GetOption) (*tag.Tag, error) {
	return ts.store.Tag().Get(ctx, example, option)
}

func (ts *tagSrv) Update(ctx context.Context, example *tag.Tag, option *meta.UpdateOption) error {
	return ts.store.Tag().Update(ctx, example, option)
}

func (ts *tagSrv) Del(ctx context.Context, example *tag.Tag, option *meta.DeleteOption) error {
	return ts.store.Tag().Delete(ctx, example, option)
}

func (ts *tagSrv) DeleteCollection(ctx context.Context, examples []*tag.Tag, option *meta.DeleteCollectionOption) []error {
	return ts.store.Tag().DeleteCollection(ctx, examples, option)
}

func (ts *tagSrv) List(ctx context.Context, example *tag.Tag, option *meta.ListOption) ([]*tag.Tag, int64, error) {
	total, err := ts.store.Tag().Count(ctx, example, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}
	list, err := ts.store.Tag().List(ctx, example, option)
	return list, total, err
}

func (ts *tagSrv) ListByWhereNode(ctx context.Context, example *tag.Tag, node *meta.WhereNode, option *meta.ListOption) ([]*tag.Tag, int64, error) {
	total, err := ts.store.Tag().CountComplex(ctx, example, node, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}
	list, err := ts.store.Tag().ListComplex(ctx, example, node, option)
	return list, total, err
}

func newTagSrv(store store.Factory) TagSrv {
	return &tagSrv{
		store: store,
	}
}
