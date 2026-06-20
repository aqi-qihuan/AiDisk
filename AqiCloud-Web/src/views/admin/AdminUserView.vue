<template>
  <div class="admin-user-view">
    <!-- 浮动背景装饰 -->
    <div class="bg-decoration">
      <div class="bg-orb bg-orb-1"></div>
      <div class="bg-orb bg-orb-2"></div>
      <div class="bg-orb bg-orb-3"></div>
    </div>

    <!-- 页面头部 -->
    <div class="page-header glass-card" style="animation-delay: 0s">
      <div class="header-content">
        <div class="icon-wrapper">
          <div class="icon-glow"></div>
          <el-icon :size="28"><UserFilled /></el-icon>
        </div>
        <div class="title-wrapper">
          <h1>用户管理</h1>
          <p class="subtitle">管理系统用户和权限</p>
        </div>
      </div>
      <div class="header-stats">
        <div class="stat-card">
          <span class="stat-number">{{ total }}</span>
          <span class="stat-label">总用户</span>
        </div>
      </div>
    </div>

    <!-- 搜索区域 -->
    <div class="search-section">
      <div class="search-card glass-card" style="animation-delay: 0.1s">
        <div class="search-header">
          <div class="search-title">
            <el-icon><Filter /></el-icon>
            <span>筛选条件</span>
          </div>
        </div>
        <el-form :model="formSearchParams" inline class="search-form">
          <el-form-item>
            <template #label>
              <span class="form-label">
                <el-icon><User /></el-icon>
                用户名
              </span>
            </template>
            <el-input
              v-model="formSearchParams.username"
              placeholder="请输入用户名"
              clearable
              class="search-input"
              :prefix-icon="User"
            />
          </el-form-item>
          <el-form-item>
            <template #label>
              <span class="form-label">
                <el-icon><Document /></el-icon>
                用户ID
              </span>
            </template>
            <el-input
              v-model="formSearchParams.userId"
              placeholder="请输入用户ID"
              clearable
              class="search-input"
              :prefix-icon="Document"
            />
          </el-form-item>
          <el-form-item class="search-actions">
            <DSButton variant="primary" @click="doSearch" class="search-btn">
              <el-icon><Search /></el-icon>
              搜索
            </DSButton>
            <DSButton variant="outline" @click="resetSearch" class="reset-btn">
              <el-icon><Refresh /></el-icon>
              重置
            </DSButton>
          </el-form-item>
        </el-form>
      </div>
    </div>

    <!-- 数据表格 -->
    <div class="table-section">
      <div class="table-card glass-card" style="animation-delay: 0.2s">
        <div class="table-header">
          <div class="table-title">
            <el-icon><User /></el-icon>
            <span>用户列表</span>
          </div>
          <div class="table-actions">
            <DSButton variant="outline" size="sm" @click="loadData">
              <el-icon><Refresh /></el-icon>
              刷新
            </DSButton>
          </div>
        </div>
        <el-table
          :data="dataList"
          style="width: 100%"
          stripe
          highlight-current-row
          class="user-table"
          v-loading="loading"
          element-loading-text="加载中..."
        >
          <el-table-column
            type="index"
            label="序号"
            width="60"
            align="center"
            :index="
              (index: number) =>
                index +
                1 +
                ((searchParams.current ?? 1) - 1) *
                  (searchParams.pageSize ?? 10)
            "
          />
          <el-table-column label="用户信息" min-width="200">
            <template #default="{ row }">
              <div class="user-info-cell">
                <el-image
                  :src="row.avatarUrl || getDefaultAvatar(row.userId)"
                  class="user-avatar"
                  fit="cover"
                />
                <div class="user-details">
                  <div class="user-name">{{ row.username }}</div>
                  <div class="user-email">{{ row.email || "未设置邮箱" }}</div>
                </div>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="存储空间" min-width="240">
            <template #default="{ row }">
              <div class="storage-cell">
                <div class="storage-header">
                  <span class="storage-title">存储空间</span>
                  <span
                    class="storage-percentage"
                    :class="getStoragePercentageClass(row.storage)"
                  >
                    {{ calculateStoragePercentage(row.storage) }}%
                  </span>
                </div>
                <div class="storage-progress-wrapper">
                  <el-progress
                    :percentage="calculateStoragePercentage(row.storage)"
                    :status="getStorageStatus(row.storage)"
                    :stroke-width="6"
                    :show-text="false"
                  />
                </div>
                <div class="storage-details">
                  <div class="storage-item">
                    <span class="storage-label">已使用</span>
                    <span class="storage-value used">{{
                      formatUsedStorage(row.storage)
                    }}</span>
                  </div>
                  <div class="storage-item">
                    <span class="storage-label">总容量</span>
                    <span class="storage-value total">{{
                      formatTotalStorage(row.storage)
                    }}</span>
                  </div>
                </div>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="角色" width="100" align="center">
            <template #default="{ row }">
              <DSTag :variant="getRoleVariant(row.role)" class="role-tag">
                {{ formatRole(row.role) }}
              </DSTag>
            </template>
          </el-table-column>
          <el-table-column label="更新时间" width="160">
            <template #default="{ row }">
              <div class="time-cell">
                <el-icon><Clock /></el-icon>
                <span>{{ formatDate(row.updateTime) }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="创建时间" width="160">
            <template #default="{ row }">
              <div class="time-cell">
                <el-icon><Calendar /></el-icon>
                <span>{{ formatDate(row.createTime) }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column
            label="操作"
            width="100"
            fixed="right"
            align="center"
          >
            <template #default="{ row }">
              <DSButton
                variant="primary"
                size="sm"
                @click="openEditDialog(row)"
                class="edit-btn"
              >
                <el-icon><Edit /></el-icon>
              </DSButton>
            </template>
          </el-table-column>
        </el-table>

        <!-- 空状态 -->
        <div v-if="dataList.length === 0 && !loading" class="empty-state">
          <div class="empty-icon-wrapper">
            <svg
              class="empty-svg"
              viewBox="0 0 96 96"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <defs>
                <linearGradient
                  id="emptyGrad"
                  x1="0%"
                  y1="0%"
                  x2="100%"
                  y2="100%"
                >
                  <stop offset="0%" stop-color="#F472B6" />
                  <stop offset="100%" stop-color="#C084FC" />
                </linearGradient>
              </defs>
              <circle
                cx="48"
                cy="36"
                r="20"
                stroke="url(#emptyGrad)"
                stroke-width="3"
                fill="none"
                opacity="0.6"
              />
              <circle
                cx="48"
                cy="34"
                r="8"
                fill="url(#emptyGrad)"
                opacity="0.3"
              />
              <path
                d="M24 72 C24 60 34 54 48 54 C62 54 72 60 72 72"
                stroke="url(#emptyGrad)"
                stroke-width="3"
                fill="none"
                stroke-linecap="round"
                opacity="0.6"
              />
              <circle cx="40" cy="32" r="2.5" fill="url(#emptyGrad)" />
              <circle cx="56" cy="32" r="2.5" fill="url(#emptyGrad)" />
              <path
                d="M42 40 Q48 44 54 40"
                stroke="url(#emptyGrad)"
                stroke-width="2"
                fill="none"
                stroke-linecap="round"
              />
            </svg>
          </div>
          <p class="empty-text">暂无用户数据</p>
          <p class="empty-subtext">请尝试调整搜索条件</p>
        </div>
      </div>
    </div>

    <!-- 分页 -->
    <div class="pagination-section">
      <div class="pagination-card glass-card" style="animation-delay: 0.3s">
        <div class="pagination-info">
          <span class="pagination-total"
            >共 <strong>{{ total }}</strong> 个用户</span
          >
          <span class="pagination-pages"
            >第 {{ searchParams.current }} /
            {{ Math.ceil(total / (searchParams.pageSize || 10)) || 1 }} 页</span
          >
        </div>
        <el-pagination
          background
          layout="prev, pager, next"
          :total="total"
          :page-size="searchParams.pageSize"
          :current-page="searchParams.current"
          @current-change="onPageChange"
          class="custom-pagination"
        />
      </div>
    </div>

    <!-- 编辑对话框 -->
    <el-dialog
      title="编辑用户信息"
      v-model="editDialogVisible"
      width="600px"
      class="ds-dialog"
    >
      <div class="edit-form-container">
        <el-form :model="editForm" label-width="100px" class="edit-form">
          <div class="form-row">
            <el-form-item label="用户ID">
              <el-input v-model="editForm.userId" disabled />
            </el-form-item>
          </div>

          <div class="form-row">
            <el-form-item label="用户名">
              <el-input
                v-model="editForm.username"
                placeholder="请输入用户名"
              />
            </el-form-item>
          </div>

          <div class="form-row">
            <el-form-item label="邮箱">
              <el-input v-model="editForm.email" placeholder="请输入邮箱" />
            </el-form-item>
          </div>

          <div class="form-row">
            <el-form-item label="容量 (MB)">
              <el-input-number
                v-model="editForm.capacity"
                :min="1"
                :step="100"
                style="width: 100%"
              />
            </el-form-item>
          </div>

          <div class="form-row">
            <el-form-item label="角色">
              <el-select
                v-model="editForm.role"
                placeholder="请选择角色"
                style="width: 100%"
              >
                <el-option label="管理员" :value="1" />
                <el-option label="普通用户" :value="0" />
                <el-option label="封号" :value="-1" />
              </el-select>
            </el-form-item>
          </div>

          <div class="form-row">
            <el-form-item label="头像">
              <div class="avatar-upload">
                <ImageUpload v-model="editForm.avatarUrl" />
              </div>
            </el-form-item>
          </div>
        </el-form>
      </div>

      <template #footer>
        <div class="dialog-footer">
          <DSButton variant="outline" @click="editDialogVisible = false">
            取消
          </DSButton>
          <DSButton
            variant="primary"
            @click="updateUserInfo"
            :loading="updating"
          >
            保存
          </DSButton>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { listUsersByPage, updateUser } from "@/api/user";
import ImageUpload from "@/components/common/ImageUpload.vue";
import { ElMessage } from "element-plus";
import { ref, watchEffect } from "vue";
import {
  UserFilled,
  User,
  Search,
  Refresh,
  Edit,
  Filter,
  Document,
  Folder,
  Clock,
  Calendar,
} from "@element-plus/icons-vue";
import DSButton from "@/components/design-system/DSButton.vue";
import DSTag from "@/components/design-system/DSTag.vue";

const formSearchParams = ref<API.UserQueryRequest>({});
const updating = ref(false);
const loading = ref(false);

// 初始化搜索条件
const initSearchParams = {
  current: 1,
  pageSize: 10,
};

const searchParams = ref<API.UserQueryRequest>({
  ...initSearchParams,
});
const dataList = ref<API.UserDTO[]>([]);
const total = ref<number>(10);

// 编辑对话框相关
const editDialogVisible = ref(false);
const editForm = ref<API.UserUpdateRequest>({});

/**
 * 加载数据
 */
const loadData = async () => {
  loading.value = true;
  try {
    const res = await listUsersByPage(searchParams.value);
    if (res.data.code === 0) {
      dataList.value = res.data?.data?.records || [];
      total.value = Number(res.data?.data?.total) || 0;
    } else {
      ElMessage.error("获取数据失败，" + res.data?.msg);
    }
  } catch (error) {
    ElMessage.error("获取数据失败，请稍后重试");
  } finally {
    loading.value = false;
  }
};

/**
 * 执行搜索
 */
const doSearch = () => {
  searchParams.value = {
    ...initSearchParams,
    ...formSearchParams.value,
  };
};

/**
 * 重置搜索
 */
const resetSearch = () => {
  formSearchParams.value = {};
  searchParams.value = {
    ...initSearchParams,
  };
};

/**
 * 当分页变化时，改变搜索条件，触发数据加载
 * @param page
 */
const onPageChange = (page: number) => {
  searchParams.value = {
    ...searchParams.value,
    current: page,
  };
};

/**
 * 格式化日期
 * @param date
 */
const formatDate = (date: string) => {
  return new Date(date).toLocaleString();
};

/**
 * 格式化角色
 * @param role
 */
const formatRole = (role: number) => {
  switch (role) {
    case 1:
      return "管理员";
    case 0:
      return "普通用户";
    case -1:
      return "封号";
    default:
      return "未知";
  }
};

/**
 * 获取角色标签变体
 * @param role
 */
const getRoleVariant = (role: number) => {
  switch (role) {
    case 1:
      return "success";
    case 0:
      return "info";
    case -1:
      return "danger";
    default:
      return "default";
  }
};

/**
 * 格式化已使用存储
 * @param storage
 */
const formatUsedStorage = (storage: API.StorageDTO | undefined) => {
  if (!storage) return "0 B";
  return calculateStorageSize(storage.storageSize || 0);
};

/**
 * 格式化总存储
 * @param storage
 */
const formatTotalStorage = (storage: API.StorageDTO | undefined) => {
  if (!storage) return "0 B";
  return calculateStorageSize(storage.totalStorageSize || 0);
};

/**
 * 获取存储百分比样式类
 * @param storage
 */
const getStoragePercentageClass = (storage: API.StorageDTO | undefined) => {
  const percentage = Number(calculateStoragePercentage(storage));
  if (percentage >= 80) return "danger";
  if (percentage >= 50) return "warning";
  return "success";
};

/**
 * 计算存储使用百分比
 * @param storage
 */
const calculateStoragePercentage = (storage: API.StorageDTO | undefined) => {
  if (!storage || !storage.totalStorageSize) return 0;
  return (
    ((storage.storageSize || 0) / storage.totalStorageSize) *
    100
  ).toFixed(2);
};

/**
 * 获取存储状态
 * @param storage
 */
const getStorageStatus = (storage: API.StorageDTO | undefined) => {
  const percentage = Number(calculateStoragePercentage(storage));
  if (percentage > 80) return "exception";
  if (percentage > 50) return "warning";
  return "success";
};

/**
 * 计算储存大小
 * @param size
 */
const calculateStorageSize = (size: number | undefined) => {
  if (size === undefined || isNaN(size)) {
    return "未知";
  }

  let sizeNum = Number(size);
  const units = ["B", "KB", "MB", "GB", "TB"];
  let index = 0;
  while (sizeNum >= 1024 && index < units.length - 1) {
    sizeNum /= 1024;
    index++;
  }
  return `${sizeNum.toFixed(index > 0 ? 1 : 0)} ${units[index]}`;
};

/**
 * 获取默认头像
 * @param userId
 */
const getDefaultAvatar = (userId?: string) => {
  const seed = userId || Math.random().toString(36).substring(2, 15);
  return `https://api.dicebear.com/9.x/micah/svg?seed=${seed}`;
};

/**
 * 打开编辑对话框
 * @param user
 */
const openEditDialog = (user: API.UserDTO) => {
  const { createTime, updateTime, storage, ...updateInfo } = user;
  editForm.value = {
    ...updateInfo,
  };
  editDialogVisible.value = true;
};

/**
 * 更新用户信息
 */
const updateUserInfo = async () => {
  updating.value = true;
  try {
    // 确保 editForm 中包含了最新的 avatarUrl
    const updateData: API.UserUpdateRequest = {
      ...editForm.value,
      avatarUrl: editForm.value.avatarUrl, // 确保包含了头像 URL
    };

    const res = await updateUser(updateData);
    if (res.data.code === 0) {
      ElMessage.success("用户信息更新成功");
      editDialogVisible.value = false;
      loadData();
    } else {
      ElMessage.error("更新失败，" + res.data?.msg);
    }
  } catch (error) {
    ElMessage.error("更新失败，请稍后重试");
  } finally {
    updating.value = false;
  }
};

/**
 * 监听 searchParams 变量，改变时触发数据的重新加载
 */
watchEffect(() => {
  loadData();
});
</script>

<style scoped>
.admin-user-view {
  min-height: 100vh;
  padding: var(--ds-spacing-lg);
  background: linear-gradient(135deg, #0B0B10 0%, #1A1A24 100%);
}

/* ===== 页面头部 ===== */
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--spacing-lg);
  padding: var(--spacing-lg);
  background: var(--color-surface);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-md);
  border: 1px solid var(--color-border);
}

.header-content {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.icon-wrapper {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 56px;
  height: 56px;
  background: linear-gradient(
    135deg,
    var(--color-primary) 0%,
    var(--color-secondary) 100%
  );
  border-radius: var(--radius-lg);
  color: white;
  overflow: hidden;
}

.icon-glow {
  position: absolute;
  inset: 0;
  background: radial-gradient(
    circle at 30% 30%,
    rgba(255, 255, 255, 0.3) 0%,
    transparent 50%
  );
  pointer-events: none;
}

.title-wrapper h1 {
  margin: 0;
  font-size: var(--text-xl);
  font-weight: var(--font-semibold);
  color: var(--color-text-primary);
  letter-spacing: var(--tracking-tight);
}

.subtitle {
  margin: var(--spacing-xs) 0 0 0;
  font-size: var(--text-sm);
  color: var(--color-text-secondary);
}

.header-stats {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.stat-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: var(--spacing-sm) var(--spacing-lg);
  background: linear-gradient(
    135deg,
    var(--color-primary) 0%,
    var(--color-secondary) 100%
  );
  border-radius: var(--radius-lg);
  color: white;
  min-width: 80px;
}

.stat-number {
  font-size: var(--text-2xl);
  font-weight: var(--font-bold);
  line-height: 1;
}

.stat-label {
  font-size: var(--text-xs);
  opacity: 0.9;
  margin-top: var(--spacing-xs);
}

/* ===== 搜索区域 ===== */
.search-section {
  margin-bottom: var(--spacing-lg);
}

.search-card {
  background: var(--color-surface);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-sm);
  border: 1px solid var(--color-border);
  overflow: hidden;
  transition: box-shadow var(--transition-base);
}

.search-card:hover {
  box-shadow: var(--shadow-md);
}

.search-header {
  display: flex;
  align-items: center;
  padding: var(--spacing-md) var(--spacing-lg);
  background: linear-gradient(
    135deg,
    rgba(99, 102, 241, 0.05) 0%,
    rgba(129, 140, 248, 0.05) 100%
  );
  border-bottom: 1px solid var(--color-border);
}

.search-title {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  font-size: var(--text-sm);
  font-weight: var(--font-medium);
  color: var(--color-text-primary);
}

.search-title .el-icon {
  color: var(--color-primary);
}

.search-form {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  gap: var(--spacing-md);
  padding: var(--spacing-lg);
}

.form-label {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  font-size: var(--text-sm);
  color: var(--color-text-secondary);
}

.form-label .el-icon {
  color: var(--color-primary);
  font-size: 14px;
}

.search-input {
  width: 220px;
}

.search-input :deep(.el-input__wrapper) {
  border-radius: var(--radius-md);
  box-shadow: 0 0 0 1px var(--color-border) inset;
  transition: all var(--transition-base);
}

.search-input :deep(.el-input__wrapper:hover) {
  box-shadow: 0 0 0 1px var(--color-primary) inset;
}

.search-input :deep(.el-input__wrapper.is-focus) {
  box-shadow:
    0 0 0 2px rgba(99, 102, 241, 0.2) inset,
    0 0 0 1px var(--color-primary) inset;
}

.search-actions {
  margin-left: auto;
  display: flex;
  gap: var(--spacing-sm);
}

.search-btn,
.reset-btn {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
}

/* ===== 表格区域 ===== */
.table-section {
  margin-bottom: var(--spacing-lg);
}

.table-card {
  background: var(--color-surface);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-sm);
  border: 1px solid var(--color-border);
  overflow: hidden;
  transition: box-shadow var(--transition-base);
}

.table-card:hover {
  box-shadow: var(--shadow-md);
}

.table-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-md) var(--spacing-lg);
  background: linear-gradient(
    135deg,
    rgba(99, 102, 241, 0.05) 0%,
    rgba(129, 140, 248, 0.05) 100%
  );
  border-bottom: 1px solid var(--color-border);
}

.table-title {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  font-size: var(--text-sm);
  font-weight: var(--font-medium);
  color: var(--color-text-primary);
}

.table-title .el-icon {
  color: var(--color-primary);
}

.table-actions {
  display: flex;
  gap: var(--spacing-sm);
}

.user-table {
  --el-table-header-bg-color: transparent;
  --el-table-row-hover-bg-color: rgba(99, 102, 241, 0.04);
}

.user-table :deep(.el-table__header th) {
  font-weight: var(--font-semibold);
  color: var(--color-text-primary);
  background: transparent;
  border-bottom: 1px solid var(--color-border);
  padding: var(--spacing-sm) var(--spacing-md);
}

.user-table :deep(.el-table__row) {
  transition: background-color var(--transition-fast);
}

.user-table :deep(.el-table__cell) {
  padding: var(--spacing-md);
}

/* 用户信息单元格 */
.user-info-cell {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.user-avatar {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  border: 2px solid var(--color-border);
  box-shadow: var(--shadow-sm);
  transition: transform var(--transition-base);
}

.user-avatar:hover {
  transform: scale(1.05);
}

.user-details {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.user-name {
  font-weight: var(--font-medium);
  color: var(--color-text-primary);
  font-size: var(--text-sm);
}

.user-email {
  font-size: var(--text-xs);
  color: var(--color-text-tertiary);
}

/* 存储空间单元格 */
.storage-cell {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
  padding: var(--spacing-xs) 0;
}

.storage-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.storage-title {
  font-size: var(--text-xs);
  color: var(--color-text-secondary);
  font-weight: var(--font-medium);
}

.storage-percentage {
  font-size: var(--text-xs);
  font-weight: var(--font-semibold);
  padding: 2px 8px;
  border-radius: var(--radius-full);
  transition: all var(--transition-base);
}

.storage-percentage.success {
  color: var(--color-success);
  background: rgba(16, 185, 129, 0.1);
}

.storage-percentage.warning {
  color: var(--color-warning);
  background: rgba(245, 158, 11, 0.1);
}

.storage-percentage.danger {
  color: var(--color-error);
  background: rgba(239, 68, 68, 0.1);
}

.storage-progress-wrapper {
  width: 100%;
}

.storage-progress-wrapper :deep(.el-progress-bar__outer) {
  border-radius: var(--radius-full);
  background-color: rgba(99, 102, 241, 0.08);
  overflow: hidden;
}

.storage-progress-wrapper :deep(.el-progress-bar__inner) {
  border-radius: var(--radius-full);
  transition: width var(--transition-slow) ease;
}

.storage-details {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: var(--spacing-xs);
  padding-top: var(--spacing-xs);
  border-top: 1px dashed var(--color-border);
}

.storage-item {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.storage-label {
  font-size: 11px;
  color: var(--color-text-tertiary);
}

.storage-value {
  font-size: var(--text-sm);
  font-weight: var(--font-semibold);
}

.storage-value.used {
  color: var(--color-text-primary);
}

.storage-value.total {
  color: var(--color-primary);
}

/* 角色标签 */
.role-tag {
  font-size: var(--text-xs);
  padding: 2px 10px;
}

/* 时间单元格 */
.time-cell {
  display: flex;
  align-items: center;
  gap: var(--spacing-xs);
  font-size: var(--text-xs);
  color: var(--color-text-secondary);
}

.time-cell .el-icon {
  color: var(--color-text-tertiary);
  font-size: 12px;
}

/* 编辑按钮 */
.edit-btn {
  opacity: 0.8;
  transition: all var(--transition-base);
}

.edit-btn:hover {
  opacity: 1;
  transform: scale(1.05);
}

.user-table :deep(.el-table__row:hover) .edit-btn {
  opacity: 1;
}

/* 空状态 */
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: var(--spacing-xxl);
  text-align: center;
}

.empty-icon {
  color: var(--color-text-tertiary);
  opacity: 0.5;
  margin-bottom: var(--spacing-md);
}

.empty-text {
  font-size: var(--text-base);
  font-weight: var(--font-medium);
  color: var(--color-text-secondary);
  margin: 0 0 var(--spacing-xs) 0;
}

.empty-subtext {
  font-size: var(--text-sm);
  color: var(--color-text-tertiary);
  margin: 0;
}

/* ===== 分页区域 ===== */
.pagination-section {
  margin-bottom: var(--spacing-lg);
}

.pagination-card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--spacing-md) var(--spacing-lg);
  background: var(--color-surface);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-sm);
  border: 1px solid var(--color-border);
}

.pagination-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.pagination-total {
  font-size: var(--text-sm);
  color: var(--color-text-secondary);
}

.pagination-total strong {
  color: var(--color-primary);
  font-weight: var(--font-semibold);
}

.pagination-pages {
  font-size: var(--text-xs);
  color: var(--color-text-tertiary);
}

.custom-pagination :deep(.el-pagination__prev),
.custom-pagination :deep(.el-pagination__next),
.custom-pagination :deep(.el-pager li) {
  border-radius: var(--radius-md);
  transition: all var(--transition-base);
}

.custom-pagination :deep(.el-pager li.is-active) {
  background: linear-gradient(
    135deg,
    var(--color-primary) 0%,
    var(--color-secondary) 100%
  );
  box-shadow: var(--shadow-sm);
}

/* ===== 编辑对话框 ===== */
.edit-form-container {
  padding: var(--spacing-lg);
}

.edit-form {
  max-width: 100%;
}

.edit-form :deep(.el-form-item__label) {
  font-weight: var(--font-medium);
  color: var(--color-text-primary);
}

.edit-form :deep(.el-input__wrapper),
.edit-form :deep(.el-input-number .el-input__wrapper) {
  border-radius: var(--radius-md);
  box-shadow: 0 0 0 1px var(--color-border) inset;
  transition: all var(--transition-base);
}

.edit-form :deep(.el-input__wrapper:hover),
.edit-form :deep(.el-input-number .el-input__wrapper:hover) {
  box-shadow: 0 0 0 1px var(--color-primary) inset;
}

.edit-form :deep(.el-input__wrapper.is-focus),
.edit-form :deep(.el-input-number .el-input__wrapper.is-focus) {
  box-shadow:
    0 0 0 2px rgba(99, 102, 241, 0.2) inset,
    0 0 0 1px var(--color-primary) inset;
}

.edit-form :deep(.el-select .el-input__wrapper) {
  border-radius: var(--radius-md);
}

.form-row {
  margin-bottom: var(--spacing-md);
}

.avatar-upload {
  width: 120px;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: var(--spacing-md);
  padding: var(--spacing-md) var(--spacing-lg);
  border-top: 1px solid var(--color-border);
}

/* 对话框样式优化 */
:deep(.ds-dialog .el-dialog__header) {
  padding: var(--spacing-lg);
  border-bottom: 1px solid var(--color-border);
  margin-right: 0;
}

:deep(.ds-dialog .el-dialog__title) {
  font-weight: var(--font-semibold);
  color: var(--color-text-primary);
}

:deep(.ds-dialog .el-dialog__body) {
  padding: 0;
}

:deep(.ds-dialog .el-dialog__footer) {
  padding: 0;
}

/* ===== 页面容器 ===== */
.admin-user-view {
  min-height: 100vh;
  padding: var(--spacing-lg);
  background: var(--color-bg);
}

/* ===== 动画 ===== */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.page-header,
.search-section,
.table-section,
.pagination-section {
  animation: fadeInUp var(--transition-slow) ease forwards;
}

.search-section {
  animation-delay: 0.1s;
}

.table-section {
  animation-delay: 0.2s;
}

.pagination-section {
  animation-delay: 0.3s;
}

/* ===== 响应式设计 ===== */
@media (max-width: 1200px) {
  .search-input {
    width: 200px;
  }
}

@media (max-width: 968px) {
  .page-header {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--spacing-md);
  }

  .header-stats {
    width: 100%;
    justify-content: flex-start;
  }

  .search-form {
    flex-direction: column;
    align-items: stretch;
  }

  .search-input {
    width: 100%;
  }

  .search-actions {
    margin-left: 0;
    width: 100%;
    justify-content: flex-end;
  }

  .pagination-card {
    flex-direction: column;
    gap: var(--spacing-md);
  }
}

@media (max-width: 768px) {
  .admin-user-view {
    padding: var(--spacing-md);
  }

  .icon-wrapper {
    width: 48px;
    height: 48px;
  }

  .icon-wrapper :deep(.el-icon) {
    font-size: 24px;
  }

  .title-wrapper h1 {
    font-size: var(--text-lg);
  }

  .subtitle {
    font-size: var(--text-xs);
  }

  .search-card,
  .table-card,
  .pagination-card {
    border-radius: var(--radius-lg);
  }

  .user-info-cell {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--spacing-sm);
  }

  .user-avatar {
    width: 36px;
    height: 36px;
  }
}

@media (max-width: 576px) {
  .page-header {
    padding: var(--spacing-md);
  }

  .header-content {
    flex-direction: column;
    align-items: flex-start;
    text-align: left;
    gap: var(--spacing-sm);
  }

  .title-wrapper h1 {
    font-size: var(--text-base);
  }

  .subtitle {
    font-size: var(--text-xs);
  }

  .stat-card {
    padding: var(--spacing-xs) var(--spacing-md);
    min-width: 60px;
  }

  .stat-number {
    font-size: var(--text-lg);
  }

  .table-header {
    flex-direction: column;
    gap: var(--spacing-sm);
    align-items: flex-start;
  }

  .pagination-info {
    flex-direction: column;
    gap: var(--spacing-xs);
    align-items: center;
  }
}
</style>
