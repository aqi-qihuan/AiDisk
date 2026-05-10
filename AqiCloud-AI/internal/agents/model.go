package agents

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/core"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
	openai "github.com/sashabaranov/go-openai"
)

// ChatModel OpenAI 兼容的 Eino ChatModel 实现
type ChatModel struct {
	client *openai.Client
	model  string
	temp   float64
	tools  []*schema.ToolInfo
}

// NewChatModel 创建 Eino ChatModel（对标 Python core/llm.py）
func NewChatModel(provider ...string) *ChatModel {
	var client *openai.Client
	if len(provider) > 0 && provider[0] != "" {
		client = core.NewLLMClientWithProvider(provider[0])
	} else {
		client = core.NewLLMClient()
	}
	c := core.GetConfig()
	return &ChatModel{client: client, model: c.LLMModelName, temp: c.LLMTemperature}
}

// GetType 返回模型类型
func (m *ChatModel) GetType() string {
	return "ChatModel"
}

// IsCallbacksEnabled 是否启用回调
func (m *ChatModel) IsCallbacksEnabled() bool {
	return true
}

// Generate 非流式生成（实现 model.BaseChatModel）
func (m *ChatModel) Generate(ctx context.Context, input []*schema.Message, opts ...model.Option) (*schema.Message, error) {
	opt := model.GetCommonOptions(&model.Options{}, opts...)

	messages := toOpenAIMessages(input)
	req := openai.ChatCompletionRequest{
		Model:       m.model,
		Messages:    messages,
		Temperature: float32(m.temp),
	}

	// 优先使用 option 传入的 tools，否则使用 model 自带的
	tools := m.tools
	if opt.Tools != nil {
		tools = opt.Tools
	}
	if len(tools) > 0 {
		req.Tools = toOpenAITools(tools)
	}

	resp, err := m.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("LLM 调用失败: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("LLM 未返回有效响应")
	}

	choice := resp.Choices[0]
	msg := &schema.Message{
		Role:    schema.Assistant,
		Content: choice.Message.Content,
	}

	if len(choice.Message.ToolCalls) > 0 {
		msg.ToolCalls = toEinoToolCalls(choice.Message.ToolCalls)
	}

	return msg, nil
}

// Stream 流式生成（实现 model.BaseChatModel）
func (m *ChatModel) Stream(ctx context.Context, input []*schema.Message, opts ...model.Option) (*schema.StreamReader[*schema.Message], error) {
	opt := model.GetCommonOptions(&model.Options{}, opts...)

	messages := toOpenAIMessages(input)
	req := openai.ChatCompletionRequest{
		Model:       m.model,
		Messages:    messages,
		Temperature: float32(m.temp),
		Stream:      true,
	}

	tools := m.tools
	if opt.Tools != nil {
		tools = opt.Tools
	}
	if len(tools) > 0 {
		req.Tools = toOpenAITools(tools)
	}

	stream, err := m.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("LLM 流式调用失败: %w", err)
	}

	sr, sw := schema.Pipe[*schema.Message](0)
	go func() {
		defer sw.Close()
		defer stream.Close()

		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				return
			}
			if len(resp.Choices) == 0 {
				continue
			}

			delta := resp.Choices[0].Delta
			if delta.Content != "" {
				sw.Send(&schema.Message{
					Role:    schema.Assistant,
					Content: delta.Content,
				}, nil)
			}

			if len(delta.ToolCalls) > 0 {
				toolCalls := toEinoToolCalls(delta.ToolCalls)
				sw.Send(&schema.Message{
					Role:      schema.Assistant,
					ToolCalls: toolCalls,
				}, nil)
			}

			if resp.Choices[0].FinishReason != "" {
				return
			}
		}
	}()

	return sr, nil
}

// BindTools 绑定工具到模型（实现 model.ChatModel）
func (m *ChatModel) BindTools(tools []*schema.ToolInfo) error {
	m.tools = tools
	return nil
}

// WithTools 返回带工具的新实例（实现 model.ToolCallingChatModel）
func (m *ChatModel) WithTools(tools []*schema.ToolInfo) (model.ToolCallingChatModel, error) {
	clone := *m
	clone.tools = tools
	return &clone, nil
}

// ========== 类型转换 ==========

func toOpenAIMessages(input []*schema.Message) []openai.ChatCompletionMessage {
	msgs := make([]openai.ChatCompletionMessage, 0, len(input))
	for _, msg := range input {
		role := string(msg.Role)
		m := openai.ChatCompletionMessage{Role: role, Content: msg.Content}
		if len(msg.ToolCalls) > 0 {
			m.ToolCalls = make([]openai.ToolCall, 0, len(msg.ToolCalls))
			for _, tc := range msg.ToolCalls {
				m.ToolCalls = append(m.ToolCalls, openai.ToolCall{
					ID:   tc.ID,
					Type: openai.ToolType(tc.Type),
					Function: openai.FunctionCall{
						Name:      tc.Function.Name,
						Arguments: tc.Function.Arguments,
					},
				})
			}
		}
		if msg.ToolCallID != "" {
			m.ToolCallID = msg.ToolCallID
		}
		msgs = append(msgs, m)
	}
	return msgs
}

func toOpenAITools(tools []*schema.ToolInfo) []openai.Tool {
	defs := make([]openai.Tool, 0, len(tools))
	for _, t := range tools {
		var params json.RawMessage
		if t.ParamsOneOf != nil {
			js, _ := t.ParamsOneOf.ToJSONSchema()
			if js != nil {
				params, _ = json.Marshal(js)
			}
		}
		defs = append(defs, openai.Tool{
			Type: openai.ToolTypeFunction,
			Function: &openai.FunctionDefinition{
				Name:        t.Name,
				Description: t.Desc,
				Parameters:  params,
			},
		})
	}
	return defs
}

func toEinoToolCalls(calls []openai.ToolCall) []schema.ToolCall {
	result := make([]schema.ToolCall, 0, len(calls))
	for _, c := range calls {
		result = append(result, schema.ToolCall{
			ID:   c.ID,
			Type: string(c.Type),
			Function: schema.FunctionCall{
				Name:      c.Function.Name,
				Arguments: c.Function.Arguments,
			},
		})
	}
	return result
}
