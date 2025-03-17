package controller

import (
	"intelli_dl_onling_logo/internal/utils"
	"intelli_dl_onling_logo/pkg/logger"

	"github.com/gin-gonic/gin"
)

// TestController 测试控制器
type TestController struct{}

// NewTestController 创建测试控制器实例
func NewTestController() *TestController {
	return &TestController{}
}

// Test 测试接口
func (c *TestController) Test(ctx *gin.Context) {
	clientIP := ctx.ClientIP()
	logger.Info("收到测试请求，客户端IP: %s", clientIP)

	utils.Success(ctx, gin.H{
		"message": "测试接口调用成功",
		"time":    utils.GetCurrentTime(),
		"ip":      clientIP,
	})
}
