<template>
  <div class="global-header">
    <!-- Left Section: Logo & Navigation -->
    <div class="header-left">
      <div class="logo-section">
        <img src="@/assets/logo.svg" alt="小七云盘 Logo" class="logo" />
        <h1 class="site-title">{{ t("app.name") }}</h1>
      </div>

      <!-- Navigation Menu - Desktop -->
      <nav class="nav-menu">
        <button
          v-for="item in visibleRoutes"
          :key="item.path"
          class="nav-item"
          :class="{ active: isActiveRoute(item.path) }"
          @click="doMenuClick(item.path)"
        >
          {{ translateRouteName(item.name as string) }}
        </button>
      </nav>

      <!-- Mobile Menu Button -->
      <button class="mobile-menu-btn" @click="mobileMenuOpen = !mobileMenuOpen">
        <svg
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
        >
          <line x1="3" y1="6" x2="21" y2="6" />
          <line x1="3" y1="12" x2="21" y2="12" />
          <line x1="3" y1="18" x2="21" y2="18" />
        </svg>
      </button>
    </div>

    <!-- Mobile Navigation Menu -->
    <transition name="slide-down">
      <nav v-show="mobileMenuOpen" class="mobile-nav-menu">
        <button
          v-for="item in visibleRoutes"
          :key="item.path"
          class="mobile-nav-item"
          :class="{ active: isActiveRoute(item.path) }"
          @click="handleMobileNavClick(item.path)"
        >
          {{ translateRouteName(item.name as string) }}
        </button>
      </nav>
    </transition>

    <!-- Right Section: User Actions -->
    <div class="header-right">
      <!-- Language Switcher -->
      <LanguageSwitcher />

      <!-- Logged In -->
      <template v-if="loginStatus">
        <DSDropdown trigger="hover" placement="bottom-end">
          <div class="user-profile">
            <DSAvatar
              :size="38"
              :src="loginUser.avatarUrl || getDefaultAvatar(loginUser.id)"
              class="user-avatar"
            />
            <span class="username">{{ loginUser.username }}</span>
          </div>
          <template #dropdown>
            <DSDropdownItem @click="openSettingsDialog">
              <template #icon>
                <el-icon><User /></el-icon>
              </template>
              {{ t("user.profile") }}
            </DSDropdownItem>
            <DSDropdownItem @click="handleLogout" divided>
              <template #icon>
                <el-icon><SwitchButton /></el-icon>
              </template>
              {{ t("auth.logout") }}
            </DSDropdownItem>
          </template>
        </DSDropdown>
      </template>

      <!-- Not Logged In -->
      <template v-else>
        <DSButton
          variant="primary"
          size="md"
          @click="doMenuClick('/user/login')"
        >
          {{ t("auth.login") }}
        </DSButton>
      </template>
    </div>
  </div>

  <!-- Settings Dialog -->
  <DSDialog
    v-model="settingsDialogVisible"
    :title="t('user.profile')"
    size="medium"
    :show-footer="true"
    :confirm-text="t('common.save')"
    :cancel-text="t('common.cancel')"
    @cancel="settingsDialogVisible = false"
    @confirm="updateUserSettings"
  >
    <div class="settings-content">
      <!-- 头像区域 -->
      <div class="avatar-section">
        <div class="avatar-wrapper-large">
          <ImageUpload v-model="userSettingsForm.avatarUrl" />
        </div>
        <p class="avatar-hint">{{ t("user.clickToChangeAvatar") }}</p>
      </div>

      <!-- 表单区域 -->
      <div class="form-section">
        <DSForm>
          <DSFormItem :label="t('user.username')">
            <DSInput
              v-model="userSettingsForm.username"
              :placeholder="t('user.usernamePlaceholder')"
              size="large"
            />
          </DSFormItem>
        </DSForm>
      </div>
    </div>
  </DSDialog>
</template>

<script lang="ts" setup>
/**
 * GlobalHeader - 全局顶部导航组件
 *
 * 特性：
 * - 使用设计系统组件和样式
 * - 响应式导航菜单
 * - 用户下拉菜单
 * - 个人设置对话框
 */
import { User, SwitchButton, Menu } from "@element-plus/icons-vue";
import checkAccess from "@/access/checkAccess";
import { updateUserInfo } from "@/api/user";
import ImageUpload from "@/components/common/ImageUpload.vue";
import LanguageSwitcher from "@/components/LanguageSwitcher.vue";
import {
  DSButton,
  DSAvatar,
  DSDropdown,
  DSDropdownItem,
  DSDialog,
  DSForm,
  DSFormItem,
  DSInput,
} from "@/components/design-system";
import { useLoginUserStore } from "@/store/user";
import { ElMessage } from "element-plus";
import { computed, ref, watch } from "vue";
import { useRouter } from "vue-router";
import { routes } from "../router/routes";
import { useI18n } from "vue-i18n";

const router = useRouter();
const { t } = useI18n();
const { loginStatus, logout, loginUser, fetchLoginUser } = useLoginUserStore();

// Mobile menu state
const mobileMenuOpen = ref(false);

/**
 * 路由名称翻译映射
 */
const routeNameMap: Record<string, string> = {
  主页: "nav.home",
  文件: "nav.files",
  我的分享: "nav.share",
  回收站: "nav.recycle",
  AI网盘智答: "ai.answer",
  "AI聊天智能助理-可联网": "ai.chat",
  AI在线文档助手: "ai.document",
  图片管理: "nav.pictures",
  关于: "nav.about",
  搜索: "nav.search",
  管理: "nav.admin",
};

/**
 * 翻译路由名称
 */
const translateRouteName = (name: string): string => {
  const key = routeNameMap[name];
  return key ? t(key) : name;
};

/**
 * Handle mobile navigation click
 * @param path - Target route path
 */
const handleMobileNavClick = (path: string): void => {
  mobileMenuOpen.value = false;
  doMenuClick(path);
};

/**
 * 处理菜单点击，进行路由跳转
 * @param path 目标路由路径
 */
const doMenuClick = (path: string): void => {
  router.push({ path });
};

/**
 * 处理退出登录
 */
const handleLogout = async (): Promise<void> => {
  await logout();
  router.push("/user/login");
};

/**
 * 判断当前路由是否激活
 * @param path 路由路径
 */
const isActiveRoute = (path: string): boolean => {
  return router.currentRoute.value.path === path;
};

/**
 * 根据权限过滤可见路由
 */
const visibleRoutes = computed(() => {
  return routes.filter((item) => {
    if (item.meta?.hideInMenu) {
      return false;
    }
    if (!checkAccess(loginUser, item.meta?.access as number)) {
      return false;
    }
    return true;
  });
});

// Settings Dialog
const settingsDialogVisible = ref(false);
const userSettingsForm = ref<API.UserUpdateMyRequest>({});

/**
 * 打开个人设置对话框
 */
const openSettingsDialog = (): void => {
  userSettingsForm.value = {
    userId: useLoginUserStore().loginUser.id,
    username: useLoginUserStore().loginUser.username,
    avatarUrl: useLoginUserStore().loginUser.avatarUrl,
  };
  settingsDialogVisible.value = true;
};

/**
 * 更新用户设置
 */
const updateUserSettings = async (): Promise<void> => {
  try {
    const res = await updateUserInfo({}, userSettingsForm.value);
    if (res.data.code === 0) {
      ElMessage.success(t("message.profileUpdateSuccess"));
      await useLoginUserStore().fetchLoginUser();
      const updatedUser = useLoginUserStore().loginUser;
      loginUser.username = updatedUser.username;
      loginUser.avatarUrl = updatedUser.avatarUrl;
      settingsDialogVisible.value = false;
    } else {
      ElMessage.error(t("message.updateFailed") + "，" + res.data?.msg);
    }
  } catch (error) {
    ElMessage.error(t("message.updateFailedRetry"));
  }
};

// Watch for username changes
watch(
  () => useLoginUserStore().loginUser.username,
  (newValue) => {
    if (newValue) {
      userSettingsForm.value.username = newValue;
      loginUser.username = newValue;
    }
  },
);

/**
 * 获取默认头像URL
 * @param userId 用户ID
 */
const getDefaultAvatar = (userId?: number): string => {
  const seed =
    userId?.toString() || Math.random().toString(36).substring(2, 15);
  return `https://api.dicebear.com/9.x/micah/svg?seed=${seed}`;
};
</script>

<style scoped>
/* 引入 Inter 字体 */
@import url("https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap");

/* CSS变量定义 - HOK Dark Theme */
:root {
  --header-primary: #db2777;
  --header-secondary: #f472b6;
  --header-cta: #d97706;
  --header-bg: rgba(20, 20, 28, 0.85);
  --header-text: #f8fafc;
  --header-font: var(--font-primary);
}

.global-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 32px;
  height: 68px;
  /* Dark Glassmorphism */
  background: rgba(20, 20, 28, 0.85);
  backdrop-filter: blur(24px) saturate(180%);
  -webkit-backdrop-filter: blur(24px) saturate(180%);
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
  box-shadow: 0 4px 24px rgba(0, 0, 0, 0.3);
  position: sticky;
  top: 0;
  z-index: 100;
  font-family: var(--header-font);
}

/* ==================== Left Section ==================== */
.header-left {
  display: flex;
  align-items: center;
  gap: 40px;
}

.logo-section {
  display: flex;
  align-items: center;
  gap: 12px;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  padding: 4px;
}

.logo {
  width: 44px;
  height: 44px;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  filter: drop-shadow(0 2px 4px rgba(219, 39, 119, 0.1));
}

.logo:hover {
  transform: rotate(-5deg) scale(1.08);
  filter: drop-shadow(0 4px 8px rgba(219, 39, 119, 0.2));
}

.site-title {
  font-size: 1.25rem;
  font-weight: 700;
  font-family: var(--header-font);
  margin: 0;
  background: linear-gradient(135deg, #db2777 0%, #ca8a04 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  letter-spacing: -0.02em;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.logo-section:hover .site-title {
  letter-spacing: 0;
}

/* ==================== Navigation Menu ==================== */
.nav-menu {
  display: flex;
  align-items: center;
  gap: 8px;
}

.nav-item {
  padding: 8px 16px;
  font-size: 0.9375rem;
  font-weight: 500;
  font-family: var(--header-font);
  color: var(--color-text-secondary, #94a3b8);
  background: transparent;
  border: none;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
  white-space: nowrap;
}

.nav-item::before {
  content: "";
  position: absolute;
  top: 50%;
  left: 50%;
  width: 0;
  height: 0;
  background: radial-gradient(
    circle,
    rgba(219, 39, 119, 0.15) 0%,
    transparent 70%
  );
  border-radius: 50%;
  transform: translate(-50%, -50%);
  transition: all 0.4s ease;
  z-index: -1;
}

.nav-item:hover::before {
  width: 200%;
  height: 200%;
}

.nav-item:hover {
  color: #f472b6;
  transform: translateY(-1px);
  background: rgba(219, 39, 119, 0.06);
}

.nav-item.active {
  color: #ffffff;
  background: linear-gradient(
    135deg,
    rgba(219, 39, 119, 0.12) 0%,
    rgba(217, 119, 6, 0.08) 100%
  );
  font-weight: 600;
}

.nav-item.active::after {
  content: "";
  position: absolute;
  bottom: 4px;
  left: 50%;
  transform: translateX(-50%);
  width: 24px;
  height: 3px;
  background: linear-gradient(90deg, #db2777 0%, #d97706 100%);
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(219, 39, 119, 0.4);
}

/* ==================== Right Section ==================== */
.header-right {
  display: flex;
  align-items: center;
  gap: 16px;
}

/* ==================== User Profile ==================== */
.user-profile {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 6px 14px 6px 6px;
  border-radius: 24px;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  background: rgba(255, 255, 255, 0.06);
  border: 1.5px solid rgba(219, 39, 119, 0.2);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
}

.user-profile:hover {
  background: rgba(255, 255, 255, 0.1);
  border-color: rgba(219, 39, 119, 0.4);
  box-shadow: 0 6px 20px rgba(219, 39, 119, 0.25);
  transform: translateY(-2px);
}

.user-avatar {
  border: 2px solid transparent;
  background:
    linear-gradient(#1a1a24, #1a1a24) padding-box,
    linear-gradient(135deg, #db2777 0%, #d97706 100%) border-box;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
}

.user-profile:hover .user-avatar {
  transform: scale(1.05) rotate(5deg);
  box-shadow: 0 4px 16px rgba(219, 39, 119, 0.35);
}

.username {
  font-size: 0.9375rem;
  font-weight: 600;
  font-family: var(--header-font);
  color: var(--color-text-primary, #f8fafc);
  max-width: 120px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  letter-spacing: -0.01em;
}

.user-profile:hover .username {
  color: #f8fafc;
}
/* ==================== Dropdown Styles ==================== */
:deep(.el-dropdown) {
  color: inherit;
}

:deep(.el-dropdown:focus-visible) {
  outline: none;
}

:deep(.el-dropdown-menu__item) {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  padding: var(--spacing-sm) var(--spacing-md);
  font-size: var(--text-sm);
  color: var(--color-text-secondary);
  transition: all var(--transition-base);
}

:deep(.el-dropdown-menu__item:hover) {
  background-color: var(--color-bg);
  color: var(--color-primary);
}

:deep(.el-dropdown-menu__item .el-icon) {
  font-size: 16px;
}

/* ==================== Dialog Footer ==================== */
.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: var(--spacing-sm);
}

/* ==================== Mobile Menu Button ==================== */
.mobile-menu-btn {
  display: none;
  align-items: center;
  justify-content: center;
  width: 44px;
  height: 44px;
  background: linear-gradient(
    135deg,
    rgba(219, 39, 119, 0.08) 0%,
    rgba(202, 138, 4, 0.08) 100%
  );
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  border: 1.5px solid rgba(219, 39, 119, 0.2);
  border-radius: 12px;
  color: #db2777;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  margin-left: 12px;
  position: relative;
  overflow: hidden;
}

.mobile-menu-btn::before {
  content: "";
  position: absolute;
  inset: 0;
  background: radial-gradient(
    circle at center,
    rgba(219, 39, 119, 0.15) 0%,
    transparent 70%
  );
  opacity: 0;
  transition: opacity 0.3s ease;
}

.mobile-menu-btn:hover {
  background: linear-gradient(
    135deg,
    rgba(219, 39, 119, 0.12) 0%,
    rgba(202, 138, 4, 0.12) 100%
  );
  border-color: rgba(219, 39, 119, 0.35);
  box-shadow: 0 4px 16px rgba(219, 39, 119, 0.2);
  transform: translateY(-2px);
}

.mobile-menu-btn:hover::before {
  opacity: 1;
}

.mobile-menu-btn svg {
  position: relative;
  z-index: 1;
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.mobile-menu-btn:hover svg {
  transform: rotate(90deg);
}

/* 菜单打开状态 */
.mobile-menu-btn.menu-open {
  background: linear-gradient(
    135deg,
    rgba(219, 39, 119, 0.15) 0%,
    rgba(202, 138, 4, 0.15) 100%
  );
}

.mobile-menu-btn.menu-open svg {
  transform: rotate(90deg);
}

/* ==================== Mobile Navigation Menu ==================== */
.mobile-nav-menu {
  display: none;
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  background: rgba(20, 20, 28, 0.97);
  backdrop-filter: blur(24px) saturate(180%);
  -webkit-backdrop-filter: blur(24px) saturate(180%);
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.4);
  padding: 12px 16px;
  z-index: 99;
  max-height: calc(100vh - 60px);
  overflow-y: auto;
}

.mobile-nav-item {
  display: block;
  width: 100%;
  padding: 14px 20px;
  font-size: 1rem;
  font-weight: 500;
  font-family: var(--header-font);
  color: var(--color-text-secondary, #94a3b8);
  background: transparent;
  border: none;
  border-radius: 12px;
  text-align: left;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  margin-bottom: 6px;
  position: relative;
  overflow: hidden;
}

.mobile-nav-item::before {
  content: "";
  position: absolute;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  width: 0;
  height: 70%;
  background: linear-gradient(135deg, #db2777 0%, #d97706 100%);
  border-radius: 0 6px 6px 0;
  transition: width 0.3s ease;
}

.mobile-nav-item:last-child {
  margin-bottom: 0;
}

.mobile-nav-item:hover {
  color: #f472b6;
  background: linear-gradient(
    135deg,
    rgba(219, 39, 119, 0.08) 0%,
    rgba(217, 119, 6, 0.06) 100%
  );
  transform: translateX(4px);
  padding-left: 24px;
}

.mobile-nav-item:hover::before {
  width: 4px;
}

.mobile-nav-item.active {
  color: #ffffff;
  background: linear-gradient(
    135deg,
    rgba(219, 39, 119, 0.12) 0%,
    rgba(217, 119, 6, 0.08) 100%
  );
  font-weight: 600;
  padding-left: 24px;
}

.mobile-nav-item.active::before {
  width: 4px;
}

/* Slide down animation */
.slide-down-enter-active,
.slide-down-leave-active {
  transition: all 0.35s cubic-bezier(0.4, 0, 0.2, 1);
}

.slide-down-enter-from,
.slide-down-leave-to {
  opacity: 0;
  transform: translateY(-12px);
}

/* ==================== Responsive ==================== */
@media (max-width: 1200px) {
  .header-left {
    gap: 32px;
  }

  .nav-menu {
    gap: 4px;
  }

  .nav-item {
    padding: 8px 12px;
    font-size: 0.875rem;
  }
}

@media (max-width: 1024px) {
  .global-header {
    padding: 12px 24px;
  }

  .site-title {
    font-size: 1.125rem;
  }

  .nav-item {
    padding: 6px 12px;
    font-size: 0.875rem;
  }

  .header-left {
    gap: 24px;
  }
}

@media (max-width: 768px) {
  .global-header {
    padding: 12px 20px;
    height: 64px;
  }

  .site-title {
    display: none;
  }

  .nav-menu {
    display: none;
  }

  .mobile-menu-btn {
    display: flex;
  }

  .mobile-nav-menu {
    display: block;
    top: 64px;
  }

  .username {
    display: none;
  }

  .logo {
    width: 40px;
    height: 40px;
  }

  .user-profile {
    padding: 6px;
    background: rgba(255, 255, 255, 0.7);
  }

  .header-left {
    gap: 16px;
  }
}

@media (max-width: 576px) {
  .global-header {
    padding: 10px 16px;
    height: 60px;
  }

  .logo {
    width: 38px;
    height: 38px;
  }

  .mobile-nav-menu {
    top: 60px;
    padding: 10px 12px;
  }

  .mobile-nav-item {
    padding: 12px 16px;
    font-size: 0.9375rem;
    margin-bottom: 4px;
  }

  .mobile-menu-btn {
    width: 42px;
    height: 42px;
  }

  .header-left {
    gap: 12px;
  }
}

/* ==================== Animations ==================== */
@keyframes fadeInDown {
  from {
    opacity: 0;
    transform: translateY(-12px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.global-header {
  animation: fadeInDown 0.5s ease-out;
}

/* 移动端菜单项淡入动画 */
@keyframes slideInFromLeft {
  from {
    opacity: 0;
    transform: translateX(-12px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.mobile-nav-item {
  animation: slideInFromLeft 0.3s ease-out;
}

.mobile-nav-item:nth-child(1) {
  animation-delay: 0.05s;
}
.mobile-nav-item:nth-child(2) {
  animation-delay: 0.1s;
}
.mobile-nav-item:nth-child(3) {
  animation-delay: 0.15s;
}
.mobile-nav-item:nth-child(4) {
  animation-delay: 0.2s;
}
.mobile-nav-item:nth-child(5) {
  animation-delay: 0.25s;
}

/* ==================== Settings Dialog ==================== */
.settings-content {
  padding: var(--spacing-lg);
}

.avatar-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: var(--spacing-xl);
  padding-bottom: var(--spacing-lg);
  border-bottom: 1px solid var(--color-border);
}

.avatar-wrapper-large {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  overflow: hidden;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transition: transform var(--transition-base);
}

.avatar-wrapper-large:hover {
  transform: scale(1.05);
}

.avatar-hint {
  margin-top: var(--spacing-sm);
  font-size: var(--text-sm);
  color: var(--color-text-tertiary);
}

.form-section {
  max-width: 400px;
  margin: 0 auto;
}
</style>
