package base

import (
	"context"

	meta "go-blog/pkg/meta/option"

	"gorm.io/gorm"
)

type BaseModel any

type PgModel[T BaseModel] struct {
	DB *gorm.DB
}

type CreateMixin[T BaseModel] interface {
	Create(ctx context.Context, t *T, option *meta.CreateOption) error
	CreateCollection(ctx context.Context, t []*T, option *meta.CreateCollectionOption) []error
}

type UpdateMixin[T BaseModel] interface {
	Update(ctx context.Context, t *T, option *meta.UpdateOption) error
	UpdateByWhere(ctx context.Context, node *meta.WhereNode, example *T, option *meta.UpdateOption) error
	UpdateCollection(ctx context.Context, t []*T, option *meta.UpdateCollectionOption) []error
	Save(ctx context.Context, t *T, option *meta.UpdateOption) error
}

type GetMixin[T BaseModel] interface {
	Get(ctx context.Context, t *T, option *meta.GetOption) (*T, error)
	Count(ctx context.Context, t *T, option *meta.GetOption) (int64, error)
	CountComplex(ctx context.Context, example *T, condition *meta.WhereNode, option *meta.GetOption) (int64, error)
	List(ctx context.Context, t *T, option *meta.ListOption) ([]*T, error)
	ListComplex(ctx context.Context, example *T, condition *meta.WhereNode, option *meta.ListOption) ([]*T, error)
}

type DeleteMixin[T BaseModel] interface {
	Delete(ctx context.Context, t *T, option *meta.DeleteOption) error
	DeleteCollection(ctx context.Context, t []*T, option *meta.DeleteCollectionOption) []error
	DeleteByIds(ctx context.Context, ids []uint) error
}

type BasicCURD[T BaseModel] interface {
	CreateMixin[T]
	DeleteMixin[T]
	UpdateMixin[T]
	GetMixin[T]
}
