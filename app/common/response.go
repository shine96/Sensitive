package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
