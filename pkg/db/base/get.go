package base

import (
	"context"

	meta "go-blog/pkg/meta/option"
)

func (p *PgModel[T]) Get(ctx context.Context, t *T, option *meta.GetOption) (*T, error) {
	return nil, nil
}
