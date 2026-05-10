<template>
  <div
    :class="[
      'ds-card',
      {
        'ds-card-hoverable': hoverable,
        'ds-card-bordered': bordered,
        'ds-card-shadow': shadow,
      }
    ]"
  >
    <div v-if="$slots.header || title" class="ds-card-header">
      <slot name="header">
        <h3 v-if="title" class="ds-card-title">{{ title }}</h3>
      </slot>
      <div v-if="$slots.extra" class="ds-card-extra">
        <slot name="extra"></slot>
      </div>
    </div>
    <div :class="['ds-card-body', { 'ds-card-body-padding': !noPadding }]">
      <slot></slot>
    </div>
    <div v-if="$slots.footer" class="ds-card-footer">
      <slot name="footer"></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';

interface Props {
  title?: string;
  hoverable?: boolean;
  bordered?: boolean;
  shadow?: boolean;
  noPadding?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  title: '',
  hoverable: false,
  bordered: true,
  shadow: false,
  noPadding: false,
});
</script>

<style scoped>
.ds-card {
  background-color: #FFFFFF;
  border-radius: 12px;
  overflow: hidden;
  transition: all 0.2s ease;
}

/* Bordered */
.ds-card-bordered {
  border: 1px solid #E2E8F0;
}

/* Shadow */
.ds-card-shadow {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
}

/* Hoverable */
.ds-card-hoverable {
  cursor: pointer;
}

.ds-card-hoverable:hover {
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.08);
  transform: translateY(-2px);
}

/* Header */
.ds-card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 20px 24px;
  border-bottom: 1px solid #E2E8F0;
  background-color: #F8F8FF;
}

.ds-card-title {
  font-size: 20px;
  font-weight: 600;
  color: #1E1B4B;
  margin: 0;
}

.ds-card-extra {
  display: flex;
  align-items: center;
  gap: 8px;
}

/* Body */
.ds-card-body {
  padding: 24px;
}

.ds-card-body-padding {
  padding: 0;
}

/* Footer */
.ds-card-footer {
  padding: 16px 24px;
  border-top: 1px solid #E2E8F0;
  background-color: #F5F3FF;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 12px;
}
</style>
