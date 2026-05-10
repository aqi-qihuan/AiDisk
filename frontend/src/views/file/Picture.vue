<template>
  <div class="picture-view">
    <!-- 浮动背景装饰 -->
    <div class="bg-decoration">
      <div class="bg-orb bg-orb-1"></div>
      <div class="bg-orb bg-orb-2"></div>
      <div class="bg-orb bg-orb-3"></div>
    </div>
    <!-- 页面标题 -->
    <div class="page-header">
      <div class="header-content">
        <div class="icon-wrapper">
          <el-icon :size="32"><Picture /></el-icon>
        </div>
        <div class="title-wrapper">
          <h1>图片管理</h1>
          <p class="subtitle">浏览和管理您的图片文件</p>
        </div>
      </div>
      <div class="header-actions">
        <DSTag variant="info" class="stats-tag">
          共 {{ total }} 张图片
        </DSTag>
        <DSButton
          variant="primary"
          size="sm"
          @click="openUploadDialog"
          class="upload-btn"
        >
          <el-icon><Upload /></el-icon>
          上传图片
        </DSButton>
      </div>
    </div>

    <!-- 工具栏 -->
    <div class="toolbar">
      <div class="left-tools">
        <el-input
          v-model="searchKeyword"
          placeholder="搜索图片..."
          prefix-icon="Search"
          clearable
          class="search-input"
          @clear="handleSearch"
          @keyup.enter="handleSearch"
        >
          <template #append>
            <el-button @click="handleSearch" :icon="Search">搜索</el-button>
          </template>
        </el-input>
        
        <div class="filter-group">
          <span class="filter-label">排序:</span>
          <el-select v-model="sortBy" placeholder="排序方式" class="sort-select" @change="handleSort">
            <el-option label="按名称" value="name" />
            <el-option label="按大小" value="size" />
            <el-option label="按日期" value="date" />
          </el-select>
        </div>
      </div>
      
      <div class="right-tools">
        <div class="view-mode-group">
          <el-tooltip content="网格视图" placement="top">
            <el-button
              :type="viewMode === 'grid' ? 'primary' : 'default'"
              :icon="Grid"
              circle
              @click="viewMode = 'grid'"
            />
          </el-tooltip>
          <el-tooltip content="列表视图" placement="top">
            <el-button
              :type="viewMode === 'list' ? 'primary' : 'default'"
              :icon="List"
              circle
              @click="viewMode = 'list'"
            />
          </el-tooltip>
        </div>
      </div>
    </div>

    <!-- 图片网格视图 -->
    <div v-if="viewMode === 'grid'" class="picture-grid">
      <div
        v-for="picture in pictureList"
        :key="picture.id"
        class="picture-card"
        @click="handlePictureClick(picture)"
      >
        <div class="picture-thumbnail">
          <img
            :src="getFileUrl(picture)"
            :alt="picture.fileName"
            @error="handleImageError"
          />
          <div class="picture-overlay">
            <div class="overlay-actions">
              <el-tooltip content="预览" placement="top">
                <el-button
                  circle
                  :icon="View"
                  @click.stop="handlePreview(picture)"
                />
              </el-tooltip>
              <el-tooltip content="下载" placement="top">
                <el-button
                  circle
                  :icon="Download"
                  @click.stop="handleDownload(picture)"
                />
              </el-tooltip>
              <el-tooltip content="删除" placement="top">
                <el-button
                  circle
                  type="danger"
                  :icon="Delete"
                  @click.stop="handleDelete(picture)"
                />
              </el-tooltip>
            </div>
          </div>
        </div>
        <div class="picture-info">
          <div class="picture-name" :title="picture.fileName">
            {{ picture.fileName }}
          </div>
          <div class="picture-meta">
            <span class="picture-size">{{ formatFileSize(picture.fileSize) }}</span>
            <span class="picture-date">{{ formatDate(picture.updateTime) }}</span>
          </div>
        </div>
      </div>

      <!-- 空状态 -->
      <div v-if="pictureList.length === 0 && !loading" class="empty-state">
        <div class="empty-icon">
          <el-icon :size="80"><Picture /></el-icon>
        </div>
        <h3>暂无图片</h3>
        <p>上传您的第一张图片开始使用吧</p>
        <DSButton variant="primary" @click="openUploadDialog">
          <el-icon><Upload /></el-icon>
          上传图片
        </DSButton>
      </div>
    </div>

    <!-- 图片列表视图 -->
    <div v-else class="picture-list">
      <el-table
        :data="pictureList"
        style="width: 100%"
        @row-click="handlePictureClick"
        v-loading="loading"
      >
        <el-table-column label="预览" width="100">
          <template #default="{ row }">
            <div class="table-thumbnail">
              <img
                :src="getFileUrl(row)"
                :alt="row.fileName"
                @error="handleImageError"
              />
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="fileName" label="文件名" min-width="200">
          <template #default="{ row }">
            <div class="file-name-cell">
              <el-icon><Picture /></el-icon>
              <span class="file-name">{{ row.fileName }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="fileSize" label="大小" width="120">
          <template #default="{ row }">
            {{ formatFileSize(row.fileSize) }}
          </template>
        </el-table-column>
        <el-table-column prop="updateTime" label="修改时间" width="180">
          <template #default="{ row }">
            {{ formatDate(row.updateTime) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <div class="table-actions">
              <el-button
                link
                type="primary"
                :icon="View"
                @click.stop="handlePreview(row)"
              >
                预览
              </el-button>
              <el-button
                link
                type="primary"
                :icon="Download"
                @click.stop="handleDownload(row)"
              >
                下载
              </el-button>
              <el-button
                link
                type="danger"
                :icon="Delete"
                @click.stop="handleDelete(row)"
              >
                删除
              </el-button>
            </div>
          </template>
        </el-table-column>
      </el-table>

      <!-- 空状态 -->
      <div v-if="pictureList.length === 0 && !loading" class="empty-state">
        <div class="empty-icon">
          <el-icon :size="80"><Picture /></el-icon>
        </div>
        <h3>暂无图片</h3>
        <p>上传您的第一张图片开始使用吧</p>
        <DSButton variant="primary" @click="openUploadDialog">
          <el-icon><Upload /></el-icon>
          上传图片
        </DSButton>
      </div>
    </div>

    <!-- 分页 -->
    <div v-if="total > 0" class="pagination-container">
      <DSTag variant="info" class="total-info">
        共 {{ total }} 张图片
      </DSTag>
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :total="total"
        :page-sizes="[20, 40, 60, 100]"
        layout="sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>

    <!-- 上传对话框 -->
    <el-dialog
      v-model="uploadDialogVisible"
      title="上传图片"
      :width="isMobile ? '95%' : '60%'"
      :fullscreen="isMobile"
      class="ds-dialog"
    >
      <FileUpload
        :currentPath="'/'"
        :acceptTypes="'image/*'"
        @upload-success="handleUploadSuccess"
      />
    </el-dialog>

    <!-- 图片预览对话框 -->
    <el-dialog
      v-model="previewDialogVisible"
      :title="currentPicture?.fileName"
      width="80%"
      class="preview-dialog"
    >
      <div class="preview-content">
        <img
          v-if="currentPicture"
          :src="getFileUrl(currentPicture)"
          :alt="currentPicture.fileName"
          class="preview-image"
        />
      </div>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from 'vue';
import { ElMessage, ElMessageBox } from 'element-plus';
import {
  Picture,
  Upload,
  Download,
  Delete,
  View,
  Search,
  Grid,
  List
} from '@element-plus/icons-vue';
import { useFileStore } from '@/store/file';
import DSButton from '@/components/design-system/DSButton.vue';
import DSTag from '@/components/design-system/DSTag.vue';
import FileUpload from '@/components/file/FileUpload.vue';
import request from '@/utils/request';

const fileStore = useFileStore();

// 响应式数据
const pictureList = ref<API.FileDTO[]>([]);
const loading = ref(false);
const searchKeyword = ref('');
const sortBy = ref('date');
const viewMode = ref<'grid' | 'list'>('grid');
const currentPage = ref(1);
const pageSize = ref(20);
const total = ref(0);
const uploadDialogVisible = ref(false);
const previewDialogVisible = ref(false);
const currentPicture = ref<API.FileDTO | null>(null);
const imageBlobUrls = ref<Map<number, string>>(new Map());

// 计算属性
const isMobile = computed(() => {
  return window.innerWidth < 768;
});

// 获取图片 Blob URL（与 PDF 预览同机制，使用 myAxios 带 token header）
const getFileUrl = (file: API.FileDTO): string => {
  if (imageBlobUrls.value.has(file.id)) {
    return imageBlobUrls.value.get(file.id)!;
  }
  loadBlobUrl(file);
  return '';
};

const loadBlobUrl = async (file: API.FileDTO) => {
  try {
    const response = await request(`/file/v1/preview`, {
      method: 'GET',
      params: { fileId: file.id },
      responseType: 'blob',
    });
    const mimeType = response.headers['content-type'] || 'image/jpeg';
    const url = URL.createObjectURL(new Blob([response.data], { type: mimeType }));
    imageBlobUrls.value.set(file.id, url);
  } catch (error) {
    console.error('获取图片预览失败:', error);
  }
};

const cleanupBlobUrls = () => {
  imageBlobUrls.value.forEach(url => URL.revokeObjectURL(url));
  imageBlobUrls.value.clear();
};

// 格式化文件大小
const formatFileSize = (bytes: number) => {
  if (bytes === 0) return '0 B';
  const k = 1024;
  const sizes = ['B', 'KB', 'MB', 'GB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return Math.round(bytes / Math.pow(k, i) * 100) / 100 + ' ' + sizes[i];
};

// 格式化日期
const formatDate = (dateString: string) => {
  const date = new Date(dateString);
  return date.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  });
};

// 获取图片列表
const fetchPictures = async () => {
  loading.value = true;
  cleanupBlobUrls();
  try {
    // 调用API获取图片类型的文件
    const response = await fileStore.getFileList({
      current: currentPage.value,
      size: pageSize.value,
      path: '/',
      keyword: searchKeyword.value,
      fileType: 'image' // 图片类型
    });

    if (response.data.data) {
      pictureList.value = response.data.data;
      total.value = response.data.data.length || 0;
    }
  } catch (error) {
    console.error('获取图片列表失败:', error);
    ElMessage.error('获取图片列表失败');
  } finally {
    loading.value = false;
  }
};

// 搜索处理
const handleSearch = () => {
  currentPage.value = 1;
  fetchPictures();
};

// 排序处理
const handleSort = () => {
  // 根据排序方式对列表进行排序
  const list = [...pictureList.value];
  
  switch (sortBy.value) {
    case 'name':
      list.sort((a, b) => a.fileName.localeCompare(b.fileName));
      break;
    case 'size':
      list.sort((a, b) => b.fileSize - a.fileSize);
      break;
    case 'date':
      list.sort((a, b) => new Date(b.updateTime).getTime() - new Date(a.updateTime).getTime());
      break;
  }
  
  pictureList.value = list;
};

// 图片点击处理
const handlePictureClick = (picture: API.FileDTO) => {
  currentPicture.value = picture;
  previewDialogVisible.value = true;
};

// 预览处理
const handlePreview = (picture: API.FileDTO) => {
  currentPicture.value = picture;
  previewDialogVisible.value = true;
};

// 下载处理
const handleDownload = async (picture: API.FileDTO) => {
  try {
    const response = await request(`/file/v1/preview`, {
      method: 'GET',
      params: { fileId: picture.id },
      responseType: 'blob',
    });
    const blob = new Blob([response.data]);
    const url = URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;
    link.download = picture.fileName;
    link.click();
    URL.revokeObjectURL(url);
    ElMessage.success('开始下载');
  } catch (error) {
    console.error('下载失败:', error);
    ElMessage.error('下载失败');
  }
};

// 删除处理
const handleDelete = async (picture: API.FileDTO) => {
  try {
    await ElMessageBox.confirm(
      `确定要删除图片 "${picture.fileName}" 吗?`,
      '删除确认',
      {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }
    );

    // 调用删除API
    await fileStore.deleteFile(picture.id);
    ElMessage.success('删除成功');
    fetchPictures();
  } catch (error) {
    if (error !== 'cancel') {
      console.error('删除失败:', error);
      ElMessage.error('删除失败');
    }
  }
};

// 图片加载错误处理
const handleImageError = (e: Event) => {
  const target = e.target as HTMLImageElement;
  target.src = '/image/default-image.png';
};

// 打开上传对话框
const openUploadDialog = () => {
  uploadDialogVisible.value = true;
};

// 上传成功处理
const handleUploadSuccess = () => {
  uploadDialogVisible.value = false;
  ElMessage.success('上传成功');
  fetchPictures();
};

// 分页处理
const handleSizeChange = (size: number) => {
  pageSize.value = size;
  fetchPictures();
};

const handleCurrentChange = (page: number) => {
  currentPage.value = page;
  fetchPictures();
};

// 初始化
onMounted(() => {
  fetchPictures();
});

onUnmounted(() => {
  cleanupBlobUrls();
});
</script>

<style scoped>
.picture-view {
  min-height: 100vh;
  padding: var(--ds-spacing-lg);
  background: linear-gradient(135deg, #FDF2F8 0%, #F5F3FF 50%, #FCE7F3 100%);
  position: relative;
  overflow: hidden;
}

/* 浮动背景装饰球 */
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
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--ds-spacing-xl);
  padding: var(--ds-spacing-lg);
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: var(--ds-radius-lg);
  box-shadow: var(--ds-shadow-md);
  position: relative;
  z-index: 1;
}

.header-content {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-md);
}

.icon-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 64px;
  height: 64px;
  background: linear-gradient(135deg, #DB2777 0%, #A855F7 100%);
  border-radius: var(--ds-radius-lg);
  color: white;
  animation: bounce 2s infinite;
}

.title-wrapper h1 {
  margin: 0;
  font-size: var(--ds-text-size-xxl);
  font-weight: var(--ds-text-weight-bold);
  color: var(--ds-color-text-primary);
}

.subtitle {
  margin: var(--ds-spacing-xs) 0 0 0;
  font-size: var(--ds-text-size-md);
  color: var(--ds-color-text-secondary);
}

.header-actions {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-md);
}

.stats-tag {
  font-size: var(--ds-text-size-sm);
}

.upload-btn {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-xs);
}

/* 工具栏 */
.toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--ds-spacing-lg);
  padding: var(--ds-spacing-md);
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: var(--ds-radius-lg);
  box-shadow: var(--ds-shadow-sm);
  position: relative;
  z-index: 1;
}

.left-tools {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-md);
  flex: 1;
}

.search-input {
  width: 300px;
}

.filter-group {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-sm);
}

.filter-label {
  font-size: var(--ds-text-size-sm);
  color: var(--ds-color-text-secondary);
}

.sort-select {
  width: 120px;
}

.right-tools {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-md);
}

.view-mode-group {
  display: flex;
  gap: var(--ds-spacing-xs);
}

/* 图片网格视图 */
.picture-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: var(--ds-spacing-lg);
  margin-bottom: var(--ds-spacing-xl);
}

.picture-card {
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: var(--ds-radius-lg);
  overflow: hidden;
  box-shadow: var(--ds-shadow-sm);
  cursor: pointer;
  transition: all var(--ds-transition-base);
  position: relative;
  z-index: 1;
}

.picture-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 20px 40px rgba(219, 39, 119, 0.15);
}

.picture-thumbnail {
  position: relative;
  width: 100%;
  height: 200px;
  overflow: hidden;
  background: #f5f5f5;
}

.picture-thumbnail img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform var(--ds-transition-base);
}

.picture-card:hover .picture-thumbnail img {
  transform: scale(1.1);
}

.picture-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.6);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity var(--ds-transition-base);
}

.picture-card:hover .picture-overlay {
  opacity: 1;
}

.overlay-actions {
  display: flex;
  gap: var(--ds-spacing-sm);
}

.picture-info {
  padding: var(--ds-spacing-md);
}

.picture-name {
  font-size: var(--ds-text-size-md);
  font-weight: var(--ds-text-weight-medium);
  color: var(--ds-color-text-primary);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.picture-meta {
  display: flex;
  justify-content: space-between;
  margin-top: var(--ds-spacing-xs);
  font-size: var(--ds-text-size-sm);
  color: var(--ds-color-text-secondary);
}

/* 图片列表视图 */
.picture-list {
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: var(--ds-radius-lg);
  overflow: hidden;
  box-shadow: var(--ds-shadow-sm);
  margin-bottom: var(--ds-spacing-xl);
  position: relative;
  z-index: 1;
}

.table-thumbnail {
  width: 80px;
  height: 60px;
  border-radius: var(--ds-radius-sm);
  overflow: hidden;
}

.table-thumbnail img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.file-name-cell {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-sm);
}

.file-name {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.table-actions {
  display: flex;
  gap: var(--ds-spacing-xs);
}

/* 空状态 */
.empty-state {
  grid-column: 1 / -1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: var(--ds-spacing-xxl) var(--ds-spacing-lg);
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: var(--ds-radius-lg);
  box-shadow: var(--ds-shadow-sm);
  position: relative;
  z-index: 1;
}

.empty-icon {
  margin-bottom: var(--ds-spacing-lg);
  color: var(--ds-color-primary);
  animation: float 3s ease-in-out infinite;
}

.empty-state h3 {
  margin: 0 0 var(--ds-spacing-sm) 0;
  font-size: var(--ds-text-size-xl);
  font-weight: var(--ds-text-weight-medium);
  color: var(--ds-color-text-primary);
}

.empty-state p {
  margin: 0 0 var(--ds-spacing-lg) 0;
  font-size: var(--ds-text-size-md);
  color: var(--ds-color-text-secondary);
}

/* 分页 */
.pagination-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--ds-spacing-md);
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: var(--ds-radius-lg);
  box-shadow: var(--ds-shadow-sm);
  position: relative;
  z-index: 1;
}

.total-info {
  font-size: var(--ds-text-size-sm);
}

/* 预览对话框 */
.preview-dialog :deep(.el-dialog__body) {
  padding: 0;
}

.preview-content {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 400px;
  background: #f5f5f5;
}

.preview-image {
  max-width: 100%;
  max-height: 70vh;
  object-fit: contain;
}

/* 动画 */
@keyframes bounce {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
}

@keyframes float {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-20px);
  }
}

/* 响应式设计 */
@media (max-width: 1024px) {
  .picture-grid {
    grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
  }
  
  .search-input {
    width: 200px;
  }
}

@media (max-width: 768px) {
  .picture-view {
    padding: var(--ds-spacing-md);
  }
  
  .page-header {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--ds-spacing-md);
  }
  
  .header-actions {
    width: 100%;
    justify-content: space-between;
  }
  
  .toolbar {
    flex-direction: column;
    gap: var(--ds-spacing-md);
  }
  
  .left-tools {
    width: 100%;
    flex-direction: column;
  }
  
  .search-input {
    width: 100%;
  }
  
  .filter-group {
    width: 100%;
  }
  
  .sort-select {
    flex: 1;
  }
  
  .right-tools {
    width: 100%;
    justify-content: center;
  }
  
  .picture-grid {
    grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
    gap: var(--ds-spacing-md);
  }
  
  .picture-thumbnail {
    height: 140px;
  }
  
  .icon-wrapper {
    width: 48px;
    height: 48px;
  }
  
  .icon-wrapper :deep(.el-icon) {
    font-size: 24px;
  }
  
  .title-wrapper h1 {
    font-size: var(--ds-text-size-xl);
  }
}
</style>
