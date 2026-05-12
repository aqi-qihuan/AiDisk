# 🎨 AqiCloud-AiPan 全新 UI 设计方案

> **版本**: v2.0.0 | **日期**: 2026-05-11 | **风格**: Glassmorphism + Rose Pink Core

---

## 📋 目录

- [1. 设计语言与品牌规范](#1-设计语言与品牌规范)
- [2. 色彩系统](#2-色彩系统)
- [3. 字体排印系统](#3-字体排印系统)
- [4. 间距与布局系统](#4-间距与布局系统)
- [5. 玻璃态设计语言](#5-玻璃态设计语言)
- [6. 图标与插图风格](#6-图标与插图风格)
- [7. 组件库架构](#7-组件库架构)
- [8. 核心页面布局](#8-核心页面布局)
- [9. 用户流程设计](#9-用户流程设计)
- [10. 交互与动效规范](#10-交互与动效规范)
- [11. 响应式策略](#11-响应式策略)
- [12. 无障碍设计](#12-无障碍设计)
- [13. 暗色模式](#13-暗色模式)
- [14. 设计交付与开发协作](#14-设计交付与开发协作)
- [15. 迁移路线图](#15-迁移路线图)

---

## 1. 设计语言与品牌规范

### 1.1 设计哲学

```
纯净 · 智能 · 高效
```

| 维度 | 描述 |
|------|------|
| **风格** | 扁平化基础上的轻玻璃态 (Light Glassmorphism) |
| **气质** | 科技感 + 亲和力，专业而不冰冷 |
| **核心视觉** | 粉色主调 + 金色点缀 + 毛玻璃效果 |
| **设计原则** | 一致性、清晰性、反馈性、容错性 |

### 1.2 品牌色板

- **主色 (Primary)**: #DB2777 — 代表活力与亲和
- **点缀色 (CTA)**: #D97706 — 代表行动与温暖
- **辅助色 (Secondary)**: #F472B6 — 主色的柔和延伸

### 1.3 标语与应用名称

| 场景 | 内容 |
|------|------|
| 产品名 | 小七云盘 (AqiCloud) |
| 英文标语 | Smart Cloud, Smarter You |
| 中文标语 | 让文件管理更智能 |

---

## 2. 色彩系统

### 2.1 主色板

```css
/* Rose Pink Palette */
--pink-50:  #FDF2F8;
--pink-100: #FCE7F3;
--pink-200: #FBCFE8;
--pink-300: #F9A8D4;
--pink-400: #F472B6;   /* Secondary */
--pink-500: #EC4899;
--pink-600: #DB2777;   /* Primary */
--pink-700: #BE185D;
--pink-800: #9D174D;
--pink-900: #831843;
```

### 2.2 强调色板

```css
/* Amber CTA Palette */
--amber-50:  #FFFBEB;
--amber-100: #FEF3C7;
--amber-200: #FDE68A;
--amber-300: #FCD34D;
--amber-400: #FBBF24;
--amber-500: #D97706;   /* CTA */
--amber-600: #B45309;
--amber-700: #92400E;
```

### 2.3 中性色板

```css
/* Neutral Palette */
--neutral-50:  #FAFAFA;
--neutral-100: #F4F4F5;
--neutral-200: #E4E4E7;
--neutral-300: #D4D4D8;
--neutral-400: #A1A1AA;
--neutral-500: #71717A;
--neutral-600: #52525B;
--neutral-700: #3F3F46;
--neutral-800: #27272A;
--neutral-900: #18181B;
```

### 2.4 功能色板

```css
/* Functional Colors */
--color-success: #10B981;     /* 成功 - Emerald */
--color-warning: #F59E0B;     /* 警告 - Amber */
--color-error: #EF4444;       /* 错误 - Red */
--color-info: #3B82F6;        /* 信息 - Blue */
```

### 2.5 语义色值对照表

| 用途 | CSS 变量 | 亮色值 | 暗色值 | 用途说明 |
|------|----------|--------|--------|----------|
| 页面背景 | `--color-bg` | `#FDF2F8` | `#0F172A` | 主背景色 |
| 卡片表面 | `--color-surface` | `#FFFFFF` | `#1E293B` | 卡片、表单背景 |
| 主文本 | `--color-text-primary` | `#831843` | `#F8FAFC` | 标题、正文 |
| 副文本 | `--color-text-secondary` | `#64748B` | `#94A3B8` | 说明文字 |
| 弱文本 | `--color-text-tertiary` | `#94A3B8` | `#64748B` | 占位符、禁用文字 |
| 边框 | `--color-border` | `#E2E8F0` | `#334155` | 分割线、边框 |

### 2.6 WCAG 对比度合规

| 组合 | 对比度 | 等级 | 适用场景 |
|------|--------|------|----------|
| Primary 上白色 | 6.5:1 | ✅ AA | 主按钮文字 |
| 主文本上白色背景 | 10.8:1 | ✅ AAA | 正文阅读 |
| 副文本上白色背景 | 4.6:1 | ✅ AA | 辅助说明 |
| Primary 上 Pink-100 | 3.5:1 | ❌ 不足 | 禁用状态 |

> **规则**: 所有可交互元素和正文文本必须满足 WCAG AA (4.5:1) 标准。

---

## 3. 字体排印系统

### 3.1 字体家族

```css
/* Primary Font (Body) */
--font-primary: 'Fira Sans', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;

/* Heading Font */
--font-heading: 'Fira Code', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;

/* Monospace */
--font-mono: 'Fira Code', 'Courier New', monospace;
```

> **设计理由**: Fira Sans 具有优秀的可读性和现代感，Fira Code 在标题上赋予独特的科技感。两者同属 Fira 家族，和谐统一。

### 3.2 字体层级

| 层级 | Size | Weight | Line Height | Letter Spacing | 用途 |
|------|------|--------|-------------|----------------|------|
| **Display 1** | 48px | 700 (Bold) | 1.1 | -0.03em | Hero 标题 |
| **Display 2** | 36px | 700 (Bold) | 1.15 | -0.02em | 页面大标题 |
| **Heading 1** | 24px | 600 (Semibold) | 1.25 | -0.02em | 区块标题 |
| **Heading 2** | 20px | 600 (Semibold) | 1.3 | -0.01em | 卡片标题 |
| **Heading 3** | 18px | 600 (Semibold) | 1.35 | normal | 子标题 |
| **Body Large** | 16px | 400 (Regular) | 1.6 | normal | 正文 |
| **Body** | 14px | 400 (Regular) | 1.6 | normal | 次要正文 |
| **Small** | 12px | 400 (Regular) | 1.5 | normal | 辅助文字、标签 |
| **Caption** | 11px | 500 (Medium) | 1.4 | 0.02em | 数据标注、时间戳 |

### 3.3 CSS 变量映射

```css
--text-xs: 11px;
--text-sm: 12px;
--text-base: 14px;     /* 以14px为基准 */
--text-lg: 16px;
--text-xl: 18px;
--text-2xl: 20px;
--text-3xl: 24px;
--text-4xl: 36px;
--text-5xl: 48px;
```

---

## 4. 间距与布局系统

### 4.1 8px 网格系统

采用 **8px 基准网格** (base unit = 8px)，所有间距、边距、组件尺寸均以 8 的倍数递增。

```css
--space-1: 2px;    /* 微间距 */
--space-2: 4px;    /* 密集间距 */
--space-3: 8px;    /* 基准单位 */
--space-4: 12px;   /* 紧凑间距 */
--space-5: 16px;   /* 标准间距 */
--space-6: 20px;   /* 舒适间距 */
--space-7: 24px;   /* 组件间距 */
--space-8: 32px;   /* 区块间距 */
--space-9: 40px;   /* 大区块间距 */
--space-10: 48px;  /* 页面间距 */
--space-11: 64px;  /* 大页面间距 */
--space-12: 80px;  /* Hero 间距 */
```

> **注意**: 此系统将替代现有的 `--spacing-xs/sm/md/lg/xl/xxl` 命名，提供更细粒度的控制。过渡期两者并存。

### 4.2 布局网格

```
Desktop (>1024px): 12列网格, gutter 24px
Tablet (768-1024px): 8列网格, gutter 20px
Mobile (<768px): 4列网格, gutter 16px
```

### 4.3 容器宽度

| 断点 | 容器最大宽度 | 侧边距 |
|------|-------------|--------|
| Mobile | 100% | 16px |
| Tablet | 100% | 24px |
| Desktop | 1200px | 32px |
| Large | 1440px | auto |

---

## 5. 玻璃态设计语言

### 5.1 玻璃层级体系

| 层级 | 背景透明度 | 模糊程度 | 边框 | 阴影 | 用途 |
|------|-----------|----------|------|------|------|
| **Subtle** | rgba(255,255,255,0.6) | blur(10px) | 1px solid rgba(255,255,255,0.4) | 无 | 次要卡片 |
| **Default** | rgba(255,255,255,0.8) | blur(20px) saturate(180%) | 1px solid rgba(255,255,255,0.6) | `--glass-shadow` | 主要卡片、导航 |
| **Strong** | rgba(255,255,255,0.9) | blur(30px) saturate(200%) | 1px solid rgba(255,255,255,0.8) | `--glass-shadow` + y:4px | 模态框、菜单 |
| **Primary Tint** | rgba(219,39,119,0.08) | blur(20px) | 1px solid rgba(219,39,119,0.2) | `--glass-shadow` | 选中状态、特色卡片 |

### 5.2 圆角体系

```css
--radius-sm: 4px;     /* 标签、小元素 */
--radius-md: 8px;     /* 按钮、输入框 */
--radius-lg: 12px;    /* 卡片、对话框 */
--radius-xl: 16px;    /* 大卡片、模态框 */
--radius-2xl: 24px;   /* 大面板 */
--radius-full: 9999px; /* 胶囊、头像 */
```

### 5.3 阴影体系

```css
--shadow-sm: 0 1px 2px rgba(0,0,0,0.05);
--shadow-md: 0 4px 6px rgba(0,0,0,0.07), 0 2px 4px rgba(0,0,0,0.04);
--shadow-lg: 0 10px 15px rgba(0,0,0,0.08), 0 4px 6px rgba(0,0,0,0.04);
--shadow-xl: 0 20px 25px rgba(0,0,0,0.1), 0 10px 10px rgba(0,0,0,0.04);
  --shadow-glow: 0 0 20px rgba(219,39,119,0.15); /* 粉色发光 */
  --shadow-glow-cta: 0 0 20px rgba(217,119,6,0.2); /* 金色发光 */
```

---

## 6. 图标与插图风格

### 6.1 图标系统

| 属性 | 规格 |
|------|------|
| **风格** | 线性 (Outline) + 圆角端点 |
| **线条粗细** | 1.5px (小图标) / 2px (大图标) |
| **尺寸规格** | 16px / 20px / 24px / 32px / 48px |
| **颜色** | 继承 currentColor |
| **库选择** | Lucide Icons (已在 package.json 中引入 `lucide-vue-next`) |

### 6.2 图标引入方式

```vue
<!-- 推荐方式: 使用 Lucide 组件 -->
<script setup>
import { File, Folder, Upload, Share2, Trash2 } from 'lucide-vue-next'
</script>

<template>
  <File :size="20" class="icon-primary" />
</template>

<style scoped>
.icon-primary { color: var(--color-primary); }
</style>
```

### 6.3 文件类型图标映射

| 文件类型 | 图标 | 语义色 |
|----------|------|--------|
| 文件夹 | `Folder` | --color-primary |
| 图片 | `Image` | --color-success |
| 文档 | `FileText` | --color-info |
| 视频 | `Video` | --color-warning |
| 音频 | `Music` | --color-error (pink) |
| 压缩包 | `Archive` | --color-text-secondary |
| 代码 | `Code` | --color-primary |
| PDF | `File` | --color-error |
| 其他 | `File` | --color-text-tertiary |

---

## 7. 组件库架构

### 7.1 组件分层

```
┌──────────────────────────────────┐
│         基础原子组件              │
│  Button, Input, Tag, Card, Icon  │
├──────────────────────────────────┤
│         复合分子组件              │
│  Form, Dialog, Dropdown, Table,  │
│  Pagination, Tabs, Alert         │
├──────────────────────────────────┤
│         业务有机组件              │
│  FileTable, FileGrid, FileUpload │
│  FolderTree, BreadCrumb,         │
│  ContextMenu, OperationBar       │
├──────────────────────────────────┤
│         页面级组件                │
│  HomePage, FilePage, SharePage   │
│  LoginPage, ChatPage, etc.       │
└──────────────────────────────────┘
```

### 7.2 现有组件状态（需迁移）

| 组件 | 当前状态 | 迁移内容 | 优先级 |
|------|---------|----------|--------|
| DSButton | ✅ 已迁移 | 无需修改 | - |
| DSTag | ✅ 已迁移 | 无需修改 | - |
| DSInput | ✅ 已迁移 | 无需修改 | - |
| DSCard | ✅ 已迁移 | 无需修改 | - |
| **GlobalHeader** | ⚠️ 本地样式 | 替换为设计系统变量 | 🔴 高 |
| **HomeView** | ⚠️ 本地粉色主题 | 统一为紫色玻璃态 | 🔴 高 |
| **FileTable** | ⚠️ 部分未对齐 | 统一悬停/选中样式 | 🟡 中 |
| **Login/Register** | ⚠️ 本地样式 | 统一为设计系统 | 🔴 高 |

### 7.3 新增组件规格

#### DSDropdown (下拉菜单)

```vue
<template>
  <div class="ds-dropdown" :class="{ 'ds-dropdown--open': isOpen }">
    <slot name="trigger" :toggle="toggle" />
    <transition name="fade-slide">
      <div v-show="isOpen" class="ds-dropdown__menu glass-strong" @click="handleMenuClick">
        <slot name="menu" />
      </div>
    </transition>
  </div>
</template>

<style scoped>
.ds-dropdown__menu {
  position: absolute;
  min-width: 180px;
  padding: 4px;
  border-radius: var(--radius-lg);
  z-index: var(--z-dropdown);
}
</style>
```

#### DSModal (模态对话框)

| 属性 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| visible | boolean | false | 显示控制 |
| title | string | - | 标题 |
| width | string | '480px' | 宽度 |
| showFooter | boolean | true | 显示底部操作栏 |
| confirmText | string | '确认' | 确认按钮文本 |
| cancelText | string | '取消' | 取消按钮文本 |
| confirmLoading | boolean | false | 确认按钮加载状态 |

#### DSFileIcon (文件类型图标)

| 属性 | 类型 | 默认值 | 说明 |
|------|------|--------|------|
| fileSuffix | string | '' | 文件后缀名 |
| isFolder | boolean | false | 是否为文件夹 |
| size | 'small'\|'medium'\|'large' | 'medium' | 图标尺寸 |

#### DSFileGridItem (文件网格卡片)

文件管理区的网格视图卡片，支持:

- 文件图标 + 文件名
- 多选状态（勾选框）
- 拖动排序
- 右键菜单触发
- 文件大小/日期信息

### 7.4 组件状态设计规范

每个交互组件必须定义以下状态:

```
┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────┐
│ Default  │  │  Hover   │  │  Active  │  │ Disabled │  │Loading   │
│  正常    │→ │  悬停    │→ │  按下    │→ │  禁用    │  │  加载    │
└──────────┘  └──────────┘  └──────────┘  └──────────┘  └──────────┘

                 ┌──────────┐  ┌──────────┐
                 │ Focus    │  │  Error   │
                 │  聚焦    │  │  错误    │
                 └──────────┘  └──────────┘
```

---

## 8. 核心页面布局

### 8.1 页面布局架构

```
┌────────────────────────────────────────────────┐
│              GlobalHeader (60px)                │
│  [Logo] [导航菜单]              [语言] [用户]   │
├──────────┬─────────────────────────────────────┤
│          │                                      │
│  Sidebar │          Main Content                │
│  (240px) │                                      │
│          │  ┌────────────────────────────┐      │
│  📁 全部  │  │ OperationBar               │      │
│  🖼️ 图片  │  │ [新建][上传][下载][更多]   │      │
│  📄 文档  │  ├────────────────────────────┤      │
│  🎬 视频  │  │ BreadCrumb                 │      │
│  🎵 音乐  │  │ 全部 > 文档 > 项目         │      │
│  📦 压缩  │  ├────────────────────────────┤      │
│          │  │                              │      │
│  存储空间 │  │    FileTable / FileGrid     │      │
│  ████ 60% │  │    (文件列表/网格展示)      │      │
│          │  │                              │      │
└──────────┴──┴──────────────────────────────┘      │
```

### 8.2 页面布局分类

| 布局类别 | 适用页面 | 说明 |
|----------|---------|------|
| **Public Layout** | 首页、登录、注册、关于 | 全宽页面，无侧边栏 |
| **App Layout** | 文件管理、分享、回收站 | 含侧边栏的文件管理布局 |
| **Chat Layout** | AI 聊天、AI 文档 | 含侧边栏 + 对话面板 |
| **Admin Layout** | 用户管理 | 窄侧边栏 + 数据表格 |

### 8.3 关键页面设计

#### 8.3.1 首页 (HomePage) — 全新设计

```
┌──────────────────────────────────────────────────────┐
│                   Hero Section                        │
│  ┌────────────────────┐  ┌─────────────────────┐     │
│  │  ✨ AI Powered     │  │    Glass Card       │     │
│  │                    │  │  [插入网盘管理插图]  │     │
│  │  AI Pan 智能云盘    │  │                     │     │
│  │                    │  └─────────────────────┘     │
│  │  让文件管理更智能、  │                               │
│  │  更高效、更安全      │                               │
│  │                    │                               │
│  │  [🚀 开始使用] [📖]│                               │
│  └────────────────────┘                               │
├──────────────────────────────────────────────────────┤
│                  Features Grid                        │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌──────────┐ │
│  │  Glass   │ │  Glass   │ │  Glass   │ │  Glass   │ │
│  │  Card    │ │  Card    │ │  Card    │ │  Card    │ │
│  │ AI 问答  │ │ 文档助手  │ │ AI聊天   │ │ 智能搜索 │ │
│  └──────────┘ └──────────┘ └──────────┘ └──────────┘ │
├──────────────────────────────────────────────────────┤
│   Stats Row (渐变紫色背景 + 白色玻璃卡片)              │
│  ┌──────┐ ┌──────┐ ┌──────┐ ┌──────┐                │
│  │100万+│ │99.9% │ │50PB+ │ │24/7  │                │
│  │ 用户  │ │可用性 │ │存储量 │ │ 支持  │                │
│  └──────┘ └──────┘ └──────┘ └──────┘                │
├──────────────────────────────────────────────────────┤
│  CTA Section                                         │
│  开始你的智能云盘之旅                                  │
│  [立即注册] [登录]                                    │
├──────────────────────────────────────────────────────┤
│  Footer                                              │
└──────────────────────────────────────────────────────┘
```

**设计要点**:
1. Hero 背景使用粉色到金色的渐变光晕动画 (gradient orbs)
2. 所有卡片使用 `glass-card` 类
3. Stats 区域使用主色渐变背景 + 半透明白色玻璃卡片
4. CTA 区域使用金色高亮按钮

#### 8.3.2 登录/注册页面 — 全新设计

```
┌─────────────────────────────────────────────────┐
│                 Auth Layout                      │
│                                                  │
│  ┌─────────────────────────────────────┐         │
│  │          Glass Card (居中)          │         │
│  │                                     │         │
│  │      ┌─────────────────┐            │         │
│  │      │      Logo       │            │         │
│  │      └─────────────────┘            │         │
│  │                                     │         │
│  │     欢迎回来                         │         │
│  │     登录你的小七云盘账户               │         │
│  │                                     │         │
│  │  ┌─────────────────────────┐        │         │
│  │  │ 📧 手机号 / 用户名      │        │         │
│  │  └─────────────────────────┘        │         │
│  │  ┌─────────────────────────┐        │         │
│  │  │ 🔒 密码                 │        │         │
│  │  └─────────────────────────┘        │         │
│  │                                     │         │
│  │  [✔️ 记住我]       [忘记密码?]       │         │
│  │                                     │         │
│  │  ┌─────────────────────────┐        │         │
│  │  │      🚀 登 录          │        │         │
│  │  └─────────────────────────┘        │         │
│  │                                     │         │
│  │  还没有账户? [立即注册 →]            │         │
│  │                                     │         │
│  └─────────────────────────────────────┘         │
│                                                  │
│  Background: 紫色渐变 + 浮动光晕                  │
└─────────────────────────────────────────────────┘
```

**设计要点**:
1. 全屏粉色渐变背景，带有浮动光晕动画
2. 居中玻璃卡片作为表单容器
3. 输入框使用设计系统的 DSInput，带前置图标
4. 登录按钮使用 cta (金色) 变体，增加视觉动感

#### 8.3.3 文件管理页面 — 核心体验

```
┌────────────────────────────────────────────────────────────┐
│  OperationBar                                              │
│  [← 返回]  [📁 新建文件夹]  [⬆️ 上传]  [⬇️ 下载]  [更多▼]  │
│  [选中的文件: 3项]       [视图切换: ▦表格 ⊞网格]           │
├────────────────────────────────────────────────────────────┤
│  BreadCrumb                                                │
│  🏠 全部 > 📂 文档 > 📂 项目 > 📂 2026                     │
├────────────────────────────────────────────────────────────┤
│  SearchBar                                                 │
│  🔍 在当前目录搜索文件...                              [搜索]│
├────────────────────────────────────────────────────────────┤
│  File Table (或 File Grid)                                 │
│  ┌──┬────────┬───────┬────────┬────────────┐               │
│  │☐│ 文件名  │ 类型  │  大小  │  修改日期   │               │
│  ├──┼────────┼───────┼────────┼────────────┤               │
│  │☐│📁 项目  │文件夹 │  --    │ 05-11 15:30│               │
│  │☐│📁 图片  │文件夹 │  --    │ 05-10 09:12│               │
│  │☐│📄 报告  │ PDF   │ 2.3MB  │ 05-09 14:00│               │
│  │☐│🖼️ 设计稿│ 图片  │ 4.1MB  │ 05-08 11:22│               │
│  │☐│🎞️ 教程  │ 视频  │ 156MB  │ 05-07 16:45│               │
│  └──┴────────┴───────┴────────┴────────────┘               │
│                                                             │
│  已选 3 项                                    第1页/共5页    │
└────────────────────────────────────────────────────────────┘
```

**设计要点**:
1. 操作栏竖直居中，按钮使用设计系统 DSButton
2. 面包屑可点击，使用 `#` 图标装饰
3. 搜索框集成在当前目录上方，支持实时过滤
4. 表格行悬停有浅粉色背景 (`rgba(219,39,119,0.03)`)
5. 选中行有粉色左边框装饰 (`3px solid var(--color-primary)`)
6. 网格视图文件卡片使用 `glass-card` 样式

#### 8.3.4 AI 聊天页面

```
┌─────────────────────────────────────────────────────┐
│ Sidebar             │        Chat Panel             │
│                     │                                │
│  📝 新对话          │  ┌────────────────────────┐   │
│                     │  │   AI Assistant         │   │
│  历史对话           │  │                        │   │
│  ─────────          │  │   用户消息气泡          │   │
│  项目讨论 05/10     │  │   ┌──────────────┐    │   │
│  代码分析 05/09     │  │   │ 用户消息内容   │    │   │
│  文档总结 05/08     │  │   └──────────────┘    │   │
│                     │  │                        │   │
│  搜索对话...        │  │   ┌──────────────┐    │   │
│                     │  │   │ 🤖 AI 回复    │    │   │
│                     │  │   │ 包含代码块     │    │   │
│                     │  │   │ 和 Markdown   │    │   │
│                     │  │   └──────────────┘    │   │
│                     │  │                        │   │
│                     │  ├────────────────────────┤   │
│                     │  │ 输入框 [📎] [输入...] [➤] │   │
│                     │  └────────────────────────┘   │
└─────────────────────────────────────────────────────┘
```

---

## 9. 用户流程设计

### 9.1 核心用户流程

```
                                 ┌──────────┐
                                 │  访问网站  │
                                 └────┬─────┘
                                      │
                         ┌────────────┴────────────┐
                         ▼                         ▼
                   ┌──────────┐            ┌──────────────┐
                   │  首页浏览  │            │  直接访问分享  │
                   └────┬─────┘            └──────┬───────┘
                        │                         │
               ┌────────┴────────┐                 │
               ▼                 ▼                 ▼
         ┌──────────┐     ┌──────────┐     ┌──────────────┐
         │  登录     │     │  注册    │     │ 分享提取码验证 │
         └────┬─────┘     └────┬─────┘     └──────┬───────┘
              │                │                   │
              └────────┬───────┘                   │
                       ▼                           ▼
                 ┌──────────┐              ┌──────────────┐
                 │ 文件管理   │              │ 浏览分享文件  │
                 │  Dashboard│              │ 或转存至网盘  │
                 └────┬─────┘              └──────────────┘
                      │
        ┌─────────────┼──────────────────────────────┐
        ▼             ▼              ▼                ▼
  ┌─────────┐  ┌──────────┐  ┌──────────┐    ┌──────────┐
  │ 上传文件  │  │ 新建文件夹 │  │ 分享文件  │    │ AI 处理  │
  └────┬────┘  └──────────┘  └────┬─────┘    └────┬─────┘
       │                           │               │
       ▼                           ▼               ▼
  ┌──────────┐              ┌──────────┐    ┌──────────┐
  │ 拖拽/选择  │              │ 设置权限  │    │ 智能问答  │
  │ 文件上传  │              │ 生成链接  │    │ 文档分析  │
  │ 进度展示  │              │ 复制分享  │    │ 内容生成  │
  └──────────┘              └──────────┘    └──────────┘
```

### 9.2 关键交互流程

#### 上传文件流程

```
[点击上传] → [文件选择器打开]
    │
    ├── 小文件 (<10MB) → [秒传检测] → [直接上传] → [进度条] → [完成提示]
    │
    └── 大文件 (>10MB) → [MD5计算] → [分片上传]
                                    ├── 已存在 → [秒传成功]
                                    └── 不存在 → [分片上传] → [合并] → [完成提示]
```

#### 分享文件流程

```
[选中文件] → [点击分享] → [分享弹窗]
    │
    ├── [选择有效期: 永久/7天/30天]
    ├── [设置提取码: 4-6位]
    │
    └── [确认分享] → [生成链接] → [复制链接] → [分享成功提示]
```

---

## 10. 交互与动效规范

### 10.1 动效时长

| 场景 | 时长 | 缓动函数 | 说明 |
|------|------|---------|------|
| 微交互 (按钮悬停) | 150ms | ease | 最小反馈 |
| 组件过渡 | 200ms | ease | 一般交互动画 |
| 面板滑动 | 300ms | cubic-bezier(0.4, 0, 0.2, 1) | 侧边栏、弹窗 |
| 页面切换 | 400ms | cubic-bezier(0.4, 0, 0.2, 1) | 路由过渡 |
| 大背景动画 (orbs) | 20s | ease-in-out | 背景光晕浮动 |

### 10.2 微交互动效清单

| 交互 | 效果 | CSS 实现 |
|------|------|----------|
| 按钮悬停 | 上移1px + 阴影增强 | `transform: translateY(-1px); box-shadow: ...` |
| 按钮点击 | 缩放 0.97 | `transform: scale(0.97)` |
| 卡片悬停 | 上移2px + 阴影增强 | `transform: translateY(-2px); box-shadow: var(--shadow-lg)` |
| 链接悬停 | 下划线渐入 | `::after { width: 0 → 100% }` |
| 侧边栏展开 | 平滑宽度过渡 | `transition: width 300ms cubic-bezier(...)` |
| 弹窗出现 | 缩放淡入 | `opacity: 0 → 1; transform: scale(0.95 → 1)` |
| 消息提示 | 从上方滑入 | `transform: translateY(-20px → 0)` |
| 文件操作 | 行高亮闪烁 | 选中行背景变色动画 |
| 加载中 | 渐变脉冲 | `@keyframes pulse` |

### 10.3 页面过渡

```css
/* 路由页面切换 */
.page-enter-active,
.page-leave-active {
  transition: opacity 300ms ease, transform 300ms ease;
}

.page-enter-from {
  opacity: 0;
  transform: translateY(12px);
}

.page-leave-to {
  opacity: 0;
  transform: translateY(-12px);
}
```

---

## 11. 响应式策略

### 11.1 响应式断点

| 断点名称 | 最小宽度 | 最大宽度 | 目标设备 |
|----------|---------|---------|----------|
| `xs` | 0 | 575px | 小屏手机 (iPhone SE) |
| `sm` | 576px | 767px | 大屏手机/小平板 |
| `md` | 768px | 1023px | 平板 (iPad) |
| `lg` | 1024px | 1279px | 桌面 |
| `xl` | 1280px | 1439px | 大桌面 |
| `2xl` | 1440px | - | 超大屏幕 |

### 11.2 响应式适配规则

| 元素 | Mobile | Tablet | Desktop |
|------|--------|--------|---------|
| Header 高度 | 52-56px | 60px | 60px |
| Sidebar | 抽屉式(遮罩) | 固定/可折叠 | 固定/可折叠 |
| 文件视图 | 列表(隐藏部分列) | 列表/网格切换 | 列表/网格切换 |
| 操作栏 | 图标+文字(折叠) | 完整显示 | 完整显示 |
| 弹窗 | 全屏/底部弹窗 | 居中弹窗 | 居中弹窗 |
| 卡片网格 | 1列 | 2-3列 | 3-4列 |
| 字号 | 缩小1档 | 正常 | 正常 |
| 触摸目标 | ≥44px | ≥40px | ≥36px |

### 11.3 移动端特殊处理

```css
/* 移动端文件列表: 隐藏类型列和日期列，文件名占满宽度 */
@media (max-width: 767px) {
  .type-col,
  .date-col {
    display: none;
  }
  
  .file-name-text {
    max-width: none;
    flex: 1;
  }
  
  /* 底部操作栏固定 */
  .mobile-action-bar {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    background: var(--glass-strong);
    backdrop-filter: blur(20px);
    padding: 12px 16px;
    z-index: var(--z-fixed);
  }
  
  /* 侧边栏: 全屏抽屉 */
  .layout-sidebar {
    position: fixed;
    left: 0;
    top: 56px;
    width: 280px;
    height: calc(100vh - 56px);
    z-index: var(--z-fixed);
    transform: translateX(-100%);
    transition: transform 300ms cubic-bezier(0.4, 0, 0.2, 1);
  }
  
  .layout-sidebar.open {
    transform: translateX(0);
  }
}
```

---

## 12. 无障碍设计

### 12.1 色彩对比度标准

所有文本和交互元素必须满足 WCAG 2.1 AA 标准:

| 文本类型 | 最小对比度 | 目标 |
|----------|-----------|------|
| 正文 (>=14px) | 4.5:1 | ✅ AAA 7:1 |
| 大文本 (>=18px Bold / >=24px) | 3:1 | ✅ AA 4.5:1 |
| 禁用文本 | - | 不要求，但需可识别 |
| 图标按钮 | 3:1 | ✅ AA |

### 12.2 键盘导航

| 快捷键 | 功能 |
|--------|------|
| `Tab` | 在交互元素间跳转 |
| `Enter` / `Space` | 激活当前元素 |
| `Esc` | 关闭弹窗/菜单/下拉 |
| `↑` / `↓` | 列表/菜单项导航 |
| `Ctrl + F` | 搜索文件 |
| `Delete` | 删除选中文件 |
| `Ctrl + A` | 全选文件 |
| `Ctrl + C` / `V` | 复制/粘贴文件 |

### 12.3 ARIA 属性标准

| 组件 | 必需 ARIA 属性 | 说明 |
|------|---------------|------|
| 按钮 | `role="button"` | 非 `<button>` 元素 |
| 弹窗 | `role="dialog"`, `aria-modal="true"`, `aria-labelledby` | 模态框 |
| 标签页 | `role="tablist"`, `role="tab"`, `aria-selected` | 选项卡 |
| 菜单 | `role="menu"`, `role="menuitem"` | 下拉/右键菜单 |
| 进度条 | `role="progressbar"`, `aria-valuenow` | 上传进度 |
| 通知 | `role="alert"`, `aria-live="polite"` | Toast/消息 |
| 错误 | `aria-invalid="true"`, `aria-describedby` | 表单验证 |
| 开关 | `role="switch"`, `aria-checked` | 切换开关 |

### 12.4 焦点管理

```typescript
// 弹窗打开时自动聚焦到关闭按钮
onMounted(() => {
  nextTick(() => {
    closeButtonRef.value?.focus();
  });
});

// 弹窗关闭时聚焦回到触发元素
const handleClose = () => {
  triggerElementRef.value?.focus();
};
```

---

## 13. 暗色模式

### 13.1 暗色模式色值

| CSS 变量 | 亮色模式 | 暗色模式 |
|----------|---------|---------|
| `--color-bg` | `#FDF2F8` | `#0F172A` |
| `--color-surface` | `#FFFFFF` | `#1E293B` |
| `--color-text-primary` | `#831843` | `#F8FAFC` |
| `--color-text-secondary` | `#64748B` | `#94A3B8` |
| `--color-border` | `#E2E8F0` | `#334155` |
| `--glass-bg` | `rgba(255,255,255,0.8)` | `rgba(30,41,59,0.8)` |
| `--glass-shadow` | `0 8px 32px rgba(0,0,0,0.12)` | `0 8px 32px rgba(0,0,0,0.4)` |

### 13.2 暗色模式切换实现

```typescript
// store/theme.ts
import { defineStore } from 'pinia';
import { ref, watch } from 'vue';

export const useThemeStore = defineStore('theme', () => {
  const isDark = ref(false);

  function toggleTheme() {
    isDark.value = !isDark.value;
    applyTheme();
  }

  function applyTheme() {
    if (isDark.value) {
      document.documentElement.setAttribute('data-theme', 'dark');
    } else {
      document.documentElement.removeAttribute('data-theme');
    }
    localStorage.setItem('theme', isDark.value ? 'dark' : 'light');
  }

  function initTheme() {
    const saved = localStorage.getItem('theme');
    if (saved === 'dark') {
      isDark.value = true;
      applyTheme();
    }
  }

  return { isDark, toggleTheme, initTheme };
});
```

### 13.3 图片适配

```css
/* 暗色模式下的图片处理 */
[data-theme="dark"] img {
  filter: brightness(0.9) contrast(1.1);
}

/* SVG 图标颜色继承 */
[data-theme="dark"] .icon-invert {
  filter: invert(1) hue-rotate(180deg);
}
```

---

## 14. 设计交付与开发协作

### 14.1 组件文档规范

每个组件文档需包含:

```markdown
## [组件名]

### 概述
一句话描述组件用途

### Props
| 属性 | 类型 | 默认值 | 必填 | 说明 |

### Events
| 事件名 | 参数 | 说明 |

### Slots
| 插槽名 | 作用域 | 说明 |

### 使用示例
```vue
<template>
  <DSButton variant="primary" size="medium" @click="handleClick">
    按钮文字
  </DSButton>
</template>
```

### 状态预览
- Default / Hover / Active / Disabled / Loading

### 设计说明
- 使用场景、布局建议、注意事项
```

### 14.2 开发协作 CheckList

```
□ 组件是否使用设计系统 CSS 变量?
□ 组件是否定义了完整的 TypeScript 类型?
□ 组件是否覆盖了所有交互状态 (hover/active/disabled/loading/focus/error)?
□ 组件是否支持键盘导航和 ARIA 属性?
□ 组件是否适配移动端 (触摸目标 ≥44px)?
□ 组件是否支持暗色模式?
□ 组件动画是否遵循 prefers-reduced-motion?
□ 组件是否移除了所有 console.log 调试代码?
□ 组件是否添加了 JSDoc 注释?
□ 组件是否在移动端和桌面端都进行了测试?
```

### 14.3 设计验收标准

| 标准 | 要求 | 检查方法 |
|------|------|---------|
| 像素级还原 | 与设计稿偏差 ≤ 2px | 视觉对比 |
| 组件一致性 | 同组件在不同页面表现一致 | 组件库审计 |
| 响应式完整 | 4个断点均正常显示 | 设备模拟 |
| 交互完整 | 所有状态存在且正确 | 交互测试 |
| 无障碍 | 满足 WCAG AA | Lighthouse 检查 |

---

## 15. 迁移路线图

### Phase 1: 设计系统统一 (预计 2天)
- [ ] 统一 HomeView 为紫色玻璃态主题
- [ ] 统一 GlobalHeader 为设计系统变量
- [ ] 统一 LoginView / RegisterView 为设计系统
- [ ] 验证 4 个 DS 组件在所有页面的兼容性

### Phase 2: 组件增强 (预计 3天)
- [ ] 新增 DSModal 组件
- [ ] 新增 DSFileIcon 组件 (文件类型图标)
- [ ] 新增 DSFileGridItem 组件 (网格文件卡片)
- [ ] 优化 FileTable 悬停和选中样式
- [ ] 完善 DSDropdown 组件

### Phase 3: 交互升级 (预计 2天)
- [ ] 添加页面切换过渡动画
- [ ] 完善微交互动效
- [ ] 实现上传进度动画
- [ ] 优化移动端手势交互

### Phase 4: 暗色模式 (预计 1天)
- [ ] 实现暗色模式切换
- [ ] 验证所有组件的暗色兼容性
- [ ] 优化暗色模式下的图片显示

### Phase 5: 无障碍优化 (预计 1天)
- [ ] 检查并修复所有组件的键盘导航
- [ ] 添加 ARIA 属性标注
- [ ] 验证色彩对比度
- [ ] Lighthouse 审计

---

## 附录 A: 设计资源

### 字体资源
- [Fira Sans (Google Fonts)](https://fonts.google.com/specimen/Fira+Sans)
- [Fira Code (Google Fonts)](https://fonts.google.com/specimen/Fira+Code)

### 图标资源
- [Lucide Icons](https://lucide.dev/icons/) — 已在项目中安装

### 设计工具
- 建议使用 Figma 创建组件原型
- 设计令牌可用 Figma Tokens 插件管理

### 参考设计
- [Google Material Design 3](https://m3.material.io/)
- [Apple Human Interface Guidelines](https://developer.apple.com/design/human-interface-guidelines/)

---

## 附录 B: CSS 变量对照表 (新旧兼容)

| 新变量 | 旧变量 (已弃用) | 值 |
|--------|----------------|-----|
| `--color-primary` | `--ds-color-primary` | `#DB2777` |
| `--color-secondary` | `--ds-color-secondary` | `#F472B6` |
| `--color-cta` | `--ds-color-cta` | `#D97706` |
| `--text-base` | `--ds-text-size-base` | `14px` |
| `--radius-md` | `--ds-radius-md` | `8px` |
| `--space-5` | `--spacing-md` | `16px` |

> 过渡期内新旧变量并存。新组件优先使用新变量命名。

---

*本设计规范由 UI Designer 于 2026-05-11 创建，适用于 AqiCloud-AiPan v2.0.0 设计系统更新（粉色主题）。*
