package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// APIToken 定义API令牌模型
type APIToken struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID      primitive.ObjectID `json:"user_id" bson:"user_id"`                     // 关联的用户ID
	User        User               `json:"user" bson:"user"`                           // 关联的用户
	Name        string             `json:"name" bson:"name"`                           // 令牌名称
	Token       string             `json:"token" bson:"token"`                         // 令牌值
	Permissions string             `json:"permissions" bson:"permissions"`             // 权限范围
	LastUsedAt  *time.Time         `json:"last_used_at" bson:"last_used_at,omitempty"` // 最后使用时间
	ExpiresAt   *time.Time         `json:"expires_at" bson:"expires_at,omitempty"`     // 过期时间
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt   *time.Time         `json:"deleted_at" bson:"deleted_at,omitempty"`
}
