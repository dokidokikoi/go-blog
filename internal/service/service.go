package service

import (
	"go-blog/internal/db/store"
	"sync"
)

type Service interface {
}

type service struct {
	store store.Factory
}

var (
	serviceClient Service
	once          sync.Once
)

func NewService(store *store.Factory) Service {
	once.Do(func() {
		serviceClient = &service{
			store: store,
		}
	})

	return serviceClient
}
