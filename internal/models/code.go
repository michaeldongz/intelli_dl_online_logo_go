package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 验证码类型常量
const (
	CODE_TYPE_EMAIL = 1 // 邮箱验证码
	CODE_TYPE_SMS   = 2 // 短信验证码
)

// 验证码状态常量
const (
	CODE_STATUS_UNUSED  = 1 // 未使用
	CODE_STATUS_USED    = 2 // 已使用
	CODE_STATUS_EXPIRED = 3 // 已过期
)

// Code 验证码模型
type Code struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type      int                `bson:"type" json:"type"`       // 验证码类型：1-邮箱，2-短信
	Code      string             `bson:"code" json:"code"`       // 验证码内容
	Content   string             `bson:"content" json:"content"` // 验证码文案
	Email     string             `bson:"email" json:"email"`     // 邮箱地址
	Phone     string             `bson:"phone" json:"phone"`     // 手机号
	Status    int                `bson:"status" json:"status"`   // 状态：1-未使用，2-已使用，3-已过期
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	UsedAt    *time.Time         `bson:"used_at" json:"used_at"`       // 使用时间
	ValidFrom time.Time          `bson:"valid_from" json:"valid_from"` // 生效时间
	ExpiredAt time.Time          `bson:"expired_at" json:"expired_at"` // 过期时间
}
