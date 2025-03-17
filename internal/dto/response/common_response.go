package response

// CommonResponse 通用响应结构
type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 成功响应
func Success(data interface{}) CommonResponse {
	return CommonResponse{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

// Error 错误响应
func Error(code int, message string) CommonResponse {
	return CommonResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
