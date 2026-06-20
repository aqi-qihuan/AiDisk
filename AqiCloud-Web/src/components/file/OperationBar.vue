<template>
  <div class="operation-bar">
    <!-- 常规操作按钮 -->
    <div v-if="isAllFiles && !isSelectionMode" class="normal-actions">
      <DSButton variant="golden" size="medium" @click="openCreateFolderDialog">
        <template #icon>
          <svg
            width="16"
            height="16"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2.5"
          >
            <path d="M12 5v14M5 12h14" />
          </svg>
        </template>
        {{ t("file.newFolder") }}
      </DSButton>
      <DSButton variant="secondary" size="medium" @click="openUploadDialog">
        <template #icon>
          <svg
            width="16"
            height="16"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
          >
            <path
              d="M21 15v4a2 2 0 01-2 2H5a2 2 0 01-2-2v-4M17 8l-5-5-5 5M12 3v12"
            />
          </svg>
        </template>
        {{ t("file.uploadFile") }}
      </DSButton>
    </div>

    <!-- 批量操作按钮 -->
    <div
      v-if="isSelectionMode && selectedFiles.length > 0"
      class="batch-actions"
    >
      <DSButton variant="primary" size="small" @click="handleBatchDownload">
        <template #icon
          ><el-icon><Download /></el-icon
        ></template>
        {{ t("common.download") }}
      </DSButton>
      <DSButton variant="danger" size="small" @click="handleBatchDelete">
        <template #icon
          ><el-icon><Delete /></el-icon
        ></template>
        {{ t("common.delete") }}
      </DSButton>
      <DSButton variant="secondary" size="small" @click="handleMove">
        <template #icon
          ><el-icon><Position /></el-icon
        ></template>
        {{ t("file.move") }}
      </DSButton>
      <DSButton variant="secondary" size="small" @click="handleCopy">
        <template #icon
          ><el-icon><DocumentCopy /></el-icon
        ></template>
        {{ t("file.copy") }}
      </DSButton>
      <DSButton variant="cta" size="small" @click="handleShare">
        <template #icon
          ><el-icon><Share /></el-icon
        ></template>
        {{ t("file.share") }}
      </DSButton>
    </div>

    <!-- 右侧操作区域 -->
    <div class="right-section">
      <!-- 搜索框 -->
      <div class="search-container">
        <DSInput
          v-model="searchQuery"
          :placeholder="t('file.searchPlaceholder')"
          size="medium"
          clearable
          @input="handleSearch"
        >
          <template #prefix>
            <el-icon><Search /></el-icon>
          </template>
        </DSInput>
      </div>

      <!-- 视图和操作按钮 -->
      <div class="right-actions">
        <el-tooltip :content="t('file.refresh')" placement="top">
          <DSButton variant="ghost" size="small" @click="$emit('refresh')">
            <el-icon><Refresh /></el-icon>
          </DSButton>
        </el-tooltip>

        <el-tooltip
          v-if="viewMode === 'grid'"
          :content="t('file.selectionMode')"
          placement="top"
        >
          <DSButton variant="ghost" size="small" @click="toggleSelectionMode">
            <el-icon><Finished /></el-icon>
          </DSButton>
        </el-tooltip>

        <div class="divider"></div>

        <el-tooltip :content="t('file.listView')" placement="top">
          <DSButton
            variant="ghost"
            size="small"
            @click="handleViewModeChange('table')"
          >
            <el-icon><List /></el-icon>
          </DSButton>
        </el-tooltip>

        <el-tooltip :content="t('file.gridView')" placement="top">
          <DSButton
            variant="ghost"
            size="small"
            @click="handleViewModeChange('grid')"
          >
            <el-icon><Grid /></el-icon>
          </DSButton>
        </el-tooltip>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  Delete,
  Download,
  Finished,
  Folder,
  Grid,
  List,
  Position,
  Refresh,
  Search,
  Share,
  DocumentCopy,
  Upload,
} from "@element-plus/icons-vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { computed, ref } from "vue";
import { useI18n } from "vue-i18n";
import { getFileTree } from "@/api/file";
import type { FileDTO } from "@/api/types";
import { DSButton, DSInput } from "@/components/design-system";

const { t } = useI18n();

/**
 * OperationBar 组件 - 操作栏组件
 * 提供文件上传、新建文件夹、搜索、视图切换等功能
 */

interface OperationBarProps {
  currentPath: string;
  selectedFiles: FileDTO[];
  fileType?: number | null;
  viewMode: "table" | "grid";
  isSelectionMode: boolean;
}

const props = defineProps<OperationBarProps>();

const emit = defineEmits<{
  refresh: [];
  viewModeChange: [mode: "table" | "grid"];
  selectionModeChange: [value: boolean];
  openUploadDialog: [];
  openCreateFolderDialog: [];
  search: [query: string];
  batchDelete: [files: FileDTO[]];
  batchMove: [files: FileDTO[], path: string];
  batchCopy: [files: FileDTO[], path: string];
  shareFiles: [files: FileDTO[]];
  batchDownload: [files: FileDTO[]];
}>();

const searchQuery = ref("");

/**
 * 判断是否显示所有文件
 */
const isAllFiles = computed(
  () => props.fileType === undefined || props.fileType === null,
);

/**
 * 处理视图模式切换
 */
const handleViewModeChange = (newMode: "table" | "grid"): void => {
  emit("viewModeChange", newMode);
};

/**
 * 切换选择模式
 */
const toggleSelectionMode = (): void => {
  emit("selectionModeChange", !props.isSelectionMode);
};

/**
 * 处理批量下载
 */
const handleBatchDownload = (): void => {
  ElMessage.success(`正在下载 ${props.selectedFiles.length} 个文件`);
  emit("batchDownload", props.selectedFiles);
};

/**
 * 处理批量删除
 */
const handleBatchDelete = (): void => {
  ElMessageBox.confirm(
    `确定要删除选中的 ${props.selectedFiles.length} 个文件吗？`,
    "警告",
    {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "warning",
    },
  )
    .then(() => {
      emit("batchDelete", props.selectedFiles);
    })
    .catch(() => {
      // 用户取消操作
    });
};

/**
 * 处理文件移动
 */
const handleMove = (): void => {
  emit("batchMove", props.selectedFiles, props.currentPath);
  getFileTree();
};

/**
 * 处理文件复制
 */
const handleCopy = (): void => {
  emit("batchCopy", props.selectedFiles, props.currentPath);
  getFileTree();
};

/**
 * 处理文件分享
 */
const handleShare = (): void => {
  emit("shareFiles", props.selectedFiles);
};

/**
 * 打开上传对话框
 */
const openUploadDialog = (): void => {
  emit("openUploadDialog");
};

/**
 * 打开新建文件夹对话框
 */
const openCreateFolderDialog = (): void => {
  emit("openCreateFolderDialog");
};

/**
 * 处理搜索
 */
const handleSearch = (): void => {
  emit("search", searchQuery.value);
};
</script>

<style scoped>
.operation-bar {
  margin-bottom: 12px;
  display: flex;
  align-items: center;
  padding: 8px 16px;
  gap: 16px;
}

.normal-actions,
.batch-actions {
  display: flex;
  gap: 8px;
}

/* 发光按钮样式 */
.glowing-btn {
  position: relative;
  overflow: hidden;
  border-radius: 12px;
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  font-weight: 500;
}

.glowing-btn::before {
  content: "";
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: linear-gradient(
    115deg,
    transparent,
    rgba(255, 255, 255, 0.3),
    transparent
  );
  transform: rotate(-45deg);
  animation: glowing 3s linear infinite;
}

.glowing-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.upload-btn,
.folder-btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-color: transparent;
}

@keyframes glowing {
  0% {
    left: -50%;
    top: -50%;
  }
  100% {
    left: 150%;
    top: 150%;
  }
}

/* 批量操作按钮样式 */
.custom-button {
  border-radius: 20px;
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  display: flex;
  align-items: center;
  padding: 8px 16px;
  font-size: 13px;
  font-weight: 500;
  border: none;
  color: white;
}

.custom-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.custom-button.primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.custom-button.danger {
  background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%);
}

.custom-button.warning {
  background: linear-gradient(135deg, #ffecd2 0%, #fcb69f 100%);
  color: #856404;
}

.custom-button.copy {
  background: linear-gradient(135deg, #a1c4fd 0%, #c2e9fb 100%);
  color: #004085;
}

.custom-button.success {
  background: linear-gradient(135deg, #84fab0 0%, #8fd3f4 100%);
  color: #155724;
}

/* 右侧操作区域 */
.right-section {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  gap: 12px;
  margin-left: auto;
}

.search-container {
  flex: 1;
  max-width: 300px;
}

.search-input {
  width: 100%;
  border-radius: 20px;
}

::v-deep(.search-input .el-input__wrapper) {
  border-radius: 20px;
  padding: 0 15px;
}

.right-actions {
  display: flex;
  align-items: center;
}

.right-actions .el-button {
  padding: 8px;
  height: 36px;
  width: 36px;
  border-radius: 8px;
}

.right-actions .el-button:hover {
  background-color: rgba(255, 255, 255, 0.04);
}

.right-actions .el-button.is-active {
  color: #409eff;
  background-color: #ecf5ff;
}

.divider {
  height: 20px;
  margin: 0 4px;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .operation-bar {
    flex-direction: column;
    align-items: stretch;
    padding: 12px;
    gap: 12px;
  }

  .normal-actions,
  .batch-actions {
    flex-wrap: wrap;
    gap: 8px;
    justify-content: center;
  }

  .normal-actions :deep(.ds-button),
  .batch-actions :deep(.ds-button) {
    flex: 1;
    min-width: 120px;
    max-width: 200px;
  }

  .right-section {
    flex-direction: column;
    align-items: stretch;
    margin-left: 0;
    gap: 12px;
  }

  .search-container {
    max-width: 100%;
    order: -1;
  }

  .right-actions {
    justify-content: center;
    gap: 8px;
  }
}

@media (max-width: 480px) {
  .normal-actions :deep(.ds-button),
  .batch-actions :deep(.ds-button) {
    min-width: 100px;
    font-size: 12px;
  }

  .normal-actions :deep(.ds-button) .el-icon,
  .batch-actions :deep(.ds-button) .el-icon {
    font-size: 14px;
  }
}

/* ===== HOK 移动端增强 - 批量操作折叠 ===== */
@media (max-width: 640px) {
  .batch-actions {
    flex-wrap: nowrap;
    overflow-x: auto;
    -webkit-overflow-scrolling: touch;
    scrollbar-width: none;
    -ms-overflow-style: none;
    padding-bottom: 4px;
    gap: 6px;
  }

  .batch-actions::-webkit-scrollbar {
    display: none;
  }

  .batch-actions :deep(.ds-button) {
    flex: 0 0 auto;
    min-width: auto;
    white-space: nowrap;
    font-size: 12px;
    padding: 0 12px;
  }

  .normal-actions :deep(.ds-button) {
    flex: 1;
    min-width: 0;
  }

  /* 小屏隐藏按钮文字只留图标 */
  .normal-actions :deep(.ds-button .ds-btn-text),
  .batch-actions :deep(.ds-button .ds-btn-text) {
    display: none;
  }

  .normal-actions :deep(.ds-button),
  .batch-actions :deep(.ds-button) {
    padding: 0 14px;
    min-width: 44px;
    justify-content: center;
  }
}
</style>
