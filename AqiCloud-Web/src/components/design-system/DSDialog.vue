<template>
  <Teleport to="body">
    <Transition name="ds-dialog">
      <div
        v-if="modelValue"
        class="ds-dialog-overlay"
        @click.self="handleOverlayClick"
      >
        <div
          :class="[
            'ds-dialog',
            `ds-dialog-${size}`,
            {
              'ds-dialog-fullscreen': effectiveFullscreen,
              'ds-dialog-centered': centered && !effectiveFullscreen,
            },
          ]"
          :style="dialogStyle"
        >
          <!-- Header -->
          <div v-if="showHeader" class="ds-dialog-header">
            <slot name="header">
              <h3 class="ds-dialog-title">{{ title }}</h3>
            </slot>
            <button
              v-if="showClose"
              class="ds-dialog-close"
              @click="handleClose"
            >
              <svg viewBox="0 0 24 24" width="16" height="16">
                <path
                  fill="currentColor"
                  d="M19 6.41L17.59 5 12 10.59 6.41 5 5 6.41 10.59 12 5 17.59 6.41 19 12 13.41 17.59 19 19 17.59 13.41 12z"
                />
              </svg>
            </button>
          </div>

          <!-- Body -->
          <div class="ds-dialog-body" :style="bodyStyle">
            <slot></slot>
          </div>

          <!-- Footer -->
          <div v-if="showFooter" class="ds-dialog-footer">
            <slot name="footer">
              <DSButton variant="secondary" @click="handleCancel">
                {{ cancelText }}
              </DSButton>
              <DSButton
                variant="primary"
                @click="handleConfirm"
                :loading="confirmLoading"
              >
                {{ confirmText }}
              </DSButton>
            </slot>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<script setup lang="ts">
import { computed, watch, onMounted, onUnmounted, ref } from "vue";
import DSButton from "./DSButton.vue";

interface Props {
  modelValue: boolean;
  title?: string;
  size?: "small" | "medium" | "large" | "fullscreen";
  width?: string | number;
  fullscreen?: boolean;
  showClose?: boolean;
  showHeader?: boolean;
  showFooter?: boolean;
  closeOnClickOverlay?: boolean;
  closeOnPressEscape?: boolean;
  confirmText?: string;
  cancelText?: string;
  confirmLoading?: boolean;
  centered?: boolean;
  /** 移动端是否自动 fullscreen (默认 true) */
  mobileFullscreen?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  title: "",
  size: "medium",
  width: "",
  fullscreen: false,
  showClose: true,
  showHeader: true,
  showFooter: false,
  closeOnClickOverlay: true,
  closeOnPressEscape: true,
  confirmText: "确定",
  cancelText: "取消",
  confirmLoading: false,
  centered: false,
  mobileFullscreen: true,
});

const emit = defineEmits<{
  "update:modelValue": [value: boolean];
  close: [];
  confirm: [];
  cancel: [];
}>();

// 移动端检测
const isMobile = ref(false);
const checkMobile = () => {
  isMobile.value = window.innerWidth <= 768;
};
onMounted(() => {
  checkMobile();
  window.addEventListener("resize", checkMobile);
});
onUnmounted(() => {
  window.removeEventListener("resize", checkMobile);
});

// 是否实际 fullscreen（移动端自动或手动 fullscreen）
const effectiveFullscreen = computed(() => {
  return props.fullscreen || (props.mobileFullscreen && isMobile.value);
});

const dialogStyle = computed(() => {
  if (effectiveFullscreen.value) return {};
  if (props.width) {
    const width =
      typeof props.width === "number" ? `${props.width}px` : props.width;
    return { width, maxWidth: "90vw" };
  }
  return {};
});

const bodyStyle = computed(() => {
  if (effectiveFullscreen.value) {
    return { maxHeight: "calc(100vh - 120px)", overflow: "auto" };
  }
  return {};
});

const handleClose = () => {
  emit("update:modelValue", false);
  emit("close");
};

const handleOverlayClick = () => {
  if (props.closeOnClickOverlay) {
    handleClose();
  }
};

const handleConfirm = () => {
  emit("confirm");
};

const handleCancel = () => {
  emit("update:modelValue", false);
  emit("cancel");
};

// ESC key handler
watch(
  () => props.modelValue,
  (val) => {
    if (val && props.closeOnPressEscape) {
      const handleEscape = (e: KeyboardEvent) => {
        if (e.key === "Escape") {
          handleClose();
        }
      };
      document.addEventListener("keydown", handleEscape);
      return () => document.removeEventListener("keydown", handleEscape);
    }
  },
);

// Lock body scroll when dialog is open
watch(
  () => props.modelValue,
  (val) => {
    if (val) {
      document.body.style.overflow = "hidden";
    } else {
      document.body.style.overflow = "";
    }
  },
);
</script>

<style scoped>
.ds-dialog-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.6);
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
  display: flex;
  align-items: flex-start;
  justify-content: center;
  padding: 10vh 20px 20px;
  z-index: 2000;
  overflow: auto;
}

.ds-dialog-overlay.ds-dialog-centered {
  align-items: center;
  padding: 20px;
}

.ds-dialog {
  background: var(--color-bg-card, #1a1a24);
  border-radius: var(--radius-xl, 16px);
  box-shadow:
    0 24px 80px rgba(0, 0, 0, 0.5),
    0 8px 32px rgba(0, 0, 0, 0.3);
  border: 1px solid rgba(255, 255, 255, 0.06);
  max-height: 90vh;
  display: flex;
  flex-direction: column;
  animation: ds-dialog-enter 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

@keyframes ds-dialog-enter {
  from {
    opacity: 0;
    transform: scale(0.9) translateY(-20px);
  }
  to {
    opacity: 1;
    transform: scale(1) translateY(0);
  }
}

/* Sizes */
.ds-dialog-small {
  width: 400px;
  max-width: 90vw;
}

.ds-dialog-medium {
  width: 560px;
  max-width: 90vw;
}

.ds-dialog-large {
  width: 720px;
  max-width: 90vw;
}

.ds-dialog-fullscreen {
  width: 100vw;
  height: 100vh;
  max-width: 100vw;
  max-height: 100vh;
  border-radius: 0;
}

/* Header - Pink Gradient */
.ds-dialog-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 20px 24px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
  flex-shrink: 0;
  background: linear-gradient(
    135deg,
    var(--color-primary) 0%,
    var(--color-secondary) 50%,
    var(--color-gold-light) 100%
  );
  border-radius: var(--radius-xl) var(--radius-xl) 0 0;
  position: relative;
  overflow: hidden;
}

.ds-dialog-header::after {
  content: "";
  position: absolute;
  top: -50%;
  right: -20%;
  width: 200px;
  height: 200px;
  background: rgba(255, 255, 255, 0.08);
  border-radius: 50%;
  filter: blur(40px);
  pointer-events: none;
}

.ds-dialog-title {
  margin: 0;
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
  line-height: 1.4;
  position: relative;
  z-index: 1;
}

.ds-dialog-close {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  padding: 0;
  border: none;
  background: rgba(255, 255, 255, 0.1);
  color: rgba(255, 255, 255, 0.8);
  cursor: pointer;
  border-radius: 8px;
  transition: all 0.2s ease;
  position: relative;
  z-index: 1;
  backdrop-filter: blur(4px);
}

.ds-dialog-close:hover {
  background: rgba(255, 255, 255, 0.2);
  color: #ffffff;
  transform: rotate(90deg);
}

/* Body */
.ds-dialog-body {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
  color: var(--color-text-primary);
}

/* Footer */
.ds-dialog-footer {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 12px;
  padding: 16px 24px;
  border-top: 1px solid rgba(255, 255, 255, 0.06);
  background: var(--color-bg-elevated, #22222e);
  border-radius: 0 0 var(--radius-xl) var(--radius-xl);
  flex-shrink: 0;
}

.ds-dialog-fullscreen .ds-dialog-footer {
  border-radius: 0;
}

/* Transition */
.ds-dialog-enter-active,
.ds-dialog-leave-active {
  transition: opacity 0.3s ease;
}

.ds-dialog-enter-active .ds-dialog,
.ds-dialog-leave-active .ds-dialog {
  transition:
    transform 0.3s ease,
    opacity 0.3s ease;
}

.ds-dialog-enter-from,
.ds-dialog-leave-to {
  opacity: 0;
}

.ds-dialog-enter-from .ds-dialog,
.ds-dialog-leave-to .ds-dialog {
  opacity: 0;
  transform: scale(0.9) translateY(-20px);
}

/* Mobile */
@media (max-width: 768px) {
  .ds-dialog-overlay {
    padding: 0;
    align-items: flex-end;
  }

  .ds-dialog {
    width: 100% !important;
    max-width: 100%;
    border-radius: var(--radius-2xl, 24px) var(--radius-2xl, 24px) 0 0;
    max-height: 85vh;
  }

  .ds-dialog.ds-dialog-centered {
    margin: auto;
    align-self: center;
    border-radius: var(--radius-xl, 16px);
  }

  .ds-dialog-header {
    padding: 16px 20px;
    border-radius: var(--radius-2xl, 24px) var(--radius-2xl, 24px) 0 0;
  }

  .ds-dialog-body {
    padding: 20px;
  }

  .ds-dialog-footer {
    padding: 12px 20px;
    flex-wrap: wrap;
    border-radius: 0;
  }

  .ds-dialog-footer :deep(.ds-btn) {
    flex: 1;
    min-width: 120px;
  }
}

/* Small Mobile */
@media (max-width: 576px) {
  .ds-dialog-body {
    padding: 16px;
  }

  .ds-dialog-header {
    padding: 14px 16px;
  }

  .ds-dialog-title {
    font-size: 16px;
  }
}
</style>
