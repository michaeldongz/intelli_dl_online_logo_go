package response

// CodeResponse 验证码响应
type CodeResponse struct {
	Message string `json:"message"` // 响应消息
}

// NewCodeResponse 创建验证码响应
func NewCodeResponse(message string) CodeResponse {
	return CodeResponse{
		Message: message,
	}
}
