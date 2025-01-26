# Mango CMS

基于Gin和GraphQL的现代化内容管理系统。

## 项目结构

```
.
├── cmd                 # 应用程序入口
├── internal           # 私有应用程序和库代码
│   ├── config        # 配置
│   ├── handler       # HTTP处理器
│   ├── middleware    # 中间件
│   ├── model        # 数据模型
│   ├── repository   # 数据访问层
│   └── service      # 业务逻辑层
├── pkg              # 可以被外部应用程序使用的库代码
└── schema          # GraphQL schema定义

```

## 特性

- 基于Gin的高性能Web框架
- GraphQL API支持
- 清晰的分层架构
- 完整的错误处理
- 中间件支持
- 配置管理

## 快速开始

1. 克隆仓库
```bash
git clone https://github.com/Mango-CMS/mango-cms.git
```

2. 安装依赖
```bash
go mod download
```

3. 运行应用
```bash
go run cmd/main.go
```

4. 配置环境
```bash
# 复制环境配置文件
cp .env.example .env

# 修改.env文件中的数据库配置
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=mango_cms
```

5. 数据库迁移
```bash
# 执行数据库迁移
go run cmd/migrate/main.go

# 回滚最近一次迁移
go run cmd/migrate/main.go --rollback

# 回滚指定次数的迁移
go run cmd/migrate/main.go --rollback --steps=2
```

## API文档

### GraphQL API

访问 `/playground` 路径可以打开交互式GraphQL Playground界面，在这里你可以：
- 查看完整的Schema文档
- 测试API请求
- 实时查看查询结果

### 常用查询示例

1. 获取所有文章
```graphql
query {
  articles {
    id
    title
    content
    slug
    status
    createdAt
    updatedAt
  }
}
```

2. 获取单篇文章
```graphql
query {
  article(id: 1) {
    id
    title
    content
    slug
    status
  }
}
```

3. 创建文章
```graphql
mutation {
  createArticle(
    title: "文章标题"
    content: "文章内容"
    slug: "article-slug"
    status: "published"
  ) {
    id
    title
    slug
  }
}
```

4. 更新文章
```graphql
mutation {
  updateArticle(
    id: 1
    title: "更新的标题"
    content: "更新的内容"
  ) {
    id
    title
    content
    updatedAt
  }
}
```

5. 删除文章
```graphql
mutation {
  deleteArticle(id: 1)
}
```

## 许可证

本项目采用 Apache 2.0 许可证