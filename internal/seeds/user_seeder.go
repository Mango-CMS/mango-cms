package seeds

import (
	"context"
	"log"
	"time"

	"github.com/Mango-CMS/mango-cms/internal/model"
	"github.com/Mango-CMS/mango-cms/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserSeeder 用户数据填充器
type UserSeeder struct {
	repo *repository.UserRepository
}

// NewUserSeeder 创建用户数据填充器
func NewUserSeeder() *UserSeeder {
	return &UserSeeder{
		repo: repository.NewUserRepository(),
	}
}

// Seed 填充用户数据
func (s *UserSeeder) Seed(ctx context.Context) error {
	// 创建默认管理员用户
	adminPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin := &model.User{
		Username:  "admin",
		Email:     "admin@example.com",
		Password:  string(adminPassword),
		Role:      "admin",
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.CreateUser(admin); err != nil {
		log.Printf("Failed to create admin user: %v", err)
		return err
	}

	// 创建测试编辑用户
	editorPassword, _ := bcrypt.GenerateFromPassword([]byte("editor123"), bcrypt.DefaultCost)
	editor := &model.User{
		Username:  "editor",
		Email:     "editor@example.com",
		Password:  string(editorPassword),
		Role:      "editor",
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.CreateUser(editor); err != nil {
		log.Printf("Failed to create editor user: %v", err)
		return err
	}

	return nil
}

// Clear 清除用户数据
func (s *UserSeeder) Clear(ctx context.Context) error {
	// 删除所有测试用户
	users, err := s.repo.GetUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		if err := s.repo.DeleteUser(user.ID.String()); err != nil {
			log.Printf("Failed to delete user %s: %v", user.ID, err)
			return err
		}
	}

	return nil
}
