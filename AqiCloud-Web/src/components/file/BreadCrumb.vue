<template>
  <div class="breadcrumb-wrapper">
    <div class="title">{{ t("file.currentPath") }}</div>
    <div class="breadcrumb-box" :class="{ 'able-input': isAllFiles }">
      <el-breadcrumb separator=">">
        <el-breadcrumb-item v-if="!isAllFiles">
          {{ currentFileTypeName }}
        </el-breadcrumb-item>
        <template v-else>
          <el-breadcrumb-item @click="handleRootClick">
            {{ t("file.rootDirectory") }}
          </el-breadcrumb-item>
          <el-breadcrumb-item
            v-for="(item, index) in pathParts"
            :key="index"
            @click="handlePartClick(index)"
          >
            {{ item }}
          </el-breadcrumb-item>
        </template>
      </el-breadcrumb>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted } from "vue";
import { useI18n } from "vue-i18n";
import { useFileStores } from "@/store/filePath";
import { useLoginUserStore } from "@/store/user";

const { t } = useI18n();

/**
 * BreadCrumb 组件 - 面包屑导航组件
 * 显示当前文件路径，支持点击导航
 */

interface BreadCrumbProps {
  currentPath: string;
  fileType?: number;
}

const props = defineProps<BreadCrumbProps>();

const emit = defineEmits<{
  pathChange: [path: string, parentId: string | null];
}>();

const inputFilePath = ref(localStorage.getItem("Path"));

const fileTypeMap: Record<number, string> = {
  1: "全部图片",
  2: "全部视频",
  3: "全部音频",
  4: "全部文档",
  5: "压缩文件",
  6: "其他",
};

/**
 * 判断是否显示所有文件
 */
const isAllFiles = computed(() => props.fileType === undefined);

/**
 * 获取当前文件类型名称
 */
const currentFileTypeName = computed(() => {
  return props.fileType !== undefined
    ? fileTypeMap[props.fileType]
    : "全部文件";
});

/**
 * 解析路径为部分数组
 */
const pathParts = computed(() => {
  return props.currentPath.split("/").filter(Boolean);
});

/**
 * 处理根目录点击
 */
const handleRootClick = (): void => {
  const { loginUser } = useLoginUserStore();
  emit("pathChange", "/", loginUser.rootFileId);
};

/**
 * 处理路径部分点击
 * @param index - 路径部分的索引
 */
const handlePartClick = (index: number): void => {
  const newPath = "/" + pathParts.value.slice(0, index + 1).join("/") + "/";
  const { filePaths } = useFileStores();
  emit("pathChange", newPath, filePaths[index + 1]);
};

/**
 * 组件挂载时初始化路径
 */
onMounted(() => {
  if (inputFilePath.value !== props.currentPath) {
    const newPath = inputFilePath.value?.endsWith("/")
      ? inputFilePath.value
      : `${inputFilePath.value}/`;

    emit("pathChange", newPath, localStorage.getItem("parent_id"));
  }
});

/**
 * 监听 fileType 的变化
 */
watch(
  () => props.fileType,
  () => {
    // 可以在这里添加类型切换时的逻辑
  },
);

/**
 * 监听 currentPath 的变化
 */
watch(
  () => props.currentPath,
  (newPath) => {
    inputFilePath.value = newPath;
    localStorage.setItem("Path", inputFilePath.value);
  },
);
</script>

<style scoped>
.breadcrumb-wrapper {
  display: flex;
  align-items: center;
  height: 40px;
  padding: 0 16px;
  border-radius: 8px;
  position: relative;
  overflow: hidden;
  background: rgba(20, 22, 40, 0.85);
  backdrop-filter: blur(16px) saturate(180%);
  -webkit-backdrop-filter: blur(16px) saturate(180%);
  border: 1px solid rgba(212, 168, 83, 0.12);
  color: #e8d5b0;
}

.title {
  margin-right: 12px;
  font-size: 14px;
  color: #8b8878;
  font-weight: 500;
}

.breadcrumb-box {
  flex: 1;
  display: flex;
  align-items: center;
}

.breadcrumb-box.able-input {
  cursor: pointer;
}

::v-deep(.el-breadcrumb__item) {
  cursor: pointer;
}

::v-deep(.el-breadcrumb__inner) {
  color: rgba(232, 213, 176, 0.75);
  font-size: 14px;
  padding: 4px 10px;
  border-radius: 4px;
  transition: all 0.2s;
}

::v-deep(.el-breadcrumb__inner:hover) {
  background-color: rgba(212, 168, 83, 0.12);
  color: #d4a853;
}

::v-deep(.el-breadcrumb__separator) {
  margin: 0 6px;
  font-size: 14px;
  color: rgba(139, 136, 120, 0.6);
}

::v-deep(.el-breadcrumb__item:last-child .el-breadcrumb__inner) {
  color: #e8d5b0;
  font-weight: 600;
  cursor: default;
}

::v-deep(.el-breadcrumb__item:last-child .el-breadcrumb__inner:hover) {
  background-color: transparent;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .breadcrumb-wrapper {
    height: 36px;
    padding: 0 12px;
  }

  .title {
    font-size: 12px;
  }

  ::v-deep(.el-breadcrumb__inner) {
    font-size: 12px;
  }
}
</style>
