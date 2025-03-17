package middleware

import (
	"bytes"
	"intelli_dl_onling_logo/pkg/logger"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 请求方法
		method := c.Request.Method

		// 请求路径
		path := c.Request.URL.Path

		// 请求IP
		clientIP := c.ClientIP()

		// 如果是POST/PUT/PATCH请求，记录请求体
		if method == "POST" || method == "PUT" || method == "PATCH" {
			// 读取请求体
			var bodyBytes []byte
			if c.Request.Body != nil {
				bodyBytes, _ = io.ReadAll(c.Request.Body)
			}

			// 重新设置请求体，因为读取后会清空
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			// 记录请求体，但不记录敏感信息
			if len(bodyBytes) > 0 {
				// 这里可以添加敏感信息过滤逻辑
				logger.Debug("请求体: %s", string(bodyBytes))
			}
		}

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latency := endTime.Sub(startTime)

		// 状态码
		statusCode := c.Writer.Status()

		// 记录请求日志
		logger.Info("请求: %s %s | 状态: %d | IP: %s | 耗时: %v",
			method, path, statusCode, clientIP, latency)
	}
}
