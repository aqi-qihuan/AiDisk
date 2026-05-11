<template>
  <input
    :type="type"
    :value="modelValue"
    :placeholder="placeholder"
    :disabled="disabled"
    :class="['ds-input', `ds-input--${size}`, { 'ds-input--error': error }]"
    @input="onInput"
    @focus="onFocus"
    @blur="onBlur"
  />
</template>

<script setup lang="ts">
import { computed } from 'vue';

interface Props {
  modelValue?: string;
  type?: 'text' | 'email' | 'password' | 'number';
  placeholder?: string;
  size?: 'sm' | 'md' | 'lg';
  disabled?: boolean;
  error?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: '',
  type: 'text',
  placeholder: '',
  size: 'md',
  disabled: false,
  error: false,
});

const emit = defineEmits<{
  'update:modelValue': [value: string];
  focus: [event: FocusEvent];
  blur: [event: FocusEvent];
}>();

const onInput = (event: Event) => {
  const target = event.target as HTMLInputElement;
  emit('update:modelValue', target.value);
};

const onFocus = (event: FocusEvent) => {
  emit('focus', event);
};

const onBlur = (event: FocusEvent) => {
  emit('blur', event);
};
</script>

<style scoped>
.ds-input {
  width: 100%;
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
  font-family: var(--font-primary);
  color: var(--color-text-primary);
  background: var(--color-surface);
  transition: border-color var(--transition-base), box-shadow var(--transition-base);
  box-sizing: border-box;
}

.ds-input:focus {
  outline: none;
  border-color: var(--color-primary);
  box-shadow: 0 0 0 3px rgba(124, 58, 237, 0.1);
}

.ds-input:hover:not(:disabled):not(:focus) {
  border-color: var(--color-text-secondary);
}

.ds-input:disabled {
  background: var(--color-bg);
  color: var(--color-text-tertiary);
  cursor: not-allowed;
  opacity: 0.6;
}

.ds-input--error {
  border-color: var(--color-error);
}

.ds-input--error:focus {
  box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.1);
}

/* Sizes */
.ds-input--sm {
  padding: 6px 10px;
  font-size: var(--text-sm);
}

.ds-input--md {
  padding: 8px 12px;
  font-size: var(--text-base);
}

.ds-input--lg {
  padding: 12px 16px;
  font-size: var(--text-lg);
}

/* Responsive */
@media (max-width: 768px) {
  .ds-input--lg {
    font-size: var(--text-base);
    padding: 10px 14px;
  }
}

/* Accessibility */
@media (prefers-reduced-motion: reduce) {
  .ds-input {
    transition: none;
  }
}
</style>
