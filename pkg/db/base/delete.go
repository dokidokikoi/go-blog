package base

import (
	"context"

	meta "go-blog/pkg/meta/option"
)

func (p *PgModel[T]) Delete(ctx context.Context, t *T, option *meta.DeleteOption) error {
	return nil
}

func (p *PgModel[T]) DeleteCollection(ctx context.Context, t []*T, option *meta.DeleteCollectionOption) []error {
	return nil
}

func (p PgModel[T]) DeleteByIds(ctx context.Context, ids []uint) error {
	return nil
}
