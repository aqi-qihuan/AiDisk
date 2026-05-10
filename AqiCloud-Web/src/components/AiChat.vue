<template>
  <div class="chat-container">
    <!-- 精美头部 -->
    <div class="chat-header">
      <div class="header-content">
        <div class="header-avatar">
          <div class="avatar-glow"></div>
          <div class="avatar-inner">🤖</div>
        </div>
        <div class="header-info">
          <h3 class="header-title">AI 聊天助理</h3>
          <p class="header-subtitle">
            <span class="status-dot"></span>
            在线
          </p>
        </div>
      </div>
      <div class="header-actions">
        <button class="header-btn" @click="clearChat" title="清空对话">
          <el-icon><Delete /></el-icon>
        </button>
      </div>
    </div>

    <!-- 消息区域 -->
    <div class="chat-content" ref="messageContainer">
      <!-- 欢迎界面 -->
      <div v-if="messages.length === 0 || (messages.length === 1 && messages[0].isHistory)" class="welcome-section">
        <div class="welcome-card">
          <div class="welcome-avatar-large">
            <div class="welcome-ring"></div>
            <div class="welcome-icon">🤖</div>
          </div>
          <h2 class="welcome-title">你好，我是 AI 助手</h2>
          <p class="welcome-desc">有什么我可以帮您的吗？</p>
          <div class="quick-actions">
            <div class="quick-action-item" @click="quickSend('帮我写一段代码')">
              <span class="action-icon">💻</span>
              <span class="action-text">写代码</span>
            </div>
            <div class="quick-action-item" @click="quickSend('解释一下这个概念')">
              <span class="action-icon">📚</span>
              <span class="action-text">解释概念</span>
            </div>
            <div class="quick-action-item" @click="quickSend('帮我优化这段文字')">
              <span class="action-icon">✍️</span>
              <span class="action-text">优化文字</span>
            </div>
            <div class="quick-action-item" @click="quickSend('给我一些建议')">
              <span class="action-icon">💡</span>
              <span class="action-text">获取建议</span>
            </div>
          </div>
        </div>
      </div>

      <!-- 消息列表 -->
      <template v-else>
        <div v-for="(message, index) in messages"
             :key="index"
             :class="['message-item', message.type, { 'first-message': index === 0 }]">
          <div class="message-avatar-wrapper" v-if="message.type === 'ai'">
            <div class="message-avatar ai-avatar">
              <span>🤖</span>
            </div>
          </div>
          <div class="message-bubble">
            <div class="message-sender">{{ message.type === 'ai' ? 'AI 助手' : '我' }}</div>
            <div class="message-content">
              <div v-if="message.type === 'ai' && !message.isHistory" class="typewriter">
                {{ message.displayText }}
              </div>
              <span v-else>{{ message.content }}</span>
            </div>
            <div class="message-time">{{ formatTime(message.timestamp) }}</div>
          </div>
          <div class="message-avatar-wrapper" v-if="message.type === 'user'">
            <div class="message-avatar user-avatar">
              <span>👤</span>
            </div>
          </div>
        </div>
      </template>

      <!-- 加载状态 -->
      <div v-if="isLoading" class="message-item ai loading-item">
        <div class="message-avatar-wrapper">
          <div class="message-avatar ai-avatar">
            <span>🤖</span>
          </div>
        </div>
        <div class="message-bubble">
          <div class="message-sender">AI 助手</div>
          <div class="message-content loading">
            <div class="typing-indicator">
              <span class="dot"></span>
              <span class="dot"></span>
              <span class="dot"></span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 输入区域 -->
    <div class="chat-input-area">
      <div class="input-wrapper">
        <textarea
          v-model="inputMessage"
          placeholder="输入消息..."
          :disabled="isLoading"
          @keyup.enter.prevent="sendMessage"
          @keydown.enter.shift.prevent="inputMessage += '\n'"
          rows="1"
          ref="textareaRef"
        ></textarea>
        <div class="input-actions">
          <div class="input-tips">
            <span class="tip-badge">Enter 发送</span>
            <span class="tip-badge">Shift+Enter 换行</span>
          </div>
          <button
            class="send-btn"
            :class="{ 'active': inputMessage.trim(), 'loading': isLoading }"
            @click="sendMessage"
            :disabled="!inputMessage.trim() || isLoading"
          >
            <svg v-if="!isLoading" class="send-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M22 2L11 13M22 2l-7 20-4-9-9-4 20-7z"/>
            </svg>
            <div v-else class="btn-spinner"></div>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick, watch } from 'vue'
import axios from 'axios'
import { ElMessage } from 'element-plus'
import { Delete } from '@element-plus/icons-vue'
import { useLoginUserStore } from "@/store/user";

const { token } = useLoginUserStore();
const messages = ref([])
const textareaRef = ref(null)
const inputMessage = ref('')
const messageContainer = ref(null)
const isLoading = ref(false)
// 格式化时间
const formatTime = (timestamp) => {
  if (!timestamp) return ''
  const date = new Date(timestamp)
  return date.toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit' })
}

// 快捷发送
const quickSend = (text) => {
  inputMessage.value = text
  sendMessage()
}

// 清空对话
const clearChat = () => {
  messages.value = []
}

// 获取历史记录
const fetchChatHistory = async () => {
  try {
    const response = await axios.get('/api/chat/history', {
      headers: { 'Authorization': `Bearer ${token}` }
    });
    
    if (response.data.messages && Array.isArray(response.data.messages)) {
      const historyMessages = response.data.messages.map(msg => ({
        type: msg.role === 'user' ? 'user' : 'ai',
        content: msg.content,
        displayText: msg.content,
        timestamp: msg.timestamp,
        isHistory: true
      }))
      messages.value = historyMessages
      await nextTick()
      scrollToBottom()
    }
  } catch (error) {
    if (error.response?.status === 422) {
      const errorDetail = error.response.data.detail
      ElMessage.error(`验证错误: ${errorDetail.map(d => d.msg).join(', ')}`)
    } else {
      ElMessage.error('获取历史记录失败')
    }
    console.error('获取历史记录错误:', error)
  }
}

// 发送消息
const sendMessage = async () => {
  if (!inputMessage.value.trim() || isLoading.value) return

  isLoading.value = true

  // 添加用户消息
  const userMessage = {
    type: 'user',
    content: inputMessage.value,
    displayText: inputMessage.value,
    timestamp: new Date().toISOString(),
    isHistory: false
  }
  messages.value.push(userMessage)

  try {
    // 发送聊天请求
    const response = await fetch('/api/chat/stream', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify({
        message: inputMessage.value  // 修改为符合接口要求的参数
      })
    })

    if (response.status === 422) {
      const errorData = await response.json()
      throw new Error(`验证错误: ${errorData.detail.map(d => d.msg).join(', ')}`)
    }

    if (!response.ok) {
      throw new Error(`请求失败，状态码: ${response.status}`)
    }

    // 处理流式响应
    await handleStreamResponse(response)
  } catch (error) {
    isLoading.value = false
    ElMessage.error(error.message || 'AI响应失败')
    console.error('发送消息错误:', error)

    // 添加错误提示消息
    messages.value.push({
      type: 'ai',
      content: '很抱歉，处理您的请求时出现错误。请稍后再试。',
      displayText: '很抱歉，处理您的请求时出现错误。请稍后再试。',
      timestamp: new Date().toISOString(),
      isHistory: false
    })
  }

  inputMessage.value = ''
  await nextTick()
  scrollToBottom()
}

// 处理流式响应
const handleStreamResponse = async (response) => {
  const reader = response.body.getReader()
  const decoder = new TextDecoder()
  let aiMessage = {
    type: 'ai',
    content: '',
    displayText: '',
    timestamp: new Date().toISOString(),
    isHistory: false
  }
  messages.value.push(aiMessage)
  isLoading.value = false

  try {
    while (true) {
      const { done, value } = await reader.read()
      if (done) break

      const chunk = decoder.decode(value)
      // 直接使用返回的字符串，不需要 JSON 解析
      aiMessage.content += chunk
      aiMessage.displayText += chunk
      scrollToBottom()
    }
  } catch (error) {
    console.error('处理流式响应错误:', error)
    ElMessage.error('接收AI响应出错')
  }
}

// 滚动到底部
const scrollToBottom = () => {
  const container = messageContainer.value
  if (container) {
    container.scrollTop = container.scrollHeight
  }
}

// 打字机效果函数
const typeWriter = (message, speed = 50) => {
  let index = 0
  message.displayText = ''

  const timer = setInterval(() => {
    if (index < message.content.length) {
      message.displayText += message.content[index]
      index++
    } else {
      clearInterval(timer)
    }
  }, speed)
}

// 自动调整文本框高度
watch(inputMessage, () => {
  if (textareaRef.value) {
    textareaRef.value.style.height = 'auto'
    textareaRef.value.style.height = Math.min(textareaRef.value.scrollHeight, 120) + 'px'
  }
})

// 组件加载时获取历史记录
onMounted(async () => {
  await fetchChatHistory()
})

// 监听消息变化，自动滚动到底部
watch(messages, () => {
  scrollToBottom()
}, {deep: true})
</script>

<style scoped>
/* ========== 整体布局 ========== */
.chat-container {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  background: linear-gradient(180deg, #f8fafc 0%, #ffffff 100%);
  position: relative;
  overflow: hidden;
}

/* ========== 头部 ========== */
.chat-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 24px;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid rgba(102, 126, 234, 0.1);
  position: relative;
  z-index: 10;
}

.chat-header::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(102, 126, 234, 0.2), transparent);
}

.header-content {
  display: flex;
  align-items: center;
  gap: 14px;
}

.header-avatar {
  position: relative;
  width: 44px;
  height: 44px;
}

.avatar-glow {
  position: absolute;
  inset: -2px;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  opacity: 0.4;
  filter: blur(8px);
  animation: pulse 2s ease-in-out infinite;
}

.avatar-inner {
  position: relative;
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.header-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.header-title {
  font-size: 17px;
  font-weight: 600;
  color: #1e293b;
  margin: 0;
}

.header-subtitle {
  font-size: 13px;
  color: #64748b;
  display: flex;
  align-items: center;
  gap: 6px;
  margin: 0;
}

.status-dot {
  width: 8px;
  height: 8px;
  background: #10b981;
  border-radius: 50%;
  box-shadow: 0 0 0 2px rgba(16, 185, 129, 0.2);
  animation: pulse-green 2s ease-in-out infinite;
}

.header-actions {
  display: flex;
  gap: 8px;
}

.header-btn {
  width: 36px;
  height: 36px;
  border: none;
  background: rgba(102, 126, 234, 0.08);
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: #64748b;
  transition: all 0.2s ease;
}

.header-btn:hover {
  background: rgba(102, 126, 234, 0.15);
  color: #667eea;
  transform: translateY(-1px);
}

/* ========== 消息区域 ========== */
.chat-content {
  flex: 1;
  overflow-y: auto;
  padding: 24px;
  display: flex;
  flex-direction: column;
  gap: 20px;
}

/* 欢迎界面 */
.welcome-section {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100%;
  padding: 40px 20px;
}

.welcome-card {
  text-align: center;
  max-width: 480px;
  animation: fadeInUp 0.6s ease-out;
}

.welcome-avatar-large {
  position: relative;
  width: 100px;
  height: 100px;
  margin: 0 auto 24px;
}

.welcome-ring {
  position: absolute;
  inset: 0;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 50%, #f093fb 100%);
  animation: rotate 4s linear infinite;
}

.welcome-ring::before {
  content: '';
  position: absolute;
  inset: 3px;
  border-radius: 50%;
  background: #f8fafc;
}

.welcome-icon {
  position: absolute;
  inset: 6px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 48px;
  animation: bounce 2s ease-in-out infinite;
}

.welcome-title {
  font-size: 26px;
  font-weight: 700;
  color: #1e293b;
  margin: 0 0 8px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.welcome-desc {
  font-size: 16px;
  color: #64748b;
  margin: 0 0 32px;
}

.quick-actions {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 12px;
}

.quick-action-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px;
  background: white;
  border: 1px solid rgba(102, 126, 234, 0.1);
  border-radius: 14px;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  text-align: left;
}

.quick-action-item:hover {
  border-color: rgba(102, 126, 234, 0.3);
  box-shadow: 0 8px 24px rgba(102, 126, 234, 0.12);
  transform: translateY(-2px);
}

.action-icon {
  width: 40px;
  height: 40px;
  background: linear-gradient(135deg, rgba(102, 126, 234, 0.1) 0%, rgba(118, 75, 162, 0.1) 100%);
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
}

.action-text {
  font-size: 14px;
  font-weight: 500;
  color: #334155;
}

/* 消息项 */
.message-item {
  display: flex;
  gap: 12px;
  max-width: 85%;
  animation: fadeInUp 0.3s ease-out;
}

.message-item.user {
  margin-left: auto;
  flex-direction: row-reverse;
}

.message-item.ai {
  margin-right: auto;
}

.message-avatar-wrapper {
  flex-shrink: 0;
}

.message-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
}

.message-avatar.ai-avatar {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.message-avatar.user-avatar {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
  box-shadow: 0 4px 12px rgba(245, 87, 108, 0.3);
}

.message-bubble {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.message-sender {
  font-size: 12px;
  color: #94a3b8;
  font-weight: 500;
  padding: 0 4px;
}

.message-item.user .message-sender {
  text-align: right;
}

.message-content {
  padding: 14px 18px;
  border-radius: 18px;
  font-size: 15px;
  line-height: 1.6;
  word-break: break-word;
  max-width: 100%;
}

.message-item.ai .message-content {
  background: white;
  border: 1px solid rgba(102, 126, 234, 0.1);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  border-bottom-left-radius: 6px;
}

.message-item.user .message-content {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
  border-bottom-right-radius: 6px;
}

.message-time {
  font-size: 11px;
  color: #cbd5e1;
  padding: 0 4px;
}

.message-item.user .message-time {
  text-align: right;
}

/* 加载动画 */
.loading-item {
  opacity: 0.8;
}

.typing-indicator {
  display: flex;
  gap: 4px;
  padding: 8px 4px;
}

.typing-indicator .dot {
  width: 6px;
  height: 6px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  animation: typing 1.4s ease-in-out infinite;
}

.typing-indicator .dot:nth-child(1) { animation-delay: 0s; }
.typing-indicator .dot:nth-child(2) { animation-delay: 0.2s; }
.typing-indicator .dot:nth-child(3) { animation-delay: 0.4s; }

/* ========== 输入区域 ========== */
.chat-input-area {
  padding: 20px 24px;
  background: linear-gradient(180deg, rgba(255,255,255,0) 0%, rgba(255,255,255,0.9) 20%, white 100%);
  position: relative;
}

.chat-input-area::before {
  content: '';
  position: absolute;
  top: 0;
  left: 24px;
  right: 24px;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(102, 126, 234, 0.15), transparent);
}

.input-wrapper {
  position: relative;
  max-width: 800px;
  margin: 0 auto;
  background: linear-gradient(135deg, #f8fafc 0%, #ffffff 100%);
  border-radius: 20px;
  border: 1px solid rgba(102, 126, 234, 0.12);
  padding: 16px 70px 16px 20px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.04), inset 0 1px 0 rgba(255, 255, 255, 0.8);
}

.input-wrapper:hover {
  border-color: rgba(102, 126, 234, 0.2);
  box-shadow: 0 8px 30px rgba(102, 126, 234, 0.08), inset 0 1px 0 rgba(255, 255, 255, 0.8);
}

.input-wrapper:focus-within {
  border-color: rgba(102, 126, 234, 0.35);
  box-shadow: 0 0 0 4px rgba(102, 126, 234, 0.06), 0 8px 30px rgba(102, 126, 234, 0.1);
  transform: translateY(-1px);
}

.input-wrapper textarea {
  width: 100%;
  border: none;
  background: transparent;
  resize: none;
  outline: none;
  font-size: 15px;
  line-height: 1.6;
  max-height: 120px;
  min-height: 24px;
  overflow-y: auto;
  font-family: inherit;
  color: #334155;
}

.input-wrapper textarea::placeholder {
  color: #94a3b8;
}

.input-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 10px;
}

.input-tips {
  display: flex;
  gap: 8px;
}

.tip-badge {
  font-size: 11px;
  color: #64748b;
  background: rgba(102, 126, 234, 0.06);
  padding: 4px 10px;
  border-radius: 6px;
  border: 1px solid rgba(102, 126, 234, 0.1);
}

.send-btn {
  position: absolute;
  right: 12px;
  bottom: 12px;
  width: 44px;
  height: 44px;
  border: none;
  border-radius: 14px;
  background: linear-gradient(135deg, #e2e8f0 0%, #f1f5f9 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.06);
}

.send-btn:not(:disabled):hover {
  transform: scale(1.05);
}

.send-btn.active {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
}

.send-btn.active:hover {
  transform: scale(1.05) translateY(-2px);
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.5);
}

.send-btn:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}

.send-icon {
  width: 20px;
  height: 20px;
  color: #94a3b8;
  transition: all 0.3s ease;
}

.send-btn.active .send-icon {
  color: white;
}

.btn-spinner {
  width: 20px;
  height: 20px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

/* ========== 滚动条 ========== */
.chat-content::-webkit-scrollbar {
  width: 6px;
}

.chat-content::-webkit-scrollbar-thumb {
  background: linear-gradient(180deg, rgba(102, 126, 234, 0.3) 0%, rgba(118, 75, 162, 0.3) 100%);
  border-radius: 3px;
}

.chat-content::-webkit-scrollbar-track {
  background: transparent;
}

/* ========== 动画 ========== */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes pulse {
  0%, 100% {
    opacity: 0.4;
    transform: scale(1);
  }
  50% {
    opacity: 0.6;
    transform: scale(1.05);
  }
}

@keyframes pulse-green {
  0%, 100% {
    opacity: 1;
    box-shadow: 0 0 0 2px rgba(16, 185, 129, 0.2);
  }
  50% {
    opacity: 0.8;
    box-shadow: 0 0 0 4px rgba(16, 185, 129, 0.15);
  }
}

@keyframes rotate {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

@keyframes bounce {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-8px);
  }
}

@keyframes typing {
  0%, 60%, 100% {
    transform: translateY(0);
    opacity: 1;
  }
  30% {
    transform: translateY(-6px);
    opacity: 0.5;
  }
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

/* ========== 响应式 ========== */
@media (max-width: 768px) {
  .chat-header {
    padding: 12px 16px;
  }

  .header-avatar {
    width: 38px;
    height: 38px;
  }

  .avatar-inner {
    font-size: 20px;
  }

  .header-title {
    font-size: 15px;
  }

  .chat-content {
    padding: 16px;
    gap: 16px;
  }

  .welcome-avatar-large {
    width: 80px;
    height: 80px;
  }

  .welcome-icon {
    font-size: 36px;
  }

  .welcome-title {
    font-size: 22px;
  }

  .quick-actions {
    grid-template-columns: 1fr;
  }

  .message-item {
    max-width: 92%;
    gap: 10px;
  }

  .message-avatar {
    width: 36px;
    height: 36px;
    font-size: 16px;
  }

  .message-content {
    padding: 12px 14px;
    font-size: 14px;
  }

  .chat-input-area {
    padding: 16px;
  }

  .input-wrapper {
    padding: 14px 60px 14px 16px;
    border-radius: 16px;
  }

  .send-btn {
    width: 40px;
    height: 40px;
    right: 10px;
    bottom: 10px;
    border-radius: 12px;
  }

  .send-icon {
    width: 18px;
    height: 18px;
  }

  .input-tips {
    display: none;
  }
}

.typewriter {
  white-space: pre-wrap;
  word-break: break-word;
}
</style>
