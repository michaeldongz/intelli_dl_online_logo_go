package database

import (
	"context"
	"encoding/json"
	"time"
)

// RedisUtils 提供Redis常用操作的工具类
type RedisUtils struct{}

// Set 设置键值对
func (r *RedisUtils) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	var val string
	switch v := value.(type) {
	case string:
		val = v
	default:
		bytes, err := json.Marshal(value)
		if err != nil {
			return err
		}
		val = string(bytes)
	}
	return RedisClient.Set(ctx, key, val, expiration).Err()
}

// Get 获取值
func (r *RedisUtils) Get(ctx context.Context, key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

// GetObj 获取并解析为对象
func (r *RedisUtils) GetObj(ctx context.Context, key string, obj interface{}) error {
	val, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), obj)
}

// Delete 删除键
func (r *RedisUtils) Delete(ctx context.Context, keys ...string) error {
	return RedisClient.Del(ctx, keys...).Err()
}

// Exists 检查键是否存在
func (r *RedisUtils) Exists(ctx context.Context, key string) (bool, error) {
	result, err := RedisClient.Exists(ctx, key).Result()
	return result > 0, err
}

// Expire 设置过期时间
func (r *RedisUtils) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return RedisClient.Expire(ctx, key, expiration).Err()
}

// NewRedisUtils 创建Redis工具类实例
func NewRedisUtils() *RedisUtils {
	return &RedisUtils{}
}
