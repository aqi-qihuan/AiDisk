<template>
  <div class="aside-container">
    <DSMenu :collapse="isCollapse" class="aside-menu">
      <!-- AI 功能区 - 置顶突出 -->
      <div class="sub-menu ai-menu">
        <div
          class="sub-menu-title ai-title"
          :class="{ 'is-collapse': isCollapse, 'is-open': isAiMenuOpen }"
          @click="toggleAiMenu"
        >
          <div class="ai-icon-wrapper">
            <el-icon class="ai-main-icon"><Star /></el-icon>
            <div class="ai-icon-glow"></div>
            <div class="ai-icon-sparkle sparkle-1">✦</div>
            <div class="ai-icon-sparkle sparkle-2">✦</div>
          </div>
          <span v-if="!isCollapse" class="menu-text ai-text">{{
            t("ai.title")
          }}</span>
          <el-icon v-if="!isCollapse" class="arrow-icon ai-arrow"
            ><ArrowDown
          /></el-icon>
        </div>
        <div
          v-show="isAiMenuOpen || isCollapse"
          class="sub-menu-items ai-items"
        >
          <DSMenuItem
            index="/Answer"
            :active="activeIndex === '/Answer'"
            :collapse="isCollapse"
            @click="handleSelect('/Answer')"
          >
            <template #icon>
              <div class="sub-item-icon">
                <el-icon><Document /></el-icon>
              </div>
            </template>
            {{ t("ai.answer") }}
          </DSMenuItem>
          <DSMenuItem
            index="/Chat"
            :active="activeIndex === '/Chat'"
            :collapse="isCollapse"
            @click="handleSelect('/Chat')"
          >
            <template #icon>
              <div class="sub-item-icon">
                <el-icon><ChatDotRound /></el-icon>
              </div>
            </template>
            {{ t("ai.chat") }}
          </DSMenuItem>
          <DSMenuItem
            index="/Document"
            :active="activeIndex === '/Document'"
            :collapse="isCollapse"
            @click="handleSelect('/Document')"
          >
            <template #icon>
              <div class="sub-item-icon">
                <el-icon><Edit /></el-icon>
              </div>
            </template>
            {{ t("ai.document") }}
          </DSMenuItem>
        </div>
      </div>

      <div class="menu-divider"></div>

      <DSMenuItem
        index="/file"
        :active="activeIndex === '/file'"
        :collapse="isCollapse"
        @click="handleSelect('/file')"
      >
        <template #icon>
          <el-icon><FolderOpened /></el-icon>
        </template>
        {{ t("nav.files") }}
      </DSMenuItem>

      <DSMenuItem
        index="/share"
        :active="activeIndex === '/share'"
        :collapse="isCollapse"
        @click="handleSelect('/share')"
      >
        <template #icon>
          <el-icon><Share /></el-icon>
        </template>
        {{ t("nav.share") }}
      </DSMenuItem>

      <DSMenuItem
        index="/recycle"
        :active="activeIndex === '/recycle'"
        :collapse="isCollapse"
        @click="handleSelect('/recycle')"
      >
        <template #icon>
          <el-icon><Delete /></el-icon>
        </template>
        {{ t("nav.recycle") }}
      </DSMenuItem>
    </DSMenu>

    <!-- 容量进度条组件 -->
    <div class="capacity-container">
      <CapacityProcess :is-collapse="isCollapse" />
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  Delete,
  Document,
  FolderOpened,
  Share,
  ArrowDown,
  Magic,
  ChatDotRound,
  Edit,
  Star,
} from "@element-plus/icons-vue";
import { ref, watch } from "vue";
import { useI18n } from "vue-i18n";
import { useRoute, useRouter } from "vue-router";
import CapacityProcess from "@/components/storage/CapacityProcess.vue";
import { DSMenu, DSMenuItem } from "@/components/design-system";

const { t } = useI18n();

/**
 * AsideMenu 组件 - 侧边栏菜单组件
 * 提供文件、分享、回收站、AI功能的导航
 */

interface AsideMenuProps {
  isCollapse: boolean;
}

const props = defineProps<AsideMenuProps>();

const emit = defineEmits<{
  select: [fileType: number | null];
}>();

const route = useRoute();
const router = useRouter();
const activeIndex = ref("");
const isAiMenuOpen = ref(false);

/**
 * 切换 AI 菜单展开状态
 */
const toggleAiMenu = () => {
  isAiMenuOpen.value = !isAiMenuOpen.value;
};

/**
 * 处理菜单项选择事件
 * @param index - 选中菜单项的索引
 */
const handleSelect = (index: string): void => {
  activeIndex.value = index;

  // 处理文件类型选择
  if (index === "0" || !isNaN(Number(index))) {
    emit("select", index === "0" ? null : Number(index));
    router.push("/file");
    return;
  }

  // 处理页面跳转
  if (index.startsWith("/")) {
    router.push(index);
  }
};

/**
 * 监听路由变化
 */
watch(
  () => route.path,
  (newPath) => {
    if (newPath === "/file") {
      activeIndex.value = "/file";
    } else {
      activeIndex.value = newPath;
    }
  },
  { immediate: true },
);
</script>

<style scoped>
.aside-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  background-color: var(--color-surface);
}

.aside-menu {
  flex: 1;
  border-right: none;
  overflow-y: auto;
  background-color: var(--color-surface);
  padding: 8px 0;
}

/* 子菜单样式 */
.sub-menu {
  display: flex;
  flex-direction: column;
}

.sub-menu-title {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 20px;
  cursor: pointer;
  transition: all var(--transition-base);
  color: var(--color-text-secondary);
  font-size: 14px;
  font-weight: 500;
}

.sub-menu-title:hover {
  color: var(--color-primary);
  background-color: rgba(99, 102, 241, 0.08);
}

.sub-menu-title.is-collapse {
  justify-content: center;
  padding: 14px;
}

.sub-menu-title.is-open .arrow-icon {
  transform: rotate(180deg);
}

.menu-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 24px;
  height: 24px;
  font-size: 18px;
  flex-shrink: 0;
}

.menu-text {
  flex: 1;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.arrow-icon {
  transition: transform var(--transition-base);
  font-size: 14px;
}

.sub-menu-items {
  display: flex;
  flex-direction: column;
}

/* AI 菜单特殊样式 - 卡片式设计 */
.ai-menu {
  margin: 12px;
}

.ai-title {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 16px;
  background: linear-gradient(
    135deg,
    rgba(99, 102, 241, 0.08) 0%,
    rgba(129, 140, 248, 0.08) 100%
  );
  border-radius: var(--radius-xl);
  border: 1px solid rgba(99, 102, 241, 0.12);
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.ai-title:hover {
  background: linear-gradient(
    135deg,
    rgba(99, 102, 241, 0.12) 0%,
    rgba(129, 140, 248, 0.12) 100%
  );
  border-color: rgba(99, 102, 241, 0.2);
  transform: translateY(-1px);
  box-shadow: 0 4px 16px rgba(99, 102, 241, 0.12);
}

.ai-title.is-open {
  background: linear-gradient(
    135deg,
    rgba(99, 102, 241, 0.15) 0%,
    rgba(129, 140, 248, 0.15) 100%
  );
  border-color: rgba(99, 102, 241, 0.25);
}

.ai-icon-wrapper {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 44px;
  height: 44px;
  background: linear-gradient(
    135deg,
    var(--color-primary) 0%,
    var(--color-secondary) 100%
  );
  border-radius: var(--radius-lg);
  color: white;
  font-size: 22px;
  box-shadow:
    0 4px 12px rgba(99, 102, 241, 0.3),
    0 0 0 1px rgba(255, 255, 255, 0.1) inset;
  flex-shrink: 0;
  overflow: hidden;
}

.ai-main-icon {
  position: relative;
  z-index: 2;
  animation: starPulse 2s ease-in-out infinite;
}

.ai-icon-glow {
  position: absolute;
  inset: 0;
  background: radial-gradient(
    circle at 30% 30%,
    rgba(255, 255, 255, 0.3) 0%,
    transparent 50%
  );
  z-index: 1;
}

.ai-icon-sparkle {
  position: absolute;
  color: rgba(255, 255, 255, 0.8);
  font-size: 8px;
  z-index: 3;
  animation: sparkle 1.5s ease-in-out infinite;
}

.sparkle-1 {
  top: 6px;
  right: 8px;
  animation-delay: 0s;
}

.sparkle-2 {
  bottom: 8px;
  left: 6px;
  font-size: 6px;
  animation-delay: 0.75s;
}

@keyframes starPulse {
  0%,
  100% {
    transform: scale(1);
    filter: drop-shadow(0 0 4px rgba(255, 255, 255, 0.5));
  }
  50% {
    transform: scale(1.1);
    filter: drop-shadow(0 0 8px rgba(255, 255, 255, 0.8));
  }
}

@keyframes sparkle {
  0%,
  100% {
    opacity: 0.4;
    transform: scale(0.8);
  }
  50% {
    opacity: 1;
    transform: scale(1.2);
  }
}

.ai-text {
  flex: 1;
  font-weight: var(--font-semibold);
  font-size: var(--text-base);
  color: var(--color-text-primary);
}

.ai-arrow {
  color: var(--color-primary);
  font-size: 14px;
  transition: transform 0.3s ease;
}

.ai-title.is-open .ai-arrow {
  transform: rotate(180deg);
}

/* AI 子菜单项 */
.ai-items {
  margin-top: 8px;
  padding-left: 8px;
}

.ai-items :deep(.ds-menu-item) {
  margin: 4px 0;
  padding: 12px 16px;
  border-radius: var(--radius-lg);
  background: transparent;
  transition: all 0.25s ease;
}

.ai-items :deep(.ds-menu-item:hover) {
  background: rgba(99, 102, 241, 0.06);
  transform: translateX(4px);
}

.ai-items :deep(.ds-menu-item.is-active) {
  background: rgba(99, 102, 241, 0.1);
}

.ai-items :deep(.ds-menu-item.is-active::before) {
  content: "";
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 3px;
  height: 20px;
  background: linear-gradient(
    180deg,
    var(--color-primary) 0%,
    var(--color-secondary) 100%
  );
  border-radius: 0 3px 3px 0;
}

.sub-item-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  color: var(--color-text-secondary);
  transition: color 0.25s ease;
}

.ai-items :deep(.ds-menu-item:hover) .sub-item-icon {
  color: var(--color-primary);
}

.ai-items :deep(.ds-menu-item.is-active) .sub-item-icon {
  color: var(--color-primary);
}

/* 菜单分隔线 */
.menu-divider {
  height: 1px;
  background: linear-gradient(
    90deg,
    transparent 0%,
    var(--color-border) 20%,
    var(--color-border) 80%,
    transparent 100%
  );
  margin: 16px 20px;
}

/* 优化普通菜单项样式 */
:deep(.ds-menu-item) {
  margin: 4px 12px;
  padding: 12px 16px;
  border-radius: var(--radius-lg);
  transition: all 0.25s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
}

:deep(.ds-menu-item:hover) {
  background: rgba(99, 102, 241, 0.06);
  transform: translateX(4px);
}

:deep(.ds-menu-item.is-active) {
  background: rgba(99, 102, 241, 0.1);
  font-weight: var(--font-medium);
}

:deep(.ds-menu-item.is-active::before) {
  content: "";
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 3px;
  height: 20px;
  background: linear-gradient(
    180deg,
    var(--color-primary) 0%,
    var(--color-secondary) 100%
  );
  border-radius: 0 3px 3px 0;
}

:deep(.ds-menu-item .menu-item-icon) {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  border-radius: var(--radius-md);
  background: rgba(99, 102, 241, 0.08);
  color: var(--color-text-secondary);
  transition: all 0.25s ease;
}

:deep(.ds-menu-item:hover .menu-item-icon) {
  background: rgba(99, 102, 241, 0.12);
  color: var(--color-primary);
}

:deep(.ds-menu-item.is-active .menu-item-icon) {
  background: linear-gradient(
    135deg,
    var(--color-primary) 0%,
    var(--color-secondary) 100%
  );
  color: white;
  box-shadow: 0 4px 12px rgba(99, 102, 241, 0.25);
}

.capacity-container {
  padding: 16px;
  border-top: 1px solid var(--color-border);
}

/* 移动端适配 */
@media (max-width: 768px) {
  .ai-title {
    margin: 6px 8px;
    padding: 14px 16px;
  }

  .ai-icon-wrapper {
    width: 36px;
    height: 36px;
    font-size: 18px;
  }

  .menu-divider {
    margin: 12px 16px;
  }

  :deep(.ds-menu-item) {
    margin: 3px 8px;
  }

  .sub-menu-title {
    padding-left: 15px;
  }

  .ai-items :deep(.ds-menu-item) {
    padding-left: 12px;
    margin-left: 8px;
  }
}
</style>
