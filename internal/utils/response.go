package utils

import (
	"myapp/internal/dto/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success 返回成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, response.Success(data))
}

// ErrorResponse 返回错误响应
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, response.Error(code, message))
}
