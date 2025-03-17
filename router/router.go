package router

import (
	"myapp/internal/controller"
	"myapp/internal/middleware"
	"myapp/internal/models"
	"myapp/pkg/logger"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// 使用自定义日志中间件替代gin.Logger()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	// 控制器实例
	userController := controller.NewUserController()
	testController := controller.NewTestController()

	logger.Info("初始化路由...")

	// 健康检查
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 用户相关路由 - 无需认证
	userGroup := r.Group("/api/user")
	{
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/login", userController.Login)
	}

	// 需要认证的路由
	authGroup := r.Group("/api")
	authGroup.Use(middleware.JWTAuth()) // 所有路由都需要JWT认证
	{
		// 获取当前用户信息
		authGroup.GET("/user/info", userController.GetUserInfo)
	}

	// 测试路由 - 需要管理员或超级管理员权限
	testGroup := r.Group("/api/test")
	testGroup.Use(middleware.JWTAuth())                                             // 先进行JWT认证
	testGroup.Use(middleware.CheckRole(models.ROLE_ADMIN, models.ROLE_SUPER_ADMIN)) // 然后检查角色权限
	{
		testGroup.GET("/", testController.Test)
	}

	logger.Info("路由初始化完成")
	return r
}
