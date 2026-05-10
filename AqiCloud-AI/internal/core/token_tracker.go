package core

import (
	"sync"
	"time"
)

// TokenUsage Token 使用记录
type TokenUsage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
	Timestamp        time.Time
}

// TokenTracker Token 使用追踪器（对标 Python core/token_tracker.py）
type TokenTracker struct {
	mu            sync.Mutex
	dailyLimit    int
	hourlyLimit   int
	windowMinutes int
	usageRecords  map[string][]TokenUsage
	totalUsage    TokenUsage
}

var (
	tracker     *TokenTracker
	trackerOnce sync.Once
)

func GetTokenTracker() *TokenTracker {
	trackerOnce.Do(func() {
		c := GetConfig()
		tracker = &TokenTracker{
			dailyLimit:    c.TokenDailyLimit,
			hourlyLimit:   c.TokenHourlyLimit,
			windowMinutes: 60,
			usageRecords:  make(map[string][]TokenUsage),
		}
	})
	return tracker
}

// RecordUsage 记录 token 使用
func (t *TokenTracker) RecordUsage(userID string, promptTokens, completionTokens int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	usage := TokenUsage{
		PromptTokens:     promptTokens,
		CompletionTokens: completionTokens,
		TotalTokens:      promptTokens + completionTokens,
		Timestamp:        time.Now(),
	}

	t.usageRecords[userID] = append(t.usageRecords[userID], usage)
	t.totalUsage.TotalTokens += usage.TotalTokens
	t.cleanupOld(userID)
}

// CheckLimit 检查是否超出 token 限制
func (t *TokenTracker) CheckLimit(userID string, estimatedTokens int) (bool, string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	hourly := t.getUsageInWindow(userID, 60)
	daily := t.getUsageInWindow(userID, 1440)

	if hourly+estimatedTokens > t.hourlyLimit {
		return false, "小时token限制 exceeded"
	}
	if daily+estimatedTokens > t.dailyLimit {
		return false, "日token限制 exceeded"
	}
	return true, "OK"
}

// GetUserStats 获取用户 token 使用统计
func (t *TokenTracker) GetUserStats(userID string) map[string]any {
	t.mu.Lock()
	defer t.mu.Unlock()

	hourly := t.getUsageInWindow(userID, 60)
	daily := t.getUsageInWindow(userID, 1440)

	remaining := t.dailyLimit - daily
	if remaining < 0 {
		remaining = 0
	}

	pct := 0.0
	if t.dailyLimit > 0 {
		pct = float64(daily) / float64(t.dailyLimit) * 100
	}

	return map[string]any{
		"user_id":            userID,
		"hourly_usage":       hourly,
		"hourly_limit":       t.hourlyLimit,
		"hourly_remaining":   t.hourlyLimit - hourly,
		"daily_usage":        daily,
		"daily_limit":        t.dailyLimit,
		"daily_remaining":    remaining,
		"usage_percent":      pct,
	}
}

// GetGlobalStats 获取全局 token 使用统计
func (t *TokenTracker) GetGlobalStats() map[string]any {
	t.mu.Lock()
	defer t.mu.Unlock()

	return map[string]any{
		"total_prompt_tokens":     t.totalUsage.PromptTokens,
		"total_completion_tokens": t.totalUsage.CompletionTokens,
		"total_tokens":            t.totalUsage.TotalTokens,
		"active_users":            len(t.usageRecords),
	}
}

func (t *TokenTracker) getUsageInWindow(userID string, minutes int) int {
	records, ok := t.usageRecords[userID]
	if !ok {
		return 0
	}

	cutoff := time.Now().Add(-time.Duration(minutes) * time.Minute)
	total := 0
	for _, u := range records {
		if u.Timestamp.After(cutoff) {
			total += u.TotalTokens
		}
	}
	return total
}

func (t *TokenTracker) cleanupOld(userID string) {
	records := t.usageRecords[userID]
	cutoff := time.Now().Add(-time.Duration(t.windowMinutes) * time.Minute)
	filtered := records[:0]
	for _, u := range records {
		if u.Timestamp.After(cutoff) {
			filtered = append(filtered, u)
		}
	}
	t.usageRecords[userID] = filtered
}

// EstimateTokens 估算文本 token 数量（对标 Python estimate_tokens）
func EstimateTokens(text string) int {
	if text == "" {
		return 0
	}

	chinese := 0
	for _, c := range text {
		if c >= '一' && c <= '鿿' {
			chinese++
		}
	}
	other := len(text) - chinese

	estimated := chinese/2 + other/4
	if estimated < 1 {
		return 1
	}
	return estimated
}
