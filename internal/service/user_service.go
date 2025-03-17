package service

import (
	"context"
	"errors"
	"myapp/internal/dto/request"
	"myapp/internal/dto/response"
	"myapp/internal/models"
	"myapp/internal/repository"
	"myapp/pkg/logger"
)

// UserService 用户业务逻辑层
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService 创建用户服务实例
func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(),
	}
}

// Register 用户注册
func (s *UserService) Register(ctx context.Context, req *request.UserRegisterRequest) (*models.User, error) {
	logger.Info("用户注册请求: %s, 昵称: %s", req.Email, req.Nickname)

	// 检查邮箱是否已存在
	existingUser, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		logger.Warn("邮箱已被注册: %s", req.Email)
		return nil, errors.New("邮箱已被注册")
	}

	// 创建新用户
	user := &models.User{
		Email:    req.Email,
		Nickname: req.Nickname,
		Password: req.Password,
	}

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		logger.Error("用户注册失败: %s, 错误: %v", req.Email, err)
		return nil, err
	}

	logger.Info("用户注册成功: %s, ID: %s", user.Email, user.ID.Hex())
	return user, nil
}

// Login 用户登录
func (s *UserService) Login(ctx context.Context, req *request.UserLoginRequest) (*response.UserLoginResponse, error) {
	logger.Info("用户登录请求: %s", req.Email)

	// 查找用户
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		logger.Warn("登录失败，用户不存在: %s", req.Email)
		return nil, errors.New("用户不存在")
	}

	// 验证密码
	if !s.userRepo.CheckPassword(req.Password, user.Password) {
		logger.Warn("登录失败，密码错误: %s", req.Email)
		return nil, errors.New("密码错误")
	}

	// 生成JWT令牌
	token, err := GenerateToken(user.ID.Hex(), user.Email)
	if err != nil {
		logger.Error("生成令牌失败: %s, 错误: %v", req.Email, err)
		return nil, err
	}

	logger.Info("用户登录成功: %s, ID: %s", user.Email, user.ID.Hex())

	// 创建登录响应
	loginResp := response.NewUserLoginResponse(user, token)
	return &loginResp, nil
}

// GetUserByID 通过ID获取用户
func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	logger.Debug("获取用户信息，ID: %s", id)
	return s.userRepo.FindByID(ctx, id)
}
