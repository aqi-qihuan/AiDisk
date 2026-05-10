package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/agents"
	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/core"
	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/models"
	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/services"
	"github.com/cloudwego/eino/schema"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			token = c.GetHeader("Authorization")
		}
		log.Printf("[AuthMiddleware] token=%q", token)
		if token == "" {
			c.AbortWithStatusJSON(http.StatusOK, models.Error("Token不能为空", -1))
			return
		}

		claims, err := core.ParseJWT(token)
		if err != nil {
			log.Printf("[AuthMiddleware] ParseJWT error: %v", err)
			c.AbortWithStatusJSON(http.StatusOK, models.Error(err.Error(), -1))
			return
		}

		c.Set("account_id", claims.AccountID)
		c.Set("username", claims.Username)
		c.Next()
	}
}

// RegisterRoutes 注册所有路由
func RegisterRoutes(r *gin.Engine) {
	// 全局 CORS 中间件
	r.Use(core.CorsMiddleware())
	// 根路径
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":          "欢迎使用AI智能体中心API",
			"version":          "1.0.0",
			"available_agents": []string{"chat", "doc", "pan"},
		})
	})

	// ===== Chat 路由 =====
	chat := r.Group("/api/chat")
	{
		chat.POST("/stream", AuthMiddleware(), chatStream)
		chat.GET("/providers", getProviders)
		chat.POST("/switch-provider", AuthMiddleware(), switchProvider)
		chat.GET("/history", AuthMiddleware(), getChatHistory)
		chat.DELETE("/history", AuthMiddleware(), clearChatHistory)
		chat.GET("/token-usage", AuthMiddleware(), getTokenUsage)
		chat.GET("/token-usage/global", AuthMiddleware(), getGlobalTokenUsage)
	}

	// ===== Document 路由 =====
	doc := r.Group("/api/document")
	{
		doc.POST("/stream", docStream)
		doc.GET("/providers", getDocProviders)
	}

	// ===== Pan 路由 =====
	pan := r.Group("/api/pan")
	{
		pan.POST("/query", AuthMiddleware(), panQuery)
		pan.GET("/providers", getPanProviders)
	}
}

// ===== Chat Handler =====

func chatStream(c *gin.Context) {
	var req models.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, models.Error("请求参数错误", -1))
		return
	}

	accountID, _ := c.Get("account_id")
	provider := req.Provider
	if provider == "" {
		provider = core.GetConfig().LLMProvider
	}

	chatSvc := services.GetChatService(provider)
	summary := chatSvc.GenerateSummary(c.Request.Context(), fmt.Sprintf("%d", accountID))

	// 构建带摘要的系统提示
	systemPrompt := fmt.Sprintf("%s\n\n对话摘要: %s", agents.ChatAgentSystemPrompt(), summary)

	// 构建历史消息
	var history []*schema.Message
	for _, h := range req.History {
		if h.Role == "user" {
			history = append(history, &schema.Message{Role: schema.User, Content: h.Content})
		} else {
			history = append(history, &schema.Message{Role: schema.Assistant, Content: h.Content})
		}
	}

	// 使用 AgentPool 复用实例
	pool := agents.GetAgentPool()
	var agent *agents.ChatAgent
	var err error
	if pool != nil {
		agent, err = pool.GetChat(c.Request.Context(), provider)
	} else {
		agent, err = agents.NewChatAgent(c.Request.Context(), provider)
	}
	if err != nil {
		c.JSON(http.StatusOK, models.Error(fmt.Sprintf("创建 Agent 失败: %v", err), 500))
		return
	}

	messages := agents.BuildMessages(systemPrompt, history, req.Message)

	// 动态上下文裁剪，避免超长对话浪费 token
	if maxTokens := core.GetConfig().MaxContextTokens; maxTokens > 0 {
		messages = agents.TrimMessages(messages, maxTokens)
	}

	// 流式输出
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Flush()

	stream, err := agent.Stream(c.Request.Context(), messages)
	if err != nil {
		resp := models.Error(fmt.Sprintf("流式调用失败: %v", err), 500)
		data, _ := json.Marshal(resp)
		fmt.Fprintf(c.Writer, "data: %s\n\n", data)
		return
	}

	var fullResponse strings.Builder
	for {
		chunks, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		for _, chunk := range chunks {
			if chunk.Content != "" {
				fullResponse.WriteString(chunk.Content)
				resp := models.StreamData(chunk.Content)
				data, _ := json.Marshal(resp)
				fmt.Fprintf(c.Writer, "data: %s\n\n", data)
				c.Writer.Flush()
			}
		}
	}
	stream.Close()

	// 保存对话历史（使用独立 context，避免请求结束后 Redis 写入失败）
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := chatSvc.SaveChatMessages(ctx, fmt.Sprintf("%d", accountID), req.Message, fullResponse.String()); err != nil {
			log.Printf("[SaveChatMessages] error: %v", err)
		}
	}()
}

func getProviders(c *gin.Context) {
	cfg := core.GetConfig()
	c.JSON(http.StatusOK, gin.H{
		"providers": []string{"ollama", "openai_compatible"},
		"current":   cfg.LLMProvider,
		"model":     cfg.LLMModelName,
	})
}

func switchProvider(c *gin.Context) {
	var req struct {
		Provider string `json:"provider"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, models.Error("请求参数错误", -1))
		return
	}

	if req.Provider != "ollama" && req.Provider != "openai_compatible" {
		c.JSON(http.StatusOK, models.Error("无效的模型提供商，支持: ollama, openai_compatible", -1))
		return
	}

	core.GetConfig().SetLLMProvider(req.Provider)
	log.Printf("模型提供商已切换为: %s", req.Provider)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("模型提供商已切换为: %s", req.Provider),
		"success": true,
	})
}

func getChatHistory(c *gin.Context) {
	accountID, _ := c.Get("account_id")
	chatSvc := services.GetChatService()
	messages, err := chatSvc.GetChatHistory(c.Request.Context(), fmt.Sprintf("%d", accountID))
	if err != nil {
		c.JSON(http.StatusOK, models.Error("获取聊天历史失败", -1))
		return
	}
	c.JSON(http.StatusOK, gin.H{"history": messages})
}

func clearChatHistory(c *gin.Context) {
	accountID, _ := c.Get("account_id")
	chatSvc := services.GetChatService()
	if err := chatSvc.ClearChatHistory(c.Request.Context(), fmt.Sprintf("%d", accountID)); err != nil {
		c.JSON(http.StatusOK, models.Error("清空聊天历史失败", -1))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "聊天历史已清空", "success": true})
}

func getTokenUsage(c *gin.Context) {
	accountID, _ := c.Get("account_id")
	stats := core.GetTokenTracker().GetUserStats(fmt.Sprintf("%d", accountID))
	c.JSON(http.StatusOK, models.Success(stats))
}

func getGlobalTokenUsage(c *gin.Context) {
	stats := core.GetTokenTracker().GetGlobalStats()
	c.JSON(http.StatusOK, models.Success(stats))
}

// ===== Document Handler =====

type docChunkResult struct {
	index int
	text  string
}

func docStream(c *gin.Context) {
	var req models.DocumentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, models.Error("请求参数错误", -1))
		return
	}

	ds := &services.DocumentService{}
	title, content, _, err := ds.FetchDocument(req.URL)
	if err != nil {
		c.JSON(http.StatusOK, models.Error(fmt.Sprintf("获取文档失败: %v", err), 500))
		return
	}

	chunks := ds.ChunkContent(content, 1000000)

	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Flush()

	// 单 chunk 直接流式处理
	if len(chunks) <= 1 {
		streamDocChunk(c, req, title, chunks[0])
		fmt.Fprintf(c.Writer, "data: [DONE]\n\n")
		return
	}

	// 多 chunk 并行 LLM 调用，按序合并流式输出
	resultCh := make(chan docChunkResult, len(chunks))

	for i, chunk := range chunks {
		go func(idx int, content string) {
			input := services.BuildDocInput(req, title, content)
			messages := []*schema.Message{
				{Role: schema.System, Content: agents.DocAgentSystemPrompt()},
				{Role: schema.User, Content: input},
			}

			pool := agents.GetAgentPool()
			var agent *agents.DocAgent
			var err error
			if pool != nil {
				agent, err = pool.GetDoc(context.Background(), core.GetConfig().LLMProvider)
			} else {
				agent, err = agents.NewDocAgent(context.Background(), core.GetConfig().LLMProvider)
			}
			if err != nil {
				resultCh <- docChunkResult{index: idx, text: fmt.Sprintf("[分块 %d 处理失败: %v]", idx, err)}
				return
			}

			// 非流式调用，获取完整响应
			stream, err := agent.Stream(context.Background(), messages)
			if err != nil {
				resultCh <- docChunkResult{index: idx, text: fmt.Sprintf("[分块 %d 处理失败: %v]", idx, err)}
				return
			}

			var fullText strings.Builder
			for {
				msgs, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					break
				}
				for _, msg := range msgs {
					fullText.WriteString(msg.Content)
				}
			}
			stream.Close()

			resultCh <- docChunkResult{index: idx, text: fullText.String()}
		}(i, chunk)
	}

	// 按序等待并流式输出
	for i := 0; i < len(chunks); i++ {
		r := <-resultCh
		if i > 0 {
			fmt.Fprintf(c.Writer, "data: %s\n\n", models.StreamData("\n\n--- 下一部分 ---\n\n"))
			c.Writer.Flush()
		}
		for _, ch := range r.text {
			fmt.Fprintf(c.Writer, "data: %s\n\n", models.StreamData(string(ch)))
		}
		c.Writer.Flush()
	}

	fmt.Fprintf(c.Writer, "data: [DONE]\n\n")
}

// streamDocChunk 单 chunk 流式处理
func streamDocChunk(c *gin.Context, req models.DocumentRequest, title, content string) {
	input := services.BuildDocInput(req, title, content)
	messages := []*schema.Message{
		{Role: schema.System, Content: agents.DocAgentSystemPrompt()},
		{Role: schema.User, Content: input},
	}

	pool := agents.GetAgentPool()
	var agent *agents.DocAgent
	var err error
	if pool != nil {
		agent, err = pool.GetDoc(c.Request.Context(), core.GetConfig().LLMProvider)
	} else {
		agent, err = agents.NewDocAgent(c.Request.Context(), core.GetConfig().LLMProvider)
	}
	if err != nil {
		resp := models.Error(fmt.Sprintf("创建 Agent 失败: %v", err), 500)
		data, _ := json.Marshal(resp)
		fmt.Fprintf(c.Writer, "data: %s\n\n", data)
		return
	}

	stream, err := agent.Stream(c.Request.Context(), messages)
	if err != nil {
		resp := models.Error(fmt.Sprintf("文档处理错误: %v", err), 500)
		data, _ := json.Marshal(resp)
		fmt.Fprintf(c.Writer, "data: %s\n\n", data)
		return
	}

	for {
		msgs, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		for _, msg := range msgs {
			if msg.Content != "" {
				data, _ := json.Marshal(models.StreamData(msg.Content))
				fmt.Fprintf(c.Writer, "data: %s\n\n", data)
				c.Writer.Flush()
			}
		}
	}
	stream.Close()
}

func getDocProviders(c *gin.Context) {
	cfg := core.GetConfig()
	c.JSON(http.StatusOK, gin.H{
		"providers": []string{"ollama", "openai_compatible"},
		"current":   cfg.LLMProvider,
		"model":     cfg.LLMModelName,
	})
}

// ===== Pan Handler =====

func panQuery(c *gin.Context) {
	var req models.PanQueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, models.Error("请求参数错误", -1))
		return
	}

	accountID, _ := c.Get("account_id")
	aid := accountID.(int)

	log.Printf("用户%d开始查询网盘: %s", aid, req.Query)

	// 关键词快速意图分类，避免不必要的 LLM 调用
	intention := agents.FastIntentParse(req.Query)

	panSvc := services.GetPanService()

	var result any
	var respType string

	switch strings.ToLower(intention) {
	case "storage", "storage_info":
		info, err := panSvc.QueryStorage(c.Request.Context(), aid)
		if err != nil {
			c.JSON(http.StatusOK, models.Error(fmt.Sprintf("查询失败: %v", err), -1))
			return
		}
		respType = "storage_info"
		result = info

	case "file_statistics", "statistics":
		count, _ := panSvc.QueryFileCount(c.Request.Context(), aid)
		totalSize, _ := panSvc.QueryTotalSize(c.Request.Context(), aid)
		types, _ := panSvc.QueryFileTypes(c.Request.Context(), aid)
		respType = "file_statistics"
		result = models.FileStatistics{
			TotalFiles: int(count),
			TotalSize:  totalSize,
			FileTypes:  types,
		}

	default: // file_list, search
		files, err := panSvc.QueryFiles(c.Request.Context(), aid, "", 50)
		if err != nil {
			c.JSON(http.StatusOK, models.Error(fmt.Sprintf("查询失败: %v", err), -1))
			return
		}
		respType = "file_list"
		result = files
	}

	c.JSON(http.StatusOK, models.Success(map[string]any{
		"type": respType,
		"data": result,
	}))
}

func getPanProviders(c *gin.Context) {
	cfg := core.GetConfig()
	c.JSON(http.StatusOK, gin.H{
		"providers": []string{"ollama", "openai_compatible"},
		"current":   cfg.LLMProvider,
		"model":     cfg.LLMModelName,
	})
}
