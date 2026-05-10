package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aqi/AqiCloud-AgentPan-Go/internal/config"
	"github.com/aqi/AqiCloud-AgentPan-Go/internal/model"
)

type ChatService struct{}

func NewChatService() *ChatService { return &ChatService{} }

func (s *ChatService) SendMessage(ctx context.Context, message, provider string) (*model.AiResponseDto, error) {
	cfg := config.GetConfig()

	switch provider {
	case "ollama":
		return s.callOllama(ctx, message, cfg.OllamaBaseURL)
	case "dashscope":
		return s.callDashScope(ctx, message, cfg.DashScopeAPIKey, cfg.DashScopeBase)
	default:
		return s.callOllama(ctx, message, cfg.OllamaBaseURL)
	}
}

func (s *ChatService) callDashScope(ctx context.Context, message, apiKey, baseURL string) (*model.AiResponseDto, error) {
	reqBody, _ := json.Marshal(map[string]interface{}{
		"model": "qwen-max",
		"messages": []map[string]string{
			{"role": "user", "content": message},
		},
	})

	req, err := http.NewRequestWithContext(ctx, "POST", baseURL+"/chat/completions", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var openAIResp struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		return nil, err
	}
	if len(openAIResp.Choices) == 0 {
		return nil, fmt.Errorf("AI未返回有效响应")
	}

	return &model.AiResponseDto{
		Content:  openAIResp.Choices[0].Message.Content,
		Model:    "qwen-max",
		Provider: "dashscope",
		Success:  true,
	}, nil
}

func (s *ChatService) callOllama(ctx context.Context, message, baseURL string) (*model.AiResponseDto, error) {
	reqBody, _ := json.Marshal(map[string]interface{}{
		"model": "qwen3:8b",
		"messages": []map[string]string{
			{"role": "user", "content": message},
		},
		"stream": false,
	})

	req, err := http.NewRequestWithContext(ctx, "POST", baseURL+"/api/chat", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var ollamaResp struct {
		Model   string `json:"model"`
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
		Done bool `json:"done"`
	}
	if err := json.Unmarshal(body, &ollamaResp); err != nil {
		return nil, err
	}

	return &model.AiResponseDto{
		Content:  ollamaResp.Message.Content,
		Model:    ollamaResp.Model,
		Provider: "ollama",
		Success:  true,
	}, nil
}
