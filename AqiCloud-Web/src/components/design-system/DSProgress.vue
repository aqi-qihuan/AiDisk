<template>
  <div :class="['ds-progress', `ds-progress-${type}`, `ds-progress-${size}`]">
    <!-- 线性进度条 -->
    <div v-if="type === 'line'" class="ds-progress-line">
      <div class="ds-progress-line-bg">
        <div
          class="ds-progress-line-fill"
          :style="lineStyle"
          role="progressbar"
          :aria-valuenow="percent"
          aria-valuemin="0"
          aria-valuemax="100"
        >
          <span
            v-if="showInfo && textInside"
            class="ds-progress-line-text-inside"
          >
            {{ formatPercent }}
          </span>
        </div>
      </div>
      <span v-if="showInfo && !textInside" class="ds-progress-line-text">
        {{ formatPercent }}
      </span>
    </div>

    <!-- 环形进度条 -->
    <div
      v-else-if="type === 'circle' || type === 'dashboard'"
      class="ds-progress-circle"
    >
      <svg
        :width="circleSize"
        :height="circleSize"
        :viewBox="`0 0 ${circleSize} ${circleSize}`"
      >
        <!-- 背景圆 -->
        <circle
          :cx="circleSize / 2"
          :cy="circleSize / 2"
          :r="radius"
          fill="none"
          :stroke="trackColor"
          :stroke-width="strokeWidth"
        />
        <!-- 进度圆 -->
        <circle
          :cx="circleSize / 2"
          :cy="circleSize / 2"
          :r="radius"
          fill="none"
          :stroke="strokeColor"
          :stroke-width="strokeWidth"
          :stroke-dasharray="circumference"
          :stroke-dashoffset="dashOffset"
          stroke-linecap="round"
          :style="circleStyle"
          role="progressbar"
          :aria-valuenow="percent"
          aria-valuemin="0"
          aria-valuemax="100"
        />
      </svg>
      <div v-if="showInfo" class="ds-progress-circle-text">
        <slot name="format">{{ formatPercent }}</slot>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";

interface Props {
  percent?: number;
  type?: "line" | "circle" | "dashboard";
  size?: "sm" | "md" | "lg";
  strokeWidth?: number;
  showInfo?: boolean;
  textInside?: boolean;
  status?: "normal" | "success" | "exception" | "active";
  color?: string;
  trackColor?: string;
  format?: (percent: number) => string;
}

const props = withDefaults(defineProps<Props>(), {
  percent: 0,
  type: "line",
  size: "md",
  strokeWidth: 6,
  showInfo: true,
  textInside: false,
  status: "normal",
  color: "",
  trackColor: "#E5E7EB",
  format: (percent: number) => `${percent}%`,
});

// 计算百分比（限制在 0-100）
const clampedPercent = computed(() => {
  return Math.min(100, Math.max(0, props.percent));
});

// 格式化百分比文本
const formatPercent = computed(() => {
  return props.format(clampedPercent.value);
});

// 线性进度条样式
const lineStyle = computed(() => {
  const style: Record<string, string> = {
    width: `${clampedPercent.value}%`,
    background: strokeColor.value,
  };
  return style;
});

// 环形进度条尺寸
const circleSize = computed(() => {
  const sizes = {
    sm: 80,
    md: 120,
    lg: 160,
  };
  return sizes[props.size];
});

const radius = computed(() => {
  return (circleSize.value - props.strokeWidth) / 2;
});

const circumference = computed(() => {
  return 2 * Math.PI * radius.value;
});

const dashOffset = computed(() => {
  return circumference.value * (1 - clampedPercent.value / 100);
});

const circleStyle = computed(() => {
  return {
    transition: "stroke-dashoffset 0.3s ease",
  };
});

// 进度条颜色
const strokeColor = computed(() => {
  if (props.color) return props.color;

  if (props.status === "success") return "var(--color-success)";
  if (props.status === "exception") return "var(--color-error)";
  if (props.status === "active") return "var(--color-primary)";

  // 根据百分比渐变
  if (clampedPercent.value < 30) return "var(--color-error)";
  if (clampedPercent.value < 70) return "var(--color-warning)";
  return "var(--color-success)";
});
</script>

<style scoped>
.ds-progress {
  display: inline-block;
}

/* 线性进度条 */
.ds-progress-line {
  display: flex;
  align-items: center;
  gap: 8px;
}

.ds-progress-line-bg {
  flex: 1;
  height: 8px;
  background: var(--color-border);
  border-radius: 4px;
  overflow: hidden;
}

.ds-progress-sm .ds-progress-line-bg {
  height: 4px;
}

.ds-progress-lg .ds-progress-line-bg {
  height: 12px;
}

.ds-progress-line-fill {
  height: 100%;
  border-radius: 4px;
  transition:
    width 0.3s ease,
    background 0.3s ease;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
}

.ds-progress-line-text-inside {
  font-size: 10px;
  color: white;
  font-weight: 600;
  white-space: nowrap;
}

.ds-progress-line-text {
  font-size: 14px;
  color: var(--color-text-primary);
  min-width: 40px;
  text-align: right;
}

/* 环形进度条 */
.ds-progress-circle {
  position: relative;
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.ds-progress-circle svg {
  transform: rotate(-90deg);
}

.ds-progress-circle-text {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  font-size: 14px;
  font-weight: 600;
  color: var(--color-text-primary);
}

.ds-progress-sm .ds-progress-circle-text {
  font-size: 12px;
}

.ds-progress-lg .ds-progress-circle-text {
  font-size: 18px;
}

/* 状态样式 */
.ds-progress-success .ds-progress-line-fill {
  background: var(--color-success) !important;
}

.ds-progress-exception .ds-progress-line-fill {
  background: var(--color-error) !important;
}

/* 动画 */
@keyframes ds-progress-active {
  0% {
    opacity: 0.6;
  }
  50% {
    opacity: 1;
  }
  100% {
    opacity: 0.6;
  }
}

.ds-progress-active .ds-progress-line-fill {
  animation: ds-progress-active 2s ease-in-out infinite;
}

/* 响应式 */
@media (max-width: 768px) {
  .ds-progress-circle-text {
    font-size: 12px;
  }

  .ds-progress-sm .ds-progress-circle-text {
    font-size: 10px;
  }

  .ds-progress-lg .ds-progress-circle-text {
    font-size: 14px;
  }
}

/* 无障碍访问性 */
@media (prefers-reduced-motion: reduce) {
  .ds-progress-line-fill,
  .ds-progress circle {
    transition: none !important;
  }

  .ds-progress-active .ds-progress-line-fill {
    animation: none;
  }
}
</style>
