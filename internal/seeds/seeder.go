package seeds

import (
	"context"
	"log"
)

// Seeder 定义数据填充器接口
type Seeder interface {
	Seed(ctx context.Context) error
	Clear(ctx context.Context) error
}

// Manager 管理所有数据填充器
type Manager struct {
	seeders []Seeder
}

// NewManager 创建一个新的数据填充管理器
func NewManager() *Manager {
	return &Manager{}
}

// Register 注册一个数据填充器
func (m *Manager) Register(seeder Seeder) {
	m.seeders = append(m.seeders, seeder)
}

// SeedAll 执行所有数据填充
func (m *Manager) SeedAll(ctx context.Context) error {
	for _, seeder := range m.seeders {
		if err := seeder.Seed(ctx); err != nil {
			log.Printf("Failed to seed data: %v", err)
			return err
		}
	}
	return nil
}

// ClearAll 清除所有填充的数据
func (m *Manager) ClearAll(ctx context.Context) error {
	// 反向清除数据，以处理数据依赖关系
	for i := len(m.seeders) - 1; i >= 0; i-- {
		if err := m.seeders[i].Clear(ctx); err != nil {
			log.Printf("Failed to clear data: %v", err)
			return err
		}
	}
	return nil
}
