<template>
  <div class="ai-chat-view">
    <!-- Header -->
    <div class="chat-header">
      <h1 class="page-title">AI 聊天助手</h1>
      <DSButton variant="ghost" size="sm" @click="clearChat">
        <svg
          width="16"
          height="16"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <path
            d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4V4a2 2 0 00-2-2h-.586a1 1 0 01-.707-.293l-1-1A1 1 0 003.586 2H8a2 2 0 012 2v2m5 4h.01M12 12h.01"
          />
        </svg>
        清空对话
      </DSButton>
    </div>

    <!-- Chat Messages -->
    <div class="chat-messages" ref="messagesContainer">
      <div
        v-for="(msg, index) in messages"
        :key="index"
        :class="['message', msg.role]"
      >
        <div class="message-avatar">
          <div :class="['avatar', msg.role]">
            {{ msg.role === "user" ? "我" : "AI" }}
          </div>
        </div>
        <div class="message-content">
          <div class="message-role">
            {{ msg.role === "user" ? "我" : "AI 助手" }}
          </div>
          <div class="message-text">{{ msg.content }}</div>
          <div class="message-time">{{ msg.time }}</div>
        </div>
      </div>

      <!-- Loading -->
      <div v-if="isLoading" class="message assistant">
        <div class="message-avatar">
          <div class="avatar assistant">AI</div>
        </div>
        <div class="message-content">
          <div class="loading-dots">
            <span></span>
            <span></span>
            <span></span>
          </div>
        </div>
      </div>
    </div>

    <!-- Input Area -->
    <div class="chat-input-area">
      <div class="input-wrapper">
        <textarea
          v-model="userInput"
          placeholder="输入消息，Enter 发送，Shift+Enter 换行..."
          class="chat-input"
          rows="1"
          @keydown="handleKeydown"
          @input="autoResize"
        ></textarea>
        <DSButton
          variant="primary"
          size="md"
          :disabled="!userInput.trim() || isLoading"
          @click="sendMessage"
          class="send-button"
        >
          <svg
            width="18"
            height="18"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
          >
            <path d="M22 2L11 13M22 2l-7 20-4-9-9-4 20-7z" />
          </svg>
        </DSButton>
      </div>
      <div class="input-hint">Enter 发送，Shift + Enter 换行</div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, nextTick } from "vue";
import DSButton from "@/components/design-system/DSButton.vue";

interface Message {
  role: "user" | "assistant";
  content: string;
  time: string;
}

// Mock data
const messages = ref<Message[]>([
  {
    role: "assistant",
    content:
      "你好！我是 AI 助手，可以帮你查询文件、整理资料、回答问题。有什么可以帮你的吗？",
    time: "14:30",
  },
]);

const userInput = ref("");
const isLoading = ref(false);
const messagesContainer = ref<HTMLElement | null>(null);

const sendMessage = async () => {
  if (!userInput.value.trim() || isLoading.value) return;

  // Add user message
  messages.value.push({
    role: "user",
    content: userInput.value,
    time: new Date().toLocaleTimeString("zh-CN", {
      hour: "2-digit",
      minute: "2-digit",
    }),
  });

  const input = userInput.value;
  userInput.value = "";
  isLoading.value = true;

  // Scroll to bottom
  await nextTick();
  scrollToBottom();

  // Simulate AI response
  setTimeout(() => {
    messages.value.push({
      role: "assistant",
      content: `我理解你的问题："${input}"。作为 AI 助手，我会帮你分析并给出建议。`,
      time: new Date().toLocaleTimeString("zh-CN", {
        hour: "2-digit",
        minute: "2-digit",
      }),
    });
    isLoading.value = false;

    nextTick(() => {
      scrollToBottom();
    });
  }, 1500);
};

const clearChat = () => {
  messages.value = [];
};

const handleKeydown = (event: KeyboardEvent) => {
  if (event.key === "Enter" && !event.shiftKey) {
    event.preventDefault();
    sendMessage();
  }
};

const autoResize = (event: Event) => {
  const textarea = event.target as HTMLTextAreaElement;
  textarea.style.height = "auto";
  textarea.style.height = Math.min(textarea.scrollHeight, 200) + "px";
};

const scrollToBottom = () => {
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight;
  }
};
</script>

<style scoped>
.ai-chat-view {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background: var(--color-gray-50, #f9fafb);
}

/* Header */
.chat-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 20px 32px;
  background: white;
  border-bottom: 1px solid var(--color-gray-200, #e5e7eb);
}

.page-title {
  font-size: 20px;
  font-weight: 600;
  color: var(--color-gray-900, #111827);
  margin: 0;
}

/* Chat Messages */
.chat-messages {
  flex: 1;
  overflow-y: auto;
  padding: 32px;
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.message {
  display: flex;
  gap: 12px;
  max-width: 80%;
}

.message.user {
  flex-direction: row-reverse;
  align-self: flex-end;
}

.message-avatar {
  flex-shrink: 0;
}

.avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 600;
}

.avatar.user {
  background: var(--color-primary, #3b82f6);
  color: white;
}

.avatar.assistant {
  background: var(--color-gray-100, #f3f4f6);
  color: var(--color-gray-700, #374151);
  border: 1px solid var(--color-gray-200, #e5e7eb);
}

.message-content {
  flex: 1;
}

.message.user .message-content {
  text-align: right;
}

.message-role {
  font-size: 12px;
  font-weight: 500;
  color: var(--color-gray-500, #6b7280);
  margin-bottom: 4px;
}

.message-text {
  padding: 12px 16px;
  border-radius: 12px;
  font-size: 14px;
  line-height: 1.6;
}

.message.user .message-text {
  background: var(--color-primary, #3b82f6);
  color: white;
  border-bottom-right-radius: 4px;
}

.message.assistant .message-text {
  background: white;
  color: var(--color-gray-900, #111827);
  border: 1px solid var(--color-gray-200, #e5e7eb);
  border-bottom-left-radius: 4px;
}

.message-time {
  font-size: 11px;
  color: var(--color-gray-400, #9ca3af);
  margin-top: 4px;
}

/* Loading Dots */
.loading-dots {
  display: flex;
  gap: 4px;
  padding: 12px 16px;
  background: white;
  border: 1px solid var(--color-gray-200, #e5e7eb);
  border-radius: 12px;
  border-bottom-left-radius: 4px;
}

.loading-dots span {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: var(--color-gray-400, #9ca3af);
  animation: bounce 1.4s infinite ease-in-out both;
}

.loading-dots span:nth-child(1) {
  animation-delay: -0.32s;
}
.loading-dots span:nth-child(2) {
  animation-delay: -0.16s;
}

@keyframes bounce {
  0%,
  80%,
  100% {
    transform: scale(0.6);
    opacity: 0.5;
  }
  40% {
    transform: scale(1);
    opacity: 1;
  }
}

/* Input Area */
.chat-input-area {
  padding: 20px 32px;
  background: white;
  border-top: 1px solid var(--color-gray-200, #e5e7eb);
}

.input-wrapper {
  display: flex;
  gap: 12px;
  align-items: flex-end;
}

.chat-input {
  flex: 1;
  padding: 12px 16px;
  border: 1px solid var(--color-gray-200, #e5e7eb);
  border-radius: 8px;
  font-size: 14px;
  font-family: "Inter", sans-serif;
  line-height: 1.5;
  resize: none;
  overflow-y: auto;
  max-height: 200px;
  transition: border-color 0.15s ease;
  box-sizing: border-box;
}

.chat-input:focus {
  outline: none;
  border-color: var(--color-primary, #3b82f6);
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.send-button {
  flex-shrink: 0;
}

.input-hint {
  font-size: 12px;
  color: var(--color-gray-400, #9ca3af);
  margin-top: 8px;
}

/* Responsive */
@media (max-width: 768px) {
  .chat-header {
    padding: 16px;
  }

  .chat-messages {
    padding: 16px;
    gap: 16px;
  }

  .message {
    max-width: 90%;
  }

  .chat-input-area {
    padding: 16px;
  }

  .page-title {
    font-size: 18px;
  }
}
</style>
