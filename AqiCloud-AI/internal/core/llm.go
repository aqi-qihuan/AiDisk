package core

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

// NewLLMClient 根据配置创建 OpenAI 兼容客户端
// Ollama 的 /v1/chat/completions 接口兼容 OpenAI 协议
func NewLLMClient() *openai.Client {
	c := GetConfig()

	if c.LLMProvider == "ollama" {
		config := openai.DefaultConfig("ollama")
		config.BaseURL = c.LLMOllamaBaseURL + "/v1"
		return openai.NewClientWithConfig(config)
	}

	config := openai.DefaultConfig(c.LLMAPIKey)
	config.BaseURL = c.LLMBaseURL
	return openai.NewClientWithConfig(config)
}

// NewLLMClientWithProvider 按指定提供商创建客户端
func NewLLMClientWithProvider(provider string) *openai.Client {
	c := GetConfig()

	if provider == "ollama" {
		config := openai.DefaultConfig("ollama")
		config.BaseURL = c.LLMOllamaBaseURL + "/v1"
		return openai.NewClientWithConfig(config)
	}

	config := openai.DefaultConfig(c.LLMAPIKey)
	config.BaseURL = c.LLMBaseURL
	return openai.NewClientWithConfig(config)
}

// BuildChatRequest 构建 OpenAI 聊天请求
func BuildChatRequest(messages []openai.ChatCompletionMessage) openai.ChatCompletionRequest {
	c := GetConfig()
	return openai.ChatCompletionRequest{
		Model:       c.LLMModelName,
		Messages:    messages,
		Temperature: float32(c.LLMTemperature),
	}
}

// StreamChatCompletion 流式聊天
func StreamChatCompletion(ctx context.Context, client *openai.Client, req openai.ChatCompletionRequest) (*openai.ChatCompletionStream, error) {
	req.Stream = true
	return client.CreateChatCompletionStream(ctx, req)
}
