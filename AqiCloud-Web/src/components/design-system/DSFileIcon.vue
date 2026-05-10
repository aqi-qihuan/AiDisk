<template>
  <div :class="['ds-file-icon', `ds-file-icon-${size}`]">
    <component :is="iconComponent" :class="['file-icon', iconColorClass]" />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import {
  Folder,
  FolderOpened,
  Document,
  Picture,
  VideoCamera,
  Headset,
  DocumentCopy,
  Collection,
  Tickets,
  Reading,
  Files,
  Film,
  Coffee,
  Box,
  Cpu,
  Memo,
} from '@element-plus/icons-vue';

/**
 * DSFileIcon - 设计系统文件图标组件
 * 根据文件类型显示相应的图标
 */

interface Props {
  fileSuffix?: string;
  isFolder?: boolean;
  size?: 'small' | 'medium' | 'large';
}

const props = withDefaults(defineProps<Props>(), {
  fileSuffix: '',
  isFolder: false,
  size: 'medium',
});

// 文件类型映射到图标
const fileTypeIconMap: Record<string, any> = {
  // 图片
  jpg: Picture,
  jpeg: Picture,
  png: Picture,
  gif: Picture,
  bmp: Picture,
  webp: Picture,
  svg: Picture,
  
  // 视频
  mp4: Film,
  avi: Film,
  mov: Film,
  wmv: Film,
  flv: Film,
  mkv: Film,
  webm: Film,
  
  // 音频
  mp3: Headset,
  wav: Headset,
  flac: Headset,
  aac: Headset,
  ogg: Headset,
  wma: Headset,
  m4a: Headset,
  
  // 文档
  pdf: Reading,
  doc: DocumentCopy,
  docx: DocumentCopy,
  xls: Tickets,
  xlsx: Tickets,
  ppt: Collection,
  pptx: Collection,
  txt: Document,
  md: Memo,
  
  // 代码
  js: Document,
  ts: Document,
  vue: Document,
  html: Document,
  css: Document,
  json: Document,
  java: Coffee,
  py: Document,
  go: Document,
  cpp: Cpu,
  c: Cpu,
  h: Cpu,
  
  // 压缩文件
  zip: Files,
  rar: Files,
  '7z': Files,
  tar: Files,
  gz: Files,
  bz2: Files,
  
  // 可执行文件
  jar: Box,
  war: Box,
  ear: Box,
  exe: Cpu,
  dll: Cpu,
  so: Cpu,
  sh: Document,
  bat: Document,
};

// 文件类型颜色映射
const fileTypeColorMap: Record<string, string> = {
  // 图片 - 蓝色
  jpg: 'color-primary',
  jpeg: 'color-primary',
  png: 'color-primary',
  gif: 'color-primary',
  bmp: 'color-primary',
  webp: 'color-primary',
  svg: 'color-primary',
  
  // 视频 - 紫色
  mp4: 'color-purple',
  avi: 'color-purple',
  mov: 'color-purple',
  wmv: 'color-purple',
  flv: 'color-purple',
  mkv: 'color-purple',
  webm: 'color-purple',
  
  // 音频 - 绿色
  mp3: 'color-success',
  wav: 'color-success',
  flac: 'color-success',
  aac: 'color-success',
  ogg: 'color-success',
  wma: 'color-success',
  m4a: 'color-success',
  
  // 文档 - 橙色
  pdf: 'color-warning',
  doc: 'color-warning',
  docx: 'color-warning',
  xls: 'color-warning',
  xlsx: 'color-warning',
  ppt: 'color-warning',
  pptx: 'color-warning',
  txt: 'color-gray',
  
  // Markdown - 蓝色
  md: 'color-info',
  
  // 代码 - 靛蓝色
  js: 'color-indigo',
  ts: 'color-indigo',
  vue: 'color-indigo',
  html: 'color-indigo',
  css: 'color-indigo',
  json: 'color-indigo',
  java: 'color-indigo',
  py: 'color-indigo',
  go: 'color-indigo',
  
  // 压缩文件 - 黄色
  zip: 'color-yellow',
  rar: 'color-yellow',
  '7z': 'color-yellow',
  tar: 'color-yellow',
  gz: 'color-yellow',
  bz2: 'color-yellow',
  
  // 可执行文件 - 深橙色
  jar: 'color-orange',
  war: 'color-orange',
  ear: 'color-orange',
  exe: 'color-red',
  dll: 'color-red',
  so: 'color-red',
  sh: 'color-gray',
  bat: 'color-gray',
  
  // 代码文件 - 靛蓝色
  cpp: 'color-indigo',
  c: 'color-indigo',
  h: 'color-indigo',
};

/**
 * 计算图标组件
 */
const iconComponent = computed(() => {
  // 如果是文件夹，返回文件夹图标
  if (props.isFolder) {
    return Folder;
  }
  
  // 根据文件后缀返回对应图标
  const suffix = props.fileSuffix.toLowerCase();
  return fileTypeIconMap[suffix] || Document;
});

/**
 * 计算图标颜色类
 */
const iconColorClass = computed(() => {
  // 如果是文件夹，返回黄色
  if (props.isFolder) {
    return 'color-folder';
  }
  
  // 根据文件后缀返回对应颜色
  const suffix = props.fileSuffix.toLowerCase();
  return fileTypeColorMap[suffix] || 'color-default';
});
</script>

<style scoped>
.ds-file-icon {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.ds-file-icon-small {
  width: 24px;
  height: 24px;
}

.ds-file-icon-small .file-icon {
  font-size: 20px;
}

.ds-file-icon-medium {
  width: 32px;
  height: 32px;
}

.ds-file-icon-medium .file-icon {
  font-size: 28px;
}

.ds-file-icon-large {
  width: 48px;
  height: 48px;
}

.ds-file-icon-large .file-icon {
  font-size: 40px;
}

.file-icon {
  transition: all var(--transition-fast, 0.15s) ease;
}

/* 颜色类 */
.color-folder {
  color: #fbbf24;
}

.color-primary {
  color: var(--color-primary, #6366f1);
}

.color-purple {
  color: #8b5cf6;
}

.color-success {
  color: var(--color-success, #10b981);
}

.color-warning {
  color: var(--color-warning, #f59e0b);
}

.color-gray {
  color: #64748b;
}

.color-indigo {
  color: #6366f1;
}

.color-yellow {
  color: #eab308;
}

.color-default {
  color: #94a3b8;
}

.color-orange {
  color: #f97316;
}

.color-red {
  color: #ef4444;
}

.color-info {
  color: #3b82f6;
}
</style>
