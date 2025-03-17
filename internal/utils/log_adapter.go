package utils

import (
	"intelli_dl_onling_logo/pkg/logger"
)

// 以下函数是为了兼容现有代码而提供的适配器

// Debug 调试日志
func Debug(msg string, args ...interface{}) {
	logger.Debug(msg, args...)
}

// Info 信息日志
func Info(msg string, args ...interface{}) {
	logger.Info(msg, args...)
}

// Warn 警告日志
func Warn(msg string, args ...interface{}) {
	logger.Warn(msg, args...)
}

// Error 错误日志
func Error(msg string, args ...interface{}) {
	logger.Error(msg, args...)
}

// Fatal 致命错误日志
func Fatal(msg string, args ...interface{}) {
	logger.Fatal(msg, args...)
}
