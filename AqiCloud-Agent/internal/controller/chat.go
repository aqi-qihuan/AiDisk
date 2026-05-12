package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/aqi/AqiCloud-Agent/internal/config"
	"github.com/aqi/AqiCloud-Agent/internal/model"
	"github.com/aqi/AqiCloud-Agent/internal/service"
	"github.com/gin-gonic/gin"
)

type ChatController struct {
	svc *service.ChatService
}

func NewChatController() *ChatController {
	return &ChatController{svc: service.NewChatService()}
}

func (c *ChatController) Register(r *gin.RouterGroup) {
	group := r.Group("/api/chat")
	group.POST("/message", c.message)
	group.POST("/message/provider", c.messageWithProvider)
	group.POST("/stream", c.stream)

	// Additional AI endpoints (compatibility with Java version)
	r.POST("/v1/chat/completions", c.openAICompletions)
	r.GET("/ai_write", c.aiWrite)
	r.POST("/api/document/stream", c.documentStream)
	r.POST("/api/pan/query", c.panQuery)
}

// message AI聊天（默认提供商）
// @Summary      AI聊天（默认提供商）
// @Description  向默认AI提供商发送消息并获取响应
// @Tags         AI聊天
// @Accept       plain
// @Produce      json
// @Security     Token
// @Param        body  body      string  true  "聊天消息内容"
// @Success      200   {object}  model.JsonData
// @Router       /api/chat/message [post]
func (c *ChatController) message(ctx *gin.Context) {
	body, _ := io.ReadAll(ctx.Request.Body)
	msg := string(body)
	result, err := c.svc.SendMessage(ctx.Request.Context(), msg, "")
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), 500))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(result))
}

// messageWithProvider AI聊天（指定提供商）
// @Summary      AI聊天（指定提供商）
// @Description  向指定AI提供商发送消息并获取响应
// @Tags         AI聊天
// @Accept       plain
// @Produce      json
// @Security     Token
// @Param        body      body     string  true  "聊天消息内容"
// @Param        provider  query    string  true  "AI提供商 (dashscope 或 ollama)"
// @Success      200       {object}  model.JsonData
// @Router       /api/chat/message/provider [post]
func (c *ChatController) messageWithProvider(ctx *gin.Context) {
	provider := ctx.Query("provider")
	body, _ := io.ReadAll(ctx.Request.Body)
	msg := string(body)
	result, err := c.svc.SendMessage(ctx.Request.Context(), msg, provider)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error(err.Error(), 500))
		return
	}
	ctx.JSON(http.StatusOK, model.Success(result))
}

// stream AI聊天（SSE流式）
// @Summary      AI聊天（SSE流式）
// @Description  将请求代理到外部AI Agent服务，以SSE流式返回响应
// @Tags         AI聊天
// @Accept       json
// @Produce      text/event-stream
// @Security     Token
// @Param        body  body      string  true  "聊天消息内容"
// @Success      200   {string}  string  "SSE流式响应"
// @Router       /api/chat/stream [post]
func (c *ChatController) stream(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(http.StatusOK, model.Error("Token不能为空", model.CodeNotLogin))
		return
	}

	body, _ := io.ReadAll(ctx.Request.Body)

	cfg := config.GetConfig()
	if cfg.StreamBaseURL == "" {
		ctx.JSON(http.StatusOK, model.Error("SSE代理未配置", 500))
		return
	}

	// Forward request to external AI agent service
	targetURL := cfg.StreamBaseURL + cfg.StreamChatPath
	req, err := http.NewRequestWithContext(ctx.Request.Context(), "POST", targetURL, bytes.NewReader(body))
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("创建请求失败", 500))
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", token)
	req.Header.Set("Accept", "text/event-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("请求外部服务失败: "+err.Error(), 500))
		return
	}
	defer resp.Body.Close()

	// Stream SSE response back to client
	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")
	ctx.Writer.Flush()

	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			ctx.Writer.Write(buf[:n])
			ctx.Writer.Flush()
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			break
		}
	}
}

// openAICompletions OpenAI兼容接口（代理到DashScope）
// @Summary      OpenAI兼容接口
// @Description  代理到DashScope的OpenAI兼容接口
// @Tags         AI聊天
// @Accept       json
// @Produce      json
// @Param        body  body      string  true  "OpenAI格式请求体"
// @Success      200   {object}  string  "OpenAI格式响应"
// @Router       /v1/chat/completions [post]
func (c *ChatController) openAICompletions(ctx *gin.Context) {
	cfg := config.GetConfig()
	if cfg.DashScopeBase == "" {
		ctx.JSON(http.StatusOK, model.Error("DashScope未配置", 500))
		return
	}

	body, _ := io.ReadAll(ctx.Request.Body)

	targetURL := cfg.DashScopeBase + "/chat/completions"
	req, err := http.NewRequestWithContext(ctx.Request.Context(), "POST", targetURL, bytes.NewReader(body))
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("创建请求失败", 500))
		return
	}
	req.Header.Set("Content-Type", "application/json")
	if cfg.DashScopeAPIKey != "" {
		req.Header.Set("Authorization", "Bearer "+cfg.DashScopeAPIKey)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("请求DashScope失败: "+err.Error(), 500))
		return
	}
	defer resp.Body.Close()

	ctx.Writer.Header().Set("Content-Type", "application/json")
	io.Copy(ctx.Writer, resp.Body)
}

// aiWrite AI写作接口（SSE流式）
// @Summary      AI写作
// @Description  以SSE流式返回AI响应
// @Tags         AI聊天
// @Produce      text/event-stream
// @Param        query  query  string  true  "查询内容"
// @Success      200   {string}  string  "SSE流式响应"
// @Router       /ai_write [get]
func (c *ChatController) aiWrite(ctx *gin.Context) {
	query := ctx.Query("query")
	if query == "" {
		ctx.JSON(http.StatusOK, model.Error("缺少query参数", 500))
		return
	}

	body, _ := json.Marshal(map[string]string{"query": query})
	c.forwardSSE(ctx, body)
}

// documentStream AI文档流式接口
// @Summary      AI文档流式
// @Description  代理到AI Agent服务的文档分析接口
// @Tags         AI聊天
// @Accept       json
// @Produce      text/event-stream
// @Param        body  body      string  true  "文档分析请求"
// @Success      200   {string}  string  "SSE流式响应"
// @Router       /api/document/stream [post]
func (c *ChatController) documentStream(ctx *gin.Context) {
	body, _ := io.ReadAll(ctx.Request.Body)
	c.forwardSSE(ctx, body)
}

// panQuery AI网盘问答接口
// @Summary      AI网盘问答
// @Description  代理到AI Agent服务的网盘问答接口
// @Tags         AI聊天
// @Accept       json
// @Produce      text/event-stream
// @Param        body  body      string  true  "网盘查询请求"
// @Success      200   {string}  string  "SSE流式响应"
// @Router       /api/pan/query [post]
func (c *ChatController) panQuery(ctx *gin.Context) {
	body, _ := io.ReadAll(ctx.Request.Body)
	c.forwardSSE(ctx, body)
}

func (c *ChatController) forwardSSE(ctx *gin.Context, body []byte) {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(http.StatusOK, model.Error("Token不能为空", model.CodeNotLogin))
		return
	}

	cfg := config.GetConfig()
	if cfg.StreamBaseURL == "" {
		ctx.JSON(http.StatusOK, model.Error("SSE代理未配置", 500))
		return
	}

	// Determine target path based on current request
	targetPath := ctx.Request.URL.Path
	targetURL := cfg.StreamBaseURL + targetPath
	req, err := http.NewRequestWithContext(ctx.Request.Context(), "POST", targetURL, bytes.NewReader(body))
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("创建请求失败", 500))
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", token)
	req.Header.Set("Accept", "text/event-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusOK, model.Error("请求外部服务失败: "+err.Error(), 500))
		return
	}
	defer resp.Body.Close()

	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")
	ctx.Writer.Flush()

	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			ctx.Writer.Write(buf[:n])
			ctx.Writer.Flush()
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			break
		}
	}
}
