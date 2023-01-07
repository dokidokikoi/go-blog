package category

import (
	"context"
	"fmt"
	"go-blog/internal/db/model/category"
	"go-blog/internal/db/store"
	meta "go-blog/pkg/meta/option"
)

type CategorySrv interface {
	Create(ctx context.Context, a *category.Category) error
	IsExist(ctx context.Context, a *category.Category) (bool, error)
	Update(ctx context.Context, a *category.Category) error
	DeleteById(ctx context.Context, id uint) error
	DeleteByIds(ctx context.Context, ids []uint) error
	GetById(ctx context.Context, id uint, option *meta.GetOption) (*category.Category, error)
	List(ctx context.Context, keyword string, t int, option *meta.ListOption) ([]*category.Category, int64, error)
}

type categorySrv struct {
	store store.Factory
}

func (c categorySrv) Create(ctx context.Context, a *category.Category) error {
	a.ID = 0
	return c.store.Category().Create(ctx, a, nil)
}

func (c categorySrv) IsExist(ctx context.Context, a *category.Category) (bool, error) {
	if a.ID == 0 {
		cate, err := c.store.Category().Get(ctx, a, nil)
		if err != nil {
			return false, err
		}
		if cate == nil {
			return false, nil
		}
		a.ID = cate.ID
		return true, nil
	}
	cate, err := c.store.Category().Get(ctx, &category.Category{ID: a.ID}, &meta.GetOption{Include: []string{"id"}})
	if err != nil {
		return false, err
	}
	if cate == nil {
		return false, err
	}
	return true, nil
}

func (c categorySrv) Update(ctx context.Context, a *category.Category) error {
	return c.store.Category().Update(ctx, a, nil)
}

func (c categorySrv) DeleteById(ctx context.Context, id uint) error {
	return c.store.Category().Delete(ctx, &category.Category{ID: id}, nil)
}

func (c categorySrv) DeleteByIds(ctx context.Context, ids []uint) error {
	return c.store.Category().DeleteByIds(ctx, ids)
}

func (c categorySrv) GetById(ctx context.Context, id uint, option *meta.GetOption) (*category.Category, error) {
	return c.store.Category().Get(ctx, &category.Category{ID: id}, option)
}

func (c categorySrv) List(ctx context.Context, keyword string, t int, option *meta.ListOption) ([]*category.Category, int64, error) {
	if keyword == "" {
		count, _ := c.store.Category().Count(ctx, &category.Category{}, &option.GetOption)
		result, err := c.store.Category().List(ctx, &category.Category{}, option)
		return result, count, err
	}

	root := &meta.WhereNode{}
	root.Conditions = append(root.Conditions,
		&meta.Condition{Field: "category_name", Operator: meta.LIKE, Value: fmt.Sprintf("%s%s%s", "%", keyword, "%")})
	root.Next = &meta.WhereNode{Conditions: []*meta.Condition{{Field: "type", Operator: meta.EQUAL, Value: t}}}
	count, _ := c.store.Category().CountComplex(ctx, &category.Category{}, root, &option.GetOption)
	result, err := c.store.Category().ListComplex(ctx, &category.Category{}, root, option)
	return result, count, err
}

func NewCategorySrv(store store.Factory) CategorySrv {
	return &categorySrv{store: store}
}
