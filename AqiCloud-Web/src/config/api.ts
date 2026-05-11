/**
 * FastAPI 服务配置（AI 相关功能）
 * 端口: 8000 (开发) / ai-api.a4i.icu (生产)
 * 用于: AI聊天、文档助手、网盘智答等功能
 */
// API配置
export const API_BASE_URL = 'http://127.0.0.1:8000/api'

// API路径配置
export const API_PATHS = {
  CHAT: '/chat/stream',
  DOCUMENT: '/document/stream',
  PAN_QUERY: '/pan/query',
  REWRITE: '/rewrite/stream'
}

// 获取完整的API URL
export const getApiUrl = (path: string) => `${API_BASE_URL}${path}` 