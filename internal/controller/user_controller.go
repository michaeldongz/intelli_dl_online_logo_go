package controller

import (
	"intelli_dl_onling_logo/internal/constants"
	"intelli_dl_onling_logo/internal/dto/request"
	"intelli_dl_onling_logo/internal/dto/response"
	"intelli_dl_onling_logo/internal/service"
	"intelli_dl_onling_logo/internal/utils"
	"intelli_dl_onling_logo/pkg/logger"

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

// GetUserInfo 获取当前用户信息
func (c *UserController) GetUserInfo(ctx *gin.Context) {
	// 从上下文中获取用户ID
	userID, _ := ctx.Get("userID")

	logger.Info("获取用户信息请求，用户ID: %s", userID)

	// 获取用户信息
	user, err := c.userService.GetUserByID(ctx, userID.(string))
	if err != nil {
		logger.Error("获取用户信息失败: %v", err)
		utils.NotFound(ctx, constants.MSG_USER_NOT_EXIST)
		return
	}

	// 转换为响应格式
	userResp := response.NewUserResponse(user)
	utils.Success(ctx, userResp)
}
