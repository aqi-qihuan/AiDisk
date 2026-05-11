# AqiCloud-Ai

<div align="center">

**基于 Go + Eino 的轻量级 AI 智能体中心**

[![Go Version](https://img.shields.io/badge/Go-1.26-blue.svg)](https://golang.org/)
[![Gin](https://img.shields.io/badge/Gin-v1.12.0-00A6ED.svg)](https://gin-gonic.com/)
[![License](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![Eino](https://img.shields.io/badge/Eino-v0.8.13-green.svg)](https://github.com/cloudwego/eino)

</div>

## 📖 项目简介

AqiCloud-Ai 是一个基于 Go 语言和 Eino AI Agent 框架开发的轻量级 AI 智能体中心，提供聊天、文档处理、网盘查询三大 Agent 能力。系统采用 RESTful API 设计，集成多种 AI 服务（Ollama、OpenAI Compatible），空载内存仅 ~55MB，适合资源受限环境部署。

### ✨ 核心特性

- 🤖 **AI 聊天助理** — 对话历史、上下文摘要、实时搜索（Web Search Tool）
- 📄 **AI 文档助手** — HTML / PDF / 纯文本智能分析与概要生成
- 🔍 **AI 网盘智答** — 自然语言查询网盘，意图识别自动路由
- 🔐 **JWT 认证** — 与同生态网盘服务共享认证体系
- 🔄 **多 LLM 后端** — 支持 Ollama 与 OpenAI Compatible 协议，运行时可切换
- 📊 **Token 统计** — 按用户 / 全局维度追踪消耗
- 🐳 **容器化部署** — 提供 Dockerfile，支持容器化部署

## 🛠️ 技术栈

### 后端框架
- **语言**: Go 1.26
- **Web 框架**: Gin v1.12.0
- **AI 框架**: Eino v0.8.13 (CloudWeGo)
- **认证**: JWT (golang-jwt/jwt/v5)

### 数据存储
- **数据库**: MySQL 8.0+
- **缓存**: Redis 7.0+
- **ORM**: GORM v1.31.1

### AI 服务
- **Ollama** (本地 LLM 部署)
- **OpenAI Compatible** (DashScope 等)

### 文档与工具
- **配置管理**: godotenv
- **文档处理**: goquery / ledongthuc/pdf

## 📂 项目结构

```
AqiCloud-Ai/
├── cmd/agent/main.go            # 程序入口
├── internal/
│   ├── agents/                  # ChatAgent / DocAgent / PanAgent
│   │   ├── agent.go            # Agent 基类
│   │   ├── model.go            # LLM 模型管理
│   │   └── pool.go             # Agent 池
│   ├── core/                    # 配置、认证、CORS、Token 统计
│   │   ├── config.go           # 配置管理
│   │   ├── auth.go             # JWT 认证
│   │   ├── cors.go             # CORS 中间件
│   │   ├── db.go               # 数据库初始化
│   │   ├── llm.go              # LLM 配置
│   │   └── token_tracker.go   # Token 统计
│   ├── handlers/                # 路由注册
│   │   └── routes.go          # 路由定义
│   ├── models/                  # 请求 / 响应结构体
│   │   ├── chat_schemas.go    # 聊天请求/响应
│   │   ├── doc_schemas.go     # 文档请求/响应
│   │   ├── pan_schemas.go     # 网盘请求/响应
│   │   ├── chat_log.go         # 聊天记录模型
│   │   └── json_response.go   # 统一响应格式
│   ├── services/                # 业务逻辑
│   │   ├── chat_service.go    # 聊天服务
│   │   ├── doc_service.go     # 文档服务
│   │   └── pan_service.go     # 网盘服务
│   ├── callbacks/               # Eino 回调
│   │   └── cost.go            # 成本统计回调
│   └── tools/                   # AI 工具
│       └── web_search.go      # Web 搜索工具
├── .env.example                 # 环境变量示例
├── Dockerfile                   # Docker 构建文件
├── Makefile                     # 构建脚本
├── go.mod                       # Go 模块文件
└── go.sum                       # 依赖校验文件
```

## 🚀 安装步骤

### 环境要求

**方式一：Docker 部署**
- Docker 20.10+
- Docker Compose 2.0+

**方式二：本地开发**
- Go 1.26+
- Redis 7.0+
- MySQL 8.0+
- Ollama 或 OpenAI Compatible 服务

### 方式一：Docker 部署（推荐）

最适合生产环境快速部署。

#### 1. 克隆项目

```bash
git clone https://github.com/aqi-qihuan/AqiCloud-Ai.git
cd AqiCloud-Ai
```

#### 2. 配置环境变量

```bash
cp .env.example .env
# 编辑 .env 文件，修改配置项
```

#### 3. 使用 Docker Compose 启动

创建 `docker-compose.yml`：

```yaml
version: '3.8'

services:
  ai-agent:
    build: .
    ports:
      - "8000:8000"
    env_file:
      - .env
    depends_on:
      - redis
      - mysql
    networks:
      - ai-network

  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    volumes:
      - redis-data:/data
    networks:
      - ai-network

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: your-password
      MYSQL_DATABASE: aqi-cloud-pan
    ports:
      - "3306:3306"
    volumes:
      - mysql-data:/var/lib/mysql
    networks:
      - ai-network

networks:
  ai-network:
    driver: bridge

volumes:
  redis-data:
  mysql-data:
```

启动服务：

```bash
docker-compose up -d
```

#### 4. 查看日志

```bash
docker-compose logs -f ai-agent
```

### 方式二：从源码构建

适合开发调试。

#### 1. 克隆项目

```bash
git clone https://github.com/aqi-qihuan/AqiCloud-Ai.git
cd AqiCloud-Ai
```

#### 2. 安装依赖

```bash
go mod download
```

#### 3. 配置环境变量

```bash
cp .env.example .env
# 根据实际情况编辑 .env 文件
```

#### 4. 启动依赖服务（Redis + MySQL）

使用 Docker 快速启动依赖：

```bash
docker run -d --name redis -p 6379:6379 redis:7-alpine
docker run -d --name mysql -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=your-password \
  -e MYSQL_DATABASE=aqi-cloud-pan \
  mysql:8.0
```

#### 5. 运行项目

```bash
# 开发模式（支持热重载，需安装 air 等工具）
go run ./cmd/agent/main.go

# 或使用 Makefile
make run
```

#### 6. 构建生产二进制

```bash
# 使用 Makefile
make build

# 或直接使用 go build
go build -o bin/ai-agent ./cmd/agent
./bin/ai-agent
```

### 方式三：直接使用编译好的二进制

适合快速测试。

```bash
# 从 Release 页面下载预编译二进制
# 或使用 go install
go install github.com/aqi/AqiCloud-Ai/cmd/agent@latest

# 配置并运行
cp .env.example .env
./agent
```

服务默认监听 `:8000`。

## ⚙️ 配置说明

### 完整配置项

| 配置项 | 说明 | 默认值 | 必填 |
|--------|------|--------|------|
| `APP_NAME` | 应用名称 | AI智能体中心API服务 | ❌ |
| `DEBUG` | 调试模式 | false | ❌ |
| `LISTEN_ADDR` | 监听地址 | :8000 | ❌ |
| `REDIS_HOST` | Redis 主机地址 | 127.0.0.1 | ✅ |
| `REDIS_PORT` | Redis 端口 | 6379 | ❌ |
| `REDIS_DB` | Redis 数据库 | 0 | ❌ |
| `REDIS_PASSWORD` | Redis 密码 | - | ❌ |
| `REDIS_MAX_CONNECTIONS` | Redis 最大连接数 | 10 | ❌ |
| `REDIS_CHAT_HISTORY_TTL` | 聊天历史 TTL（秒） | 86400 | ❌ |
| `MYSQL_HOST` | MySQL 主机地址 | 127.0.0.1 | ✅ |
| `MYSQL_PORT` | MySQL 端口 | 3306 | ❌ |
| `MYSQL_USER` | MySQL 用户名 | root | ✅ |
| `MYSQL_PASSWORD` | MySQL 密码 | - | ✅ |
| `MYSQL_DATABASE` | MySQL 数据库名 | aqi-cloud-pan | ✅ |
| `JWT_SECRET_KEY` | JWT 密钥 | - | ✅ |
| `JWT_ALGORITHM` | JWT 算法 | HS256 | ❌ |
| `JWT_ACCESS_TOKEN_EXPIRE_MINUTES` | JWT 过期时间（分钟） | 30 | ❌ |
| `JWT_LOGIN_SUBJECT` | JWT 主题 | AQI | ❌ |
| `LLM_PROVIDER` | LLM 提供商 | ollama | ❌ |
| `LLM_MODEL_NAME` | LLM 模型名称 | qwen3.5:9b | ❌ |
| `LLM_OLLAMA_BASE_URL` | Ollama API 地址 | http://127.0.0.1:11434 | ❌ |
| `LLM_BASE_URL` | OpenAI Compatible API 地址 | https://dashscope.aliyuncs.com/compatible-mode/v1 | ❌ |
| `LLM_API_KEY` | LLM API Key | - | ❌ |
| `LLM_TEMPERATURE` | LLM 温度参数 | 0.7 | ❌ |
| `LLM_STREAMING` | LLM 流式输出 | true | ❌ |
| `FRONTEND_BASE_URL` | 前端基础 URL | http://127.0.0.1:8081 | ❌ |
| `TOKEN_DAILY_LIMIT` | 每日 Token 限制 | 100000 | ❌ |
| `TOKEN_HOURLY_LIMIT` | 每小时 Token 限制 | 10000 | ❌ |
| `MAX_CHAT_HISTORY_MESSAGES` | 最大聊天历史消息数 | 20 | ❌ |
| `MAX_SUMMARY_LENGTH` | 最大摘要长度 | 500 | ❌ |
| `SUMMARY_TRIGGER_THRESHOLD` | 摘要触发阈值 | 10 | ❌ |

### 配置示例

<details>
<summary>点击查看完整 .env.example</summary>

```env
# Redis
REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_DB=0
REDIS_PASSWORD=
REDIS_MAX_CONNECTIONS=10
REDIS_CHAT_HISTORY_TTL=86400

# MySQL
MYSQL_HOST=127.0.0.1
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=
MYSQL_DATABASE=aqi-cloud-pan

# JWT（需与 AqiCloud-Agent 保持一致）
JWT_SECRET_KEY=your-jwt-secret-key-here
JWT_ALGORITHM=HS256
JWT_ACCESS_TOKEN_EXPIRE_MINUTES=30
JWT_LOGIN_SUBJECT=AQI

# LLM
LLM_PROVIDER=ollama
LLM_MODEL_NAME=qwen3.5:9b
LLM_OLLAMA_BASE_URL=http://127.0.0.1:11434
LLM_BASE_URL=https://dashscope.aliyuncs.com/compatible-mode/v1
LLM_API_KEY=your-api-key-here
LLM_TEMPERATURE=0.7
LLM_STREAMING=true

# App
APP_NAME=AI智能体中心API服务
DEBUG=false
LISTEN_ADDR=:8000
FRONTEND_BASE_URL=http://127.0.0.1:8081

# Token Limits
TOKEN_DAILY_LIMIT=100000
TOKEN_HOURLY_LIMIT=10000
MAX_CHAT_HISTORY_MESSAGES=20
MAX_SUMMARY_LENGTH=500
SUMMARY_TRIGGER_THRESHOLD=10
```
</details>

## 📚 使用示例

### cURL 示例

**1. 流式对话**

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

**2. 查看可用 LLM 提供商**

```bash
curl http://localhost:8000/api/chat/providers
```

响应示例：

```json
{
  "code": 0,
  "data": {
    "providers": ["ollama", "openai_compatible"],
    "current": "ollama"
  }
}
```

**3. 切换 LLM 提供商**

```bash
curl -X POST http://localhost:8000/api/chat/switch-provider \
  -H "Content-Type: application/json" \
  -H "token: YOUR_JWT_TOKEN" \
  -d '{"provider": "openai_compatible"}'
```

**4. 查看 Token 用量**

```bash
curl http://localhost:8000/api/chat/token-usage \
  -H "token: YOUR_JWT_TOKEN"
```

响应示例：

```json
{
  "code": 0,
  "data": {
    "user_id": "123",
    "daily_used": 1500,
    "daily_limit": 100000,
    "hourly_used": 200,
    "hourly_limit": 10000
  }
}
```

**5. 文档分析**

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

**6. 网盘查询**

```bash
curl -X POST http://localhost:8000/api/pan/query \
  -H "Content-Type: application/json" \
  -H "token: YOUR_JWT_TOKEN" \
  -d '{"query": "查看我的存储空间使用情况"}'
```

### JavaScript/TypeScript 示例

**流式对话（使用 Fetch API）**

```javascript
async function chatWithAI(message, token) {
  const response = await fetch('http://localhost:8000/api/chat/stream', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'token': token
    },
    body: JSON.stringify({
      message: message,
      history: [],
      provider: 'ollama'
    })
  });

  const reader = response.body.getReader();
  const decoder = new TextDecoder();

  while (true) {
    const { done, value } = await reader.read();
    if (done) break;

    const chunk = decoder.decode(value);
    const lines = chunk.split('\n');

    for (const line of lines) {
      if (line.startsWith('data: ')) {
        const data = JSON.parse(line.slice(6));
        if (data.code === 0) {
          console.log(data.data); // 流式输出内容
        }
      }
    }
  }
}

// 使用示例
chatWithAI('你好！', 'YOUR_JWT_TOKEN');
```

**使用 Axios**

```javascript
import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8000',
  headers: {
    'Content-Type': 'application/json'
  }
});

// 设置 token
api.defaults.headers.common['token'] = 'YOUR_JWT_TOKEN';

// 获取可用提供商
async function getProviders() {
  const response = await api.get('/api/chat/providers');
  console.log(response.data);
}

// 发送聊天消息
async function sendMessage(message) {
  const response = await api.post('/api/chat/stream', {
    message: message,
    history: [],
    provider: 'ollama'
  });

  return response.data;
}
```

### Python 示例

```python
import requests

BASE_URL = "http://localhost:8000"
TOKEN = "YOUR_JWT_TOKEN"

headers = {
    "Content-Type": "application/json",
    "token": TOKEN
}

# 1. 获取可用 LLM 提供商
def get_providers():
    response = requests.get(f"{BASE_URL}/api/chat/providers")
    return response.json()

# 2. 流式对话
def chat_stream(message, history=None):
    if history is None:
        history = []

    response = requests.post(
        f"{BASE_URL}/api/chat/stream",
        headers=headers,
        json={
            "message": message,
            "history": history,
            "provider": "ollama"
        },
        stream=True
    )

    for line in response.iter_lines():
        if line:
            line = line.decode('utf-8')
            if line.startswith('data: '):
                print(line[6:])  # 处理 SSE 数据

# 3. 切换 LLM 提供商
def switch_provider(provider):
    response = requests.post(
        f"{BASE_URL}/api/chat/switch-provider",
        headers=headers,
        json={"provider": provider}
    )
    return response.json()

# 4. 查看 Token 用量
def get_token_usage():
    response = requests.get(
        f"{BASE_URL}/api/chat/token-usage",
        headers=headers
    )
    return response.json()

# 使用示例
if __name__ == "__main__":
    # 获取提供商列表
    providers = get_providers()
    print(providers)

    # 发送聊天消息
    chat_stream("你好，请介绍一下自己")
```

## 📡 API 文档

### 主要 API 端点

#### 根路径

| Method | Path | Auth | 说明 |
|--------|-------|------|------|
| GET    | `/`   | -    | 服务信息 |

#### AI 聊天

| Method   | Path                           | Auth | 说明               |
|----------|--------------------------------|------|--------------------|
| POST     | `/api/chat/stream`             | JWT  | SSE 流式对话       |
| GET      | `/api/chat/providers`          | -    | 可用 LLM 提供商    |
| POST     | `/api/chat/switch-provider`    | JWT  | 切换 LLM 提供商    |
| GET      | `/api/chat/history`            | JWT  | 聊天历史           |
| DELETE   | `/api/chat/history`            | JWT  | 清空聊天历史       |
| GET      | `/api/chat/token-usage`        | JWT  | 用户 Token 用量    |
| GET      | `/api/chat/token-usage/global` | JWT  | 全局 Token 用量    |

#### AI 文档

| Method | Path                     | Auth | 说明               |
|--------|--------------------------|------|--------------------|
| POST   | `/api/document/stream`   | -    | SSE 流式文档分析   |
| GET    | `/api/document/providers`| -    | 可用 LLM 提供商    |

#### AI 网盘

| Method | Path                 | Auth | 说明             |
|--------|----------------------|------|------------------|
| POST   | `/api/pan/query`     | JWT  | 自然语言查询网盘 |
| GET    | `/api/pan/providers` | -    | 可用 LLM 提供商  |

### 认证

需要认证的接口通过 `token` 请求头传递 JWT，也兼容 `Authorization` 头。

### 请求示例

**流式对话** — `POST /api/chat/stream`

```json
{
  "message": "今天天气怎么样？",
  "history": [
    {"role": "user", "content": "你好"},
    {"role": "assistant", "content": "你好！有什么可以帮你的？"}
  ],
  "provider": "ollama"
}
```

SSE 响应：

```
data: {"code":0,"data":"今天"}

data: {"code":0,"data":"天气"}

data: {"code":0,"data":"晴朗"}
```

**网盘查询** — `POST /api/pan/query`

```json
{
  "query": "查看我的存储空间"
}
```

后端自动识别意图（storage / file_list / file_statistics）并路由查询。

**文档分析** — `POST /api/document/stream`

```json
{
  "url": "https://example.com/report.pdf",
  "summary_type": "简洁",
  "language": "中文",
  "length": "200字",
  "additional_instructions": "关注结论部分"
}
```

## 🔄 LLM 后端切换

运行时通过 API 切换提供商：

```bash
curl -X POST http://localhost:8000/api/chat/switch-provider \
  -H "Content-Type: application/json" \
  -H "token: YOUR_JWT_TOKEN" \
  -d '{"provider": "openai_compatible"}'
```

## 🧪 运行测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test github.com/aqi/AqiCloud-Ai/internal/service

# 查看测试覆盖率
go test -cover ./...

# 使用 Makefile
make test
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
git clone https://github.com/YOUR_USERNAME/AqiCloud-Ai.git
cd AqiCloud-Ai
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
go mod tidy  # 清理依赖
gofmt -w .   # 格式化代码
go test ./... # 运行测试

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
git commit -m "feat: 添加文档分析功能"
git commit -m "fix: 修复流式对话断开的问题"
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
- 运行 `go vet` 和 `go lint` 检查代码质量

### 报告 Bug

提交 Issue 时，请包含以下信息：

- 🐛 **Bug 描述**：简洁明确地描述问题
- 🔄 **复现步骤**：详细描述如何复现该 Bug
- ✅ **预期行为**：你认为的正确行为
- ❌ **实际行为**：实际发生的情况
- 🖥️ **环境信息**：操作系统、Go 版本、MySQL 版本等
- 📝 **相关日志**：错误日志或截图

**Bug 报告模板：**

```markdown
## Bug 描述
简要描述 Bug

## 复现步骤
1. 步骤 1
2. 步骤 2
3. 步骤 3

## 预期行为
描述期望的正确行为

## 实际行为
描述实际发生的错误行为

## 环境信息
- OS: [e.g., Windows 10, Ubuntu 20.04]
- Go 版本: [e.g., 1.26]
- 依赖版本: [e.g., Gin 1.12.0]

## 错误信息/日志
粘贴相关错误日志

## 截图
如果适用，添加截图
```

### 功能建议

提交功能建议时，请说明：

- 💡 **功能描述**：详细描述你希望添加的功能
- 🎯 **使用场景**：为什么需要这个功能
- 📐 **实现思路**：如果有，提供你的实现思路
- 📷 **参考资料**：提供相关的截图或文档链接

## 📄 许可证

本项目采用 **MIT 许可证**。详见 [LICENSE](LICENSE) 文件。

MIT 许可证允许：
- ✅ 商用
- ✅ 修改
- ✅ 分发
- ✅ 私有使用

仅需保留版权声明和许可证声明。

## 📧 联系方式

- **作者**: 阿七
- **邮箱**: 2316364297@qq.com
- **GitHub**: https://github.com/aqi-qihuan/AiDisk
- **Issue Tracker**: [GitHub Issues](https://github.com/aqi-qihuan/AqiCloud-Ai/issues)
- **Discussions**: [GitHub Discussions](https://github.com/aqi-qihuan/AqiCloud-Ai/discussions)

## 🙏 致谢

感谢以下开源项目的支持：

- [Gin](https://gin-gonic.com/) - Web 框架
- [Eino](https://github.com/cloudwego/eino) - AI Agent 框架
- [GORM](https://gorm.io/) - ORM 库
- [Redis](https://redis.io/) - 缓存数据库
- [JWT](https://github.com/golang-jwt/jwt) - JWT 认证

## 🔗 相关项目

| 项目 | 端口 | 说明 |
|------|------|------|
| AqiCloud-Ai | 8000 | AI 智能体中心（本仓库） |
| AqiCloud-Agent | 8080 | 网盘管理服务 |
| AqiCloud-Web | 8081 | Vue 3 前端 |

---

<div align="center">

**⭐ 如果这个项目对你有帮助，请给它一个 Star！⭐**

</div>
