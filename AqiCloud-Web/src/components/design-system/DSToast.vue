<template>
  <Teleport to="body">
    <Transition name="ds-toast-fade">
      <div
        v-if="visible"
        :class="['ds-toast', `ds-toast-${type}`, `ds-toast-${position}`]"
        :style="toastStyle"
        role="alert"
        :aria-live="type === 'error' ? 'assertive' : 'polite'"
      >
        <!-- Lucide 图标 -->
        <div class="ds-toast-icon">
          <Check v-if="type === 'success'" :size="20" />
          <X v-else-if="type === 'error'" :size="20" />
          <AlertTriangle v-else-if="type === 'warning'" :size="20" />
          <Info v-else-if="type === 'info'" :size="20" />
          <MessageSquare v-else :size="20" />
        </div>

        <!-- 内容 -->
        <div class="ds-toast-content">
          <div v-if="title" class="ds-toast-title">{{ title }}</div>
          <div class="ds-toast-message">{{ message }}</div>
        </div>

        <!-- 关闭按钮 -->
        <button
          v-if="closable"
          class="ds-toast-close"
          @click="close"
          aria-label="关闭提示"
        >
          <X :size="16" />
        </button>

        <!-- 进度条 -->
        <div
          v-if="duration > 0"
          class="ds-toast-progress"
          :style="{ animationDuration: `${duration}ms` }"
        ></div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from "vue";
import { Check, X, AlertTriangle, Info, MessageSquare } from "lucide-vue-next";

interface Props {
  type?: "success" | "error" | "warning" | "info" | "default";
  title?: string;
  message: string;
  duration?: number;
  closable?: boolean;
  position?:
    | "top-right"
    | "top-center"
    | "top-left"
    | "bottom-right"
    | "bottom-center"
    | "bottom-left";
  offset?: number;
}

const props = withDefaults(defineProps<Props>(), {
  type: "default",
  title: "",
  duration: 3000,
  closable: true,
  position: "top-right",
  offset: 20,
});

const emit = defineEmits<{
  (e: "close"): void;
}>();

const visible = ref(false);
let timer: number | null = null;

// 计算样式
const toastStyle = computed(() => {
  const style: Record<string, string> = {};

  if (props.position.includes("top")) {
    style.top = `${props.offset}px`;
  } else {
    style.bottom = `${props.offset}px`;
  }

  if (props.position.includes("right")) {
    style.right = `${props.offset}px`;
  } else if (props.position.includes("left")) {
    style.left = `${props.offset}px`;
  } else {
    style.left = "50%";
    style.transform = "translateX(-50%)";
  }

  return style;
});

// 显示
const show = () => {
  visible.value = true;

  if (props.duration > 0) {
    timer = window.setTimeout(() => {
      close();
    }, props.duration);
  }
};

// 关闭
const close = () => {
  visible.value = false;

  if (timer) {
    clearTimeout(timer);
    timer = null;
  }

  emit("close");
};

// 组件挂载后显示
onMounted(() => {
  show();
});

// 组件卸载前清理
onUnmounted(() => {
  if (timer) {
    clearTimeout(timer);
  }
});

// 暴露方法
defineExpose({
  close,
  show,
});
</script>

<style scoped>
.ds-toast {
  position: fixed;
  z-index: 9999;
  display: flex;
  align-items: flex-start;
  gap: 12px;
  padding: 16px 20px;
  min-width: 300px;
  max-width: 450px;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: 12px;
  box-shadow:
    0 8px 32px rgba(0, 0, 0, 0.12),
    0 2px 8px rgba(0, 0, 0, 0.04);
  border: 1px solid rgba(255, 255, 255, 0.6);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

/* 类型样式 */
.ds-toast-success {
  border-left: 4px solid #10b981;
}

.ds-toast-success .ds-toast-icon {
  color: #10b981;
}

.ds-toast-error {
  border-left: 4px solid #ef4444;
}

.ds-toast-error .ds-toast-icon {
  color: #ef4444;
}

.ds-toast-warning {
  border-left: 4px solid #f59e0b;
}

.ds-toast-warning .ds-toast-icon {
  color: #f59e0b;
}

.ds-toast-info {
  border-left: 4px solid #3b82f6;
}

.ds-toast-info .ds-toast-icon {
  color: #3b82f6;
}

.ds-toast-default {
  border-left: 4px solid #DB2777;
}

.ds-toast-default .ds-toast-icon {
  color: #DB2777;
}

/* 图标 */
.ds-toast-icon {
  font-size: 20px;
  font-weight: bold;
  flex-shrink: 0;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 内容 */
.ds-toast-content {
  flex: 1;
  min-width: 0;
}

.ds-toast-title {
  font-size: 14px;
  font-weight: 600;
  color: #1e1b4b;
  margin-bottom: 4px;
  line-height: 1.4;
}

.ds-toast-message {
  font-size: 13px;
  color: #64748b;
  line-height: 1.5;
  word-wrap: break-word;
}

/* 关闭按钮 */
.ds-toast-close {
  position: absolute;
  top: 8px;
  right: 8px;
  width: 24px;
  height: 24px;
  border: none;
  background: transparent;
  color: #94a3b8;
  font-size: 18px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 4px;
  transition: all 0.2s ease;
  padding: 0;
}

.ds-toast-close:hover {
  background: rgba(0, 0, 0, 0.05);
  color: #64748b;
}

/* 进度条 */
.ds-toast-progress {
  position: absolute;
  bottom: 0;
  left: 0;
  height: 3px;
  background: linear-gradient(90deg, #DB2777, #F472B6);
  border-radius: 0 0 12px 12px;
  animation: ds-toast-progress-shrink linear forwards;
  width: 100%;
}

@keyframes ds-toast-progress-shrink {
  from {
    width: 100%;
  }
  to {
    width: 0%;
  }
}

/* 动画 */
.ds-toast-fade-enter-active {
  animation: ds-toast-slide-in 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.ds-toast-fade-leave-active {
  animation: ds-toast-slide-out 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

@keyframes ds-toast-slide-in {
  from {
    opacity: 0;
    transform: translateX(100px) scale(0.9);
  }
  to {
    opacity: 1;
    transform: translateX(0) scale(1);
  }
}

@keyframes ds-toast-slide-out {
  from {
    opacity: 1;
    transform: translateX(0) scale(1);
  }
  to {
    opacity: 0;
    transform: translateX(100px) scale(0.9);
  }
}

/* 响应式 */
@media (max-width: 768px) {
  .ds-toast {
    min-width: auto;
    max-width: calc(100vw - 40px);
    width: calc(100vw - 40px);
    padding: 12px 16px;
  }

  .ds-toast-title {
    font-size: 13px;
  }

  .ds-toast-message {
    font-size: 12px;
  }
}

/* 无障碍访问性 */
@media (prefers-reduced-motion: reduce) {
  .ds-toast,
  .ds-toast-fade-enter-active,
  .ds-toast-fade-leave-active {
    animation: none;
    transition: opacity 0.01ms;
  }

  .ds-toast-progress {
    animation: none;
  }
}
</style>
