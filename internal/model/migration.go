package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Migration 记录数据库迁移历史
type Migration struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Batch     int                `json:"batch" bson:"batch"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time         `json:"deleted_at" bson:"deleted_at,omitempty"`
}
