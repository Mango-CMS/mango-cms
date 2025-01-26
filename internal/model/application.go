package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Application 应用模块定义
type Application struct {
	ID          primitive.ObjectID `json:"id,string" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`               // 模块名称
	Slug        string             `json:"slug" bson:"slug"`               // 模块标识
	Sign        string             `json:"sign" bson:"sign"`               // 模块签名
	Description string             `json:"description" bson:"description"` // 模块描述
	Fields      []ApplicationField `json:"fields" bson:"fields"`           // 字段定义
	Models      []ApplicationModel `json:"models" bson:"models"`           // 模型定义
	Status      string             `json:"status" bson:"status"`           // 模块状态：active, inactive
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt   *time.Time         `json:"deleted_at" bson:"deleted_at,omitempty"`
}

// ApplicationField 应用模块字段定义
type ApplicationField struct {
	ID          primitive.ObjectID `json:"id,string" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`               // 字段名称
	Slug        string             `json:"slug" bson:"slug"`               // 字段标识
	Type        string             `json:"type" bson:"type"`               // 字段类型：string, number, boolean, date, etc
	Required    bool               `json:"required" bson:"required"`       // 是否必填
	Description string             `json:"description" bson:"description"` // 字段描述
	Default     interface{}        `json:"default" bson:"default"`         // 默认值
	Validation  interface{}        `json:"validation" bson:"validation"`   // 验证规则
}

type ApplicationModel struct {
	ID          primitive.ObjectID `json:"id,string" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`               // 模型名称
	Slug        string             `json:"slug" bson:"slug"`               // 模型标识
	Description string             `json:"description" bson:"description"` // 模型描述
	Fields      []ApplicationField `json:"fields" bson:"fields"`           // 字段定义
	Content     interface{}        `json:"content" bson:"content"`         // 模型内容,实际上模型的内容应该与Fields保持一致
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

// ApplicationPermission 应用模块权限定义
type ApplicationPermission struct {
	ID            primitive.ObjectID `json:"id,string" bson:"_id,omitempty"`
	ApplicationID primitive.ObjectID `json:"application_id,string" bson:"application_id"` // 关联的应用模块ID
	RoleID        string             `json:"role_id" bson:"role_id"`                      // 角色ID
	Permissions   []string           `json:"permissions" bson:"permissions"`              // 权限列表：create, read, update, delete
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
}
