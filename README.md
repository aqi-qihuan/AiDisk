# AiDisk - AI 智能云盘系统

<div align="center">

**AI 驱动的个人云盘系统，集成智能聊天、文档分析、网盘智答三大 AI 能力**

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.26-blue.svg)](https://golang.org/)
[![Vue](https://img.shields.io/badge/Vue-3.x-42b883.svg)](https://vuejs.org/)
[![GitHub Stars](https://img.shields.io/github/stars/aqi-qihuan/AiDisk?style=social)](https://github.com/aqi-qihuan/AiDisk)

</div>

---

## 📖 项目简介

AiDisk 是一个基于 Go + Vue 3 开发的 AI 智能云盘系统，提供完整的文件管理、分享、回收站功能，并集成了三大 AI 能力：

- 🤖 **AI 智能聊天** - 支持 Ollama 和 OpenAI 兼容接口，流式响应
- 📄 **AI 文档分析** - 智能分析网页和 PDF 文档，生成摘要
- 🔍 **AI 网盘智答** - 自然语言查询网盘，意图识别自动路由

系统采用微服务模式部署，各服务职责清晰，支持 Docker Compose 一键部署。

---

## ✨ 核心特性

### 📁 文件管理（AqiCloud-Agent）

- ✅ 用户注册/登录（JWT 认证）
- ✅ 文件上传/下载/删除/重命名/移动/复制
- ✅ **分块上传** - 大文件分块上传、秒传、断点续传
- ✅ **文件分享** - 支持有效期和提取码的分享链接
- ✅ **回收站** - 软删除、恢复、永久删除
- ✅ **存储空间管理** - 用户配额管理、存储统计

### 🤖 AI 能力（AqiCloud-AI）

- ✅ **AI 聊天助理** - 对话历史、上下文摘要、实时搜索
- ✅ **AI 文档助手** - HTML/PDF/纯文本智能分析与概要生成
- ✅ **AI 网盘智答** - 自然语言查询网盘，意图识别自动路由
- ✅ **多 LLM 后端** - 支持 Ollama 与 OpenAI Compatible，运行时可切换
- ✅ **Token 统计** - 按用户/全局维度追踪消耗

### 🎨 前端界面（AqiCloud-Web）

- ✅ **现代设计** - 基于专业设计系统，精美 UI
- ✅ **响应式布局** - 完美适配桌面端和移动端
- ✅ **TypeScript** - 100% TypeScript 覆盖率
- ✅ **AI 集成** - 无缝集成 AI 聊天、文档分析、网盘智答

---

## 🏗️ 项目结构

```
AiDisk/
├── AqiCloud-Agent/    # 云盘后端 API（Go + Gin）
│   ├── cmd/server/          # 程序入口
│   ├── internal/           # 业务逻辑（controller/service/model）
│   ├── docs/               # Swagger API 文档
│   ├── Dockerfile           # Docker 构建文件
│   └── README.md           # 详细文档
│
├── AqiCloud-AI/       # AI 智能体中心（Go + Eino）
│   ├── cmd/agent/          # 程序入口
│   ├── internal/           # Agent/chat/doc/pan 四大模块
│   ├── Dockerfile           # Docker 构建文件
│   └── README.md           # 详细文档
│
├── AqiCloud-Web/      # 前端界面（Vue 3 + Vite）
│   ├── src/                # 源代码
│   ├── public/             # 静态资源
│   ├── Dockerfile           # Docker 构建文件
│   └── README.md           # 详细文档
│
├── docker-compose.yml  # Docker Compose 配置
├── deploy.sh           # 构建 & 打包脚本
├── Makefile             # 根构建命令
└── README.md           # 本文档
```

---

## 🚀 快速开始

### 方式一：Docker Compose 部署（推荐）

**最适合生产环境快速部署。**

#### 1. 克隆项目

```bash
git clone https://github.com/aqi-qihuan/AiDisk.git
cd AiDisk
```

#### 2. 配置环境变量

```bash
# 复制示例配置
cp AqiCloud-Agent/.env.example AqiCloud-Agent/.env
cp AqiCloud-AI/.env.example AqiCloud-AI/.env

# 编辑配置文件，填写必要参数
# 详见下方「配置说明」部分
```

#### 3. 启动服务

```bash
# 使用 Docker Compose 启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f
```

**访问地址：**

- **前端界面**: http://localhost:8081
- **Agent API**: http://localhost:9090
- **AI API**: http://localhost:8000
- **MinIO 控制台**: http://localhost:9001
- **Swagger 文档**: http://localhost:9090/swagger/index.html

---

### 方式二：本地开发环境

**适合开发调试。**

#### 1. 环境要求

| 依赖 | 版本要求 | 说明 |
|------|---------|------|
| **Go** | >= 1.26 | 后端语言 |
| **Node.js** | >= 18 | 前端构建 |
| **MySQL** | >= 8.0 | 数据库 |
| **Redis** | >= 7.0 | 缓存（AI 服务需要）|
| **MinIO** | latest | 对象存储 |
| **(可选) Ollama** | latest | 本地 LLM 部署 |

#### 2. 安装依赖

```bash
# 安装 Go 依赖
cd AqiCloud-Agent && go mod download
cd ../AqiCloud-AI && go mod download

# 安装 Node.js 依赖
cd AqiCloud-Web && npm install
```

#### 3. 启动依赖服务（MySQL + Redis + MinIO）

**使用 Docker 快速启动：**

```bash
# MySQL
docker run -d --name mysql \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=你的密码 \
  -e MYSQL_DATABASE=aqi-cloud-pan \
  mysql:8.0

# Redis
docker run -d --name redis \
  -p 6379:6379 \
  redis:7-alpine

# MinIO
docker run -d --name minio \
  -p 9000:9000 \
  -p 9001:9001 \
  -e MINIO_ROOT_USER=你的key \
  -e MINIO_ROOT_PASSWORD=你的密码 \
  minio/minio server /data --console-address ":9001"
```

#### 4. 配置环境变量

编辑 `AqiCloud-Agent/.env` 和 `AqiCloud-AI/.env`，填写数据库连接信息等。

**关键配置项：**

```bash
# AqiCloud-Agent/.env
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=你的密码
MYSQL_DATABASE=aqi-cloud-pan

MINIO_ENDPOINT=localhost:9000
MINIO_ACCESS_KEY=你的key
MINIO_SECRET_KEY=你的密码
MINIO_BUCKET=ai-pan

JWT_SECRET=你的加密字符串
```

```bash
# AqiCloud-AI/.env
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=你的密码

MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=你的密码
MYSQL_DATABASE=aqi-cloud-pan

JWT_SECRET=你的加密字符串

LLM_PROVIDER=ollama
LLM_OLLAMA_BASE_URL=http://localhost:11434
```

#### 5. 运行服务

```bash
# 终端 1：启动 Agent 服务
cd AqiCloud-Agent
go run cmd/server/main.go

# 终端 2：启动 AI 服务
cd AqiCloud-AI
go run cmd/agent/main.go

# 终端 3：启动前端
cd AqiCloud-Web
npm run dev
```

#### 6. 构建生产二进制

```bash
# 使用 Makefile 构建所有服务
make build-all

# 或单独构建
make build-agent    # Agent 服务
make build-ai       # AI 服务
make build-frontend # 前端
```

构建产物在 `bin/` 目录下。

---

## ⚙️ 配置说明

### AqiCloud-Agent 配置项

| 配置项 | 说明 | 默认值 | 必填 |
|--------|------|--------|------|
| `APP_NAME` | 应用名称 | AqiCloud-AgentPan API | ❌ |
| `DEBUG` | 调试模式 | false | ❌ |
| `LISTEN_ADDR` | 监听地址 | :9090 | ❌ |
| `MYSQL_HOST` | MySQL 主机 | - | ✅ |
| `MYSQL_PORT` | MySQL 端口 | 3306 | ❌ |
| `MYSQL_USER` | MySQL 用户 | - | ✅ |
| `MYSQL_PASSWORD` | MySQL 密码 | - | ✅ |
| `MYSQL_DATABASE` | MySQL 数据库 | - | ✅ |
| `MINIO_ENDPOINT` | MinIO 端点 | - | ✅ |
| `MINIO_ACCESS_KEY` | MinIO 访问密钥 | - | ✅ |
| `MINIO_SECRET_KEY` | MinIO 秘密密钥 | - | ✅ |
| `MINIO_BUCKET` | MinIO 桶名 | ai-pan | ❌ |
| `JWT_SECRET` | JWT 密钥（至少 32 字符） | - | ✅ |
| `JWT_ALGORITHM` | JWT 算法 | HS256 | ❌ |
| `JWT_EXPIRE_DAYS` | Token 过期天数 | 7 | ❌ |
| `DASHSCOPE_API_KEY` | 阿里云百炼 API Key | - | ❌ |
| `OLLAMA_BASE_URL` | Ollama API 地址 | - | ❌ |

### AqiCloud-AI 配置项

| 配置项 | 说明 | 默认值 | 必填 |
|--------|------|--------|------|
| `APP_NAME` | 应用名称 | AI智能体中心API服务 | ❌ |
| `DEBUG` | 调试模式 | false | ❌ |
| `LISTEN_ADDR` | 监听地址 | :8000 | ❌ |
| `REDIS_HOST` | Redis 主机 | - | ✅ |
| `REDIS_PORT` | Redis 端口 | 6379 | ❌ |
| `REDIS_PASSWORD` | Redis 密码 | - | ❌ |
| `MYSQL_HOST` | MySQL 主机 | - | ✅ |
| `MYSQL_PORT` | MySQL 端口 | 3306 | ❌ |
| `MYSQL_USER` | MySQL 用户 | - | ✅ |
| `MYSQL_PASSWORD` | MySQL 密码 | - | ✅ |
| `JWT_SECRET` | JWT 密钥（与 Agent 一致） | - | ✅ |
| `JWT_ALGORITHM` | JWT 算法 | HS256 | ❌ |
| `LLM_PROVIDER` | LLM 提供商 | ollama | ❌ |
| `LLM_MODEL_NAME` | LLM 模型名称 | qwen3.5:9b | ❌ |
| `LLM_OLLAMA_BASE_URL` | Ollama API 地址 | - | ❌ |

**完整配置说明详见：**

- AqiCloud-Agent: [AqiCloud-Agent/README.md](AqiCloud-Agent/README.md#配置说明)
- AqiCloud-AI: [AqiCloud-AI/README.md](AqiCloud-AI/README.md#配置说明)

---

## 📚 使用示例

### 1. 注册账号

```bash
curl -X POST http://localhost:9090/api/account/v1/register \
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
curl -X POST http://localhost:9090/api/account/v1/login \
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
    "token": "AQIeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "account_id": 1234567890,
    "username": "testuser"
  }
}
```

### 3. 上传文件

```bash
# 普通上传
curl -X POST http://localhost:9090/api/file/v1/upload \
  -H "token: YOUR_JWT_TOKEN" \
  -F "file=@/path/to/your/file.txt" \
  -F "parent_id=0"
```

### 4. AI 聊天

```bash
curl -X POST http://localhost:8000/api/chat/stream \
  -H "Content-Type: application/json" \
  -H "token: YOUR_JWT_TOKEN" \
  -d '{
    "message": "你好，请介绍一下自己",
    "history": [],
    "provider": "ollama"
  }'
```

### 5. 文档分析

```bash
curl -X POST http://localhost:8000/api/document/stream \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://example.com/report.pdf",
    "summary_type": "简洁",
    "language": "中文",
    "length": "200字"
  }'
```

### 6. 网盘查询

```bash
curl -X POST http://localhost:8000/api/pan/query \
  -H "Content-Type: application/json" \
  -H "token: YOUR_JWT_TOKEN" \
  -d '{"query": "查看我的存储空间使用情况"}'
```

**更多示例详见：**

- AqiCloud-Agent API 示例: [AqiCloud-Agent/README.md](AqiCloud-Agent/README.md#使用示例)
- AqiCloud-AI API 示例: [AqiCloud-AI/README.md](AqiCloud-AI/README.md#使用示例)

---

## 📖 API 文档

### AqiCloud-Agent（云盘 API）

项目集成了 Swagger API 文档，启动服务后访问：

```
http://localhost:9090/swagger/index.html
```

**主要 API 端点：**

| Method | Path | Auth | 说明 |
|--------|-------|------|------|
| POST | `/api/account/v1/register` | - | 用户注册 |
| POST | `/api/account/v1/login` | - | 用户登录 |
| GET | `/api/file/v1/list` | JWT | 文件列表 |
| POST | `/api/file/v1/upload` | JWT | 上传文件 |
| POST | `/api/file/v1/chunk/init` | JWT | 初始化分块上传 |
| POST | `/api/share/v1/create` | JWT | 创建分享 |
| GET | `/api/recycle/v1/list` | JWT | 回收站列表 |

### AqiCloud-AI（AI API）

**主要 API 端点：**

| Method | Path | Auth | 说明 |
|--------|-------|------|------|
| POST | `/api/chat/stream` | JWT | SSE 流式聊天 |
| GET | `/api/chat/providers` | - | 可用 LLM 提供商 |
| POST | `/api/chat/switch-provider` | JWT | 切换 LLM 提供商 |
| POST | `/api/document/stream` | - | SSE 流式文档分析 |
| POST | `/api/pan/query` | JWT | 自然语言查询网盘 |

---

## 🧪 运行测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test github.com/aqi/AqiCloud-Agent/internal/service
go test github.com/aqi/AqiCloud-AI/internal/service

# 查看测试覆盖率
go test -cover ./...

# 使用 Makefile
make test
```

---

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
git clone https://github.com/YOUR_USERNAME/AiDisk.git
cd AiDisk
```

#### 3. 创建功能分支

```bash
git checkout -b feature/your-feature-name
```

**分支命名规范：**

- `feature/` - 新功能
- `fix/` - Bug 修复
- `docs/` - 文档更新
- `refactor/` - 代码重构
- `test/` - 测试相关

#### 4. 开发和提交

```bash
# 进行你的修改
go fmt ./...
go mod tidy

# 运行测试
go test ./...

# 提交
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

- 遵循 Go 官方代码风格（使用 `go fmt` 格式化代码）
- 添加必要的注释，尤其是导出函数和复杂逻辑
- 确保每个函数都有适当的错误处理
- 编写单元测试，测试覆盖率不低于 60%
- 运行 `go vet` 和 `go lint` 检查代码质量
- Vue 组件使用 TypeScript，遵循 Vue 3 风格指南

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

---

## 📂 项目部署

### 生产环境部署

#### 1. 构建生产包

```bash
# 使用部署脚本
./deploy.sh

# 或手动构建
make build-all
```

#### 2. 上传到服务器

```bash
# 使用 SCP 上传
scp deploy-*.tar.gz user@server:/opt/aidisk/

# 或使用 rsync
rsync -avz bin/ user@server:/opt/aidisk/bin/
```

#### 3. 在服务器上解压并启动

```bash
# SSH 登录服务器
ssh user@server

# 解压
cd /opt/aidisk
tar xzf deploy-*.tar.gz

# 使用 Docker Compose 启动
docker-compose up -d

# 或直接运行二进制
./bin/AqiCloud-Agent &
./bin/AqiCloud-AI &
```

**详细部署指南详见：**

- [docs/architecture/lightweight-architecture.md](docs/architecture/lightweight-architecture.md)
- [docs/configuration/jwt-unification.md](docs/configuration/jwt-unification.md)

---

## 📄 许可证

本项目采用 **MIT 许可证**。详见 [LICENSE](LICENSE) 文件。

MIT 许可证允许：

- ✅ 商用
- ✅ 修改
- ✅ 分发
- ✅ 私有使用

仅需保留版权声明和许可证声明。

---

## 📧 联系方式

- **作者**: 阿七
- **邮箱**: 2316364297@qq.com
- **网站**: https://pan.aqi125.cn
- **GitHub**: https://github.com/aqi-qihuan/AiDisk
- **Issue Tracker**: [GitHub Issues](https://github.com/aqi-qihuan/AiDisk/issues)
- **Discussions**: [GitHub Discussions](https://github.com/aqi-qihuan/AiDisk/discussions)

---

## 🙏 致谢

感谢以下开源项目的支持：

### 后端

- [Gin](https://gin-gonic.com/) - Web 框架
- [GORM](https://gorm.io/) - ORM 库
- [MinIO](https://min.io/) - 对象存储
- [Redis](https://redis.io/) - 缓存数据库
- [Eino](https://github.com/cloudwego/eino) - AI Agent 框架
- [JWT](https://github.com/golang-jwt/jwt) - JWT 认证

### 前端

- [Vue 3](https://vuejs.org/) - 前端框架
- [Vite](https://vitejs.dev/) - 构建工具
- [Element Plus](https://element-plus.org/) - UI 组件库
- [TypeScript](https://www.typescriptlang.org/) - 类型系统

---

## 🔗 相关项目

| 项目 | 端口 | 说明 |
|------|------|------|
| AqiCloud-Agent | 9090 | 云盘后端 API（本仓库） |
| AqiCloud-AI | 8000 | AI 智能体中心 |
| AqiCloud-Web | 8081 | Vue 3 前端 |

---

## 🎯 Roadmap

### v1.0.0（当前版本）

- ✅ 完整的文件管理功能
- ✅ AI 聊天、文档分析、网盘智答
- ✅ Docker Compose 一键部署
- ✅ Swagger API 文档

### v1.1.0（计划中）

- ⏳ 文件版本管理
- ⏳ 文件协作编辑
- ⏳ AI 智能标签
- ⏳ 移动端 APP（React Native）

---

<div align="center">

**⭐ 如果这个项目对你有帮助，请给它一个 Star！⭐**

</div>
