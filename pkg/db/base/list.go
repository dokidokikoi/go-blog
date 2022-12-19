package base

import (
	"context"

	meta "go-blog/pkg/meta/option"
)

func (p *PgModel[T]) Count(ctx context.Context, t *T, option *meta.GetOption) (int64, error) {
	return 0, nil
}

func (p *PgModel[T]) CountComplex(ctx context.Context, example *T, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	return 0, nil
}

func (p *PgModel[T]) List(ctx context.Context, t *T, option *meta.ListOption) ([]*T, error) {
	return nil, nil
}

func (p *PgModel[T]) ListComplex(ctx context.Context, example *T, condition *meta.WhereNode, option *meta.ListOption) ([]*T, error) {
	return nil, nil
}
