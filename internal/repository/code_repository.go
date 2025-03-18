package repository

import (
	"context"
	"fmt"
	"intelli_dl_onling_logo/internal/database"
	"intelli_dl_onling_logo/internal/models"
	"intelli_dl_onling_logo/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

// Delete 删除验证码
func (r *CodeRepository) Delete(ctx context.Context, key string) error {
	return r.redis.Delete(ctx, key)
}

// UpdateCodeUsed 更新验证码使用记录
func (r *CodeRepository) UpdateCodeUsed(ctx context.Context, code *models.Code) error {
	logger.Debug("开始 更新验证码使用记录: %s", code.Email)

	currentTime := time.Now()
	update := bson.M{
		"$set": bson.M{
			"updated_at": currentTime,
			"used_at":    &currentTime,
			"status":     models.CODE_STATUS_USED,
		},
	}

	_, err := r.mongo.UpdateByID(ctx, string(code.ID.Hex()), update)
	if err != nil {
		logger.Error("更新验证码使用记录失败: %v", err)
		return err
	}

	logger.Info("更新验证码使用记录成功: %s, ID: %s", code.Email, code.ID.Hex())
	return nil
}

// FindByEmailAndCode 通过邮箱和验证码查找验证码记录
func (r *CodeRepository) FindByEmailAndCode(ctx context.Context, email string, code string) (*models.Code, error) {
	logger.Debug("开始 通过邮箱和验证码查找验证码记录: %s, %s", email, code)

	filter := bson.M{
		"email": email,
		"code":  code,
	}

	var codeRecord models.Code
	err := r.mongo.FindOne(ctx, filter, &codeRecord)
	if err != nil {
		logger.Error("通过邮箱和验证码查找验证码记录失败: %v", err)
		return &models.Code{}, err
	}

	return &codeRecord, nil
}
