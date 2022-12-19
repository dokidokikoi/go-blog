package controller

import (
	"go-blog/internal/db/store"
	"go-blog/internal/service"
)

type Controller struct {
}

func (c *Controller) NewController(store *store.Factory) service.Service {
	return service.NewService(store)
}
