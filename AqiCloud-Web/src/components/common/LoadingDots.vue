<template>
  <div class="loading-dots">
    <span class="dot"></span>
    <span class="dot"></span>
    <span class="dot"></span>
  </div>
</template>

<script setup lang="ts">
/**
 * LoadingDots 组件 - 加载动画组件
 * 显示三个点的加载动画效果
 */
</script>

<style scoped>
.loading-dots {
  display: inline-flex;
  align-items: center;
  gap: 6px;
}

.dot {
  width: 8px;
  height: 8px;
  background-color: var(--el-color-primary);
  border-radius: 50%;
  display: inline-block;
  animation: bounce 1.4s infinite ease-in-out both;
}

.dot:nth-child(1) {
  animation-delay: -0.32s;
}

.dot:nth-child(2) {
  animation-delay: -0.16s;
}

.dot:nth-child(3) {
  animation-delay: 0s;
}

@keyframes bounce {
  0%,
  80%,
  100% {
    transform: scale(0);
  }
  40% {
    transform: scale(1);
  }
}

/* 移动端适配 */
@media (max-width: 768px) {
  .dot {
    width: 6px;
    height: 6px;
  }

  .loading-dots {
    gap: 4px;
  }
}
</style>
