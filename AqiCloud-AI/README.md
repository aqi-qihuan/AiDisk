# AqiCloud-Ai-Agent-Go

> 基于 Go + Eino 的轻量级 AI 智能体中心，提供聊天、文档处理、网盘查询三大 Agent 能力。空载内存 ~55MB。

## 功能特性

- **AI 聊天助理** — 对话历史、上下文摘要、实时搜索（Web Search Tool）
- **AI 文档助手** — HTML / PDF / 纯文本智能分析与概要生成
- **AI 网盘智答** — 自然语言查询网盘，意图识别自动路由
- **JWT 认证** — 与同生态网盘服务共享认证体系
- **多 LLM 后端** — 支持 Ollama 与 OpenAI Compatible 协议，运行时可切换
- **Token 统计** — 按用户 / 全局维度追踪消耗

## 技术栈

| 组件 | 用途 |
|------|------|
| Go 1.26 | 语言运行时 |
| Gin | HTTP 框架 |
| Eino | AI Agent 框架（CloudWeGo） |
| go-openai | OpenAI SDK，对接 Ollama / DashScope |
| JWT v5 | Token 认证 |
| Redis | 聊天历史存储（TTL） |
| MySQL + GORM | 网盘数据查询 |
| goquery / ledongthuc/pdf | 文档文本提取 |

## 快速开始

### 环境要求

- Go 1.26+
- Redis
- MySQL
- Ollama 或 OpenAI Compatible 服务

### 安装运行

```bash
git clone https://github.com/aqi-qihuan/AqiCloud-Ai-Agent-Go.git
cd AqiCloud-Ai-Agent-Go
go mod download

# 开发模式
go run ./cmd/agent/main.go

# 生产模式
go build -o ai-agent ./cmd/agent/main.go
./ai-agent
```

服务默认监听 `:8000`。

### 配置

复制 `.env.example` 为 `.env`，按需修改配置项。完整的环境变量参考见下方 `.env.example`。

### `.env.example`

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

# JWT（需与 AqiCloud-AgentPan-Go 保持一致）
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

## API 文档

### 根路径

| Method | Path       | Auth | 说明     |
|--------|------------|------|----------|
| GET    | `/`        | -    | 服务信息 |

### AI 聊天

| Method   | Path                           | Auth | 说明               |
|----------|--------------------------------|------|--------------------|
| POST     | `/api/chat/stream`             | JWT  | SSE 流式对话       |
| GET      | `/api/chat/providers`          | -    | 可用 LLM 提供商    |
| POST     | `/api/chat/switch-provider`    | JWT  | 切换 LLM 提供商    |
| GET      | `/api/chat/history`            | JWT  | 聊天历史           |
| DELETE   | `/api/chat/history`            | JWT  | 清空聊天历史       |
| GET      | `/api/chat/token-usage`        | JWT  | 用户 Token 用量    |
| GET      | `/api/chat/token-usage/global` | JWT  | 全局 Token 用量    |

### AI 文档

| Method | Path                     | Auth | 说明               |
|--------|--------------------------|------|--------------------|
| POST   | `/api/document/stream`   | -    | SSE 流式文档分析   |
| GET    | `/api/document/providers`| -    | 可用 LLM 提供商    |

### AI 网盘

| Method | Path                   | Auth | 说明             |
|--------|------------------------|------|------------------|
| POST   | `/api/pan/query`       | JWT  | 自然语言查询网盘 |
| GET    | `/api/pan/providers`   | -    | 可用 LLM 提供商  |

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

## LLM 后端切换

运行时通过 API 切换提供商：

```bash
curl -X POST http://localhost:8000/api/chat/switch-provider \
  -H "Content-Type: application/json" \
  -H "token: YOUR_JWT_TOKEN" \
  -d '{"provider": "openai_compatible"}'
```

## 项目结构

```
├── cmd/agent/main.go            # 入口
├── internal/
│   ├── agents/                  # ChatAgent / DocAgent / PanAgent
│   ├── core/                    # 配置、认证、CORS、Token 统计
│   ├── handlers/                # 路由注册
│   ├── models/                  # 请求 / 响应结构体
│   ├── services/                # 业务逻辑（聊天、文档、网盘）
│   ├── callbacks/               # Eino 回调（成本统计）
│   └── tools/                   # Web Search 等工具
└── .env.example                 # 配置模板
```

## 相关项目

| 项目 | 端口 | 说明 |
|------|------|------|
| AqiCloud-Ai-Agent-Go | 8000 | AI 智能体中心（本仓库） |
| AqiCloud-AgentPan-Go | 8080 | 网盘管理服务 |
| AqiCloud-AiPan | 8081 | Vue 3 前端 |
