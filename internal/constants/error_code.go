package constants

// 响应码常量
const (
	// 成功
	SUCCESS = 200

	// 客户端错误 (4xx)
	BAD_REQUEST          = 400
	UNAUTHORIZED         = 401
	FORBIDDEN            = 403
	NOT_FOUND            = 404
	METHOD_NOT_ALLOWED   = 405
	REQUEST_TIMEOUT      = 408
	CONFLICT             = 409
	UNPROCESSABLE_ENTITY = 422
	TOO_MANY_REQUESTS    = 429

	// 服务器错误 (5xx)
	INTERNAL_SERVER_ERROR = 500
	SERVICE_UNAVAILABLE   = 503
	GATEWAY_TIMEOUT       = 504
)
