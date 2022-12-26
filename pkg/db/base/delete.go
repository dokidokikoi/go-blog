package base

import (
	"context"

	meta "go-blog/pkg/meta/option"
)

func (p *PgModel[T]) Delete(ctx context.Context, t *T, option *meta.DeleteOption) error {
	var example T
	err := p.DB.Where(t).Delete(example).Error
	return err
}

func (p *PgModel[T]) DeleteCollection(ctx context.Context, t []*T, option *meta.DeleteCollectionOption) []error {
	var errors []error
	var example T
	for _, entity := range t {
		if e := p.DB.Where(entity).Delete(example).Error; e != nil {
			errors = append(errors, e)
		}
	}
	return errors
}

func (p PgModel[T]) DeleteByIds(ctx context.Context, ids []uint) error {
	var entities T
	err := p.DB.Delete(&entities, ids).Error
	return err
}
