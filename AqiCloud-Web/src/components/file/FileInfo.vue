<template>
  <el-dialog
    v-model="dialogVisible"
    title="文件详情"
    width="560px"
    class="file-info-dialog"
  >
    <!-- 文件图标和名称头部 -->
    <div class="file-header">
      <div class="file-icon" :class="getFileIconClass(fileInfo)">
        <el-icon :size="40">
          <component :is="getFileIcon(fileInfo)" />
        </el-icon>
      </div>
      <div class="file-title">
        <div class="file-name">{{ fileInfo.fileName }}</div>
        <div class="file-meta">
          {{ getFileTypeName(fileInfo) }} ·
          {{ formatFileSize(fileInfo.fileSize) }}
        </div>
      </div>
    </div>

    <el-divider />

    <!-- 文件详细信息 -->
    <el-descriptions :column="1" border class="file-descriptions">
      <el-descriptions-item label="文件ID">
        <span class="code-text">{{ fileInfo.fileId || "-" }}</span>
      </el-descriptions-item>
      <el-descriptions-item label="文件类型">
        <el-tag size="small" :type="getFileTagType(fileInfo)">
          {{ getFileTypeName(fileInfo) }}
        </el-tag>
        <span v-if="fileInfo.fileSuffix" class="file-suffix"
          >.{{ fileInfo.fileSuffix }}</span
        >
      </el-descriptions-item>
      <el-descriptions-item label="文件大小">
        <span class="size-highlight">{{
          formatFileSize(fileInfo.fileSize)
        }}</span>
        <span v-if="fileInfo.fileSize" class="size-bytes"
          >({{ fileInfo.fileSize.toLocaleString() }} 字节)</span
        >
      </el-descriptions-item>
      <el-descriptions-item label="存储位置">
        <div class="storage-info">
          <div v-if="fileInfo.fileBucketName" class="storage-item">
            <span class="label">存储桶:</span>
            <span class="code-text">{{ fileInfo.fileBucketName }}</span>
          </div>
          <div v-if="fileInfo.objectKey" class="storage-item">
            <span class="label">对象键:</span>
            <span class="code-text truncate">{{ fileInfo.objectKey }}</span>
          </div>
        </div>
      </el-descriptions-item>
      <el-descriptions-item label="文件路径">
        <span class="path-text">{{ fileInfo.filePath || "/" }}</span>
      </el-descriptions-item>
      <el-descriptions-item label="文件标识">
        <span class="code-text truncate">{{
          fileInfo.fileIdentifier || "-"
        }}</span>
      </el-descriptions-item>
      <el-descriptions-item label="修改时间">
        <span class="time-text">
          <el-icon><Clock /></el-icon>
          {{ formatDateTime(fileInfo.updateTime || fileInfo.gmtModified) }}
        </span>
      </el-descriptions-item>
    </el-descriptions>
  </el-dialog>
</template>

<script setup lang="ts">
import { computed } from "vue";
import {
  formatFileSize,
  formatDateTime,
  getFileTypeName,
} from "@/utils/format";
import type { FileDTO } from "@/api/types";
import {
  Document,
  Folder,
  Picture,
  VideoCamera,
  Headset,
  Box,
  Clock,
} from "@element-plus/icons-vue";

/**
 * FileInfo 组件 - 文件详情对话框组件
 * 显示文件的详细信息
 */

interface FileInfoProps {
  modelValue: boolean;
  fileInfo: FileDTO;
}

const props = defineProps<FileInfoProps>();

const emit = defineEmits<{
  "update:modelValue": [value: boolean];
}>();

/**
 * 双向绑定对话框的可见状态
 */
const dialogVisible = computed({
  get: () => props.modelValue,
  set: (value) => emit("update:modelValue", value),
});

/**
 * 获取文件图标
 */
const getFileIcon = (file: FileDTO) => {
  if (file.isDir || file.isDirectory) return Folder;
  const suffix = file.fileSuffix?.toLowerCase() || "";
  const type = file.fileType?.toLowerCase() || "";

  if (["jpg", "jpeg", "png", "gif", "webp", "bmp"].includes(suffix))
    return Picture;
  if (["mp4", "avi", "mov", "wmv", "flv", "mkv"].includes(suffix))
    return VideoCamera;
  if (["mp3", "wav", "flac", "aac", "ogg"].includes(suffix)) return Headset;
  if (["zip", "rar", "7z", "tar", "gz"].includes(suffix)) return Box;

  return Document;
};

/**
 * 获取文件图标样式类
 */
const getFileIconClass = (file: FileDTO) => {
  if (file.isDir || file.isDirectory) return "icon-folder";
  const suffix = file.fileSuffix?.toLowerCase() || "";

  if (["jpg", "jpeg", "png", "gif", "webp"].includes(suffix))
    return "icon-image";
  if (["mp4", "avi", "mov", "wmv"].includes(suffix)) return "icon-video";
  if (["mp3", "wav", "flac", "aac"].includes(suffix)) return "icon-audio";
  if (["zip", "rar", "7z", "tar"].includes(suffix)) return "icon-archive";

  return "icon-document";
};

/**
 * 获取文件标签类型
 */
const getFileTagType = (
  file: FileDTO,
): "" | "success" | "warning" | "info" | "danger" => {
  if (file.isDir || file.isDirectory) return "warning";
  const suffix = file.fileSuffix?.toLowerCase() || "";

  if (["jpg", "jpeg", "png", "gif", "webp"].includes(suffix)) return "success";
  if (["mp4", "avi", "mov", "wmv"].includes(suffix)) return "danger";
  if (["mp3", "wav", "flac", "aac"].includes(suffix)) return "";
  if (["zip", "rar", "7z", "tar"].includes(suffix)) return "info";

  return "";
};
</script>

<style scoped>
/* 对话框整体样式 */
:deep(.el-dialog__header) {
  padding: 20px 24px;
  border-bottom: 1px solid var(--el-border-color-lighter);
}

:deep(.el-dialog__title) {
  font-weight: 600;
  font-size: 18px;
  color: var(--el-text-color-primary);
}

:deep(.el-dialog__body) {
  padding: 24px;
}

/* 文件头部区域 */
.file-header {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 8px 0;
}

.file-icon {
  width: 64px;
  height: 64px;
  border-radius: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  background: linear-gradient(135deg, #1A1A24 0%, #22222E 100%);
  color: var(--color-text-tertiary, #64748B);
  transition: all 0.3s ease;
}

.file-icon.icon-folder {
  background: linear-gradient(135deg, #fef9e7 0%, #f9e79f 100%);
  color: #f39c12;
}

.file-icon.icon-image {
  background: linear-gradient(135deg, #e8f8f5 0%, #abebc6 100%);
  color: #27ae60;
}

.file-icon.icon-video {
  background: linear-gradient(135deg, #fdedec 0%, #f5b7b1 100%);
  color: #e74c3c;
}

.file-icon.icon-audio {
  background: linear-gradient(135deg, #f5eef8 0%, #d7bde2 100%);
  color: #8e44ad;
}

.file-icon.icon-archive {
  background: linear-gradient(135deg, #eaf2f8 0%, #aed6f1 100%);
  color: #3498db;
}

.file-icon.icon-document {
  background: linear-gradient(135deg, #f8f9f9 0%, #d5dbdb 100%);
  color: #5d6d7e;
}

.file-title {
  flex: 1;
  min-width: 0;
}

.file-name {
  font-size: 18px;
  font-weight: 600;
  color: var(--el-text-color-primary);
  word-break: break-all;
  line-height: 1.4;
  margin-bottom: 6px;
}

.file-meta {
  font-size: 13px;
  color: var(--el-text-color-secondary);
}

/* 分隔线 */
:deep(.el-divider) {
  margin: 20px 0;
}

/* 描述列表样式 */
.file-descriptions :deep(.el-descriptions__label) {
  width: 100px;
  font-weight: 600;
  color: var(--color-text-secondary, #94A3B8);
  background-color: var(--color-bg-card, #1A1A24);
  padding: 14px 16px;
}

.file-descriptions :deep(.el-descriptions__content) {
  padding: 14px 16px;
  word-break: break-all;
  color: #303133;
}

/* 各类文本样式 */
.code-text {
  font-family: "Consolas", "Monaco", "Courier New", monospace;
  font-size: 12px;
  color: var(--color-text-secondary, #94A3B8);
  background: rgba(255, 255, 255, 0.04);
  padding: 2px 8px;
  border-radius: 4px;
}

.truncate {
  display: inline-block;
  max-width: 300px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  vertical-align: middle;
}

.file-suffix {
  margin-left: 8px;
  color: var(--color-text-tertiary, #64748B);
  font-size: 12px;
}

.size-highlight {
  font-weight: 600;
  color: #667eea;
  font-size: 15px;
}

.size-bytes {
  margin-left: 8px;
  color: var(--color-text-tertiary, #64748B);
  font-size: 12px;
}

.storage-info {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.storage-item {
  display: flex;
  align-items: center;
  gap: 8px;
}

.storage-item .label {
  color: var(--color-text-tertiary, #64748B);
  font-size: 12px;
  min-width: 56px;
}

.path-text {
  color: var(--color-text-secondary, #94A3B8);
  font-family: "Consolas", monospace;
  font-size: 13px;
}

.time-text {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  color: var(--color-text-secondary, #94A3B8);
}

.time-text .el-icon {
  color: var(--color-text-tertiary, #64748B);
}

/* 移动端适配 */
@media (max-width: 768px) {
  :deep(.el-dialog) {
    width: 92% !important;
    margin: 0 auto;
  }

  :deep(.el-dialog__body) {
    padding: 16px;
  }

  .file-header {
    flex-direction: column;
    text-align: center;
    gap: 12px;
  }

  .file-icon {
    width: 56px;
    height: 56px;
  }

  .file-name {
    font-size: 16px;
  }

  .truncate {
    max-width: 200px;
  }

  .file-descriptions :deep(.el-descriptions__label) {
    width: 80px;
    padding: 12px;
  }

  .file-descriptions :deep(.el-descriptions__content) {
    padding: 12px;
  }
}
</style>
