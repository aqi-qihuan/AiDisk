<template>
  <div class="ds-form-item">
    <label
      v-if="label"
      class="form-item-label"
      :class="{ 'is-required': required }"
    >
      {{ label }}
    </label>
    <div class="form-item-content">
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
/**
 * DSFormItem 表单项组件
 * 提供表单中的单个表单项
 */

interface DSFormItemProps {
  label?: string;
  required?: boolean;
}

withDefaults(defineProps<DSFormItemProps>(), {
  required: false,
});
</script>

<style scoped>
.ds-form-item {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-item-label {
  font-size: 14px;
  font-weight: 500;
  color: var(--color-text-primary);
}

.form-item-label.is-required::after {
  content: " *";
  color: #ef4444;
}

.form-item-content {
  width: 100%;
}
</style>
