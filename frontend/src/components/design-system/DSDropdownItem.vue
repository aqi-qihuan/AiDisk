<template>
  <div class="ds-dropdown-item" :class="{ 'is-divided': divided }" @click="handleClick">
    <span v-if="$slots.icon" class="dropdown-item-icon">
      <slot name="icon" />
    </span>
    <span class="dropdown-item-text">
      <slot />
    </span>
  </div>
</template>

<script setup lang="ts">
/**
 * DSDropdownItem 下拉菜单项组件
 * 提供下拉菜单中的单个选项
 */

interface DSDropdownItemProps {
  divided?: boolean;
}

withDefaults(defineProps<DSDropdownItemProps>(), {
  divided: false,
});

const emit = defineEmits<{
  click: [];
}>();

const handleClick = () => {
  emit("click");
};
</script>

<style scoped>
.ds-dropdown-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 16px;
  cursor: pointer;
  transition: all var(--transition-base);
  color: var(--color-text-secondary);
  font-size: 14px;
}

.ds-dropdown-item:hover {
  background-color: var(--color-bg);
  color: var(--color-primary);
}

.ds-dropdown-item.is-divided {
  border-top: 1px solid var(--color-border);
  margin-top: 6px;
  padding-top: 10px;
}

.dropdown-item-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 16px;
}

.dropdown-item-text {
  white-space: nowrap;
}
</style>
