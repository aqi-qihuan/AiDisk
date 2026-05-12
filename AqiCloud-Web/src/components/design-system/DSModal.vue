<template>
  <Teleport to="body">
    <Transition name="ds-modal-fade">
      <div
        v-if="visible"
        class="ds-modal-overlay glass-overlay"
        @click="handleOverlayClick"
        role="dialog"
        :aria-label="title || 'Modal'"
        aria-modal="true"
      >
        <Transition name="ds-modal-scale">
          <div
            v-if="visible"
            :class="['ds-modal', `ds-modal-${size}`]"
            @click.stop
            ref="modalRef"
          >
            <!-- 头部 -->
            <div v-if="title || $slots.header" class="ds-modal-header">
              <slot name="header">
                <h3 class="ds-modal-title">{{ title }}</h3>
              </slot>
              <button
                v-if="closable"
                class="ds-modal-close"
                @click="close"
                aria-label="关闭对话框"
              >
                <X :size="20" />
              </button>
            </div>

            <!-- 内容 -->
            <div class="ds-modal-body">
              <slot />
            </div>

            <!-- 底部 -->
            <div v-if="$slots.footer" class="ds-modal-footer">
              <slot name="footer" />
            </div>
          </div>
        </Transition>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch, nextTick } from "vue";
import { X } from "lucide-vue-next";

interface Props {
  modelValue: boolean;
  title?: string;
  size?: "sm" | "md" | "lg" | "xl" | "full";
  closable?: boolean;
  maskClosable?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  title: "",
  size: "md",
  closable: true,
  maskClosable: true,
});

const emit = defineEmits<{
  (e: "update:modelValue", value: boolean): void;
  (e: "close"): void;
  (e: "open"): void;
}>();

const visible = ref(props.modelValue);
const modalRef = ref<HTMLElement | null>(null);
let previousFocusElement: HTMLElement | null = null;

// 监听 modelValue 变化
watch(
  () => props.modelValue,
  (val) => {
    visible.value = val;
    if (val) {
      open();
    } else {
      close();
    }
  },
);

// 打开模态框
const open = () => {
  // 保存当前焦点元素
  previousFocusElement = document.activeElement as HTMLElement;

  visible.value = true;
  emit("update:modelValue", true);
  emit("open");

  // 焦点陷阱
  nextTick(() => {
    trapFocus();
  });
};

// 关闭模态框
const close = () => {
  visible.value = false;
  emit("update:modelValue", false);
  emit("close");

  // 恢复焦点
  if (previousFocusElement) {
    previousFocusElement.focus();
  }
};

// 点击遮罩层关闭
const handleOverlayClick = () => {
  if (props.maskClosable) {
    close();
  }
};

// 焦点陷阱
const trapFocus = () => {
  if (!modalRef.value) return;

  const focusableElements = modalRef.value.querySelectorAll(
    'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])',
  );

  if (focusableElements.length > 0) {
    (focusableElements[0] as HTMLElement).focus();
  }
};

// 键盘事件处理
const handleKeydown = (event: KeyboardEvent) => {
  if (!visible.value) return;

  // ESC 关闭
  if (event.key === "Escape" && props.closable) {
    close();
  }

  // Tab 焦点陷阱
  if (event.key === "Tab" && modalRef.value) {
    const focusableElements = modalRef.value.querySelectorAll(
      'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])',
    );

    const firstElement = focusableElements[0] as HTMLElement;
    const lastElement = focusableElements[
      focusableElements.length - 1
    ] as HTMLElement;

    if (event.shiftKey) {
      if (document.activeElement === firstElement) {
        lastElement.focus();
        event.preventDefault();
      }
    } else {
      if (document.activeElement === lastElement) {
        firstElement.focus();
        event.preventDefault();
      }
    }
  }
};

// 组件挂载
onMounted(() => {
  document.addEventListener("keydown", handleKeydown);
  if (props.modelValue) {
    open();
  }
});

// 组件卸载
onUnmounted(() => {
  document.removeEventListener("keydown", handleKeydown);
});

// 暴露方法
defineExpose({
  open,
  close,
});
</script>

<style scoped>
.ds-modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: var(--z-modal-backdrop);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

.ds-modal {
  position: relative;
  z-index: var(--z-modal);
  background: var(--color-surface);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-xl);
  max-height: 90vh;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

/* 尺寸 */
.ds-modal-sm {
  width: 100%;
  max-width: 400px;
}

.ds-modal-md {
  width: 100%;
  max-width: 560px;
}

.ds-modal-lg {
  width: 100%;
  max-width: 800px;
}

.ds-modal-xl {
  width: 100%;
  max-width: 1140px;
}

.ds-modal-full {
  width: 100%;
  max-width: 100%;
  height: 100%;
  max-height: 100%;
  border-radius: 0;
}

/* 头部 */
.ds-modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 20px 24px;
  border-bottom: 1px solid var(--color-border);
}

.ds-modal-title {
  font-family: var(--font-heading);
  font-size: var(--text-xl);
  font-weight: var(--font-semibold);
  color: var(--color-text-primary);
  margin: 0;
}

.ds-modal-close {
  width: 32px;
  height: 32px;
  border: none;
  background: transparent;
  color: var(--color-text-secondary);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: var(--radius-md);
  transition: all var(--transition-fast);
  padding: 0;
}

.ds-modal-close:hover {
  background: rgba(0, 0, 0, 0.05);
  color: var(--color-text-primary);
}

/* 内容 */
.ds-modal-body {
  padding: 24px;
  overflow-y: auto;
  flex: 1;
}

/* 底部 */
.ds-modal-footer {
  padding: 16px 24px;
  border-top: 1px solid var(--color-border);
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

/* 动画 */
.ds-modal-fade-enter-active,
.ds-modal-fade-leave-active {
  transition: opacity 0.3s ease;
}

.ds-modal-fade-enter-from,
.ds-modal-fade-leave-to {
  opacity: 0;
}

.ds-modal-scale-enter-active,
.ds-modal-scale-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.ds-modal-scale-enter-from,
.ds-modal-scale-leave-to {
  opacity: 0;
  transform: scale(0.9);
}

/* 响应式 */
@media (max-width: 768px) {
  .ds-modal-overlay {
    padding: 10px;
  }

  .ds-modal {
    max-height: 95vh;
  }

  .ds-modal-header {
    padding: 16px;
  }

  .ds-modal-body {
    padding: 16px;
  }

  .ds-modal-footer {
    padding: 12px 16px;
  }

  .ds-modal-title {
    font-size: var(--text-lg);
  }
}

/* 无障碍访问性 */
@media (prefers-reduced-motion: reduce) {
  .ds-modal-fade-enter-active,
  .ds-modal-fade-leave-active,
  .ds-modal-scale-enter-active,
  .ds-modal-scale-leave-active {
    transition: none;
  }
}
</style>
