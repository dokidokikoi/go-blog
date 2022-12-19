package base

import (
	"context"

	meta "go-blog/pkg/meta/option"
)

func (p *PgModel[T]) Create(ctx context.Context, t *T, option *meta.CreateOption) error {
	return nil
}

func (p *PgModel[T]) CreateCollection(ctx context.Context, t []*T, option *meta.CreateCollectionOption) []error {
	return nil
}
