<template>
  <div :class="['ds-input-wrapper', { 'ds-input-wrapper-disabled': disabled, 'ds-input-wrapper-error': error }]">
    <label v-if="label" :class="['ds-input-label', { 'ds-input-label-required': required }]">
      {{ label }}
    </label>
    <div class="ds-input-container">
      <span v-if="$slots.prefix" class="ds-input__prefix">
        <slot name="prefix"></slot>
      </span>
      <input
        ref="inputRef"
        :type="type"
        :value="modelValue"
        :placeholder="placeholder"
        :disabled="disabled"
        :readonly="readonly"
        :maxlength="maxlength"
        :class="[
          'ds-input',
          `ds-input-${size}`,
          {
            'ds-input-disabled': disabled,
            'ds-input-error': error,
            'ds-input-with-prefix': $slots.prefix,
            'ds-input-with-suffix': $slots.suffix,
          }
        ]"
        @input="handleInput"
        @change="handleChange"
        @focus="handleFocus"
        @blur="handleBlur"
        @keyup.enter="handleEnter"
      />
      <span v-if="$slots.suffix" class="ds-input__suffix">
        <slot name="suffix"></slot>
      </span>
      <span v-if="showClear && modelValue" class="ds-input__clear" @click="handleClear">
        ×
      </span>
      <span v-if="clearable && modelValue" class="ds-input__clear" @click="handleClear">
        ×
      </span>
    </div>
    <div v-if="error || hint" :class="['ds-input-message', { 'ds-input-message-error': error }]">
      {{ error || hint }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, defineProps, defineEmits, defineExpose } from 'vue';

interface Props {
  modelValue?: string | number;
  label?: string;
  type?: string;
  placeholder?: string;
  disabled?: boolean;
  readonly?: boolean;
  maxlength?: number;
  clearable?: boolean;
  showClear?: boolean;
  required?: boolean;
  error?: string;
  hint?: string;
  size?: 'small' | 'medium' | 'large';
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: '',
  type: 'text',
  placeholder: '',
  disabled: false,
  readonly: false,
  clearable: false,
  showClear: false,
  required: false,
  error: '',
  hint: '',
  size: 'medium',
});

const emit = defineEmits<{
  (e: 'update:modelValue', value: string | number): void;
  (e: 'change', value: string | number): void;
  (e: 'focus', event: FocusEvent): void;
  (e: 'blur', event: FocusEvent): void;
  (e: 'enter', value: string | number): void;
  (e: 'clear'): void;
}>();

const inputRef = ref<HTMLInputElement>();

const handleInput = (event: Event) => {
  const target = event.target as HTMLInputElement;
  emit('update:modelValue', target.value);
};

const handleChange = (event: Event) => {
  const target = event.target as HTMLInputElement;
  emit('change', target.value);
};

const handleFocus = (event: FocusEvent) => {
  emit('focus', event);
};

const handleBlur = (event: FocusEvent) => {
  emit('blur', event);
};

const handleEnter = () => {
  emit('enter', props.modelValue);
};

const handleClear = () => {
  emit('update:modelValue', '');
  emit('clear');
};

const focus = () => {
  inputRef.value?.focus();
};

const blur = () => {
  inputRef.value?.blur();
};

defineExpose({
  focus,
  blur,
});
</script>

<style scoped>
.ds-input-wrapper {
  display: flex;
  flex-direction: column;
  gap: 8px;
  width: 100%;
}

.ds-input-label {
  font-size: 14px;
  font-weight: 500;
  color: #1E1B4B;
  line-height: 1.5;
}

.ds-input-label-required::after {
  content: ' *';
  color: #EF4444;
}

.ds-input-container {
  position: relative;
  display: flex;
  align-items: center;
  width: 100%;
}

.ds-input {
  flex: 1;
  width: 100%;
  padding: 0 16px;
  border: 1px solid #E2E8F0;
  border-radius: 8px;
  font-size: 16px;
  color: #1E1B4B;
  background-color: #FFFFFF;
  transition: all 0.2s ease;
  outline: none;
  box-sizing: border-box;
}

/* Sizes */
.ds-input-small {
  height: 32px;
  font-size: 14px;
  padding: 0 12px;
}

.ds-input-medium {
  height: 40px;
  font-size: 16px;
  padding: 0 16px;
}

.ds-input-large {
  height: 48px;
  font-size: 18px;
  padding: 0 20px;
}

/* States */
.ds-input:hover:not(.ds-input-disabled) {
  border-color: #6366F1;
}

.ds-input:focus {
  border-color: #6366F1;
  box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.1);
}

.ds-input-disabled {
  background-color: #F5F3FF;
  cursor: not-allowed;
  opacity: 0.6;
}

.ds-input-error {
  border-color: #EF4444;
}

.ds-input-error:focus {
  box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.1);
}

/* Prefix & Suffix */
.ds-input-with-prefix {
  padding-left: 40px;
}

.ds-input-with-suffix {
  padding-right: 40px;
}

.ds-input__prefix,
.ds-input__suffix {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  display: flex;
  align-items: center;
  color: #64748B;
  pointer-events: none;
}

.ds-input__prefix {
  left: 12px;
}

.ds-input__suffix {
  right: 12px;
}

/* Clear Button */
.ds-input__clear {
  position: absolute;
  right: 36px;
  top: 50%;
  transform: translateY(-50%);
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  color: #94A3B8;
  cursor: pointer;
  transition: color 0.2s ease;
  line-height: 1;
}

.ds-input__clear:hover {
  color: #64748B;
}

/* Message */
.ds-input-message {
  font-size: 12px;
  color: #64748B;
  line-height: 1.5;
  min-height: 18px;
}

.ds-input-message-error {
  color: #EF4444;
}
</style>
