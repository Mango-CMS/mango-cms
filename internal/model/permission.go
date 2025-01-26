package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Permission struct {
	ID          primitive.ObjectID `json:"id,string" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Slug        string             `json:"slug" bson:"slug"`
	Description string             `json:"description" bson:"description"`
	Module      string             `json:"module" bson:"module"` // 模块：article, user, etc.
	Action      string             `json:"action" bson:"action"` // 操作：create, read, update, delete
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt   *time.Time         `json:"deleted_at" bson:"deleted_at,omitempty"`
}

type RolePermission struct {
	ID           primitive.ObjectID `json:"id,string" bson:"_id,omitempty"`
	Role         string             `json:"role" bson:"role"`
	PermissionID primitive.ObjectID `json:"permission_id,string" bson:"permission_id"`
	Permission   Permission         `json:"permission" bson:"permission"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt    *time.Time         `json:"deleted_at" bson:"deleted_at,omitempty"`
}
