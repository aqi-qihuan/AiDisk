<template>
  <button
    :class="[
      'ds-btn',
      `ds-btn-${variant}`,
      `ds-btn-${size}`,
      {
        'ds-btn-disabled': disabled,
        'ds-btn-loading': loading,
        'ds-btn-block': block,
        'ds-btn-ghost': ghost,
        'ds-btn-rounded': rounded,
      },
    ]"
    :disabled="disabled || loading"
    @click="handleClick"
  >
    <span v-if="loading" class="ds-btn__loading">
      <span class="ds-btn__loading-icon"></span>
    </span>
    <span v-if="$slots.icon" class="ds-btn__icon">
      <slot name="icon"></slot>
    </span>
    <span class="ds-btn__content">
      <slot></slot>
    </span>
  </button>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from "vue";

interface Props {
  variant?:
    | "primary"
    | "secondary"
    | "cta"
    | "golden"
    | "danger"
    | "ghost"
    | "success";
  size?: "small" | "medium" | "large";
  disabled?: boolean;
  loading?: boolean;
  block?: boolean;
  ghost?: boolean;
  rounded?: boolean;
  type?: "button" | "submit" | "reset";
}

const props = withDefaults(defineProps<Props>(), {
  variant: "primary",
  size: "medium",
  disabled: false,
  loading: false,
  block: false,
  ghost: false,
  rounded: false,
  type: "button",
});

const emit = defineEmits<{
  (e: "click", event: MouseEvent): void;
}>();

const handleClick = (event: MouseEvent) => {
  if (!props.disabled && !props.loading) {
    emit("click", event);
  }
};
</script>

<style scoped>
.ds-btn {
  position: relative;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 0 24px;
  border: none;
  border-radius: var(--radius-md);
  font-family: var(--font-primary);
  font-size: 16px;
  font-weight: 500;
  line-height: 1.5;
  cursor: pointer;
  transition: all var(--transition-fast);
  white-space: nowrap;
  outline: none;
  box-sizing: border-box;
  overflow: hidden;
}

/* Sizes */
.ds-btn-small {
  height: 32px;
  padding: 0 16px;
  font-size: 14px;
}

.ds-btn-medium {
  height: 40px;
  padding: 0 24px;
  font-size: 16px;
}

.ds-btn-large {
  height: 48px;
  padding: 0 32px;
  font-size: 18px;
}

/* ====================
   Variants - HOK Dark Theme
   ==================== */

/* Primary - Pink Gradient + Glow */
.ds-btn-primary {
  background: linear-gradient(
    135deg,
    var(--color-primary) 0%,
    var(--color-secondary) 100%
  );
  color: #ffffff;
  box-shadow: var(--glow-primary-btn);
}

.ds-btn-primary:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 30px rgba(219, 39, 119, 0.5);
}

.ds-btn-primary:active:not(:disabled) {
  transform: translateY(0) scale(0.98);
  box-shadow: 0 2px 12px rgba(219, 39, 119, 0.3);
}

/* Secondary - Dark elevated */
.ds-btn-secondary {
  background: var(--color-bg-elevated);
  color: var(--color-text-primary);
  border: 1px solid var(--color-border);
}

.ds-btn-secondary:hover:not(:disabled) {
  border-color: var(--color-primary);
  background: rgba(219, 39, 119, 0.08);
  transform: translateY(-1px);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.3);
}

.ds-btn-secondary:active:not(:disabled) {
  transform: translateY(0);
}

/* CTA - Gold Gradient + Glow */
.ds-btn-cta {
  background: linear-gradient(
    135deg,
    var(--color-cta) 0%,
    var(--color-cta-light) 50%,
    var(--color-gold) 100%
  );
  color: #1a1a24;
  font-weight: 600;
  box-shadow: var(--glow-cta-btn);
}

.ds-btn-cta:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 30px rgba(217, 119, 6, 0.5);
}

.ds-btn-cta:active:not(:disabled) {
  transform: translateY(0) scale(0.98);
  box-shadow: 0 2px 12px rgba(217, 119, 6, 0.3);
}

/* Golden - Gold bg, dark text */
.ds-btn-golden {
  background: linear-gradient(
    135deg,
    var(--color-gold) 0%,
    var(--color-cta-light) 100%
  );
  color: #0b0b10;
  font-weight: 600;
  box-shadow: 0 4px 20px rgba(251, 191, 36, 0.3);
}

.ds-btn-golden:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 6px 30px rgba(251, 191, 36, 0.5);
}

.ds-btn-golden:active:not(:disabled) {
  transform: translateY(0) scale(0.98);
}

/* Danger */
.ds-btn-danger {
  background: linear-gradient(135deg, var(--color-error) 0%, #f87171 100%);
  color: #ffffff;
  box-shadow: 0 4px 16px rgba(239, 68, 68, 0.3);
}

.ds-btn-danger:hover:not(:disabled) {
  transform: translateY(-1px);
  box-shadow: 0 6px 24px rgba(239, 68, 68, 0.4);
}

.ds-btn-danger:active:not(:disabled) {
  transform: translateY(0) scale(0.98);
}

/* Ghost - Dark theme outline */
.ds-btn-ghost {
  background: transparent;
  color: var(--color-primary);
  border: 1px solid rgba(219, 39, 119, 0.3);
}

.ds-btn-ghost:hover:not(:disabled) {
  background: rgba(219, 39, 119, 0.1);
  border-color: var(--color-primary);
  transform: translateY(-1px);
}

.ds-btn-ghost:active:not(:disabled) {
  transform: translateY(0);
}

/* Success */
.ds-btn-success {
  background: linear-gradient(135deg, var(--color-success) 0%, #34d399 100%);
  color: #ffffff;
  box-shadow: 0 4px 16px rgba(16, 185, 129, 0.3);
}

.ds-btn-success:hover:not(:disabled) {
  transform: translateY(-1px);
  box-shadow: 0 6px 24px rgba(16, 185, 129, 0.4);
}

.ds-btn-success:active:not(:disabled) {
  transform: translateY(0) scale(0.98);
}

/* Disabled State */
.ds-btn-disabled {
  opacity: 0.4;
  cursor: not-allowed;
  transform: none !important;
  box-shadow: none !important;
  background: var(--color-bg-elevated) !important;
  color: var(--color-text-tertiary) !important;
  border-color: transparent !important;
}

/* Loading State */
.ds-btn-loading {
  pointer-events: none;
}

.ds-btn__loading {
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.ds-btn__loading-icon {
  width: 14px;
  height: 14px;
  border: 2px solid currentColor;
  border-top-color: transparent;
  border-radius: 50%;
  animation: ds-spin 0.6s linear infinite;
}

@keyframes ds-spin {
  to {
    transform: rotate(360deg);
  }
}

/* Block Button */
.ds-btn-block {
  width: 100%;
}

/* Rounded Button */
.ds-btn-rounded {
  border-radius: 9999px;
}

/* Icon */
.ds-btn__icon {
  display: inline-flex;
  align-items: center;
}

.ds-btn__content {
  display: inline-flex;
  align-items: center;
}
</style>
