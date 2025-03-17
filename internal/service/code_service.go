package service

import (
	"context"
	"errors"
	"intelli_dl_onling_logo/internal/constants"
	"intelli_dl_onling_logo/internal/dto/request"
	"intelli_dl_onling_logo/internal/dto/response"
	"intelli_dl_onling_logo/internal/models"
	"intelli_dl_onling_logo/internal/repository"
	"intelli_dl_onling_logo/internal/utils"
	"intelli_dl_onling_logo/pkg/logger"
	"math/rand"
	"time"
)

// CodeService 验证码业务逻辑层
type CodeService struct {
	codeRepo    *repository.CodeRepository
	emailClient *utils.EmailClient
}

// NewCodeService 创建验证码服务实例
func NewCodeService() *CodeService {
	return &CodeService{
		codeRepo:    repository.NewCodeRepository(),
		emailClient: utils.NewEmailClient(),
	}
}

// SendEmailCode 发送邮箱验证码
func (s *CodeService) SendEmailCode(ctx context.Context, req *request.SendCodeRequest) (*response.CodeResponse, error) {
	logger.Info("发送邮箱验证码请求: %s", req.Email)

	// 检查邮箱是否已有验证码
	exists, err := s.codeRepo.CheckEmailHasCode(ctx, req.Email)
	if err != nil && err != context.DeadlineExceeded {
		logger.Error("检查邮箱验证码失败: %s, 错误: %v", req.Email, err)
		return nil, errors.New(constants.MSG_EMAIL_FORMAT_ERROR)
	}

	if exists {
		logger.Warn("邮箱已存在未过期的验证码: %s", req.Email)
		return nil, errors.New(constants.MSG_EMAIL_CODE_EXIST)
	}

	// 生成6位随机验证码
	verifyCode := s.generateRandomCode(6)
	logger.Debug("生成验证码: %s, 邮箱: %s", verifyCode, req.Email)

	// 设置验证码过期时间为5分钟
	expiration := 5 * time.Minute
	now := time.Now()
	expiredAt := now.Add(expiration)

	// 创建验证码记录
	codeModel := &models.Code{
		Type:      models.CODE_TYPE_EMAIL,
		Code:      verifyCode,
		Content:   "您的验证码是: " + verifyCode + "，有效期5分钟，请勿泄露给他人。",
		Email:     req.Email,
		ValidFrom: now,
		ExpiredAt: expiredAt,
	}

	// 发送验证码邮件
	params := map[string]string{
		"code": verifyCode,
	}
	emailTemplate := constants.GetEmailTemplate(constants.EMAIL_TEMPLATE_VERIFY_CODE, params)

	// 发送HTML邮件
	err = s.emailClient.SendHTMLEmail([]string{req.Email}, "验证码", emailTemplate)
	if err != nil {
		logger.Error("发送验证码邮件失败: %s, 错误: %v", req.Email, err)
		return nil, errors.New(constants.MSG_EMAIL_SEND_FAILED)
	}
	logger.Info("验证码邮件发送成功: %s", req.Email)

	// 保存到MongoDB
	err = s.codeRepo.Create(ctx, codeModel)
	if err != nil {
		logger.Error("保存验证码记录失败: %s, 错误: %v", req.Email, err)
		return nil, errors.New(constants.MSG_SERVER_ERROR)
	}

	// 保存到Redis，设置过期时间
	err = s.codeRepo.SaveToRedis(ctx, req.Email, verifyCode, expiration)
	if err != nil {
		logger.Error("保存验证码到Redis失败: %s, 错误: %v", req.Email, err)
		return nil, errors.New(constants.MSG_SERVER_ERROR)
	}

	logger.Info("验证码发送成功: %s", req.Email)
	return &response.CodeResponse{Message: "验证码发送成功，请查收邮件"}, nil
}

// generateRandomCode 生成指定长度的随机验证码
func (s *CodeService) generateRandomCode(length int) string {
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(code)
}
