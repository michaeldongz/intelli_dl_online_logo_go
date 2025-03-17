package constants

// 邮件模板类型常量
const (
	// 验证码邮件模板
	EMAIL_TEMPLATE_VERIFY_CODE = "verify_code"
	// 欢迎邮件模板
	EMAIL_TEMPLATE_WELCOME = "welcome"
	// 密码重置邮件模板
	EMAIL_TEMPLATE_RESET_PASSWORD = "reset_password"
)

// 邮件相关错误信息
const (
	MSG_EMAIL_SEND_FAILED = "邮件发送失败"
)

// GetEmailTemplate 根据模板类型获取邮件模板
func GetEmailTemplate(templateType string, params map[string]string) string {
	switch templateType {
	case EMAIL_TEMPLATE_VERIFY_CODE:
		return getVerifyCodeTemplate(params)
	case EMAIL_TEMPLATE_WELCOME:
		return getWelcomeTemplate(params)
	case EMAIL_TEMPLATE_RESET_PASSWORD:
		return getResetPasswordTemplate(params)
	default:
		return ""
	}
}

// getVerifyCodeTemplate 获取验证码邮件模板
func getVerifyCodeTemplate(params map[string]string) string {
	code := params["code"]
	return `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>邮箱注册验证码</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
        }
        .container {
            border: 1px solid #ddd;
            border-radius: 5px;
            padding: 20px;
            background-color: #f9f9f9;
        }
        .code {
            font-size: 24px;
            font-weight: bold;
            color: #007bff;
            letter-spacing: 5px;
            text-align: center;
            margin: 20px 0;
            padding: 10px;
            background-color: #e9f5ff;
            border-radius: 5px;
        }
        .footer {
            font-size: 12px;
            color: #999;
            margin-top: 30px;
            text-align: center;
        }
    </style>
</head>
<body>
    <div class="container">
        <h2>邮箱注册验证码</h2>
        <p>您好，</p>
        <p>您的验证码如下，有效期为5分钟，请勿泄露给他人：</p>
        <div class="code">` + code + `</div>
        <p>如果您没有请求此验证码，请忽略此邮件。</p>
        <p>谢谢！</p>
        <div class="footer">
            <p>此邮件由系统自动发送，请勿回复。</p>
        </div>
    </div>
</body>
</html>
`
}

// getWelcomeTemplate 获取欢迎邮件模板
func getWelcomeTemplate(params map[string]string) string {
	username := params["username"]
	return `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>欢迎加入</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
        }
        .container {
            border: 1px solid #ddd;
            border-radius: 5px;
            padding: 20px;
            background-color: #f9f9f9;
        }
        .footer {
            font-size: 12px;
            color: #999;
            margin-top: 30px;
            text-align: center;
        }
    </style>
</head>
<body>
    <div class="container">
        <h2>欢迎加入</h2>
        <p>亲爱的 ` + username + `，</p>
        <p>感谢您注册我们的服务！我们很高兴您能加入我们的社区。</p>
        <p>如果您有任何问题或需要帮助，请随时联系我们的客服团队。</p>
        <p>祝您使用愉快！</p>
        <div class="footer">
            <p>此邮件由系统自动发送，请勿回复。</p>
        </div>
    </div>
</body>
</html>
`
}

// getResetPasswordTemplate 获取密码重置邮件模板
func getResetPasswordTemplate(params map[string]string) string {
	link := params["link"]
	return `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>密码重置</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 600px;
            margin: 0 auto;
            padding: 20px;
        }
        .container {
            border: 1px solid #ddd;
            border-radius: 5px;
            padding: 20px;
            background-color: #f9f9f9;
        }
        .button {
            display: inline-block;
            padding: 10px 20px;
            background-color: #007bff;
            color: white;
            text-decoration: none;
            border-radius: 5px;
            margin: 20px 0;
        }
        .footer {
            font-size: 12px;
            color: #999;
            margin-top: 30px;
            text-align: center;
        }
    </style>
</head>
<body>
    <div class="container">
        <h2>密码重置</h2>
        <p>您好，</p>
        <p>我们收到了您的密码重置请求。请点击下面的按钮重置您的密码：</p>
        <p><a href="` + link + `" class="button">重置密码</a></p>
        <p>如果您没有请求重置密码，请忽略此邮件。</p>
        <p>谢谢！</p>
        <div class="footer">
            <p>此邮件由系统自动发送，请勿回复。</p>
            <p>链接有效期为24小时。</p>
        </div>
    </div>
</body>
</html>
`
}
