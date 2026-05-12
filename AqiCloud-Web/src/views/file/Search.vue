<template>
  <div class="search-container">
    <!-- 消息列表区域 -->
    <div class="message-list" ref="messageList">
      <div
        v-for="(message, index) in messages"
        :key="index"
        :class="[
          'message-item',
          message.type === 'question' ? 'user-message' : 'assistant-message',
        ]"
      >
        <!-- 头像 -->
        <div class="avatar">
          <span v-if="message.type === 'question'"
            ><svg width="20" height="20" viewBox="0 0 24 24" fill="none">
              <circle
                cx="12"
                cy="8"
                r="4"
                stroke="#DB2777"
                stroke-width="1.5"
              />
              <path
                d="M20 21a8 8 0 10-16 0"
                stroke="#A855F7"
                stroke-width="1.5"
                stroke-linecap="round"
              /></svg
          ></span>
          <span v-else
            ><svg width="20" height="20" viewBox="0 0 24 24" fill="none">
              <rect
                x="3"
                y="11"
                width="18"
                height="10"
                rx="2"
                stroke="#A855F7"
                stroke-width="1.5"
              />
              <circle cx="9" cy="15" r="1" fill="#DB2777" />
              <circle cx="15" cy="15" r="1" fill="#DB2777" />
              <path
                d="M8 7V4a4 4 0 018 0v3"
                stroke="#DB2777"
                stroke-width="1.5"
              /></svg
          ></span>
        </div>

        <!-- 消息内容 -->
        <div class="message-content-wrapper">
          <div class="message-content" v-html="message.content"></div>
        </div>
      </div>

      <!-- 空状态提示 -->
      <div v-if="messages.length === 0" class="welcome-message">
        <div class="welcome-icon">
          <svg
            width="64"
            height="64"
            viewBox="0 0 24 24"
            fill="none"
            stroke="url(#searchGrad)"
            stroke-width="1.5"
          >
            <defs>
              <linearGradient
                id="searchGrad"
                x1="0%"
                y1="0%"
                x2="100%"
                y2="100%"
              >
                <stop offset="0%" stop-color="#DB2777" />
                <stop offset="100%" stop-color="#A855F7" />
              </linearGradient>
            </defs>
            <circle cx="11" cy="11" r="8" />
            <path d="M21 21l-4.35-4.35" stroke-linecap="round" />
          </svg>
        </div>
        <h3 class="welcome-title">AI 智能搜索</h3>
        <p class="welcome-description">输入问题，AI 为您智能解答</p>
        <div class="suggestions">
          <DSButton
            v-for="(suggestion, idx) in suggestions"
            :key="idx"
            variant="outline"
            size="sm"
            @click="question = suggestion"
          >
            {{ suggestion }}
          </DSButton>
        </div>
      </div>
    </div>

    <!-- 输入区域 -->
    <div class="input-area">
      <div class="input-wrapper ds-card">
        <input
          v-model="question"
          placeholder="请输入你的问题... (Enter 发送)"
          @keyup.enter.exact.prevent="submitQuestion"
          class="ds-input"
          ref="inputRef"
        />

        <div class="input-actions">
          <DSTag color="info" size="sm">Enter 发送</DSTag>

          <DSButton
            variant="primary"
            size="sm"
            @click="submitQuestion"
            :loading="isLoading"
            :disabled="!question.trim()"
          >
            提问
          </DSButton>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, nextTick } from "vue";
import DSButton from "@/components/design-system/DSButton.vue";
import DSTag from "@/components/design-system/DSTag.vue";

// 存储问答消息
const messages = ref<{ type: "question" | "answer"; content: string }[]>([]);
// 用户输入的问题
const question = ref("");
// 加载状态
const isLoading = ref(false);
// 消息列表引用
const messageList = ref<HTMLElement | null>(null);
// 输入框引用
const inputRef = ref<HTMLInputElement | null>(null);

// 建议问题
const suggestions = ref([
  "什么是人工智能？",
  "如何学习编程？",
  "云存储的优势是什么？",
]);

// 滚动到底部
const scrollToBottom = () => {
  nextTick(() => {
    if (messageList.value) {
      messageList.value.scrollTop = messageList.value.scrollHeight;
    }
  });
};

// 提交问题
const submitQuestion = async () => {
  if (question.value.trim() === "" || isLoading.value) return;

  const userQuestion = question.value;

  // 添加用户问题到消息列表
  messages.value.push({ type: "question", content: userQuestion });
  question.value = "";

  // 添加加载中的答案
  messages.value.push({
    type: "answer",
    content: '<div class="loading-dots">思考中...</div>',
  });

  isLoading.value = true;
  scrollToBottom();

  try {
    const response = await fetch("http://127.0.0.1:8081/ai/chat", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ question: userQuestion }),
    });

    if (!response.body) throw new Error("Response stream not available.");

    const reader = response.body.getReader();
    const decoder = new TextDecoder("utf-8");

    let answerContent = "";

    while (true) {
      const { done, value } = await reader.read();
      if (done) break;

      const chunk = decoder.decode(value, { stream: true });
      answerContent += chunk;
      messages.value[messages.value.length - 1].content = answerContent;
      scrollToBottom();
    }
  } catch (error) {
    console.error("请求AI接口出错:", error);
    messages.value[messages.value.length - 1].content =
      "❌ 请求出错，请稍后重试。";
  } finally {
    isLoading.value = false;
  }

  scrollToBottom();
};
</script>

<style scoped>
.search-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background: linear-gradient(135deg, #db2777 0%, #a855f7 100%);
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
  max-width: 80%;
  animation: fadeInUp 0.3s ease-out;
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
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  flex-shrink: 0;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

/* 消息内容包裹器 */
.message-content-wrapper {
  display: flex;
  flex-direction: column;
  gap: var(--ds-spacing-1);
  max-width: 100%;
}

/* 消息内容 */
.message-content {
  padding: var(--ds-spacing-3) var(--ds-spacing-4);
  border-radius: var(--ds-radius-lg);
  font-size: var(--ds-text-size-base);
  line-height: 1.6;
  word-break: break-word;
  min-height: 40px;
}

.message-item.user-message .message-content {
  background: linear-gradient(135deg, #db2777 0%, #a855f7 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(219, 39, 119, 0.3);
}

.message-item.assistant-message .message-content {
  background: white;
  color: var(--ds-color-text-primary);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}

/* 欢迎消息 */
.welcome-message {
  text-align: center;
  padding: var(--ds-spacing-8) var(--ds-spacing-4);
  animation: fadeIn 0.5s ease-out;
}

.welcome-icon {
  font-size: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: var(--ds-spacing-4);
  animation: bounce 2s infinite;
}

.welcome-title {
  font-size: var(--ds-text-size-xl);
  font-weight: 600;
  color: white;
  margin-bottom: var(--ds-spacing-2);
}

.welcome-description {
  font-size: var(--ds-text-size-sm);
  color: rgba(255, 255, 255, 0.8);
  margin-bottom: var(--ds-spacing-6);
}

.suggestions {
  display: flex;
  gap: var(--ds-spacing-2);
  flex-wrap: wrap;
  justify-content: center;
}

/* 输入区域 */
.input-area {
  padding: var(--ds-spacing-4);
  background: white;
  box-shadow: 0 -4px 20px rgba(0, 0, 0, 0.1);
}

.input-wrapper {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-3);
  padding: var(--ds-spacing-2) var(--ds-spacing-4);
  border: 2px solid var(--ds-color-border);
  transition: all 0.3s ease;
}

.input-wrapper:focus-within {
  border-color: var(--ds-color-primary);
  box-shadow: 0 0 0 3px rgba(219, 39, 119, 0.1);
}

.ds-input {
  flex: 1;
  border: none;
  background: transparent;
  font-size: var(--ds-text-size-base);
  color: var(--ds-color-text-primary);
  outline: none;
  padding: var(--ds-spacing-2) 0;
}

.ds-input::placeholder {
  color: var(--ds-color-text-tertiary);
}

.input-actions {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-2);
  flex-shrink: 0;
}

/* 加载动画 */
.loading-dots {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-1);
  color: var(--ds-color-text-tertiary);
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

@keyframes bounce {
  0%,
  100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
}

/* 滚动条 */
::v-deep ::-webkit-scrollbar {
  width: 6px;
  height: 6px;
}

::v-deep ::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.1);
  border-radius: 3px;
}

::v-deep ::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.3);
  border-radius: 3px;
}

::v-deep ::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0.5);
}

/* 响应式 */
@media (max-width: 640px) {
  .message-item {
    max-width: 90%;
  }

  .welcome-icon {
    font-size: 48px;
  }

  .suggestions {
    flex-direction: column;
  }

  .suggestions .ds-button {
    width: 100%;
  }

  .input-wrapper {
    flex-wrap: wrap;
  }

  .ds-input {
    width: 100%;
  }

  .input-actions {
    width: 100%;
    justify-content: space-between;
  }
}
</style>
