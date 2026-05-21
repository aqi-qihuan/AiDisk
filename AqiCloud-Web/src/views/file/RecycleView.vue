<template>
  <div class="recycle-view">
    <!-- 浮动背景装饰 -->
    <div class="bg-decoration">
      <div class="bg-orb bg-orb-1"></div>
      <div class="bg-orb bg-orb-2"></div>
      <div class="bg-orb bg-orb-3"></div>
    </div>
    <!-- 顶部操作栏卡片（有文件时才显示） -->
    <div v-if="fileList.length > 0" class="ds-card operation-card">
      <div class="operation-bar">
        <DSButton
          variant="success"
          size="large"
          :disabled="!selectedFiles.length"
          @click="handleBatchRestore"
          :loading="restoring"
        >
          <RefreshLeft />
          {{ t("file.restoreRecycle")
          }}<template v-if="selectedFiles.length">
            ({{ selectedFiles.length }})</template
          >
        </DSButton>

        <DSButton
          variant="danger"
          size="large"
          :disabled="!selectedFiles.length"
          @click="handleBatchDelete"
          :loading="deleting"
        >
          <Delete />
          {{ t("file.deletePermanentlyRecycle")
          }}<template v-if="selectedFiles.length">
            ({{ selectedFiles.length }})</template
          >
        </DSButton>

        <div class="info-tags">
          <DSTag color="info">
            <svg
              width="16"
              height="16"
              viewBox="0 0 24 24"
              fill="none"
              style="vertical-align: -3px; margin-right: 4px"
            >
              <path
                d="M22 19a2 2 0 01-2 2H4a2 2 0 01-2-2V5a2 2 0 012-2h5l2 3h9a2 2 0 012 2z"
                stroke="currentColor"
                stroke-width="1.5"
              /></svg
            >{{ t("file.totalFiles", { count: fileList.length }) }}
          </DSTag>
          <DSTag color="warning" v-if="selectedFiles.length > 0">
            <svg
              width="16"
              height="16"
              viewBox="0 0 24 24"
              fill="none"
              style="vertical-align: -3px; margin-right: 4px"
            >
              <polyline
                points="20 6 9 17 4 12"
                stroke="currentColor"
                stroke-width="2.5"
                stroke-linecap="round"
                stroke-linejoin="round"
              /></svg
            >{{ t("file.selectedFiles", { count: selectedFiles.length }) }}
          </DSTag>
        </div>
      </div>
    </div>

    <!-- 文件列表（有数据时显示表格） -->
    <div v-if="fileList.length > 0" class="ds-card file-list-card">
      <FileTable :fileList="fileList" @selection-change="handleSelectionChange">
        <template #operation="{ row }">
          <div class="operation-buttons">
            <DSButton
              variant="primary"
              size="sm"
              @click="handleRestore(row)"
              :loading="row.restoring"
            >
              <RefreshLeft />
              {{ t("file.restoreRecycle") }}
            </DSButton>
            <DSButton
              variant="danger"
              size="sm"
              @click="handleDelete(row)"
              :loading="row.deleting"
            >
              <Delete />
              {{ t("file.delete") }}
            </DSButton>
          </div>
        </template>
      </FileTable>
    </div>

    <!-- 空状态（无数据时独立显示） -->
    <div v-if="fileList.length === 0" class="ds-card file-list-card">
      <div class="empty-state">
        <div class="empty-icon">
          <svg
            class="empty-svg"
            width="120"
            height="120"
            viewBox="0 0 120 120"
            fill="none"
          >
            <defs>
              <linearGradient
                id="recycleGrad1"
                x1="0%"
                y1="0%"
                x2="100%"
                y2="100%"
              >
                <stop offset="0%" stop-color="#D4A853" />
                <stop offset="100%" stop-color="#C9A96E" />
              </linearGradient>
              <linearGradient
                id="recycleGrad2"
                x1="0%"
                y1="0%"
                x2="100%"
                y2="100%"
              >
                <stop offset="0%" stop-color="#B8943F" />
                <stop offset="100%" stop-color="#D4A853" />
              </linearGradient>
              <filter id="glow" x="-50%" y="-50%" width="200%" height="200%">
                <feGaussianBlur stdDeviation="3" result="coloredBlur" />
                <feMerge>
                  <feMergeNode in="coloredBlur" />
                  <feMergeNode in="SourceGraphic" />
                </feMerge>
              </filter>
            </defs>
            <!-- 外圈装饰 -->
            <circle
              cx="60"
              cy="60"
              r="55"
              stroke="url(#recycleGrad2)"
              stroke-width="2"
              opacity="0.3"
            />
            <!-- 垃圾桶主体 -->
            <path
              d="M35 40h50l-5 60H40l-5-60z"
              fill="rgba(20,22,40,0.5)"
              stroke="url(#recycleGrad1)"
              stroke-width="2"
            />
            <!-- 垃圾桶盖 -->
            <path d="M30 35h60v5H30z" fill="url(#recycleGrad2)" opacity="0.6" />
            <rect
              x="50"
              y="28"
              width="20"
              height="8"
              rx="2"
              fill="url(#recycleGrad1)"
            />
            <!-- 内部线条装饰 -->
            <line
              x1="45"
              y1="50"
              x2="48"
              y2="90"
              stroke="url(#recycleGrad2)"
              stroke-width="2"
              stroke-linecap="round"
              opacity="0.6"
            />
            <line
              x1="60"
              y1="50"
              x2="60"
              y2="90"
              stroke="url(#recycleGrad2)"
              stroke-width="2"
              stroke-linecap="round"
              opacity="0.6"
            />
            <line
              x1="75"
              y1="50"
              x2="72"
              y2="90"
              stroke="url(#recycleGrad2)"
              stroke-width="2"
              stroke-linecap="round"
              opacity="0.6"
            />
            <!-- 勾选标记 -->
            <circle
              cx="85"
              cy="35"
              r="15"
              fill="rgba(212, 168, 83, 0.15)"
              stroke="#D4A853"
              stroke-width="2"
            />
            <path
              d="M77 35l5 5 10-10"
              stroke="#D4A853"
              stroke-width="3"
              stroke-linecap="round"
              stroke-linejoin="round"
              fill="none"
            />
          </svg>
        </div>
        <p class="empty-title">{{ t("file.recycleEmpty") }}</p>
        <p class="empty-description">{{ t("file.recycleEmptyDesc") }}</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useI18n } from "vue-i18n";
import { ElMessage, ElMessageBox } from "element-plus";
import { Delete, RefreshLeft } from "@element-plus/icons-vue";
import { listRecycleFiles, batchRestore, batchDelete } from "@/api/recycle";
import FileTable from "@/components/file/FileTable.vue";
import { useStorageStore } from "@/store/storage";
import DSButton from "@/components/design-system/DSButton.vue";
import DSTag from "@/components/design-system/DSTag.vue";

const { t } = useI18n();

const fileList = ref<API.FileDTO[]>([]);
const selectedFiles = ref<API.FileDTO[]>([]);
const storageStore = useStorageStore();
const restoring = ref(false);
const deleting = ref(false);

// 获取回收站文件列表
const fetchRecycleFiles = async () => {
  try {
    const response = await listRecycleFiles();
    if (response.data?.success) {
      fileList.value = (response.data?.data ?? []).map((file: API.FileDTO) => ({
        ...file,
        fileType: file.isDir === 1 ? "folder" : file.fileType || "other",
        fileSuffix: file.isDir === 1 ? "folder" : file.fileSuffix || "other",
        restoring: false,
        deleting: false,
      }));
    } else {
      ElMessage.error(response.data?.msg || "获取回收站文件列表失败");
    }
  } catch (error) {
    console.error("获取回收站文件列表失败:", error);
    ElMessage.error("获取回收站文件列表失败");
  }
};

// 处理文件选择
const handleSelectionChange = (selection: API.FileDTO[]) => {
  selectedFiles.value = selection;
};

// 还原单个文件
const handleRestore = async (file: API.FileDTO) => {
  file.restoring = true;
  try {
    const response = await batchRestore({
      fileIds: [String(file.id)],
    });

    if (response.data?.success) {
      ElMessage.success("文件还原成功");
      fetchRecycleFiles();
      storageStore.updateStorageInfo();
    } else {
      ElMessage.error(response.data?.msg || "文件还原失败");
    }
  } catch (error) {
    console.error("文件还原失败:", error);
    ElMessage.error("文件还原失败");
  } finally {
    file.restoring = false;
  }
};

// 批量还原文件
const handleBatchRestore = async () => {
  if (!selectedFiles.value.length) return;

  restoring.value = true;
  try {
    const response = await batchRestore({
      fileIds: selectedFiles.value.map((file) => String(file.id)),
    });

    if (response.data?.success) {
      ElMessage.success("文件还原成功");
      fetchRecycleFiles();
      selectedFiles.value = [];
      storageStore.updateStorageInfo();
    } else {
      ElMessage.error(response.data?.msg || "文件还原失败");
    }
  } catch (error) {
    console.error("文件还原失败:", error);
    ElMessage.error("文件还原失败");
  } finally {
    restoring.value = false;
  }
};

// 彻底删除单个文件
const handleDelete = async (file: API.FileDTO) => {
  file.deleting = true;
  try {
    await ElMessageBox.confirm(
      "此操作将永久删除该文件, 是否继续?",
      "永久删除警告",
      {
        confirmButtonText: "确定删除",
        cancelButtonText: "取消",
        type: "warning",
        customClass: "ds-dialog",
      },
    );

    const response = await batchDelete({
      fileIds: [String(file.id)],
    });

    if (response.data?.success) {
      ElMessage.success("文件已永久删除");
      fetchRecycleFiles();
    } else {
      ElMessage.error(response.data?.msg || "文件删除失败");
    }
  } catch (error) {
    if (error !== "cancel") {
      console.error("文件删除失败:", error);
      ElMessage.error("文件删除失败");
    }
  } finally {
    file.deleting = false;
  }
};

// 批量彻底删除文件
const handleBatchDelete = async () => {
  if (!selectedFiles.value.length) return;

  deleting.value = true;
  try {
    await ElMessageBox.confirm(
      `此操作将永久删除选中的 ${selectedFiles.value.length} 个文件, 是否继续?`,
      "批量永久删除警告",
      {
        confirmButtonText: "确定删除",
        cancelButtonText: "取消",
        type: "warning",
        customClass: "ds-dialog",
      },
    );

    const response = await batchDelete({
      fileIds: selectedFiles.value.map((file) => String(file.id)),
    });

    if (response.data?.success) {
      ElMessage.success("文件已永久删除");
      fetchRecycleFiles();
      selectedFiles.value = [];
    } else {
      ElMessage.error(response.data?.msg || "文件删除失败");
    }
  } catch (error) {
    if (error !== "cancel") {
      console.error("文件删除失败:", error);
      ElMessage.error("文件删除失败");
    }
  } finally {
    deleting.value = false;
  }
};

onMounted(() => {
  fetchRecycleFiles();
});
</script>

<style scoped>
.recycle-view {
  padding: var(--ds-spacing-4);
  background: linear-gradient(160deg, #0a0a1a 0%, #141428 40%, #1a1530 100%);
  min-height: calc(100vh - 64px);
  position: relative;
  overflow: hidden;
}

/* 背景装饰球 */
.bg-decoration {
  position: absolute;
  inset: 0;
  overflow: hidden;
  pointer-events: none;
  z-index: 0;
}

.bg-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(60px);
  opacity: 0.4;
  animation: floatOrb 20s ease-in-out infinite;
}

.bg-orb-1 {
  width: 300px;
  height: 300px;
  background: radial-gradient(
    circle,
    rgba(212, 168, 83, 0.12) 0%,
    transparent 70%
  );
  top: -100px;
  right: -50px;
  animation-delay: 0s;
}

.bg-orb-2 {
  width: 250px;
  height: 250px;
  background: radial-gradient(
    circle,
    rgba(201, 169, 110, 0.08) 0%,
    transparent 70%
  );
  bottom: -80px;
  left: -60px;
  animation-delay: -7s;
}

.bg-orb-3 {
  width: 200px;
  height: 200px;
  background: radial-gradient(
    circle,
    rgba(184, 148, 63, 0.08) 0%,
    transparent 70%
  );
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  animation-delay: -14s;
}

@keyframes floatOrb {
  0%,
  100% {
    transform: translate(0, 0) scale(1);
  }
  33% {
    transform: translate(30px, -20px) scale(1.05);
  }
  66% {
    transform: translate(-20px, 15px) scale(0.95);
  }
}

/* 操作卡片 */
.operation-card {
  margin-bottom: var(--ds-spacing-4);
  padding: var(--ds-spacing-3) var(--ds-spacing-4);
  border-left: 4px solid #d4a853;
  background: rgba(20, 22, 40, 0.7);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border: 1px solid rgba(212, 168, 83, 0.12);
  animation: fadeInUp 0.5s ease-out;
  position: relative;
  z-index: 1;
}

.operation-bar {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-3);
  flex-wrap: wrap;
}

.info-tags {
  margin-left: auto;
  display: flex;
  gap: var(--ds-spacing-2);
}

/* 文件列表卡片 */
.file-list-card {
  padding: var(--ds-spacing-4);
  background: rgba(20, 22, 40, 0.7);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border: 1px solid rgba(212, 168, 83, 0.1);
  animation: fadeInUp 0.5s ease-out 0.1s both;
  position: relative;
  z-index: 1;
}

/* 操作按钮组 */
.operation-buttons {
  display: flex;
  gap: var(--ds-spacing-2);
}

/* 空状态 */
.empty-state {
  text-align: center;
  padding: var(--ds-spacing-12) var(--ds-spacing-4);
}

.empty-icon {
  font-size: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: var(--ds-spacing-4);
  animation: float 3s ease-in-out infinite;
  filter: drop-shadow(0 10px 20px rgba(212, 168, 83, 0.1));
}

.empty-svg {
  opacity: 0;
  animation: fadeInScale 0.6s ease-out forwards 0.3s;
}

.empty-title {
  font-size: var(--ds-text-size-lg);
  font-weight: 600;
  color: #e8d5b0;
  margin-bottom: var(--ds-spacing-2);
}

.empty-description {
  font-size: var(--ds-text-size-sm);
  color: #8b8878;
}

/* 浮动动画 */
@keyframes float {
  0%,
  100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
}

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

@keyframes fadeInScale {
  from {
    opacity: 0;
    transform: scale(0.8);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}

/* 响应式 */
@media (max-width: 640px) {
  .recycle-view {
    padding: var(--ds-spacing-2);
  }

  .operation-card {
    padding: var(--ds-spacing-2) var(--ds-spacing-3);
  }

  .operation-bar {
    flex-direction: column;
    align-items: stretch;
  }

  .info-tags {
    margin-left: 0;
    margin-top: var(--ds-spacing-2);
  }

  .operation-buttons {
    flex-direction: column;
  }

  .empty-icon {
    font-size: 60px;
  }
}
</style>
