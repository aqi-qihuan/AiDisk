# Eino 生产级实践：AiDisk 云盘 AI 服务技术解析

> 本文基于真实上线项目 AiDisk，分享使用字节 Eino 框架构建生产级 AI 服务的完整经验。

---

## 一、为什么选 Eino

在 AiDisk 的 AI 服务技术选型阶段，我评估了三个方案：

| 框架 | 语言 | 维护状态 | 决策 |
|------|------|---------|------|
| LangChainGo | Go | 不活跃，Python 特性缺失 | ❌ 放弃 |
| oai-go | Go | 仅封装 OpenAI 协议，无 Agent 抽象 | ❌ 太薄 |
| **Eino** | Go | CloudWeGo 团队持续维护 | ✅ 选用 |

Eino 是 CloudWeGo 出品的 AI 应用开发框架，原生 Go 实现，组件化设计彻底，Callback 和 Stream 接口符合 Go 的惯用风格。最关键的是有团队持续维护，长期看更稳。

---

## 二、架构概览

```
前端 (Vue 3)
    ↕ JWT 认证
AQICloud-AI 服务 (Go + Eino)
    ├── ChatAgent    → 通用对话
    ├── DocAgent     → 文档分析（URL/PDF，自动分块）
    ├── PanAgent     → 网盘智能查询（意图识别 → 查 MySQL）
    └── Agent 池     → 复用 LLM 客户端，降响应时间 30%
         ↕
    Ollama (本地)  ↔  DashScope (云端)  运行时切换
```

两个关键点：
1. **Agent 服务与文件服务分离**：AI 服务因模型调用耗时长，不影响文件上传下载
2. **多 LLM 后端**：通过 Eino 的 `ChatModel` 抽象，一行代码切换 provider

---

## 三、Eino 四层用法详解

### 3.1 ChatModel：多后端切换

```go
// 运行时切换，无需重启
func SwitchProvider(provider string) (chatmodel.ChatModel, error) {
    switch provider {
    case "ollama":
        return ollama.NewChatModel(ctx, &ollama.Config{
            Model:  "qwen2:7b",
            BaseURL: "http://localhost:11434/v1",
        })
    case "dashscope":
        return dashscope.NewChatModel(ctx, &dashscope.Config{
            APIKey: os.Getenv("DASHSCOPE_API_KEY"),
            Model:  "qwen-max",
        })
    default:
        return nil, fmt.Errorf("unknown provider: %s", provider)
    }
}
```

生产环境里，Ollama 作为默认（免费、低延迟），DashScope 作为降级（Ollama 挂了自动切换）。

### 3.2 Agent：三个独立 Agent

每个 Agent 继承 Eino 的 `Agent` 接口，有自己的系统提示词和工具集：

```go
type PanAgent struct {
    chatModel chatmodel.ChatModel
    tools      []tool.BaseTool  // 查文件列表、查存储空间等工具
    systemPrompt string
}

func (a *PanAgent) Invoke(ctx context.Context, input []*message.Message) (*message.Placeholder, error) {
    // 意图识别 → 选工具 → 执行 → 生成回答
}
```

PanAgent 的意图识别是纯提示词工程：给 LLM 一个工具列表，让它自己决定调哪个工具，再把工具返回的数据整理成自然语言。

### 3.3 Callback：零侵入 Token 统计

```go
type TokenCounterCallback struct{}

func (t *TokenCounterCallback) OnEnd(ctx context.Context, info *callbacks.RunInfo, output callbacks.CallbackOutput) {
    if output == nil {
        return
    }
    // 从 output 里取 token 数，写入 Redis
    usage := output.(*callbacks.CallbackOutput).TokenUsage
    redisClient.Incr(ctx, fmt.Sprintf("token:%s", userID), usage.TotalTokens)
}
```

把这个 Callback 注册到 ChatModel，所有 LLM 调用的 token 消耗自动被记录，业务代码完全不用改。

### 3.4 Stream：SSE 流式输出

```go
stream, err := chatModel.Stream(ctx, messages)
if err != nil { return err }
defer stream.Close()  // 必须 Close，否则泄露连接

for {
    chunk, err := stream.Recv()
    if errors.Is(err, io.EOF) {
        break
    }
    // 通过 SSE 推送给前端
    sendSSE(w, chunk.Content)
}
```

Eino 的 Stream 接口把底层的 chunk 处理封装好了，只需要 `Recv()` 循环读取。

---

## 四、生产踩坑实录

### 坑 1：长对话上下文爆炸

**问题**：对话超过 20 轮，`messages` 里的 token 数超过模型上下文窗口，请求直接失败。

**解决**：用 Eino 的 `TrimMessages` 函数，配合自写摘要机制：

```go
// 超阈值时，把旧消息压缩成摘要
if len(messages) > maxRounds {
    summary := generateSummary(ctx, messages[:len(messages)-maxRounds])
    messages = append([]*message.Message{
        message.NewSystemMessage(summary),
    }, messages[len(messages)-maxRounds:]...)
}
```

摘要本身也用 LLM 生成，塞回系统提示词里，AI 不会"失忆"。

### 坑 2：ChatModel 并发不安全

**问题**：文档分析并行处理多个 chunk，每个 chunk 创建一个 LLM 客户端，并发一高直接把内存吃光，且多 goroutine 读写同一实例会 panic。

**解决**：实现 Agent 池：

```go
type AgentPool struct {
    mu      sync.RWMutex
    agents  map[string]*AgentInstance  // provider -> instance
    maxConn int
}

func (p *AgentPool) Get(provider string) (*AgentInstance, error) {
    p.mu.RLock()
    if agent, ok := p.agents[provider]; ok {
        p.mu.RUnlock()
        return agent, nil
    }
    p.mu.RUnlock()

    p.mu.Lock()
    defer p.mu.Unlock()
    // 创建新实例，加入池
    agent := newAgentInstance(provider)
    p.agents[provider] = agent
    return agent, nil
}
```

优化后，平均响应时间降了 30%，内存峰值降了 40%。

### 坑 3：Stream 不 Close 泄露连接

**问题**：文档分析并发高时，文件句柄耗尽，新请求全部超时。

**根治**：所有 `stream` 获取处，必须 `defer stream.Close()`，且在所有错误分支（early return）之前都要 Close。还加了 panic recovery：

```go
func streamWithRecovery(w http.ResponseWriter, model chatmodel.ChatModel, msgs []*message.Message) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("[stream] panic recovered: %v", r)
        }
    }()

    stream, err := model.Stream(ctx, msgs)
    if err != nil { return }
    defer stream.Close()

    // ... Recv 循环
}
```

---

## 五、部署与监控

整套服务 Docker Compose 编排，跑在 4 核 4GB 服务器上：

| 组件 | 内存限制 | 说明 |
|------|---------|------|
| MySQL | 512MB | buffer pool 64MB |
| Redis | 80MB | maxmemory 64MB |
| MinIO | 256MB | 小内存模式 |
| AQICloud-Agent | 1GB | 文件服务 |
| AQICloud-AI | 512MB | Eino AI 服务 |
| Nginx | 64MB | 反向代理 |
| **系统预留** | **~1.5GB** | OS + 其他 |

CPU 长期 1-2%，内存使用不到 2GB，4GB 服务器绰绰有余。

Token 配额通过 Redis 实时监控，前端在「AI 助手」页面展示当日已用 token 数，超配额时 API 直接拒绝。

---

## 六、总结与资源

AiDisk 是 Eino 在生产环境的完整实践案例，四层接口（ChatModel / Agent / Callback / Stream）都有真实业务场景对应。

**如果你正在用 Eino 做 AI 应用，可以参考：**

- **GitHub（完整代码）**：https://github.com/aqi-qihuan/AiDisk
- **线上 Demo（可体验 AI 聊天和文档分析）**：https://pan.aqi125.cn
- **Eino 官方仓库**：https://github.com/cloudwego/eino

生产环境踩的坑我都写在了代码注释里，希望能帮大家少走点弯路。

---

*作者：阿七，CloudWeGo 社区成员，全栈开发者，AiDisk 项目作者。*
