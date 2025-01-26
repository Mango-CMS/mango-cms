package main

import (
	"context"
	"flag"
	"log"

	"github.com/Mango-CMS/mango-cms/internal/config"
	"github.com/Mango-CMS/mango-cms/internal/repository"
	"github.com/Mango-CMS/mango-cms/internal/seeds"
)

func main() {
	// 定义命令行参数
	clear := flag.Bool("clear", false, "清除所有填充数据")
	flag.Parse()

	// 加载环境变量配置
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	// 初始化数据库连接
	if err := repository.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 创建数据填充管理器
	manager := seeds.NewManager()

	// 注册数据填充器
	manager.Register(seeds.NewUserSeeder())

	ctx := context.Background()

	// 根据命令行参数执行操作
	if *clear {
		// 清除所有填充数据
		if err := manager.ClearAll(ctx); err != nil {
			log.Fatalf("Failed to clear seed data: %v", err)
		}
		log.Println("Successfully cleared all seed data")
	} else {
		// 执行数据填充
		if err := manager.SeedAll(ctx); err != nil {
			log.Fatalf("Failed to seed data: %v", err)
		}
		log.Println("Successfully seeded all data")
	}
}
