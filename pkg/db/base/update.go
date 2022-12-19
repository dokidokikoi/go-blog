package base

import (
	"context"

	meta "go-blog/pkg/meta/option"
)

func (p *PgModel[T]) Update(ctx context.Context, t *T, option *meta.UpdateOption) error {
	return nil
}

func (p *PgModel[T]) UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *T, option *meta.UpdateOption) error {
	return nil
}

func (p *PgModel[T]) UpdateCollection(ctx context.Context, t []*T, option *meta.UpdateCollectionOption) []error {
	return nil
}

func (p *PgModel[T]) Save(ctx context.Context, t *T, option *meta.UpdateOption) error {
	return nil
}
