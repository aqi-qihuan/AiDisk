<template>
  <teleport to="body">
    <div v-if="visible" class="context-menu" :style="menuStyle" @click.stop>
      <div class="context-menu-item" @click="handleAction('view')">
        <el-icon><View /></el-icon>
        <span>{{ t("file.view") }}</span>
      </div>
      <div class="context-menu-item" @click="handleAction('copy')">
        <el-icon><DocumentCopy /></el-icon>
        <span>{{ t("file.copy") }}</span>
      </div>
      <div class="context-menu-item" @click="handleAction('move')">
        <el-icon><Position /></el-icon>
        <span>{{ t("file.move") }}</span>
      </div>
      <div class="context-menu-item" @click="handleAction('share')">
        <el-icon><Share /></el-icon>
        <span>{{ t("file.share") }}</span>
      </div>
      <div class="context-menu-item" @click="handleAction('rename')">
        <el-icon><Edit /></el-icon>
        <span>{{ t("file.rename") }}</span>
      </div>
      <div class="context-menu-item" @click="handleAction('info')">
        <el-icon><InfoFilled /></el-icon>
        <span>{{ t("file.info") }}</span>
      </div>
      <div class="context-menu-divider"></div>
      <div class="context-menu-item delete" @click="handleAction('delete')">
        <el-icon><Delete /></el-icon>
        <span>{{ t("file.delete") }}</span>
      </div>
    </div>
  </teleport>
</template>

<script setup lang="ts">
import {
  View,
  Position,
  Share,
  Edit,
  InfoFilled,
  DocumentCopy,
  Delete,
} from "@element-plus/icons-vue";
import { computed, onMounted, onUnmounted } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

/**
 * ContextMenu 组件 - 右键菜单组件
 * 提供文件操作的快捷菜单功能
 */

interface ContextMenuProps {
  visible: boolean;
  x: number;
  y: number;
}

const props = defineProps<ContextMenuProps>();

const emit = defineEmits<{
  "update:visible": [value: boolean];
  action: [action: string];
}>();

/**
 * 计算右键菜单的定位样式
 * 自动调整位置避免超出屏幕边界
 */
const menuStyle = computed(() => {
  const menuWidth = 150;
  const menuHeight = 280;
  const padding = 10;

  let left = props.x;
  let top = props.y;

  // 避免超出右边界
  if (left + menuWidth + padding > window.innerWidth) {
    left = window.innerWidth - menuWidth - padding;
  }

  // 避免超出下边界
  if (top + menuHeight + padding > window.innerHeight) {
    top = window.innerHeight - menuHeight - padding;
  }

  return {
    position: "fixed" as const,
    top: `${top}px`,
    left: `${left}px`,
  };
});

/**
 * 处理右键菜单项点击事件
 * @param action - 操作类型
 */
const handleAction = (action: string): void => {
  emit("action", action);
  emit("update:visible", false);
};

/**
 * 关闭右键菜单
 */
const closeContextMenu = (): void => {
  emit("update:visible", false);
};

/**
 * 组件挂载时，添加全局点击事件监听器
 */
onMounted(() => {
  document.addEventListener("click", closeContextMenu);
});

/**
 * 组件卸载时，移除全局点击事件监听器
 */
onUnmounted(() => {
  document.removeEventListener("click", closeContextMenu);
});
</script>

<style scoped>
.context-menu {
  position: fixed;
  background: var(--color-surface, #ffffff);
  border-radius: var(--radius-lg, 12px);
  padding: var(--spacing-xs, 6px) 0;
  min-width: 160px;
  box-shadow: var(--shadow-xl, 0 10px 40px rgba(0, 0, 0, 0.15));
  z-index: var(--z-index-dropdown, 2000);
  animation: fadeIn 0.2s var(--ease-out, cubic-bezier(0.4, 0, 0.2, 1));
  border: 1px solid var(--color-border, #e2e8f0);
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: scale(0.95) translateY(-4px);
  }
  to {
    opacity: 1;
    transform: scale(1) translateY(0);
  }
}

.context-menu-item {
  padding: var(--spacing-sm, 10px) var(--spacing-md, 16px);
  cursor: pointer;
  display: flex;
  align-items: center;
  transition: all var(--transition-fast, 0.15s) ease;
  user-select: none;
  margin: 0 var(--spacing-xs, 4px);
  border-radius: var(--radius-md, 8px);
}

.context-menu-item:hover {
  background-color: var(--color-primary-50, #eef2ff);
}

.context-menu-item:active {
  background-color: var(--color-primary-100, #e0e7ff);
  transform: scale(0.98);
}

.context-menu-item .el-icon {
  margin-right: var(--spacing-sm, 10px);
  font-size: 16px;
  color: var(--color-text-tertiary, #64748b);
  transition: color var(--transition-fast, 0.15s) ease;
}

.context-menu-item:hover .el-icon {
  color: var(--color-primary, #6366f1);
}

.context-menu-item span {
  font-size: 14px;
  color: var(--color-text-primary, #1e1b4b);
  font-weight: 500;
  transition: color var(--transition-fast, 0.15s) ease;
}

.context-menu-item:hover span {
  color: var(--color-primary, #6366f1);
}

.context-menu-divider {
  margin: var(--spacing-xs, 6px) var(--spacing-sm, 8px);
  height: 1px;
  background-color: var(--color-border, #e2e8f0);
}

.context-menu-item.delete {
  color: var(--color-error, #ef4444);
}

.context-menu-item.delete .el-icon {
  color: var(--color-error, #ef4444);
}

.context-menu-item.delete:hover {
  background-color: var(--color-error-50, #fef2f2);
}

.context-menu-item.delete:hover .el-icon,
.context-menu-item.delete:hover span {
  color: var(--color-error, #ef4444);
}

.context-menu-item.delete:active {
  background-color: var(--color-error-100, #fee2e2);
}
</style>
