package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 用户角色常量
const (
	ROLE_ADMIN       = 1  // 管理员
	ROLE_SUPER_ADMIN = 2  // 超级管理员
	ROLE_USER        = 10 // 普通用户
)

// User 用户模型
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nickname  string             `bson:"nickname" json:"nickname"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"` // 不在JSON中返回密码
	Role      int                `bson:"role" json:"role"`  // 用户角色
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
