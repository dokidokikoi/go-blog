package user

import (
	"go-blog/internal/db/store"
	"go-blog/internal/service"
)

type CreateUser struct {
	Account  string `json:"account" binding:"required"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email" binding:"required"`
	NickName string `json:"nick_name" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginParam struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type Controller struct {
	srv service.Service
}

func NewController(store store.Factory) *Controller {
	return &Controller{service.NewService(store)}
}
