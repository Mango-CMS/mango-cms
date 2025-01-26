package main

import (
	"log"

	"github.com/Mango-CMS/mango-cms/internal/config"
	"github.com/Mango-CMS/mango-cms/internal/repository"
	"github.com/Mango-CMS/mango-cms/internal/service"
	"github.com/Mango-CMS/mango-cms/schema"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func main() {
	// 加载环境变量配置
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	// 初始化Gin引擎
	r := gin.Default()

	// 初始化数据库连接
	if err := repository.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 初始化服务 包含了数据库初始化
	userService := service.NewUserService()
	apiTokenService := service.NewAPITokenService()
	permissionService := service.NewPermissionService()
	applicationService := service.NewApplicationService()

	// 初始化GraphQL schema
	schema.InitServices(
		userService,
		apiTokenService,
		permissionService,
		applicationService,
	)

	// 设置GraphQL处理器
	h := handler.New(&handler.Config{
		Schema: &schema.Schema,
		Pretty: true,
	})

	// 添加GraphQL路由
	r.POST("/graphql", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	// 添加GraphQL Playground路由
	r.GET("/playground", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.File("internal/templates/playground.html")
	})

	// 启动服务器
	if err := r.Run(":" + config.GetEnv("SERVER_PORT", "8080")); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
