package router

import (
	"myapp/internal/controller"
	"myapp/internal/middleware"
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

	// 用户相关路由
	userGroup := r.Group("/api/user")
	{
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/login", userController.Login)
	}

	// 需要认证的路由
	authGroup := r.Group("/api")
	authGroup.Use(middleware.JWTAuth())
	{
		// 这里可以添加需要认证的路由
	}

	// 测试路由
	testGroup := r.Group("/api/test")
	{
		testGroup.GET("/", testController.Test)
	}

	logger.Info("路由初始化完成")
	return r
}
