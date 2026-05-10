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
      }
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
import { defineProps, defineEmits } from 'vue';

interface Props {
  variant?: 'primary' | 'secondary' | 'cta' | 'danger' | 'ghost' | 'success';
  size?: 'small' | 'medium' | 'large';
  disabled?: boolean;
  loading?: boolean;
  block?: boolean;
  ghost?: boolean;
  rounded?: boolean;
  type?: 'button' | 'submit' | 'reset';
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'primary',
  size: 'medium',
  disabled: false,
  loading: false,
  block: false,
  ghost: false,
  rounded: false,
  type: 'button',
});

const emit = defineEmits<{
  (e: 'click', event: MouseEvent): void;
}>();

const handleClick = (event: MouseEvent) => {
  if (!props.disabled && !props.loading) {
    emit('click', event);
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
  border-radius: 8px;
  font-size: 16px;
  font-weight: 500;
  line-height: 1.5;
  cursor: pointer;
  transition: all 0.2s ease;
  white-space: nowrap;
  outline: none;
  box-sizing: border-box;
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

/* Variants */
.ds-btn-primary {
  background-color: #6366F1;
  color: #FFFFFF;
}

.ds-btn-primary:hover:not(:disabled) {
  background-color: #4F46E5;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.3);
}

.ds-btn-primary:active:not(:disabled) {
  transform: translateY(0);
}

.ds-btn-secondary {
  background-color: #FFFFFF;
  color: #6366F1;
  border: 1px solid #E2E8F0;
}

.ds-btn-secondary:hover:not(:disabled) {
  border-color: #6366F1;
  background-color: #F5F3FF;
  transform: translateY(-1px);
}

.ds-btn-secondary:active:not(:disabled) {
  transform: translateY(0);
}

.ds-btn-cta {
  background-color: #10B981;
  color: #FFFFFF;
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.ds-btn-cta:hover:not(:disabled) {
  background-color: #059669;
  box-shadow: 0 6px 16px rgba(16, 185, 129, 0.4);
  transform: translateY(-2px);
}

.ds-btn-cta:active:not(:disabled) {
  transform: translateY(-1px);
}

.ds-btn-danger {
  background-color: #EF4444;
  color: #FFFFFF;
}

.ds-btn-danger:hover:not(:disabled) {
  background-color: #DC2626;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(239, 68, 68, 0.3);
}

.ds-btn-ghost {
  background-color: transparent;
  color: #6366F1;
  border: 1px solid #6366F1;
}

.ds-btn-ghost:hover:not(:disabled) {
  background-color: rgba(99, 102, 241, 0.1);
}

.ds-btn-success {
  background-color: #10B981;
  color: #FFFFFF;
}

.ds-btn-success:hover:not(:disabled) {
  background-color: #059669;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(16, 185, 129, 0.3);
}

.ds-btn-success:active:not(:disabled) {
  transform: translateY(0);
}

/* Disabled State */
.ds-btn-disabled {
  opacity: 0.5;
  cursor: not-allowed;
  transform: none !important;
  box-shadow: none !important;
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
