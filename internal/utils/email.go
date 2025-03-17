package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"intelli_dl_onling_logo/config"
	"mime"
	"mime/multipart"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
)

// EmailClient 邮件客户端结构体
type EmailClient struct {
	Host     string
	Port     int
	Username string
	Password string
	FromName string
}

// NewEmailClient 创建一个新的邮件客户端
func NewEmailClient() *EmailClient {
	Debug("初始化邮件客户端")
	return &EmailClient{
		Host:     config.GlobalConfig.Email.Host,
		Port:     config.GlobalConfig.Email.Port,
		Username: config.GlobalConfig.Email.Username,
		Password: config.GlobalConfig.Email.Password,
		FromName: config.GlobalConfig.Email.FromName,
	}
}

// SendTextEmail 发送纯文本邮件
func (c *EmailClient) SendTextEmail(to []string, subject, body string) error {
	Debug("发送纯文本邮件给: %v", to)

	// 构建邮件头
	header := make(map[string]string)
	header["From"] = fmt.Sprintf("%s <%s>", c.FromName, c.Username)
	header["To"] = strings.Join(to, ",")
	header["Subject"] = subject
	header["Content-Type"] = "text/plain; charset=UTF-8"

	// 构建邮件内容
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// 发送邮件
	err := c.sendMail(to, []byte(message))
	if err != nil {
		Error("发送纯文本邮件失败: %v", err)
		return err
	}

	Info("成功发送纯文本邮件给: %v", to)
	return nil
}

// SendHTMLEmail 发送HTML邮件
func (c *EmailClient) SendHTMLEmail(to []string, subject, htmlBody string) error {
	Debug("发送HTML邮件给: %v", to)

	// 构建邮件头
	header := make(map[string]string)
	header["From"] = fmt.Sprintf("%s <%s>", c.FromName, c.Username)
	header["To"] = strings.Join(to, ",")
	header["Subject"] = subject
	header["Content-Type"] = "text/html; charset=UTF-8"

	// 构建邮件内容
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + htmlBody

	// 发送邮件
	err := c.sendMail(to, []byte(message))
	if err != nil {
		Error("发送HTML邮件失败: %v", err)
		return err
	}

	Info("成功发送HTML邮件给: %v", to)
	return nil
}

// SendEmailWithAttachment 发送带附件的邮件
func (c *EmailClient) SendEmailWithAttachment(to []string, subject, body string, attachmentPath string, isHTML bool) error {
	Debug("发送带附件的邮件给: %v, 附件: %s", to, attachmentPath)

	// 创建一个新的multipart写入器
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()

	// 构建邮件头
	header := make(map[string]string)
	header["From"] = fmt.Sprintf("%s <%s>", c.FromName, c.Username)
	header["To"] = strings.Join(to, ",")
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("multipart/mixed; boundary=\"%s\"", boundary)

	// 写入邮件头
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n"

	// 添加邮件正文
	contentType := "text/plain"
	if isHTML {
		contentType = "text/html"
	}

	// 写入正文部分
	part, _ := writer.CreatePart(map[string][]string{
		"Content-Type": {fmt.Sprintf("%s; charset=UTF-8", contentType)},
	})
	part.Write([]byte(body))

	// 添加附件
	filename := filepath.Base(attachmentPath)
	fileData, err := ReadFile(attachmentPath)
	if err != nil {
		Error("读取附件失败: %v", err)
		return err
	}

	// 写入附件部分
	part, _ = writer.CreatePart(map[string][]string{
		"Content-Type":              {fmt.Sprintf("%s; name=\"%s\"", mime.TypeByExtension(filepath.Ext(filename)), filename)},
		"Content-Disposition":       {fmt.Sprintf("attachment; filename=\"%s\"", filename)},
		"Content-Transfer-Encoding": {"base64"},
	})

	// 对附件进行base64编码
	encoded := base64.StdEncoding.EncodeToString(fileData)
	part.Write([]byte(encoded))

	// 关闭writer
	writer.Close()

	// 发送邮件
	err = c.sendMail(to, append([]byte(message), buf.Bytes()...))
	if err != nil {
		Error("发送带附件的邮件失败: %v", err)
		return err
	}

	Info("成功发送带附件的邮件给: %v, 附件: %s", to, attachmentPath)
	return nil
}

// ReadFile 读取文件内容
func ReadFile(filePath string) ([]byte, error) {
	Debug("读取文件: %s", filePath)

	// 使用os包读取文件内容
	data, err := os.ReadFile(filePath)
	if err != nil {
		Error("读取文件失败: %v", err)
		return nil, err
	}

	return data, nil
}

// sendMail 发送邮件的底层方法
func (c *EmailClient) sendMail(to []string, message []byte) error {
	// 构建邮件服务器地址
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)

	// 创建SMTP客户端
	auth := smtp.PlainAuth("", c.Username, c.Password, c.Host)

	// 尝试使用TLS发送
	tlsConfig := &tls.Config{
		ServerName: c.Host,
	}

	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("连接SMTP服务器失败: %w", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, c.Host)
	if err != nil {
		return fmt.Errorf("创建SMTP客户端失败: %w", err)
	}
	defer client.Quit()

	// 验证身份
	if err = client.Auth(auth); err != nil {
		return err
	}

	// 设置发件人
	if err = client.Mail(c.Username); err != nil {
		return err
	}

	// 设置收件人
	for _, recipient := range to {
		if err = client.Rcpt(recipient); err != nil {
			return err
		}
	}

	// 发送邮件内容
	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(message)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return client.Quit()
}
