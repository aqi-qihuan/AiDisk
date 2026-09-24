---
marp: true
theme: default
paginate: true
backgroundColor: '#0a0a0a'
backgroundImage: url('./assets/hok-bg.svg')
style: |
  /* ===== 全局基础 - HOK 暗黑史诗 ===== */
  section {
    font-family: 'Cinzel', 'Source Han Serif SC', 'Songti SC', 'PingFang SC', serif;
    font-size: 26px;
    color: #d8c8a8;
    background-image: url('./assets/hok-bg.svg');
    background-size: cover;
    background-position: center;
    padding: 70px 80px 60px 80px;
    position: relative;
  }

  /* ===== 页眉装饰条 ===== */
  section::before {
    content: '';
    position: absolute;
    top: 18px;
    left: 0;
    right: 0;
    height: 24px;
    background-image: url('./assets/hok-header.svg');
    background-repeat: no-repeat;
    background-position: center;
    background-size: 100% 24px;
    pointer-events: none;
    z-index: 1;
  }

  /* ===== 页脚装饰条 ===== */
  section::after {
    content: '';
    position: absolute;
    bottom: 18px;
    left: 0;
    right: 0;
    height: 24px;
    background-image: url('./assets/hok-footer.svg');
    background-repeat: no-repeat;
    background-position: center;
    background-size: 100% 24px;
    pointer-events: none;
    z-index: 1;
    opacity: 0.5;
  }

  /* Marp 默认页码隐藏（被装饰条替代） */
  section > footer {
    display: none;
  }

  /* ===== 标题 - 金色史诗感 ===== */
  h1 {
    color: #D97706;
    font-size: 54px;
    font-weight: 900;
    letter-spacing: 6px;
    text-shadow:
      0 1px 0 #fbbf24,
      0 2px 4px #000000aa,
      0 0 20px #D9770644;
    border-bottom: 2px solid #D97706;
    border-image: linear-gradient(90deg, transparent, #D97706, #fbbf24, #D97706, transparent) 1;
    padding-bottom: 18px;
    margin-bottom: 28px;
    margin-top: 20px;
    font-family: 'Cinzel', 'Source Han Serif SC', serif;
    position: relative;
    z-index: 2;
  }

  h2 {
    color: #fbbf24;
    font-size: 38px;
    font-weight: 700;
    letter-spacing: 4px;
    text-shadow:
      0 1px 0 #fde68a,
      0 2px 4px #000000aa;
    margin-top: 14px;
    margin-bottom: 22px;
    font-family: 'Cinzel', 'Source Han Serif SC', serif;
    position: relative;
    z-index: 2;
  }

  h3 {
    color: #D97706;
    font-size: 28px;
    font-weight: 600;
    letter-spacing: 2px;
    font-family: 'Cinzel', 'Source Han Serif SC', serif;
    position: relative;
    z-index: 2;
  }

  /* ===== 正文 ===== */
  p, li, td, th {
    color: #c8b896;
    line-height: 1.75;
    font-family: 'PingFang SC', 'Source Han Sans SC', sans-serif;
    position: relative;
    z-index: 2;
  }

  strong {
    color: #fbbf24;
    font-weight: 700;
    text-shadow: 0 1px 2px #000000aa;
  }

  /* ===== 代码块 - 暗金代码框 ===== */
  pre {
    background: #0d0d0d !important;
    border: 1px solid #D9770666;
    border-left: 4px solid #D97706;
    border-radius: 4px;
    padding: 20px 24px;
    box-shadow:
      inset 0 0 30px #D9770608,
      0 4px 12px #00000088;
    overflow: hidden;
    position: relative;
    z-index: 2;
  }

  pre code {
    color: #fbbf24 !important;
    font-size: 16px;
    font-family: 'Courier New', monospace;
    line-height: 1.6;
    text-shadow: 0 0 2px #D9770644;
  }

  code {
    background: #D9770618;
    color: #fbbf24;
    padding: 2px 8px;
    border-radius: 3px;
    font-size: 0.9em;
    border: 1px solid #D9770633;
    font-family: 'Courier New', monospace;
  }

  /* ===== 表格 - 史诗金边 ===== */
  table {
    width: 100%;
    border-collapse: collapse;
    margin: 14px 0;
    border: 1px solid #D9770644;
    position: relative;
    z-index: 2;
  }

  th {
    background: linear-gradient(180deg, #1a1410 0%, #0a0a0a 100%);
    color: #fbbf24;
    padding: 12px 18px;
    font-size: 22px;
    font-weight: 700;
    letter-spacing: 3px;
    text-shadow: 0 1px 2px #000000aa;
    border: 1px solid #D9770666;
    font-family: 'Cinzel', serif;
  }

  td {
    border: 1px solid #D9770633;
    padding: 10px 18px;
    color: #c8b896;
    font-size: 22px;
    background: #0d0a07;
  }

  tr:nth-child(even) td {
    background: #110d08;
  }

  /* ===== 引用块 ===== */
  blockquote {
    border-left: 4px solid #D97706;
    background: #D977060a;
    padding: 14px 22px;
    margin: 14px 0;
    color: #d8c8a8;
    font-style: italic;
    position: relative;
    z-index: 2;
  }

  /* ===== 列表 ===== */
  ul li::marker, ol li::marker {
    color: #D97706;
  }

  /* ===== 链接 ===== */
  a {
    color: #fbbf24;
    text-decoration: none;
    border-bottom: 1px solid #D9770644;
    transition: 0.2s;
  }
  a:hover {
    color: #fde68a;
    border-bottom-color: #fbbf24;
  }

  /* ===== 分隔线 - 金色渐变 ===== */
  hr {
    border: none;
    height: 2px;
    background: linear-gradient(90deg, transparent, #D97706, #fbbf24, #D97706, transparent);
    margin: 24px 0;
    box-shadow: 0 0 8px #D9770666;
    position: relative;
    z-index: 2;
  }

  /* ===== lead 页居中 ===== */
  section.lead {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    text-align: center;
  }

  /* ===== 章节序号圆圈 ===== */
  .chap {
    display: inline-block;
    width: 38px;
    height: 38px;
    line-height: 34px;
    text-align: center;
    border: 2px solid #D97706;
    border-radius: 50%;
    color: #fbbf24;
    font-family: 'Cinzel', serif;
    font-size: 18px;
    font-weight: 700;
    letter-spacing: 0;
    margin-right: 12px;
    text-shadow: 0 0 6px #D9770666;
    box-shadow:
      inset 0 0 8px #D9770622,
      0 0 6px #D9770644;
    background: radial-gradient(circle, #1a1410 0%, #0a0a0a 100%);
    vertical-align: middle;
  }

  /* ===== 二维码金边外框 ===== */
  .qr-frame {
    display: inline-block;
    padding: 10px;
    background: linear-gradient(135deg, #1a1410, #0a0a0a);
    border: 2px solid #D97706;
    border-radius: 6px;
    box-shadow:
      0 0 0 1px #fbbf2444,
      inset 0 0 12px #D9770622,
      0 4px 16px #000000aa;
    position: relative;
    z-index: 2;
  }

  .qr-frame::before,
  .qr-frame::after {
    content: '';
    position: absolute;
    width: 14px;
    height: 14px;
    border: 2px solid #fbbf24;
  }
  .qr-frame::before {
    top: -3px;
    left: -3px;
    border-right: none;
    border-bottom: none;
  }
  .qr-frame::after {
    bottom: -3px;
    right: -3px;
    border-left: none;
    border-top: none;
  }

  /* ===== 内容容器 - 确保在装饰条之上 ===== */
  section > * {
    position: relative;
    z-index: 2;
  }
---

<!-- _class: lead -->

# <span style="color:#D97706">AiDisk</span>

<span style="color:#fbbf24;font-size:30px;letter-spacing:8px;text-shadow:0 2px 4px #000000aa">全栈 AI 云盘系统</span>
<span style="color:#7a6a4a;font-size:18px;letter-spacing:4px">— 30 分钟技术分享 —</span>

<br>

<span style="color:#fbbf24;font-size:18px;letter-spacing:2px">阿七</span>
<span style="color:#4a3a2a">│</span>
<span style="color:#D97706;font-size:18px;letter-spacing:2px">CloudWeGo 社区</span>

<br>

<div class="qr-frame">
<img src="https://api.qrserver.com/v1/create-qr-code/?size=120x120&data=https://github.com/aqi-qihuan/AiDisk" width="120" height="120" alt="QR"/>
</div>

<span style="color:#D97706;font-size:13px;letter-spacing:1px">github.com/aqi-qihuan/AiDisk</span>
<span style="color:#fbbf24;font-size:13px;letter-spacing:1px">pan.aqi125.cn</span>

---

## <span style="color:#fbbf24">今 日 议 程</span>

<br>

| | 章节 | 时长 |
|:-:|-------|-------|
| <span class="chap">I</span> | 开场 & 自我介绍 | 2 min |
| <span class="chap">II</span> | 为什么做 AiDisk | 3 min |
| <span class="chap">III</span> | 技术架构概览 | 5 min |
| <span class="chap">IV</span> | **现场 Demo** | 7 min |
| <span class="chap">V</span> | 深入技术亮点（含 Eino） | 7 min |
| <span class="chap">VI</span> | 开源 & 社区 | 2 min |
| <span class="chap">VII</span> | 总结 & Q&A | 4 min |

<br>

<hr>

<span style="color:#D9770688;font-size:14px;letter-spacing:3px">— 史诗级全栈实践分享 —</span>

---

## <span style="color:#D97706">为 何 而 来</span>

<br>

```
  ╔══════════════════════════════════════════════╗
  ║  [ I ]  网盘收费限速                          ║
  ║  百度网盘不充会员 30KB/s                     ║
  ║  文件在别人的服务器上 → 隐私？              ║
  ╠══════════════════════════════════════════════╣
  ║  [ II ]  AI 与云盘割裂                       ║
  ║  文件在云盘，AI 在另一个网站                ║
  ║  下载 → 上传 → 下载，绕三圈               ║
  ╠══════════════════════════════════════════════╣
  ║  [ III ]  练手 Go 语言的完整项目             ║
  ║  学一门语言最好的方式：                     ║
  ║  用它从 0 到 1 做完一个能上线的项目      ║
  ╚══════════════════════════════════════════════╝
```

---

## <span style="color:#fbbf24">技 术 架 构</span>

<br>

```
  ┌─────────────────────────────────────────┐
  │         AQICloud-Web (Vue 3)            │
  │           pan.aqi125.cn                  │
  └─────────┬──────────┬────────────────────┘
             ▼ JWT        ▼ JWT
  ┌──────────────┐  ┌──────────────────────────┐
  │ AQICloud-Agent│  │  AQICloud-AI (Eino)      │
  │ Gin + MySQL   │  │  ChatModel / Agent /      │
  │ MinIO + Redis │  │  Callback / Stream        │
  └─────┬────────┘  └─────┬───────────────────┘
         │                    │
         └────────┬─────────┘
                  ▼
     MySQL  Redis  MinIO  Elasticsearch
```

---

## <span style="color:#D97706">项 目 规 模</span>

<br>

| 指标 | 数据 |
|------|------|
| Go 源文件 | **51 个** |
| Go 代码行数 | **~7800 行** |
| API 端点 | **50+** |
| 前端组件 | **93 个** (Vue 3 + TS) |
| 部署方式 | Docker Compose |
| 服务器 | **4 核 4GB** 轻量云 |
| 运行时内存 | **< 2GB** |
| CPU 空闲 | **1-2%** |

<hr>

<span style="color:#D9770688;font-size:14px;letter-spacing:2px">▲ 4GB 内存跑 10+ 容器，资源利用极致</span>

---

<!-- _class: lead -->

# <span style="color:#fbbf24">Demo 时 间</span>

<br>

<span style="color:#D97706;font-size:26px;letter-spacing:4px">pan.aqi125.cn</span>

<span style="color:#7a6a4a;font-size:15px;letter-spacing:2px">打开浏览器，现场体验</span>

<br>

<div class="qr-frame">
<img src="https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=https://pan.aqi125.cn" width="150" height="150" alt="QR"/>
</div>

---

## <span style="color:#D97706">Demo 流 程</span>

<br>

```
  ①  登录注册    (1 min)   JWT 认证 · Token 有效期 7 天
       ↓
  ②  文件上传    (1 min)   分块上传 · 5MB/块 · 秒传(SHA-256)
       ↓
  ③  文件分享    (1 min)   有效期 + 提取码 · 生成分享链接
       ↓
  ④  AI 聊天    (2 min)   自然语言查网盘 · Agent 查 MySQL
       ↓
  ⑤  文档分析    (1 min)   URL/PDF · 分块 · LLM 摘要
       ↓
  ⑥  模型切换    (1 min)   Ollama ↔ DashScope · 无需重启
```

---

## <span style="color:#fbbf24">技 术 亮 点 Ⅰ：分 块 上 传</span>

<br>

```
  前端                        后端 (Agent)
    │                          │
    ├── 算文件哈希 ──────────→ 查重？──yes──→ 返回成功（秒传）
    │                          │
    ├── 初始化上传 ──────────→ 返回 fileID
    │                          │
    ├── 上传块 1 ───────────→ Redis 记录 block:1 = done
    ├── 上传块 2 ───────────→ Redis 记录 block:2 = done
    │       ...                │
    ├── 上传块 N ───────────→ Redis 记录 block:N = done
    │                          │
    ├── 合并请求 ───────────→ MinIO 合并所有块 → 存 MySQL 元数据
```

<span style="color:#D9770688;font-size:14px">▲ 5MB/块 · Redis 记录进度 · 断点续传 · 秒传去重</span>

---

## <span style="color:#D97706">技 术 亮 点 Ⅱ：Agent 池 & Token 追踪</span>

<br>

**Agent 池**

三个独立 Agent：`Chat` · `Doc` · `Pan`
复用 LLM 客户端实例

<span style="color:#fbbf24;font-size:20px">平均响应时间 ↓ 30% · 内存峰值 ↓ 40%</span>

<br>

**Token 追踪** <span style="color:#D9770688;font-size:14px">（Eino Callback 实现）</span>

```
  LLM 调用 → OnStart 埋点 → 调用完成 → OnEnd 统计 token
                  ↓
          写入 Redis（按用户/按小时）
                  ↓
          前端实时展示配额
```

---

<!-- _class: lead -->

# <span style="color:#fbbf24">Eino 生 产 级 实 践</span>

<br>

<span style="color:#d8c8a8;font-size:22px;letter-spacing:2px">

**AiDisk = Eino 在生产环境的完整实践案例**
<br>
<span style="color:#D97706">非 demo · 非 example · 真实上线 · 真实用户</span>

</span>

<br>

<div class="qr-frame">
<img src="https://api.qrserver.com/v1/create-qr-code/?size=140x140&data=https://github.com/aqi-qihuan/AiDisk" width="140" height="140" alt="QR"/>
</div>

<span style="color:#fbbf24;font-family:monospace;font-size:15px">github.com/aqi-qihuan/AiDisk</span>
<span style="color:#D97706;font-family:monospace;font-size:15px">pan.aqi125.cn</span>

---

## <span style="color:#D97706">Eino 四 层 用 法</span>

<br>

| 层 | Eino 接口 | AiDisk 里的实现 |
|:-:|-----------|-------------------|
| <span class="chap">I</span> | **ChatModel** | 多 LLM 后端切换 · Ollama ↔ DashScope · 运行时切换 · 无需重启 |
| <span class="chap">II</span> | **Agent** | 三个独立 Agent · Chat / Doc / Pan · 各自系统提示词 + 工具集 |
| <span class="chap">III</span> | **Callback** | Token 统计埋点 · 零侵入 · 业务代码不用改 |
| <span class="chap">IV</span> | **Stream** | SSE 流式输出 · 前端逐字渲染 · Eino 封装 chunk 处理 |

<br>

```go
// 切换模型：一行代码
chatModel, _ := einoagent.GetChatModelWithProvider("dashscope")
```

---

## <span style="color:#fbbf24">生 产 踩 坑 记 录</span>

<br>

**坑 Ⅰ：长对话上下文爆炸**

<span style="color:#c8b896;font-size:20px">现象：> 20 轮对话，token 直接爆掉</span>
<span style="color:#fbbf24;font-size:20px">解决：`eino.TrimMessages` + 自写摘要机制</span>

<br>

**坑 Ⅱ：ChatModel 并发不安全**

<span style="color:#c8b896;font-size:20px">现象：多 goroutine 调同一实例 → 数据竞争</span>
<span style="color:#fbbf24;font-size:20px">解决：Agent 池 + 每个实例加锁 / 从池取不同实例</span>

<br>

**坑 Ⅲ：Stream 不 Close 会泄露连接**

<span style="color:#c8b896;font-size:20px">现象：并发一高，文件句柄耗尽</span>
<span style="color:#fbbf24;font-size:20px">解决：`defer stream.Close()` 必须写 + panic recovery</span>

---

## <span style="color:#D97706">开 源 & 社 区</span>

<br>

<span style="color:#fbbf24;font-size:22px">▶ GitHub</span>
<span style="color:#c8b896;font-size:18px">github.com/aqi-qihuan/AiDisk</span>

README 包含：完整安装步骤 · API 文档 · 配置说明 · 贡献指南

<br>

<span style="color:#fbbf24;font-size:22px">▶ 线上 Demo</span>
<span style="color:#c8b896;font-size:18px">pan.aqi125.cn · 扫码可体验注册/上传/AI聊天</span>

<br>

<span style="color:#D97706;font-size:20px">做开源的三个收获</span>

<span class="chap">I</span> 逼自己写出更好的代码

<span class="chap">II</span> 建立个人品牌

<span class="chap">III</span> 认识志同道合的人

---

<!-- _class: lead -->

# <span style="color:#D97706">总 结</span>

<br>

<span style="color:#c8b896;font-size:24px;text-align:left;display:block;margin:0 auto;width:80%">

**AiDisk = 一个网盘 + 一套 AI 能力 + 一套完整工程实践**

- 集成字节 **Eino** 框架的生产级项目（非 demo）
- 代码开源 · 线上可访问 · 踩坑记录在注释里
- 适合作为：Eino 实践参考 / 全栈学习素材 / Go 练手项目

</span>

<br>

<div class="qr-frame">
<img src="https://api.qrserver.com/v1/create-qr-code/?size=130x130&data=https://github.com/aqi-qihuan/AiDisk" width="130" height="130" alt="QR"/>
</div>

<span style="color:#fbbf24;font-size:18px;letter-spacing:2px">github.com/aqi-qihuan/AiDisk</span>
<span style="color:#D97706;font-size:18px;letter-spacing:2px">pan.aqi125.cn</span>

<br>

<span style="color:#fbbf24;font-size:26px;letter-spacing:6px">Q & A</span>
