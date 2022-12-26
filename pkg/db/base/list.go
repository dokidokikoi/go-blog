package base

import (
	"context"
	"errors"

	myErrors "go-blog/internal/errors"
	meta "go-blog/pkg/meta/option"

	"gorm.io/gorm"
)

func (p *PgModel[T]) Count(ctx context.Context, t *T, option *meta.GetOption) (int64, error) {
	var result int64
	if option != nil && len(option.Include) > 0 {
		var fields []any
		for _, i := range option.Include {
			fields = append(fields, i)
		}
		err := p.DB.Model(t).Where(t, fields...).Count(&result).Error
		return result, err
	}
	err := p.DB.Model(t).Where(t).Count(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = myErrors.RecordNotFound
	}
	return result, err
}

func (p *PgModel[T]) CountComplex(ctx context.Context, example *T, condition *meta.WhereNode, option *meta.GetOption) (int64, error) {
	var result int64
	var t T
	var db = p.DB.Model(&t)
	if option != nil {
		if option.Include == nil {
			db = db.Where(example)
		} else {
			var fields []any
			for _, i := range option.Include {
				fields = append(fields, i)
			}
			db = db.Where(example, fields...)
		}
	}
	err := CompositeQuery(db, condition).Count(&result).Error
	return result, err
}

func (p *PgModel[T]) List(ctx context.Context, t *T, option *meta.ListOption) ([]*T, error) {
	var tList []*T
	err := CommonDealList(p.DB, t, option).Find(&tList).Error
	return tList, err
}

func (p *PgModel[T]) ListComplex(ctx context.Context, example *T, condition *meta.WhereNode, option *meta.ListOption) ([]*T, error) {
	var tList []*T
	err := CompositeQuery(CommonDealList(p.DB, example, option), condition).Find(&tList).Error
	return tList, err
}
