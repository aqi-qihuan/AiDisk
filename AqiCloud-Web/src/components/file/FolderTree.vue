<template>
  <div class="folder-tree">
    <el-tree
      ref="treeRef"
      :data="treeData"
      :load="loadNode"
      @node-click="handleNodeClick"
      :highlight-current="true"
      node-key="id"
      :props="defaultProps"
    >
      <template #default="{ node }">
        <span class="custom-tree-node">
          <el-icon><Folder /></el-icon>
          <span>{{ node.label }}</span>
        </span>
      </template>
    </el-tree>
  </div>
</template>

<script setup lang="ts">
import { getFileTree } from "@/api/file";
import { Folder } from "@element-plus/icons-vue";
import { ElTree } from "element-plus";
import { onMounted, ref } from "vue";
import type { TreeNodeDTO } from "@/api/types";

/**
 * FolderTree 组件 - 文件夹树组件
 * 提供文件夹树形结构展示和选择功能
 */

interface TreeNode {
  level: number;
  data?: any;
}

const emit = defineEmits<{
  select: [node: any];
}>();

const treeRef = ref<InstanceType<typeof ElTree>>();
const treeData = ref<TreeNodeDTO[]>([]);

/**
 * 树节点的配置属性
 */
const defaultProps = {
  children: "children",
  label: "label",
  isLeaf: (data: TreeNodeDTO) => data.state === "closed",
};

/**
 * 处理节点点击事件
 * @param data - 被点击的节点数据
 * @param node - 节点实例
 */
const handleNodeClick = (data: TreeNodeDTO, node: any): void => {
  emit("select", {
    ...data,
    path: data.label,
  });
};

/**
 * 加载节点数据
 * @param node - 当前节点实例
 * @param resolve - 加载完成的回调函数
 */
const loadNode = async (
  node: TreeNode,
  resolve: (data: TreeNodeDTO[]) => void,
): Promise<void> => {
  if (node.level === 0) {
    try {
      const response = await getFileTree();

      if (response.data?.success && response.data.data) {
        treeData.value = response.data.data || [];
        resolve(treeData.value);
      } else {
        resolve([]);
      }
    } catch (error) {
      console.error("获取文件树失败:", error);
      resolve([]);
    }
  } else {
    // 子节点直接使用现有的 children 数据
    resolve(node.data?.children || []);
  }
};

/**
 * 组件挂载时初始化文件树
 */
onMounted(() => {
  loadNode({ level: 0 }, (data) => {
    treeData.value = data;
  });
});
</script>

<style scoped>
.folder-tree {
  max-height: 400px;
  overflow-y: auto;
  border-radius: 4px;
}

.custom-tree-node {
  display: flex;
  align-items: center;
  font-size: 14px;
  padding: 4px 0;
}

.custom-tree-node .el-icon {
  margin-right: 6px;
  font-size: 16px;
  color: var(--color-text-tertiary, #64748B);
}

::v-deep(.el-tree-node__content) {
  height: 36px;
  transition: background-color 0.2s;
}

::v-deep(.el-tree-node__content:hover) {
  background-color: rgba(219, 39, 119, 0.08);
}

::v-deep(.el-tree-node.is-current > .el-tree-node__content) {
  background-color: #ecf5ff;
  color: #409eff;
}

::v-deep(.el-tree-node.is-current .custom-tree-node .el-icon) {
  color: #409eff;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .custom-tree-node {
    font-size: 13px;
  }

  ::v-deep(.el-tree-node__content) {
    height: 32px;
  }
}
</style>
