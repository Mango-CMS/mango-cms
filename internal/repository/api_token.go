package repository

import (
	"context"
	"time"

	"github.com/Mango-CMS/mango-cms/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const APITokensCollection = "api_tokens"

type APITokenRepository struct{}

func NewAPITokenRepository() *APITokenRepository {
	return &APITokenRepository{}
}

// GetAPITokenByID 根据ID获取API令牌
func (r *APITokenRepository) GetAPITokenByID(id primitive.ObjectID) (*model.APIToken, error) {
	var token model.APIToken
	err := DB.Collection(APITokensCollection).FindOne(
		context.Background(),
		bson.M{"_id": id},
	).Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// GetAPITokenByToken 根据令牌值获取API令牌
func (r *APITokenRepository) GetAPITokenByToken(token string) (*model.APIToken, error) {
	var apiToken model.APIToken
	err := DB.Collection(APITokensCollection).FindOne(
		context.Background(),
		bson.M{"token": token},
	).Decode(&apiToken)
	if err != nil {
		return nil, err
	}
	return &apiToken, nil
}

// GetAPITokensByUserID 获取用户的所有API令牌
func (r *APITokenRepository) GetAPITokensByUserID(userID primitive.ObjectID) ([]model.APIToken, error) {
	var tokens []model.APIToken
	cursor, err := DB.Collection(APITokensCollection).Find(
		context.Background(),
		bson.M{"user_id": userID},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &tokens); err != nil {
		return nil, err
	}
	return tokens, nil
}

// CreateAPIToken 创建新的API令牌
func (r *APITokenRepository) CreateAPIToken(token *model.APIToken) error {
	token.CreatedAt = time.Now()
	token.UpdatedAt = time.Now()
	_, err := DB.Collection(APITokensCollection).InsertOne(context.Background(), token)
	return err
}

// UpdateAPIToken 更新API令牌信息
func (r *APITokenRepository) UpdateAPIToken(token *model.APIToken) error {
	token.UpdatedAt = time.Now()
	_, err := DB.Collection(APITokensCollection).UpdateOne(
		context.Background(),
		bson.M{"_id": token.ID},
		bson.M{"$set": token},
	)
	return err
}

// DeleteAPIToken 删除API令牌
func (r *APITokenRepository) DeleteAPIToken(id primitive.ObjectID) error {
	_, err := DB.Collection(APITokensCollection).DeleteOne(
		context.Background(),
		bson.M{"_id": id},
	)
	return err
}
