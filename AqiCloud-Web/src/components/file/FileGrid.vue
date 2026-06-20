<template>
  <div class="file-grid">
    <el-row :gutter="20">
      <el-col
        v-for="file in fileList"
        :key="file.id"
        :xs="12"
        :sm="8"
        :md="6"
        :lg="4"
        :xl="3"
      >
        <div
          class="file-item"
          :class="{ 'selection-mode': isSelectionMode }"
          @click="handleFileClick(file)"
          @contextmenu.prevent="handleContextMenu($event, file)"
        >
          <div v-if="isSelectionMode" class="file-checkbox">
            <el-checkbox v-model="selectedFileIds" :label="file.id" @click.stop>
              <!-- 空的标签内容，不显示任何文本 -->
            </el-checkbox>
          </div>
          <div class="file-icon-wrapper">
            <DSFileIcon
              :fileSuffix="file.fileSuffix"
              :isFolder="file.fileType === 'folder' || file.fileType === 'DIR'"
              size="large"
            />
          </div>
          <div class="file-name" :title="file.fileName">
            {{ file.fileName }}
          </div>
        </div>
      </el-col>
    </el-row>

    <ContextMenu
      v-model:visible="showContextMenu"
      :x="contextMenuPosition.x"
      :y="contextMenuPosition.y"
      @action="handleContextMenuAction"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from "vue";
import { Document } from "@element-plus/icons-vue";
import { getFileIcon } from "@/utils/fileIcons";
import ContextMenu from "./ContextMenu.vue";
import type { FileDTO } from "@/api/types";
import { DSFileIcon } from "@/components/design-system";

/**
 * FileGrid 组件 - 文件网格视图组件
 * 提供文件的网格布局展示、多选功能、右键菜单等交互
 */

const props = defineProps<{
  fileList: FileDTO[];
  isSelectionMode: boolean;
  currentPath?: string;
}>();

const emit = defineEmits<{
  fileClick: [file: FileDTO];
  selectionChange: [files: FileDTO[]];
  openRenameDialog: [file: FileDTO];
  moveFile: [files: FileDTO[], path: string];
  deleteFile: [file: FileDTO];
  shareFiles: [files: FileDTO[]];
  openFileInfo: [file: FileDTO];
  copyFile: [files: FileDTO[], path: string];
}>();

// 存储被选中的文件ID数组
const selectedFileIds = ref<number[]>([]);

// 右键菜单相关状态
const showContextMenu = ref(false);
const contextMenuPosition = ref({ x: 0, y: 0 });
const selectedFile = ref<FileDTO | null>(null);

/**
 * 处理文件点击事件
 * 仅在非选择模式下触发文件点击事件
 * @param file - 被点击的文件对象
 */
const handleFileClick = (file: FileDTO): void => {
  if (!props.isSelectionMode) {
    emit("fileClick", file);
  } else {
    // 在选择模式下，切换文件的选中状态
    const index = selectedFileIds.value.indexOf(file.id);
    if (index > -1) {
      selectedFileIds.value.splice(index, 1);
    } else {
      selectedFileIds.value.push(file.id);
    }
    updateSelection();
  }
};

/**
 * 处理右键菜单事件
 * @param event - 鼠标事件
 * @param file - 当前操作的文件对象
 */
const handleContextMenu = (event: MouseEvent, file: FileDTO): void => {
  event.preventDefault();
  showContextMenu.value = true;
  contextMenuPosition.value = { x: event.clientX, y: event.clientY };
  selectedFile.value = file;
};

/**
 * 处理右键菜单选项点击
 * @param action - 操作类型
 */
const handleContextMenuAction = (action: string): void => {
  if (!selectedFile.value) return;

  const actions: Record<string, () => void> = {
    view: () => emit("fileClick", selectedFile.value!),
    copy: () =>
      emit("copyFile", [selectedFile.value!], props.currentPath || "/"),
    move: () =>
      emit("moveFile", [selectedFile.value!], props.currentPath || "/"),
    share: () => emit("shareFiles", [selectedFile.value!]),
    rename: () => emit("openRenameDialog", selectedFile.value!),
    info: () => emit("openFileInfo", selectedFile.value!),
    delete: () => emit("deleteFile", selectedFile.value!),
  };

  const handler = actions[action];
  if (handler) handler();
};

/**
 * 更新选中的文件列表
 * 根据选中的文件ID过滤出完整的文件对象列表，并触发选择变更事件
 */
const updateSelection = (): void => {
  const selectedFileObjects = props.fileList.filter((file) =>
    selectedFileIds.value.includes(file.id),
  );
  emit("selectionChange", selectedFileObjects);
};

/**
 * 监听选中文件变化，触发选择更新
 */
watch(
  selectedFileIds,
  () => {
    updateSelection();
  },
  { immediate: true },
);

/**
 * 监听选择模式变化
 * 当退出选择模式时，清空所有选中状态
 */
watch(
  () => props.isSelectionMode,
  (newMode) => {
    if (!newMode) {
      selectedFileIds.value = [];
      updateSelection();
    }
  },
);
</script>

<style scoped>
.file-grid {
  margin-top: 20px;
  width: 100%;
  overflow-x: hidden;
}

.file-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  padding: 10px;
  position: relative;
  box-sizing: border-box;
  border-radius: 8px;
  user-select: none;
  /* HOK 触摸优化: 最小触控 44px */
  min-height: 44px;
}

.file-item:hover {
  background-color: rgba(255, 255, 255, 0.04);
  transform: translateY(-2px);
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.2);
  border: 1px solid rgba(255, 255, 255, 0.06);
}

.file-item:active {
  transform: translateY(0);
  background-color: rgba(255, 255, 255, 0.06);
}

.file-item.selection-mode {
  cursor: default;
}

.file-checkbox {
  position: absolute;
  top: 5px;
  left: 5px;
  z-index: 1;
}

::v-deep(.file-checkbox .el-checkbox__label) {
  display: none;
}

.file-icon-wrapper {
  width: 80px;
  height: 80px;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 8px;
}

.file-icon {
  max-width: 100%;
  max-height: 100%;
  object-fit: contain;
}

.image-error {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  background-color: rgba(255, 255, 255, 0.03);
  border-radius: 4px;
}

.image-error .el-icon {
  font-size: 40px;
  color: #64748b;
}

.file-name {
  width: 100%;
  text-align: center;
  word-break: break-all;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  font-size: 12px;
  color: #94a3b8;
  line-height: 1.4;
}

@media (max-width: 768px) {
  .file-icon-wrapper {
    width: 60px;
    height: 60px;
  }

  .file-name {
    font-size: 11px;
  }

  .file-item {
    padding: 8px;
    min-height: 44px;
    /* 触摸优化: 增大触控热区 */
    -webkit-tap-highlight-color: transparent;
  }

  .file-item:hover {
    /* 移动端 hover 效果减弱，用 active 替代 */
    transform: none;
    box-shadow: none;
  }

  .file-item:active {
    background-color: rgba(219, 39, 119, 0.08);
    transform: scale(0.97);
  }
}

@media (max-width: 480px) {
  .file-icon-wrapper {
    width: 48px;
    height: 48px;
    margin-bottom: 6px;
  }

  .file-item {
    padding: 6px;
  }

  .file-name {
    font-size: 10px;
    -webkit-line-clamp: 1;
    line-clamp: 1;
  }
}
</style>
