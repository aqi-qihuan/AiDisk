package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aqi/AqiCloud-Ai/internal/agents"
	"github.com/aqi/AqiCloud-Ai/internal/core"
	"github.com/aqi/AqiCloud-Ai/internal/handlers"
	"github.com/aqi/AqiCloud-Ai/internal/util"
	"github.com/gin-gonic/gin"
)

func main() {
	start := time.Now()

	// 加载 .env 文件（必须在 GetConfig 之前，否则 sync.Once 会缓存空值）
	loadEnvFile()

	// 初始化配置
	cfg := core.GetConfig()

	// 初始化 Token 追踪器
	core.GetTokenTracker()

	// Health checks
	checks := []util.CheckResult{
		pingMySQL(cfg),
		pingRedis(cfg),
		pingLLM(cfg),
	}

	// 初始化 MySQL（对话记录持久化）
	core.GetDB()

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

	elapsed := time.Since(start).Milliseconds()

	port := cfg.ListenAddr
	if port[0] == ':' {
		port = port[1:]
	}
	util.PrintStartupBanner(cfg.AppName, port, elapsed, checks)

	log.Printf("%s 启动成功，监听地址: %s", cfg.AppName, cfg.ListenAddr)
	if err := r.Run(cfg.ListenAddr); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}

// loadEnvFile 简易 .env 加载（对标 Python pydantic-settings 的 env_file）
func pingMySQL(cfg *core.Config) util.CheckResult {
	start := time.Now()
	dsn := cfg.MySQLDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return util.CheckResult{Name: "MySQL", OK: false, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
	}
	defer db.Close()
	err = db.Ping()
	return util.CheckResult{Name: "MySQL", OK: err == nil, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
}

func pingRedis(cfg *core.Config) util.CheckResult {
	start := time.Now()
	addr := cfg.RedisAddr()
	conn, err := (&net.Dialer{Timeout: 3 * time.Second}).Dial("tcp", addr)
	if err != nil {
		return util.CheckResult{Name: "Redis", OK: false, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
	}
	conn.Close()
	return util.CheckResult{Name: "Redis", OK: true, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
}

func pingLLM(cfg *core.Config) util.CheckResult {
	start := time.Now()

	if cfg.LLMProvider == "ollama" {
		if cfg.LLMOllamaBaseURL == "" {
			return util.CheckResult{Name: "Ollama", Skip: true}
		}
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		req, _ := http.NewRequestWithContext(ctx, "GET", cfg.LLMOllamaBaseURL, nil)
		resp, err := http.DefaultClient.Do(req)
		if err != nil || resp.StatusCode > 399 {
			return util.CheckResult{Name: "Ollama", OK: false, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
		}
		defer resp.Body.Close()
		return util.CheckResult{Name: "Ollama", OK: true, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
	}

	if cfg.LLMBaseURL == "" || cfg.LLMAPIKey == "" {
		return util.CheckResult{Name: "LLM(API)", Skip: true}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, "GET", cfg.LLMBaseURL+"/models", nil)
	req.Header.Set("Authorization", "Bearer "+cfg.LLMAPIKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode > 399 {
		return util.CheckResult{Name: "LLM(API)", OK: false, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
	}
	defer resp.Body.Close()
	return util.CheckResult{Name: "LLM(API)", OK: true, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
}

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
