package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

// Config 日志配置
type Config struct {
	Level      string
	Format     string
	Output     string
	Directory  string
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

// InitLogger 初始化日志
func InitLogger(config Config) error {
	// 设置日志级别
	var level zapcore.Level
	switch config.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	// 设置编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var encoder zapcore.Encoder
	if config.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 设置日志输出
	var writeSyncer zapcore.WriteSyncer
	if config.Output == "file" {
		// 确保日志目录存在
		if err := os.MkdirAll(config.Directory, 0755); err != nil {
			return fmt.Errorf("创建日志目录失败: %w", err)
		}

		// 设置日志轮转
		logPath := filepath.Join(config.Directory, fmt.Sprintf("%s.log", config.Filename))
		lumberJackLogger := &lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    config.MaxSize, // MB
			MaxBackups: config.MaxBackups,
			MaxAge:     config.MaxAge, // 天
			Compress:   config.Compress,
		}
		writeSyncer = zapcore.AddSync(lumberJackLogger)
	} else {
		writeSyncer = zapcore.AddSync(os.Stdout)
	}

	// 创建核心
	core := zapcore.NewCore(encoder, writeSyncer, level)

	// 创建Logger
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Sugar = Logger.Sugar()

	// 替换全局Logger
	zap.ReplaceGlobals(Logger)

	return nil
}

// Debug 调试日志
func Debug(msg string, args ...interface{}) {
	Sugar.Debugf(msg, args...)
}

// Info 信息日志
func Info(msg string, args ...interface{}) {
	Sugar.Infof(msg, args...)
}

// Warn 警告日志
func Warn(msg string, args ...interface{}) {
	Sugar.Warnf(msg, args...)
}

// Error 错误日志
func Error(msg string, args ...interface{}) {
	Sugar.Errorf(msg, args...)
}

// Fatal 致命错误日志
func Fatal(msg string, args ...interface{}) {
	Sugar.Fatalf(msg, args...)
}

// Sync 同步日志缓冲区到输出
func Sync() {
	_ = Logger.Sync()
}
