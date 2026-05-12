package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/aqi/AqiCloud-Agent/docs"
	"github.com/aqi/AqiCloud-Agent/internal/config"
	"github.com/aqi/AqiCloud-Agent/internal/controller"
	"github.com/aqi/AqiCloud-Agent/internal/middleware"
	"github.com/aqi/AqiCloud-Agent/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           AqiCloud-AgentPan API
// @version         1.0
// @description     AI智能化云盘系统后端API，支持文件管理、分享、回收站、AI聊天等功能
// @termsOfService  https://www.aqi125.cn
// @contact.name    aqi
// @contact.email   2316364297@qq.com
// @contact.url     https://www.aqi125.cn
// @license.name    Apache 2.0
// @license.url     https://www.aqi125.cn
// @BasePath        /
// @securityDefinitions.apikey  Token
// @in              header
// @name            token
func main() {
	start := time.Now()

	// 加载 .env 文件（按优先级尝试: 当前目录 -> 可执行文件所在目录）
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠ .env 加载失败(当前目录): %v", err)
		// 回退: 尝试使用可执行文件所在目录
		if ex, exe := os.Executable(); exe == nil {
			dir := filepath.Dir(ex)
			if err := godotenv.Load(filepath.Join(dir, ".env")); err != nil {
				log.Printf("⚠ .env 加载失败(二进制目录): %v", err)
			}
		}
	} else {
		log.Println("✅ .env 加载成功")
	}
	cfg := config.GetConfig()
	util.InitSnowflake(1)
	config.AutoMigrate()

	// Health checks
	checks := []util.CheckResult{
		pingMySQL(cfg),
		pingMinIO(cfg),
		pingOllama(cfg),
		pingStream(cfg),
	}

	elapsed := time.Since(start).Milliseconds()

	// Setup Gin
	gin.SetMode(gin.ReleaseMode)
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	r.Use(middleware.CorsMiddleware())

	// Swagger UI
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Global auth middleware: public paths skip auth check but still get account_id if token present
	r.Use(func(ctx *gin.Context) {
		skip := middleware.ShouldSkipAuth(ctx.Request.URL.Path)
		if !skip {
			middleware.AuthMiddleware()(ctx)
			if ctx.IsAborted() {
				return
			}
		} else {
			// 公共路径：如果有 token 则尝试解析，设置 account_id
			token := ctx.GetHeader("token")
			if token == "" {
				token = ctx.Query("token")
			}
			if token != "" {
				claims, err := util.ParseLoginToken(token)
				if err == nil && claims.AccountID != 0 {
					ctx.Set("account_id", claims.AccountID)
					ctx.Set("username", claims.Username)
				}
			}
		}
	})

	// Register controllers
	controller.NewAccountController().Register(r.Group(""))
	controller.NewFileController().Register(r.Group(""))
	controller.NewShareController().Register(r.Group(""))
	controller.NewRecycleController().Register(r.Group(""))
	controller.NewChatController().Register(r.Group(""))

	// Add Swagger paths to public skip list
	middleware.AddPublicPath("/swagger/index.html")
	middleware.AddPublicPath("/swagger/doc.json")
	middleware.AddPublicPath("/swagger/any")

	// Print banner
	port := cfg.ListenAddr
	if port[0] == ':' {
		port = port[1:]
	}
	util.PrintStartupBanner(port, elapsed, checks)

	log.Printf("%s 启动成功，监听地址: %s", cfg.AppName, cfg.ListenAddr)
	if err := r.Run(cfg.ListenAddr); err != nil {
		log.Fatal("服务启动失败:", err)
	}
}

func pingMySQL(cfg *config.Config) util.CheckResult {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=3s",
		cfg.MySQLUser, cfg.MySQLPassword, cfg.MySQLHost, cfg.MySQLPort, cfg.MySQLDatabase)
	start := time.Now()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return util.CheckResult{Name: "MySQL", OK: false, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
	}
	defer db.Close()
	err = db.Ping()
	return util.CheckResult{Name: "MySQL", OK: err == nil, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
}

func pingMinIO(cfg *config.Config) util.CheckResult {
	addr := cfg.MinIOEndpoint
	if addr[0] == ':' {
		addr = "localhost" + addr
	}
	start := time.Now()
	conn, err := (&net.Dialer{Timeout: 3 * time.Second}).Dial("tcp", addr)
	if err != nil {
		return util.CheckResult{Name: "MinIO", OK: false, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
	}
	conn.Close()
	return util.CheckResult{Name: "MinIO", OK: true, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
}

func pingOllama(cfg *config.Config) util.CheckResult {
	if cfg.OllamaBaseURL == "" {
		return util.CheckResult{Name: "Ollama", Skip: true}
	}
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, "GET", cfg.OllamaBaseURL, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode > 399 {
		return util.CheckResult{Name: "Ollama", OK: false, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
	}
	defer resp.Body.Close()
	return util.CheckResult{Name: "Ollama", OK: true, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
}

func pingStream(cfg *config.Config) util.CheckResult {
	if cfg.StreamBaseURL == "" {
		return util.CheckResult{Name: "StreamAgent", Skip: true}
	}
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, "GET", cfg.StreamBaseURL, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode > 399 {
		return util.CheckResult{Name: "StreamAgent", OK: false, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
	}
	defer resp.Body.Close()
	return util.CheckResult{Name: "StreamAgent", OK: true, Latency: fmt.Sprintf("%dms", time.Since(start).Milliseconds())}
}
