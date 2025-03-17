package repository

import (
	"context"
	"intelli_dl_onling_logo/internal/database"
	"intelli_dl_onling_logo/internal/models"
	"intelli_dl_onling_logo/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository 用户数据访问层
type UserRepository struct {
	mongo *database.MongoUtils
}

// NewUserRepository 创建用户仓库实例
func NewUserRepository() *UserRepository {
	return &UserRepository{
		mongo: database.NewMongoUtils("users"),
	}
}

// Create 创建用户
func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	logger.Debug("开始创建用户: %s", user.Email)

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("密码加密失败: %v", err)
		return err
	}

	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	id, err := r.mongo.InsertOne(ctx, user)
	if err != nil {
		logger.Error("用户创建失败: %v", err)
		return err
	}

	user.ID = id
	logger.Info("用户创建成功: %s, ID: %s", user.Email, user.ID.Hex())
	return nil
}

// FindByEmail 通过邮箱查找用户
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	logger.Debug("通过邮箱查找用户: %s", email)

	var user models.User
	err := r.mongo.FindOne(ctx, bson.M{"email": email}, &user)
	if err != nil {
		logger.Debug("未找到用户: %s, 错误: %v", email, err)
		return nil, err
	}

	logger.Debug("找到用户: %s, ID: %s", email, user.ID.Hex())
	return &user, nil
}

// FindByID 通过ID查找用户
func (r *UserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	logger.Debug("通过ID查找用户: %s", id)

	var user models.User
	err := r.mongo.FindByID(ctx, id, &user)
	if err != nil {
		logger.Debug("未找到用户ID: %s, 错误: %v", id, err)
		return nil, err
	}

	logger.Debug("找到用户ID: %s, 邮箱: %s", id, user.Email)
	return &user, nil
}

// CheckPassword 检查密码是否正确
func (r *UserRepository) CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
