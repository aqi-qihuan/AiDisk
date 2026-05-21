<template>
  <div class="answer-container">
    <!-- 消息列表区域 -->
    <div class="message-list" ref="messageList">
      <div
        v-for="(message, index) in messages"
        :key="index"
        :class="[
          'message-item',
          message.isUser ? 'user-message' : 'assistant-message',
        ]"
      >
        <!-- 头像 -->
        <div class="avatar">
          <img
            :src="message.isUser ? userAvatar : assistantAvatar"
            alt="avatar"
          />
        </div>

        <!-- 消息内容 -->
        <div class="message-content-wrapper">
          <div class="message-content" v-html="message.content"></div>
        </div>
      </div>

      <!-- 空状态提示 - 精美欢迎界面 -->
      <div v-if="messages.length === 0" class="welcome-message">
        <div class="welcome-card">
          <div class="welcome-header">
            <div class="welcome-avatar">
              <div class="avatar-ring"></div>
              <div class="avatar-inner">
                <el-icon :size="32"><Cpu /></el-icon>
              </div>
            </div>
            <h3 class="welcome-title">{{ t("ai.assistantTitle") }}</h3>
            <p class="welcome-subtitle">{{ t("ai.assistantSubtitle") }}</p>
          </div>

          <div class="welcome-divider"></div>

          <p class="welcome-description">{{ t("ai.welcomeDesc") }}</p>

          <div class="suggestions">
            <div
              v-for="(suggestion, idx) in suggestions"
              :key="idx"
              class="suggestion-card"
              @click="inputMessage = suggestion.text + suggestion.desc"
            >
              <div class="suggestion-icon">
                <el-icon :size="24">
                  <component :is="suggestion.icon" />
                </el-icon>
              </div>
              <div class="suggestion-content">
                <span class="suggestion-text">{{ suggestion.text }}</span>
                <span class="suggestion-desc">{{ suggestion.desc }}</span>
              </div>
              <div class="suggestion-arrow">→</div>
            </div>
          </div>

          <div class="welcome-footer">
            <div class="typing-indicator">
              <span class="dot"></span>
              <span class="dot"></span>
              <span class="dot"></span>
            </div>
            <span class="footer-text">{{ t("ai.alwaysReady") }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- 输入区域 -->
    <div class="input-area">
      <div class="input-wrapper ds-card">
        <textarea
          v-model="inputMessage"
          :placeholder="t('ai.inputPlaceholder')"
          @keyup.enter.exact.prevent="sendMessage"
          @keydown.enter.shift.prevent="inputMessage += '\n'"
          rows="1"
          ref="textareaRef"
        ></textarea>

        <div class="input-actions">
          <div class="input-tips">
            <DSTag color="info" size="sm">Enter 发送</DSTag>
            <DSTag color="info" size="sm">Shift+Enter 换行</DSTag>
          </div>

          <div
            class="send-button"
            @click="sendMessage"
            :class="{ active: inputMessage.trim(), loading: isLoading }"
          >
            <svg
              v-if="!isLoading"
              class="icon"
              viewBox="0 0 1024 1024"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M931.4 481.3L815.9 396l0.3-0.4-604.7-249c-3.8-1.6-7.9-1.9-11.9-1l-0.3-0.7L135 171.3c-12.3 2.7-20.3 14.8-17.6 27.1 0.9 4.3 3.1 8 6.2 10.9l-0.3 0.7 155.6 130.3L492.3 436l-213.4 95.7-140.9 118c-9.7 8.1-11 22.6-2.8 32.3 3.9 4.6 9.3 7.5 15.2 8.1l0.3 0.7 63.9 26.3c4 1.6 8.1 1.9 12.1 1l0.3 0.7 201.4-83 139.3 57.4c4 1.6 8.1 1.9 12.1 0.9l0.3 0.7 63.9-26.3c12.3-2.7 20.3-14.8 17.6-27.1-0.9-4.3-3.1-8-6.2-10.9l0.3-0.7-102.8-86 389.4-174.6c11.3-5.1 16.3-18.3 11.3-29.6-2.3-4.9-6-8.6-10.8-11z"
              ></path>
            </svg>
            <div v-else class="loading-spinner"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { useLoginUserStore } from "@/store/user";
import { defineComponent, ref, onMounted, nextTick, watch } from "vue";
import { useI18n } from "vue-i18n";
import LoadingDots from "@/components/common/LoadingDots.vue";
import DSButton from "@/components/design-system/DSButton.vue";
import DSTag from "@/components/design-system/DSTag.vue";
import { getApiUrl, API_PATHS } from "@/config/api";
import { ElMessage } from "element-plus";
import { downloadUrlParam } from "@/api/file";
import JSONBIG from "json-bigint";
import userAvatar from "@/assets/user-avatar.png";
import assistantAvatar from "@/assets/assistant-avatar.png";
import {
  Cpu,
  FolderOpened,
  Document,
  VideoCamera,
  Headset,
  Picture,
  DataAnalysis,
  Files,
} from "@element-plus/icons-vue";

const jsonParser = JSONBIG({ storeAsString: true });

// 扩展 Window 接口
declare global {
  interface Window {
    handleFileDownload: (
      fileId: string,
      fileName: string,
      isFolder: boolean,
    ) => Promise<void>;
  }
}

interface FileItem {
  id: string;
  file_id: string | null;
  file_name: string;
  file_type: string | null;
  file_suffix: string | null;
  file_size: number | null;
  gmt_create: string;
  gmt_modified: string;
}

export default defineComponent({
  name: "Answer",
  components: {
    LoadingDots,
    DSButton,
    DSTag,
    Cpu,
    FolderOpened,
    Document,
    VideoCamera,
    Headset,
    Picture,
    DataAnalysis,
    Files,
  },
  setup() {
    const { t } = useI18n();
    const messages = ref<Array<{ isUser: boolean; content: string }>>([]);
    const inputMessage = ref("");
    const messageList = ref<HTMLElement | null>(null);
    const textareaRef = ref<HTMLTextAreaElement | null>(null);
    const isLoading = ref(false);

    const suggestions = ref([
      {
        icon: "FolderOpened",
        text: t("ai.storageSpace"),
        desc: t("ai.viewUsage"),
      },
      {
        icon: "Document",
        text: t("ai.recentFiles"),
        desc: t("ai.browseLatest"),
      },
      {
        icon: "DataAnalysis",
        text: t("ai.fileStats"),
        desc: t("ai.typeAnalysis"),
      },
    ]);

    const scrollToBottom = () => {
      setTimeout(() => {
        if (messageList.value) {
          messageList.value.scrollTop = messageList.value.scrollHeight;
        }
      }, 100);
    };

    watch(inputMessage, () => {
      if (textareaRef.value) {
        textareaRef.value.style.height = "auto";
        textareaRef.value.style.height =
          Math.min(textareaRef.value.scrollHeight, 120) + "px";
      }
    });

    // 格式化文件大小
    const formatFileSize = (size: number) => {
      if (!size) return "-";
      const units = ["B", "KB", "MB", "GB", "TB"];
      let index = 0;
      let fileSize = size;

      while (fileSize >= 1024 && index < units.length - 1) {
        fileSize /= 1024;
        index++;
      }

      return `${fileSize.toFixed(2)} ${units[index]}`;
    };

    // 获取文件图标组件
    const getFileIcon = (fileType: string) => {
      switch (fileType?.toUpperCase()) {
        case "IMG":
          return "Picture";
        case "DOC":
        case "DOCX":
          return "Document";
        case "PDF":
          return "Files";
        case "VIDEO":
          return "VideoCamera";
        case "AUDIO":
          return "Headset";
        case "文件夹":
          return "FolderOpened";
        default:
          return "Document";
      }
    };

    const sendMessage = async () => {
      if (!inputMessage.value.trim() || isLoading.value) return;

      const userMessage = inputMessage.value;
      messages.value.push({ isUser: true, content: userMessage });
      messages.value.push({
        isUser: false,
        content: '<div class="loading-dots">思考中...</div>',
      });
      inputMessage.value = "";

      if (textareaRef.value) {
        textareaRef.value.style.height = "auto";
      }

      scrollToBottom();
      isLoading.value = true;

      try {
        const response = await fetch(getApiUrl(API_PATHS.PAN_QUERY), {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            token: useLoginUserStore().token || "",
          },
          body: JSON.stringify({
            account_id: "0",
            query: userMessage,
          }),
        });

        if (!response.ok) {
          throw new Error("请求失败: " + response.status);
        }

        const text = await response.text();
        const responseData = jsonParser.parse(text);

        if (responseData.code === 0) {
          if (responseData.data && responseData.data.type === "storage_info") {
            const storageData = responseData.data.data;

            if (storageData) {
              const storageInfo = {
                used_size: Number(storageData.used_size) || 0,
                total_size: Number(storageData.total_size) || 10485760,
                used_percentage: Number(storageData.used_percentage) || 0,
              };

              const usedSize = formatFileSize(storageInfo.used_size);
              const totalSize = formatFileSize(storageInfo.total_size);
              const percentage = storageInfo.used_percentage.toFixed(2);

              const formattedContent = `
                <div class="storage-info">
                  <div class="storage-header">
                    <span class="storage-icon" data-icon="storage"></span>
                    <span class="storage-title">存储空间使用情况</span>
                    <span class="storage-percentage">${percentage}%</span>
                  </div>
                  <div class="storage-progress">
                    <div class="progress-bar" style="width: ${percentage}%"></div>
                  </div>
                  <div class="storage-details">
                    <div class="storage-item">
                      <span class="storage-label">已使用</span>
                      <span class="storage-value">${usedSize}</span>
                    </div>
                    <div class="storage-item">
                      <span class="storage-label">总容量</span>
                      <span class="storage-value">${totalSize}</span>
                    </div>
                  </div>
                </div>
              `;
              messages.value[messages.value.length - 1].content =
                formattedContent;
            }
          } else if (
            responseData.type === "text" &&
            responseData.data.type === undefined
          ) {
            messages.value[messages.value.length - 1].content =
              responseData.data.content || "";
          } else if (
            responseData.data &&
            responseData.data.type === "file_list"
          ) {
            const fileList = responseData.data.data as FileItem[];
            let formattedContent = '<div class="file-list">\n';

            fileList.forEach((file: FileItem) => {
              const isFolder = file.file_size === null;
              const fileType = file.file_type || (isFolder ? "文件夹" : "文件");
              const fileName = file.file_name;
              const fileId = String(file.id);

              formattedContent += `
                <div class="file-item">
                  <span class="file-icon" data-icon="${getFileIcon(fileType)}"></span>
                  ${
                    isFolder
                      ? `<span class="file-name">${fileName}</span>`
                      : `<span class="file-name file-link" onclick="handleFileDownload('${String(fileId)}', '${fileName}', false)">${fileName}</span>`
                  }
                </div>
              `;
            });

            formattedContent += "</div>";
            messages.value[messages.value.length - 1].content =
              formattedContent;

            window.handleFileDownload = async (
              fileId: string,
              fileName: string,
              isFolder: boolean,
            ) => {
              try {
                if (isFolder) {
                  ElMessage.warning("文件夹不支持下载");
                  return;
                }

                const fileIds = [fileId];
                const response = await downloadUrlParam({
                  fileIds,
                });

                if (response.data && response.data.success) {
                  const downloadUrls = response.data.data;

                  for (const downloadInfo of downloadUrls) {
                    const response = await fetch(downloadInfo.downloadUrl);
                    const blob = await response.blob();
                    const link = document.createElement("a");
                    const url = URL.createObjectURL(blob);
                    link.href = url;
                    link.download = downloadInfo.fileName;
                    document.body.appendChild(link);
                    link.click();
                    URL.revokeObjectURL(url);
                    document.body.removeChild(link);
                  }

                  ElMessage.success("开始下载文件");
                } else {
                  ElMessage.error(response.data?.msg || "获取下载链接失败");
                }
              } catch (error) {
                console.error("下载文件出错:", error);
                ElMessage.error("下载文件失败");
              }
            };
          } else if (
            responseData.data &&
            responseData.data.type === "file_statistics"
          ) {
            const stats = responseData.data.data;
            const fileTypeMap: {
              [key: string]: { icon: string; name: string };
            } = {
              IMG: { icon: "Picture", name: "图片" },
              DOC: { icon: "Document", name: "文档" },
              VIDEO: { icon: "VideoCamera", name: "视频" },
              AUDIO: { icon: "Headset", name: "音频" },
              PDF: { icon: "Files", name: "PDF" },
              other: { icon: "FolderOpened", name: "其他" },
            };

            let typeStatsHtml = "";
            if (stats && stats.file_types) {
              // 处理数组格式的 file_types
              if (Array.isArray(stats.file_types)) {
                stats.file_types.forEach((item: any) => {
                  const suffix = item.suffix?.toUpperCase() || "other";
                  const typeInfo = fileTypeMap[suffix] || {
                    icon: "📁",
                    name: item.suffix || "其他",
                  };
                  const sizeStr = item.size ? formatFileSize(item.size) : "";
                  typeStatsHtml += `
                    <div class="file-type-item">
                      <span class="type-icon" data-icon="${typeInfo.icon}"></span>
                      <div class="type-info">
                        <span class="type-name">${typeInfo.name}</span>
                        <span class="type-count">${item.count} 个文件 ${sizeStr ? "(" + sizeStr + ")" : ""}</span>
                      </div>
                    </div>
                  `;
                });
              } else {
                // 兼容对象格式
                Object.entries(stats.file_types).forEach(([type, count]) => {
                  const typeInfo = fileTypeMap[type] || {
                    icon: "📁",
                    name: type,
                  };
                  typeStatsHtml += `
                    <div class="file-type-item">
                      <span class="type-icon" data-icon="${typeInfo.icon}"></span>
                      <div class="type-info">
                        <span class="type-name">${typeInfo.name}</span>
                        <span class="type-count">${count} 个文件</span>
                      </div>
                    </div>
                  `;
                });
              }
            }

            const formattedContent = `
              <div class="file-statistics">
                <div class="stats-header">
                  <div class="total-files">
                    <span class="total-icon" data-icon="DataAnalysis"></span>
                    <span class="total-number">${stats.total_files || 0}</span>
                    <span class="total-label">总文件数</span>
                  </div>
                </div>
                ${typeStatsHtml ? `<div class="file-types">${typeStatsHtml}</div>` : ""}
              </div>
            `;
            messages.value[messages.value.length - 1].content =
              formattedContent;
          } else {
            messages.value[messages.value.length - 1].content =
              "未知的响应类型";
          }
        } else {
          messages.value[messages.value.length - 1].content =
            responseData.msg || "请求失败";
        }
      } catch (error) {
        console.error("Error:", error);
        messages.value[messages.value.length - 1].content =
          "抱歉，发生了错误：" +
          (error instanceof Error ? error.message : "未知错误");
      } finally {
        isLoading.value = false;
      }

      scrollToBottom();
    };

    return {
      t,
      messages,
      inputMessage,
      messageList,
      textareaRef,
      sendMessage,
      userAvatar,
      assistantAvatar,
      isLoading,
      suggestions,
    };
  },
});
</script>

<style scoped>
.answer-container {
  display: flex;
  flex-direction: column;
  height: 100%;
  background: linear-gradient(160deg, #0a0a1a 0%, #141428 40%, #1a1530 100%);
  position: relative;
  overflow: hidden;
}

/* 消息列表 */
.message-list {
  flex: 1;
  overflow-y: auto;
  padding: var(--ds-spacing-4);
  display: flex;
  flex-direction: column;
  gap: var(--ds-spacing-4);
}

/* 消息项 */
.message-item {
  display: flex;
  gap: var(--ds-spacing-3);
  max-width: 75%;
  animation: fadeInUp 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.message-item.user-message {
  margin-left: auto;
  flex-direction: row-reverse;
}

.message-item.assistant-message {
  margin-right: auto;
}

/* 头像 */
.avatar {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  overflow: hidden;
  flex-shrink: 0;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  border: 2px solid rgba(255, 255, 255, 0.3);
  transition: transform 0.3s ease;
}

.avatar:hover {
  transform: scale(1.05);
}

.avatar img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

/* 消息内容包裹器 */
.message-content-wrapper {
  display: flex;
  flex-direction: column;
  gap: var(--ds-spacing-1);
  max-width: 100%;
}

/* 消息内容 */
.message-content {
  padding: 14px 20px;
  border-radius: 20px;
  font-size: 15px;
  line-height: 1.7;
  word-break: break-word;
  min-height: 44px;
  max-width: 100%;
  position: relative;
  transition: all 0.3s ease;
}

/* 用户消息 - 粉紫色渐变气泡 */
.message-item.user-message .message-content {
  background: linear-gradient(135deg, #c9a96e 0%, #d4a853 100%);
  color: #0a0a1a;
  box-shadow:
    0 8px 24px rgba(212, 168, 83, 0.25),
    0 2px 4px rgba(0, 0, 0, 0.3);
  border-bottom-right-radius: 6px;
  font-weight: 500;
}

.message-item.user-message .message-content:hover {
  box-shadow:
    0 12px 32px rgba(212, 168, 83, 0.35),
    0 4px 8px rgba(0, 0, 0, 0.25);
  transform: translateY(-1px);
}

/* AI 消息 - 暗色毛玻璃气泡 */
.message-item.assistant-message .message-content {
  background: rgba(20, 20, 45, 0.85);
  color: #e8d5b0;
  box-shadow:
    0 4px 16px rgba(0, 0, 0, 0.3),
    0 1px 3px rgba(0, 0, 0, 0.2),
    inset 0 1px 0 rgba(212, 168, 83, 0.1);
  border-bottom-left-radius: 6px;
  backdrop-filter: blur(16px);
  border: 1px solid rgba(212, 168, 83, 0.15);
}

.message-item.assistant-message .message-content:hover {
  box-shadow:
    0 8px 24px rgba(0, 0, 0, 0.4),
    0 2px 6px rgba(0, 0, 0, 0.25),
    inset 0 1px 0 rgba(212, 168, 83, 0.15);
  transform: translateY(-1px);
  border-color: rgba(212, 168, 83, 0.25);
}

/* 欢迎消息 - 全新设计 */
.welcome-message {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100%;
  padding: var(--ds-spacing-6);
  animation: fadeIn 0.6s ease-out;
}

.welcome-card {
  background: rgba(20, 22, 40, 0.8);
  backdrop-filter: blur(24px);
  border-radius: 24px;
  padding: var(--ds-spacing-8);
  max-width: 480px;
  width: 100%;
  border: 1px solid rgba(212, 168, 83, 0.2);
  box-shadow:
    0 20px 60px rgba(0, 0, 0, 0.5),
    0 0 0 1px rgba(212, 168, 83, 0.1) inset,
    0 0 40px rgba(212, 168, 83, 0.05);
  animation: slideUp 0.6s ease-out;
}

.welcome-header {
  text-align: center;
  margin-bottom: var(--ds-spacing-6);
}

.welcome-avatar {
  position: relative;
  width: 80px;
  height: 80px;
  margin: 0 auto var(--ds-spacing-4);
}

.avatar-ring {
  position: absolute;
  inset: 0;
  border-radius: 50%;
  background: linear-gradient(135deg, #d4a853 0%, #c9a96e 50%, #b8943f 100%);
  animation: rotate 3s linear infinite;
}

.avatar-ring::before {
  content: "";
  position: absolute;
  inset: 3px;
  border-radius: 50%;
  background: #0f1125;
}

.avatar-inner {
  position: absolute;
  inset: 6px;
  border-radius: 50%;
  background: linear-gradient(135deg, #c9a96e 0%, #b8943f 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  color: #0a0a1a;
  font-size: 36px;
  animation: bounce 2s ease-in-out infinite;
}

.welcome-title {
  font-size: 28px;
  font-weight: 700;
  background: linear-gradient(135deg, #d4a853 0%, #f0c060 50%, #c9a96e 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: var(--ds-spacing-2);
  text-shadow: 0 2px 10px rgba(212, 168, 83, 0.3);
}

.welcome-subtitle {
  font-size: var(--ds-text-size-base);
  color: #c4b998;
  font-weight: 500;
}

.welcome-divider {
  height: 1px;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(212, 168, 83, 0.25),
    transparent
  );
  margin: var(--ds-spacing-6) 0;
}

.welcome-description {
  font-size: var(--ds-text-size-base);
  color: #8b8878;
  text-align: center;
  margin-bottom: var(--ds-spacing-5);
  font-weight: 500;
}

.suggestions {
  display: flex;
  flex-direction: column;
  gap: var(--ds-spacing-3);
  margin-bottom: var(--ds-spacing-6);
}

.suggestion-card {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-4);
  padding: var(--ds-spacing-4);
  background: rgba(20, 22, 40, 0.5);
  border-radius: 16px;
  border: 1px solid rgba(212, 168, 83, 0.1);
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.suggestion-card:hover {
  background: rgba(30, 32, 55, 0.7);
  border-color: rgba(212, 168, 83, 0.35);
  box-shadow: 0 8px 24px rgba(212, 168, 83, 0.12);
  transform: translateY(-2px);
}

.suggestion-icon {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #c9a96e 0%, #b8943f 100%);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  color: #0a0a1a;
  box-shadow: 0 4px 12px rgba(212, 168, 83, 0.25);
}

.suggestion-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.suggestion-text {
  font-size: var(--ds-text-size-base);
  font-weight: 600;
  color: #e8d5b0;
}

.suggestion-desc {
  font-size: var(--ds-text-size-sm);
  color: #8b8878;
}

.suggestion-arrow {
  width: 32px;
  height: 32px;
  background: rgba(212, 168, 83, 0.1);
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #d4a853;
  font-weight: 600;
  transition: all 0.3s ease;
}

.suggestion-card:hover .suggestion-arrow {
  background: linear-gradient(135deg, #d4a853 0%, #c9a96e 100%);
  color: #0a0a1a;
  transform: translateX(4px);
}

.welcome-footer {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--ds-spacing-3);
  padding-top: var(--ds-spacing-4);
  border-top: 1px solid rgba(212, 168, 83, 0.15);
}

.typing-indicator {
  display: flex;
  gap: 4px;
}

.typing-indicator .dot {
  width: 6px;
  height: 6px;
  background: linear-gradient(135deg, #d4a853 0%, #c9a96e 100%);
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
  font-size: var(--ds-text-size-sm);
  color: #8b8878;
}

/* 输入区域 - HOK暗黑史诗风格 */
.input-area {
  padding: var(--ds-spacing-4) var(--ds-spacing-6);
  background: linear-gradient(
    180deg,
    rgba(10, 10, 30, 0) 0%,
    rgba(20, 22, 40, 0.8) 20%,
    rgba(15, 18, 35, 0.95) 100%
  );
  position: relative;
}

.input-area::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(
    90deg,
    transparent,
    rgba(212, 168, 83, 0.2),
    transparent
  );
}

.input-wrapper {
  position: relative;
  max-width: 900px;
  margin: 0 auto;
  background: rgba(20, 22, 40, 0.7);
  border-radius: 20px;
  border: 1px solid rgba(212, 168, 83, 0.15);
  padding: var(--ds-spacing-4) var(--ds-spacing-16) var(--ds-spacing-4)
    var(--ds-spacing-5);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  display: flex;
  flex-direction: column;
  gap: var(--ds-spacing-3);
  box-shadow:
    0 4px 20px rgba(0, 0, 0, 0.4),
    0 1px 3px rgba(0, 0, 0, 0.3),
    inset 0 1px 0 rgba(212, 168, 83, 0.05);
}

.input-wrapper:hover {
  border-color: rgba(212, 168, 83, 0.25);
  box-shadow:
    0 8px 30px rgba(212, 168, 83, 0.08),
    0 2px 8px rgba(0, 0, 0, 0.35),
    inset 0 1px 0 rgba(212, 168, 83, 0.08);
}

.input-wrapper:focus-within {
  border-color: rgba(212, 168, 83, 0.45);
  box-shadow:
    0 0 0 4px rgba(212, 168, 83, 0.06),
    0 8px 30px rgba(212, 168, 83, 0.1),
    inset 0 1px 0 rgba(212, 168, 83, 0.1);
  transform: translateY(-1px);
}

.input-wrapper textarea {
  width: 100%;
  border: none;
  background: transparent;
  resize: none;
  outline: none;
  font-size: 15px;
  line-height: 1.7;
  max-height: 140px;
  min-height: 24px;
  overflow-y: auto;
  font-family: inherit;
  color: #e8d5b0;
}

.input-wrapper textarea::placeholder {
  color: #6b6b78;
  font-size: 14px;
}

.input-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.input-tips {
  display: flex;
  gap: var(--ds-spacing-2);
}

.input-tips :deep(.ds-tag) {
  background: rgba(212, 168, 83, 0.08);
  border: 1px solid rgba(212, 168, 83, 0.15);
  color: #c9a96e;
  font-size: 11px;
  padding: 3px 10px;
  border-radius: 6px;
  font-weight: 500;
  transition: all 0.2s ease;
}

.input-tips :deep(.ds-tag):hover {
  background: rgba(212, 168, 83, 0.15);
  transform: translateY(-1px);
}

/* 发送按钮 - HOK暗黑史诗风格 */
.send-button {
  position: absolute;
  right: 12px;
  bottom: 12px;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  border-radius: 12px;
  background: rgba(30, 32, 55, 0.8);
  border: 1px solid rgba(212, 168, 83, 0.1);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  flex-shrink: 0;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
}

.send-button svg {
  width: 18px;
  height: 18px;
  fill: #6b6b78;
  transition: all 0.3s ease;
}

.send-button.active {
  background: linear-gradient(135deg, #c9a96e 0%, #d4a853 100%);
  border-color: transparent;
  box-shadow: 0 4px 15px rgba(212, 168, 83, 0.35);
}

.send-button.active svg {
  fill: #0a0a1a;
  transform: translateX(1px);
}

.send-button:hover:not(.loading) {
  transform: scale(1.05);
}

.send-button.active:hover:not(.loading) {
  transform: scale(1.05) translateY(-2px);
  box-shadow: 0 6px 20px rgba(212, 168, 83, 0.45);
}

.send-button:active:not(.loading) {
  transform: scale(0.95);
}

.send-button.loading {
  cursor: wait;
  background: linear-gradient(135deg, #c9a96e 0%, #d4a853 100%);
}

.loading-spinner {
  width: 18px;
  height: 18px;
  border: 2px solid rgba(10, 10, 26, 0.3);
  border-top-color: #0a0a1a;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

/* 存储信息样式 */
.storage-info {
  background: rgba(20, 22, 40, 0.7);
  border: 1px solid rgba(212, 168, 83, 0.15);
  border-radius: var(--ds-radius-lg);
  padding: var(--ds-spacing-4);
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

.storage-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--ds-spacing-3);
}

.storage-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 24px;
  height: 24px;
  font-size: 0;
}

.storage-icon::before {
  content: "";
  width: 24px;
  height: 24px;
  background-size: contain;
  background-repeat: no-repeat;
  background-position: center;
}

.storage-icon[data-icon="storage"]::before {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23C9A96E'%3E%3Cpath d='M5 6.2C5 5.07 5.9 4 7.05 4h9.9C18.1 4 19 5.07 19 6.2c0 .59-.26 1.16-.74 1.54L13.5 11.4c-.88.7-2.12.7-3 0L5.74 7.74C5.26 7.36 5 6.79 5 6.2zM5 10.8v6.4C5 18.67 6.33 20 7.95 20h8.1C17.67 20 19 18.67 19 17.2v-6.4l-3.5 2.8c-1.48 1.18-3.52 1.18-5 0L5 10.8z'/%3E%3C/svg%3E");
}

.storage-title {
  font-size: var(--ds-text-size-base);
  font-weight: 600;
  color: #e8d5b0;
  flex: 1;
  margin-left: var(--ds-spacing-2);
}

.storage-percentage {
  font-size: var(--ds-text-size-lg);
  font-weight: 600;
  color: #d4a853;
}

.storage-progress {
  height: 8px;
  background: rgba(20, 22, 40, 0.5);
  border-radius: 4px;
  overflow: hidden;
  margin-bottom: var(--ds-spacing-3);
}

.progress-bar {
  height: 100%;
  background: linear-gradient(90deg, #c9a96e 0%, #d4a853 100%);
  border-radius: 4px;
  transition: width 0.3s ease;
  box-shadow: 0 0 8px rgba(212, 168, 83, 0.3);
}

.storage-details {
  display: flex;
  justify-content: space-between;
}

.storage-item {
  display: flex;
  flex-direction: column;
  gap: var(--ds-spacing-1);
}

.storage-label {
  font-size: var(--ds-text-size-sm);
  color: #8b8878;
}

.storage-value {
  font-size: var(--ds-text-size-base);
  font-weight: 600;
  color: #e8d5b0;
}

/* 文件列表样式 */
.file-list {
  display: flex;
  flex-direction: column;
  gap: var(--ds-spacing-2);
}

.file-item {
  display: flex;
  align-items: center;
  padding: var(--ds-spacing-2);
  border-radius: var(--ds-radius-md);
  background: rgba(20, 22, 40, 0.4);
  border: 1px solid transparent;
  transition: all 0.3s ease;
}

.file-item:hover {
  background: rgba(30, 32, 55, 0.6);
  border-color: rgba(212, 168, 83, 0.15);
  transform: translateX(4px);
}

.file-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 20px;
  height: 20px;
  margin-right: var(--ds-spacing-2);
  font-size: 0;
}

.file-icon::before {
  content: "";
  width: 20px;
  height: 20px;
  background-size: contain;
  background-repeat: no-repeat;
  background-position: center;
}

/* 文件类型图标 SVG */
.file-icon[data-icon="Picture"]::before {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23D4A853'%3E%3Cpath d='M4 5a2 2 0 012-2h12a2 2 0 012 2v14a2 2 0 01-2 2H6a2 2 0 01-2-2V5zm2 0v14h12V5H6zm2 10l3-3 2 2 4-4v6H8v-1z'/%3E%3C/svg%3E");
}

.file-icon[data-icon="Document"]::before {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23C9A96E'%3E%3Cpath d='M7 2a2 2 0 00-2 2v16a2 2 0 002 2h10a2 2 0 002-2V8l-6-6H7zm1 2h4v4h4v12H8V4zm2 8v2h6v-2h-6zm0 4v2h4v-2h-4z'/%3E%3C/svg%3E");
}

.file-icon[data-icon="Files"]::before {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23B8943F'%3E%3Cpath d='M3 6a2 2 0 012-2h6l2 2h8a2 2 0 012 2v10a2 2 0 01-2 2H5a2 2 0 01-2-2V6z'/%3E%3C/svg%3E");
}

.file-icon[data-icon="VideoCamera"]::before {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23D4A853'%3E%3Cpath d='M4 8a2 2 0 012-2h8a2 2 0 012 2v8a2 2 0 01-2 2H6a2 2 0 01-2-2V8zm14 0l4-2v12l-4-2V8z'/%3E%3C/svg%3E");
}

.file-icon[data-icon="Headset"]::before {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23C9A96E'%3E%3Cpath d='M12 3a7 7 0 00-7 7v4a3 3 0 003 3h1v-6H6v-1a6 6 0 1112 0v1h-3v6h3a3 3 0 003-3v-4a7 7 0 00-7-7z'/%3E%3C/svg%3E");
}

.file-icon[data-icon="FolderOpened"]::before {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23C9A96E'%3E%3Cpath d='M3 7a2 2 0 012-2h4.586a1 1 0 01.707.293l2.414 2.414a1 1 0 00.707.293H19a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2V7z'/%3E%3C/svg%3E");
}

.file-name {
  flex: 1;
  color: #e8d5b0;
}

.file-link {
  color: #d4a853;
  cursor: pointer;
  text-decoration: underline;
}

.file-link:hover {
  color: #f0c060;
}

/* 文件统计样式 */
.file-statistics {
  background: rgba(20, 22, 40, 0.7);
  border: 1px solid rgba(212, 168, 83, 0.15);
  border-radius: var(--ds-radius-lg);
  padding: var(--ds-spacing-4);
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
}

.stats-header {
  text-align: center;
  margin-bottom: var(--ds-spacing-4);
}

.total-files {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--ds-spacing-1);
}

.total-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  font-size: 0;
}

.total-icon::before {
  content: "";
  width: 32px;
  height: 32px;
  background-size: contain;
  background-repeat: no-repeat;
  background-position: center;
}

.total-icon[data-icon="DataAnalysis"]::before {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23D4A853'%3E%3Cpath d='M3 3v18h18v-2H5V3H3zm14 4v10h2V7h-2zm-6 4v6h2v-6h-2zm-4 2v4h2v-4H7z'/%3E%3C/svg%3E");
}

.type-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 24px;
  height: 24px;
  margin-right: var(--ds-spacing-2);
  font-size: 0;
}

.type-icon::before {
  content: "";
  width: 24px;
  height: 24px;
  background-size: contain;
  background-repeat: no-repeat;
  background-position: center;
}

.type-icon[data-icon="Picture"]::before {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23D4A853'%3E%3Cpath d='M4 5a2 2 0 012-2h12a2 2 0 012 2v14a2 2 0 01-2 2H6a2 2 0 01-2-2V5zm2 0v14h12V5H6zm2 10l3-3 2 2 4-4v6H8v-1z'/%3E%3C/svg%3E");
}

.type-icon[data-icon="Document"]::before {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23C9A96E'%3E%3Cpath d='M7 2a2 2 0 00-2 2v16a2 2 0 002 2h10a2 2 0 002-2V8l-6-6H7zm1 2h4v4h4v12H8V4zm2 8v2h6v-2h-6zm0 4v2h4v-2h-4z'/%3E%3C/svg%3E");
}

.type-icon[data-icon="Files"]::before {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23B8943F'%3E%3Cpath d='M3 6a2 2 0 012-2h6l2 2h8a2 2 0 012 2v10a2 2 0 01-2 2H5a2 2 0 01-2-2V6z'/%3E%3C/svg%3E");
}

.type-icon[data-icon="VideoCamera"]::before {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23D4A853'%3E%3Cpath d='M4 8a2 2 0 012-2h8a2 2 0 012 2v8a2 2 0 01-2 2H6a2 2 0 01-2-2V8zm14 0l4-2v12l-4-2V8z'/%3E%3C/svg%3E");
}

.type-icon[data-icon="Headset"]::before {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23C9A96E'%3E%3Cpath d='M12 3a7 7 0 00-7 7v4a3 3 0 003 3h1v-6H6v-1a6 6 0 1112 0v1h-3v6h3a3 3 0 003-3v-4a7 7 0 00-7-7z'/%3E%3C/svg%3E");
}

.type-icon[data-icon="FolderOpened"]::before {
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%23C9A96E'%3E%3Cpath d='M3 7a2 2 0 012-2h4.586a1 1 0 01.707.293l2.414 2.414a1 1 0 00.707.293H19a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2V7z'/%3E%3C/svg%3E");
}

.total-number {
  font-size: 36px;
  font-weight: 600;
  color: #d4a853;
}

.total-label {
  font-size: var(--ds-text-size-sm);
  color: #8b8878;
}

.file-types {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  gap: var(--ds-spacing-2);
}

.file-type-item {
  display: flex;
  align-items: center;
  padding: var(--ds-spacing-2);
  background: rgba(20, 22, 40, 0.4);
  border: 1px solid transparent;
  border-radius: var(--ds-radius-md);
  transition: all 0.3s ease;
}

.file-type-item:hover {
  background: rgba(212, 168, 83, 0.08);
  border-color: rgba(212, 168, 83, 0.15);
  transform: translateY(-2px);
}

.type-icon {
  font-size: 24px;
  margin-right: var(--ds-spacing-2);
}

.type-info {
  display: flex;
  flex-direction: column;
}

.type-name {
  font-size: var(--ds-text-size-sm);
  font-weight: 500;
  color: #e8d5b0;
}

.type-count {
  font-size: var(--ds-text-size-xs);
  color: #8b8878;
}

/* 加载动画 */
.loading-dots {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-1);
  color: #8b8878;
}

/* 动画 */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(10px);
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

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(30px) scale(0.95);
  }
  to {
    opacity: 1;
    transform: translateY(0) scale(1);
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

@keyframes typing {
  0%,
  60%,
  100% {
    transform: translateY(0);
    opacity: 1;
  }
  30% {
    transform: translateY(-8px);
    opacity: 0.5;
  }
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

/* 响应式 */
@media (max-width: 768px) {
  .answer-container {
    height: calc(100vh - 56px);
  }

  .message-list {
    padding: var(--ds-spacing-3);
    gap: var(--ds-spacing-3);
  }

  .message-item {
    max-width: 95%;
    gap: var(--ds-spacing-2);
  }

  .avatar {
    width: 32px;
    height: 32px;
  }

  .message-content {
    padding: var(--ds-spacing-2) var(--ds-spacing-3);
    font-size: var(--ds-text-size-sm);
  }

  .input-area {
    padding: var(--ds-spacing-3) var(--ds-spacing-4);
  }

  .input-wrapper {
    padding: var(--ds-spacing-3) var(--ds-spacing-12) var(--ds-spacing-3)
      var(--ds-spacing-4);
    border-radius: 16px;
  }

  .input-wrapper textarea {
    font-size: 14px;
    min-height: 20px;
  }

  .send-button {
    width: 36px;
    height: 36px;
    right: 10px;
    bottom: 10px;
    border-radius: 10px;
  }

  .send-button svg {
    width: 16px;
    height: 16px;
  }

  .welcome-card {
    padding: var(--ds-spacing-5);
    border-radius: 20px;
  }

  .welcome-avatar {
    width: 64px;
    height: 64px;
  }

  .avatar-inner {
    font-size: 28px;
  }

  .welcome-title {
    font-size: 22px;
  }

  .welcome-subtitle {
    font-size: var(--ds-text-size-sm);
  }

  .welcome-description {
    font-size: var(--ds-text-size-sm);
  }

  .suggestion-card {
    padding: var(--ds-spacing-3);
  }

  .suggestion-icon {
    width: 40px;
    height: 40px;
    font-size: 20px;
  }

  .suggestion-text {
    font-size: var(--ds-text-size-sm);
  }

  .suggestion-desc {
    font-size: var(--ds-text-size-xs);
  }

  .input-tips {
    display: none;
  }

  .file-types {
    grid-template-columns: 1fr;
  }

  .storage-info {
    padding: var(--ds-spacing-3);
  }

  .storage-header {
    margin-bottom: var(--ds-spacing-2);
  }

  .storage-title {
    font-size: var(--ds-text-size-sm);
  }

  .storage-percentage {
    font-size: var(--ds-text-size-base);
  }

  .storage-value {
    font-size: var(--ds-text-size-sm);
  }

  .file-statistics {
    padding: var(--ds-spacing-3);
  }

  .total-number {
    font-size: 28px;
  }
}

@media (max-width: 480px) {
  .answer-container {
    height: calc(100vh - 52px);
  }

  .message-item {
    max-width: 98%;
  }

  .avatar {
    width: 28px;
    height: 28px;
  }

  .message-content {
    padding: var(--ds-spacing-2);
    font-size: var(--ds-text-size-xs);
  }

  .input-area {
    padding: var(--ds-spacing-2) var(--ds-spacing-3);
  }

  .input-wrapper {
    padding: var(--ds-spacing-3) var(--ds-spacing-10) var(--ds-spacing-3)
      var(--ds-spacing-3);
    border-radius: 14px;
  }

  .input-wrapper textarea {
    font-size: 14px;
  }

  .input-tips {
    display: none;
  }

  .send-button {
    width: 34px;
    height: 34px;
    right: 8px;
    bottom: 8px;
    border-radius: 10px;
  }

  .send-button svg {
    width: 16px;
    height: 16px;
  }
}
</style>
