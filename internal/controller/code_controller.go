package controller

import (
	"intelli_dl_onling_logo/internal/dto/request"
	"intelli_dl_onling_logo/internal/service"
	"intelli_dl_onling_logo/internal/utils"
	"intelli_dl_onling_logo/pkg/logger"

	"github.com/gin-gonic/gin"
)

// CodeController 验证码控制器
type CodeController struct {
	codeService *service.CodeService
}

// NewCodeController 创建验证码控制器实例
func NewCodeController() *CodeController {
	return &CodeController{
		codeService: service.NewCodeService(),
	}
}

// SendEmailCode 发送邮箱验证码
func (c *CodeController) SendEmailCode(ctx *gin.Context) {
	var req request.SendCodeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Warn("发送验证码请求参数错误: %v", err)
		utils.BadRequest(ctx, err)
		return
	}

	logger.Info("收到发送验证码请求: %s", req.Email)
	resp, err := c.codeService.SendEmailCode(ctx, &req)
	if err != nil {
		logger.Error("发送验证码失败: %s, 错误: %v", req.Email, err)
		utils.ErrorResponse(ctx, 500, err.Error())
		return
	}

	logger.Info("验证码发送成功: %s", req.Email)
	utils.Success(ctx, resp)
}
