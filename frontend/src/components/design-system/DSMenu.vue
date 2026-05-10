<template>
  <nav class="ds-menu" :class="{ 'is-collapse': collapse }">
    <slot />
  </nav>
</template>

<script setup lang="ts">
/**
 * DSMenu 菜单组件
 * 提供垂直导航菜单容器
 */

interface DSMenuProps {
  collapse?: boolean;
}

withDefaults(defineProps<DSMenuProps>(), {
  collapse: false,
});
</script>

<style scoped>
.ds-menu {
  display: flex;
  flex-direction: column;
  background-color: var(--color-surface);
  border-right: 1px solid var(--color-border);
  height: 100%;
  overflow-y: auto;
}

.ds-menu.is-collapse {
  width: 64px;
}

.ds-menu:not(.is-collapse) {
  width: 240px;
}
</style>
