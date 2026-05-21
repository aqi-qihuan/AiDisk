<template>
  <div class="capacity-process" :class="{ 'is-collapse': isCollapse }">
    <template v-if="isCollapse">
      <!-- 圆形进度条 -->
      <div class="circle-container">
        <el-tooltip placement="right">
          <template #content>
            {{ t("user.used") }}: {{ formatStorage(usedStorage) }}<br />
            {{ t("user.totalCapacity") }}: {{ formatStorage(totalStorage) }}
          </template>
          <div class="progress-wrapper">
            <el-progress
              type="circle"
              :percentage="usagePercentage"
              :status="progressStatus"
              :width="36"
              :stroke-width="6"
              :show-text="false"
            />
            <span class="percentage-text">{{ usagePercentage }}%</span>
          </div>
        </el-tooltip>
        <div class="storage-info-mini">
          <span>{{ formatStorage(usedStorage) }}</span>
        </div>
      </div>
    </template>
    <template v-else>
      <!-- 条形进度条 -->
      <div class="storage-title">{{ t("user.storage") }}</div>
      <el-progress
        :percentage="usagePercentage"
        :status="progressStatus"
        :stroke-width="8"
        :show-text="false"
      />
      <div class="storage-info">
        <div class="used-info">
          <span class="label">{{ t("user.used") }}</span>
          <span class="value">{{ formatStorage(usedStorage) }}</span>
        </div>
        <div class="total-info">
          <span class="value">{{ formatStorage(totalStorage) }}</span>
          <span class="label">{{ t("user.totalCapacity") }}</span>
        </div>
      </div>
    </template>
  </div>
</template>

<script setup lang="ts">
import { onMounted, computed } from "vue";
import { useI18n } from "vue-i18n";
import { useStorageStore } from "@/store/storage";

const { t } = useI18n();

// 接收父组件传递的折叠状态
const props = defineProps({
  isCollapse: {
    type: Boolean,
    default: false,
  },
});

const storageStore = useStorageStore();

// 计算已使用存储空间
const usedStorage = computed(() => {
  return Number(storageStore.storageSize || 0);
});

// 计算总存储空间
const totalStorage = computed(() => {
  return Number(storageStore.totalStorageSize || 0);
});

// 计算使用百分比
const usagePercentage = computed(() => {
  if (!totalStorage.value) return 0;
  return Math.round((usedStorage.value / totalStorage.value) * 100);
});

// 根据使用比例决定进度条状态
const progressStatus = computed(() => {
  const percentage = usagePercentage.value;
  if (percentage >= 90) return "exception";
  if (percentage >= 70) return "warning";
  return "success";
});

// 格式化存储大小显示
const formatStorage = (bytes: number): string => {
  if (!bytes) return "0 B";
  const units = ["B", "KB", "MB", "GB", "TB"];
  let value = bytes;
  let unitIndex = 0;

  while (value >= 1024 && unitIndex < units.length - 1) {
    value /= 1024;
    unitIndex++;
  }

  return `${value.toFixed(2)} ${units[unitIndex]}`;
};

onMounted(() => {
  storageStore.fetchStorageInfo();
});
</script>

<style scoped>
.capacity-process {
  padding: 16px;
  background-color: rgba(15, 18, 35, 0.6);
  border: 1px solid rgba(212, 168, 83, 0.1);
  border-radius: 12px;
  transition: all 0.3s;
}

.capacity-process.is-collapse {
  padding: 8px 4px;
}

.circle-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
}

.progress-wrapper {
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
}

.percentage-text {
  position: absolute;
  font-size: 9px;
  color: #e8d5b0;
  font-weight: normal;
}

.storage-info-mini {
  font-size: 8px;
  color: #8b8878;
  line-height: 1;
  margin-top: 2px;
}

.storage-title {
  color: #8b8878;
  font-size: 13px;
  margin-bottom: 12px;
}

.storage-info {
  margin-top: 12px;
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  font-size: 12px;
}

.used-info,
.total-info {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.label {
  color: #8b8878;
}

.value {
  color: #e8d5b0;
  font-weight: 500;
}

:deep(.el-progress-bar__outer) {
  background-color: rgba(20, 22, 40, 0.5);
  border-radius: 4px;
}

:deep(.el-progress-bar__inner) {
  border-radius: 4px;
  background: linear-gradient(90deg, #c9a96e 0%, #d4a853 100%) !important;
  transition: all 0.3s;
}

:deep(.el-progress-circle) {
  transform: rotate(-90deg);
}

:deep(.el-progress-circle path:first-child) {
  opacity: 0.12;
  stroke: rgba(212, 168, 83, 0.3);
}

:deep(.el-progress-circle path:last-child) {
  stroke-linecap: round;
  stroke-width: 6;
}

:deep(.el-progress.is-success path:last-child) {
  stroke: #d4a853;
}

:deep(.el-progress.is-warning path:last-child) {
  stroke: #d4a853;
}

:deep(.el-progress.is-exception path:last-child) {
  stroke: #ef4444;
}
</style>
