package repository

import (
	"context"
	"fmt"
	"intelli_dl_onling_logo/internal/database"
	"intelli_dl_onling_logo/internal/models"
	"intelli_dl_onling_logo/pkg/logger"
	"time"
)

// CodeRepository 验证码数据访问层
type CodeRepository struct {
	mongo *database.MongoUtils
	redis *database.RedisUtils
}

// NewCodeRepository 创建验证码仓库实例
func NewCodeRepository() *CodeRepository {
	return &CodeRepository{
		mongo: database.NewMongoUtils("code_log"),
		redis: database.NewRedisUtils(),
	}
}

// Create 创建验证码记录
func (r *CodeRepository) Create(ctx context.Context, code *models.Code) error {
	logger.Debug("开始创建验证码记录: %s", code.Email)

	code.CreatedAt = code.ValidFrom
	code.UpdatedAt = code.ValidFrom
	code.Status = models.CODE_STATUS_UNUSED

	id, err := r.mongo.InsertOne(ctx, code)
	if err != nil {
		logger.Error("验证码记录创建失败: %v", err)
		return err
	}

	code.ID = id
	logger.Info("验证码记录创建成功: %s, ID: %s", code.Email, code.ID.Hex())
	return nil
}

// SaveToRedis 保存验证码到Redis
func (r *CodeRepository) SaveToRedis(ctx context.Context, email string, code string, expiration time.Duration) error {
	key := fmt.Sprintf("code:email:%s", email)
	return r.redis.Set(ctx, key, code, expiration)
}

// GetFromRedis 从Redis获取验证码
func (r *CodeRepository) GetFromRedis(ctx context.Context, email string) (string, error) {
	key := fmt.Sprintf("code:email:%s", email)
	return r.redis.Get(ctx, key)
}

// CheckEmailHasCode 检查邮箱是否已有验证码
func (r *CodeRepository) CheckEmailHasCode(ctx context.Context, email string) (bool, error) {
	key := fmt.Sprintf("code:email:%s", email)
	return r.redis.Exists(ctx, key)
}
