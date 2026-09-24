# AiDisk 全栈升级方案：Go 1.27.x + 中间件对齐

> 生成日期：2026-09-24
> 服务器：广州 Lighthouse `lhins-8w1w37al`（134.175.206.158）
> 方案基于实际扫描结果（本地代码 + 服务器在线验证），非猜测。

---

## 一、扫描结果

### 1.1 Go 工具链现状

| 项 | AqiCloud-Agent | AqiCloud-AI | 目标 |
|---|---|---|---|
| go.mod 声明 | `go 1.26.2` | `go 1.26` | **`go 1.27.1`** |
| Dockerfile 构建镜像 | `golang:1.26-alpine` | `golang:1.26-alpine` | `golang:1.27.1-alpine` |
| Dockerfile 运行镜像 | `alpine:latest`（漂移） | `alpine:latest`（漂移） | `alpine:3.22`（固定） |
| compose 运行镜像 | `alpine:3.20` | `alpine:3.20` | `alpine:3.22`（对齐 aurora-go） |

### 1.2 服务器中间件实际版本（2026-09-24 在线核实）

| 中间件 | 运行镜像 | 状态 | AiDisk 是否使用 |
|---|---|---|---|
| MySQL | `mysql:8.4.10` | ✅ 新版 | 是（两服务共用） |
| Redis | `redis:8.2.9-alpine` | ✅ 新版 | 是（AI 服务） |
| 对象存储 | `rustfs/rustfs:1.0.0`（**已替代 MinIO**） | ✅ | 是（Agent，经 `aurora-minio` 别名） |
| Nginx | `nginx:1.28.0-alpine` | ✅ 新版 | 是（反代入口） |
| RabbitMQ | `rabbitmq:4.3.6-management` | ✅ 新版 | 否 |
| Elasticsearch | `elasticsearch:9.5.3` | ✅ 新版 | 否 |

> **结论：服务器侧中间件已在 aurora-go 升级时全部到位（8.4.10 / 8.2.9 / 1.28 / RustFS 1.0.0）。AiDisk 侧没有中间件需要升级，需要做的是运行时镜像对齐 + 配置注释修正。**

### 1.3 Go 依赖现状 vs 最新（2026-09）

| 依赖 | Agent | AI | 最新 stable | 判定 |
|---|---|---|---|---|
| gin | v1.12.0 | v1.12.0 | v1.12.0 | ✅ 已最新 |
| gorm | v1.31.1 | v1.31.1 | v1.31.x | ✅ 基本最新 |
| gorm/mysql driver | v1.6.0 | v1.6.0 | v1.6.x | ✅ |
| jwt/v5 | v5.3.1 | v5.3.1 | v5.3.1 | ✅ |
| **eino** | — | **v0.8.13** | **v0.9.21** | 🔺 落后一个 minor |
| go-redis/v9 | — | v9.19.0 | v9.x | ✅ |
| go-openai | — | v1.41.2 | v1.4x | ✅ |
| aws-sdk-go-v2/s3 | v1.100.0 | — | 持续滚动 | 🔺 随 go get 更新 |

### 1.4 扫描中发现的隐患（按优先级）

| 级别 | 问题 | 说明 |
|---|---|---|
| 🟡 P1 | **gin 传递依赖 CVE** | Agent go.mod 锁定 `quic-go v0.59.0`（CVE-2026-40898 HTTP/3 QPACK 内存耗尽，v0.59.1 修复）和 `golang.org/x/net v0.51.0`（GO-2026-4918 HTTP/2 无限循环，v0.53.0 修复）。govulncheck 判定可从 gin 的服务路径触达 |
| 🟡 P2 | **eino v0.8.13 → v0.9.21** | minor 升级，ADK/checkpoint/stream 有行为修正；v0.10.0 尚在 alpha，**不要上** |
| 🟢 P3 | 运行镜像漂移 | Dockerfile 用 `alpine:latest`（不可复现），compose 锁 3.20（落后 aurora-go 的 3.22） |
| 🟢 P4 | compose 注释误导 | `MINIO_ENDPOINT=aurora-minio:9000` 实际由 RustFS 承载（网络别名验证通过，健康端点返回 rustfs v1.0.0）——功能正常，但注释会误导后人 |
| 🟢 P5 | Makefile `deps` 目标路径过期 | 引用 `backend-agent/backend-ai/frontend`，实际目录是 `AqiCloud-Agent/AqiCloud-AI/AqiCloud-Web` |
| 🟢 P6 | JWT 变量名不统一 | compose 中 AI 用 `JWT_SECRET_KEY`，Agent 用 `JWT_SECRET`；代码已做 fallback 兼容，建议顺手统一为 `JWT_SECRET` |

---

## 二、Go 1.27 升级要点

- **Go 1.27.0**：2026-08-19 发布；**1.27.1**：2026-09-01（当前最新 patch）
- 1.26 线仍在维护（最新 1.26.8），本次升级非安全强制，属常规对齐
- 兼容性：1.26 → 1.27 遵循 Go 1 兼容性承诺，无 breaking change 预期
- 亮点（可选启用，默认关闭）：`encoding/json/v2`（GOEXPERIMENT=jsonv2）、generic methods、后量子 crypto
- 升级动作只有 4 处：两个 go.mod、两个 Dockerfile，其余靠 `go mod tidy` 自动收敛

---

## 三、实施方案（4 阶段）

### Phase 0：基线与备份（10 min）

```bash
# 本地
git checkout -b upgrade/go1.27-deps
git tag pre-upgrade-20260924 && git push origin pre-upgrade-20260924

# 服务器快照（控制台或 API）：对 lhins-8w1w37al 创建快照
# 快照是回滚底线，务必等快照完成再动服务器
```

### Phase 1：Go 工具链 + CVE 修复（30 min）

```bash
# 1. go.mod 版本声明
cd AqiCloud-Agent && go mod edit -go=1.27.1
cd ../AqiCloud-AI  && go mod edit -go=1.27.1

# 2. 修复 gin 传递依赖 CVE（显式提升间接依赖）
cd AqiCloud-Agent
go get github.com/quic-go/quic-go@v0.59.1 golang.org/x/net@v0.53.0
go mod tidy && go build ./... && go vet ./...

cd ../AqiCloud-AI
go get golang.org/x/net@v0.53.0
go mod tidy && go build ./... && go vet ./...

# 3. 本地全量测试
go test ./... -count=1
```

**验证点**：`go version -m bin/*` 确认构建产物为 go1.27.1；`govulncheck ./...` 无 HIGH 以上告警。

### Phase 2：eino minor 升级（30 min，仅 AI 服务）

```bash
cd AqiCloud-AI
go get github.com/cloudwego/eino@v0.9.21
go mod tidy
go build ./... && go test ./...
```

**回归重点**（eino 改动集中在 ADK/checkpoint/stream）：
1. AI 聊天流式输出（SSE 逐字渲染）
2. 文档分析（多 chunk 并行 + Stream Close）
3. Token 统计埋点（Callback 路径）
4. 模型切换（Ollama ↔ DashScope）

若 v0.9 引入 API 变更导致编译失败：优先适配；改动超过 100 行则降级回 v0.8.13 并记录到 TODO（不阻塞本次升级）。

### Phase 3：镜像与配置对齐（20 min）

**两个 Dockerfile 统一改法**（Agent 与 AI 相同模式）：

```dockerfile
FROM golang:1.27.1-alpine AS builder
# ... 原有构建步骤不变 ...
FROM alpine:3.22          # ← latest 改为固定版本
RUN apk --no-cache add ca-certificates tzdata
# ... 其余不变 ...
```

**docker-compose.yml 两处微调**：

```yaml
# agent/ai 两个服务：
image: alpine:3.22        # ← 3.20 升 3.22，对齐 aurora-go

# agent 的 MINIO 区块加注释（endpoint 不改，别名已验证）：
MINIO_ENDPOINT: "aurora-minio:9000"   # 注意：实际由 aurora-rustfs (RustFS 1.0.0) 承载，兼容 S3

# ai 的 JWT 统一（代码已有 fallback，改了更干净）：
JWT_SECRET: "a4i.icu..."              # ← JWT_SECRET_KEY 改名
```

**Makefile 修复**：

```makefile
deps:
	cd AqiCloud-Agent && go mod tidy
	cd AqiCloud-AI && go mod tidy
	cd AqiCloud-Web && npm install
```

### Phase 4：服务器部署与回归（40 min）

```bash
# 1. 服务器上构建（或在本地交叉编译后上传二进制）
ssh 服务器 "cd /opt/aipan-llm && git pull && make build-all"

# 2. 滚动重启（先 AI 后 Agent，或逐个）
docker compose up -d --force-recreate ai
docker compose up -d --force-recreate agent

# 3. 健康检查
docker ps --format '{{.Names}}\t{{.Status}}' | grep aqicloud
curl -fsS http://localhost:9090/health
curl -fsS http://localhost:8000/

# 4. 功能回归（线上 https://pan.aqi125.cn）
#   □ 登录 → 文件列表
#   □ 上传一个小文件 → 秒传/分块
#   □ 分享链接 + 提取码
#   □ AI 聊天（流式）→ Token 计数增长
#   □ 文档分析 → 摘要返回
#   □ 运行时切换 Ollama → DashScope
```

**验证点**：`docker exec aqicloud-agent /app/AqiCloud-Agent --version`（若有）；日志无 panic；监控 CPU/内存无异常跳变。

---

## 四、回滚策略

| 阶段 | 回滚方式 |
|---|---|
| Phase 1-3（本地） | `git checkout pre-upgrade-20260924`，go.mod/Dockerfile 原样恢复 |
| Phase 4（服务器） | 服务器快照回滚（最快但整机回退）；或仅把旧二进制文件重新挂载 + `docker compose restart`（分钟级，推荐） |

---

## 五、执行顺序总览

```
P0 备份快照 ──→ P1 Go 1.27.1 + CVE 修复 ──→ P2 eino v0.9.21 ──→ P3 镜像 3.22 + 配置修正 ──→ P4 部署回归
   (10min)        (30min)                      (30min)             (20min)                    (40min)
                                                                                    总计 ≈ 2.2h
```

**明确不做的事**：
- ❌ 不升级 eino 到 v0.10.0-alpha（未 stable）
- ❌ 不动服务器上 aurora-go 的中间件容器（已是目标版本，AiDisk 是共享方不是所有方）
- ❌ 不改 `MINIO_ENDPOINT` 指向（RustFS 别名已验证，改了反而多余）
- ❌ 不启用 GOEXPERIMENT=jsonv2（等 1.28 默认化再评估）
