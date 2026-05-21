<template>
  <div class="share-view min-h-screen">
    <!-- 浮动背景装饰 -->
    <div class="bg-decoration">
      <div class="bg-orb bg-orb-1"></div>
      <div class="bg-orb bg-orb-2"></div>
      <div class="bg-orb bg-orb-3"></div>
    </div>
    <!-- 验证页面 -->
    <div v-if="needVerify" class="verify-page">
      <div class="verify-left">
        <!-- 用户信息卡片 -->
        <div class="user-card">
          <div class="user-card-avatar">
            <el-avatar
              :size="72"
              :src="shareInfo?.shareAccountDTO?.avatarUrl"
              class="user-avatar"
            />
            <div class="avatar-ring"></div>
          </div>
          <div class="user-card-info">
            <h2 class="user-card-name">
              {{ shareInfo?.shareAccountDTO?.username }}
            </h2>
            <p class="user-card-share">
              分享了
              <span class="share-count">{{
                shareInfo?.fileCount || shareDetail?.fileDTOList?.length || 0
              }}</span>
              个文件
            </p>
          </div>
        </div>

        <!-- 提取码输入区域 -->
        <div class="verify-form">
          <div class="verify-input-wrapper">
            <el-input
              v-model="verifyForm.shareCode"
              placeholder="请输入提取码"
              class="verify-input"
              size="large"
              maxlength="6"
              @keyup.enter="handleVerify"
            >
              <template #prefix>
                <el-icon><Lock /></el-icon>
              </template>
            </el-input>
          </div>
          <DSButton
            variant="primary"
            size="large"
            class="verify-button"
            @click="handleVerify"
            :loading="verifying"
          >
            <el-icon><Unlock /></el-icon>
            提取文件
          </DSButton>
        </div>
      </div>
      <div class="verify-right">
        <div class="illustration-container">
          <img src="@/assets/verify_bg.svg" alt="分享" class="verify-image" />
          <div class="floating-elements">
            <div class="float-item float-1">
              <svg
                width="32"
                height="32"
                viewBox="0 0 24 24"
                fill="none"
                stroke="#DB2777"
                stroke-width="1.5"
                opacity="0.6"
              >
                <path
                  d="M14 2H6a2 2 0 00-2 2v16a2 2 0 002 2h12a2 2 0 002-2V8z"
                />
                <polyline points="14 2 14 8 20 8" />
              </svg>
            </div>
            <div class="float-item float-2">
              <svg
                width="32"
                height="32"
                viewBox="0 0 24 24"
                fill="none"
                stroke="#A855F7"
                stroke-width="1.5"
                opacity="0.6"
              >
                <path
                  d="M22 19a2 2 0 01-2 2H4a2 2 0 01-2-2V5a2 2 0 012-2h5l2 3h9a2 2 0 012 2z"
                />
              </svg>
            </div>
            <div class="float-item float-3">
              <svg
                width="32"
                height="32"
                viewBox="0 0 24 24"
                fill="none"
                stroke="#F472B6"
                stroke-width="1.5"
                opacity="0.6"
              >
                <path d="M18 10h-1.26A8 8 0 109 20h9a5 5 0 000-10z" />
              </svg>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 文件列表 -->
    <div v-if="shareDetail && !needVerify" class="file-container">
      <div class="file-content">
        <!-- 分享信息栏 -->
        <div class="share-info-bar">
          <div class="share-user">
            <el-avatar
              :size="40"
              :src="shareInfo?.shareAccountDTO?.avatarUrl"
              class="user-avatar"
            />
            <div class="user-info">
              <span class="user-name">{{
                shareInfo?.shareAccountDTO?.username
              }}</span>
              <DSTag variant="info" size="small">
                {{
                  formatShareEndTime(
                    shareInfo?.shareEndTime,
                    shareInfo?.shareDayType,
                  )
                }}
              </DSTag>
            </div>
          </div>
          <div class="share-actions">
            <DSButton variant="primary" @click="handleSaveClick">
              <el-icon><FolderAdd /></el-icon>
              保存{{
                selectedFiles.length ? `(${selectedFiles.length}个文件)` : ""
              }}
            </DSButton>
          </div>
        </div>

        <!-- 面包屑导航 -->
        <div class="breadcrumb">
          <el-breadcrumb separator="/">
            <el-breadcrumb-item
              class="breadcrumb-item"
              @click="handlePathChange('/')"
            >
              全部文件
            </el-breadcrumb-item>
            <template v-for="(path, index) in currentPathArray" :key="index">
              <el-breadcrumb-item
                class="breadcrumb-item"
                @click="handlePathChange(getFullPath(index))"
              >
                {{ path }}
              </el-breadcrumb-item>
            </template>
          </el-breadcrumb>
        </div>

        <!-- 文件列表 -->
        <div class="file-list">
          <FileTable
            :fileList="currentFileList"
            @fileClick="handleFileClick"
            @selection-change="handleSelectionChange"
          />
          <div class="pagination-container">
            <div class="pagination-wrapper">
              <DSTag variant="info" class="total-info">
                {{ t("common.total", { count: pagination.total }) }}
              </DSTag>
              <el-pagination
                v-model:current-page="pagination.current"
                v-model:page-size="pagination.pageSize"
                :total="pagination.total"
                :page-sizes="[10, 20, 50, 100]"
                layout="sizes, prev, pager, next, jumper"
                @size-change="handleSizeChange"
                @current-change="handleCurrentChange"
              />
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 底部信息 -->
    <div class="footer-info">
      <div class="footer-links">
        <router-link to="/terms" class="footer-link">服务协议</router-link>
        <span class="divider">|</span>
        <router-link to="/privacy" class="footer-link">隐私政策</router-link>
        <span class="divider">|</span>
        <a href="#" target="_blank" class="footer-link">帮助中心</a>
        <span class="divider">|</span>
        <a href="#" target="_blank" class="footer-link">问题反馈</a>
      </div>
      <div class="copyright">©2024 小七云盘 版权所有</div>
    </div>

    <!-- 保存对话框 -->
    <el-dialog
      v-model="saveDialogVisible"
      title="保存到我的网盘"
      :width="isMobile ? '90%' : '500px'"
      :fullscreen="isMobile"
      class="ds-dialog"
    >
      <div class="save-dialog-content">
        <div class="save-file-info">
          <template v-if="selectedFiles.length === 1">
            <p class="file-name">{{ selectedFiles[0]?.fileName }}</p>
          </template>
          <template v-else>
            <p class="file-count">已选择 {{ selectedFiles.length }} 个文件</p>
          </template>
          <p class="path-label">选择要保存到的文件夹：</p>
        </div>
        <div class="folder-tree-container">
          <FolderTree @select="handleFolderSelect" />
        </div>
        <div class="selected-path">
          <span class="path-label">当前选择：</span>
          <DSTag variant="info" class="path-tag">{{ saveForm.filePath }}</DSTag>
        </div>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <DSButton variant="secondary" @click="saveDialogVisible = false">
            取消
          </DSButton>
          <DSButton
            variant="primary"
            @click="delayedHandleSave"
            :loading="isSaving"
          >
            确定
          </DSButton>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from "vue-i18n";
import {
  checkShareCode,
  getShareDetail,
  getShareFileList,
  saveFiles,
  visitShare,
} from "@/api/share";
import FileTable from "@/components/file/FileTable.vue";
import FolderTree from "@/components/file/FolderTree.vue";
import { FolderAdd, Lock, Unlock } from "@element-plus/icons-vue";
import { ElMessage, ElMessageBox } from "element-plus";
import { computed, onMounted, ref, watch } from "vue";
import { useRoute } from "vue-router";
import { useLoginUserStore } from "@/store/user";
import { DSCard, DSInput, DSButton, DSTag } from "@/components/design-system";

const { t } = useI18n();
const { loginUser, token } = useLoginUserStore();
const accountId = loginUser?.id || "";

const route = useRoute();
const shareId = route.params.shareId as string;

// 移动端检测
const isMobile = computed(() => {
  return window.innerWidth <= 768;
});

const shareInfo = ref<API.ShareDetailDTO | API.ShareSimpleDTO>();
const shareDetail = ref<API.ShareDetailDTO>();
const needVerify = ref(false);
const verifying = ref(false);
const verifyForm = ref({
  shareCode: "",
});

// 格式化日期
const formatDate = (dateStr: string | undefined) => {
  if (!dateStr) return "-";
  return new Date(dateStr).toLocaleString();
};

// 获取和设置 Share-Token 的工具函数
const getShareToken = (shareId: string) => {
  return localStorage.getItem(`share_token_${shareId}`);
};

const setShareToken = (shareId: string, token: string) => {
  localStorage.setItem(`share_token_${shareId}`, token);
};

// 访问分享链接
const visitShareLink = async () => {
  if (!shareId) {
    ElMessage.error("分享ID无效");
    return;
  }

  try {
    // 先检查本地是否有 token
    const storedToken = getShareToken(shareId);
    if (storedToken) {
      const detailResponse = await getShareDetail({
        headers: {
          "Share-Token": storedToken,
        },
      });

      if (detailResponse.data?.code === 0 && detailResponse.data.data) {
        shareDetail.value = detailResponse.data.data as API.ShareDetailDTO;
        shareInfo.value = detailResponse.data.data as API.ShareDetailDTO;
        needVerify.value = false;
        return;
      }
    }

    const response = await visitShare({
      shareId: shareId,
    });

    if (response.data?.code === 0 && response.data.data) {
      const data = response.data.data as API.ShareDetailDTO;
      shareInfo.value = data;
      setShareToken(shareId, response.data.data.shareToken);

      // 获取详细的分享信息（包含文件列表）
      const resp = await getShareDetail({
        headers: {
          "Share-Token": getShareToken(shareId),
        },
      });

      if (resp.data?.code === 0 && resp.data.data) {
        shareDetail.value = resp.data.data as API.ShareDetailDTO;
        shareInfo.value = resp.data.data as API.ShareDetailDTO;
      }

      // 判断是否需要验证
      if (data.shareType === "no_code") {
        needVerify.value = false;
      } else {
        needVerify.value = true;
      }
    } else {
      ElMessage.error(response.data?.message || "访问分享失败");
    }
  } catch (error) {
    console.error("访问分享出错:", error);
    ElMessage.error("访问分享失败");
  }
};

// 验证分享码
const handleVerify = async () => {
  if (!verifyForm.value.shareCode) {
    ElMessage.warning("请输入分享码");
    return;
  }

  verifying.value = true;
  try {
    const response = await checkShareCode({
      shareId: shareId,
      shareCode: verifyForm.value.shareCode,
    });

    if (response.data?.code === 0 && response.data.data) {
      const shareToken = response.data.data.toString();
      setShareToken(shareId, shareToken);

      const detailResponse = await getShareDetail({
        headers: {
          "Share-Token": shareToken,
        },
      });

      if (detailResponse.data?.code === 0 && detailResponse.data.data) {
        shareDetail.value = detailResponse.data.data as API.ShareDetailDTO;
        shareInfo.value = detailResponse.data.data as API.ShareDetailDTO;
        needVerify.value = false;
      } else {
        ElMessage.error(detailResponse.data?.message || "获取分享详情失败");
      }
    } else {
      ElMessage.error(response.data?.message || "验证码错误");
    }
  } catch (error) {
    ElMessage.error("验证失败");
  } finally {
    verifying.value = false;
  }
};

// 路径相关的响应式变量
const currentPath = ref("/");
const currentFileList = ref<API.FileDTO[]>([]);

// 计算当前路径数组
const currentPathArray = computed(() => {
  return currentPath.value.split("/").filter(Boolean);
});

// 获取完整路径
const getFullPath = (index: number) => {
  return "/" + currentPathArray.value.slice(0, index + 1).join("/") + "/";
};

// 分页相关的响应式变量
const pagination = ref({
  current: 1,
  pageSize: 10,
  total: 0,
});

// 更新文件列表
const updateFileList = async () => {
  if (!shareDetail.value?.fileDTOList) return;

  if (currentPath.value === "/") {
    currentFileList.value = shareDetail.value.fileDTOList.map((file) => ({
      ...file,
      fileType: file.isDir === 1 ? "folder" : file.fileType || "other",
      fileSuffix: file.isDir === 1 ? "folder" : file.fileSuffix || "other",
    }));
    pagination.value.total = shareDetail.value.fileDTOList.length;
    return;
  }

  const queryParams: API.ShareFileQueryRequest = {
    shareId: shareId,
    parentId: String(localStorage.getItem("shareid")),
  };

  try {
    const storedToken = getShareToken(shareId);
    let response;

    if (storedToken) {
      response = await getShareFileList(queryParams, {
        headers: {
          "Share-Token": storedToken,
        },
      });
    } else {
      response = await getShareFileList(queryParams, {
        headers: {
          "Share-Token": getShareToken(shareId),
        },
      });
    }

    if (response.data?.code === 0 && response.data.data) {
      currentFileList.value = response.data.data.map((file: API.FileDTO) => ({
        ...file,
        fileSuffix: file.isDir === 1 ? "folder" : file.fileSuffix || "other",
        fileType: file.isDir === 1 ? "folder" : file.fileType || "other",
      }));
      pagination.value.total = Number(response.data.data.total) || 0;
    } else {
      ElMessage.error(response.data?.message || "获取文件列表失败");
    }
  } catch (error) {
    console.error("获取文件列表失败:", error);
    ElMessage.error("获取文件列表失败");
  }
};

// 处理文件点击
const handleFileClick = async (file: API.FileDTO) => {
  if (file.isDir === 1) {
    currentPath.value = `${file.fileName}/`;
    pagination.value.current = 1;
    localStorage.setItem("shareid", String(file.id));
    localStorage.setItem("parentId", String(file.parentId));
    await updateFileList();
  } else {
    console.log("File clicked:", file);
  }
};

// 处理路径变化
const handlePathChange = async (path: string) => {
  currentPath.value = path;
  pagination.value.current = 1;
  await updateFileList();
};

// 监听分享详情变化
watch(
  () => shareDetail.value,
  (newVal) => {
    if (newVal) {
      currentPath.value = "/";
      updateFileList();
    }
  },
  { immediate: false },
);

// 分页处理函数
const handleSizeChange = async (size: number) => {
  pagination.value.pageSize = size;
  pagination.value.current = 1;
  await updateFileList();
};

const handleCurrentChange = async (page: number) => {
  pagination.value.current = page;
  await updateFileList();
};

// 格式化有效期
const formatShareEndTime = (
  endTime: string | undefined,
  dayType: number | undefined,
) => {
  if (!endTime) return "-";
  if (dayType === 0) {
    return "永久有效";
  }
  return new Date(endTime).toLocaleString();
};

// 保存相关的响应式变量
const saveDialogVisible = ref(false);
const saveForm = ref({
  filePath: "/",
  fileIds: [] as any,
});
const selectedFile = ref<API.FileDTO | null>(null);
const selectedFiles = ref<API.FileDTO[]>([]);

// 文件选择变化的处理函数
const handleSelectionChange = (selection: API.FileDTO[]) => {
  selectedFiles.value = selection;
  saveForm.value.fileIds = selection.map((file) => file.id);
};

// 保存按钮的点击处理函数
const handleSaveClick = () => {
  if (!selectedFiles.value.length) {
    ElMessage.warning("请选择要保存的文件");
    return;
  }

  if (!token) {
    ElMessageBox.confirm(t("common.loginRequired"), t("common.info"), {
      confirmButtonText: "去登录",
      cancelButtonText: "取消",
      type: "warning",
    })
      .then(() => {
        window.location.href = `/account/v1/login?redirect=${window.location.href}`;
      })
      .catch(() => {});
    return;
  }

  saveForm.value.fileIds = selectedFiles.value.map((file) => String(file.id));
  saveDialogVisible.value = true;
};

const handleFolderSelect = (node: API.TreeNodeDTO & { path: string }) => {
  saveForm.value.filePath = node.path;
  localStorage.setItem("fileid", String(node.id ?? ""));
};

const handleSave = async () => {
  if (!saveForm.value.filePath) {
    ElMessage.warning("请选择保存位置");
    return;
  }

  try {
    const storedToken = getShareToken(shareId);
    let response;
    const saveParams = {
      shareId: shareId,
      fileIds: saveForm.value.fileIds,
      parentId: String(localStorage.getItem("fileid") || ""),
      accountId: accountId,
    };

    if (storedToken) {
      response = await saveFiles(saveParams, {
        headers: {
          "Share-Token": storedToken,
        },
      });
    } else {
      response = await saveFiles(saveParams, {
        headers: {
          "Share-Token": getShareToken(shareId),
        },
      });
    }

    if (response.data?.code === 0) {
      ElMessage.success("文件转存成功");
      saveDialogVisible.value = false;
      saveForm.value = {
        filePath: "/",
        fileIds: [],
      };
    } else {
      if (loginUser.storageDTO.usedSize > loginUser.storageDTO.totalSize) {
        ElMessage.error("转存失败,存储空间不足");
      } else {
        ElMessage.error(response.data?.message || "转存失败");
      }
    }
  } catch (error) {
    console.error("转存文件失败:", error);
    ElMessage.error("转存失败");
  }
};

const isSaving = ref(false);

// 延迟调用和按钮状态
const delayedHandleSave = () => {
  if (isSaving.value) return;

  isSaving.value = true;

  setTimeout(async () => {
    await handleSave();
    isSaving.value = false;
  }, 500);
};

onMounted(() => {
  visitShareLink();
});
</script>

<style scoped>
.share-view {
  position: relative;
  background: linear-gradient(135deg, #fdf2f8 0%, #f5f3ff 50%, #fce7f3 100%);
}

/* 浮动背景装饰球 */
.bg-decoration {
  position: fixed;
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
    rgba(219, 39, 119, 0.4) 0%,
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
    rgba(168, 85, 247, 0.3) 0%,
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
    rgba(244, 114, 182, 0.3) 0%,
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

/* 验证页面 */
.verify-page {
  display: flex;
  min-height: 100vh;
  position: relative;
  overflow: hidden;
}

.verify-left {
  flex: 0.4;
  padding: 80px 60px;
  display: flex;
  flex-direction: column;
  align-items: center;
  position: relative;
  z-index: 2;
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
}

/* 用户卡片 */
.user-card {
  width: 100%;
  max-width: 400px;
  display: flex;
  align-items: center;
  gap: var(--spacing-lg);
  padding: var(--spacing-xl);
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: var(--radius-xl);
  border: 1px solid rgba(219, 39, 119, 0.1);
  box-shadow: var(--shadow-md);
  animation: fadeInUp 0.6s ease;
  transition: all var(--transition-base);
}

.user-card:hover {
  box-shadow: 0 20px 40px rgba(219, 39, 119, 0.15);
  transform: translateY(-2px);
}

.user-card-avatar {
  position: relative;
  flex-shrink: 0;
}

.user-card-avatar .user-avatar {
  border: 3px solid var(--color-surface);
  box-shadow: var(--shadow-md);
}

.avatar-ring {
  position: absolute;
  inset: -4px;
  border-radius: 50%;
  background: linear-gradient(
    135deg,
    var(--color-primary) 0%,
    var(--color-secondary) 100%
  );
  z-index: -1;
  opacity: 0.5;
}

.user-card-info {
  display: flex;
  flex-direction: column;
  gap: var(--spacing-xs);
}

.user-card-name {
  font-size: var(--text-xl);
  font-weight: var(--font-semibold);
  color: var(--color-text-primary);
  margin: 0;
}

.user-card-share {
  font-size: var(--text-sm);
  color: var(--color-text-secondary);
  margin: 0;
}

.share-count {
  color: var(--color-primary);
  font-weight: var(--font-semibold);
}

/* 验证表单 */
.verify-form {
  width: 100%;
  max-width: 400px;
  display: flex;
  flex-direction: column;
  gap: var(--spacing-lg);
  animation: fadeInUp 0.6s ease 0.2s both;
}

.verify-input-wrapper {
  position: relative;
}

.verify-input :deep(.el-input__wrapper) {
  border-radius: var(--radius-lg);
  box-shadow: 0 0 0 1px var(--color-border) inset;
  padding: 8px 16px;
  transition: all var(--transition-base);
}

.verify-input :deep(.el-input__wrapper:hover) {
  box-shadow: 0 0 0 1px var(--color-primary) inset;
}

.verify-input :deep(.el-input__wrapper.is-focus) {
  box-shadow:
    0 0 0 2px rgba(99, 102, 241, 0.2) inset,
    0 0 0 1px var(--color-primary) inset;
}

.verify-input :deep(.el-input__inner) {
  font-size: var(--text-lg);
  letter-spacing: 4px;
  text-align: center;
}

.verify-input :deep(.el-input__prefix) {
  color: var(--color-text-tertiary);
  font-size: 18px;
}

.verify-button {
  width: 100%;
  height: 52px;
  font-size: var(--text-base);
  font-weight: var(--font-semibold);
  border-radius: var(--radius-lg);
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--spacing-sm);
}

.verify-button .el-icon {
  font-size: 18px;
}

.verify-right {
  position: absolute;
  right: 0;
  top: 0;
  height: 100%;
  width: 60%;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1;
  opacity: 0.9;
}

.illustration-container {
  position: relative;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.verify-image {
  width: 700px;
  max-width: 90%;
  animation: fadeInRight 0.8s ease;
}

/* 浮动装饰元素 */
.floating-elements {
  position: absolute;
  inset: 0;
  pointer-events: none;
}

.float-item {
  position: absolute;
  font-size: 0;
  opacity: 0.6;
  animation: float 3s ease-in-out infinite;
}

.float-1 {
  top: 20%;
  left: 15%;
  animation-delay: 0s;
}

.float-2 {
  top: 35%;
  right: 20%;
  animation-delay: 1s;
}

.float-3 {
  bottom: 25%;
  left: 25%;
  animation-delay: 2s;
}

@keyframes float {
  0%,
  100% {
    transform: translateY(0) rotate(0deg);
  }
  50% {
    transform: translateY(-15px) rotate(5deg);
  }
}

/* 文件列表容器 */
.file-container {
  min-height: 100vh;
  padding-bottom: 100px;
}

.file-content {
  max-width: 1400px;
  margin: 0 auto;
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  min-height: 100vh;
  box-shadow: 0 4px 24px rgba(0, 0, 0, 0.06);
  position: relative;
  z-index: 1;
}

.share-info-bar {
  padding: 24px 32px;
  border-bottom: 1px solid #e2e8f0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: linear-gradient(135deg, #d4a853 0%, #a855f7 100%);
  color: white;
}

.share-info-bar .user-name {
  color: white;
  font-size: 16px;
}

.share-user {
  display: flex;
  align-items: center;
  gap: var(--spacing-md);
}

.user-info {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
}

.share-actions {
  display: flex;
  gap: var(--spacing-md);
}

/* 面包屑 */
.breadcrumb {
  padding: var(--spacing-md) var(--spacing-xl);
  border-bottom: 1px solid var(--border-color);
  background: var(--bg-color);
}

.breadcrumb-item {
  cursor: pointer;
  color: var(--text-secondary);
  transition: color 0.2s;
}

.breadcrumb-item:hover {
  color: var(--color-primary);
}

/* 文件列表 */
.file-list {
  padding: var(--spacing-md) var(--spacing-xl);
  display: flex;
  flex-direction: column;
  flex: 1;
}

.pagination-container {
  padding: var(--spacing-lg) 0;
  border-top: 1px solid var(--border-color);
  background: var(--bg-color);
}

.pagination-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: var(--spacing-lg);
  flex-wrap: wrap;
}

.total-info {
  flex-shrink: 0;
}

/* 底部信息 */
.footer-info {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  padding: var(--spacing-md) 0;
  text-align: center;
  font-size: 13px;
  color: var(--text-secondary);
  background: linear-gradient(
    to bottom,
    transparent,
    rgba(255, 255, 255, 0.95) 20%
  );
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  z-index: 100;
}

.footer-links {
  margin-bottom: var(--spacing-sm);
}

.footer-link {
  color: var(--text-secondary);
  text-decoration: none;
  transition: color 0.3s;
}

.footer-link:hover {
  color: var(--color-primary);
}

.divider {
  margin: 0 var(--spacing-sm);
  color: var(--border-color);
}

.copyright {
  color: var(--text-disabled);
}

/* 对话框 */
.save-dialog-content {
  padding: var(--spacing-md) 0;
}

.save-file-info {
  margin-bottom: var(--spacing-lg);
}

.file-name,
.file-count,
.path-label {
  margin: 0;
  color: var(--text-primary);
  font-size: 14px;
  margin-bottom: var(--spacing-sm);
}

.file-name {
  font-weight: 500;
}

.folder-tree-container {
  border: 1px solid rgba(219, 39, 119, 0.1);
  border-radius: var(--radius-md);
  padding: var(--spacing-md);
  margin-bottom: var(--spacing-lg);
  max-height: 350px;
  overflow-y: auto;
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
}

.selected-path {
  display: flex;
  align-items: center;
  gap: var(--spacing-sm);
  color: var(--text-primary);
  font-size: 14px;
}

.path-tag {
  flex: 1;
}

.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: var(--spacing-md);
}

/* 对话框样式优化 */
:deep(.ds-dialog .el-dialog__header) {
  background: linear-gradient(
    135deg,
    var(--color-secondary) 0%,
    var(--color-primary) 100%
  );
  color: white;
  padding: var(--spacing-lg);
}

:deep(.ds-dialog .el-dialog__title) {
  color: white;
  font-weight: 600;
}

:deep(.ds-dialog .el-dialog__body) {
  padding: var(--spacing-xl);
}

:deep(.ds-dialog .el-dialog__footer) {
  padding: var(--spacing-lg);
  border-top: 1px solid var(--border-color);
}

/* 动画 */
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

@keyframes fadeInRight {
  from {
    opacity: 0;
    transform: translateX(30px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

/* 响应式 */
@media (max-width: 1200px) {
  .verify-right {
    width: 70%;
  }

  .verify-image {
    width: 600px;
  }
}

@media (max-width: 900px) {
  .verify-right {
    width: 80%;
    opacity: 0.7;
  }

  .verify-image {
    width: 500px;
  }
}

@media (max-width: 768px) {
  .verify-left {
    flex: 1;
    padding: var(--spacing-xl) var(--spacing-md);
  }

  .share-info-card,
  .verify-form {
    max-width: 100%;
  }

  .verify-right {
    display: none;
  }

  .file-content {
    box-shadow: none;
  }

  .share-info-bar {
    padding: var(--spacing-md);
  }

  .share-info-bar .user-name {
    font-size: 14px;
  }

  .breadcrumb,
  .file-list {
    padding-left: var(--spacing-md);
    padding-right: var(--spacing-md);
  }

  .pagination-wrapper {
    flex-direction: column;
    gap: var(--spacing-md);
    padding: 0 var(--spacing-md);
  }

  .total-info {
    font-size: 14px;
    width: 100%;
    text-align: center;
  }

  :deep(.el-pagination) {
    width: 100%;
  }

  :deep(.el-pagination__sizes),
  :deep(.el-pagination__jump) {
    display: none;
  }

  .footer-info {
    font-size: 12px;
    padding: var(--spacing-sm) 0;
  }

  .footer-links {
    display: none;
  }
}
</style>
