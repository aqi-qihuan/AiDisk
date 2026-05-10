<template>
  <div class="basic-layout">
    <!-- Header -->
    <header class="layout-header">
      <GlobalHeader />
    </header>

    <!-- Main Container -->
    <div class="layout-container" :class="{ 'sidebar-open': showSidebar && !isCollapse }">
      <!-- Mobile Overlay -->
      <div
        v-if="showSidebar && isMobile"
        class="mobile-overlay"
        :class="{ visible: !isCollapse }"
        @click="closeSidebar"
      ></div>

      <!-- Sidebar (conditionally shown) -->
      <aside
        v-if="showSidebar"
        class="layout-sidebar"
        :class="{ collapsed: isCollapse }"
      >
        <AsideMenu :is-collapse="isCollapse" @select="handleFileTypeSelect" />
      </aside>

      <!-- Sidebar Toggle Button -->
      <div v-if="showSidebar && !isMobile" class="sidebar-toggle" :style="{ left: toggleButtonLeft }">
        <el-tooltip
          effect="dark"
          :content="isCollapse ? '展开侧边栏' : '收起侧边栏'"
          placement="right"
        >
          <button class="toggle-button" @click="toggleCollapse">
            <div class="toggle-line"></div>
            <el-icon class="toggle-arrow" :size="12">
              <component :is="isCollapse ? DArrowRight : DArrowLeft" />
            </el-icon>
          </button>
        </el-tooltip>
      </div>

      <!-- Mobile Sidebar Toggle Button -->
      <div v-if="showSidebar && isMobile" class="mobile-sidebar-toggle" :class="{ hidden: !isCollapse }">
        <button class="mobile-toggle-button" @click="openSidebar">
          <el-icon :size="20"><Menu /></el-icon>
        </button>
      </div>

      <!-- Main Content -->
      <main class="layout-main">
        <router-view :fileType="fileType" />
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
/**
 * BasicLayout - 统一布局组件
 * 
 * 特性：
 * - 使用设计系统 CSS 变量
 * - 响应式侧边栏
 * - 流畅的动画过渡
 * - 移动端适配
 */
import GlobalHeader from "@/components/GlobalHeader.vue";
import AsideMenu from "@/components/file/AsideMenu.vue";
import { DArrowLeft, DArrowRight, Menu } from "@element-plus/icons-vue";
import { computed, onMounted, onUnmounted, ref } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();
const fileType = ref<number | null>(null);
const isCollapse = ref(false);
const windowWidth = ref(window.innerWidth);

// 计算是否为移动端
const isMobile = computed((): boolean => windowWidth.value <= 768);

// 监听窗口大小变化
const handleResize = () => {
  windowWidth.value = window.innerWidth;
  // 切换到桌面端时自动展开侧边栏
  if (!isMobile.value) {
    isCollapse.value = false;
  }
};

window.addEventListener('resize', handleResize);

// 组件卸载时移除监听
onUnmounted(() => {
  window.removeEventListener('resize', handleResize);
});

/**
 * 计算侧边栏切换按钮的位置
 */
const toggleButtonLeft = computed((): string => {
  // 平板模式下侧边栏宽度为 200px，收起时为 56px
  // 桌面模式下侧边栏宽度为 240px，收起时为 64px
  const isTablet = windowWidth.value <= 1024;
  if (isCollapse.value) {
    return isTablet ? '56px' : '64px';
  }
  return isTablet ? '200px' : '240px';
});

/**
 * 判断当前路由是否需要显示侧边栏
 * 文件管理、分享、回收站等页面显示侧边栏
 */
const showSidebar = computed((): boolean => {
  const path = route.path;
  const sidebarRoutes = [
    "/file",
    "/share",
    "/recycle",
    "/Search",
    "/Answer",
    "/Chat",
    "/Document",
    "/picture"
  ];
  return sidebarRoutes.some(route => path.startsWith(route) || path === route);
});

/**
 * 处理文件类型选择
 * @param type 文件类型ID
 */
const handleFileTypeSelect = (type: number | null): void => {
  fileType.value = type;
};

/**
 * 切换侧边栏展开/收起状态
 */
const toggleCollapse = (): void => {
  isCollapse.value = !isCollapse.value;
  localStorage.setItem("asideMenuCollapsed", isCollapse.value.toString());
};

/**
 * 打开侧边栏（移动端）
 */
const openSidebar = (): void => {
  isCollapse.value = false;
};

/**
 * 关闭侧边栏（移动端）
 */
const closeSidebar = (): void => {
  isCollapse.value = true;
};

/**
 * 组件挂载时恢复侧边栏状态
 */
onMounted(() => {
  // 移动端默认收起侧边栏
  if (isMobile.value) {
    isCollapse.value = true;
  } else {
    isCollapse.value = localStorage.getItem("asideMenuCollapsed") === "true";
  }
});
</script>

<style scoped>
.basic-layout {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  background-color: var(--color-bg);
}

/* ==================== Header ==================== */
.layout-header {
  position: sticky;
  top: 0;
  z-index: var(--z-sticky);
  height: 60px;
  background-color: var(--color-surface);
  box-shadow: var(--shadow-sm);
  border-bottom: 1px solid var(--color-border);
}

/* ==================== Container ==================== */
.layout-container {
  display: flex;
  flex: 1;
  overflow: hidden;
  position: relative;
}

/* ==================== Sidebar ==================== */
.layout-sidebar {
  width: 240px;
  background-color: var(--color-surface);
  border-right: 1px solid var(--color-border);
  transition: width var(--transition-slow);
  overflow-x: hidden;
  overflow-y: auto;
  flex-shrink: 0;
}

.layout-sidebar.collapsed {
  width: 64px;
}

/* Sidebar Scrollbar */
.layout-sidebar::-webkit-scrollbar {
  width: 6px;
}

.layout-sidebar::-webkit-scrollbar-track {
  background: transparent;
}

.layout-sidebar::-webkit-scrollbar-thumb {
  background-color: var(--color-border);
  border-radius: var(--radius-full);
}

.layout-sidebar::-webkit-scrollbar-thumb:hover {
  background-color: var(--color-text-tertiary);
}

/* ==================== Sidebar Toggle ==================== */
.sidebar-toggle {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  z-index: var(--z-fixed);
  transition: left var(--transition-slow);
}

.toggle-button {
  width: 12px;
  height: 60px;
  background: transparent;
  border: none;
  cursor: pointer;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 0;
  position: relative;
  transition: all var(--transition-base);
}

/* 细线 */
.toggle-line {
  position: absolute;
  left: 0;
  width: 2px;
  height: 40px;
  background: linear-gradient(180deg, transparent, var(--color-primary), transparent);
  border-radius: 1px;
  transition: all var(--transition-base);
}

/* 箭头图标 */
.toggle-arrow {
  color: var(--color-primary);
  opacity: 0;
  transform: translateX(-2px);
  transition: all var(--transition-base);
}

/* 悬停效果 */
.toggle-button:hover {
  width: 20px;
}

.toggle-button:hover .toggle-line {
  height: 50px;
  background: linear-gradient(180deg, transparent, var(--color-secondary), transparent);
}

.toggle-button:hover .toggle-arrow {
  opacity: 1;
  transform: translateX(0);
  color: var(--color-secondary);
}

.toggle-button:active .toggle-arrow {
  transform: scale(0.9);
}

/* ==================== Main Content ==================== */
.layout-main {
  flex: 1;
  background-color: var(--color-surface);
  overflow-x: hidden;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
}

/* Main Scrollbar */
.layout-main::-webkit-scrollbar {
  width: 8px;
}

.layout-main::-webkit-scrollbar-track {
  background: var(--color-bg);
}

.layout-main::-webkit-scrollbar-thumb {
  background-color: var(--color-border);
  border-radius: var(--radius-full);
}

.layout-main::-webkit-scrollbar-thumb:hover {
  background-color: var(--color-text-tertiary);
}

/* ==================== Responsive ==================== */
/* Tablet */
@media (max-width: 1024px) {
  .layout-sidebar {
    width: 200px;
  }

  .layout-sidebar.collapsed {
    width: 56px;
  }
}

/* Mobile */
@media (max-width: 768px) {
  .layout-header {
    height: 56px;
  }

  .layout-sidebar {
    position: fixed;
    left: 0;
    top: 56px;
    height: calc(100vh - 56px);
    width: 240px;
    z-index: var(--z-fixed);
    transform: translateX(-100%);
    transition: transform var(--transition-slow);
  }

  .layout-sidebar:not(.collapsed) {
    transform: translateX(0);
  }

  .layout-sidebar.collapsed {
    width: 240px;
    transform: translateX(-100%);
  }

  .sidebar-toggle {
    display: none;
  }

  /* Mobile Overlay */
  .mobile-overlay {
    position: fixed;
    top: 56px;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.5);
    opacity: 0;
    pointer-events: none;
    transition: opacity var(--transition-slow);
    z-index: calc(var(--z-fixed) - 1);
  }

  .mobile-overlay.visible {
    opacity: 1;
    pointer-events: auto;
  }

  /* Mobile Sidebar Toggle Button */
  .mobile-sidebar-toggle {
    position: fixed;
    left: 16px;
    top: 72px;
    z-index: calc(var(--z-fixed) - 2);
    transition: all var(--transition-base);
  }

  .mobile-sidebar-toggle.hidden {
    opacity: 0;
    pointer-events: none;
    transform: translateX(-20px);
  }

  .mobile-toggle-button {
    width: 44px;
    height: 44px;
    background: var(--color-primary);
    border: none;
    border-radius: var(--radius-md);
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    box-shadow: var(--shadow-lg);
    transition: all var(--transition-base);
  }

  .mobile-toggle-button:hover {
    background-color: var(--color-secondary);
    transform: scale(1.05);
  }

  .mobile-toggle-button:active {
    transform: scale(0.95);
  }

  .mobile-toggle-button :deep(.el-icon) {
    color: white;
  }
}

/* Small Mobile */
@media (max-width: 576px) {
  .layout-header {
    height: 52px;
  }

  .layout-sidebar {
    top: 52px;
    height: calc(100vh - 52px);
    width: 200px;
  }

  .mobile-overlay {
    top: 52px;
  }

  .mobile-sidebar-toggle {
    top: 68px;
    left: 12px;
  }

  .mobile-toggle-button {
    width: 40px;
    height: 40px;
  }
}

/* ==================== Animations ==================== */
@keyframes slideInLeft {
  from {
    transform: translateX(-100%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

/* Sidebar Animation */
.layout-sidebar {
  animation: slideInLeft var(--transition-slow) ease-out;
}

/* Main Content Animation */
.layout-main {
  animation: fadeIn var(--transition-slow) ease-out;
}
</style>
