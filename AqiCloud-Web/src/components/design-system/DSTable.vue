<template>
  <div class="ds-table-wrapper">
    <table class="ds-table">
      <thead>
        <tr>
          <th v-if="selectable" class="checkbox-col">
            <input 
              type="checkbox" 
              :checked="isAllSelected"
              @change="toggleSelectAll"
            />
          </th>
          <th 
            v-for="col in columns" 
            :key="col.key"
            :class="['col-' + col.key, { sortable: col.sortable }]"
            @click="col.sortable && sort(col.key)"
          >
            {{ col.label }}
            <span v-if="col.sortable" class="sort-icon">
              {{ sortKey === col.key ? (sortOrder === 'asc' ? '↑' : '↓') : '↕' }}
            </span>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr 
          v-for="(row, index) in sortedData" 
          :key="row.id || index"
          :class="{ selected: row.selected }"
        >
          <td v-if="selectable" class="checkbox-col">
            <input 
              type="checkbox" 
              v-model="row.selected"
            />
          </td>
          <td v-for="col in columns" :key="col.key">
            <slot :name="'cell-' + col.key" :row="row" :value="row[col.key]">
              {{ row[col.key] }}
            </slot>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';

interface Column {
  key: string;
  label: string;
  sortable?: boolean;
}

interface Props {
  columns: Column[];
  data: Record<string, any>[];
  selectable?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  selectable: false
});

const emit = defineEmits<{
  'update:data': [data: Record<string, any>[]];
  'row-click': [row: Record<string, any>];
}>();

const sortKey = ref('');
const sortOrder = ref<'asc' | 'desc'>('asc');

const isAllSelected = computed(() => {
  return props.data.length > 0 && props.data.every(row => row.selected);
});

const sortedData = computed(() => {
  if (!sortKey.value) return props.data;
  
  return [...props.data].sort((a, b) => {
    const aVal = a[sortKey.value];
    const bVal = b[sortKey.value];
    const modifier = sortOrder.value === 'asc' ? 1 : -1;
    return aVal > bVal ? modifier : -modifier;
  });
});

const sort = (key: string) => {
  if (sortKey.value === key) {
    sortOrder.value = sortOrder.value === 'asc' ? 'desc' : 'asc';
  } else {
    sortKey.value = key;
    sortOrder.value = 'asc';
  }
};

const toggleSelectAll = (event: Event) => {
  const checked = (event.target as HTMLInputElement).checked;
  const updatedData = props.data.map(row => ({ ...row, selected: checked }));
  emit('update:data', updatedData);
};
</script>

<style scoped>
.ds-table-wrapper {
  background: white;
  border: 1px solid var(--color-gray-200, #E5E7EB);
  border-radius: 8px;
  overflow: hidden;
}

.ds-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 14px;
}

.ds-table thead tr {
  background: var(--color-gray-50, #F9FAFB);
  border-bottom: 1px solid var(--color-gray-200, #E5E7EB);
}

.ds-table th {
  padding: 12px 16px;
  text-align: left;
  font-weight: 500;
  font-size: 13px;
  color: var(--color-gray-500, #6B7280);
}

.ds-table th.sortable {
  cursor: pointer;
  user-select: none;
}

.ds-table th.sortable:hover {
  color: var(--color-gray-900, #111827);
}

.sort-icon {
  margin-left: 4px;
  font-size: 12px;
}

.ds-table td {
  padding: 12px 16px;
  color: var(--color-gray-900, #111827);
  border-bottom: 1px solid var(--color-gray-100, #F3F4F6);
}

.ds-table tbody tr {
  transition: background-color 0.15s ease;
}

.ds-table tbody tr:hover {
  background: var(--color-gray-50, #F9FAFB);
}

.ds-table tbody tr.selected {
  background: var(--color-primary-light, #EFF6FF);
}

.checkbox-col {
  width: 48px;
  text-align: center;
}
</style>
