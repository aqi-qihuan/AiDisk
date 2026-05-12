<template>
  <div class="file-view">
    <div class="file-view-bg">
      <div class="bg-orb bg-orb-1"></div>
      <div class="bg-orb bg-orb-2"></div>
    </div>
    <div class="flex-grow overflow-hidden">
      <!-- 操作栏 -->
      <OperationBar
        :currentPath="currentPath"
        :selectedFiles="selectedFiles"
        :fileType="fileType"
        :viewMode="fileStore.viewMode"
        :isSelectionMode="isSelectionMode"
        @refresh="refreshFiles"
        @viewModeChange="handleViewModeChange"
        @selectionModeChange="handleSelectionModeChange"
        @openUploadDialog="openUploadDialog"
        @openCreateFolderDialog="openCreateFolderDialog"
        @search="handleSearch"
        @batchDelete="handleBatchDelete"
        @batchMove="openMoveDialog"
        @batchCopy="openCopyDialog"
        @shareFiles="openShareDialog"
        @batchDownload="handleBatchDownload"
      />

      <!-- 面包屑导航 -->
      <BreadCrumb
        :currentPath="currentPath"
        :fileType="fileType"
        @pathChange="handlePathChange"
      />

      <!-- 文件列表容器 -->
      <div class="file-list-container overflow-auto">
        <div class="file-list-wrapper glass-panel">
          <!-- 表格视图 -->
          <FileTable
            v-if="fileStore.viewMode === 'table'"
            :key="'table-' + fileStore.viewMode"
            :fileList="fileList"
            :currentPath="currentPath"
            @refresh="refreshFiles"
            @fileClick="handleFileClick"
            @selectionChange="handleSelectionChange"
            @openRenameDialog="openRenameDialog"
            @moveFile="openMoveDialog"
            @copyFile="openCopyDialog"
            @deleteFile="handleDeleteFile"
            @shareFiles="openShareDialog"
            @openFileInfo="openFileInfo"
          />

          <!-- 网格视图 -->
          <FileGrid
            v-else
            :key="'grid-' + fileStore.viewMode"
            :fileList="fileList"
            :isSelectionMode="isSelectionMode"
            @refresh="refreshFiles"
            @fileClick="handleFileClick"
            @selectionChange="handleSelectionChange"
          />
        </div>
      </div>
    </div>

    <!-- 分页容器 -->
    <div class="pagination-container glass-bar">
      <div class="pagination-wrapper">
        <DSPagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :total="total"
          :page-sizes="[10, 20, 50, 100]"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
        />
      </div>
    </div>

    <!-- 上传文件对话框 -->
    <DSDialog
      v-model="uploadDialogVisible"
      title="上传文件"
      :size="isMobile ? 'fullscreen' : 'large'"
    >
      <FileUpload
        :currentPath="currentPath"
        @upload-success="handleUploadSuccess"
      />
    </DSDialog>

    <!-- 新建文件夹对话框 -->
    <DSDialog
      v-model="createFolderDialogVisible"
      title="新建文件夹"
      :size="isMobile ? 'fullscreen' : 'small'"
      :show-footer="true"
      @cancel="createFolderDialogVisible = false"
      @confirm="handleCreateFolder"
      :confirm-loading="loading"
    >
      <div class="form-item">
        <label class="form-label">文件夹名称</label>
        <DSInput
          v-model="folderForm.folderName"
          placeholder="请输入文件夹名称"
          clearable
        />
      </div>
    </DSDialog>

    <!-- 移动文件对话框 -->
    <DSDialog
      v-model="moveDialogVisible"
      title="移动文件"
      :size="isMobile ? 'fullscreen' : 'large'"
      :show-footer="true"
      @cancel="banch"
      @confirm="handleBatchMove"
      :confirm-loading="loading"
    >
      <div class="form-item">
        <label class="form-label">当前路径</label>
        <DSInput v-model="moveForm.sourcePath" disabled />
      </div>
      <div class="form-item">
        <label class="form-label">目标路径</label>
        <DSInput
          v-model="moveForm.targetPath"
          placeholder="请选择目标路径"
          readonly
        />
      </div>
      <FolderTree @select="handleFolderSelect" />
    </DSDialog>

    <!-- 复制文件对话框 -->
    <DSDialog
      v-model="CopyDialogVisible"
      title="复制文件"
      :size="isMobile ? 'fullscreen' : 'large'"
      :show-footer="true"
      @cancel="cancelCopy"
      @confirm="handleBatchCopy"
      :confirm-loading="loading"
    >
      <div class="form-item">
        <label class="form-label">当前路径</label>
        <DSInput v-model="moveForm.sourcePath" disabled />
      </div>
      <div class="form-item">
        <label class="form-label">目标路径</label>
        <DSInput
          v-model="moveForm.targetPath"
          placeholder="请选择目标路径"
          readonly
        />
      </div>
      <FolderTree @select="handleFolderSelect" />
    </DSDialog>

    <!-- 分享文件对话框 -->
    <DSDialog
      v-model="shareDialogVisible"
      title="分享文件"
      :size="isMobile ? 'fullscreen' : 'medium'"
      :show-footer="true"
      @cancel="shareDialogVisible = false"
      @confirm="handleCreateShare"
      :confirm-loading="loading"
    >
      <div class="form-item">
        <label class="form-label">分享名称</label>
        <DSInput
          v-model="shareForm.shareName"
          placeholder="请输入分享名称"
          clearable
        />
      </div>
      <div class="form-item">
        <label class="form-label">分享类型</label>
        <el-select
          v-model="shareForm.shareType"
          placeholder="请选择分享类型"
          class="w-full"
        >
          <el-option :value="no_code" label="公开分享"></el-option>
          <el-option :value="need_code" label="私密分享"></el-option>
        </el-select>
      </div>
      <div class="form-item">
        <label class="form-label">有效期</label>
        <el-select
          v-model="shareForm.shareDayType"
          placeholder="请选择有效期"
          class="w-full"
        >
          <el-option :value="0" label="永久有效"></el-option>
          <el-option :value="1" label="7天"></el-option>
          <el-option :value="2" label="30天"></el-option>
        </el-select>
      </div>
    </DSDialog>

    <!-- 分享结果对话框 -->
    <DSDialog
      v-model="shareResultDialogVisible"
      title="分享成功"
      :size="isMobile ? 'fullscreen' : 'medium'"
      :show-footer="true"
      @cancel="shareResultDialogVisible = false"
      @confirm="copyAllShareInfo"
      confirm-text="复制全部信息"
      cancel-text="关闭"
    >
      <div class="form-item">
        <label class="form-label">分享链接</label>
        <div class="input-with-button">
          <DSInput v-model="shareUrl" readonly :show-clear="false" />
          <DSButton
            variant="secondary"
            size="small"
            @click="copyToClipboard(shareUrl)"
          >
            复制
          </DSButton>
        </div>
      </div>
      <div v-if="shareResult.shareCode" class="form-item">
        <label class="form-label">分享码</label>
        <div class="input-with-button">
          <DSInput
            v-model="shareResult.shareCode"
            readonly
            :show-clear="false"
          />
          <DSButton
            variant="secondary"
            size="small"
            @click="copyToClipboard(shareResult.shareCode)"
          >
            复制
          </DSButton>
        </div>
      </div>
    </DSDialog>

    <!-- 重命名对话框 -->
    <DSDialog
      v-model="renameDialogVisible"
      title="重命名"
      :size="isMobile ? 'fullscreen' : 'small'"
      :show-footer="true"
      @cancel="renameDialogVisible = false"
      @confirm="handleRename"
      :confirm-loading="loading"
    >
      <div class="form-item">
        <label class="form-label">新名称</label>
        <DSInput
          v-model="renameForm.newFileName"
          placeholder="请输入新名称"
          clearable
        />
      </div>
    </DSDialog>

    <!-- 文件详情 -->
    <FileInfo v-model="fileInfoDialogVisible" :fileInfo="fileInfo" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from "vue";
import { useLoginUserStore } from "@/store/user";
import { ElLoading, ElMessage, ElMessageBox } from "element-plus";
import FileTable from "@/components/file/FileTable.vue";
import FileGrid from "@/components/file/FileGrid.vue";
import BreadCrumb from "@/components/file/BreadCrumb.vue";
import OperationBar from "@/components/file/OperationBar.vue";
import FileUpload from "@/components/file/FileUpload.vue";
import {
  DSTag,
  DSInput,
  DSButton,
  DSDialog,
  DSPagination,
} from "@/components/design-system";
import {
  searchFilesByName,
  listFiles,
  createFolder,
  delFiles,
  moveFiles,
  copyFiles,
  renameFile,
  getFolderContents,
} from "@/api/file";
import { useFileStore } from "@/store/file";
import FolderTree from "@/components/file/FolderTree.vue";
import { createShare } from "@/api/share";
import FileInfo from "@/components/file/FileInfo.vue";
import { useStorageStore } from "@/store/storage";
import request from "@/utils/request";

const props = defineProps<{
  fileType: number | null;
}>();

const fileStore = useFileStore();

// 移动端检测
const isMobile = computed(() => {
  return window.innerWidth <= 768;
});

// 加载状态
const loading = ref(false);

const no_code = ref("no_code");
const need_code = ref("need_code");
// 全量文件列表（从API获取）
const allFiles = ref<API.FileDTO[]>([]);
// 当前显示的分页切片
const fileList = ref<API.FileDTO[]>([]);
const currentPage = ref(1);
const loginUserStore = useLoginUserStore();
const parent_id = ref(localStorage.getItem("parent_id"));
const pageSize = ref(10);
const total = ref(0);
const currentPath = ref("/");
const fileType = ref<number | undefined>(undefined);
const isSelectionMode = ref(false);
const selectedFiles = ref<API.FileDTO[]>([]);
const uploadDialogVisible = ref(false);
const searchQuery = ref("");
const createFolderDialogVisible = ref(false);
const accountId = loginUserStore.loginUser.id;
import { useFileStores } from "@/store/filePath";

const { filePaths } = useFileStores();
const shareUrl = ref("");

const folderForm = ref({
  folderName: "",
});
const moveDialogVisible = ref(false);
const CopyDialogVisible = ref(false);
const moveForm = ref({
  sourcePath: "",
  targetPath: "",
  targetParentId: "",
});
const shareDialogVisible = ref(false);
const shareResultDialogVisible = ref(false);
const shareForm = ref({
  shareName: "",
  shareType: "no_code",
  shareDayType: 0,
});
const shareResult = ref<API.ShareUrlDTO>({});

// 重命名相关的响应式变量
const renameDialogVisible = ref(false);
const renameForm = ref({
  fileId: undefined as string | undefined,
  accountId: loginUserStore.loginUser.id,
  newFileName: "",
});
const currentRenameFile = ref<API.FileDTO | null>(null);

// 文件详情相关的响应式变量
const fileInfoDialogVisible = ref(false);
const fileInfo = ref<API.FileDTO>({});

const storageStore = useStorageStore();

/**
 * 根据条件查询文件列表
 */
async function searchFiles() {
  try {
    const response = await searchFilesByName({
      search: searchQuery.value,
    });

    if (!response.data) {
      ElMessage.error("API 响应为空");
      return;
    }

    if (!response.data.success) {
      ElMessage.error(response.data.msg || "API 请求失败");
      return;
    }

    if (!response.data.data || !Array.isArray(response.data.data)) {
      ElMessage.error("API 数据格式不正确");
      return;
    }

    // 搜索结果也走全量存储 + 客户端分页
    allFiles.value = response.data.data.map((file: API.FileDTO) => ({
      ...file,
      fileSuffix: file.isDir === 1 ? "folder" : file.fileSuffix || "other",
      fileType: file.isDir === 1 ? "folder" : file.fileType || "other",
    }));

    total.value = allFiles.value.length;
    currentPage.value = 1; // 新搜索重置到第1页
    updatePagedData();
  } catch (error) {
    if (error instanceof Error) {
      ElMessage.error(`获取文件列表失败: ${error.message}`);
    } else {
      ElMessage.error("获取文件列表失败");
    }
  }
}

/**
 * 获取文件列表
 */
async function fetchFiles() {
  try {
    const response = await listFiles({
      parent_id: parent_id.value,
    });

    if (!response.data) {
      ElMessage.error("API 响应为空");
      return;
    }

    if (!response.data.success) {
      ElMessage.error(response.data.msg || "API 请求失败");
      return;
    }

    if (!response.data.data || !Array.isArray(response.data.data)) {
      ElMessage.error("API 数据格式不正确");
      return;
    }

    // 存储全量数据
    allFiles.value = response.data.data.map((file: API.FileDTO) => ({
      ...file,
      fileSuffix: file.isDir === 1 ? "folder" : file.fileSuffix || "other",
      fileType: file.isDir === 1 ? "folder" : file.fileType || "other",
    }));

    // 更新总数（基于全量）
    total.value = allFiles.value.length;

    // 计算当前页的切片
    updatePagedData();
  } catch (error) {
    if (error instanceof Error) {
      ElMessage.error(`获取文件列表失败: ${error.message}`);
    } else {
      ElMessage.error("获取文件列表失败");
    }
  }
}

/**
 * 根据当前页码和每页大小，从 allFiles 中切出当前页数据
 */
function updatePagedData() {
  const start = (currentPage.value - 1) * pageSize.value;
  const end = start + pageSize.value;
  fileList.value = allFiles.value.slice(start, end);
}

onMounted(() => {
  fetchFiles();
});

const refreshFiles = () => {
  fetchFiles();
};

const handleSizeChange = (val: number) => {
  pageSize.value = val;
  // 重置到第1页，防止越界
  currentPage.value = 1;
  updatePagedData();
};

const handleCurrentChange = (val: number) => {
  currentPage.value = val;
  updatePagedData();
};

/**
 * 点击文件/文件夹
 * 文件夹：进入目录
 * 文件：打开预览
 */
const handleFileClick = (file: API.FileDTO) => {
  if (file.isDir === 1) {
    // 文件夹：进入目录
    currentPath.value =
      currentPath.value === "/"
        ? `/${file.fileName}/`
        : `${currentPath.value}${file.fileName}/`;
    parent_id.value = String(file.id);
    currentPage.value = 1;
    fetchFiles();
  } else {
    // 文件：打开预览
    openFilePreview(file);
  }
};

/**
 * 打开文件预览
 * 根据文件类型选择不同的预览方式
 */
const openFilePreview = (file: API.FileDTO) => {
  const suffix = (file.fileSuffix || "").toLowerCase();

  // 图片文件 - 使用图片预览
  const imageTypes = ["jpg", "jpeg", "png", "gif", "bmp", "webp", "svg"];
  if (imageTypes.includes(suffix)) {
    previewImage(file);
    return;
  }

  // 视频文件 - 使用视频预览
  const videoTypes = ["mp4", "avi", "mov", "wmv", "flv", "mkv", "webm"];
  if (videoTypes.includes(suffix)) {
    previewVideo(file);
    return;
  }

  // 音频文件 - 使用音频预览
  const audioTypes = ["mp3", "wav", "flac", "aac", "ogg", "wma", "m4a"];
  if (audioTypes.includes(suffix)) {
    previewAudio(file);
    return;
  }

  // 文本文件 - 使用文本预览
  const textTypes = [
    "txt",
    "md",
    "json",
    "xml",
    "html",
    "css",
    "js",
    "ts",
    "java",
    "py",
    "go",
    "c",
    "cpp",
    "h",
  ];
  if (textTypes.includes(suffix)) {
    previewText(file);
    return;
  }

  // PDF 文件 - 使用 PDF 预览
  if (suffix === "pdf") {
    previewPdf(file);
    return;
  }

  // 其他文件 - 提示下载
  ElMessage.info(`该文件类型 (${suffix || "未知"}) 暂不支持预览，请下载后查看`);
};

/**
 * 获取文件直接访问 URL（使用 myAxios 实例，与 PDF 预览同机制）
 */
const getStreamUrl = async (file: API.FileDTO): Promise<string> => {
  const response = await request(`/file/v1/preview`, {
    method: "GET",
    params: { fileId: file.id },
    responseType: "blob",
  });
  const mimeType =
    response.headers["content-type"] || "application/octet-stream";
  const url = URL.createObjectURL(
    new Blob([response.data], { type: mimeType }),
  );
  return url;
};

/**
 * 预览图片
 */
const previewImage = async (file: API.FileDTO) => {
  try {
    const url = await getStreamUrl(file);
    ElMessageBox.alert(
      `<div style="text-align: center;">
        <img src="${url}" style="max-width: 100%; max-height: 70vh; border-radius: 8px; object-fit: contain;" />
      </div>`,
      file.fileName || "图片预览",
      {
        dangerouslyUseHTMLString: true,
        showConfirmButton: true,
        confirmButtonText: "关闭",
        showCancelButton: false,
        customClass: "preview-dialog",
      },
    ).catch(() => {});
  } catch {
    ElMessage.error("图片预览失败");
  }
};

/**
 * 预览视频
 */
const previewVideo = async (file: API.FileDTO) => {
  try {
    const url = await getStreamUrl(file);
    ElMessageBox.alert(
      `<div style="text-align: center;">
        <video controls autoplay style="max-width: 100%; max-height: 70vh; border-radius: 8px;">
          <source src="${url}" />
        </video>
      </div>`,
      file.fileName || "视频预览",
      {
        dangerouslyUseHTMLString: true,
        showConfirmButton: true,
        confirmButtonText: "关闭",
        showCancelButton: false,
        customClass: "preview-dialog",
      },
    ).catch(() => {});
  } catch {
    ElMessage.error("视频预览失败");
  }
};

/**
 * 预览音频
 */
const previewAudio = async (file: API.FileDTO) => {
  try {
    const url = await getStreamUrl(file);
    ElMessageBox.alert(
      `<div style="text-align: center; padding: 20px;">
        <p style="color: #1e1b4b; margin-bottom: 16px; font-weight: 600; font-size: 16px;">${escapeHtml(file.fileName || "")}</p>
        <audio controls autoplay style="width: 100%; max-width: 500px;">
          <source src="${url}" />
        </audio>
      </div>`,
      file.fileName || "音频预览",
      {
        dangerouslyUseHTMLString: true,
        showConfirmButton: true,
        confirmButtonText: "关闭",
        showCancelButton: false,
        customClass: "preview-dialog",
      },
    ).catch(() => {});
  } catch {
    ElMessage.error("音频预览失败");
  }
};

/**
 * 预览文本
 */
const previewText = async (file: API.FileDTO) => {
  try {
    const response = await request(`/file/v1/preview`, {
      method: "GET",
      params: { fileId: file.id },
      responseType: "blob",
    });
    const text = await (response.data as Blob).text();
    ElMessageBox.alert(
      `<div style="text-align: left; padding: 10px;">
        <pre style="background: #1e1b4b; color: #e2e8f0; padding: 16px; border-radius: 8px;
                    max-height: 70vh; overflow: auto; font-size: 13px; line-height: 1.6;
                    white-space: pre-wrap; word-wrap: break-word;">${escapeHtml(text)}</pre>
      </div>`,
      file.fileName || "文本预览",
      {
        dangerouslyUseHTMLString: true,
        showConfirmButton: true,
        confirmButtonText: "关闭",
        showCancelButton: false,
        customClass: "preview-dialog",
      },
    );
  } catch {
    ElMessage.error("文本预览失败");
  }
};

/**
 * 预览 PDF
 */
const previewPdf = async (file: API.FileDTO) => {
  try {
    const response = await request(`/file/v1/preview`, {
      method: "GET",
      params: { fileId: file.id },
      responseType: "blob",
    });
    const url = URL.createObjectURL(
      new Blob([response.data], { type: "application/pdf" }),
    );
    window.open(url, "_blank");
  } catch {
    ElMessage.error("PDF 预览失败");
  }
};

/**
 * HTML 转义
 */
const escapeHtml = (text: string): string => {
  const div = document.createElement("div");
  div.textContent = text;
  return div.innerHTML;
};

const handlePathChange = (newPath: string, ID: string) => {
  currentPath.value = newPath;
  parent_id.value = ID;
  currentPage.value = 1;
  const startIndex = filePaths.indexOf(ID);
  if (startIndex !== -1) {
    useFileStores().filePaths = filePaths.slice(0, startIndex);
  }
  fetchFiles();
};

/**
 * 切换视图模式
 */
const handleViewModeChange = (newMode: "table" | "grid") => {
  fileStore.setViewMode(newMode);
  isSelectionMode.value = false;
  selectedFiles.value = [];
};

const handleSelectionModeChange = (mode: boolean) => {
  isSelectionMode.value = mode;
  if (!mode) {
    selectedFiles.value = [];
  }
};

const handleSelectionChange = (selection: API.FileDTO[]) => {
  selectedFiles.value = selection;
  isSelectionMode.value = selection.length > 0;
};

watch(
  () => props.fileType,
  (newFileType) => {
    fileType.value = newFileType ?? undefined;
    currentPath.value = "/";
    currentPage.value = 1;
    fetchFiles();
  },
  { immediate: true },
);

const openUploadDialog = () => {
  uploadDialogVisible.value = true;
};

const handleUploadSuccess = ({ file, path }: { file: any; path: string }) => {
  refreshFiles();
  storageStore.updateStorageInfo();
};

const handleSearch = (query: string) => {
  searchQuery.value = query;
  currentPage.value = 1; // 新搜索重置到第1页
  if (searchQuery.value == null || searchQuery.value == "") {
    fetchFiles();
  } else {
    searchFiles();
  }
};

const openCreateFolderDialog = () => {
  createFolderDialogVisible.value = true;
};

const handleCreateFolder = async () => {
  if (!folderForm.value.folderName) {
    ElMessage.warning("请输入文件夹名称");
    return;
  }

  loading.value = true;
  try {
    const response = await createFolder({
      folderName: folderForm.value.folderName,
      parentId: String(parent_id.value),
      accountId: accountId,
    });

    if (response.data && response.data.success) {
      ElMessage.success("文件夹创建成功");
      createFolderDialogVisible.value = false;
      folderForm.value.folderName = "";
      refreshFiles();
    } else {
      ElMessage.error(response.data?.msg || "创建文件夹失败");
    }
  } catch (error) {
    ElMessage.error("创建文件夹失败");
  } finally {
    loading.value = false;
  }
};

/**
 * 批量下载 - 使用后端 ZIP 打包下载
 * 支持文件夹递归下载，保持文件夹结构
 */
const handleBatchDownload = async (filesToDownload: API.FileDTO[]) => {
  try {
    const loadingInstance = ElLoading.service({
      lock: true,
      text: `正在准备下载 ${filesToDownload.length} 个文件...`,
      background: "rgba(0, 0, 0, 0.7)",
    });

    // 收集所有文件ID（包括文件夹内的文件）
    const fileIds: string[] = [];

    for (const file of filesToDownload) {
      // 兼容 isDir 和 isDirectory 两种字段名
      const isDirValue = file.isDir ?? file.isDirectory;
      console.log(
        `处理文件/文件夹: ${file.fileName}, isDir: ${file.isDir}, isDirectory: ${file.isDirectory}, 判断结果: ${isDirValue}`,
      );

      // 添加文件/文件夹ID到列表（后端会处理文件夹递归）
      if (file.id) {
        fileIds.push(String(file.id));
        console.log(
          `添加文件/文件夹ID到下载列表: ${file.fileName}, ID: ${file.id}`,
        );
      }
    }

    if (fileIds.length === 0) {
      loadingInstance.close();
      ElMessage.warning("没有可下载的文件");
      return;
    }

    loadingInstance.setText(`正在打包下载 ${fileIds.length} 个文件...`);

    // 使用后端 ZIP 打包下载接口
    const baseUrl = "http://127.0.0.1:8080/api";
    const token = loginUserStore.token || "";

    // 构建表单数据
    const formData = new FormData();
    formData.append("fileIdsStr", fileIds.join(","));

    console.log(`发送批量下载请求，文件IDs:`, fileIds);

    try {
      const response = await fetch(`${baseUrl}/file/v1/batch_download`, {
        method: "POST",
        headers: {
          token: token,
        },
        body: formData,
      });

      console.log(`批量下载响应状态:`, response.status);

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      // 获取文件名（从Content-Disposition头）
      const contentDisposition = response.headers.get("content-disposition");
      let filename = "batch_download.zip";
      if (contentDisposition) {
        const filenameMatch = contentDisposition.match(/filename="(.+)"/);
        if (filenameMatch) {
          filename = filenameMatch[1];
        }
      }

      // 下载ZIP文件
      const blob = await response.blob();
      console.log(`ZIP文件大小:`, blob.size);

      if (blob.size === 0) {
        throw new Error("下载的ZIP文件大小为0");
      }

      // 创建下载链接
      const link = document.createElement("a");
      const url = URL.createObjectURL(blob);
      link.href = url;
      link.download = filename;
      document.body.appendChild(link);
      link.click();

      // 延迟释放资源
      setTimeout(() => {
        window.URL.revokeObjectURL(url);
        if (link.parentNode) {
          document.body.removeChild(link);
        }
      }, 100);

      loadingInstance.close();
      ElMessage.success(`成功下载 ${fileIds.length} 个文件的压缩包`);
      console.log(`批量下载完成: ${filename}`);
    } catch (error) {
      console.error(`批量下载失败:`, error);
      loadingInstance.close();
      ElMessage.error("批量下载失败，请重试");
    }
  } catch (error) {
    console.error("下载文件出错:", error);
    ElMessage.error("下载文件失败");
  }
};

/**
 * 递归收集文件夹内的所有文件
 */
const collectFilesFromFolder = async (
  folder: API.FileDTO,
  fileList: { fileId: string; fileName: string }[],
  folderPath: string,
): Promise<void> => {
  try {
    console.log(`正在获取文件夹内容: ${folder.fileName}, id: ${folder.id}`);

    // 使用 getFolderContents 不修改全局状态
    const response = await getFolderContents({
      parent_id: folder.id || "",
    });

    console.log(`文件夹 ${folder.fileName} 的 API 响应:`, response.data);

    if (!response.data || !response.data.success) {
      console.warn(
        `获取文件夹 ${folder.fileName} 内容失败:`,
        response.data?.msg,
      );
      return;
    }

    const children = response.data.data || [];
    console.log(
      `文件夹 ${folder.fileName} 包含 ${children.length} 个子项:`,
      children,
    );

    // 如果文件夹为空，添加一个标记文件来创建空文件夹结构
    if (children.length === 0) {
      console.log(
        `文件夹 ${folder.fileName} 为空，添加占位文件以创建文件夹结构`,
      );
      fileList.push({
        fileId: "EMPTY_FOLDER_MARKER",
        fileName: `${folderPath}/.gitkeep`,
      });
      return;
    }

    for (const child of children) {
      const childPath = `${folderPath}/${child.fileName}`;
      // 详细打印子项的所有字段，用于排查问题
      console.log(`处理子项详情:`, JSON.stringify(child));
      console.log(
        `处理子项: ${child.fileName}, isDir: ${child.isDir}, isDirectory: ${child.isDirectory}, id: ${child.id}`,
      );

      // 兼容 isDir 和 isDirectory 两种字段名
      const isDirValue = child.isDir ?? child.isDirectory;
      if (isDirValue === 1) {
        // 递归处理子文件夹
        await collectFilesFromFolder(child, fileList, childPath);
      } else {
        // 添加文件到列表 - 保持ID为字符串，避免BigInt精度丢失
        fileList.push({
          fileId: String(child.id),
          fileName: childPath,
        });
        console.log(`添加文件到下载列表: ${childPath}, fileId: ${child.id}`);
      }
    }
  } catch (error) {
    console.error(`获取文件夹 ${folder.fileName} 内容出错:`, error);
  }
};

const handleBatchDelete = async (filesToDelete: API.FileDTO[]) => {
  try {
    const fileIds = filesToDelete.map((file) => String(file.id));

    const requestBody = {
      accountId,
      fileIds,
    };

    const response = await delFiles(requestBody);

    if (response.data && response.data.success) {
      ElMessage.success("文件删除成功");
      refreshFiles();
      isSelectionMode.value = false;
      selectedFiles.value = [];
      storageStore.updateStorageInfo();
    } else {
      ElMessage.error(response.data?.msg || "删除文件失败");
    }
  } catch (error) {
    console.error("删除文件出错:", error);
    ElMessage.error("删除文件失败");
  }
};

const openMoveDialog = (filesToMove: API.FileDTO[], sourcePath: string) => {
  moveDialogVisible.value = true;
  selectedFiles.value = filesToMove;
  moveForm.value.sourcePath = sourcePath;
  moveForm.value.targetPath = sourcePath;
};

const openCopyDialog = (filesToMove: API.FileDTO[], sourcePath: string) => {
  CopyDialogVisible.value = true;
  selectedFiles.value = filesToMove;
  moveForm.value.sourcePath = sourcePath;
  moveForm.value.targetPath = sourcePath;
};

const cancelCopy = async () => {
  CopyDialogVisible.value = false;
};

const handleBatchCopy = async () => {
  if (!moveForm.value.targetPath) {
    ElMessage.warning("请输入目标路径");
    return;
  }

  loading.value = true;
  try {
    const fileIds = selectedFiles.value.map((file) => String(file.id));
    const { filepath } = useFileStores();
    const targetParentId = filepath;
    const response = await copyFiles({
      fileIds,
      accountId,
      targetParentId,
    });
    if (response.data && response.data.success) {
      ElMessage.success("文件复制成功");
      storageStore.updateStorageInfo();
      CopyDialogVisible.value = false;
      moveForm.value.sourcePath = "";
      moveForm.value.targetPath = "";
    } else {
      ElMessage.error(response.data?.msg || "复制文件失败");
    }
  } catch (error) {
    ElMessage.error("复制文件失败");
  } finally {
    loading.value = false;
  }
};

const banch = () => {
  moveDialogVisible.value = false;
};

const handleBatchMove = async () => {
  if (!moveForm.value.targetPath) {
    ElMessage.warning("请输入目标路径");
    return;
  }

  loading.value = true;
  try {
    const fileIds = selectedFiles.value.map((file) => String(file.id));
    const { filepath } = useFileStores();
    const targetParentId = filepath;
    const response = await moveFiles({
      fileIds,
      accountId,
      targetParentId,
    });

    if (response.data && response.data.success) {
      ElMessage.success("文件移动成功");
      moveDialogVisible.value = false;
      moveForm.value.sourcePath = "";
      moveForm.value.targetPath = "";
      refreshFiles();
      isSelectionMode.value = false;
      selectedFiles.value = [];
    } else {
      ElMessage.error(response.data?.msg || "移动文件失败");
    }
  } catch (error) {
    console.error("移动文件出错:", error);
    ElMessage.error("移动文件失败");
  } finally {
    loading.value = false;
  }
};

const handleFolderSelect = (node: API.TreeNodeDTO & { path: string }) => {
  moveForm.value.targetPath = node.path;
  useFileStores().setFilePath(node.id);
};

const openShareDialog = (filesToShare: API.FileDTO[]) => {
  shareDialogVisible.value = true;
  selectedFiles.value = filesToShare;
  shareForm.value.shareName =
    filesToShare.length === 1
      ? (filesToShare[0].fileName ?? "")
      : `${filesToShare[0].fileName ?? ""}等${filesToShare.length}个文件`;
};

const handleCreateShare = async () => {
  loading.value = true;
  try {
    const fileIds = selectedFiles.value.map((file) => String(file.id));
    const host = window.location.host;
    const response = await createShare({
      ...shareForm.value,
      fileIds,
      accountId,
      host,
    });

    if (response.data && response.data.success) {
      shareResult.value = response.data.data as API.ShareUrlDTO;
      shareUrl.value = String(
        shareResult.value.shareUrl?.replace(
          shareResult.value.shareUrl?.toString().split("/")[2],
          window.location.host,
        ),
      );
      shareDialogVisible.value = false;
      shareResultDialogVisible.value = true;
    } else {
      ElMessage.error(response.data?.msg || "创建分享失败");
    }
  } catch (error) {
    console.error("创建分享出错:", error);
    ElMessage.error("创建分享失败");
  } finally {
    loading.value = false;
  }
};

const copyToClipboard = (text: string | undefined) => {
  if (!text) {
    ElMessage.warning("没有可复制的文本");
    return;
  }

  try {
    const textArea = document.createElement("textarea");
    textArea.value = text;
    document.body.appendChild(textArea);
    textArea.select();
    textArea.setSelectionRange(0, 99999);
    let successful = document.execCommand("copy");
    document.body.removeChild(textArea);

    if (successful) {
      ElMessage.success("复制成功");
    } else {
      throw new Error("复制失败");
    }
  } catch (err) {
    console.error("Failed to copy text: ", err);
    ElMessage.error("复制失败，请手动复制");
  }
};

const copyAllShareInfo = () => {
  const shareInfo = `分享链接：${shareResult.value.shareUrl}\n${
    shareResult.value.shareCode ? `分享码：${shareResult.value.shareCode}` : ""
  }`;
  copyToClipboard(shareInfo);
};

// 重命名相关的方法
const openRenameDialog = (file: API.FileDTO) => {
  currentRenameFile.value = file;
  renameDialogVisible.value = true;
  renameForm.value = {
    fileId: file.id,
    newFileName: file.fileName || "",
    accountId: loginUserStore.loginUser.id,
  };
};

const handleRename = async () => {
  if (!renameForm.value.newFileName) {
    ElMessage.warning("请输入新名称");
    return;
  }

  loading.value = true;
  try {
    const response = await renameFile({
      fileId: renameForm.value.fileId,
      newFilename: renameForm.value.newFileName,
      accountId: renameForm.value.accountId,
    });
    // 检查响应是否成功（支持多种返回格式）
    const isSuccess =
      response.data?.success === true ||
      response.data?.code === 200 ||
      response.code === 200;
    if (isSuccess) {
      // 更新当前文件名称
      if (currentRenameFile.value) {
        currentRenameFile.value.fileName = renameForm.value.newFileName;
      }
      ElMessage.success("重命名成功");
      refreshFiles();
      renameDialogVisible.value = false;
    } else {
      ElMessage.error(
        response.data?.msg || response.data?.message || "重命名失败",
      );
    }
  } catch (error: any) {
    ElMessage.error(error?.message || "重命名失败");
  } finally {
    loading.value = false;
  }
};

// 删除单个文件的方法
const handleDeleteFile = async (file: API.FileDTO) => {
  try {
    const requestBody = {
      accountId,
      fileIds: [String(file.id)],
    };

    const response = await delFiles(requestBody);

    if (response.data && response.data.success) {
      ElMessage.success("文件删除成功");
      refreshFiles();
      storageStore.updateStorageInfo();
    } else {
      ElMessage.error(response.data?.msg || "删除文件失败");
    }
  } catch (error) {
    console.error("删除文件出错:", error);
    ElMessage.error("删除文件失败");
  }
};

// 添加格式化文件大小的函数
const formatFileSize = (size: number | undefined) => {
  if (size == null) return "-";
  const units = ["B", "KB", "MB", "GB", "TB"];
  let index = 0;
  let sizeNum = size;
  while (sizeNum >= 1024 && index < units.length - 1) {
    sizeNum /= 1024;
    index++;
  }
  return `${sizeNum.toFixed(2)} ${units[index]}`;
};

// 添加格式化日期时间的函数
const formatDateTime = (timestamp: number | undefined) => {
  if (!timestamp) return "-";
  return new Date(timestamp).toLocaleString();
};

// 添加打开文件详情的方法
const openFileInfo = (file: API.FileDTO) => {
  fileInfo.value = file;
  fileInfoDialogVisible.value = true;
};
</script>

<style scoped>
/* ===== 主容器 ===== */
.file-view {
  height: 100%;
  display: flex;
  flex-direction: column;
  position: relative;
  overflow: hidden;
  background: #f8fafc;
  font-family: var(--font-primary);
}

/* ===== 背景装饰 ===== */
.file-view-bg {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  pointer-events: none;
  overflow: hidden;
  z-index: 0;
}

.bg-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(120px);
  opacity: 0.04;
}

.bg-orb-1 {
  width: 500px;
  height: 500px;
  background: linear-gradient(135deg, #db2777 0%, #f472b6 100%);
  top: -200px;
  right: -100px;
  animation: floatOrb 30s ease-in-out infinite;
}

.bg-orb-2 {
  width: 400px;
  height: 400px;
  background: linear-gradient(135deg, #d97706 0%, #fbbf24 100%);
  bottom: -150px;
  left: -100px;
  animation: floatOrb 25s ease-in-out infinite reverse;
}

/* ===== Light Panel ===== */
.glass-panel {
  background: #ffffff;
  backdrop-filter: blur(24px) saturate(180%);
  -webkit-backdrop-filter: blur(24px) saturate(180%);
  border: 1px solid rgba(0, 0, 0, 0.06);
  border-radius: 20px;
  margin: var(--spacing-md);
  box-shadow:
    0 2px 8px rgba(0, 0, 0, 0.04),
    0 8px 32px rgba(0, 0, 0, 0.03);
  animation: fadeInUp 0.5s ease-out;
  position: relative;
  z-index: 1;
  min-height: 200px;
  transition: box-shadow 0.3s ease;
}

.glass-panel:hover {
  box-shadow:
    0 4px 16px rgba(0, 0, 0, 0.06),
    0 12px 48px rgba(0, 0, 0, 0.04);
}

/* ===== 文件列表容器 ===== */
.file-list-container {
  flex-grow: 1;
  overflow-y: auto;
  position: relative;
  z-index: 1;
}

.file-list-container::-webkit-scrollbar {
  width: 6px;
}

.file-list-container::-webkit-scrollbar-track {
  background: transparent;
}

.file-list-container::-webkit-scrollbar-thumb {
  background: linear-gradient(
    180deg,
    rgba(219, 39, 119, 0.15) 0%,
    rgba(219, 39, 119, 0.05) 100%
  );
  border-radius: 3px;
}

.file-list-container::-webkit-scrollbar-thumb:hover {
  background: linear-gradient(
    180deg,
    rgba(219, 39, 119, 0.3) 0%,
    rgba(219, 39, 119, 0.1) 100%
  );
}

/* ===== 分页容器 Light Bar ===== */
.glass-bar {
  background: #ffffff;
  backdrop-filter: blur(24px) saturate(180%);
  -webkit-backdrop-filter: blur(24px) saturate(180%);
  border-top: 1px solid rgba(0, 0, 0, 0.06);
  box-shadow: 0 -2px 8px rgba(0, 0, 0, 0.03);
  padding: 16px 24px;
  position: relative;
  z-index: 1;
  animation: slideUp 0.4s ease-out;
}

.pagination-container {
  margin-top: auto;
}

.pagination-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: var(--spacing-lg);
  flex-wrap: wrap;
}

/* ===== 表单样式增强 (Dark) ===== */
.form-item {
  margin-bottom: 20px;
  animation: fadeInUp 0.3s ease-out backwards;
}

.form-item:nth-child(1) {
  animation-delay: 0.05s;
}
.form-item:nth-child(2) {
  animation-delay: 0.1s;
}
.form-item:nth-child(3) {
  animation-delay: 0.15s;
}

.form-label {
  display: block;
  font-size: 14px;
  font-weight: 600;
  color: var(--color-text-primary, #1a1a2e);
  margin-bottom: 8px;
  letter-spacing: -0.01em;
}

.form-label.required::after {
  content: " *";
  color: #ef4444;
}

.input-with-button {
  display: flex;
  gap: 12px;
  align-items: center;
}

.input-with-button :deep(.ds-input-wrapper) {
  flex: 1;
}

/* ===== 对话框样式 ===== */
:deep(.ds-dialog .el-dialog__header) {
  background: linear-gradient(135deg, #db2777 0%, #f472b6 50%, #a855f7 100%);
  color: white;
  padding: 20px 24px;
  border-radius: 16px 16px 0 0;
  position: relative;
  overflow: hidden;
}

:deep(.ds-dialog .el-dialog__header::after) {
  content: "";
  position: absolute;
  top: -50%;
  right: -20%;
  width: 200px;
  height: 200px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 50%;
  filter: blur(40px);
}

:deep(.ds-dialog .el-dialog__title) {
  color: white;
  font-weight: 700;
  font-size: 16px;
  letter-spacing: -0.01em;
}

:deep(.ds-dialog .el-dialog__body) {
  padding: 24px 28px;
}

:deep(.ds-dialog .el-dialog__footer) {
  padding: 16px 24px;
  border-top: 1px solid rgba(0, 0, 0, 0.06);
  background: rgba(253, 242, 248, 0.5);
}

:deep(.ds-dialog .el-dialog) {
  border-radius: 16px;
  overflow: hidden;
  box-shadow:
    0 24px 80px rgba(0, 0, 0, 0.12),
    0 8px 32px rgba(219, 39, 119, 0.08);
}

/* ===== 预览对话框 ===== */
:deep(.preview-dialog) {
  border-radius: 20px;
}

:deep(.preview-dialog .el-message-box) {
  border-radius: 20px;
  padding: 32px;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(20px);
  box-shadow:
    0 24px 80px rgba(0, 0, 0, 0.1),
    0 8px 32px rgba(219, 39, 119, 0.08);
}

/* ===== Element Plus 表格穿透 (Light) ===== */
:deep(.el-table) {
  --el-table-bg-color: transparent;
  --el-table-tr-bg-color: transparent;
  --el-table-header-bg-color: rgba(219, 39, 119, 0.04);
  --el-table-row-hover-bg-color: rgba(219, 39, 119, 0.03);
  --el-table-text-color: var(--color-text-primary, #1a1a2e);
  --el-table-header-text-color: var(--color-text-primary, #1a1a2e);
  --el-table-border-color: rgba(0, 0, 0, 0.05);
  font-family: var(--font-primary);
  border-radius: 12px;
}

:deep(.el-table th.el-table__cell) {
  background: rgba(219, 39, 119, 0.04) !important;
  font-weight: 600;
  color: var(--color-text-primary, #1a1a2e);
  font-size: 13px;
  text-transform: uppercase;
  letter-spacing: 0.03em;
}

:deep(.el-table td.el-table__cell) {
  color: var(--color-text-secondary, #475569);
  font-size: 14px;
  border-bottom-color: rgba(0, 0, 0, 0.04);
}

:deep(.el-table .el-table__row) {
  transition: all 0.2s ease;
}

:deep(.el-table .el-table__row:hover > td) {
  background: rgba(219, 39, 119, 0.04) !important;
}

:deep(.el-table .el-table__row.current-row > td) {
  background: rgba(219, 39, 119, 0.06) !important;
  border-left: 3px solid #db2777;
}

/* ===== Element Plus 分页穿透 ===== */
:deep(.el-pagination) {
  --el-pagination-button-bg-color: rgba(255, 255, 255, 0.8);
  font-family: "Inter", sans-serif;
}

:deep(.el-pagination .el-pager li) {
  border-radius: 8px;
  font-weight: 500;
  transition: all 0.2s ease;
}

:deep(.el-pagination .el-pager li.is-active) {
  background: linear-gradient(135deg, #db2777 0%, #f472b6 100%);
  color: white;
  box-shadow: 0 4px 12px rgba(219, 39, 119, 0.3);
}

:deep(.el-pagination .el-pager li:hover) {
  color: #db2777;
}

:deep(.el-pagination button) {
  border-radius: 8px;
}

/* ===== Element Plus Select 穿透 ===== */
:deep(.el-select) {
  --el-select-border-color-hover: #f472b6;
}

:deep(.el-select .el-input__wrapper) {
  border-radius: 10px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.04);
  transition: all 0.2s ease;
}

:deep(.el-select .el-input__wrapper:hover) {
  box-shadow: 0 2px 8px rgba(219, 39, 119, 0.1);
}

/* ===== Element Plus Checkbox 穿透 ===== */
:deep(.el-checkbox__input.is-checked .el-checkbox__inner) {
  background: linear-gradient(135deg, #db2777 0%, #f472b6 100%);
  border-color: #db2777;
}

/* ===== 响应式 ===== */
@media (max-width: 768px) {
  .pagination-wrapper {
    flex-direction: column;
    gap: var(--spacing-md);
    padding: 0 var(--spacing-md);
  }

  .glass-panel {
    margin: var(--spacing-sm);
    border-radius: 16px;
  }

  .glass-bar {
    padding: 12px 16px;
  }

  :deep(.el-pagination) {
    width: 100%;
  }

  :deep(.el-pagination__sizes),
  :deep(.el-pagination__jump) {
    display: none;
  }

  :deep(.el-pager li) {
    min-width: 32px;
    height: 32px;
    line-height: 32px;
    font-size: 13px;
  }

  :deep(.el-pagination button) {
    min-width: 32px;
    height: 32px;
  }
}

/* ===== 动画 ===== */
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

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(16px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(12px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
