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
import { computed } from "vue";

interface Props {
  modelValue?: string;
  type?: "text" | "email" | "password" | "number";
  placeholder?: string;
  size?: "sm" | "md" | "lg";
  disabled?: boolean;
  error?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: "",
  type: "text",
  placeholder: "",
  size: "md",
  disabled: false,
  error: false,
});

const emit = defineEmits<{
  "update:modelValue": [value: string];
  focus: [event: FocusEvent];
  blur: [event: FocusEvent];
}>();

const onInput = (event: Event) => {
  const target = event.target as HTMLInputElement;
  emit("update:modelValue", target.value);
};

const onFocus = (event: FocusEvent) => {
  emit("focus", event);
};

const onBlur = (event: FocusEvent) => {
  emit("blur", event);
};
</script>

<style scoped>
.ds-input {
  width: 100%;
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 8px 16px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-radius: 8px;
  transition: all 250ms cubic-bezier(0.4, 0, 0.2, 1);
  font-family:
    "Fira Sans",
    "Plus Jakarta Sans",
    -apple-system,
    sans-serif;
  font-size: 14px;
  outline: none;
  color: #f8fafc;
  box-sizing: border-box;
}

.ds-input:focus {
  border-color: rgba(212, 168, 83, 0.3);
  box-shadow: 0 0 12px rgba(212, 168, 83, 0.08);
  background: rgba(255, 255, 255, 0.05);
}

.ds-input:hover:not(:disabled):not(:focus) {
  border-color: rgba(255, 255, 255, 0.1);
  background: rgba(255, 255, 255, 0.05);
}

.ds-input::placeholder {
  color: #64748b;
}

.ds-input:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.ds-input--error {
  border-color: rgba(239, 68, 68, 0.3);
}

.ds-input--error:focus {
  border-color: rgba(239, 68, 68, 0.5);
  box-shadow: 0 0 12px rgba(239, 68, 68, 0.08);
}

/* Sizes */
.ds-input--sm {
  padding: 6px 12px;
  font-size: 12px;
}

.ds-input--md {
  padding: 8px 16px;
  font-size: 14px;
}

.ds-input--lg {
  padding: 12px 20px;
  font-size: 16px;
}

/* Responsive */
@media (max-width: 768px) {
  .ds-input--lg {
    font-size: 14px;
    padding: 10px 16px;
  }
}

/* Accessibility */
@media (prefers-reduced-motion: reduce) {
  .ds-input {
    transition: none;
  }
}
</style>
