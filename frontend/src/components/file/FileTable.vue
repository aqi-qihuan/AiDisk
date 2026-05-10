<template>
  <div class="file-table-wrapper">
    <el-table
      :data="fileList"
      style="width: 100%"
      @selection-change="handleSelectionChange"
      :row-class-name="tableRowClassName"
      class="file-table"
    >
      <el-table-column type="selection" width="55" class-name="selection-col" />
      <el-table-column prop="fileName" label="文件名" min-width="200">
        <template #default="scope">
          <div
            class="file-name"
            @click="handleFileClick(scope.row)"
            @contextmenu.prevent="handleContextMenu($event, scope.row)"
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
      <el-table-column prop="fileType" label="类型" width="100" class-name="type-col">
        <template #default="scope">
          {{ getFileTypeName(scope.row.fileType) }}
        </template>
      </el-table-column>
      <el-table-column prop="fileSize" label="大小" width="100" align="right" class-name="size-col">
        <template #default="scope">
          {{ formatFileSize(scope.row.fileSize) }}
        </template>
      </el-table-column>
      <el-table-column
        prop="gmtModified"
        label="修改日期"
        width="160"
        align="right"
        class-name="date-col"
      >
        <template #default="scope">
          {{ formatDateTime(scope.row.gmtModified) }}
        </template>
      </el-table-column>
    </el-table>
  </div>

  <ContextMenu
    v-model:visible="showContextMenu"
    :x="contextMenuPosition.x"
    :y="contextMenuPosition.y"
    @action="handleContextMenuAction"
  />
</template>

<script setup lang="ts">
import { ref } from "vue";
import { DSFileIcon } from "@/components/design-system";
import { formatFileSize, formatDateTime, getFileTypeName } from "@/utils/format";
import ContextMenu from "./ContextMenu.vue";

/**
 * FileTable 组件
 * 以表格形式展示文件列表，支持选择、右键菜单等操作
 */

const props = defineProps<{
  fileList: API.FileDTO[];
  currentPath?: string;
}>();

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

// 右键菜单相关状态
const showContextMenu = ref(false);
const contextMenuPosition = ref({ x: 0, y: 0 });
const selectedFile = ref<API.FileDTO | null>(null);

/**
 * 处理右键菜单事件
 * @param event - 鼠标事件
 * @param file - 当前操作的文件对象
 */
const handleContextMenu = (event: MouseEvent, file: API.FileDTO): void => {
  event.preventDefault();
  showContextMenu.value = true;
  contextMenuPosition.value = { x: event.clientX, y: event.clientY };
  selectedFile.value = file;
};

/**
 * 处理右键菜单选项点击
 * @param action - 操作类型：'view'|'move'|'share'|'rename'|'info'|'delete'|'copy'
 */
const handleContextMenuAction = (action: string): void => {
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

/**
 * 处理文件点击事件
 * @param file - 被点击的文件对象
 */
const handleFileClick = (file: API.FileDTO): void => {
  emit("fileClick", file);
};

/**
 * 处理表格选择变化事件
 * @param selection - 当前选中的文件列表
 */
const handleSelectionChange = (selection: API.FileDTO[]): void => {
  emit("selectionChange", selection);
};

/**
 * 设置表格行的类名
 * @returns 自定义行类名
 */
const tableRowClassName = (): string => {
  return "custom-row";
};
</script>

<style scoped>
.file-table-wrapper {
  width: 100%;
  overflow-x: auto;
}

.file-table {
  min-width: 600px;
}

.file-name {
  display: flex;
  align-items: center;
  cursor: pointer;
  transition: color 0.3s;
}

.file-name:hover .file-name-text {
  text-decoration: underline;
  color: aqua;
}

.file-icon {
  width: 28px;
  height: 28px;
  margin-right: 14px;
  object-fit: contain;
}

.file-name-text {
  font-size: 18px;
  color: black;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

::v-deep(.custom-row) {
  height: 54px;
}

::v-deep(.el-table__header) th {
  font-size: 15px;
  font-weight: bold;
  background-color: #f5f7fa;
  padding: 14px 0;
}

::v-deep(.el-table__body) td {
  padding: 14px 0;
  font-size: 14px;
}

::v-deep(.el-table--enable-row-hover .el-table__body tr:hover > td) {
  background-color: #c6d5ed;
}

::v-deep(.el-checkbox__inner) {
  width: 18px;
  height: 18px;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .file-table-wrapper {
    overflow-x: hidden;
    -webkit-overflow-scrolling: touch;
  }

  .file-table {
    min-width: 100%;
  }

  /* 选择框列缩小 */
  ::v-deep(.selection-col) {
    width: 40px !important;
    min-width: 40px !important;
  }

  ::v-deep(.selection-col) .el-checkbox {
    margin: 0;
  }

  /* 文件名列占据主要空间 */
  ::v-deep(.el-table__body) td:nth-child(2) {
    min-width: auto;
    max-width: none;
  }

  /* 类型列 */
  ::v-deep(.type-col) {
    width: 70px !important;
    min-width: 70px !important;
  }

  /* 大小列 */
  ::v-deep(.size-col) {
    width: 80px !important;
    min-width: 80px !important;
  }

  /* 日期列 */
  ::v-deep(.date-col) {
    width: 100px !important;
    min-width: 100px !important;
  }

  .file-name {
    min-width: 0;
  }

  .file-name-text {
    font-size: 14px;
    max-width: none;
    flex: 1;
    min-width: 0;
  }

  ::v-deep(.el-table__header) th {
    font-size: 13px;
    padding: 10px 4px;
  }

  ::v-deep(.el-table__body) td {
    padding: 10px 4px;
    font-size: 13px;
  }

  ::v-deep(.custom-row) {
    height: 48px;
  }

  /* 文件图标缩小 */
  ::v-deep(.file-name) .ds-file-icon {
    width: 32px;
    height: 32px;
    flex-shrink: 0;
  }
}

@media (max-width: 480px) {
  .file-table {
    min-width: 100%;
  }

  /* 小屏幕隐藏日期列 */
  ::v-deep(.date-col) {
    display: none !important;
  }

  /* 更小屏幕选择框进一步缩小 */
  ::v-deep(.selection-col) {
    width: 36px !important;
    min-width: 36px !important;
  }

  .file-name-text {
    font-size: 13px;
  }

  ::v-deep(.el-table__header) th {
    font-size: 12px;
    padding: 8px 2px;
  }

  ::v-deep(.el-table__body) td {
    padding: 8px 2px;
    font-size: 12px;
  }

  /* 类型和大小列更紧凑 */
  ::v-deep(.type-col) {
    width: 60px !important;
    min-width: 60px !important;
  }

  ::v-deep(.size-col) {
    width: 70px !important;
    min-width: 70px !important;
  }
}
</style>
