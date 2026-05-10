package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/agents"
	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/core"
	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载 .env 文件（必须在 GetConfig 之前，否则 sync.Once 会缓存空值）
	loadEnvFile()

	// 初始化配置
	cfg := core.GetConfig()

	// 初始化 Token 追踪器
	core.GetTokenTracker()

	// 初始化 MySQL（对话记录持久化）
	core.GetDB()
	log.Println("✅ MySQL 连接成功")

	// 初始化全局 Agent 池（复用 LLM 客户端，减少每次请求开销）
	if err := agents.InitAgentPool(context.Background()); err != nil {
		log.Printf("Agent 池初始化失败（将降级为按需创建）: %v", err)
	}

	// 设置 Gin 模式
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// 注册路由
	handlers.RegisterRoutes(r)

	log.Printf("%s 启动成功，监听地址: %s", cfg.AppName, cfg.ListenAddr)
	if err := r.Run(cfg.ListenAddr); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}

// loadEnvFile 简易 .env 加载（对标 Python pydantic-settings 的 env_file）
func loadEnvFile() {
	data, err := os.ReadFile(".env")
	if err != nil {
		return // 没有 .env 文件，使用默认值
	}

	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if idx := strings.Index(line, "="); idx > 0 {
			key := strings.TrimSpace(line[:idx])
			val := strings.TrimSpace(line[idx+1:])
			os.Setenv(key, val)
		}
	}
}
