<template>
  <div
    class="ds-tooltip-wrapper"
    @mouseenter="showTooltip"
    @mouseleave="hideTooltip"
    @focus="showTooltip"
    @blur="hideTooltip"
    ref="wrapperRef"
  >
    <!-- 触发元素 -->
    <slot />

    <!-- 工具提示 -->
    <Teleport to="body">
      <Transition name="ds-tooltip-fade">
        <div
          v-if="visible"
          :class="['ds-tooltip', `ds-tooltip-${placement}`]"
          :style="tooltipStyle"
          role="tooltip"
          aria-live="polite"
          ref="tooltipRef"
        >
          <div class="ds-tooltip-content">
            <slot name="content">{{ content }}</slot>
          </div>
          <!-- 箭头 -->
          <div class="ds-tooltip-arrow"></div>
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, nextTick } from 'vue';

interface Props {
  content: string;
  placement?: 'top' | 'bottom' | 'left' | 'right';
  trigger?: 'hover' | 'click' | 'focus';
  offset?: number;
  delay?: number;
  disabled?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  placement: 'top',
  trigger: 'hover',
  offset: 8,
  delay: 200,
  disabled: false,
});

const visible = ref(false);
const tooltipStyle = ref<Record<string, string>>({});
const wrapperRef = ref<HTMLElement | null>(null);
const tooltipRef = ref<HTMLElement | null>(null);

let showTimer: number | null = null;
let hideTimer: number | null = null;

// 显示工具提示
const showTooltip = () => {
  if (props.disabled) return;
  
  if (hideTimer) {
    clearTimeout(hideTimer);
    hideTimer = null;
  }
  
  showTimer = window.setTimeout(() => {
    visible.value = true;
    nextTick(() => {
      updatePosition();
    });
  }, props.delay);
};

// 隐藏工具提示
const hideTooltip = () => {
  if (showTimer) {
    clearTimeout(showTimer);
    showTimer = null;
  }
  
  hideTimer = window.setTimeout(() => {
    visible.value = false;
  }, 100);
};

// 更新位置
const updatePosition = () => {
  if (!wrapperRef.value || !tooltipRef.value) return;
  
  const wrapperRect = wrapperRef.value.getBoundingClientRect();
  const tooltipRect = tooltipRef.value.getBoundingClientRect();
  
  const style: Record<string, string> = {};
  const scrollX = window.pageXOffset || document.documentElement.scrollLeft;
  const scrollY = window.pageYOffset || document.documentElement.scrollTop;
  
  switch (props.placement) {
    case 'top':
      style.left = `${wrapperRect.left + wrapperRect.width / 2 - tooltipRect.width / 2 + scrollX}px`;
      style.top = `${wrapperRect.top - tooltipRect.height - props.offset + scrollY}px`;
      break;
    case 'bottom':
      style.left = `${wrapperRect.left + wrapperRect.width / 2 - tooltipRect.width / 2 + scrollX}px`;
      style.top = `${wrapperRect.bottom + props.offset + scrollY}px`;
      break;
    case 'left':
      style.left = `${wrapperRect.left - tooltipRect.width - props.offset + scrollX}px`;
      style.top = `${wrapperRect.top + wrapperRect.height / 2 - tooltipRect.height / 2 + scrollY}px`;
      break;
    case 'right':
      style.left = `${wrapperRect.right + props.offset + scrollX}px`;
      style.top = `${wrapperRect.top + wrapperRect.height / 2 - tooltipRect.height / 2 + scrollY}px`;
      break;
  }
  
  tooltipStyle.value = style;
};

// 监听滚动和窗口大小变化
const handleScroll = () => {
  if (visible.value) {
    updatePosition();
  }
};

onMounted(() => {
  window.addEventListener('scroll', handleScroll, true);
  window.addEventListener('resize', handleScroll);
});

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll, true);
  window.removeEventListener('resize', handleScroll);
  
  if (showTimer) clearTimeout(showTimer);
  if (hideTimer) clearTimeout(hideTimer);
});
</script>

<style scoped>
.ds-tooltip-wrapper {
  display: inline-block;
}

.ds-tooltip {
  position: fixed;
  z-index: var(--z-tooltip);
  pointer-events: none;
}

.ds-tooltip-content {
  background: var(--color-text-primary);
  color: var(--color-surface);
  padding: 6px 12px;
  border-radius: var(--radius-md);
  font-size: var(--text-sm);
  line-height: var(--leading-snug);
  white-space: nowrap;
  box-shadow: var(--shadow-md);
  max-width: 300px;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* 箭头 */
.ds-tooltip-arrow {
  position: absolute;
  width: 0;
  height: 0;
}

.ds-tooltip-top .ds-tooltip-arrow {
  bottom: -4px;
  left: 50%;
  transform: translateX(-50%);
  border-left: 4px solid transparent;
  border-right: 4px solid transparent;
  border-top: 4px solid var(--color-text-primary);
}

.ds-tooltip-bottom .ds-tooltip-arrow {
  top: -4px;
  left: 50%;
  transform: translateX(-50%);
  border-left: 4px solid transparent;
  border-right: 4px solid transparent;
  border-bottom: 4px solid var(--color-text-primary);
}

.ds-tooltip-left .ds-tooltip-arrow {
  right: -4px;
  top: 50%;
  transform: translateY(-50%);
  border-top: 4px solid transparent;
  border-bottom: 4px solid transparent;
  border-left: 4px solid var(--color-text-primary);
}

.ds-tooltip-right .ds-tooltip-arrow {
  left: -4px;
  top: 50%;
  transform: translateY(-50%);
  border-top: 4px solid transparent;
  border-bottom: 4px solid transparent;
  border-right: 4px solid var(--color-text-primary);
}

/* 动画 */
.ds-tooltip-fade-enter-active {
  transition: opacity var(--transition-fast);
}

.ds-tooltip-fade-leave-active {
  transition: opacity var(--transition-fast);
}

.ds-tooltip-fade-enter-from,
.ds-tooltip-fade-leave-to {
  opacity: 0;
}

/* 响应式 */
@media (max-width: 768px) {
  .ds-tooltip-content {
    font-size: var(--text-xs);
    padding: 4px 8px;
    max-width: 200px;
  }
}

/* 无障碍访问性 */
@media (prefers-reduced-motion: reduce) {
  .ds-tooltip-fade-enter-active,
  .ds-tooltip-fade-leave-active {
    transition: none;
  }
}
</style>
