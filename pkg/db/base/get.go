package base

import (
	"context"
	"errors"

	myErrors "go-blog/internal/errors"
	meta "go-blog/pkg/meta/option"

	"gorm.io/gorm"
)

func (p *PgModel[T]) Get(ctx context.Context, t *T, option *meta.GetOption) (*T, error) {
	var result T
	db := p.DB
	var err error
	if option != nil {
		for _, s := range option.Preload {
			db = db.Preload(s)
		}
		if option.Include != nil {
			db = db.Where(t, option.Include)
		}
		err = db.Where(t).First(&result).Error
	} else {
		err = db.Where(t).First(&result).Error
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = myErrors.RecordNotFound
	}
	return &result, err
}
