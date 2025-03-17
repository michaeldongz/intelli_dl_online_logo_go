package utils

import (
	"context"
	"errors"
	"fmt"
	"myapp/internal/database"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	jwtSecret  = []byte("your_jwt_secret_key") // 实际应用中应该放在配置文件中
	redisUtils = database.NewRedisUtils()
)

// JWTClaims 自定义JWT声明
type JWTClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT令牌
func GenerateToken(userID, email string) (string, error) {
	Debug("为用户生成令牌: %s, ID: %s", email, userID)

	// 设置过期时间为24小时
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &JWTClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		Error("生成令牌失败: %v", err)
		return "", err
	}

	// 将令牌存储到Redis中
	ctx := context.Background()
	redisKey := fmt.Sprintf("token:%s", userID)
	err = redisUtils.Set(ctx, redisKey, tokenString, 24*time.Hour)
	if err != nil {
		Error("将令牌存储到Redis失败: %v", err)
		return "", err
	}

	Debug("令牌生成成功: %s", userID)
	return tokenString, nil
}

// ParseToken 解析JWT令牌
func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		Debug("解析令牌失败: %v", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		// 验证令牌是否在Redis中
		ctx := context.Background()
		redisKey := fmt.Sprintf("token:%s", claims.UserID)
		exists, err := redisUtils.Exists(ctx, redisKey)
		if err != nil || !exists {
			Debug("令牌在Redis中不存在: %s", claims.UserID)
			return nil, errors.New("令牌已失效")
		}

		return claims, nil
	}

	Debug("无效的令牌")
	return nil, errors.New("无效的令牌")
}

// InvalidateToken 使令牌失效
func InvalidateToken(userID string) error {
	Info("使令牌失效: %s", userID)
	ctx := context.Background()
	redisKey := fmt.Sprintf("token:%s", userID)
	return redisUtils.Delete(ctx, redisKey)
}
