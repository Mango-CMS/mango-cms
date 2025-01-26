package service

import (
	"errors"
	"time"

	"github.com/Mango-CMS/mango-cms/internal/auth"
	"github.com/Mango-CMS/mango-cms/internal/model"
	"github.com/Mango-CMS/mango-cms/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{repo: repository.NewUserRepository()}
}

// GetUserByID 获取单个用户
func (s *UserService) GetUserByID(id string) (*model.User, error) {
	return s.repo.GetUserByID(id)
}

// GetUsers 获取用户列表
func (s *UserService) GetUsers() ([]model.User, error) {
	return s.repo.GetUsers()
}

// CreateUser 创建用户
func (s *UserService) CreateUser(user *model.User) error {
	// 对密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.CreateUser(user)
}

// UpdateUser 更新用户
func (s *UserService) UpdateUser(user *model.User) error {
	return s.repo.UpdateUser(user)
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}

// Login 用户登录
func (s *UserService) Login(username, password string) (string, error) {
	// 获取用户信息
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("用户不存在")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("密码错误")
	}

	// 检查用户状态
	if user.Status != "active" {
		return "", errors.New("用户账号已被禁用")
	}

	// 更新最后登录时间
	now := time.Now()
	user.LastLogin = &now
	if err := s.repo.UpdateUser(user); err != nil {
		return "", err
	}

	// 生成JWT令牌
	token, err := auth.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserByUsername 获取单个用户
func (s *UserService) GetUserByUsername(username string) (*model.User, error) {
	return s.repo.GetUserByUsername(username)
}
