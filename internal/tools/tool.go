package tools

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	"github.com/Mango-CMS/mango-cms/internal/config"
)

// 生成随机字符串
func GenerateRandomString(length int) string {
	// 生成随机字节
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	// 使用base64编码转换为字符串，并截取指定长度
	str := base64.URLEncoding.EncodeToString(bytes)
	if len(str) > length {
		return str[:length]
	}
	return str
}

// 根据slug生成签名
func GenerateSign(slug string) string {
	salt := config.GetEnv("APPLICATION_SALT", "y+adETcLId2XLi4$=&+GK]7/)T-*|[~V")
	randomStr := GenerateRandomString(16)
	// 使用SHA-256算法生成哈希值
	hash := sha256.New()
	hash.Write([]byte(salt + slug + randomStr))
	// 将哈希值转换为十六进制字符串
	return hex.EncodeToString(hash.Sum(nil))
}
