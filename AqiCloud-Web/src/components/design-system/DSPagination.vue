<template>
  <div class="ds-pagination">
    <!-- Total Info -->
    <span v-if="showTotal" class="ds-pagination-total">
      {{ t("common.total", { count: total }) }}
    </span>

    <!-- Page Size Selector -->
    <div v-if="showSizeChanger" class="ds-pagination-size-changer">
      <select :value="pageSize" @change="handleSizeChange">
        <option v-for="size in pageSizes" :key="size" :value="size">
          {{ t("common.perPage", { size }) }}
        </option>
      </select>
    </div>

    <!-- Pager -->
    <div class="ds-pagination-pager">
      <!-- Prev Button -->
      <button
        class="ds-pagination-btn ds-pagination-prev"
        :disabled="currentPage <= 1"
        @click="handlePrev"
      >
        <svg viewBox="0 0 24 24" width="16" height="16">
          <path
            fill="currentColor"
            d="M15.41 7.41L14 6l-6 6 6 6 1.41-1.41L10.83 12z"
          />
        </svg>
      </button>

      <!-- Page Numbers -->
      <template v-for="(item, index) in pagerList" :key="index">
        <button
          v-if="item === '...'"
          class="ds-pagination-btn ds-pagination-ellipsis"
          disabled
        >
          ...
        </button>
        <button
          v-else
          :class="[
            'ds-pagination-btn',
            'ds-pagination-page',
            { 'is-active': item === currentPage },
          ]"
          @click="handlePageClick(item as number)"
        >
          {{ item }}
        </button>
      </template>

      <!-- Next Button -->
      <button
        class="ds-pagination-btn ds-pagination-next"
        :disabled="currentPage >= pageCount"
        @click="handleNext"
      >
        <svg viewBox="0 0 24 24" width="16" height="16">
          <path
            fill="currentColor"
            d="M10 6L8.59 7.41 13.17 12l-4.58 4.59L10 18l6-6z"
          />
        </svg>
      </button>
    </div>

    <!-- Jumper -->
    <div v-if="showJumper" class="ds-pagination-jumper">
      <span>{{ t("common.goTo") }}</span>
      <input
        type="number"
        :value="currentPage"
        min="1"
        :max="pageCount"
        @keyup.enter="handleJumperChange"
      />
      <span>{{ t("common.page") }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from "vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

interface Props {
  total: number;
  pageSize?: number;
  currentPage?: number;
  pageSizes?: number[];
  showTotal?: boolean;
  showSizeChanger?: boolean;
  showJumper?: boolean;
  pagerCount?: number;
}

const props = withDefaults(defineProps<Props>(), {
  total: 0,
  pageSize: 10,
  currentPage: 1,
  pageSizes: () => [10, 20, 50, 100],
  showTotal: true,
  showSizeChanger: true,
  showJumper: true,
  pagerCount: 7,
});

const emit = defineEmits<{
  "update:currentPage": [page: number];
  "update:pageSize": [size: number];
  change: [page: number, size: number];
  "size-change": [size: number];
  "current-change": [page: number];
}>();

const pageCount = computed(() => Math.ceil(props.total / props.pageSize));

const pagerList = computed(() => {
  const pages: (number | string)[] = [];
  const total = pageCount.value;
  const current = props.currentPage;
  const count = props.pagerCount;

  if (total <= count) {
    for (let i = 1; i <= total; i++) {
      pages.push(i);
    }
  } else {
    const half = Math.floor(count / 2);
    let start = Math.max(1, current - half);
    let end = Math.min(total, start + count - 1);

    if (end - start + 1 < count) {
      start = Math.max(1, end - count + 1);
    }

    if (start > 1) {
      pages.push(1);
      if (start > 2) {
        pages.push("...");
      }
    }

    for (let i = start; i <= end; i++) {
      pages.push(i);
    }

    if (end < total) {
      if (end < total - 1) {
        pages.push("...");
      }
      pages.push(total);
    }
  }

  return pages;
});

const handlePrev = () => {
  if (props.currentPage > 1) {
    const newPage = props.currentPage - 1;
    emit("update:currentPage", newPage);
    emit("current-change", newPage);
    emit("change", newPage, props.pageSize);
  }
};

const handleNext = () => {
  if (props.currentPage < pageCount.value) {
    const newPage = props.currentPage + 1;
    emit("update:currentPage", newPage);
    emit("current-change", newPage);
    emit("change", newPage, props.pageSize);
  }
};

const handlePageClick = (page: number) => {
  if (page !== props.currentPage) {
    emit("update:currentPage", page);
    emit("current-change", page);
    emit("change", page, props.pageSize);
  }
};

const handleSizeChange = (e: Event) => {
  const target = e.target as HTMLSelectElement;
  const newSize = Number(target.value);
  emit("update:pageSize", newSize);
  emit("size-change", newSize);

  // Reset to page 1 when size changes
  emit("update:currentPage", 1);
  emit("change", 1, newSize);
};

const handleJumperChange = (e: Event) => {
  const target = e.target as HTMLInputElement;
  let newPage = Number(target.value);

  if (newPage < 1) newPage = 1;
  if (newPage > pageCount.value) newPage = pageCount.value;

  if (newPage !== props.currentPage) {
    emit("update:currentPage", newPage);
    emit("current-change", newPage);
    emit("change", newPage, props.pageSize);
  }
};
</script>

<style scoped>
.ds-pagination {
  display: flex;
  align-items: center;
  gap: 16px;
  flex-wrap: wrap;
}

.ds-pagination-total {
  font-size: 14px;
  color: #64748b;
}

.ds-pagination-size-changer select {
  height: 32px;
  padding: 0 28px 0 12px;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  background-color: #ffffff;
  font-size: 14px;
  color: #1e1b4b;
  cursor: pointer;
  transition: all 0.2s ease;
  appearance: none;
  background-image: url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='%2364748B'%3E%3Cpath d='M7 10l5 5 5-5z'/%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 8px center;
  background-size: 16px;
}

.ds-pagination-size-changer select:hover {
  border-color: #DB2777;
}

.ds-pagination-pager {
  display: flex;
  align-items: center;
  gap: 4px;
}

.ds-pagination-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  min-width: 32px;
  height: 32px;
  padding: 0 8px;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  background-color: #ffffff;
  font-size: 14px;
  color: #1e1b4b;
  cursor: pointer;
  transition: all 0.2s ease;
  user-select: none;
}

.ds-pagination-btn:hover:not(:disabled) {
  border-color: #DB2777;
  color: #DB2777;
}

.ds-pagination-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.ds-pagination-btn.is-active {
  background-color: #DB2777;
  border-color: #DB2777;
  color: #ffffff;
}

.ds-pagination-ellipsis {
  border: none;
  background: transparent;
}

.ds-pagination-jumper {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 14px;
  color: #64748b;
}

.ds-pagination-jumper input {
  width: 50px;
  height: 32px;
  padding: 0 8px;
  border: 1px solid #e2e8f0;
  border-radius: 8px;
  font-size: 14px;
  text-align: center;
  color: #1e1b4b;
  transition: all 0.2s ease;
}

.ds-pagination-jumper input:hover {
  border-color: #DB2777;
}

.ds-pagination-jumper input:focus {
  border-color: #DB2777;
  box-shadow: 0 0 0 3px rgba(219, 39, 119, 0.1);
  outline: none;
}

/* Mobile */
@media (max-width: 768px) {
  .ds-pagination {
    justify-content: center;
    gap: 12px;
  }

  .ds-pagination-total {
    width: 100%;
    text-align: center;
    margin-bottom: 4px;
  }

  .ds-pagination-size-changer select {
    height: 28px;
    font-size: 12px;
    padding-left: 8px;
  }

  .ds-pagination-btn {
    min-width: 28px;
    height: 28px;
    font-size: 12px;
  }

  .ds-pagination-jumper {
    display: none;
  }
}
</style>
