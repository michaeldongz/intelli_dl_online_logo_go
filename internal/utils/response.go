package utils

import (
	"myapp/internal/constants"
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

// BadRequest 返回请求参数错误响应
func BadRequest(c *gin.Context, err error) {
	message := constants.MSG_PARAM_ERROR
	if err != nil {
		message = constants.MSG_PARAM_ERROR + ": " + err.Error()
	}
	ErrorResponse(c, constants.BAD_REQUEST, message)
}

// Unauthorized 返回未授权错误响应
func Unauthorized(c *gin.Context, message string) {
	if message == "" {
		message = constants.MSG_UNAUTHORIZED
	}
	ErrorResponse(c, constants.UNAUTHORIZED, message)
}

// Forbidden 返回禁止访问错误响应
func Forbidden(c *gin.Context, message string) {
	if message == "" {
		message = constants.MSG_FORBIDDEN
	}
	ErrorResponse(c, constants.FORBIDDEN, message)
}

// NotFound 返回资源不存在错误响应
func NotFound(c *gin.Context, message string) {
	if message == "" {
		message = constants.MSG_NOT_FOUND
	}
	ErrorResponse(c, constants.NOT_FOUND, message)
}

// ServerError 返回服务器内部错误响应
func ServerError(c *gin.Context, err error) {
	message := constants.MSG_SERVER_ERROR
	if err != nil {
		message = constants.MSG_SERVER_ERROR + ": " + err.Error()
	}
	ErrorResponse(c, constants.INTERNAL_SERVER_ERROR, message)
}
