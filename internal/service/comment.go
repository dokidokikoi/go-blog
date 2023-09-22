package service

import (
	"context"
	"go-blog/internal/db/model/comment"
	"go-blog/internal/db/store"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type CommentSrv interface {
	Create(ctx context.Context, example *comment.Comment, option *meta.CreateOption) error
	Get(ctx context.Context, example *comment.Comment, option *meta.GetOption) (*comment.Comment, error)
	Update(ctx context.Context, example *comment.Comment, option *meta.UpdateOption) error
	Del(ctx context.Context, example *comment.Comment, option *meta.DeleteOption) error
	DeleteCollection(ctx context.Context, examples []*comment.Comment, option *meta.DeleteCollectionOption) []error
	List(ctx context.Context, example *comment.Comment, option *meta.ListOption) ([]*comment.Comment, int64, error)
	ListByWhereNode(ctx context.Context, example *comment.Comment, node *meta.WhereNode, option *meta.ListOption) ([]*comment.Comment, int64, error)

	SetCommentChildren(ctx context.Context, examples []*comment.Comment) []error
}

type commentSrv struct {
	store store.Factory
}

func (cs *commentSrv) Create(ctx context.Context, example *comment.Comment, option *meta.CreateOption) error {
	return cs.store.Comments().Create(ctx, example, option)
}

func (cs *commentSrv) Get(ctx context.Context, example *comment.Comment, option *meta.GetOption) (*comment.Comment, error) {
	return cs.store.Comments().Get(ctx, example, option)
}

func (cs *commentSrv) Update(ctx context.Context, example *comment.Comment, option *meta.UpdateOption) error {
	return cs.store.Comments().Update(ctx, example, option)
}

func (cs *commentSrv) Del(ctx context.Context, example *comment.Comment, option *meta.DeleteOption) error {
	return cs.store.Comments().Delete(ctx, example, option)
}

func (cs *commentSrv) DeleteCollection(ctx context.Context, examples []*comment.Comment, option *meta.DeleteCollectionOption) []error {
	return cs.store.Comments().DeleteCollection(ctx, examples, option)
}

func (cs *commentSrv) List(ctx context.Context, example *comment.Comment, option *meta.ListOption) ([]*comment.Comment, int64, error) {
	total, err := cs.store.Comments().Count(ctx, example, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}
	list, err := cs.store.Comments().List(ctx, example, option)
	return list, total, err
}

func (cs *commentSrv) ListByWhereNode(ctx context.Context, example *comment.Comment, node *meta.WhereNode, option *meta.ListOption) ([]*comment.Comment, int64, error) {
	total, err := cs.store.Comments().CountComplex(ctx, example, node, &option.GetOption)
	if err != nil {
		return nil, 0, err
	}
	list, err := cs.store.Comments().ListComplex(ctx, example, node, option)
	return list, total, err
}

func (cs *commentSrv) SetCommentChildren(ctx context.Context, examples []*comment.Comment) []error {
	var errs []error
	for _, e := range examples {
		if e == nil {
			continue
		}
		var es []error
		e.Children, es = cs.GetCommentChildren(ctx, e)
		if es != nil {
			errs = append(errs, es...)
		}
	}
	return errs
}

func (cs commentSrv) GetCommentChildren(ctx context.Context, example *comment.Comment) ([]*comment.Comment, []error) {
	if example == nil {
		return nil, nil
	}
	option := &meta.ListOption{PageSize: 100, Page: 1}
	children, _, err := cs.List(ctx, &comment.Comment{PID: example.ID}, option)
	if err != nil {
		return nil, []error{err}
	}
	var errs []error
	for _, e := range children {
		list, es := cs.GetCommentChildren(ctx, e)
		if err != nil {
			errs = append(errs, es...)
			continue
		}
		children = append(children, list...)
	}
	return children, errs
}

func newCommentSrv(store store.Factory) CommentSrv {
	return &commentSrv{store: store}
}
