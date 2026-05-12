<template>
  <div class="chat-container">
    <!-- 消息列表区域 -->
    <div class="message-list" ref="messageList">
      <div
        v-for="(message, index) in messages"
        :key="index"
        :class="[
          'message-item',
          message.isUser ? 'user-message' : 'assistant-message',
        ]"
      >
        <!-- 头像 -->
        <div class="avatar">
          <img
            :src="message.isUser ? userAvatar : assistantAvatar"
            alt="avatar"
          />
        </div>

        <!-- 消息内容 -->
        <div class="message-content-wrapper">
          <div class="message-content">
            <template v-if="message.isLoading">
              <LoadingDots />
            </template>
            <div v-else class="message-text" v-html="message.content"></div>
          </div>

          <!-- 时间戳 -->
          <div
            class="message-time"
            v-if="!message.isLoading && message.content"
          >
            {{ formatTime(message.timestamp) }}
          </div>
        </div>
      </div>

      <!-- 空状态提示 - 精美欢迎界面 -->
      <div v-if="messages.length === 0" class="welcome-message">
        <div class="welcome-card">
          <div class="welcome-header">
            <div class="welcome-avatar">
              <div class="avatar-ring"></div>
              <div class="avatar-inner">
                <el-icon :size="36"><ChatDotRound /></el-icon>
              </div>
            </div>
            <h3 class="welcome-title">{{ t("ai.chatAssistantTitle") }}</h3>
            <p class="welcome-subtitle">{{ t("ai.chatAssistantSubtitle") }}</p>
          </div>

          <div class="welcome-divider"></div>

          <p class="welcome-description">
            {{ t("ai.chatWelcomeDesc") }}
          </p>

          <div class="suggestions">
            <div
              v-for="(suggestion, idx) in suggestions"
              :key="idx"
              class="suggestion-card"
              @click="
                inputMessage = suggestion.text;
                sendMessage();
              "
            >
              <div class="suggestion-icon">
                <el-icon :size="24">
                  <component :is="suggestion.icon" />
                </el-icon>
              </div>
              <div class="suggestion-content">
                <span class="suggestion-text">{{ suggestion.text }}</span>
              </div>
              <div class="suggestion-arrow">→</div>
            </div>
          </div>

          <div class="welcome-footer">
            <div class="typing-indicator">
              <span class="dot"></span>
              <span class="dot"></span>
              <span class="dot"></span>
            </div>
            <span class="footer-text">{{ t("ai.alwaysReady") }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- 输入区域 -->
    <div class="input-area">
      <div class="input-wrapper">
        <textarea
          v-model="inputMessage"
          :placeholder="t('ai.inputPlaceholder')"
          @keyup.enter.exact.prevent="sendMessage"
          @keydown.enter.shift.prevent="inputMessage += '\n'"
          rows="1"
          ref="textareaRef"
        ></textarea>

        <div class="input-actions">
          <div class="input-tips">
            <span class="tip-badge">{{ t("ai.enterSend") }}</span>
            <span class="tip-badge">{{ t("ai.shiftEnterNewline") }}</span>
          </div>

          <div
            class="send-btn"
            @click="sendMessage"
            :class="{ active: inputMessage.trim(), loading: isSending }"
          >
            <svg
              v-if="!isSending"
              class="icon"
              viewBox="0 0 1024 1024"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M931.4 481.3L815.9 396l0.3-0.4-604.7-249c-3.8-1.6-7.9-1.9-11.9-1l-0.3-0.7L135 171.3c-12.3 2.7-20.3 14.8-17.6 27.1 0.9 4.3 3.1 8 6.2 10.9l-0.3 0.7 155.6 130.3L492.3 436l-213.4 95.7-140.9 118c-9.7 8.1-11 22.6-2.8 32.3 3.9 4.6 9.3 7.5 15.2 8.1l0.3 0.7 63.9 26.3c4 1.6 8.1 1.9 12.1 1l0.3 0.7 201.4-83 139.3 57.4c4 1.6 8.1 1.9 12.1 0.9l0.3 0.7 63.9-26.3c12.3-2.7 20.3-14.8 17.6-27.1-0.9-4.3-3.1-8-6.2-10.9l0.3-0.7-102.8-86 389.4-174.6c11.3-5.1 16.3-18.3 11.3-29.6-2.3-4.9-6-8.6-10.8-11z"
              ></path>
            </svg>
            <div v-else class="loading-spinner"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, ref, onMounted, nextTick, watch } from "vue";
import { useI18n } from "vue-i18n";
import LoadingDots from "@/components/common/LoadingDots.vue";
import { useLoginUserStore } from "@/store/user";
import { getApiUrl, API_PATHS } from "@/config/api";
import userAvatar from "@/assets/assistant-avatar.png";
import assistantAvatar from "@/assets/user-avatar.png";
import {
  ChatDotRound,
  EditPen,
  Promotion,
  MagicStick,
} from "@element-plus/icons-vue";

export default defineComponent({
  name: "Chat",
  components: {
    LoadingDots,
    ChatDotRound,
    EditPen,
    Promotion,
    MagicStick,
  },
  setup() {
    const { t } = useI18n();
    const messages = ref<
      Array<{
        isUser: boolean;
        content: string;
        isLoading?: boolean;
        timestamp?: number;
      }>
    >([]);

    const inputMessage = ref("");
    const messageList = ref<HTMLElement | null>(null);
    const textareaRef = ref<HTMLTextAreaElement | null>(null);
    const isSending = ref(false);

    // 建议问题
    const suggestions = ref([
      { icon: "EditPen", text: t("ai.summarizeWork") },
      { icon: "Promotion", text: t("ai.improveEfficiency") },
      { icon: "MagicStick", text: t("ai.giveAdvice") },
    ]);

    const scrollToBottom = () => {
      if (messageList.value) {
        messageList.value.scrollTop = messageList.value.scrollHeight;
      }
    };

    // 格式化时间
    const formatTime = (timestamp?: number) => {
      if (!timestamp) return "";
      const date = new Date(timestamp);
      const hours = date.getHours().toString().padStart(2, "0");
      const minutes = date.getMinutes().toString().padStart(2, "0");
      return `${hours}:${minutes}`;
    };

    // 自动调整文本框高度
    watch(inputMessage, () => {
      if (textareaRef.value) {
        textareaRef.value.style.height = "auto";
        textareaRef.value.style.height =
          Math.min(textareaRef.value.scrollHeight, 120) + "px";
      }
    });

    const processStreamResponse = async (response: Response) => {
      const reader = response.body?.getReader();
      if (!reader) return;

      const currentMessageIndex = messages.value.length - 1;
      const decoder = new TextDecoder();

      try {
        while (true) {
          const { done, value } = await reader.read();
          if (done) break;

          const chunk = decoder.decode(value, { stream: true });
          const lines = chunk.split("\n");

          for (const line of lines) {
            if (line.trim() === "" || line.trim() === "data: [DONE]") continue;

            try {
              const jsonStr = line.replace(/^data: /, "").trim();
              if (!jsonStr) continue;

              const jsonData = JSON.parse(jsonStr);
              if (jsonData.code === 0 && jsonData.data) {
                messages.value[currentMessageIndex].isLoading = false;
                messages.value[currentMessageIndex].content += jsonData.data;

                await nextTick(() => {
                  scrollToBottom();
                });
              }
            } catch (e) {
              console.error("Error parsing JSON:", e);
            }
          }
        }
      } catch (error) {
        console.error("Stream reading error:", error);
        messages.value[currentMessageIndex].content = "[读取数据时发生错误]";
      } finally {
        messages.value[currentMessageIndex].isLoading = false;
        messages.value[currentMessageIndex].timestamp = Date.now();
        isSending.value = false;
        scrollToBottom();
      }
    };

    const sendMessage = async () => {
      if (!inputMessage.value.trim() || isSending.value) return;

      const userMessage = inputMessage.value;
      messages.value.push({
        isUser: true,
        content: userMessage,
        timestamp: Date.now(),
      });
      messages.value.push({
        isUser: false,
        content: "",
        isLoading: true,
        timestamp: Date.now(),
      });
      inputMessage.value = "";

      // 重置文本框高度
      if (textareaRef.value) {
        textareaRef.value.style.height = "auto";
      }

      scrollToBottom();

      try {
        isSending.value = true;
        const response = await fetch(getApiUrl(API_PATHS.CHAT), {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            token: useLoginUserStore().token || "",
          },
          body: JSON.stringify({
            message: userMessage,
          }),
        });

        if (!response.ok) {
          const errorText = await response.text();
          throw new Error(
            `Server responded with ${response.status}: ${errorText}`,
          );
        }

        await processStreamResponse(response);
      } catch (error) {
        console.error("Send message error:", error);
        messages.value.pop();
        messages.value.push({
          isUser: false,
          content: "抱歉，发生了错误，请稍后重试。",
          isLoading: false,
          timestamp: Date.now(),
        });
        isSending.value = false;
      }
      scrollToBottom();
    };

    onMounted(() => {
      scrollToBottom();
      // 聚焦到输入框
      if (textareaRef.value) {
        textareaRef.value.focus();
      }
    });

    return {
      t,
      messages,
      inputMessage,
      messageList,
      textareaRef,
      sendMessage,
      userAvatar,
      assistantAvatar,
      isSending,
      suggestions,
      formatTime,
    };
  },
});
</script>

<style scoped>
.chat-container {
  display: flex;
  flex-direction: column;
  height: 100%;
  background: var(--color-bg-surface, #14141c);
  position: relative;
  overflow: hidden;
}

/* 消息列表 */
.message-list {
  flex: 1;
  overflow-y: auto;
  padding: var(--ds-spacing-4);
  display: flex;
  flex-direction: column;
  gap: var(--ds-spacing-4);
}

/* 消息项 */
.message-item {
  display: flex;
  gap: var(--ds-spacing-3);
  max-width: 75%;
  animation: fadeInUp 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.message-item.user-message {
  margin-left: auto;
  flex-direction: row-reverse;
}

.message-item.assistant-message {
  margin-right: auto;
}

/* 头像 */
.avatar {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  overflow: hidden;
  flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  border: 2px solid rgba(255, 255, 255, 0.3);
  transition: transform 0.3s ease;
}

.avatar:hover {
  transform: scale(1.05);
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* 消息内容包裹器 */
.message-content-wrapper {
  display: flex;
  flex-direction: column;
  gap: var(--ds-spacing-1);
}

/* 消息内容 */
.message-content {
  padding: 14px 20px;
  border-radius: 20px;
  font-size: 15px;
  line-height: 1.7;
  word-break: break-word;
  min-height: 44px;
  max-width: 100%;
  position: relative;
  transition: all 0.3s ease;
}

/* 用户消息 - 金色渐变气泡（HOK设计） */
.message-item.user-message .message-content {
  background: linear-gradient(135deg, #d97706 0%, #f59e0b 50%, #fbbf24 100%);
  color: #0b0b10;
  font-weight: 500;
  box-shadow:
    0 8px 24px rgba(217, 119, 6, 0.3),
    0 2px 4px rgba(0, 0, 0, 0.1);
  border-bottom-right-radius: 6px;
}

.message-item.user-message .message-content:hover {
  box-shadow:
    0 12px 32px rgba(217, 119, 6, 0.4),
    0 4px 8px rgba(0, 0, 0, 0.15);
  transform: translateY(-1px);
}

/* AI 消息 - 粉色毛玻璃气泡（HOK设计） */
.message-item.assistant-message .message-content {
  background: rgba(219, 39, 119, 0.08);
  color: var(--color-text-primary, #f8fafc);
  box-shadow:
    0 4px 16px rgba(0, 0, 0, 0.2),
    0 1px 3px rgba(0, 0, 0, 0.1);
  border-bottom-left-radius: 6px;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(219, 39, 119, 0.15);
}

.message-item.assistant-message .message-content:hover {
  box-shadow:
    0 8px 24px rgba(219, 39, 119, 0.3),
    0 2px 6px rgba(0, 0, 0, 0.15);
  transform: translateY(-1px);
}

/* 消息文本 */
.message-text {
  white-space: pre-wrap;
}

/* 时间戳 */
.message-time {
  font-size: 11px;
  color: rgba(255, 255, 255, 0.6);
  font-weight: 500;
  padding: 0 4px;
}

.message-item.assistant-message .message-time {
  color: rgba(255, 255, 255, 0.3);
}

/* 欢迎消息 - 暗色设计 */
.welcome-message {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100%;
  padding: var(--ds-spacing-6);
  animation: fadeIn 0.6s ease-out;
}

.welcome-card {
  background: rgba(255, 255, 255, 0.04);
  backdrop-filter: blur(24px);
  border-radius: 24px;
  padding: var(--ds-spacing-8);
  max-width: 480px;
  width: 100%;
  box-shadow:
    0 20px 60px rgba(0, 0, 0, 0.3),
    0 0 0 1px rgba(255, 255, 255, 0.06);
  animation: slideUp 0.6s ease-out;
}

.welcome-header {
  text-align: center;
  margin-bottom: var(--ds-spacing-6);
}

.welcome-avatar {
  position: relative;
  width: 80px;
  height: 80px;
  margin: 0 auto var(--ds-spacing-4);
}

.avatar-ring {
  position: absolute;
  inset: 0;
  border-radius: 50%;
  background: linear-gradient(135deg, #db2777 0%, #a855f7 50%, #f472b6 100%);
  animation: rotate 3s linear infinite;
}

.avatar-ring::before {
  content: "";
  position: absolute;
  inset: 3px;
  border-radius: 50%;
  background: var(--color-bg-card, #1a1a24);
}

.avatar-inner {
  position: absolute;
  inset: 6px;
  border-radius: 50%;
  background: linear-gradient(135deg, #db2777 0%, #f472b6 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 36px;
  animation: bounce 2s ease-in-out infinite;
}

.welcome-title {
  font-size: 28px;
  font-weight: 700;
  background: linear-gradient(135deg, #db2777 0%, #fbbf24 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: var(--ds-spacing-2);
}

.welcome-subtitle {
  font-size: var(--text-base);
  color: var(--color-text-secondary, #94a3b8);
  font-weight: 500;
}

.welcome-divider {
  height: 1px;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(255, 255, 255, 0.1),
    transparent
  );
  margin: var(--ds-spacing-6) 0;
}

.welcome-description {
  font-size: var(--text-base);
  color: var(--color-text-secondary, #94a3b8);
  text-align: center;
  margin-bottom: var(--ds-spacing-5);
  font-weight: 500;
}

.suggestions {
  display: flex;
  flex-direction: column;
  gap: var(--ds-spacing-3);
  margin-bottom: var(--ds-spacing-6);
}

.suggestion-card {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-4);
  padding: var(--ds-spacing-4);
  background: rgba(255, 255, 255, 0.04);
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.06);
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.suggestion-card:hover {
  background: rgba(255, 255, 255, 0.08);
  border-color: rgba(219, 39, 119, 0.2);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.3);
  transform: translateY(-2px);
}

.suggestion-icon {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #db2777 0%, #f472b6 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
  flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(219, 39, 119, 0.3);
}

.suggestion-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.suggestion-text {
  font-size: var(--text-base);
  font-weight: 600;
  color: var(--color-text-primary, #f8fafc);
}

.suggestion-arrow {
  width: 32px;
  height: 32px;
  background: rgba(219, 39, 119, 0.1);
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--color-primary, #db2777);
  font-weight: 600;
  transition: all 0.3s ease;
}

.suggestion-card:hover .suggestion-arrow {
  background: linear-gradient(135deg, #db2777 0%, #f472b6 100%);
  color: white;
  transform: translateX(4px);
}

.welcome-footer {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--ds-spacing-3);
  padding-top: var(--ds-spacing-4);
  border-top: 1px solid rgba(255, 255, 255, 0.06);
}

.typing-indicator {
  display: flex;
  gap: 4px;
}

.typing-indicator .dot {
  width: 6px;
  height: 6px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  animation: typing 1.4s ease-in-out infinite;
}

.typing-indicator .dot:nth-child(1) {
  animation-delay: 0s;
}
.typing-indicator .dot:nth-child(2) {
  animation-delay: 0.2s;
}
.typing-indicator .dot:nth-child(3) {
  animation-delay: 0.4s;
}

.footer-text {
  font-size: var(--ds-text-size-sm);
  color: var(--ds-color-text-tertiary);
}

/* 输入区域 - 全新设计 */
.input-area {
  padding: var(--ds-spacing-4) var(--ds-spacing-6);
  background: linear-gradient(
    180deg,
    rgba(255, 255, 255, 0) 0%,
    rgba(255, 255, 255, 0.9) 20%,
    white 100%
  );
  position: relative;
}

.input-area::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(102, 126, 234, 0.2),
    transparent
  );
}

.input-wrapper {
  position: relative;
  max-width: 900px;
  margin: 0 auto;
  background: rgba(255, 255, 255, 0.03);
  border-radius: 9999px;
  border: 1px solid rgba(255, 255, 255, 0.06);
  padding: var(--ds-spacing-3) var(--ds-spacing-16) var(--ds-spacing-3)
    var(--ds-spacing-5);
  transition: all 250ms cubic-bezier(0.4, 0, 0.2, 1);
  display: flex;
  flex-direction: column;
  gap: var(--ds-spacing-3);
}

.input-wrapper:hover {
  border-color: rgba(255, 255, 255, 0.1);
  background: rgba(255, 255, 255, 0.05);
}

.input-wrapper:focus-within {
  border-color: rgba(219, 39, 119, 0.3);
  background: rgba(255, 255, 255, 0.05);
  box-shadow: 0 0 12px rgba(219, 39, 119, 0.08);
}

.input-wrapper textarea {
  width: 100%;
  border: none;
  background: transparent;
  resize: none;
  outline: none;
  font-size: 15px;
  line-height: 1.7;
  max-height: 140px;
  min-height: 24px;
  overflow-y: auto;
  font-family: inherit;
  color: var(--color-text-primary, #f8fafc);
}

.input-wrapper textarea::placeholder {
  color: var(--color-text-tertiary, #64748b);
  font-size: 14px;
}

.input-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.input-tips {
  display: flex;
  gap: var(--ds-spacing-2);
}

.tip-badge {
  font-size: 11px;
  color: var(--color-primary, #db2777);
  background: rgba(219, 39, 119, 0.1);
  border: 1px solid rgba(219, 39, 119, 0.2);
  padding: 3px 10px;
  border-radius: 6px;
  font-weight: 500;
  transition: all 0.2s ease;
}

.tip-badge:hover {
  background: rgba(219, 39, 119, 0.18);
  transform: translateY(-1px);
}

/* 发送按钮 - 暗色设计 */
.send-btn {
  position: absolute;
  right: 12px;
  bottom: 12px;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  border-radius: 12px;
  background: rgba(255, 255, 255, 0.06);
  border: 1px solid rgba(255, 255, 255, 0.08);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  flex-shrink: 0;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
}

.send-btn svg {
  width: 18px;
  height: 18px;
  fill: #64748b;
  transition: all 0.3s ease;
}

.send-btn.active {
  background: linear-gradient(135deg, #d97706 0%, #fbbf24 100%);
  border-color: transparent;
  box-shadow: 0 4px 15px rgba(217, 119, 6, 0.4);
}

.send-btn.active svg {
  fill: white;
  transform: translateX(1px);
}

.send-btn:hover:not(.loading) {
  transform: scale(1.05);
}

.send-btn.active:hover:not(.loading) {
  transform: scale(1.05) translateY(-2px);
  box-shadow: 0 6px 20px rgba(217, 119, 6, 0.5);
}

.send-btn:active:not(.loading) {
  transform: scale(0.95);
}

.send-btn:disabled {
  cursor: not-allowed;
  opacity: 0.6;
}

.loading-spinner {
  width: 18px;
  height: 18px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

/* 动画 */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(30px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
  }
}

@keyframes bounce {
  0%,
  100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-8px);
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

@keyframes typing {
  0%,
  60%,
  100% {
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

/* 响应式 */
@media (max-width: 768px) {
  .message-item {
    max-width: 90%;
  }

  .input-area {
    padding: var(--ds-spacing-3) var(--ds-spacing-4);
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

  .welcome-card {
    padding: var(--ds-spacing-5);
    border-radius: 20px;
  }

  .welcome-avatar {
    width: 64px;
    height: 64px;
  }

  .avatar-inner {
    font-size: 28px;
  }

  .welcome-title {
    font-size: 22px;
  }

  .suggestion-card {
    padding: var(--ds-spacing-3);
  }

  .suggestion-icon {
    width: 40px;
    height: 40px;
    font-size: 20px;
  }
}

@media (max-width: 480px) {
  .input-tips {
    display: none;
  }

  .input-wrapper {
    padding: 12px 50px 12px 14px;
  }

  .send-btn {
    width: 36px;
    height: 36px;
    right: 8px;
    bottom: 8px;
  }

  .send-icon {
    width: 16px;
    height: 16px;
  }
}
</style>
