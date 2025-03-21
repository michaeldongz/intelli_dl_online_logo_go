package middleware

import (
	"intelli_dl_onling_logo/internal/constants"
	"intelli_dl_onling_logo/internal/service"
	"intelli_dl_onling_logo/internal/utils"
	"intelli_dl_onling_logo/pkg/logger"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth JWT认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Warn("未提供认证令牌，IP: %s, 路径: %s", c.ClientIP(), c.Request.URL.Path)
			utils.Unauthorized(c, constants.MSG_TOKEN_REQUIRED)
			c.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			logger.Warn("认证格式错误，IP: %s, 路径: %s", c.ClientIP(), c.Request.URL.Path)
			utils.Unauthorized(c, constants.MSG_TOKEN_FORMAT_ERROR)
			c.Abort()
			return
		}

		// 解析令牌
		claims, err := service.ParseToken(parts[1])
		if err != nil {
			logger.Warn("无效的令牌: %v, IP: %s, 路径: %s", err, c.ClientIP(), c.Request.URL.Path)
			utils.Unauthorized(c, constants.MSG_TOKEN_INVALID+": "+err.Error())
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		logger.Debug("用户认证成功: %s, ID: %s, 角色: %d, IP: %s, 路径: %s",
			claims.Email, claims.UserID, claims.Role, c.ClientIP(), c.Request.URL.Path)
		c.Next()
	}
}
