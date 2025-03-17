package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User 用户模型
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nickname  string             `bson:"nickname" json:"nickname"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"` // 不在JSON中返回密码
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
