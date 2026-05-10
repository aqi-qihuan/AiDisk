package models

// ChatMessage 聊天消息
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest 聊天请求（对标 Python models/chat_schemas.py）
type ChatRequest struct {
	Message  string        `json:"message"`
	History  []ChatMessage `json:"history"`
	Model    string        `json:"model"`
	Provider string        `json:"provider"`
}

// ChatHistoryResponse 聊天历史响应
type ChatHistoryResponse struct {
	ConversationID string        `json:"conversation_id"`
	Messages       []ChatMessage `json:"messages"`
	Timestamp      string        `json:"timestamp"`
}
