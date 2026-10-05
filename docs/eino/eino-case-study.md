# Eino 在生产级云盘项目中的完整实践 —— AiDisk 案例分享

> 项目：AiDisk（开源 AI 云盘系统）
> GitHub：https://github.com/aqi-qihuan/AiDisk
> 线上 Demo：https://pan.aqi125.cn

---

## 项目背景

AiDisk 是一个带 AI 能力的个人云盘系统，后端 Go（Gin），前端 Vue 3 + TypeScript，部署在 4 核 4GB 轻量云服务器上，已在 GitHub 开源并有真实用户在使用。

AI 服务底层选用了字节开源的 **Eino** 框架（CloudWeGo 出品）。本文分享 AiDisk 在生产环境中使用 Eino 的四层实践，以及踩过的坑。

---

## 四层用法

### 1. ChatModel 抽象 —— 多 LLM 后端切换

通过 Eino 的 `ChatModel` 接口，同时接入 Ollama 本地模型和 DashScope 云端模型，两个 provider 都实现了 Eino 的模型接口。运行时调一个 API 就能切换，无需重启服务。

```go
chatModel, _ := einoagent.GetChatModelWithProvider("dashscope")
```

### 2. Agent 抽象 —— 三个独立 Agent

基于 Eino 的 Agent 接口实现了三个 Agent：

- **ChatAgent**：通用对话，带历史记忆
- **DocAgent**：文档分析，支持 URL 抓取和 PDF，自动分块处理超长文档
- **PanAgent**：网盘智能查询，意图识别 → 查 MySQL → 自然语言回答

每个 Agent 有独立的系统提示词和工具集，互不干扰。

### 3. Callback 机制 —— 零侵入 Token 统计

Eino 提供了一套回调钩子（OnStart / OnEnd / OnError）。在 LLM 调用前后埋点，把消耗的 token 数写入 Redis，前端实时展示配额。

**关键优势**：业务代码完全不用改，加一个 Callback 就行。

### 4. Stream 接口 —— SSE 流式输出

聊天和文档分析都用了 Eino 的 `Stream` 接口，封装了底层 chunk 处理，前端通过 SSE 逐字渲染 AI 回复，体验流畅。

---

## 生产踩坑

### 坑 1：长对话上下文爆炸

**现象**：对话超过 20 轮，token 直接爆掉，请求失败。

**解决**：用 Eino 的 `TrimMessages` 函数配合自写摘要机制 —— 对话超阈值时，把旧消息压缩成一份摘要塞回系统提示词，既控制 token 又不让 AI "失忆"。

### 坑 2：ChatModel 并发不安全

**现象**：多 goroutine 同时调用同一个模型实例，出现数据竞争，偶发 panic。

**解决**：实现 Agent 池，每个实例加锁，或通过池分配不同实例，避免并发读写同一模型客户端。

### 坑 3：Stream 不 Close 泄露连接

**现象**：文档分析并发一高，文件句柄耗尽，新请求全部超时。

**解决**：`defer stream.Close()` 必须写，且在所有错误分支都要 Close。还加了 panic recovery，防止异常退出时漏关。

---

## 优化效果

| 指标 | 优化前 | 优化后 |
|------|--------|--------|
| 平均响应时间 | 基线 | **↓ 30%** |
| 内存峰值 | 基线 | **↓ 40%** |
| 并发稳定性 | 偶发 panic | 零崩溃 |

---

## 总结

Eino 的四大核心接口（ChatModel / Agent / Callback / Stream）在 AiDisk 里都有真实业务场景对应，不是 example 级别的调用，而是跑在生产环境、有真实用户在使用。

如果你正在用 Eino 做 AI 应用，欢迎参考 AiDisk 的代码，生产环境踩的坑我都写在了代码注释里。

- **GitHub**：https://github.com/aqi-qihuan/AiDisk
- **线上 Demo**：https://pan.aqi125.cn
- **Eino 框架**：https://github.com/cloudwego/eino

---

*作者：阿七，CloudWeGo 社区成员，AiDisk 项目作者。*
