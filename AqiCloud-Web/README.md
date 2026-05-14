# AqiCloud-AiPan - 小七云盘

> 智能文件管理系统 - 融合 AI 技术的新一代云存储解决方案

[![Vue 3](https://img.shields.io/badge/Vue-3.5.34-42b883.svg)](https://vuejs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-6.0.3-3178c6.svg)](https://www.typescriptlang.org/)
[![Vite](https://img.shields.io/badge/Vite-8.0.12-646CFF.svg)](https://vitejs.dev/)
[![Element Plus](https://img.shields.io/badge/Element%20Plus-2.14.0-409eff.svg)](https://element-plus.org/)
[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

---

## 📖 项目简介

AqiCloud-AiPan 是一款面向个人和企业的智能文件管理系统，融合 AI 技术提供文件分类、内容识别、智能问答等功能。基于 Vue 3、TypeScript、Vite 和 Element Plus 构建，提供流畅的用户体验和现代化的界面设计。

### 核心功能

- 📁 **文件管理** - 支持拖拽上传、批量上传、文件夹管理、文件预览
- 🔗 **文件分享** - 一键分享，支持加密和权限控制
- 🤖 **AI 赋能** - 智能文件分类、内容识别、AI 问答、AIGC 文案生成
- 🔒 **安全存储** - 企业级加密技术，确保数据安全
- 📱 **响应式设计** - 完美适配桌面端和移动端
- 🎨 **现代设计** - 基于设计系统的精美 UI，支持亮色/暗色模式

---

## 🚀 快速开始

### 环境要求

| 依赖 | 版本要求 | 推荐版本 |
|------|----------|----------|
| **Node.js** | ≥ 20.19.0 或 ≥ 22.12.0 | 22.x LTS |
| **npm** | ≥ 9.0.0 | 10.x |
| **Git** | ≥ 2.30.0 | 最新稳定版 |

> ⚠️ **注意**：本项目使用 Vite 8.x（基于 Rolldown），需要 Node.js 20.19+ 或 22.12+。

### 安装步骤

#### 1. 克隆项目

```bash
git clone https://github.com/yourusername/AqiCloud-AiPan.git
cd AqiCloud-AiPan
```

#### 2. 安装依赖

```bash
npm install
```

或使用 pnpm（推荐，更快）：

```bash
pnpm install
```

#### 3. 配置环境变量

复制环境变量示例文件：

```bash
cp .env.development.example .env.development
cp .env.production.example .env.production
```

根据需要编辑 `.env.development`：

```env
# 开发环境配置
VITE_API_BASE_URL=http://127.0.0.1:9090/api
```

#### 4. 启动开发服务器

```bash
npm run dev
```

应用将在 `http://localhost:8081` 启动。

#### 5. 构建生产版本

```bash
npm run build
```

构建输出将生成在 `dist/` 目录。

---

## 💡 使用示例

### 基本用法

#### 1. 文件上传示例

```vue
<template>
  <div>
    <input type="file" @change="handleFileUpload" />
    <button @click="uploadFile">上传</button>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { uploadFile } from '@/api/file';

const selectedFile = ref<File | null>(null);

const handleFileUpload = (event: Event) => {
  const target = event.target as HTMLInputElement;
  if (target.files && target.files[0]) {
    selectedFile.value = target.files[0];
  }
};

const uploadFile = async () => {
  if (!selectedFile.value) return;
  
  try {
    const response = await uploadFile(selectedFile.value);
    console.log('上传成功:', response);
  } catch (error) {
    console.error('上传失败:', error);
  }
};
</script>
```

#### 2. API 调用示例

```typescript
// src/api/file.ts
import myAxios from '@/utils/request';

// 获取文件列表
export const getFileList = async (params: {
  page: number;
  pageSize: number;
  parentId?: string;
}) => {
  return myAxios({
    url: '/file/v1/list',
    method: 'GET',
    params,
  });
};

// 上传文件
export const uploadFile = async (file: File) => {
  const formData = new FormData();
  formData.append('file', file);
  
  return myAxios({
    url: '/file/v1/upload',
    method: 'POST',
    data: formData,
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  });
};
```

#### 3. 使用设计系统组件

```vue
<template>
  <div class="p-md">
    <h1 class="text-2xl font-bold text-primary mb-md">我的文件</h1>
    
    <DSButton variant="primary" size="md" @click="handleUpload">
      上传文件
    </DSButton>
    
    <div class="mt-lg">
      <DSInput 
        v-model="searchQuery" 
        placeholder="搜索文件..." 
        class="w-full"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { DSButton, DSInput } from '@/components/design-system';

const searchQuery = ref('');
const handleUpload = () => {
  console.log('上传文件');
};
</script>

<style scoped>
/* 使用设计系统 CSS 变量 */
.custom-style {
  background-color: var(--color-surface);
  border-radius: var(--radius-md);
  padding: var(--spacing-md);
  box-shadow: var(--shadow-md);
  transition: all var(--transition-base);
}
</style>
```

### 高级用法

#### 1. 使用 Pinia 状态管理

```typescript
// src/store/file.ts
import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import { getFileList } from '@/api/file';

export const useFileStore = defineStore('file', () => {
  const fileList = ref<any[]>([]);
  const loading = ref(false);
  const pagination = ref({
    page: 1,
    pageSize: 20,
    total: 0,
  });

  const fetchFiles = async () => {
    loading.value = true;
    try {
      const res = await getFileList({
        page: pagination.value.page,
        pageSize: pagination.value.pageSize,
      });
      if (res.success) {
        fileList.value = res.data.records;
        pagination.value.total = res.data.total;
      }
    } finally {
      loading.value = false;
    }
  };

  return {
    fileList,
    loading,
    pagination,
    fetchFiles,
  };
});
```

#### 2. 使用 Vue Router

```typescript
// src/router/routes.ts
import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: () => import('@/views/HomeView.vue'),
    },
    {
      path: '/file',
      component: () => import('@/views/file/FileView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/share/:shareId',
      component: () => import('@/views/ShareView.vue'),
    },
  ],
});

// 路由守卫
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token');
  
  if (to.meta.requiresAuth && !token) {
    next('/user/login');
  } else {
    next();
  }
});

export default router;
```

---

## 🎨 技术栈

### 前端框架
- **Vue 3.5.34** - 渐进式 JavaScript 框架（Composition API）
- **TypeScript 6.0.3** - 类型安全的 JavaScript 超集
- **Vue Router 5.0.7** - Vue.js 官方路由管理器
- **Pinia 3.0.4** - Vue 3 官方状态管理库

### UI 组件库
- **Element Plus 2.14.0** - Vue 3 组件库
- **Lucide Vue Next 1.0.0** - 图标库
- **自定义设计系统** - 统一的 UI 组件库

### 构建工具
- **Vite 8.0.12** - 新一代前端构建工具（基于 Rolldown，构建速度提升 6.6x）
- **@vitejs/plugin-vue 6.0.6** - Vue 3 支持
- **@vitejs/plugin-vue-jsx 5.1.5** - JSX 支持

### 工具库
- **Axios 1.16.0** - HTTP 客户端
- **Spark MD5 3.0.2** - 文件 MD5 计算（秒传功能）
- **JSON BigInt 1.0.0** - 大整数处理
- **Vue i18n 11.4.2** - 国际化

### 开发工具
- **TypeScript** - 类型检查
- **Prettier 3.8.3** - 代码格式化
- **ESLint** - 代码质量检查（可选）

---

## 📁 项目结构

```
AqiCloud-AiPan/
├── public/                     # 静态资源
│   ├── favicon.ico            # 网站图标
│   └── logo.svg              # 网站 Logo
├── src/
│   ├── access/               # 访问控制
│   │   ├── accessEnum.ts     # 权限枚举
│   │   ├── checkAccess.ts    # 权限检查
│   │   └── index.ts         # 导出
│   ├── api/                  # API 接口
│   │   ├── file.ts           # 文件相关 API
│   │   ├── user.ts           # 用户相关 API
│   │   ├── share.ts          # 分享相关 API
│   │   └── index.ts         # 统一导出
│   ├── assets/               # 资源文件（图片、字体等）
│   ├── components/           # 组件
│   │   ├── common/           # 通用组件
│   │   │   ├── ImageUpload.vue
│   │   │   └── LoadingDots.vue
│   │   ├── file/             # 文件相关组件
│   │   │   ├── FileTable.vue
│   │   │   ├── FileGrid.vue
│   │   │   ├── FileUpload.vue
│   │   │   └── ...
│   │   └── design-system/    # 设计系统组件
│   │       ├── DSButton.vue
│   │       ├── DSInput.vue
│   │       └── ...
│   ├── layouts/              # 布局组件
│   │   ├── BasicLayout.vue   # 主布局
│   │   └── UserLayout.vue   # 用户布局
│   ├── locales/              # 国际化
│   ├── router/               # 路由配置
│   │   └── routes.ts
│   ├── store/                # 状态管理
│   │   ├── file.ts
│   │   ├── user.ts
│   │   └── theme.ts
│   ├── styles/               # 全局样式
│   │   ├── design-system.css  # 设计系统样式
│   │   ├── design-tokens.css  # 设计令牌
│   │   ├── global-reset.css   # 全局样式重置
│   │   └── mobile.css         # 移动端样式
│   ├── utils/                # 工具函数
│   │   ├── request.ts        # Axios 实例
│   │   ├── format.ts        # 格式化工具
│   │   └── md5.ts           # MD5 工具
│   ├── views/                # 页面组件
│   │   ├── file/             # 文件管理页面
│   │   ├── user/             # 用户相关页面
│   │   ├── admin/            # 管理后台页面
│   │   └── error/            # 错误页面
│   ├── App.vue               # 根组件
│   └── main.ts               # 入口文件
├── .env.development          # 开发环境变量
├── .env.production           # 生产环境变量
├── index.html                # HTML 模板
├── package.json              # 项目配置
├── tsconfig.json             # TypeScript 配置
├── vite.config.ts           # Vite 配置
└── README.md                # 项目说明
```

---

## 🔧 开发指南

### 开发命令

| 命令 | 说明 |
|------|------|
| `npm run dev` | 启动开发服务器（端口 8081） |
| `npm run build` | 构建生产版本 |
| `npm run preview` | 预览生产版本 |

### API 配置

本项目使用 Vite 代理来解决跨域问题：

**开发环境**：
- 前端开发服务器：`http://localhost:8081`
- API 请求通过 Vite 代理到后端：`http://127.0.0.1:9090`
- Axios 配置：`baseURL: "/api"`

**生产环境**：
- 使用 Nginx 反向代理
- API 请求路径：`/api`

### 环境变量

| 变量名 | 说明 | 默认值 |
|--------|------|--------|
| `VITE_API_BASE_URL` | API 基础路径 | `/api` |

### 设计系统

本项目采用完整的设计系统，确保 UI/UX 的一致性和专业性。

**设计特点**：
- **风格**：Flat Design（扁平化设计）
- **主色调**：Indigo (#6366F1) + Emerald (#10B981)
- **字体**：Plus Jakarta Sans
- **响应式**：移动优先，完美适配所有设备

**CSS 变量系统**：

```css
/* 颜色系统 */
--color-primary: #6366F1;      /* 主色 */
--color-secondary: #818CF8;    /* 辅助色 */
--color-cta: #10B981;          /* 行动号召色 */
--color-bg: #F5F3FF;           /* 背景色 */
--color-surface: #FFFFFF;      /* 表面色 */

/* 间距系统 */
--spacing-xs: 4px;
--spacing-sm: 8px;
--spacing-md: 16px;
--spacing-lg: 24px;
--spacing-xl: 32px;

/* 圆角系统 */
--radius-sm: 4px;
--radius-md: 8px;
--radius-lg: 12px;
--radius-xl: 16px;
```

### 代码规范

#### TypeScript 规范

```typescript
/**
 * 组件 Props 接口定义
 */
interface ComponentProps {
  title: string;
  count?: number;
}

/**
 * 组件 Emits 类型定义
 */
const emit = defineEmits<{
  "update:modelValue": [value: string];
  "submit": [];
}>();
```

#### 提交规范

```
feat: 新功能
fix: 修复 bug
docs: 文档更新
style: 代码格式调整（不影响功能）
refactor: 重构
perf: 性能优化
test: 测试相关
chore: 构建/工具相关
```

---

## 🌐 核心页面

### 公开页面
- `/` - 首页
- `/user/login` - 用户登录
- `/user/register` - 用户注册
- `/share/:shareId` - 分享访问

### 文件管理（需要登录）
- `/file` - 文件列表
- `/picture` - 图片管理
- `/share` - 我的分享
- `/recycle` - 回收站
- `/search` - 文件搜索

### AI 功能
- `/chat` - AI 聊天助理
- `/document` - AI 文档助手
- `/answer` - AI 网盘智答
- `/grow` - AIGC 文案智能体

### 其他
- `/about` - 关于我们
- `/admin/user` - 用户管理（管理员）

---

## 📱 移动端适配

本项目采用移动优先的响应式设计：

### 响应式断点

| 设备类型 | 断点 | Header 高度 | 侧边栏行为 |
|---------|------|--------------|--------------|
| Desktop | > 1024px | 60px | 固定显示，可折叠 |
| Tablet | 768-1024px | 60px | 固定显示，可折叠 |
| Mobile | 576-768px | 56px | 抽屉式，带遮罩 |
| Small Mobile | < 576px | 52px | 抽屉式，带遮罩 |

### 移动端特性

- ✅ 完美适配 iPhone SE (375x667)
- ✅ 支持平板设备 (768px+)
- ✅ 优化的触摸交互
- ✅ 移动端专用样式 (`src/styles/mobile.css`)
- ✅ 抽屉式侧边栏
- ✅ 响应式导航菜单

---

## 🌈 暗色模式

设计系统内置暗色模式支持：

```vue
<script setup lang="ts">
import { ref, onMounted } from 'vue';

const isDark = ref(false);

const toggleTheme = () => {
  isDark.value = !isDark.value;
  document.documentElement.setAttribute('data-theme', isDark.value ? 'dark' : 'light');
};

onMounted(() => {
  const savedTheme = localStorage.getItem('theme');
  if (savedTheme) {
    isDark.value = savedTheme === 'dark';
  }
});
</script>
```

---

## 🤝 贡献指南

我们欢迎任何形式的贡献，包括但不限于：

- 🐛 报告 Bug
- 💡 提出新功能建议
- 📝 改进文档
- 💻 提交代码

### 贡献步骤

1. **Fork 本仓库**

   ```bash
   # 点击 GitHub 上的 Fork 按钮
   ```

2. **创建特性分支**

   ```bash
   git checkout -b feature/amazing-feature
   ```

3. **提交更改**

   ```bash
   git commit -m 'feat: Add some amazing feature'
   ```

4. **推送到分支**

   ```bash
   git push origin feature/amazing-feature
   ```

5. **创建 Pull Request**

   前往 GitHub 仓库页面，点击 "New Pull Request"

### 贡献要求

- ✅ **遵循设计系统规范** - 使用设计系统的 CSS 变量和组件
- ✅ **使用 TypeScript 类型定义** - 所有新代码必须有完整的类型定义
- ✅ **编写清晰的提交信息** - 遵循 Conventional Commits 规范
- ✅ **添加必要的测试** - 重要功能需要添加单元测试
- ✅ **更新相关文档** - 新功能需要更新文档
- ✅ **移除 console.log 调试代码** - 提交前清理调试代码
- ✅ **遵循代码风格** - 使用 Prettier 格式化代码

### 代码审查标准

| 检查项 | 要求 |
|--------|------|
| TypeScript 覆盖率 | 100% |
| 设计系统集成 | 使用 CSS 变量，不硬编码颜色/间距 |
| 代码质量 | 通过 ESLint 检查，无警告 |
| 响应式布局 | 在 4 个断点测试通过 |
| 可访问性 | 符合 WCAG 2.1 AA 标准 |

### 报告 Bug

请在 [GitHub Issues](https://github.com/yourusername/AqiCloud-AiPan/issues) 中报告 Bug，并包含以下信息：

- 📝 **Bug 描述** - 清晰描述遇到的问题
- 🔄 **复现步骤** - 详细的复现步骤
- 📷 **截图** - 如果可能，提供截图
- 🖥️ **环境信息** - 操作系统、浏览器、屏幕尺寸等
- 📱 **移动端** - 是否在移动端出现

### 提出新功能建议

请在 [GitHub Issues](https://github.com/yourusername/AqiCloud-AiPan/issues) 中提出新功能建议，并包含：

- 💡 **功能描述** - 清晰描述建议的功能
- 🎯 **使用场景** - 为什么需要这个功能
- 📐 **设计方案** - 如果可能，提供设计方案或草图
- 📝 **替代方案** - 考虑过的其他方案

---

## 📄 许可证

本项目采用 MIT 许可证。

详见 [LICENSE](LICENSE) 文件（如果不存在，请在根目录创建）。

---

## 📞 联系方式

- **项目主页**: [aqi-qihuan/AiDisk](https://github.com/aqi-qihuan/AiDisk))
- **问题反馈**: [GitHub Issues](https://github.com/yourusername/AqiCloud-AiPan/issues)
- **邮箱**: 2316364297@qq.com

---

## 🙏 致谢

感谢以下开源项目：

- [Vue.js](https://vuejs.org/)
- [Element Plus](https://element-plus.org/)
- [Pinia](https://pinia.vuejs.org/)
- [Vue Router](https://router.vuejs.org/)
- [TypeScript](https://www.typescriptlang.org/)
- [Vite](https://vitejs.dev/)
- [Axios](https://axios-http.com/)
- [Lucide](https://lucide.dev/)

---

## 📝 更新日志

### v0.1.0 (2026-05-14) - 初始版本

#### ✨ 新特性
- 🎨 完整的设计系统
- 📱 移动端适配
- 🤖 AI 功能集成
- 🧪 自动化测试
- 📚 完整文档

---

**AqiCloud-AiPan** - 让文件管理更简单，协作更高效 🚀
