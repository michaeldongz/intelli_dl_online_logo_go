package constants

// 响应信息常量
const (
	// 成功
	MSG_SUCCESS = "success"

	// 通用错误
	MSG_SERVER_ERROR     = "服务器内部错误"
	MSG_PARAM_ERROR      = "请求参数错误"
	MSG_NOT_FOUND        = "资源不存在"
	MSG_UNAUTHORIZED     = "未授权访问"
	MSG_FORBIDDEN        = "禁止访问"
	MSG_TOO_MANY_REQUEST = "请求过于频繁"

	// 用户相关错误
	MSG_USER_NOT_EXIST     = "用户不存在"
	MSG_PASSWORD_ERROR     = "密码错误"
	MSG_EMAIL_REGISTERED   = "邮箱已被注册"
	MSG_REGISTER_FAILED    = "注册失败"
	MSG_LOGIN_FAILED       = "登录失败"
	MSG_TOKEN_INVALID      = "无效的令牌"
	MSG_TOKEN_EXPIRED      = "令牌已过期"
	MSG_TOKEN_REQUIRED     = "未提供认证令牌"
	MSG_TOKEN_FORMAT_ERROR = "认证格式错误"

	// 权限相关错误
	MSG_PERMISSION_DENIED = "权限不足"
)
