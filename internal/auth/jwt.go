package auth

import (
	"errors"
	"time"

	"github.com/Mango-CMS/mango-cms/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

var (
	// 设置签名密钥，实际应用中应该从配置文件或环境变量中读取
	jwtSecret = []byte("mango-cms-secret-key")

	// TokenExpireDuration token过期时间
	TokenExpireDuration = time.Hour * 24

	// ErrInvalidToken 无效的token
	ErrInvalidToken = errors.New("invalid token")

	// ErrExpiredToken token过期
	ErrExpiredToken = errors.New("token expired")

	// ErrUnauthorized 未授权
	ErrUnauthorized = errors.New("unauthorized")
)

// Claims 自定义声明结构体
type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT令牌
func GenerateToken(user *model.User) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID:   user.ID.Hex(),
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(TokenExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析JWT令牌
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}
