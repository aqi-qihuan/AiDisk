# AqiCloud-Agent

<div align="center">

**AI 智能化云盘系统后端 API**

[![Go Version](https://img.shields.io/badge/Go-1.26.2-blue.svg)](https://golang.org/)
[![Gin](https://img.shields.io/badge/Gin-v1.12.0-00A6ED.svg)](https://gin-gonic.com/)
[![License](https://img.shields.io/badge/license-Apache%202.0-orange.svg)](LICENSE)
[![Swagger](https://img.shields.io/badge/Swagger-Enabled-green.svg)](http://localhost:8080/swagger/index.html)

</div>

## 📖 项目简介

AqiCloud-Agent 是一个基于 Go 语言开发的 AI 智能化云盘系统后端服务，提供文件管理、用户认证、文件分享、回收站、AI 聊天等完整功能。系统采用 RESTful API 设计，支持分块上传、秒传、断点续传等高级特性，并集成了多种 AI 服务（DashScope、Ollama）。

### ✨ 核心特性

- 🔐 **安全的身份认证** - 基于 JWT 的无状态认证机制
- 📁 **完整的文件管理** - 上传、下载、移动、复制、删除、重命名
- 🚀 **高性能分块上传** - 支持大文件分块上传、秒传、断点续传
- 🔗 **文件分享系统** - 支持有效期和提取码的分享链接
- 🗑️ **回收站功能** - 软删除、恢复、永久删除
- 🤖 **AI 智能聊天** - 集成 DashScope 和 Ollama，支持 SSE 流式响应
- 📊 **存储空间管理** - 用户配额管理、存储统计
- 📚 **完整 API 文档** - 基于 Swagger 的自动化 API 文档
- 🐳 **容器化部署** - 提供 Dockerfile，支持容器化部署

## 🛠️ 技术栈

### 后端框架
- **语言**: Go 1.26.2
- **Web 框架**: Gin v1.12.0
- **ORM**: GORM v1.31.1
- **认证**: JWT (golang-jwt/jwt/v5)

### 数据存储
- **数据库**: MySQL 8.0+
- **对象存储**: MinIO / AWS S3 兼容存储

### AI 服务
- **DashScope API** (阿里云百炼)
- **Ollama** (本地 LLM 部署)

### 文档与工具
- **API 文档**: Swagger (swaggo/swag)
- **配置管理**: godotenv
- **唯一 ID**: Snowflake 算法

## 📂 项目结构

```
AqiCloud-Agent/
├── cmd/
│   └── server/
│       └── main.go          # 程序入口
├── internal/
│   ├── config/              # 配置管理
│   │   ├── config.go        # 配置结构体
│   │   └── db.go           # 数据库初始化
│   ├── controller/         # 控制器层
│   │   ├── account.go      # 账号相关接口
│   │   ├── chat.go         # AI 聊天接口
│   │   ├── file.go         # 文件管理接口
│   │   ├── share.go        # 文件分享接口
│   │   └── recycle.go     # 回收站接口
│   ├── service/            # 业务逻辑层
│   │   ├── account_service.go
│   │   ├── chat_service.go
│   │   ├── file_service.go
│   │   ├── share_service.go
│   │   ├── recycle_service.go
│   │   ├── chunk_service.go # 分块上传服务
│   │   └── store_engine.go # 存储引擎
│   ├── model/              # 数据模型层
│   │   ├── account.go
│   │   ├── account_file.go
│   │   ├── file_chunk.go
│   │   ├── share.go
│   │   ├── request.go      # 请求结构体
│   │   └── response.go    # 响应结构体
│   ├── middleware/         # 中间件
│   │   ├── auth.go         # 认证中间件
│   │   └── cors.go         # CORS 中间件
│   └── util/              # 工具函数
├── docs/                   # Swagger 文档
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── .env.example            # 环境变量示例
├── Dockerfile              # Docker 构建文件
├── go.mod                  # Go 模块文件
└── go.sum                  # 依赖校验文件
```

## 🚀 安装步骤

### 环境要求

- **Go**: 1.26.2 或更高版本
- **MySQL**: 8.0 或更高版本
- **MinIO**: 最新版本（或使用 AWS S3）
- **(可选) Ollama**: 用于本地 AI 模型部署

### 方式一：从源码构建

#### 1. 克隆项目

```bash
git clone https://github.com/aqi/AqiCloud-Agent-Go.git
cd AqiCloud-Agent-Go
```

#### 2. 安装依赖

```bash
go mod download
```

#### 3. 配置环境变量

复制 `.env.example` 为 `.env` 并修改配置：

```bash
cp .env.example .env
```

编辑 `.env` 文件，配置以下参数（详见 [配置说明](#配置说明)）：

```bash
# 必填项
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=your_password
MYSQL_DATABASE=aqi-cloud-pan

MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY=your_access_key
MINIO_SECRET_KEY=your_secret_key
MINIO_BUCKET=ai-pan

JWT_SECRET=your_32_characters_secret_key_at_least

# AI 服务（可选）
DASHSCOPE_API_KEY=your_dashscope_api_key
```

#### 4. 运行服务

```bash
go run cmd/server/main.go
```

或使用编译后的二进制文件：

```bash
go build -o agentpan.exe cmd/server/main.go
./agentpan.exe
```

### 方式二：使用 Docker 部署

#### 1. 构建镜像

```bash
docker build -t AqiCloud-Agent:latest .
```

#### 2. 运行容器

```bash
docker run -d \
  -p 8080:8080 \
  --env-file .env \
  --name agentpan \
  AqiCloud-Agent:latest
```

#### 3. 使用 Docker Compose（推荐）

创建 `docker-compose.yml`：

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    env_file:
      - .env
    depends_on:
      - mysql
      - minio
    restart: unless-stopped

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: rootpassword
      MYSQL_DATABASE: aqi-cloud-pan
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql
    restart: unless-stopped

  minio:
    image: minio/minio:latest
    command: server /data --console-address ":9001"
    environment:
      MINIO_ROOT_USER: minioadmin
      MINIO_ROOT_PASSWORD: minioadmin
    ports:
      - "9000:9000"
      - "9001:9001"
    volumes:
      - minio_data:/data
    restart: unless-stopped

volumes:
  mysql_data:
  minio_data:
```

启动服务：

```bash
docker-compose up -d
```

## ⚙️ 配置说明

### 完整配置项

| 配置项 | 说明 | 默认值 | 必填 |
|--------|------|--------|------|
| `APP_NAME` | 应用名称 | AqiCloud-Agent API | ❌ |
| `DEBUG` | 调试模式 | false | ❌ |
| `LISTEN_ADDR` | 监听地址 | :8080 | ❌ |
| `MYSQL_HOST` | MySQL 主机地址 | - | ✅ |
| `MYSQL_PORT` | MySQL 端口 | 3306 | ❌ |
| `MYSQL_USER` | MySQL 用户名 | - | ✅ |
| `MYSQL_PASSWORD` | MySQL 密码 | - | ✅ |
| `MYSQL_DATABASE` | MySQL 数据库名 | - | ✅ |
| `MINIO_ENDPOINT` | MinIO 端点地址 | - | ✅ |
| `MINIO_ACCESS_KEY` | MinIO 访问密钥 | - | ✅ |
| `MINIO_SECRET_KEY` | MinIO 秘密密钥 | - | ✅ |
| `MINIO_BUCKET` | MinIO 存储桶名 | ai-pan | ❌ |
| `MINIO_AVATAR_BUCKET` | 头像存储桶名 | avatar | ❌ |
| `JWT_SECRET` | JWT 密钥（至少 32 字符） | - | ✅ |
| `JWT_LOGIN_SUBJECT` | 登录 JWT 主题 | AQI | ❌ |
| `JWT_SHARE_SUBJECT` | 分享 JWT 主题 | AQI_SHARE | ❌ |
| `DASHSCOPE_API_KEY` | 阿里云百炼 API Key | - | ❌ |
| `DASHSCOPE_BASE` | DashScope API 地址 | https://dashscope.aliyuncs.com/... | ❌ |
| `OLLAMA_BASE_URL` | Ollama API 地址 | - | ❌ |
| `STREAM_BASE_URL` | SSE 流式代理地址 | - | ❌ |
| `STREAM_CHAT_PATH` | SSE 聊天路径 | /api/chat/stream | ❌ |
| `FRONTEND_BASE_URL` | 前端基础 URL | http://127.0.0.1:8080 | ❌ |
| `DEFAULT_STORAGE_SIZE` | 默认存储空间（字节） | 10737418240 (10GB) | ❌ |
| `MAX_UPLOAD_SIZE` | 单次上传最大大小（字节） | 104857600 (100MB) | ❌ |

### MinIO 配置指南

1. **安装 MinIO**

```bash
# Docker 方式
docker run -d -p 9000:9000 -p 9001:9001 \
  -e "MINIO_ROOT_USER=minioadmin" \
  -e "MINIO_ROOT_PASSWORD=minioadmin" \
  -v minio_data:/data \
  minio/minio server /data --console-address ":9001"
```

2. **创建存储桶**

访问 MinIO 控制台：http://localhost:9001
- 使用默认凭证登录（minioadmin/minioadmin）
- 创建 `ai-pan` 和 `avatar` 两个存储桶

3. **配置 .env**

```bash
MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY=minioadmin
MINIO_SECRET_KEY=minioadmin
MINIO_BUCKET=ai-pan
MINIO_AVATAR_BUCKET=avatar
```

## 📚 使用示例

### 1. 注册账号

```bash
curl -X POST http://localhost:8080/account/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "username": "testuser",
    "password": "password123"
  }'
```

**响应示例：**

```json
{
  "code": 200,
  "message": "注册成功",
  "data": {
    "account_id": 1234567890,
    "username": "testuser",
    "email": "user@example.com"
  }
}
```

### 2. 登录获取 Token

```bash
curl -X POST http://localhost:8080/account/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

**响应示例：**

```json
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "account_id": 1234567890,
    "username": "testuser"
  }
}
```

### 3. 上传文件

```bash
# 普通上传
curl -X POST http://localhost:8080/file/upload \
  -H "token: YOUR_JWT_TOKEN" \
  -F "file=@/path/to/your/file.txt" \
  -F "parent_id=0"
```

### 4. 分块上传大文件

**步骤 1：初始化分块上传**

```bash
curl -X POST http://localhost:8080/file/chunk/init \
  -H "token: YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "file_name": "large_file.zip",
    "file_size": 1073741824,
    "parent_id": 0
  }'
```

**响应示例：**

```json
{
  "code": 200,
  "message": "初始化成功",
  "data": {
    "file_id": 1234567890,
    "chunk_size": 5242880,
    "total_chunks": 205
  }
}
```

**步骤 2：上传分块**

```bash
curl -X POST http://localhost:8080/file/chunk/upload \
  -H "token: YOUR_JWT_TOKEN" \
  -F "file_id=1234567890" \
  -F "chunk_index=0" \
  -F "chunk=@/path/to/chunk_0.dat"
```

**步骤 3：完成上传**

```bash
curl -X POST http://localhost:8080/file/chunk/complete \
  -H "token: YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "file_id": 1234567890
  }'
```

### 5. 创建文件分享

```bash
curl -X POST http://localhost:8080/share/create \
  -H "token: YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "file_id": 1234567890,
    "expire_days": 7,
    "has_password": true,
    "password": "1234"
  }'
```

**响应示例：**

```json
{
  "code": 200,
  "message": "分享创建成功",
  "data": {
    "share_id": "abc123def456",
    "share_url": "http://localhost:8080/share/abc123def456"
  }
}
```

### 6. AI 聊天

```bash
curl -X POST http://localhost:8080/chat/stream \
  -H "token: YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "message": "你好，请介绍一下自己"
  }'
```

## 📖 API 文档

项目集成了 Swagger API 文档，启动服务后访问：

```
http://localhost:8080/swagger/index.html
```

### 主要 API 端点

#### 账号相关
- `POST /account/register` - 用户注册
- `POST /account/login` - 用户登录
- `GET /account/info` - 获取用户信息
- `PUT /account/profile` - 更新用户资料

#### 文件管理
- `POST /file/upload` - 上传文件
- `GET /file/list` - 文件列表
- `GET /file/download/:file_id` - 下载文件
- `DELETE /file/delete/:file_id` - 删除文件
- `PUT /file/move` - 移动文件
- `PUT /file/rename` - 重命名文件

#### 分块上传
- `POST /file/chunk/init` - 初始化分块上传
- `POST /file/chunk/upload` - 上传分块
- `POST /file/chunk/complete` - 完成分块上传
- `GET /file/chunk/status/:file_id` - 查询上传状态

#### 文件分享
- `POST /share/create` - 创建分享
- `GET /share/info/:share_id` - 获取分享信息
- `POST /share/verify` - 验证分享密码
- `DELETE /share/delete/:share_id` - 删除分享

#### 回收站
- `GET /recycle/list` - 回收站列表
- `POST /recycle/restore/:file_id` - 恢复文件
- `DELETE /recycle/delete/:file_id` - 永久删除

#### AI 聊天
- `POST /chat/stream` - AI 流式聊天（SSE）
- `GET /chat/history` - 聊天历史记录

## 🧪 运行测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test github.com/aqi/AqiCloud-Agent-Go/internal/service

# 查看测试覆盖率
go test -cover ./...
```

## 🤝 贡献指南

我们欢迎任何形式的贡献，包括但不限于：

- 🐛 报告 Bug
- 💡 提出新功能建议
- 📖 改进文档
- 🔧 提交代码修复
- ✨ 实现新功能

### 贡献流程

#### 1. Fork 项目

点击 GitHub 页面右上角的 "Fork" 按钮，将项目 fork 到你的账户下。

#### 2. 克隆你的 Fork

```bash
git clone https://github.com/YOUR_USERNAME/AqiCloud-Agent-Go.git
cd AqiCloud-Agent-Go
```

#### 3. 创建功能分支

```bash
git checkout -b feature/your-feature-name
```

分支命名规范：
- `feature/` - 新功能
- `fix/` - Bug 修复
- `docs/` - 文档更新
- `refactor/` - 代码重构
- `test/` - 测试相关

#### 4. 开发和提交

```bash
# 进行你的修改
git add .
git commit -m "feat: 添加新功能描述"
```

**Commit Message 规范**（遵循 [Conventional Commits](https://www.conventionalcommits.org/)）：

- `feat:` - 新功能
- `fix:` - Bug 修复
- `docs:` - 文档更新
- `style:` - 代码格式调整（不影响功能）
- `refactor:` - 代码重构
- `test:` - 添加或修改测试
- `chore:` - 构建过程或辅助工具的变动

**示例：**

```bash
git commit -m "feat: 添加文件预览功能"
git commit -m "fix: 修复分块上传失败的问题"
git commit -m "docs: 更新 README 安装说明"
```

#### 5. 推送到你的 Fork

```bash
git push origin feature/your-feature-name
```

#### 6. 创建 Pull Request

1. 访问你的 Fork 页面
2. 点击 "New Pull Request"
3. 填写 PR 描述，说明你的修改内容
4. 提交 PR 等待审核

### 代码规范

- 遵循 Go 官方代码风格（使用 `gofmt` 格式化代码）
- 添加必要的注释，尤其是导出函数和复杂逻辑
- 确保每个函数都有适当的错误处理
- 编写单元测试，测试覆盖率不低于 60%
- 运行 `go lint` 检查代码质量

### 报告 Bug

提交 Issue 时，请包含以下信息：

- 🐛 **Bug 描述**：简洁明确地描述问题
- 🔄 **复现步骤**：详细描述如何复现该 Bug
- ✅ **预期行为**：你认为的正确行为
- ❌ **实际行为**：实际发生的情况
- 🖥️ **环境信息**：操作系统、Go 版本、MySQL 版本等
- 📝 **相关日志**：错误日志或截图

### 功能建议

提交功能建议时，请说明：

- 💡 **功能描述**：详细描述你希望添加的功能
- 🎯 **使用场景**：为什么需要这个功能
- 📐 **实现思路**：如果有，提供你的实现思路
- 📷 **参考资料**：提供相关的截图或文档链接

## 📄 许可证

本项目采用 Apache License 2.0 开源许可证。详见 [LICENSE](LICENSE) 文件。

## 📧 联系方式

- **作者**: 阿七
- **邮箱**: 2316364297@qq.com
- **网站**: https://pan.aqi125.cn
- **GitHub**: https://github.com/aqi-qihuan/AiDisk

## 🙏 致谢

感谢以下开源项目的支持：

- [Gin](https://gin-gonic.com/) - Web 框架
- [GORM](https://gorm.io/) - ORM 库
- [MinIO](https://min.io/) - 对象存储
- [Swag](https://github.com/swaggo/swag) - Swagger 文档生成

---

<div align="center">

**⭐ 如果这个项目对你有帮助，请给它一个 Star！⭐**

</div>
