package callbacks

import (
	"context"
	"log"

	"github.com/aqi/AqiCloud-Ai-Agent-Go/internal/core"
	"github.com/cloudwego/eino/callbacks"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/tool"
	einocb "github.com/cloudwego/eino/utils/callbacks"
)

// NewCostHandler 创建 Token 成本统计回调（商业化埋点）
func NewCostHandler(accountID string) *einocb.HandlerHelper {
	tracker := core.GetTokenTracker()
	_ = tracker

	return einocb.NewHandlerHelper().
		ChatModel(&einocb.ModelCallbackHandler{
			OnEnd: func(ctx context.Context, runInfo *callbacks.RunInfo, output *model.CallbackOutput) context.Context {
				log.Printf("[Cost] 用户 %s 完成 ChatModel 调用", accountID)
				return ctx
			},
		}).
		Tool(&einocb.ToolCallbackHandler{
			OnEnd: func(ctx context.Context, info *callbacks.RunInfo, output *tool.CallbackOutput) context.Context {
				log.Printf("[Cost] 用户 %s 完成工具调用: %s", accountID, info.Name)
				return ctx
			},
		})
}
