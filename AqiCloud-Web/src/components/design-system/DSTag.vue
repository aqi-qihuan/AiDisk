<template>
  <span
    :class="[
      'ds-tag',
      `ds-tag-${variant}`,
      `ds-tag-${size}`,
      {
        'ds-tag-closeable': closeable,
        'ds-tag-rounded': rounded,
      },
    ]"
  >
    <slot></slot>
    <span v-if="closeable" class="ds-tag__close" @click="handleClose">×</span>
  </span>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from "vue";

interface Props {
  variant?:
    | "primary"
    | "secondary"
    | "success"
    | "warning"
    | "danger"
    | "info"
    | "gold";
  size?: "small" | "medium" | "large";
  closeable?: boolean;
  rounded?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  variant: "primary",
  size: "medium",
  closeable: false,
  rounded: true,
});

const emit = defineEmits<{
  (e: "close"): void;
}>();

const handleClose = (event: MouseEvent) => {
  event.stopPropagation();
  emit("close");
};
</script>

<style scoped>
.ds-tag {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  padding: 4px 12px;
  border-radius: 9999px;
  font-size: 12px;
  font-weight: 500;
  white-space: nowrap;
  transition: all 0.2s ease;
  cursor: default;
}

/* Sizes */
.ds-tag-small {
  height: 20px;
  font-size: 11px;
  padding: 0 8px;
}

.ds-tag-medium {
  height: 24px;
  font-size: 12px;
  padding: 4px 12px;
}

.ds-tag-large {
  height: 28px;
  font-size: 13px;
  padding: 6px 16px;
}

/* Variants */
.ds-tag-primary {
  background-color: rgba(245, 158, 11, 0.1);
  color: #f59e0b;
}

.ds-tag-secondary {
  background-color: rgba(217, 119, 6, 0.1);
  color: #d97706;
}

.ds-tag-success {
  background-color: rgba(16, 185, 129, 0.1);
  color: #10b981;
}

.ds-tag-warning {
  background-color: rgba(245, 158, 11, 0.1);
  color: #f59e0b;
}

.ds-tag-danger {
  background-color: rgba(239, 68, 68, 0.1);
  color: #ef4444;
}

.ds-tag-info {
  background-color: rgba(59, 130, 246, 0.1);
  color: #3b82f6;
}

/* Tag Gold — HOK 金色标签 */
.ds-tag-gold {
  background: rgba(217, 119, 6, 0.12);
  color: var(--color-gold, #fbbf24);
  border: 1px solid rgba(217, 119, 6, 0.15);
}

/* Rounded */
.ds-tag-rounded {
  border-radius: 9999px;
}

/* Closeable */
.ds-tag-closeable {
  padding-right: 8px;
  cursor: pointer;
}

.ds-tag__close {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 16px;
  height: 16px;
  font-size: 18px;
  line-height: 1;
  opacity: 0.6;
  transition: opacity 0.2s ease;
}

.ds-tag-closeable:hover .ds-tag__close {
  opacity: 1;
}
</style>
