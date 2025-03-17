package middleware

import (
	"myapp/internal/constants"
	"myapp/internal/utils"
	"myapp/pkg/logger"

	"github.com/gin-gonic/gin"
)

// CheckRole 检查用户角色权限
func CheckRole(allowedRoles ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文中获取用户角色
		role, exists := c.Get("role")
		if !exists {
			logger.Warn("权限检查失败，未找到用户角色信息，路径: %s", c.Request.URL.Path)
			utils.Forbidden(c, constants.MSG_PERMISSION_DENIED)
			c.Abort()
			return
		}

		userRole := role.(int)
		hasPermission := false

		// 检查用户角色是否在允许的角色列表中
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			logger.Warn("权限不足，用户角色: %d, 需要角色: %v, 路径: %s",
				userRole, allowedRoles, c.Request.URL.Path)
			utils.Forbidden(c, constants.MSG_PERMISSION_DENIED)
			c.Abort()
			return
		}

		logger.Debug("权限检查通过，用户角色: %d, 路径: %s", userRole, c.Request.URL.Path)
		c.Next()
	}
}
