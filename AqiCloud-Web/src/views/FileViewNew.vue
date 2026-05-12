<template>
  <div class="file-view">
    <!-- Page Header -->
    <div class="page-header">
      <div class="header-left">
        <h1 class="page-title">文件管理</h1>
        <span class="file-count">共 {{ totalFiles }} 个文件</span>
      </div>
      <div class="header-right">
        <DSButton variant="primary" size="md" @click="showUploadDialog">
          <svg
            width="16"
            height="16"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
          >
            <path d="M12 5v14M5 12h14" />
          </svg>
          上传文件
        </DSButton>
        <DSButton variant="secondary" size="md" @click="createFolder">
          <svg
            width="16"
            height="16"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
          >
            <path
              d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z"
            />
          </svg>
          新建文件夹
        </DSButton>
      </div>
    </div>

    <!-- Toolbar -->
    <div class="toolbar">
      <div class="toolbar-left">
        <!-- Search -->
        <div class="search-box">
          <svg
            class="search-icon"
            width="18"
            height="18"
            viewBox="0 0 24 24"
            fill="none"
            stroke="#9CA3AF"
            stroke-width="2"
          >
            <circle cx="11" cy="11" r="8" />
            <path d="M21 21l-4.35-4.35" />
          </svg>
          <input
            type="text"
            v-model="searchQuery"
            placeholder="搜索文件..."
            class="search-input"
          />
        </div>
      </div>

      <div class="toolbar-right">
        <!-- View Toggle -->
        <div class="view-toggle">
          <button
            :class="['toggle-btn', { active: viewMode === 'list' }]"
            @click="viewMode = 'list'"
          >
            <svg
              width="18"
              height="18"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
            >
              <path d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
          <button
            :class="['toggle-btn', { active: viewMode === 'grid' }]"
            @click="viewMode = 'grid'"
          >
            <svg
              width="18"
              height="18"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
            >
              <rect x="3" y="3" width="7" height="7" />
              <rect x="14" y="3" width="7" height="7" />
              <rect x="3" y="14" width="7" height="7" />
              <rect x="14" y="14" width="7" height="7" />
            </svg>
          </button>
        </div>

        <!-- Sort -->
        <select v-model="sortBy" class="sort-select">
          <option value="name">按名称</option>
          <option value="size">按大小</option>
          <option value="date">按修改时间</option>
        </select>
      </div>
    </div>

    <!-- File List View -->
    <div v-if="viewMode === 'list'" class="file-list">
      <table class="file-table">
        <thead>
          <tr>
            <th class="checkbox-col">
              <input
                type="checkbox"
                :checked="isAllSelected"
                @change="toggleSelectAll"
              />
            </th>
            <th class="name-col">文件名</th>
            <th class="size-col">大小</th>
            <th class="date-col">修改时间</th>
            <th class="actions-col">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="file in filteredFiles"
            :key="file.id"
            :class="{ selected: file.selected }"
          >
            <td class="checkbox-col">
              <input type="checkbox" v-model="file.selected" />
            </td>
            <td class="name-col">
              <div class="file-name">
                <svg
                  width="20"
                  height="20"
                  viewBox="0 0 24 24"
                  fill="none"
                  :stroke="file.isFolder ? '#3B82F6' : '#6B7280'"
                  stroke-width="2"
                >
                  <path
                    :d="
                      file.isFolder
                        ? 'M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z'
                        : 'M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z'
                    "
                  />
                </svg>
                <span>{{ file.name }}</span>
              </div>
            </td>
            <td class="size-col">{{ file.size }}</td>
            <td class="date-col">{{ file.modifiedDate }}</td>
            <td class="actions-col">
              <button class="action-btn" @click="downloadFile(file)">
                下载
              </button>
              <button class="action-btn" @click="shareFile(file)">分享</button>
              <button class="action-btn danger" @click="deleteFile(file)">
                删除
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- File Grid View -->
    <div v-else class="file-grid">
      <div
        v-for="file in filteredFiles"
        :key="file.id"
        :class="['file-card', { selected: file.selected }]"
      >
        <input type="checkbox" v-model="file.selected" class="file-checkbox" />
        <div class="file-icon">
          <svg
            width="48"
            height="48"
            viewBox="0 0 24 24"
            fill="none"
            :stroke="file.isFolder ? '#3B82F6' : '#6B7280'"
            stroke-width="1.5"
          >
            <path
              :d="
                file.isFolder
                  ? 'M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z'
                  : 'M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z'
              "
            />
          </svg>
        </div>
        <div class="file-name">{{ file.name }}</div>
        <div class="file-size">{{ file.size }}</div>
      </div>
    </div>

    <!-- Batch Actions Bar -->
    <div v-if="selectedCount > 0" class="batch-actions">
      <span class="selected-count">已选择 {{ selectedCount }} 个文件</span>
      <div class="batch-buttons">
        <DSButton variant="secondary" size="sm" @click="batchDownload">
          批量下载
        </DSButton>
        <DSButton variant="secondary" size="sm" @click="batchShare">
          批量分享
        </DSButton>
        <DSButton variant="danger" size="sm" @click="batchDelete">
          批量删除
        </DSButton>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from "vue";
import DSButton from "@/components/design-system/DSButton.vue";

// Mock data
const files = ref([
  {
    id: 1,
    name: "项目方案.pdf",
    size: "2.5 MB",
    modifiedDate: "2026-05-10",
    isFolder: false,
    selected: false,
  },
  {
    id: 2,
    name: "设计素材",
    size: "--",
    modifiedDate: "2026-05-08",
    isFolder: true,
    selected: false,
  },
  {
    id: 3,
    name: "会议纪要.docx",
    size: "856 KB",
    modifiedDate: "2026-05-05",
    isFolder: false,
    selected: false,
  },
  {
    id: 4,
    name: "数据分析.xlsx",
    size: "1.2 MB",
    modifiedDate: "2026-05-03",
    isFolder: false,
    selected: false,
  },
]);

const searchQuery = ref("");
const viewMode = ref<"list" | "grid">("list");
const sortBy = ref("name");

// Computed
const totalFiles = computed(() => files.value.length);
const selectedCount = computed(
  () => files.value.filter((f) => f.selected).length,
);
const isAllSelected = computed(
  () => files.value.length > 0 && files.value.every((f) => f.selected),
);
const filteredFiles = computed(() => {
  let result = files.value;

  // Filter by search
  if (searchQuery.value) {
    result = result.filter((f) =>
      f.name.toLowerCase().includes(searchQuery.value.toLowerCase()),
    );
  }

  // Sort
  if (sortBy.value === "name") {
    result = [...result].sort((a, b) => a.name.localeCompare(b.name));
  } else if (sortBy.value === "date") {
    result = [...result].sort((a, b) =>
      b.modifiedDate.localeCompare(a.modifiedDate),
    );
  }

  return result;
});

// Methods
const toggleSelectAll = (event: Event) => {
  const checked = (event.target as HTMLInputElement).checked;
  files.value.forEach((f) => (f.selected = checked));
};

const showUploadDialog = () => {
  console.log("Show upload dialog");
};

const createFolder = () => {
  console.log("Create folder");
};

const downloadFile = (file: (typeof files.value)[0]) => {
  console.log("Download", file.name);
};

const shareFile = (file: (typeof files.value)[0]) => {
  console.log("Share", file.name);
};

const deleteFile = (file: (typeof files.value)[0]) => {
  console.log("Delete", file.name);
};

const batchDownload = () => {
  console.log("Batch download");
};

const batchShare = () => {
  console.log("Batch share");
};

const batchDelete = () => {
  console.log("Batch delete");
};
</script>

<style scoped>
.file-view {
  padding: 32px 48px;
  background: var(--color-gray-50, #f9fafb);
  min-height: 100vh;
}

/* Page Header */
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  flex-wrap: wrap;
  gap: 16px;
}

.header-left {
  display: flex;
  align-items: baseline;
  gap: 12px;
}

.page-title {
  font-size: 30px;
  font-weight: 700;
  color: var(--color-gray-900, #111827);
  margin: 0;
  letter-spacing: -0.02em;
}

.file-count {
  font-size: 14px;
  color: var(--color-gray-500, #6b7280);
}

.header-right {
  display: flex;
  gap: 12px;
}

/* Toolbar */
.toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 24px;
  flex-wrap: wrap;
  gap: 16px;
}

.toolbar-left {
  flex: 1;
  max-width: 400px;
}

.search-box {
  position: relative;
  display: flex;
  align-items: center;
}

.search-icon {
  position: absolute;
  left: 12px;
  pointer-events: none;
}

.search-input {
  width: 100%;
  padding: 8px 12px 8px 40px;
  border: 1px solid var(--color-gray-200, #e5e7eb);
  border-radius: 6px;
  font-size: 14px;
  color: var(--color-gray-900, #111827);
  background: white;
  transition: border-color 0.15s ease;
  box-sizing: border-box;
}

.search-input:focus {
  outline: none;
  border-color: var(--color-primary, #3b82f6);
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
}

.toolbar-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.view-toggle {
  display: flex;
  gap: 4px;
  background: var(--color-gray-100, #f3f4f6);
  padding: 4px;
  border-radius: 6px;
}

.toggle-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  border: none;
  background: transparent;
  color: var(--color-gray-500, #6b7280);
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.15s ease;
}

.toggle-btn:hover {
  background: white;
  color: var(--color-gray-700, #374151);
}

.toggle-btn.active {
  background: white;
  color: var(--color-primary, #3b82f6);
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}

.sort-select {
  padding: 8px 12px;
  border: 1px solid var(--color-gray-200, #e5e7eb);
  border-radius: 6px;
  font-size: 14px;
  color: var(--color-gray-700, #374151);
  background: white;
  cursor: pointer;
  transition: border-color 0.15s ease;
}

.sort-select:focus {
  outline: none;
  border-color: var(--color-primary, #3b82f6);
}

/* File Table */
.file-list {
  background: white;
  border: 1px solid var(--color-gray-200, #e5e7eb);
  border-radius: 8px;
  overflow: hidden;
}

.file-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 14px;
}

.file-table thead tr {
  background: var(--color-gray-50, #f9fafb);
  border-bottom: 1px solid var(--color-gray-200, #e5e7eb);
}

.file-table th {
  padding: 12px 16px;
  text-align: left;
  font-weight: 500;
  font-size: 13px;
  color: var(--color-gray-500, #6b7280);
}

.file-table td {
  padding: 12px 16px;
  color: var(--color-gray-900, #111827);
  border-bottom: 1px solid var(--color-gray-100, #f3f4f6);
}

.file-table tbody tr {
  transition: background-color 0.15s ease;
}

.file-table tbody tr:hover {
  background: var(--color-gray-50, #f9fafb);
}

.file-table tbody tr.selected {
  background: var(--color-primary-light, #eff6ff);
}

.checkbox-col {
  width: 48px;
  text-align: center;
}

.name-col {
  min-width: 300px;
}

.file-name {
  display: flex;
  align-items: center;
  gap: 12px;
}

.size-col {
  width: 120px;
  color: var(--color-gray-500, #6b7280) !important;
}

.date-col {
  width: 150px;
  color: var(--color-gray-500, #6b7280) !important;
}

.actions-col {
  width: 200px;
}

.action-btn {
  background: none;
  border: none;
  color: var(--color-primary, #3b82f6);
  cursor: pointer;
  font-size: 13px;
  padding: 4px 8px;
  margin-right: 8px;
  border-radius: 4px;
  transition: background-color 0.15s ease;
}

.action-btn:hover {
  background: var(--color-gray-50, #f9fafb);
}

.action-btn.danger {
  color: var(--color-error, #ef4444);
}

.action-btn.danger:hover {
  background: #fef2f2;
}

/* File Grid */
.file-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 16px;
}

.file-card {
  position: relative;
  background: white;
  border: 1px solid var(--color-gray-200, #e5e7eb);
  border-radius: 8px;
  padding: 24px 16px 16px;
  text-align: center;
  transition: border-color 0.15s ease;
  cursor: pointer;
}

.file-card:hover {
  border-color: var(--color-gray-300, #d1d5db);
}

.file-card.selected {
  border-color: var(--color-primary, #3b82f6);
  background: var(--color-primary-light, #eff6ff);
}

.file-checkbox {
  position: absolute;
  top: 8px;
  left: 8px;
}

.file-icon {
  margin-bottom: 12px;
  color: var(--color-gray-400, #9ca3af);
}

.file-name {
  font-size: 14px;
  color: var(--color-gray-900, #111827);
  margin-bottom: 4px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.file-size {
  font-size: 12px;
  color: var(--color-gray-500, #6b7280);
}

/* Batch Actions */
.batch-actions {
  position: fixed;
  bottom: 32px;
  left: 50%;
  transform: translateX(-50%);
  background: white;
  border: 1px solid var(--color-gray-200, #e5e7eb);
  border-radius: 8px;
  padding: 16px 24px;
  display: flex;
  align-items: center;
  gap: 24px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  z-index: 100;
}

.selected-count {
  font-size: 14px;
  color: var(--color-gray-700, #374151);
  font-weight: 500;
}

.batch-buttons {
  display: flex;
  gap: 8px;
}

/* Responsive */
@media (max-width: 768px) {
  .file-view {
    padding: 24px 16px;
  }

  .page-header {
    flex-direction: column;
    align-items: flex-start;
  }

  .header-right {
    width: 100%;
  }

  .header-right button {
    flex: 1;
  }

  .toolbar {
    flex-direction: column;
    align-items: stretch;
  }

  .toolbar-left {
    max-width: none;
  }

  .toolbar-right {
    justify-content: space-between;
  }

  .file-table {
    font-size: 13px;
  }

  .file-table th,
  .file-table td {
    padding: 10px 12px;
  }

  .date-col {
    display: none;
  }

  .file-grid {
    grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
    gap: 12px;
  }
}
</style>
