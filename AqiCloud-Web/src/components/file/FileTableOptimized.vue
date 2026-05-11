<template>
  <div class="file-table-wrapper" ref="wrapperRef">
    <!-- 桌面端表格视图 -->
    <el-table
      v-if="!isMobile"
      :data="fileList"
      style="width: 100%"
      @selection-change="handleSelectionChange"
      :row-class-name="tableRowClassName"
      class="file-table"
      :header-cell-style="headerCellStyle"
      :cell-style="cellStyle"
    >
      <el-table-column 
        type="selection" 
        width="55" 
        class-name="selection-col" 
      />
      
      <el-table-column 
        prop="fileName" 
        label="文件名" 
        min-width="200"
        class-name="filename-col"
      >
        <template #default="scope">
          <div
            class="file-name"
            :aria-label="`文件 ${scope.row.fileName}，点击打开`"
            @click="handleFileClick(scope.row)"
            @contextmenu.prevent="handleContextMenu($event, scope.row)"
            role="button"
            tabindex="0"
            @keydown.enter="handleFileClick(scope.row)"
            @keydown.space.prevent="handleFileClick(scope.row)"
          >
            <DSFileIcon 
              :fileSuffix="scope.row.fileSuffix"
              :isFolder="scope.row.fileType === 'folder' || scope.row.fileType === 'DIR'"
              size="medium"
            />
            <span class="file-name-text">{{ scope.row.fileName }}</span>
          </div>
        </template>
      </el-table-column>
      
      <el-table-column 
        prop="fileType" 
        label="类型" 
        width="100" 
        class-name="type-col"
      >
        <template #default="scope">
          <el-tag size="small" :type="getFileTypeTagType(scope.row.fileType)">
            {{ getFileTypeName(scope.row.fileType) }}
          </el-tag>
        </template>
      </el-table-column>
      
      <el-table-column 
        prop="fileSize" 
        label="大小" 
        width="100" 
        align="right" 
        class-name="size-col"
      >
        <template #default="scope">
          <span class="file-size">{{ formatFileSize(scope.row.fileSize) }}</span>
        </template>
      </el-table-column>
      
      <el-table-column
        prop="gmtModified"
        label="修改日期"
        width="160"
        align="right"
        class-name="date-col"
        sortable
      >
        <template #default="scope">
          <span class="file-date">{{ formatDateTime(scope.row.gmtModified) }}</span>
        </template>
      </el-table-column>
    </el-table>

    <!-- 移动端列表视图 -->
    <div 
      v-else
      class="mobile-file-list"
      @touchstart="handleTouchStart"
      @touchmove="handleTouchMove"
      @touchend="handleTouchEnd"
    >
      <!-- 全选复选框 -->
      <div class="mobile-select-all" v-if="showSelection">
        <el-checkbox 
          v-model="allSelected"
          @change="handleSelectAll"
          class="select-all-checkbox"
        >
          全选
        </el-checkbox>
      </div>

      <!-- 文件列表项 -->
      <div
        v-for="file in fileList"
        :key="file.id"
        class="mobile-file-item"
        :class="{ 'selected': selectedFiles.includes(file.id) }"
        @click="handleMobileFileClick(file)"
        :aria-label="`文件 ${file.fileName}`"
        role="button"
        tabindex="0"
        @keydown.enter="handleMobileFileClick(file)"
        @keydown.space.prevent="handleMobileFileClick(file)"
      >
        <!-- 选择框 -->
        <div class="file-checkbox" v-if="showSelection">
          <el-checkbox 
            :model-value="selectedFiles.includes(file.id)"
            @click.stop
            @change="(val: boolean) => handleMobileSelect(file, val)"
          />
        </div>

        <!-- 文件图标 -->
        <div class="file-icon-wrapper">
          <DSFileIcon 
            :fileSuffix="file.fileSuffix"
            :isFolder="file.fileType === 'folder' || file.fileType === 'DIR'"
            size="medium"
          />
        </div>

        <!-- 文件信息 -->
        <div class="file-info">
          <div class="file-name-mobile">{{ file.fileName }}</div>
          <div class="file-meta">
            <span class="file-size-mobile">{{ formatFileSize(file.fileSize) }}</span>
            <span class="file-date-mobile">{{ formatDateTime(file.gmtModified) }}</span>
          </div>
        </div>

        <!-- 滑动操作按钮 -->
        <div 
          class="swipe-actions"
          :style="{ transform: `translateX(${getSwipeOffset(file.id)}px)` }"
        >
          <button 
            class="swipe-btn swipe-btn-primary"
            @click.stop="handleSwipeAction(file, 'share')"
          >
            分享
          </button>
          <button 
            class="swipe-btn swipe-btn-warning"
            @click.stop="handleSwipeAction(file, 'rename')"
          >
            重命名
          </button>
          <button 
            class="swipe-btn swipe-btn-danger"
            @click.stop="handleSwipeAction(file, 'delete')"
          >
            删除
          </button>
        </div>
      </div>
    </div>

    <!-- 右键菜单 -->
    <ContextMenu
      v-model:visible="showContextMenu"
      :x="contextMenuPosition.x"
      :y="contextMenuPosition.y"
      @action="handleContextMenuAction"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from "vue";
import { DSFileIcon } from "@/components/design-system";
import { formatFileSize, formatDateTime, getFileTypeName } from "@/utils/format";
import ContextMenu from "./ContextMenu.vue";

/**
 * FileTableOptimized 组件 - 优化版文件列表组件
 * 特性：
 * - 响应式设计（桌面端表格 + 移动端列表）
 * - 触摸手势支持（左滑操作）
 * - 无障碍访问性（ARIA标签、键盘导航）
 * - 平滑动画过渡
 */

interface Props {
  fileList: API.FileDTO[];
  currentPath?: string;
  showSelection?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  showSelection: false,
});

const emit = defineEmits<{
  fileClick: [file: API.FileDTO];
  selectionChange: [selection: API.FileDTO[]];
  openRenameDialog: [file: API.FileDTO];
  moveFile: [files: API.FileDTO[], currentPath: string];
  deleteFile: [file: API.FileDTO];
  shareFiles: [files: API.FileDTO[]];
  openFileInfo: [file: API.FileDTO];
  copyFile: [files: API.FileDTO[], currentPath: string];
}>();

// 响应式状态
const isMobile = ref(window.innerWidth <= 768);
const wrapperRef = ref<HTMLElement | null>(null);
const showContextMenu = ref(false);
const contextMenuPosition = ref({ x: 0, y: 0 });
const selectedFile = ref<API.FileDTO | null>(null);
const selectedFiles = ref<(string | number)[]>([]);

// 触摸手势相关
const touchStartX = ref(0);
const touchStartY = ref(0);
const swipeOffsets = ref<Record<string | number, number>>({});
const currentSwipeFileId = ref<string | number | null>(null);

// 计算属性
const allSelected = computed(() => {
  return props.fileList.length > 0 && 
         props.fileList.every(file => selectedFiles.value.includes(file.id));
});

// 监听窗口大小变化
const handleResize = () => {
  isMobile.value = window.innerWidth <= 768;
};

onMounted(() => {
  window.addEventListener('resize', handleResize);
});

onUnmounted(() => {
  window.removeEventListener('resize', handleResize);
});

// 表格样式
const headerCellStyle = {
  background: 'rgba(253, 242, 248, 0.6)',
  fontWeight: '600',
  color: '#1E1B4B',
  fontSize: '13px',
  textTransform: 'uppercase',
  letterSpacing: '0.03em',
};

const cellStyle = {
  color: '#475569',
  fontSize: '14px',
  borderBottomColor: 'rgba(0, 0, 0, 0.04)',
};

// 文件类型标签样式
const getFileTypeTagType = (fileType: string) => {
  const typeMap: Record<string, string> = {
    folder: 'primary',
    image: 'success',
    video: 'warning',
    audio: 'info',
    document: 'info',
  };
  return typeMap[fileType] || 'info';
};

// 触摸手势处理
const handleTouchStart = (event: TouchEvent) => {
  const touch = event.touches[0];
  touchStartX.value = touch.clientX;
  touchStartY.value = touch.clientY;
};

const handleTouchMove = (event: TouchEvent) => {
  if (!currentSwipeFileId.value) return;
  
  const touch = event.touches[0];
  const deltaX = touch.clientX - touchStartX.value;
  
  // 只允许左滑（负值）
  if (deltaX < 0 && Math.abs(deltaX) > Math.abs(touch.clientY - touchStartY.value)) {
    event.preventDefault();
    const offset = Math.max(deltaX, -150); // 最大滑动距离150px
    swipeOffsets.value = {
      ...swipeOffsets.value,
      [currentSwipeFileId.value]: offset,
    };
  }
};

const handleTouchEnd = () => {
  // 如果滑动距离超过一半，显示操作按钮
  if (currentSwipeFileId.value) {
    const offset = swipeOffsets.value[currentSwipeFileId.value] || 0;
    if (offset < -75) {
      swipeOffsets.value = {
        ...swipeOffsets.value,
        [currentSwipeFileId.value]: -150,
      };
    } else {
      swipeOffsets.value = {
        ...swipeOffsets.value,
        [currentSwipeFileId.value]: 0,
      };
    }
  }
  
  // 重置触摸起始位置
  touchStartX.value = 0;
  touchStartY.value = 0;
};

const getSwipeOffset = (fileId: string | number) => {
  return swipeOffsets.value[fileId] || 0;
};

// 文件点击处理
const handleFileClick = (file: API.FileDTO) => {
  emit("fileClick", file);
};

const handleMobileFileClick = (file: API.FileDTO) => {
  // 如果当前有滑动打开的项，先关闭
  if (currentSwipeFileId.value && swipeOffsets.value[currentSwipeFileId.value] < -75) {
    swipeOffsets.value = {
      ...swipeOffsets.value,
      [currentSwipeFileId.value]: 0,
    };
    currentSwipeFileId.value = null;
    return;
  }
  
  emit("fileClick", file);
};

// 选择处理
const handleSelectionChange = (selection: API.FileDTO[]) => {
  emit("selectionChange", selection);
};

const handleMobileSelect = (file: API.FileDTO, selected: boolean) => {
  if (selected) {
    selectedFiles.value.push(file.id);
  } else {
    const index = selectedFiles.value.indexOf(file.id);
    if (index > -1) {
      selectedFiles.value.splice(index, 1);
    }
  }
  
  // 触发选择变更事件
  const selectedFileObjects = props.fileList.filter(file => 
    selectedFiles.value.includes(file.id)
  );
  emit("selectionChange", selectedFileObjects);
};

const handleSelectAll = (selected: boolean) => {
  if (selected) {
    selectedFiles.value = props.fileList.map(file => file.id);
  } else {
    selectedFiles.value = [];
  }
  
  const selectedFileObjects = props.fileList.filter(file => 
    selectedFiles.value.includes(file.id)
  );
  emit("selectionChange", selectedFileObjects);
};

// 滑动操作
const handleSwipeAction = (file: API.FileDTO, action: string) => {
  // 关闭滑动按钮
  if (currentSwipeFileId.value) {
    swipeOffsets.value = {
      ...swipeOffsets.value,
      [currentSwipeFileId.value]: 0,
    };
  }
  
  // 执行对应操作
  switch (action) {
    case 'share':
      emit("shareFiles", [file]);
      break;
    case 'rename':
      emit("openRenameDialog", file);
      break;
    case 'delete':
      emit("deleteFile", file);
      break;
  }
};

// 右键菜单处理
const handleContextMenu = (event: MouseEvent, file: API.FileDTO) => {
  event.preventDefault();
  showContextMenu.value = true;
  contextMenuPosition.value = { x: event.clientX, y: event.clientY };
  selectedFile.value = file;
};

const handleContextMenuAction = (action: string) => {
  if (!selectedFile.value) return;

  switch (action) {
    case "view":
      handleFileClick(selectedFile.value);
      break;
    case "copy":
      emit("copyFile", [selectedFile.value], props.currentPath || "/");
      break;
    case "move":
      emit("moveFile", [selectedFile.value], props.currentPath || "/");
      break;
    case "share":
      emit("shareFiles", [selectedFile.value]);
      break;
    case "rename":
      emit("openRenameDialog", selectedFile.value);
      break;
    case "info":
      emit("openFileInfo", selectedFile.value);
      break;
    case "delete":
      emit("deleteFile", selectedFile.value);
      break;
  }
};

// 表格行类名
const tableRowClassName = () => {
  return "custom-row";
};
</script>

<style scoped>
.file-table-wrapper {
  width: 100%;
  overflow-x: hidden;
}

/* ==================== 桌面端表格样式 ==================== */
.file-table {
  min-width: 600px;
}

:deep(.custom-row) {
  height: 54px;
  transition: all 0.2s ease;
}

:deep(.custom-row:hover) {
  background-color: rgba(219, 39, 119, 0.03) !important;
}

:deep(.el-table__header) th {
  background: rgba(253, 242, 248, 0.5) !important;
  font-weight: 600;
  color: #1E1B4B;
  font-size: 13px;
  text-transform: uppercase;
  letter-spacing: 0.03em;
}

:deep(.el-table__body) td {
  color: #475569;
  font-size: 14px;
  border-bottom-color: rgba(0, 0, 0, 0.04);
}

.file-name {
  display: flex;
  align-items: center;
  cursor: pointer;
  transition: color 0.3s;
  padding: 8px 0;
  border-radius: 8px;
}

.file-name:hover .file-name-text {
  text-decoration: underline;
  color: #DB2777;
}

.file-name-text {
  font-size: 14px;
  color: #1E1B4B;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  margin-left: 12px;
}

.file-size,
.file-date {
  font-size: 13px;
  color: #64748B;
}

/* ==================== 移动端列表样式 ==================== */
.mobile-file-list {
  padding: 8px;
}

.mobile-select-all {
  padding: 12px 16px;
  border-bottom: 1px solid rgba(0, 0, 0, 0.06);
  margin-bottom: 8px;
}

.select-all-checkbox {
  font-size: 14px;
  font-weight: 500;
}

.mobile-file-item {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  margin-bottom: 8px;
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(10px);
  border-radius: 12px;
  border: 1px solid rgba(0, 0, 0, 0.04);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
  cursor: pointer;
  min-height: 60px;
}

.mobile-file-item:hover {
  background: rgba(255, 255, 255, 0.9);
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.08);
}

.mobile-file-item:active {
  transform: scale(0.98);
  background: rgba(219, 39, 119, 0.05);
}

.mobile-file-item.selected {
  background: rgba(219, 39, 119, 0.08);
  border-color: rgba(219, 39, 119, 0.3);
}

.file-checkbox {
  margin-right: 12px;
  flex-shrink: 0;
}

.file-icon-wrapper {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 12px;
  flex-shrink: 0;
}

.file-info {
  flex: 1;
  min-width: 0;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.file-name-mobile {
  font-size: 14px;
  font-weight: 500;
  color: #1E1B4B;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.file-meta {
  display: flex;
  gap: 12px;
  font-size: 12px;
  color: #64748B;
}

.file-size-mobile,
.file-date-mobile {
  font-size: 12px;
}

/* 滑动操作按钮 */
.swipe-actions {
  position: absolute;
  right: 0;
  top: 0;
  bottom: 0;
  display: flex;
  align-items: center;
  gap: 0;
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  z-index: 2;
}

.swipe-btn {
  height: 100%;
  padding: 0 20px;
  border: none;
  color: white;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  min-width: 60px;
}

.swipe-btn:active {
  opacity: 0.8;
  transform: scale(0.95);
}

.swipe-btn-primary {
  background: linear-gradient(135deg, #DB2777 0%, #F472B6 100%);
}

.swipe-btn-warning {
  background: linear-gradient(135deg, #F59E0B 0%, #FBBF24 100%);
}

.swipe-btn-danger {
  background: linear-gradient(135deg, #EF4444 0%, #F87171 100%);
}

/* ==================== 响应式适配 ==================== */
@media (max-width: 768px) {
  .file-table-wrapper {
    overflow-x: hidden;
  }
}

@media (max-width: 480px) {
  .mobile-file-item {
    padding: 10px 12px;
    min-height: 56px;
  }
  
  .file-icon-wrapper {
    width: 36px;
    height: 36px;
    margin-right: 10px;
  }
  
  .file-name-mobile {
    font-size: 13px;
  }
  
  .file-meta {
    font-size: 11px;
    gap: 8px;
  }
  
  .swipe-btn {
    padding: 0 16px;
    font-size: 12px;
    min-width: 50px;
  }
}

/* ==================== 无障碍访问性 ==================== */
.mobile-file-item:focus,
.file-name:focus {
  outline: 2px solid #DB2777;
  outline-offset: 2px;
}

/* 减少动画（用户偏好） */
@media (prefers-reduced-motion: reduce) {
  .mobile-file-item,
  .swipe-actions,
  :deep(.custom-row) {
    transition: none;
  }
}
</style>
