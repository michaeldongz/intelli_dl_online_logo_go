package database

import (
	"context"
	"fmt"
	"intelli_dl_onling_logo/config"
	"intelli_dl_onling_logo/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

func InitMongoDB() error {
	logger.Debug("正在连接MongoDB: %s", config.GlobalConfig.MongoDB.URI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GlobalConfig.MongoDB.URI))
	if err != nil {
		logger.Error("MongoDB连接失败: %v", err)
		return fmt.Errorf("MongoDB连接失败: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Error("MongoDB Ping失败: %v", err)
		return fmt.Errorf("MongoDB Ping失败: %w", err)
	}

	MongoClient = client
	MongoDB = client.Database(config.GlobalConfig.MongoDB.Database)
	logger.Info("MongoDB连接成功，数据库: %s", config.GlobalConfig.MongoDB.Database)
	return nil
}
