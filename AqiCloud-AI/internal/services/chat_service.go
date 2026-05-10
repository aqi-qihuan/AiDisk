package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/core"
	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/models"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// chatMessage Redis List 中单条消息的存储格式
type chatMessage struct {
	Role      string `json:"role"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

// ChatService 聊天服务（对标 Python services/chat_service.py）
type ChatService struct {
	redis    *redis.Client
	provider string
	db       *gorm.DB
}

var (
	chatServiceMu     sync.Mutex
	chatServiceCache  = make(map[string]*ChatService)
	redisOnce         sync.Once
	redisClient       *redis.Client
)

func getRedisClient() *redis.Client {
	redisOnce.Do(func() {
		c := core.GetConfig()
		redisClient = redis.NewClient(&redis.Options{
			Addr:         c.RedisAddr(),
			Password:     c.RedisPassword,
			DB:           c.RedisDB,
			MaxRetries:   3,
			DialTimeout:  5 * time.Second,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 3 * time.Second,
			PoolSize:     c.RedisMaxConnections,
		})
	})
	return redisClient
}

func GetChatService(provider ...string) *ChatService {
	p := "default"
	if len(provider) > 0 && provider[0] != "" {
		p = provider[0]
	}

	chatServiceMu.Lock()
	defer chatServiceMu.Unlock()

	if svc, ok := chatServiceCache[p]; ok {
		return svc
	}

	svc := &ChatService{
		redis:    getRedisClient(),
		provider: p,
		db:       core.GetDB(),
	}
	chatServiceCache[p] = svc
	return svc
}

func (s *ChatService) chatKey(accountID string) string {
	return fmt.Sprintf("chat_history:%s", accountID)
}

func (s *ChatService) summaryKey(accountID string) string {
	return fmt.Sprintf("chat_summary:%s", accountID)
}

// SaveChatHistory 保存聊天历史到 Redis（全量覆盖，兼容旧接口）
func (s *ChatService) SaveChatHistory(ctx context.Context, accountID string, messages []map[string]any) error {
	c := core.GetConfig()

	if len(messages) > c.MaxChatHistoryMessages {
		messages = messages[len(messages)-c.MaxChatHistoryMessages:]
	}

	data, _ := json.Marshal(messages)
	return s.redis.SetEx(ctx, s.chatKey(accountID), data, time.Duration(c.RedisChatHistoryTTL)*time.Second).Err()
}

// GetChatHistory 获取用户聊天历史（使用 Redis List，LRANGE 读取）
func (s *ChatService) GetChatHistory(ctx context.Context, accountID string) ([]map[string]any, error) {
	c := core.GetConfig()
	key := s.chatKey(accountID)

	// Redis List: LPUSH 从头部插入，LRANGE 0 -1 读取全量（从旧到新）
	vals, err := s.redis.LRange(ctx, key, 0, int64(c.MaxChatHistoryMessages*2-1)).Result()
	if err == redis.Nil {
		return []map[string]any{}, nil
	}
	if err != nil {
		return nil, err
	}

	// List 中是倒序存储（最新的在头部），LRANGE 返回从头部开始，需要反转
	// LRANGE 0 N 返回 list[0]...list[N]，而 list[0] 是最新插入的
	// 所以 LRANGE 返回的是 [最新, ..., 最旧]，需要反转为 [最旧, ..., 最新]
	var messages []map[string]any
	for i := len(vals) - 1; i >= 0; i-- {
		var msg chatMessage
		if err := json.Unmarshal([]byte(vals[i]), &msg); err != nil {
			continue
		}
		messages = append(messages, map[string]any{
			"role":      msg.Role,
			"content":   msg.Content,
			"timestamp": msg.Timestamp,
		})
	}
	return messages, nil
}

// AddMessage 添加一条消息（Redis List LPUSH + LTRIM，原子操作，并发安全）
func (s *ChatService) AddMessage(ctx context.Context, accountID, role, content string) error {
	c := core.GetConfig()
	key := s.chatKey(accountID)

	msg := chatMessage{
		Role:      role,
		Content:   content,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// 原子操作：LPUSH + LTRIM + EXPIRE，无需 read-modify-write
	pipe := s.redis.Pipeline()
	pipe.LPush(ctx, key, data)
	pipe.LTrim(ctx, key, 0, int64(c.MaxChatHistoryMessages*2-1))
	pipe.Expire(ctx, key, time.Duration(c.RedisChatHistoryTTL)*time.Second)
	_, err = pipe.Exec(ctx)
	return err
}

// ClearChatHistory 清空用户聊天记录
func (s *ChatService) ClearChatHistory(ctx context.Context, accountID string) error {
	return s.redis.Del(ctx, s.chatKey(accountID), s.summaryKey(accountID)).Err()
}

// SaveChatMessages 保存一轮对话（用户消息 + 助手回复）到 Redis 和 MySQL
func (s *ChatService) SaveChatMessages(ctx context.Context, accountID, userMsg, assistantMsg string) error {
	aid, _ := strconv.ParseInt(accountID, 10, 64)

	// 写入 MySQL
	s.db.Create(&models.ChatLog{AccountID: aid, Role: "user", Content: userMsg, Provider: s.provider})
	s.db.Create(&models.ChatLog{AccountID: aid, Role: "assistant", Content: assistantMsg, Provider: s.provider})
	log.Printf("[ChatLog] 用户%s | 问: %s | 答: %s", accountID, truncate(userMsg, 100), truncate(assistantMsg, 100))

	// 写入 Redis（保留原有缓存逻辑）
	if err := s.AddMessage(ctx, accountID, "user", userMsg); err != nil {
		return err
	}
	return s.AddMessage(ctx, accountID, "assistant", assistantMsg)
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

// GenerateSummary 生成聊天摘要
func (s *ChatService) GenerateSummary(ctx context.Context, accountID string) string {
	messages, err := s.GetChatHistory(ctx, accountID)
	if err != nil || len(messages) == 0 {
		return ""
	}

	c := core.GetConfig()

	if len(messages) < c.SummaryTriggerThreshold {
		recent := messages
		if len(recent) > 5 {
			recent = recent[len(recent)-5:]
		}
		parts := make([]string, 0, len(recent))
		for _, m := range recent {
			role := "助手"
			if m["role"] == "user" {
				role = "用户"
			}
			content, _ := m["content"].(string)
			if len(content) > 100 {
				content = content[:100]
			}
			parts = append(parts, fmt.Sprintf("%s:%s", role, content))
		}
		result := ""
		for i, p := range parts {
			if i > 0 {
				result += " | "
			}
			result += p
		}
		return result
	}

	recent := messages
	if len(recent) > c.SummaryTriggerThreshold {
		recent = recent[len(recent)-c.SummaryTriggerThreshold:]
	}

	parts := make([]string, 0, len(recent))
	for _, m := range recent {
		role := "U"
		if m["role"] == "user" {
			role = "用户"
		}
		content, _ := m["content"].(string)
		if len(content) > 50 {
			content = content[:50]
		}
		parts = append(parts, fmt.Sprintf("%s:%s", role, content))
	}

	result := ""
	for i, p := range parts {
		if i > 0 {
			result += "|"
		}
		result += p
	}

	s.redis.SetEx(ctx, s.summaryKey(accountID), result, time.Duration(c.RedisChatHistoryTTL)*time.Second)
	return result
}
