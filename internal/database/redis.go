package database

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/pkg/logger"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() error {
	logger.Debug("正在连接Redis: %s:%d", config.GlobalConfig.Redis.Host, config.GlobalConfig.Redis.Port)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.GlobalConfig.Redis.Host, config.GlobalConfig.Redis.Port),
		Password: config.GlobalConfig.Redis.Password,
		DB:       config.GlobalConfig.Redis.DB,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		logger.Error("Redis连接失败: %v", err)
		return fmt.Errorf("Redis连接失败: %w", err)
	}

	logger.Info("Redis连接成功")
	return nil
}
