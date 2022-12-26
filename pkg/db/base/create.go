package base

import (
	"context"

	meta "go-blog/pkg/meta/option"
)

func (p *PgModel[T]) Create(ctx context.Context, t *T, option *meta.CreateOption) error {
	err := p.DB.Create(t).Error
	return err
}

func (p *PgModel[T]) CreateCollection(ctx context.Context, t []*T, option *meta.CreateCollectionOption) []error {
	var errors []error
	for _, up := range t {
		if e := p.Create(ctx, up, &meta.CreateOption{}); e != nil {
			errors = append(errors, e)
		}
	}
	return errors
}
