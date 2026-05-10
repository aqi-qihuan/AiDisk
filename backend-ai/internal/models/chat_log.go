package models

import (
	"time"

	"gorm.io/gorm"
)

// ChatLog 对话记录持久化模型（对标 Python 版的 chat_log 表）
type ChatLog struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID  int64     `gorm:"index;not null" json:"accountId"`
	Role       string    `gorm:"size:20;not null" json:"role"`       // user / assistant
	Content    string    `gorm:"type:text;not null" json:"content"`  // 对话内容
	Provider   string    `gorm:"size:50" json:"provider"`            // ollama / openai_compatible
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

func (ChatLog) TableName() string {
	return "t_chat_log"
}

func (c *ChatLog) BeforeCreate(tx *gorm.DB) error {
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now()
	}
	return nil
}
