package controller

import (
	"myapp/internal/constants"
	"myapp/internal/dto/request"
	"myapp/internal/dto/response"
	"myapp/internal/service"
	"myapp/internal/utils"
	"myapp/pkg/logger"

	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct {
	userService *service.UserService
}

// NewUserController 创建用户控制器实例
func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// Register 用户注册
func (c *UserController) Register(ctx *gin.Context) {
	var req request.UserRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Warn("注册请求参数错误: %v", err)
		utils.BadRequest(ctx, err)
		return
	}

	logger.Info("收到注册请求: %s", req.Email)
	user, err := c.userService.Register(ctx, &req)
	if err != nil {
		logger.Error("注册失败: %s, 错误: %v", req.Email, err)
		utils.ErrorResponse(ctx, constants.INTERNAL_SERVER_ERROR, constants.MSG_REGISTER_FAILED+": "+err.Error())
		return
	}

	logger.Info("注册成功: %s", req.Email)
	// 转换为响应格式
	userResp := response.NewUserResponse(user)
	utils.Success(ctx, userResp)
}

// Login 用户登录
func (c *UserController) Login(ctx *gin.Context) {
	var req request.UserLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Warn("登录请求参数错误: %v", err)
		utils.BadRequest(ctx, err)
		return
	}

	logger.Info("收到登录请求: %s", req.Email)
	resp, err := c.userService.Login(ctx, &req)
	if err != nil {
		logger.Warn("登录失败: %s, 错误: %v", req.Email, err)
		utils.Unauthorized(ctx, constants.MSG_LOGIN_FAILED+": "+err.Error())
		return
	}

	logger.Info("登录成功: %s", req.Email)
	utils.Success(ctx, resp)
}
