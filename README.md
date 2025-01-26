# Mango CMS

一个基于Go和Vue 3的现代化内容管理系统，采用前后端分离架构。

## 技术栈

### 后端
- Go
- GraphQL
- MongoDB
- JWT认证

### 前端
- Vue 3
- TypeScript
- Tailwind CSS
- Naive UI
- Apollo Client
- Vue Router
- Pinia

## 项目结构

```
.
├── admin/            # 前端管理系统
│   ├── src/         # 源代码
│   │   ├── assets/  # 静态资源
│   │   ├── components/ # 公共组件
│   │   ├── graphql/ # GraphQL查询和变更
│   │   ├── layouts/ # 布局组件
│   │   ├── router/  # 路由配置
│   │   ├── stores/  # Pinia状态管理
│   │   └── views/   # 页面组件
│   └── vite.config.ts # Vite配置
├── cmd/             # 后端入口程序
├── internal/        # 后端内部实现
│   ├── auth/       # 认证相关
│   ├── config/     # 配置管理
│   ├── middleware/ # 中间件
│   ├── model/      # 数据模型
│   ├── repository/ # 数据访问层
│   └── service/    # 业务逻辑层
└── schema/         # GraphQL schema定义
```

## 快速开始

### 后端

1. 克隆仓库
```bash
git clone https://github.com/yourusername/mango-cms.git
cd mango-cms
```

2. 配置环境变量
```bash
cp .env.example .env
# 修改.env中的配置信息
```

3. 安装依赖并运行
```bash
go mod download
go run cmd/main.go
```

### 前端

1. 进入前端目录
```bash
cd admin
```

2. 安装依赖
```bash
pnpm install
```

3. 开发环境运行
```bash
pnpm dev
```

4. 生产环境构建
```bash
pnpm build
```

## 功能特性

- 用户认证和授权
- 应用管理
  - 创建和管理自定义应用
  - 自定义字段配置
  - 应用权限管理
- 用户管理
- 系统设置

## GraphQL API

启动后端服务后，访问 `/playground` 路径可以打开GraphQL Playground界面，在这里你可以：
- 查看完整的Schema文档
- 测试API请求
- 实时查看查询结果

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情