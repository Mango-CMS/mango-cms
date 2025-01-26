package config

// MongoDBConfig MongoDB配置结构
type MongoDBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	AuthDB   string
	Server   bool
}

// GetMongoDBConfig 获取MongoDB配置
func GetMongoDBConfig() *MongoDBConfig {

	return &MongoDBConfig{
		Host:     GetEnv("MONGODB_HOST", "localhost"),
		Port:     GetEnvAsInt("MONGODB_PORT", 27017),
		User:     GetEnv("MONGODB_USER", ""),
		Password: GetEnv("MONGODB_PASSWORD", ""),
		DBName:   GetEnv("MONGODB_NAME", "mango_cms"),
		AuthDB:   GetEnv("MONGODB_AUTH_DB", "admin"),
		Server:   GetEnvAsBool("MONGODB_SERVER", false),
	}
}
