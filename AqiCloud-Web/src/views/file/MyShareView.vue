<template>
  <div class="my-share-view">
    <!-- 浮动背景装饰 -->
    <div class="bg-decoration">
      <div class="bg-orb bg-orb-1"></div>
      <div class="bg-orb bg-orb-2"></div>
      <div class="bg-orb bg-orb-3"></div>
    </div>
    <!-- 页面标题 -->
    <div class="page-header">
      <div class="header-title">
        <el-icon class="title-icon"><Share /></el-icon>
        <h1>{{ t('file.myShares') }}</h1>
      </div>
      <p class="header-desc">{{ t('file.manageShareDesc') }}</p>
    </div>

    <!-- 分享统计卡片 -->
    <div class="stats-row">
      <div class="stat-card">
        <div class="stat-icon public">
          <el-icon><Unlock /></el-icon>
        </div>
        <div class="stat-info">
          <span class="stat-value">{{ publicShareCount }}</span>
          <span class="stat-label">{{ t('file.publicShares') }}</span>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon private">
          <el-icon><Lock /></el-icon>
        </div>
        <div class="stat-info">
          <span class="stat-value">{{ privateShareCount }}</span>
          <span class="stat-label">{{ t('file.privateShares') }}</span>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon total">
          <el-icon><Document /></el-icon>
        </div>
        <div class="stat-info">
          <span class="stat-value">{{ shareList.length }}</span>
          <span class="stat-label">{{ t('file.totalShares') }}</span>
        </div>
      </div>
    </div>

    <!-- 分享列表 -->
    <div class="share-list">
      <div v-if="shareList.length === 0" class="empty-state">
        <div class="empty-icon">
          <el-icon><Share /></el-icon>
        </div>
        <p class="empty-title">{{ t('file.noShares') }}</p>
        <p class="empty-desc">{{ t('file.noSharesDesc') }}</p>
      </div>

      <div v-else class="share-cards">
        <div
          v-for="share in shareList"
          :key="share.id"
          class="share-card"
          :class="{ 'expired': isExpired(share.shareEndTime, share.shareDayType) }"
        >
          <!-- 卡片头部 -->
          <div class="card-header">
            <div class="share-icon">
              <el-icon><Link /></el-icon>
            </div>
            <div class="share-type-badge" :class="share.shareType?.toLowerCase()">
              <el-icon v-if="share.shareType?.toLowerCase() === 'no_code'"><Unlock /></el-icon>
              <el-icon v-else><Lock /></el-icon>
              <span>{{ share.shareType?.toLowerCase() === 'no_code' ? '公开' : '私密' }}</span>
            </div>
          </div>

          <!-- 卡片内容 -->
          <div class="card-body">
            <h3 class="share-name" :title="share.shareName">{{ share.shareName }}</h3>
            
            <div class="share-meta">
              <div class="meta-item">
                <el-icon><Clock /></el-icon>
                <span>{{ formatShareEndTime(share.shareEndTime, share.shareDayType) }}</span>
              </div>
              <div class="meta-item">
                <el-icon><Calendar /></el-icon>
                <span>{{ formatDate(share.gmtCreate || share.createTime) }}</span>
              </div>
            </div>

            <!-- 提取码显示 -->
            <div v-if="share.shareCode" class="share-code">
              <span class="code-label">提取码</span>
              <span class="code-value">{{ share.shareCode }}</span>
            </div>
          </div>

          <!-- 卡片操作 -->
          <div class="card-footer">
            <button
              class="action-btn copy"
              @click="copyShareInfo(share)"
              :title="share.shareType?.toLowerCase() === 'need_code' ? '复制链接和提取码' : '复制链接'"
            >
              <el-icon><DocumentCopy /></el-icon>
              <span>复制</span>
            </button>
            <button class="action-btn delete" @click="handleCancelShare(share)" title="取消分享">
              <el-icon><Delete /></el-icon>
              <span>取消</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import { useI18n } from 'vue-i18n';
import { ElMessage, ElMessageBox } from "element-plus";
import { 
  Share, 
  DocumentCopy, 
  Delete, 
  Link, 
  Clock, 
  Calendar,
  Unlock,
  Lock,
  Document
} from "@element-plus/icons-vue";
import { getShareUrl, cancel } from "@/api/share";
import { useLoginUserStore } from "@/store/user";

const { t } = useI18n();
const { loginUser } = useLoginUserStore();
const accountId = loginUser.id;

const shareList = ref<API.ShareDTO[]>([]);

// 计算属性：公开分享数量
const publicShareCount = computed(() => {
  return shareList.value.filter(s => s.shareType?.toLowerCase() === 'no_code').length;
});

// 计算属性：私密分享数量
const privateShareCount = computed(() => {
  return shareList.value.filter(s => s.shareType?.toLowerCase() === 'need_code').length;
});

// 检查是否过期
const isExpired = (endTime: string | undefined, dayType: number | undefined) => {
  if (!endTime || dayType === 0) return false;
  return new Date(endTime) < new Date();
};

// 格式化日期
const formatDate = (dateStr: string | undefined) => {
  if (!dateStr) return "-";
  try {
    const date = new Date(dateStr);
    return date.toLocaleDateString('zh-CN', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit'
    });
  } catch (error) {
    return dateStr;
  }
};

// 格式化有效期
const formatShareEndTime = (
  endTime: string | undefined,
  dayType: number | undefined
) => {
  if (!endTime) return "-";
  if (dayType === 0) {
    return "永久有效";
  }
  const end = new Date(endTime);
  const now = new Date();
  if (end < now) {
    return "已过期";
  }
  const diffDays = Math.ceil((end.getTime() - now.getTime()) / (1000 * 60 * 60 * 24));
  return `剩余 ${diffDays} 天`;
};

// 获取分享列表
const fetchShareList = async () => {
  try {
    const response = await getShareUrl();
    if (response.data?.code === 0) {
      shareList.value = response.data.data as API.ShareDTO[];
    } else {
      ElMessage.error(response.data?.message || "获取分享列表失败");
    }
  } catch (error) {
    console.error("获取分享列表失败:", error);
    ElMessage.error("获取分享列表失败");
  }
};

// 复制分享信息
const copyShareInfo = (share: API.ShareDTO) => {
  const shareInfo = `分享链接：${share.shareUrl}${
    share.shareCode ? `\n提取码：${share.shareCode}` : ""
  }`;

  if (!navigator.clipboard) {
    const textArea = document.createElement("textarea");
    textArea.value = shareInfo;
    document.body.appendChild(textArea);
    textArea.select();
    try {
      document.execCommand("copy");
      ElMessage.success(
        share.shareType?.toLowerCase() === 'need_code' ? "链接和提取码已复制" : "链接已复制"
      );
    } catch (err) {
      ElMessage.error("复制失败，请手动复制");
    }
    document.body.removeChild(textArea);
    return;
  }

  navigator.clipboard
    .writeText(shareInfo)
    .then(() => {
      ElMessage.success(
        share.shareType?.toLowerCase() === 'need_code' ? "链接和提取码已复制" : "链接已复制"
      );
    })
    .catch(() => {
      ElMessage.error("复制失败，请手动复制");
    });
};

// 取消分享
const handleCancelShare = async (share: API.ShareDTO) => {
  if (!share.id) {
    ElMessage.error("分享ID无效");
    return;
  }

  try {
    await ElMessageBox.confirm("确定要取消此分享吗？", "提示", {
      confirmButtonText: "确定",
      cancelButtonText: "取消",
      type: "warning",
    });

    const response = await cancel({
      shareIds: [share.id],
      accountId: accountId
    });

    if (response.data?.code === 0) {
      ElMessage.success("取消分享成功");
      fetchShareList();
    } else {
      ElMessage.error(response.data?.message || "取消分享失败");
    }
  } catch (error) {
    if (error !== "cancel") {
      console.error("取消分享失败:", error);
      ElMessage.error("取消分享失败");
    }
  }
};

onMounted(() => {
  fetchShareList();
});
</script>

<style scoped>
.my-share-view {
  padding: 24px;
  max-width: 1400px;
  margin: 0 auto;
  min-height: 100vh;
  background: linear-gradient(135deg, #FDF2F8 0%, #F5F3FF 50%, #FCE7F3 100%);
  position: relative;
  overflow: hidden;
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
  background: radial-gradient(circle, rgba(219, 39, 119, 0.4) 0%, transparent 70%);
  top: -100px;
  right: -50px;
  animation-delay: 0s;
}

.bg-orb-2 {
  width: 250px;
  height: 250px;
  background: radial-gradient(circle, rgba(168, 85, 247, 0.3) 0%, transparent 70%);
  bottom: -80px;
  left: -60px;
  animation-delay: -7s;
}

.bg-orb-3 {
  width: 200px;
  height: 200px;
  background: radial-gradient(circle, rgba(244, 114, 182, 0.3) 0%, transparent 70%);
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  animation-delay: -14s;
}

@keyframes floatOrb {
  0%, 100% { transform: translate(0, 0) scale(1); }
  33% { transform: translate(30px, -20px) scale(1.05); }
  66% { transform: translate(-20px, 15px) scale(0.95); }
}

/* 页面标题 */
.page-header {
  margin-bottom: 24px;
}

.header-title {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 8px;
}

.title-icon {
  font-size: 28px;
  color: #DB2777;
  background: linear-gradient(135deg, #DB2777 0%, #A855F7 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.header-title h1 {
  font-size: 24px;
  font-weight: 600;
  color: #1E1B4B;
  margin: 0;
}

.header-desc {
  color: #64748B;
  font-size: 14px;
  margin: 0;
  padding-left: 40px;
}

/* 统计卡片 */
.stats-row {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 16px;
  margin-bottom: 24px;
}

.stat-card {
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: 12px;
  padding: 20px;
  display: flex;
  align-items: center;
  gap: 16px;
  border: 1px solid rgba(219, 39, 119, 0.1);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);
  transition: all 0.2s ease;
  position: relative;
  z-index: 1;
}

.stat-card:hover {
  box-shadow: 0 20px 40px rgba(219, 39, 119, 0.15);
  transform: translateY(-2px);
}

.stat-icon {
  width: 48px;
  height: 48px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 24px;
}

.stat-icon.public {
  background: rgba(16, 185, 129, 0.1);
  color: #10B981;
}

.stat-icon.private {
  background: rgba(245, 158, 11, 0.1);
  color: #F59E0B;
}

.stat-icon.total {
  background: rgba(99, 102, 241, 0.1);
  color: #6366F1;
}

.stat-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.stat-value {
  font-size: 24px;
  font-weight: 700;
  color: #1E1B4B;
  line-height: 1;
}

.stat-label {
  font-size: 13px;
  color: #64748B;
}

/* 分享列表 */
.share-list {
  margin-top: 24px;
}

.empty-state {
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: 16px;
  padding: 64px 24px;
  text-align: center;
  border: 1px dashed rgba(219, 39, 119, 0.2);
  position: relative;
  z-index: 1;
}

.empty-icon {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: linear-gradient(135deg, #FDF2F8 0%, #F5F3FF 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 20px;
  font-size: 36px;
  color: #DB2777;
}

.empty-title {
  font-size: 18px;
  font-weight: 600;
  color: #1E1B4B;
  margin: 0 0 8px;
}

.empty-desc {
  font-size: 14px;
  color: #64748B;
  margin: 0;
}

/* 分享卡片网格 */
.share-cards {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}

/* 大屏幕显示4列 */
@media (min-width: 1400px) {
  .share-cards {
    grid-template-columns: repeat(4, 1fr);
  }
}

/* 中等屏幕显示2列 */
@media (max-width: 1100px) {
  .share-cards {
    grid-template-columns: repeat(2, 1fr);
  }
}

.share-card {
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: 16px;
  border: 1px solid rgba(219, 39, 119, 0.1);
  overflow: hidden;
  transition: all 0.2s ease;
  display: flex;
  flex-direction: column;
  position: relative;
  z-index: 1;
}

.share-card:hover {
  box-shadow: 0 20px 40px rgba(219, 39, 119, 0.15);
  transform: translateY(-4px);
  border-color: rgba(219, 39, 119, 0.3);
}

.share-card.expired {
  opacity: 0.7;
}

.share-card.expired .share-name {
  text-decoration: line-through;
  color: #94A3B8;
}

/* 卡片头部 */
.card-header {
  padding: 20px 20px 0;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.share-icon {
  width: 44px;
  height: 44px;
  border-radius: 12px;
  background: linear-gradient(135deg, #DB2777 0%, #A855F7 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  color: #FFFFFF;
}

.share-type-badge {
  display: flex;
  align-items: center;
  gap: 4px;
  padding: 4px 10px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 500;
}

.share-type-badge.no_code {
  background: rgba(16, 185, 129, 0.1);
  color: #10B981;
}

.share-type-badge.need_code {
  background: rgba(245, 158, 11, 0.1);
  color: #F59E0B;
}

/* 卡片内容 */
.card-body {
  padding: 16px 20px;
  flex: 1;
}

.share-name {
  font-size: 16px;
  font-weight: 600;
  color: #1E1B4B;
  margin: 0 0 12px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  line-height: 1.4;
}

.share-meta {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #64748B;
}

.meta-item .el-icon {
  font-size: 14px;
  color: #94A3B8;
}

.share-code {
  margin-top: 12px;
  padding: 10px 12px;
  background: #F8FAFC;
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.code-label {
  font-size: 12px;
  color: #64748B;
}

.code-value {
  font-size: 14px;
  font-weight: 600;
  color: #DB2777;
  font-family: monospace;
  letter-spacing: 2px;
}

/* 卡片底部操作 */
.card-footer {
  padding: 12px 20px 20px;
  display: flex;
  gap: 10px;
}

.action-btn {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  padding: 10px 16px;
  border-radius: 8px;
  border: none;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.action-btn.copy {
  background: rgba(219, 39, 119, 0.1);
  color: #DB2777;
}

.action-btn.copy:hover {
  background: linear-gradient(135deg, #DB2777 0%, #A855F7 100%);
  color: #FFFFFF;
}

.action-btn.delete {
  background: rgba(239, 68, 68, 0.1);
  color: #EF4444;
}

.action-btn.delete:hover {
  background: #EF4444;
  color: #FFFFFF;
}

/* 响应式适配 */
@media (max-width: 768px) {
  .my-share-view {
    padding: 16px;
  }

  .header-title h1 {
    font-size: 20px;
  }

  .stats-row {
    grid-template-columns: repeat(3, 1fr);
    gap: 12px;
  }

  .stat-card {
    padding: 16px;
    flex-direction: column;
    text-align: center;
    gap: 12px;
  }

  .stat-icon {
    width: 40px;
    height: 40px;
    font-size: 20px;
  }

  .stat-value {
    font-size: 20px;
  }

  .stat-label {
    font-size: 12px;
  }

  .share-cards {
    grid-template-columns: 1fr;
    gap: 16px;
  }

  .share-card {
    border-radius: 12px;
  }
}

@media (max-width: 480px) {
  .stats-row {
    grid-template-columns: repeat(3, 1fr);
  }

  .stat-card {
    padding: 12px 8px;
  }

  .stat-icon {
    width: 32px;
    height: 32px;
    font-size: 16px;
  }

  .stat-value {
    font-size: 16px;
  }

  .stat-label {
    font-size: 11px;
  }
}
</style>
