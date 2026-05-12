<template>
  <div
    :class="['ds-skeleton', `ds-skeleton-${variant}`]"
    :style="containerStyle"
  >
    <!-- 标题骨架 -->
    <template v-if="variant === 'text' || variant === '混合'">
      <div
        v-for="(line, index) in lines"
        :key="index"
        :class="[
          'ds-skeleton-line',
          { 'ds-skeleton-last': index === lines - 1 },
        ]"
        :style="getLineStyle(index)"
      ></div>
    </template>

    <!-- 矩形骨架 -->
    <template v-if="variant === 'rect'">
      <div class="ds-skeleton-rect" :style="rectStyle"></div>
    </template>

    <!-- 圆形骨架 -->
    <template v-if="variant === 'circle'">
      <div class="ds-skeleton-circle" :style="circleStyle"></div>
    </template>

    <!-- 卡片骨架 -->
    <template v-if="variant === 'card'">
      <div class="ds-skeleton-card">
        <div class="ds-skeleton-card-cover"></div>
        <div class="ds-skeleton-card-body">
          <div class="ds-skeleton-line ds-skeleton-title"></div>
          <div class="ds-skeleton-line"></div>
          <div class="ds-skeleton-line ds-skeleton-last"></div>
        </div>
      </div>
    </template>

    <!-- 列表骨架 -->
    <template v-if="variant === 'list'">
      <div v-for="item in count" :key="item" class="ds-skeleton-list-item">
        <div class="ds-skeleton-avatar"></div>
        <div class="ds-skeleton-content">
          <div class="ds-skeleton-line ds-skeleton-title"></div>
          <div class="ds-skeleton-line"></div>
        </div>
      </div>
    </template>

    <!-- 表格骨架 -->
    <template v-if="variant === 'table'">
      <div class="ds-skeleton-table">
        <div class="ds-skeleton-table-header">
          <div
            v-for="col in columns"
            :key="col"
            class="ds-skeleton-table-cell"
          ></div>
        </div>
        <div v-for="row in rows" :key="row" class="ds-skeleton-table-row">
          <div
            v-for="col in columns"
            :key="col"
            class="ds-skeleton-table-cell"
          ></div>
        </div>
      </div>
    </template>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";

interface Props {
  variant?: "text" | "rect" | "circle" | "card" | "list" | "table";
  size?: "sm" | "md" | "lg";
  lines?: number;
  count?: number;
  columns?: number;
  rows?: number;
  active?: boolean;
  round?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  variant: "text",
  size: "md",
  lines: 3,
  count: 3,
  columns: 4,
  rows: 5,
  active: true,
  round: false,
});

// 容器样式
const containerStyle = computed(() => {
  const style: Record<string, string> = {};
  if (props.variant === "rect" || props.variant === "circle") {
    style.width = `${sizeMap[props.size]}px`;
    style.height = `${sizeMap[props.size]}px`;
  }
  return style;
});

// 矩形样式
const rectStyle = computed(() => {
  return {
    borderRadius: props.round ? "50%" : "4px",
  };
});

// 圆形样式
const circleStyle = computed(() => {
  return {
    borderRadius: "50%",
  };
});

// 文本行样式
const getLineStyle = (index: number) => {
  const style: Record<string, string> = {};
  // 最后一行短一些
  if (index === props.lines - 1 && props.lines > 1) {
    style.width = "60%";
  }
  return style;
};

// 尺寸映射
const sizeMap = {
  sm: 80,
  md: 120,
  lg: 160,
};
</script>

<style scoped>
.ds-skeleton {
  width: 100%;
}

/* 动画 */
@keyframes ds-skeleton-shimmer {
  0% {
    background-position: -200% 0;
  }
  100% {
    background-position: 200% 0;
  }
}

.ds-skeleton-active .ds-skeleton-line,
.ds-skeleton-active .ds-skeleton-rect,
.ds-skeleton-active .ds-skeleton-circle,
.ds-skeleton-active .ds-skeleton-card-cover,
.ds-skeleton-active .ds-skeleton-avatar,
.ds-skeleton-active .ds-skeleton-table-cell {
  background: linear-gradient(
    90deg,
    var(--color-border) 25%,
    var(--color-bg) 50%,
    var(--color-border) 75%
  );
  background-size: 400% 100%;
  animation: ds-skeleton-shimmer 1.5s ease-in-out infinite;
}

/* 文本骨架 */
.ds-skeleton-line {
  height: 16px;
  background: var(--color-border);
  border-radius: 4px;
  margin-bottom: 12px;
  width: 100%;
}

.ds-skeleton-line:last-child {
  margin-bottom: 0;
}

.ds-skeleton-last {
  width: 60%;
}

.ds-skeleton-title {
  height: 20px;
  width: 40%;
  margin-bottom: 16px;
}

/* 矩形骨架 */
.ds-skeleton-rect {
  background: var(--color-border);
  width: 100%;
  height: 100%;
}

/* 圆形骨架 */
.ds-skeleton-circle {
  background: var(--color-border);
  border-radius: 50%;
  width: 100%;
  height: 100%;
}

/* 卡片骨架 */
.ds-skeleton-card {
  background: var(--glass-bg);
  backdrop-filter: var(--glass-blur);
  -webkit-backdrop-filter: var(--glass-blur);
  border: var(--glass-border);
  border-radius: var(--glass-radius);
  overflow: hidden;
}

.ds-skeleton-card-cover {
  width: 100%;
  height: 160px;
  background: var(--color-border);
}

.ds-skeleton-card-body {
  padding: 16px;
}

/* 列表骨架 */
.ds-skeleton-list-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 0;
  border-bottom: 1px solid var(--color-border);
}

.ds-skeleton-list-item:last-child {
  border-bottom: none;
}

.ds-skeleton-avatar {
  width: 40px;
  height: 40px;
  background: var(--color-border);
  border-radius: 50%;
  flex-shrink: 0;
}

.ds-skeleton-content {
  flex: 1;
  min-width: 0;
}

/* 表格骨架 */
.ds-skeleton-table {
  width: 100%;
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
  overflow: hidden;
}

.ds-skeleton-table-header {
  display: grid;
  grid-template-columns: repeat(var(--columns, 4), 1fr);
  gap: 1px;
  background: var(--color-border);
  padding: 1px;
}

.ds-skeleton-table-row {
  display: grid;
  grid-template-columns: repeat(var(--columns, 4), 1fr);
  gap: 1px;
  background: var(--color-border);
  padding: 1px;
  border-top: 1px solid var(--color-border);
}

.ds-skeleton-table-cell {
  background: var(--color-surface);
  height: 40px;
  padding: 8px 12px;
}

/* 尺寸 */
.ds-skeleton-sm .ds-skeleton-line {
  height: 12px;
}

.ds-skeleton-sm .ds-skeleton-title {
  height: 16px;
}

.ds-skeleton-lg .ds-skeleton-line {
  height: 20px;
}

.ds-skeleton-lg .ds-skeleton-title {
  height: 24px;
}

/* 响应式 */
@media (max-width: 768px) {
  .ds-skeleton-table {
    overflow-x: auto;
  }

  .ds-skeleton-table-header,
  .ds-skeleton-table-row {
    min-width: 500px;
  }
}

/* 无障碍访问性 */
@media (prefers-reduced-motion: reduce) {
  .ds-skeleton-active .ds-skeleton-line,
  .ds-skeleton-active .ds-skeleton-rect,
  .ds-skeleton-active .ds-skeleton-circle,
  .ds-skeleton-active .ds-skeleton-card-cover,
  .ds-skeleton-active .ds-skeleton-avatar,
  .ds-skeleton-active .ds-skeleton-table-cell {
    animation: none;
    background: var(--color-border);
  }
}
</style>
