package core

import (
	"net/http"

	"go-blog/internal/errors"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WriteResponse(c *gin.Context, appError *errors.APPError, data interface{}) {
	if appError == nil {
		c.JSON(http.StatusOK, Response{Data: data})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    appError.Code,
		Message: appError.Message,
		Data:    data,
	})
}

func WriteListResponse(c *gin.Context, appError *errors.APPError, total int64, data interface{}) {
	if appError == nil {
		c.JSON(http.StatusOK, Response{Data: map[string]any{
			"total": total,
			"list":  data,
		}})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    appError.Code,
		Message: appError.Message,
		Data:    data,
	})
}
