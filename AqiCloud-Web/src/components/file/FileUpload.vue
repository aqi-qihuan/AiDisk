<template>
  <div class="file-upload">
    <el-upload
      class="upload-area"
      drag
      action="/"
      multiple
      :auto-upload="false"
      :on-change="handleFileChange"
      :on-remove="handleRemoveFile"
      :file-list="fileList"
    >
      <el-icon class="upload-icon"><UploadFilled /></el-icon>
      <div class="upload-text">请拖拽文件到此处或 <em>点击此处上传</em></div>
      <template #tip>
        <div class="upload-tip">支持单个或批量上传，建议单文件不超过 100MB</div>
      </template>
    </el-upload>
    <div class="upload-actions">
      <DSButton
        variant="primary"
        size="large"
        @click="startUpload"
        :loading="isUploading"
      >
        {{ isUploading ? "上传中..." : "开始上传" }}
      </DSButton>
      <DSButton
        variant="secondary"
        size="large"
        @click="clearFiles"
        :disabled="fileList.length === 0"
      >
        清空列表
      </DSButton>
    </div>
  </div>
</template>

<script setup lang="ts">
import { UploadFilled } from "@element-plus/icons-vue";
import { ElNotification } from "element-plus";
import { ref, onBeforeUnmount, computed } from "vue";
import { useLoginUserStore } from "@/store/user";
import {
  initTask,
  merge,
  preSignUploadUrl,
  taskInfo,
  uploadFiles,
  secUpload,
} from "@/api/file";
import { calculateFileMD5 } from "@/utils/md5";
import type { UploadFile } from "element-plus";
import { DSButton } from "@/components/design-system";

/**
 * FileUpload 组件 - 文件上传组件
 * 支持秒传、分片上传、断点续传
 */

interface FileUploadOptions {
  file: File;
  identifier: string;
  filename: string;
  accountId: string;
  parentId: string | null;
  fileSize: number;
}

interface UploadTask {
  id?: string;
  fileName?: string;
  finished?: boolean;
  taskRecord: {
    exitPartList: Array<{ partNumber: number; size: string | number }>;
    chunkSize: number;
    chunkNum: number;
    fileIdentifier: string;
  };
}

interface UploadProgress {
  file: UploadFile;
  progress: number;
}

interface UploadSuccess {
  file: UploadFile;
  path?: string;
}

interface UploadError {
  file: UploadFile;
  error: string;
}

const props = defineProps<{
  currentPath: string;
}>();

const emit = defineEmits<{
  "upload-progress": [data: UploadProgress];
  "upload-success": [data: UploadSuccess];
  "upload-error": [data: UploadError];
}>();

const { loginUser } = useLoginUserStore();
const accountId = String(loginUser.id);
const parentId = localStorage.getItem("parent_id");

const fileList = ref<UploadFile[]>([]);
const fileUploadChunkQueue = ref<Record<string, any>>({});
const isUploading = ref(false);

// 计算属性：是否有文件正在上传
const hasUploadingFiles = computed(() =>
  fileList.value.some((file) => file.status === "uploading"),
);

/**
 * 组件卸载时清理资源
 */
onBeforeUnmount(() => {
  fileList.value = [];
  Object.values(fileUploadChunkQueue.value).forEach((queue) => {
    if (queue?.stop) queue.stop();
  });
});

/**
 * 计算文件的 MD5 标识符
 * @param file - 要计算MD5的文件
 * @returns MD5 字符串
 */
const calculateIdentifier = async (file: File): Promise<string> => {
  return await calculateFileMD5(file);
};

/**
 * 获取或创建上传任务信息
 * @param file - 要上传的文件
 * @returns 上传任务信息或 null
 */
const getTaskInfo = async (file: File): Promise<UploadTask | null> => {
  const identifier = await calculateIdentifier(file);

  // 小于5MB的文件直接上传
  if (file.size <= 5 * 1024 * 1024) {
    const uploadResult = await handleSmallFileUpload(file, identifier);
    return uploadResult.success ? ({ isDirectUpload: true } as any) : null;
  }

  try {
    const res = await taskInfo({ identifier });
    const { code, data } = res.data;

    if (code === 0 && data) {
      return {
        ...data,
        taskRecord: {
          exitPartList: data.exitPartList || [],
          chunkSize: data.chunkSize,
          chunkNum: data.chunkNum,
          fileIdentifier: identifier,
        },
      };
    }

    // 初始化新任务
    const initTaskData = {
      accountId,
      identifier,
      fileName: file.name,
      totalSize: file.size,
      chunkSize: 5 * 1024 * 1024,
    };

    const initRes = await initTask(initTaskData);

    if (initRes.data.code === 0) {
      return {
        ...initRes.data.data,
        taskRecord: {
          exitPartList: [],
          chunkSize: initTaskData.chunkSize,
          chunkNum: Math.ceil(file.size / initTaskData.chunkSize),
          fileIdentifier: identifier,
        },
      };
    } else {
      ElNotification.error({
        title: "文件上传错误",
        message: initRes.data.msg || "初始化分片任务失败",
      });
      return null;
    }
  } catch (error: any) {
    ElNotification.error({
      title: "文件上传错误",
      message: error.message || "初始化分片任务异常",
    });
    return null;
  }
};

/**
 * 处理小于5MB的文件上传
 * @param file - 要上传的文件
 * @param identifier - 文件MD5标识符
 * @returns 上传结果
 */
const handleSmallFileUpload = async (
  file: File,
  identifier: string,
): Promise<{ success: boolean; error?: string }> => {
  try {
    const response = await uploadFiles({
      file,
      identifier,
      filename: file.name,
      accountId,
      parentId,
      fileSize: file.size,
    });

    if (response.data.code === 0) {
      emit("upload-success", { file: file as any });
      ElNotification.success({
        title: "上传成功",
        message: `文件 ${file.name} 上传成功`,
      });

      // 从列表中移除已上传文件
      fileList.value = fileList.value.filter((f) => f.raw !== file);
      return { success: true };
    } else {
      throw new Error(response.data.msg || "上传文件失败");
    }
  } catch (error: any) {
    emit("upload-error", { file: file as any, error: error.message });
    ElNotification.error({
      title: "文件上传错误",
      message: error.message,
    });
    return { success: false, error: error.message };
  }
};

/**
 * 上传单个分片
 * @param partNumber - 分片序号
 * @param file - 文件对象
 * @param fileIdentifier - 文件标识符
 * @param chunkSize - 分片大小
 * @param onProgress - 进度回调
 * @param retryCount - 重试次数
 */
const uploadNext = async (
  partNumber: number,
  file: File,
  fileIdentifier: string,
  chunkSize: number,
  onProgress: (progress: { percent: number }) => void,
  retryCount = 0,
): Promise<void> => {
  try {
    const start = chunkSize * (partNumber - 1);
    const end = Math.min(start + chunkSize, file.size);
    const blob = file.slice(start, end);

    const res = await preSignUploadUrl({
      identifier: fileIdentifier,
      partNumber,
    });

    if (res.data.code === 0 && res.data.data) {
      // 使用 fetch 替代 axios，减少依赖
      const response = await fetch(res.data.data, {
        method: "PUT",
        body: blob,
        headers: { "Content-Type": "application/octet-stream" },
      });

      if (!response.ok) {
        throw new Error(`上传分片失败: ${response.statusText}`);
      }

      onProgress({ percent: Math.ceil((end / file.size) * 100) });
    } else {
      throw new Error(res.data.msg || `获取分片${partNumber}的上传地址失败`);
    }
  } catch (error: any) {
    if (retryCount < 3) {
      return uploadNext(
        partNumber,
        file,
        fileIdentifier,
        chunkSize,
        onProgress,
        retryCount + 1,
      );
    }
    throw error;
  }
};

/**
 * 处理文件上传
 * @param file - 文件对象
 * @param taskRecord - 任务记录
 * @param options - 上传选项
 */
const handleUpload = async (
  file: File,
  taskRecord: UploadTask["taskRecord"],
  options: { onProgress: (progress: { percent: number }) => void },
): Promise<void> => {
  let lastUploadedSize = 0;
  const totalSize = file.size;
  const { exitPartList, chunkSize, chunkNum, fileIdentifier } = taskRecord;

  for (let partNumber = 1; partNumber <= chunkNum; partNumber++) {
    const exitPart = exitPartList?.find(
      (part) => part.partNumber === partNumber,
    );

    if (!exitPart) {
      await uploadNext(
        partNumber,
        file,
        fileIdentifier,
        chunkSize,
        options.onProgress,
      );
    } else {
      lastUploadedSize += Number(exitPart.size);
      options.onProgress({
        percent: Math.ceil((lastUploadedSize / totalSize) * 100),
      });
    }
  }
};

/**
 * 处理文件变更事件
 * @param file - 新增的文件
 * @param uploadFileList - 文件列表
 */
const handleFileChange = (
  file: UploadFile,
  uploadFileList: UploadFile[],
): void => {
  file.percentage = 0;
  file.status = "ready";
  fileList.value = uploadFileList;
};

/**
 * 处理单个文件的上传
 * @param file - 上传文件对象
 */
const handleSingleFileUpload = async (file: UploadFile): Promise<void> => {
  if (!file.raw) return;

  file.status = "uploading";
  emit("upload-progress", { file, progress: 0 });

  try {
    const identifier = await calculateIdentifier(file.raw);

    // 尝试秒传
    const res = await secUpload({
      identifier,
      filename: file.name,
      accountId,
      parentId,
    });

    if (res.data.data !== false) {
      emit("upload-success", { file });
      ElNotification.success({
        title: "秒传成功",
        message: `文件 ${file.name} 已存在，秒传成功`,
      });
      fileList.value = fileList.value.filter((f) => f.uid !== file.uid);
      return;
    }

    // 获取或创建上传任务
    let task = await getTaskInfo(file.raw);

    // 检查是否是直接上传成功的小文件
    if ((task as any)?.isDirectUpload) {
      return;
    }

    // 检查文件是否还在列表中
    if (!fileList.value.find((f) => f.uid === file.uid)) {
      return;
    }

    if (!task) {
      file.status = "error";
      ElNotification.error({
        title: "上传失败",
        message: `文件 ${file.name} 初始化上传任务失败`,
      });
      return;
    }

    if (task.finished) {
      emit("upload-success", {
        file,
        path: `${props.currentPath}${file.name}`,
      });
      ElNotification.success({
        title: "上传成功",
        message: `文件 ${file.name} 已上传完成`,
      });
      fileList.value = fileList.value.filter((f) => f.uid !== file.uid);
      return;
    }

    // 执行分片上传
    await handleUpload(file.raw, task.taskRecord, {
      onProgress: (progress) => {
        file.percentage = progress.percent;
        emit("upload-progress", { file, progress: progress.percent });
      },
    });

    // 合并分片
    const mergeRes = await merge({
      identifier: task.taskRecord.fileIdentifier,
      accountId,
      parentId,
    });

    if (mergeRes.data.code === 0) {
      file.status = "success";
      file.percentage = 100;
      emit("upload-success", {
        file,
        path: `${props.currentPath}${file.name}`,
      });
      ElNotification.success({
        title: "上传成功",
        message: `文件 ${file.name} 上传成功`,
      });
      fileList.value = fileList.value.filter((f) => f.uid !== file.uid);
    } else {
      throw new Error(mergeRes.data.msg || "文件合并失败");
    }
  } catch (error: any) {
    file.status = "error";
    emit("upload-error", { file, error: error.message });
    ElNotification.error({
      title: "上传失败",
      message: error.message || `文件 ${file.name} 上传失败`,
    });
  }
};

/**
 * 开始上传所有文件
 */
const startUpload = async (): Promise<void> => {
  if (fileList.value.length === 0) {
    ElNotification.warning({
      title: "上传提示",
      message: "请先选择要上传的文件",
    });
    return;
  }

  if (isUploading.value) {
    ElNotification.warning({
      title: "上传提示",
      message: "已有文件正在上传，请稍候",
    });
    return;
  }

  isUploading.value = true;

  try {
    for (const file of fileList.value) {
      if (file.status !== "success") {
        await handleSingleFileUpload(file);
      }
    }
  } finally {
    isUploading.value = false;
  }
};

/**
 * 清空文件列表
 */
const clearFiles = (): void => {
  if (isUploading.value) {
    ElNotification.warning({
      title: "操作提示",
      message: "正在上传中，无法清空列表",
    });
    return;
  }
  fileList.value = [];
};

/**
 * 处理文件移除事件
 * @param file - 要移除的文件
 */
const handleRemoveFile = (file: UploadFile): void => {
  const queueObject = fileUploadChunkQueue.value[file.uid];
  if (queueObject?.stop) {
    queueObject.stop();
    fileUploadChunkQueue.value[file.uid] = undefined;
  }

  fileList.value = fileList.value.filter((f) => f.uid !== file.uid);
};
</script>

<style scoped>
.file-upload {
  width: 100%;
  padding: 20px;
  box-sizing: border-box;
}

.upload-area {
  width: 100%;
}

::v-deep(.el-upload) {
  width: 100%;
}

::v-deep(.el-upload-dragger) {
  width: 100%;
  height: 180px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  border: 2px dashed #dcdfe6;
  border-radius: 6px;
  background-color: #fafafa;
  transition: all 0.3s;
}

::v-deep(.el-upload-dragger:hover) {
  border-color: #409eff;
  background-color: #ecf5ff;
}

.upload-icon {
  font-size: 67px;
  color: #c0c4cc;
  margin-bottom: 16px;
  transition: color 0.3s;
}

::v-deep(.el-upload-dragger:hover) .upload-icon {
  color: #409eff;
}

.upload-text {
  color: #606266;
  font-size: 14px;
  text-align: center;
}

.upload-text em {
  color: #409eff;
  font-style: normal;
}

.upload-tip {
  font-size: 12px;
  color: #909399;
  margin-top: 7px;
  text-align: center;
}

.upload-actions {
  display: flex;
  gap: 12px;
  margin-top: 20px;
  justify-content: center;
}

::v-deep(.el-upload-list) {
  margin-top: 16px;
  max-height: 300px;
  overflow-y: auto;
}

::v-deep(.el-upload-list__item) {
  transition: all 0.3s;
}

::v-deep(.el-upload-list__item:hover) {
  background-color: #f5f7fa;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .file-upload {
    padding: 10px;
  }

  ::v-deep(.el-upload-dragger) {
    height: 150px;
  }

  .upload-icon {
    font-size: 50px;
  }

  .upload-text {
    font-size: 12px;
  }

  .upload-actions {
    flex-direction: column;
  }
}
</style>
