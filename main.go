package main

import (
	"flag"
	"fmt"
	"intelli_dl_onling_logo/config"
	"intelli_dl_onling_logo/internal/database"
	"intelli_dl_onling_logo/pkg/logger"
	"intelli_dl_onling_logo/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 通过命令行参数指定环境
	env := flag.String("env", "dev", "运行环境，可选：dev/prod")
	flag.Parse()

	// 初始化配置
	if err := config.InitConfig(*env); err != nil {
		log.Fatalf("配置初始化失败: %v", err)
	}

	// 初始化日志
	logConfig := logger.Config{
		Level:      config.GlobalConfig.Log.Level,
		Format:     config.GlobalConfig.Log.Format,
		Output:     config.GlobalConfig.Log.Output,
		Directory:  config.GlobalConfig.Log.Directory,
		Filename:   config.GlobalConfig.Log.Filename,
		MaxSize:    config.GlobalConfig.Log.MaxSize,
		MaxAge:     config.GlobalConfig.Log.MaxAge,
		MaxBackups: config.GlobalConfig.Log.MaxBackups,
		Compress:   config.GlobalConfig.Log.Compress,
	}
	if err := logger.InitLogger(logConfig); err != nil {
		log.Fatalf("日志初始化失败: %v", err)
	}
	defer logger.Sync()

	logger.Info("服务启动中，环境: %s", *env)

	// 初始化Redis
	if err := database.InitRedis(); err != nil {
		logger.Fatal("Redis初始化失败: %v", err)
	}
	logger.Info("Redis连接成功")

	// 初始化MongoDB
	if err := database.InitMongoDB(); err != nil {
		logger.Fatal("MongoDB初始化失败: %v", err)
	}
	logger.Info("MongoDB连接成功")

	// 设置gin模式
	gin.SetMode(config.GlobalConfig.Server.Mode)

	// 初始化路由
	r := router.InitRouter()

	// 启动服务
	addr := fmt.Sprintf(":%d", config.GlobalConfig.Server.Port)
	logger.Info("服务启动在 %s 环境，监听端口 %s", *env, addr)
	if err := r.Run(addr); err != nil {
		logger.Fatal("服务启动失败: %v", err)
	}
}
