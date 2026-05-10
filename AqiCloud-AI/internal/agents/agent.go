package agents

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/tools"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

// ========== Chat Agent ==========

var chatSystemPrompt = `你是智能助手，可日常对话、搜索实时信息、记住对话历史。
规则：
1. 实时问题（股价/天气/新闻）必须使用搜索工具
2. 不编造实时数据
3. 回答专业、友好、准确`

// ChatAgent 聊天 Agent（对标 Python agent/chat_agent.py）
// 实现 ReAct 循环：LLM -> 检查工具调用 -> 执行工具 -> 结果追加 -> 循环直到无工具调用
type ChatAgent struct {
	model *ChatModel
	tools []tool.BaseTool
}

// NewChatAgent 创建聊天 Agent
func NewChatAgent(ctx context.Context, provider ...string) (*ChatAgent, error) {
	p := "default"
	if len(provider) > 0 && provider[0] != "" {
		p = provider[0]
	}

	model := NewChatModel(p)

	webSearchTool := NewWebSearchTool()

	toolsInfo := []*schema.ToolInfo{
		{
			Name: "web_search",
			Desc: "使用此工具搜索最新的互联网信息。当你需要获取实时信息或不确定某个事实时使用",
			ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
				"query": {Type: schema.String, Desc: "搜索查询", Required: true},
			}),
		},
	}

	_, err := model.WithTools(toolsInfo)
	if err != nil {
		return nil, fmt.Errorf("绑定工具失败: %w", err)
	}

	return &ChatAgent{
		model: model,
		tools: []tool.BaseTool{webSearchTool},
	}, nil
}

// Invoke 非流式调用（ReAct 循环）
func (a *ChatAgent) Invoke(ctx context.Context, input []*schema.Message) (*schema.Message, error) {
	messages := make([]*schema.Message, len(input))
	copy(messages, input)

	for i := 0; i < 10; i++ {
		resp, err := a.model.Generate(ctx, messages)
		if err != nil {
			return nil, fmt.Errorf("LLM 调用失败: %w", err)
		}

		// 没有工具调用，返回最终结果
		if len(resp.ToolCalls) == 0 {
			return resp, nil
		}

		// 执行工具调用
		for _, tc := range resp.ToolCalls {
			result, err := a.executeTool(ctx, tc)
			if err != nil {
				messages = append(messages, &schema.Message{
					Role:       schema.Tool,
					ToolCallID: tc.ID,
					Content:    fmt.Sprintf("工具执行失败: %v", err),
				})
				continue
			}
			messages = append(messages, &schema.Message{
				Role:       schema.Tool,
				ToolCallID: tc.ID,
				Content:    result,
			})
		}
	}

	return nil, fmt.Errorf("达到最大工具调用次数限制")
}

// Stream 流式调用
// 注意：流式模式下只输出第一轮 LLM 响应，不展示工具执行过程
func (a *ChatAgent) Stream(ctx context.Context, input []*schema.Message) (*schema.StreamReader[[]*schema.Message], error) {
	messages := make([]*schema.Message, len(input))
	copy(messages, input)

	// 先做一次非流式 Generate 来获取工具调用信息
	resp, err := a.model.Generate(ctx, messages)
	if err != nil {
		return nil, fmt.Errorf("LLM 调用失败: %w", err)
	}

	// 如果有工具调用，先执行工具再流式生成最终响应
	if len(resp.ToolCalls) > 0 {
		for _, tc := range resp.ToolCalls {
			result, err := a.executeTool(ctx, tc)
			if err != nil {
				messages = append(messages, resp)
				messages = append(messages, &schema.Message{
					Role:       schema.Tool,
					ToolCallID: tc.ID,
					Content:    fmt.Sprintf("工具执行失败: %v", err),
				})
				continue
			}
			messages = append(messages, resp)
			messages = append(messages, &schema.Message{
				Role:       schema.Tool,
				ToolCallID: tc.ID,
				Content:    result,
			})
		}

		// 用工具结果再调一次 LLM，流式输出最终回答
		sr, err := a.model.Stream(ctx, messages)
		if err != nil {
			return nil, err
		}
		return wrapStreamToSlice(sr), nil
	}

	// 没有工具调用，直接流式输出
	sr, err := a.model.Stream(ctx, input)
	if err != nil {
		return nil, err
	}
	return wrapStreamToSlice(sr), nil
}

// wrapStreamToSlice 将单消息流转换为消息切片流
func wrapStreamToSlice(sr *schema.StreamReader[*schema.Message]) *schema.StreamReader[[]*schema.Message] {
	return schema.StreamReaderWithConvert(sr, func(msg *schema.Message) ([]*schema.Message, error) {
		return []*schema.Message{msg}, nil
	})
}

// executeTool 执行单个工具调用
func (a *ChatAgent) executeTool(ctx context.Context, tc schema.ToolCall) (string, error) {
	for _, t := range a.tools {
		info, err := t.Info(ctx)
		if err != nil {
			continue
		}
		if info.Name == tc.Function.Name {
			invokable, ok := t.(tool.InvokableTool)
			if !ok {
				return "", fmt.Errorf("工具 %s 不支持 InvokableTool 接口", info.Name)
			}
			return invokable.InvokableRun(ctx, tc.Function.Arguments)
		}
	}
	return "", fmt.Errorf("未找到工具: %s", tc.Function.Name)
}

// ========== Doc Agent ==========

var docSystemPrompt = `你是一个专门处理文档内容的智能助手。你可以:
1. 分析上传的文档内容
2. 回答关于文档内容的问题
3. 提取文档中的关键信息
4. 概括文档的主要内容
请根据用户的问题和提供的工具来帮助用户解决文档相关问题。`

// DocAgent 文档处理 Agent
type DocAgent struct {
	model *ChatModel
}

// NewDocAgent 创建文档 Agent
func NewDocAgent(ctx context.Context, provider ...string) (*DocAgent, error) {
	p := "default"
	if len(provider) > 0 && provider[0] != "" {
		p = provider[0]
	}

	return &DocAgent{model: NewChatModel(p)}, nil
}

// Stream 流式处理文档
func (a *DocAgent) Stream(ctx context.Context, input []*schema.Message) (*schema.StreamReader[[]*schema.Message], error) {
	sr, err := a.model.Stream(ctx, input)
	if err != nil {
		return nil, err
	}
	return wrapStreamToSlice(sr), nil
}

// ========== Pan Agent ==========

// PanAgent 网盘查询 Agent
type PanAgent struct {
	model *ChatModel
}

// NewPanAgent 创建网盘查询 Agent
func NewPanAgent(ctx context.Context, provider ...string) (*PanAgent, error) {
	p := "default"
	if len(provider) > 0 && provider[0] != "" {
		p = provider[0]
	}

	return &PanAgent{model: NewChatModel(p)}, nil
}

// Query 非流式查询
func (a *PanAgent) Query(ctx context.Context, input []*schema.Message) (*schema.Message, error) {
	return a.model.Generate(ctx, input)
}

// FastIntentParse 关键词快速意图分类，避免不必要的 LLM 调用
func FastIntentParse(query string) string {
	q := strings.ToLower(query)
	storageKeys := []string{"存储", "容量", "空间", "使用量", "已用", "剩余"}
	statKeys := []string{"统计", "多少", "总数", "总大", "分布", "类型"}
	searchKeys := []string{"搜索", "查找", "找一", "检索", "查询文件"}

	for _, k := range storageKeys {
		if strings.Contains(q, k) {
			return "storage_info"
		}
	}
	for _, k := range statKeys {
		if strings.Contains(q, k) {
			return "file_statistics"
		}
	}
	for _, k := range searchKeys {
		if strings.Contains(q, k) {
			return "search"
		}
	}
	return "file_list"
}

// IntentionParse 意图解析
func (a *PanAgent) IntentionParse(ctx context.Context, query string) (string, error) {
	messages := []*schema.Message{
		{
			Role:    schema.System,
			Content: "你是一个意图分类器，根据用户输入判断查询类型。可选类型: storage(存储空间), file_list(文件列表), file_statistics(文件统计), search(文件搜索)。只返回类型名称，不要其他内容。",
		},
		{
			Role:    schema.User,
			Content: fmt.Sprintf("根据以下查询请求，判断用户想要查询的类型，仅返回类型名称:\n%s", query),
		},
	}

	resp, err := a.model.Generate(ctx, messages)
	if err != nil {
		return "file_list", err
	}

	result := strings.TrimSpace(resp.Content)
	if result == "" {
		return "file_list", nil
	}
	return result, nil
}

// ========== WebSearch Eino Tool ==========

// WebSearchTool Eino 工具实现
type WebSearchTool struct{}

// NewWebSearchTool 创建搜索工具
func NewWebSearchTool() *WebSearchTool {
	return &WebSearchTool{}
}

// Info 返回工具元信息
func (t *WebSearchTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "web_search",
		Desc: "使用此工具搜索最新的互联网信息。当你需要获取实时信息或不确定某个事实时使用",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"query": {Type: schema.String, Desc: "搜索查询", Required: true},
		}),
	}, nil
}

// InvokableRun 执行搜索
func (t *WebSearchTool) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	var args struct {
		Query string `json:"query"`
	}
	if err := json.Unmarshal([]byte(argumentsInJSON), &args); err != nil {
		return "", fmt.Errorf("解析搜索参数失败: %w", err)
	}

	result, err := tools.WebSearch(args.Query)
	if err != nil {
		return "", fmt.Errorf("搜索失败: %w", err)
	}

	return result, nil
}

// ========== Helper ==========

// BuildMessages 构建消息列表
func BuildMessages(systemPrompt string, history []*schema.Message, userInput string) []*schema.Message {
	msgs := []*schema.Message{
		{Role: schema.System, Content: systemPrompt},
	}
	msgs = append(msgs, history...)
	msgs = append(msgs, &schema.Message{Role: schema.User, Content: userInput})
	return msgs
}

// BuildPanMessages 构建网盘查询消息
func BuildPanMessages(systemPrompt string, userInput string) []*schema.Message {
	return []*schema.Message{
		{Role: schema.System, Content: systemPrompt},
		{Role: schema.User, Content: userInput},
	}
}

// PanSystemPrompt 导出系统提示词
func PanSystemPrompt() string { return panSystemPrompt }

var panSystemPrompt = `网盘查询助手，仅查询数据，禁止修改。
规则：
1. 只读查询，禁止写入/删除/修改
2. 中文回复，简洁专业
3. 模糊请求需询问明确
4. 只查当前用户数据（account_id条件）
5. 返回JSON格式

响应格式：
{"type": "file_list|storage_info|file_statistics", "data": {...}}`

// ChatAgentSystemPrompt 导出聊天系统提示词
func ChatAgentSystemPrompt() string { return chatSystemPrompt }

// DocAgentSystemPrompt 导出文档系统提示词
func DocAgentSystemPrompt() string { return docSystemPrompt }

// EstimateTokenCount 估算文本的 token 数量（粗略估算）
func EstimateTokenCount(text string) int {
	// 中文约 1.5 token/字，英文约 1 token/词
	count := 0
	for _, r := range text {
		if r > 127 {
			count += 2 // 中文/多字节字符按 2 token 估算
		} else {
			count++
		}
	}
	return (count + 1) / 2
}

// TrimMessages 裁剪历史消息到 token 预算内，保留 system prompt 和最近的消息
func TrimMessages(messages []*schema.Message, maxTokens int) []*schema.Message {
	if len(messages) == 0 {
		return messages
	}

	// 分离 system 和其余消息
	var systemMsgs []*schema.Message
	var rest []*schema.Message
	for _, m := range messages {
		if m.Role == schema.System {
			systemMsgs = append(systemMsgs, m)
		} else {
			rest = append(rest, m)
		}
	}

	// 计算 system 消息的 token 数
	systemTokens := 0
	for _, m := range systemMsgs {
		systemTokens += EstimateTokenCount(m.Content)
	}

	budget := maxTokens - systemTokens
	if budget <= 0 {
		return systemMsgs
	}

	// 从后向前累积消息，保留最近的消息
	var kept []*schema.Message
	usedTokens := 0
	for i := len(rest) - 1; i >= 0; i-- {
		tokens := EstimateTokenCount(rest[i].Content)
		if usedTokens+tokens > budget {
			break
		}
		kept = append(kept, rest[i])
		usedTokens += tokens
	}

	// 反转恢复原始顺序
	for i, j := 0, len(kept)-1; i < j; i, j = i+1, j-1 {
		kept[i], kept[j] = kept[j], kept[i]
	}

	result := make([]*schema.Message, 0, len(systemMsgs)+len(kept))
	result = append(result, systemMsgs...)
	result = append(result, kept...)
	return result
}

// ReadToolCallsFromStream 从流式响应中累积工具调用
func ReadToolCallsFromStream(sr *schema.StreamReader[[]*schema.Message]) ([]*schema.Message, string, error) {
	defer sr.Close()

	var messages []*schema.Message
	var content strings.Builder

	for {
		chunks, err := sr.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, "", err
		}
		for _, chunk := range chunks {
			messages = append(messages, chunk)
			if chunk.Content != "" {
				content.WriteString(chunk.Content)
			}
		}
	}

	return messages, content.String(), nil
}
