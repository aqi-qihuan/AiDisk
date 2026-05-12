<template>
  <div class="document-container">
    <!-- 顶部生成内容容器 -->
    <div class="content-container">
      <template v-if="isGenerating">
        <div class="loading-container">
          <div class="loading-animation">
            <div class="loading-ring"></div>
            <div class="loading-icon">
              <el-icon :size="32"><EditPen /></el-icon>
            </div>
          </div>
          <span class="loading-text">{{ t("ai.generatingDoc") }}</span>
          <span class="loading-subtext">{{ t("ai.pleaseWait") }}</span>
        </div>
      </template>
      <div
        v-else-if="generatedContent"
        class="generated-content"
        v-html="generatedContent"
      ></div>
      <div v-else class="generated-content">
        <div class="welcome-section">
          <div class="welcome-card">
            <div class="welcome-header">
              <div class="welcome-avatar">
                <div class="avatar-ring"></div>
                <div class="avatar-inner">
                  <el-icon :size="48"><Document /></el-icon>
                </div>
              </div>
              <h3 class="welcome-title">{{ t("ai.docAssistantTitle") }}</h3>
              <p class="welcome-subtitle">{{ t("ai.docAssistantSubtitle") }}</p>
            </div>

            <div class="welcome-divider"></div>

            <div class="feature-grid">
              <div class="feature-item">
                <div class="feature-icon">
                  <el-icon><Stopwatch /></el-icon>
                </div>
                <div class="feature-text">
                  <span class="feature-title">{{ t("ai.quickGenerate") }}</span>
                  <span class="feature-desc">{{
                    t("ai.docQuickGenerateDesc")
                  }}</span>
                </div>
              </div>
              <div class="feature-item">
                <div class="feature-icon">
                  <el-icon><Aim /></el-icon>
                </div>
                <div class="feature-text">
                  <span class="feature-title">{{
                    t("ai.preciseExtraction")
                  }}</span>
                  <span class="feature-desc">{{
                    t("ai.coreContentRecognition")
                  }}</span>
                </div>
              </div>
              <div class="feature-item">
                <div class="feature-icon">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <circle cx="12" cy="12" r="10"></circle>
                    <line x1="2" y1="12" x2="22" y2="12"></line>
                    <path
                      d="M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10z"
                    ></path>
                  </svg>
                </div>
                <div class="feature-text">
                  <span class="feature-title">{{ t("ai.multiLanguage") }}</span>
                  <span class="feature-desc">{{
                    t("ai.multiLanguageDesc")
                  }}</span>
                </div>
              </div>
              <div class="feature-item">
                <div class="feature-icon">
                  <el-icon><MagicStick /></el-icon>
                </div>
                <div class="feature-text">
                  <span class="feature-title">{{ t("ai.customization") }}</span>
                  <span class="feature-desc">{{
                    t("ai.flexibleConfigDesc")
                  }}</span>
                </div>
              </div>
            </div>

            <div class="welcome-footer">
              <div class="typing-indicator">
                <span class="dot"></span>
                <span class="dot"></span>
                <span class="dot"></span>
              </div>
              <span class="footer-text">{{ t("ai.selectDocToStart") }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 底部输入区域 -->
    <div class="input-section">
      <!-- 文档选择区域 -->
      <div class="doc-select-area">
        <label class="section-label">
          <el-icon class="label-icon" :size="18"><Paperclip /></el-icon>
          <span class="label-text">{{ t("ai.selectDocument") }}</span>
        </label>
        <div class="url-input-container">
          <div class="input-wrapper">
            <input
              type="text"
              v-model="formData.url"
              placeholder="点击右侧按钮选择文档..."
              class="ds-input"
              readonly
            />
            <span class="input-icon" v-if="formData.url">
              <el-icon :size="14"><Check /></el-icon>
            </span>
          </div>
          <button class="select-file-btn" @click="openFileDialog">
            <el-icon class="btn-icon" :size="18"><FolderOpened /></el-icon>
            <span class="btn-text">{{ t("ai.selectDocument") }}</span>
          </button>
        </div>
      </div>

      <!-- 参数配置区域 -->
      <div class="params-section">
        <label class="section-label">
          <el-icon class="label-icon" :size="18"><Setting /></el-icon>
          <span class="label-text">生成参数</span>
        </label>
        <div class="input-row">
          <div class="input-item">
            <label class="input-label">
              <span class="label-dot" style="background: #667eea"></span>
              概要长度
            </label>
            <input
              type="text"
              v-model="formData.length"
              placeholder="例如：200字"
              class="ds-input"
            />
          </div>
          <div class="input-item">
            <label class="input-label">
              <span class="label-dot" style="background: #764ba2"></span>
              输出语言
            </label>
            <input
              type="text"
              v-model="formData.language"
              placeholder="例如：中文"
              class="ds-input"
            />
          </div>
          <div class="input-item">
            <label class="input-label">
              <span class="label-dot" style="background: #f093fb"></span>
              生成风格
            </label>
            <input
              type="text"
              v-model="formData.summary_type"
              placeholder="例如：简洁"
              class="ds-input"
            />
          </div>
          <div class="input-item">
            <label class="input-label">
              <span class="label-dot" style="background: #f5576c"></span>
              附加要求
            </label>
            <input
              type="text"
              v-model="formData.additional_instructions"
              placeholder="其他自定义要求"
              class="ds-input"
            />
          </div>
        </div>
      </div>

      <!-- 生成按钮 -->
      <div class="button-area">
        <button
          class="generate-btn"
          :class="{ loading: isGenerating, disabled: !formData.url }"
          @click="startGeneration"
          :disabled="!formData.url || isGenerating"
        >
          <template v-if="!isGenerating">
            <el-icon class="btn-icon" :size="20"><MagicStick /></el-icon>
            <span class="btn-text">开始生成</span>
          </template>
          <template v-else>
            <div class="btn-spinner"></div>
            <span class="btn-text">生成中...</span>
          </template>
        </button>
      </div>
    </div>

    <!-- 文件选择弹窗 -->
    <el-dialog
      v-model="fileDialogVisible"
      :title="t('ai.selectDocument')"
      width="600px"
      custom-class="ds-dialog"
    >
      <div class="file-browser">
        <!-- 面包屑导航 -->
        <div class="breadcrumb" v-if="currentPath.length > 0">
          <DSTag
            color="primary"
            class="breadcrumb-home"
            @click="navigateTo(-1)"
          >
            <el-icon :size="16"><HomeFilled /></el-icon>
            {{ t("ai.rootDirectory") }}
          </DSTag>
          <span
            v-for="(item, index) in currentPath"
            :key="item.id"
            class="breadcrumb-item"
          >
            <span class="separator">/</span>
            <DSTag
              color="info"
              class="breadcrumb-link"
              @click="navigateTo(index)"
            >
              {{ item.label }}
            </DSTag>
          </span>
        </div>

        <!-- 文件列表 -->
        <div class="file-list">
          <div
            v-for="item in currentFiles"
            :key="item.id"
            class="file-list-item"
            :class="{
              'is-folder': item.isDir === 1,
              'is-selected': selectedFile?.id === item.id,
            }"
            @click="handleItemClick(item)"
          >
            <el-icon
              ><Document v-if="item.isDir !== 1" /><Folder v-else
            /></el-icon>
            <span class="file-name">{{ item.fileName }}</span>
            <span class="file-info" v-if="item.isDir !== 1">
              {{ formatFileSize(item.fileSize) }} |
              {{ formatDate(item.updateTime) }}
            </span>
          </div>
        </div>

        <!-- 空状态 -->
        <div v-if="currentFiles.length === 0" class="empty-folder">
          <div class="empty-icon">
            <el-icon :size="48"><FolderRemove /></el-icon>
          </div>
          <p class="empty-text">{{ t("ai.emptyFolder") }}</p>
        </div>
      </div>

      <template #footer>
        <div class="dialog-footer">
          <DSButton variant="outline" @click="fileDialogVisible = false">
            取消
          </DSButton>
          <DSButton
            variant="primary"
            @click="confirmSelection"
            :disabled="!selectedFile || selectedFile.isDir === 1"
          >
            确定
          </DSButton>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import { useI18n } from "vue-i18n";
import { ElMessage } from "element-plus";
import { listFiles, downloadUrlParam } from "@/api/file";
import { useLoginUserStore } from "@/store/user";
import {
  Document,
  Folder,
  EditPen,
  Stopwatch,
  Aim,
  MagicStick,
  Paperclip,
  FolderOpened,
  Setting,
  HomeFilled,
  FolderRemove,
  Check,
} from "@element-plus/icons-vue";
import LoadingDots from "@/components/common/LoadingDots.vue";
import { getApiUrl, API_PATHS } from "@/config/api";
import DSButton from "@/components/design-system/DSButton.vue";
import DSTag from "@/components/design-system/DSTag.vue";

const { t } = useI18n();
const { loginUser } = useLoginUserStore();
const generatedContent = ref("");
const isGenerating = ref(false);
const fileDialogVisible = ref(false);
const selectedFile = ref<API.FileDTO | null>(null);
const currentPath = ref<API.TreeNodeDTO[]>([]);
const currentFiles = ref<API.FileDTO[]>([]);

const formData = reactive({
  url: "",
  length: "",
  language: "",
  summary_type: "",
  additional_instructions: "",
});

// 处理文件或文件夹点击
const handleItemClick = async (item: API.FileDTO) => {
  if (item.isDir === 1) {
    try {
      const response = await listFiles({
        parent_id: item.id,
      });
      if (response.data?.success && response.data.data) {
        currentFiles.value = response.data.data;
        currentPath.value.push({
          id: item.id,
          label: item.fileName,
        } as API.TreeNodeDTO);
      }
    } catch (error) {
      console.error("获取文件列表失败:", error);
      ElMessage.error("获取文件列表失败");
    }
  } else {
    selectedFile.value = item;
  }
};

// 导航到指定层级
const navigateTo = async (index: number) => {
  if (index === -1) {
    // 返回根目录
    currentPath.value = [];
    try {
      const response = await listFiles({
        parent_id: 0,
      });
      if (response.data?.success && response.data.data) {
        currentFiles.value = response.data.data;
      }
    } catch (error) {
      console.error("获取文件列表失败:", error);
      ElMessage.error("获取文件列表失败");
    }
    return;
  }

  if (index >= currentPath.value.length - 1) return;

  const targetFolder = currentPath.value[index];
  currentPath.value = currentPath.value.slice(0, index + 1);

  try {
    const response = await listFiles({
      parent_id: targetFolder.id,
    });
    if (response.data?.success && response.data.data) {
      currentFiles.value = response.data.data;
    }
  } catch (error) {
    console.error("获取文件列表失败:", error);
    ElMessage.error("获取文件列表失败");
  }
};

// 打开文件选择弹窗
const openFileDialog = async () => {
  fileDialogVisible.value = true;
  currentPath.value = [];
  selectedFile.value = null;

  try {
    const response = await listFiles({
      parent_id: 0,
    });
    if (response.data?.success && response.data.data) {
      currentFiles.value = response.data.data;
    }
  } catch (error) {
    console.error("获取文件列表失败:", error);
    ElMessage.error("获取文件列表失败");
  }
};

// 确认选择
const confirmSelection = async () => {
  if (!selectedFile.value || selectedFile.value.isDir === 1) {
    ElMessage.warning("请选择一个文件");
    return;
  }

  try {
    const response = await downloadUrlParam({
      fileIds: [String(selectedFile.value.id)],
    });

    if (response.data?.success && response.data.data) {
      const baseUrl = formData.url.trim();
      const downloadUrl = response.data.data[0].downloadUrl;
      formData.url = baseUrl ? `${baseUrl}${downloadUrl}` : downloadUrl;
      fileDialogVisible.value = false;
      ElMessage.success("文件选择成功");
    } else {
      ElMessage.error("获取文件下载地址失败");
    }
  } catch (error) {
    console.error("获取文件下载地址失败:", error);
    ElMessage.error("获取文件下载地址失败");
  }
};

// 格式化文件大小
const formatFileSize = (size: number) => {
  if (!size) return "0 B";
  const units = ["B", "KB", "MB", "GB", "TB"];
  let index = 0;
  let fileSize = size;

  while (fileSize >= 1024 && index < units.length - 1) {
    fileSize /= 1024;
    index++;
  }

  return `${fileSize.toFixed(2)} ${units[index]}`;
};

// 格式化日期
const formatDate = (timestamp: number) => {
  if (!timestamp) return "-";
  return new Date(timestamp).toLocaleString();
};

// 处理流式响应
const processStreamResponse = async (response: Response) => {
  const reader = response.body?.getReader();
  if (!reader) return;

  generatedContent.value = "";
  const decoder = new TextDecoder();

  try {
    while (true) {
      const { done, value } = await reader.read();
      if (done) break;

      const chunk = decoder.decode(value, { stream: true });
      const lines = chunk.split("\n");

      for (const line of lines) {
        if (line.trim() === "" || line.trim() === "data: [DONE]") continue;

        try {
          const jsonStr = line.replace(/^data: /, "").trim();
          if (!jsonStr) continue;

          const jsonData = JSON.parse(jsonStr);
          if (jsonData.code === 0) {
            isGenerating.value = false;
            generatedContent.value += jsonData.data;
          }
        } catch (e) {
          console.error("Error parsing JSON:", e);
        }
      }
    }

    // 响应完成后清空输入框
    formData.url = "";
    formData.summary_type = "";
    formData.language = "";
    formData.length = "";
    formData.additional_instructions = "";
  } catch (error) {
    console.error("Stream reading error:", error);
    generatedContent.value += "\n[读取数据时发生错误]";
  } finally {
    isGenerating.value = false;
  }
};

// 开始生成
const startGeneration = async () => {
  if (!formData.url) {
    ElMessage.warning("请选择文档");
    return;
  }

  isGenerating.value = true;
  try {
    const response = await fetch(getApiUrl(API_PATHS.DOCUMENT), {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        url: formData.url,
        summary_type: formData.summary_type,
        language: formData.language,
        length: formData.length,
        additional_instructions: formData.additional_instructions,
      }),
    });

    if (!response.ok) {
      throw new Error("Network response was not ok");
    }

    await processStreamResponse(response);
  } catch (error) {
    console.error("Error:", error);
    generatedContent.value = "生成过程中发生错误，请稍后重试";
    isGenerating.value = false;
  }
};
</script>

<style scoped>
.document-container {
  display: flex;
  flex-direction: column;
  height: 100%;
  padding: 24px;
  gap: 20px;
  background: linear-gradient(135deg, #f8fafc 0%, #e2e8f0 100%);
}

/* ========== 内容容器 ========== */
.content-container {
  flex: 1;
  padding: 32px;
  overflow-y: auto;
  background: white;
  border-radius: 24px;
  box-shadow:
    0 4px 20px rgba(0, 0, 0, 0.05),
    0 0 0 1px rgba(255, 255, 255, 0.5) inset;
  position: relative;
}

.content-container::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #667eea 0%, #764ba2 50%, #f093fb 100%);
  border-radius: 24px 24px 0 0;
}

.generated-content {
  white-space: pre-wrap;
  line-height: 1.8;
  color: #334155;
  font-size: 15px;
}

/* ========== 欢迎界面 ========== */
.welcome-section {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100%;
}

.welcome-card {
  text-align: center;
  max-width: 600px;
  width: 100%;
  animation: fadeInUp 0.6s ease-out;
}

.welcome-header {
  margin-bottom: 32px;
}

.welcome-avatar {
  position: relative;
  width: 100px;
  height: 100px;
  margin: 0 auto 24px;
}

.avatar-ring {
  position: absolute;
  inset: 0;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 50%, #f093fb 100%);
  animation: rotate 4s linear infinite;
}

.avatar-ring::before {
  content: "";
  position: absolute;
  inset: 3px;
  border-radius: 50%;
  background: white;
}

.avatar-inner {
  position: absolute;
  inset: 6px;
  border-radius: 50%;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 48px;
  animation: bounce 2s ease-in-out infinite;
}

.welcome-title {
  font-size: 32px;
  font-weight: 700;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 8px;
}

.welcome-subtitle {
  font-size: 16px;
  color: #64748b;
  font-weight: 500;
}

.welcome-divider {
  height: 1px;
  background: linear-gradient(90deg, transparent, #e2e8f0, transparent);
  margin: 32px 0;
}

/* 功能网格 */
.feature-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 16px;
  margin-bottom: 32px;
}

.feature-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 20px;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  border-radius: 16px;
  border: 1px solid rgba(102, 126, 234, 0.1);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.feature-item:hover {
  background: white;
  border-color: rgba(102, 126, 234, 0.2);
  box-shadow: 0 8px 24px rgba(102, 126, 234, 0.1);
  transform: translateY(-2px);
}

.feature-icon {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #db2777 0%, #a855f7 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(219, 39, 119, 0.3);
}

.feature-icon .el-icon {
  color: white;
  font-size: 24px;
}

.feature-text {
  display: flex;
  flex-direction: column;
  gap: 2px;
  text-align: left;
}

.feature-title {
  font-size: 15px;
  font-weight: 600;
  color: #1e293b;
}

.feature-desc {
  font-size: 13px;
  color: #64748b;
}

/* 底部指示器 */
.welcome-footer {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
}

.typing-indicator {
  display: flex;
  gap: 4px;
}

.typing-indicator .dot {
  width: 6px;
  height: 6px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  animation: typing 1.4s ease-in-out infinite;
}

.typing-indicator .dot:nth-child(1) {
  animation-delay: 0s;
}
.typing-indicator .dot:nth-child(2) {
  animation-delay: 0.2s;
}
.typing-indicator .dot:nth-child(3) {
  animation-delay: 0.4s;
}

.footer-text {
  font-size: 14px;
  color: #94a3b8;
}

/* ========== 加载状态 ========== */
.loading-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 300px;
  animation: fadeIn 0.3s ease-out;
}

.loading-animation {
  position: relative;
  width: 80px;
  height: 80px;
  margin-bottom: 24px;
}

.loading-ring {
  position: absolute;
  inset: 0;
  border-radius: 50%;
  border: 3px solid transparent;
  border-top-color: #db2777;
  border-right-color: #a855f7;
  animation: spin 1s linear infinite;
}

.loading-icon {
  position: absolute;
  inset: 12px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32px;
  animation: pulse 2s ease-in-out infinite;
}

.loading-text {
  font-size: 18px;
  font-weight: 600;
  color: #1e293b;
  margin-bottom: 8px;
}

.loading-subtext {
  font-size: 14px;
  color: #64748b;
}

/* ========== 输入区域 ========== */
.input-section {
  padding: 24px;
  background: white;
  border-radius: 20px;
  box-shadow:
    0 4px 20px rgba(0, 0, 0, 0.05),
    0 0 0 1px rgba(255, 255, 255, 0.5) inset;
}

/* 区域标签 */
.section-label {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 16px;
  font-weight: 600;
  color: #1e293b;
  font-size: 15px;
}

.label-icon {
  font-size: 18px;
}

/* 文档选择区域 */
.doc-select-area {
  margin-bottom: 24px;
}

.url-input-container {
  display: flex;
  gap: 12px;
}

.input-wrapper {
  flex: 1;
  position: relative;
}

.input-wrapper .ds-input {
  width: 100%;
  padding: 14px 44px 14px 16px;
  border: 2px solid #e2e8f0;
  border-radius: 12px;
  font-size: 15px;
  background: #f8fafc;
  color: #334155;
  transition: all 0.3s ease;
  outline: none;
}

.input-wrapper .ds-input:focus {
  border-color: #667eea;
  background: white;
  box-shadow: 0 0 0 4px rgba(102, 126, 234, 0.1);
}

.input-icon {
  position: absolute;
  right: 16px;
  top: 50%;
  transform: translateY(-50%);
  width: 24px;
  height: 24px;
  background: #10b981;
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 14px;
}

.select-file-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 14px 24px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 12px;
  font-size: 15px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3);
}

.select-file-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.4);
}

.select-file-btn:active {
  transform: translateY(0);
}

/* 参数配置区域 */
.params-section {
  margin-bottom: 24px;
}

.input-row {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
}

.input-item {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.input-label {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  font-weight: 500;
  color: #64748b;
}

.label-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}

.input-item .ds-input {
  width: 100%;
  padding: 12px 14px;
  border: 2px solid #e2e8f0;
  border-radius: 10px;
  font-size: 14px;
  background: white;
  color: #334155;
  transition: all 0.3s ease;
  outline: none;
}

.input-item .ds-input:focus {
  border-color: #667eea;
  box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
}

.input-item .ds-input::placeholder {
  color: #94a3b8;
}

/* 生成按钮 */
.button-area {
  display: flex;
  justify-content: center;
  padding-top: 8px;
}

.generate-btn {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 16px 48px;
  background: linear-gradient(135deg, #e2e8f0 0%, #f1f5f9 100%);
  color: #94a3b8;
  border: none;
  border-radius: 14px;
  font-size: 16px;
  font-weight: 600;
  cursor: not-allowed;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.generate-btn:not(.disabled) {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  cursor: pointer;
  box-shadow: 0 4px 16px rgba(102, 126, 234, 0.4);
}

.generate-btn:not(.disabled):hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 24px rgba(102, 126, 234, 0.5);
}

.generate-btn.loading {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  cursor: wait;
}

.btn-spinner {
  width: 20px;
  height: 20px;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

/* ========== 文件浏览器 ========== */
.file-browser {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.breadcrumb {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 8px;
  padding: 12px;
  background: #f8fafc;
  border-radius: 10px;
}

.file-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
  max-height: 400px;
  overflow-y: auto;
}

.file-list-item {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  border-radius: 10px;
  cursor: pointer;
  transition: all 0.3s ease;
  border: 1px solid transparent;
}

.file-list-item:hover {
  background: #f8fafc;
  border-color: #e2e8f0;
  transform: translateX(4px);
}

.file-list-item.is-selected {
  background: rgba(219, 39, 119, 0.1);
  border-color: #db2777;
}

/* ========== 动画 ========== */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
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

@keyframes bounce {
  0%,
  100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-8px);
  }
}

@keyframes rotate {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

@keyframes pulse {
  0%,
  100% {
    opacity: 1;
    transform: scale(1);
  }
  50% {
    opacity: 0.8;
    transform: scale(0.95);
  }
}

@keyframes typing {
  0%,
  60%,
  100% {
    transform: translateY(0);
    opacity: 1;
  }
  30% {
    transform: translateY(-6px);
    opacity: 0.5;
  }
}

/* ========== 响应式 ========== */
@media (max-width: 1024px) {
  .input-row {
    grid-template-columns: repeat(2, 1fr);
  }

  .feature-grid {
    grid-template-columns: 1fr;
  }
}

@media (max-width: 768px) {
  .document-container {
    padding: 16px;
  }

  .content-container {
    padding: 20px;
  }

  .welcome-title {
    font-size: 24px;
  }

  .input-section {
    padding: 16px;
  }

  .url-input-container {
    flex-direction: column;
  }

  .input-row {
    grid-template-columns: 1fr;
  }

  .generate-btn {
    width: 100%;
    justify-content: center;
  }
}
</style>
