<template>
  <div
    class="ds-menu-item"
    :class="{
      'is-active': active,
      'is-collapse': collapse,
    }"
    @click="handleClick"
  >
    <div class="menu-item-icon" v-if="$slots.icon">
      <slot name="icon" />
    </div>
    <span v-if="!collapse" class="menu-item-title">
      <slot />
    </span>
  </div>
</template>

<script setup lang="ts">
/**
 * DSMenuItem 菜单项组件
 * 提供单个菜单项
 */

interface DSMenuItemProps {
  active?: boolean;
  collapse?: boolean;
  index?: string;
}

const props = withDefaults(defineProps<DSMenuItemProps>(), {
  active: false,
  collapse: false,
});

const emit = defineEmits<{
  click: [index: string];
}>();

const handleClick = () => {
  emit("click", props.index || "");
};
</script>

<style scoped>
.ds-menu-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 20px;
  cursor: pointer;
  transition: all var(--transition-base);
  color: var(--color-text-secondary);
  position: relative;
}

.ds-menu-item:hover {
  color: var(--color-primary);
  background-color: rgba(99, 102, 241, 0.08);
}

.ds-menu-item.is-active {
  color: var(--color-primary);
  background-color: rgba(99, 102, 241, 0.12);
}

.ds-menu-item.is-active::before {
  content: "";
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 3px;
  height: 20px;
  background: var(--color-primary);
  border-radius: 0 2px 2px 0;
}

.menu-item-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 24px;
  height: 24px;
  font-size: 18px;
  flex-shrink: 0;
}

.ds-menu-item.is-collapse {
  justify-content: center;
  padding: 14px;
}

.ds-menu-item.is-collapse .menu-item-icon {
  width: 24px;
  height: 24px;
}

.menu-item-title {
  font-size: 14px;
  font-weight: 500;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>
