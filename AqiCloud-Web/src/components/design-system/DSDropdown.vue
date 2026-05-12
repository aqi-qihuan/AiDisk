<template>
  <div class="ds-dropdown" ref="dropdownRef" @mouseleave="handleMouseLeave">
    <div
      class="ds-dropdown-trigger"
      @click="handleClick"
      @mouseenter="handleMouseEnter"
    >
      <slot />
    </div>
    <Transition name="dropdown">
      <div v-show="visible" class="ds-dropdown-menu" :class="placement">
        <slot name="dropdown" />
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue";

/**
 * DSDropdown 下拉菜单组件
 * 提供下拉菜单功能
 */

interface DSDropdownProps {
  trigger?: "hover" | "click";
  placement?: "bottom" | "bottom-start" | "bottom-end";
  hideDelay?: number;
}

const props = withDefaults(defineProps<DSDropdownProps>(), {
  trigger: "hover",
  placement: "bottom-end",
  hideDelay: 800,
});

const visible = ref(false);
const dropdownRef = ref<HTMLElement | null>(null);
let hideTimeout: ReturnType<typeof setTimeout> | null = null;

const handleClick = () => {
  if (props.trigger === "click") {
    visible.value = !visible.value;
  }
};

const handleMouseEnter = () => {
  if (props.trigger === "hover") {
    // 清除关闭定时器
    if (hideTimeout) {
      clearTimeout(hideTimeout);
      hideTimeout = null;
    }
    visible.value = true;
  }
};

const handleMouseLeave = () => {
  if (props.trigger === "hover") {
    // 延迟关闭，给用户时间移动鼠标到菜单上
    hideTimeout = setTimeout(() => {
      visible.value = false;
    }, props.hideDelay);
  }
};

const handleClickOutside = (event: MouseEvent) => {
  if (dropdownRef.value && !dropdownRef.value.contains(event.target as Node)) {
    visible.value = false;
  }
};

onMounted(() => {
  document.addEventListener("click", handleClickOutside);
});

onUnmounted(() => {
  document.removeEventListener("click", handleClickOutside);
  if (hideTimeout) {
    clearTimeout(hideTimeout);
  }
});
</script>

<style scoped>
.ds-dropdown {
  position: relative;
  display: inline-block;
}

.ds-dropdown-trigger {
  cursor: pointer;
}

.ds-dropdown-menu {
  position: absolute;
  top: 100%;
  margin-top: 8px;
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-lg);
  min-width: 160px;
  z-index: var(--z-dropdown);
  padding: 6px 0;
}

.ds-dropdown-menu.bottom-start {
  left: 0;
}

.ds-dropdown-menu.bottom-end {
  right: 0;
}

.ds-dropdown-menu.bottom {
  left: 50%;
  transform: translateX(-50%);
}

/* Transition */
.dropdown-enter-active,
.dropdown-leave-active {
  transition: all var(--transition-base);
}

.dropdown-enter-from,
.dropdown-leave-to {
  opacity: 0;
  transform: translateY(-8px);
}
</style>
