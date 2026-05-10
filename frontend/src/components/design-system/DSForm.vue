<template>
  <form class="ds-form" @submit.prevent>
    <slot />
  </form>
</template>

<script setup lang="ts">
/**
 * DSForm 表单组件
 * 提供表单容器
 */
</script>

<style scoped>
.ds-form {
  display: flex;
  flex-direction: column;
  gap: 20px;
}
</style>
