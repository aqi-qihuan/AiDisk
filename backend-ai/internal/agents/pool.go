package agents

import (
	"context"
	"sync"
)

// AgentPool 全局 Agent 实例池，按 provider 索引，启动时创建，请求时复用
type AgentPool struct {
	mu   sync.RWMutex
	pool map[string]*poolEntry
}

type poolEntry struct {
	Chat *ChatAgent
	Doc  *DocAgent
	Pan  *PanAgent
}

var defaultPool *AgentPool

func GetAgentPool() *AgentPool {
	return defaultPool
}

// InitAgentPool 初始化全局 Agent 池（应在 main.go 启动时调用）
func InitAgentPool(ctx context.Context) error {
	p := &AgentPool{
		pool: make(map[string]*poolEntry),
	}

	// 为每个 provider 预创建 Agent 实例
	for _, provider := range []string{"ollama", "openai_compatible", "default"} {
		chat, err := NewChatAgent(ctx, provider)
		if err != nil {
			return err
		}
		doc, err := NewDocAgent(ctx, provider)
		if err != nil {
			return err
		}
		pan, err := NewPanAgent(ctx, provider)
		if err != nil {
			return err
		}
		p.pool[provider] = &poolEntry{Chat: chat, Doc: doc, Pan: pan}
	}

	defaultPool = p
	return nil
}

// GetChat 获取聊天 Agent，若 provider 未预创建则实时创建并缓存
func (p *AgentPool) GetChat(ctx context.Context, provider string) (*ChatAgent, error) {
	p.mu.RLock()
	entry, ok := p.pool[provider]
	p.mu.RUnlock()

	if ok {
		return entry.Chat, nil
	}

	// 懒加载
	p.mu.Lock()
	defer p.mu.Unlock()

	// double-check
	if entry, ok = p.pool[provider]; ok {
		return entry.Chat, nil
	}

	chat, err := NewChatAgent(ctx, provider)
	if err != nil {
		return nil, err
	}
	doc, err := NewDocAgent(ctx, provider)
	if err != nil {
		return nil, err
	}
	pan, err := NewPanAgent(ctx, provider)
	if err != nil {
		return nil, err
	}
	p.pool[provider] = &poolEntry{Chat: chat, Doc: doc, Pan: pan}
	return chat, nil
}

// GetDoc 获取文档 Agent
func (p *AgentPool) GetDoc(ctx context.Context, provider string) (*DocAgent, error) {
	p.mu.RLock()
	entry, ok := p.pool[provider]
	p.mu.RUnlock()

	if ok {
		return entry.Doc, nil
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if entry, ok = p.pool[provider]; ok {
		return entry.Doc, nil
	}

	chat, err := NewChatAgent(ctx, provider)
	if err != nil {
		return nil, err
	}
	doc, err := NewDocAgent(ctx, provider)
	if err != nil {
		return nil, err
	}
	pan, err := NewPanAgent(ctx, provider)
	if err != nil {
		return nil, err
	}
	p.pool[provider] = &poolEntry{Chat: chat, Doc: doc, Pan: pan}
	return doc, nil
}

// GetPan 获取网盘 Agent
func (p *AgentPool) GetPan(ctx context.Context, provider string) (*PanAgent, error) {
	p.mu.RLock()
	entry, ok := p.pool[provider]
	p.mu.RUnlock()

	if ok {
		return entry.Pan, nil
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if entry, ok = p.pool[provider]; ok {
		return entry.Pan, nil
	}

	chat, err := NewChatAgent(ctx, provider)
	if err != nil {
		return nil, err
	}
	doc, err := NewDocAgent(ctx, provider)
	if err != nil {
		return nil, err
	}
	pan, err := NewPanAgent(ctx, provider)
	if err != nil {
		return nil, err
	}
	p.pool[provider] = &poolEntry{Chat: chat, Doc: doc, Pan: pan}
	return pan, nil
}
