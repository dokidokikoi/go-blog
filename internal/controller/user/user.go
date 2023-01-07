package user

import (
	"go-blog/internal/db/store"
	"go-blog/internal/service"
)

type Controller struct {
	srv service.Service
}

func NewController(store store.Factory) *Controller {
	return &Controller{service.NewService(store)}
}
