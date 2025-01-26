package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/Mango-CMS/mango-cms/internal/model"
	"github.com/Mango-CMS/mango-cms/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type APITokenService struct {
	repo *repository.APITokenRepository
}

func NewAPITokenService() *APITokenService {
	return &APITokenService{repo: repository.NewAPITokenRepository()}
}

// GenerateToken 生成新的API令牌
func (s *APITokenService) GenerateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// CreateAPIToken 创建新的API令牌
func (s *APITokenService) CreateAPIToken(userID primitive.ObjectID, name string, permissions string, expiresIn time.Duration) (*model.APIToken, error) {
	// 生成令牌
	token, err := s.GenerateToken()
	if err != nil {
		return nil, err
	}

	// 设置过期时间
	expiresAt := time.Now().Add(expiresIn)

	// 创建令牌记录
	apiToken := &model.APIToken{
		UserID:      userID,
		Name:        name,
		Token:       token,
		Permissions: permissions,
		ExpiresAt:   &expiresAt,
	}

	if err := s.repo.CreateAPIToken(apiToken); err != nil {
		return nil, err
	}

	return apiToken, nil
}

// GetAPITokenByID 根据ID获取API令牌
func (s *APITokenService) GetAPITokenByID(id primitive.ObjectID) (*model.APIToken, error) {
	return s.repo.GetAPITokenByID(id)
}

// GetAPITokensByUserID 获取用户的所有API令牌
func (s *APITokenService) GetAPITokensByUserID(userID primitive.ObjectID) ([]model.APIToken, error) {
	return s.repo.GetAPITokensByUserID(userID)
}

// ValidateToken 验证API令牌
func (s *APITokenService) ValidateToken(tokenStr string) (*model.APIToken, error) {
	// 获取令牌信息
	token, err := s.repo.GetAPITokenByToken(tokenStr)
	if err != nil {
		return nil, errors.New("无效的令牌")
	}

	// 检查令牌是否过期
	if token.ExpiresAt != nil && token.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("令牌已过期")
	}

	// 更新最后使用时间
	now := time.Now()
	token.LastUsedAt = &now
	if err := s.repo.UpdateAPIToken(token); err != nil {
		return nil, err
	}

	return token, nil
}

// RevokeToken 撤销API令牌
func (s *APITokenService) RevokeToken(id primitive.ObjectID) error {
	return s.repo.DeleteAPIToken(id)
}

// UpdateAPIToken 更新API令牌信息
func (s *APITokenService) UpdateAPIToken(token *model.APIToken) error {
	return s.repo.UpdateAPIToken(token)
}
