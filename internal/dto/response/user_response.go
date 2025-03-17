package response

import (
	"intelli_dl_onling_logo/internal/models"
)

// UserResponse 用户信息响应
type UserResponse struct {
	ID        string `json:"id"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Role      int    `json:"role"`
	CreatedAt string `json:"created_at"`
}

// UserLoginResponse 用户登录响应
type UserLoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

// NewUserResponse 从用户模型创建用户响应
func NewUserResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:        user.ID.Hex(),
		Nickname:  user.Nickname,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

// NewUserLoginResponse 从用户模型和令牌创建登录响应
func NewUserLoginResponse(user *models.User, token string) UserLoginResponse {
	return UserLoginResponse{
		User:  NewUserResponse(user),
		Token: token,
	}
}
