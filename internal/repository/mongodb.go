package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/Mango-CMS/mango-cms/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	DB     *mongo.Database

	// 集合名称常量
	UsersCollection       = "users"
	PermissionsCollection = "permissions"
	ApiTokensCollection   = "api_tokens"
	MigrationsCollection  = "migrations"
)

// InitDB 初始化MongoDB连接
func InitDB() error {
	ctx := context.Background()
	dbConfig := config.GetMongoDBConfig()

	// 构建MongoDB连接URI
	uri := ""
	if dbConfig.Server {
		// mongodb+srv://jiabinlan:<db_password>@mangocms-dev.pmqae.mongodb.net/?retryWrites=true&w=majority&appName=MangoCMS-DEV
		uri = fmt.Sprintf("mongodb+srv://%s:%s@%s", dbConfig.User, dbConfig.Password, dbConfig.Host)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%d", dbConfig.Host, dbConfig.Port)
		if dbConfig.User != "" && dbConfig.Password != "" {
			uri = fmt.Sprintf("mongodb://%s:%s@%s:%d", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port)
		}
	}

	// 设置客户端选项
	clientOptions := options.Client().ApplyURI(uri)
	if dbConfig.AuthDB != "" {
		clientOptions.SetAuth(options.Credential{
			AuthSource: dbConfig.AuthDB,
			Username:   dbConfig.User,
			Password:   dbConfig.Password,
		})
	}

	// 连接到MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v", err)
		return err
	}

	// 验证连接
	if err = client.Ping(ctx, nil); err != nil {
		log.Printf("Failed to ping MongoDB: %v", err)
		return err
	}

	Client = client
	DB = client.Database(dbConfig.DBName)

	// 初始化集合和索引
	if err := initCollections(ctx); err != nil {
		log.Printf("Failed to initialize collections: %v", err)
		return err
	}

	log.Printf("Successfully connected to MongoDB: %s", dbConfig.DBName)
	return nil
}

// initCollections 初始化集合和索引
func initCollections(ctx context.Context) error {
	// 用户集合索引
	_, err := DB.Collection(UsersCollection).Indexes().CreateMany(ctx, []mongo.IndexModel{
		{
			Keys:    map[string]interface{}{"username": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    map[string]interface{}{"email": 1},
			Options: options.Index().SetUnique(true),
		},
	})
	if err != nil {
		return fmt.Errorf("failed to create user indexes: %v", err)
	}

	// 权限集合索引
	_, err = DB.Collection(PermissionsCollection).Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    map[string]interface{}{"slug": 1},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return fmt.Errorf("failed to create permission indexes: %v", err)
	}

	// API令牌集合索引
	_, err = DB.Collection(ApiTokensCollection).Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    map[string]interface{}{"token": 1},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return fmt.Errorf("failed to create api token indexes: %v", err)
	}

	return nil
}
