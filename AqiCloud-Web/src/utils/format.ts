/**
 * 文件格式化工具函数
 * @module utils/format
 */

/**
 * 格式化文件大小
 * 将字节大小转换为人类可读的格式（B、KB、MB、GB、TB）
 * @param size 文件大小（字节）
 * @returns 格式化后的文件大小字符串
 */
export const formatFileSize = (
  size: string | number | null | undefined,
): string => {
  if (size == null) return "-";

  let sizeNum = typeof size === "string" ? parseFloat(size) : size;
  if (typeof sizeNum !== "number" || isNaN(sizeNum)) return "-";

  const units = ["B", "KB", "MB", "GB", "TB"];
  let index = 0;

  while (sizeNum >= 1024 && index < units.length - 1) {
    sizeNum /= 1024;
    index++;
  }

  return `${sizeNum.toFixed(2)} ${units[index]}`;
};

/**
 * 格式化日期时间
 * 将时间戳或日期字符串转换为本地化的日期时间字符串
 * @param timestamp 时间戳或日期字符串
 * @returns 本地化的日期时间字符串
 */
export const formatDateTime = (
  timestamp: number | string | undefined,
): string => {
  if (!timestamp) return "-";

  if (typeof timestamp === "number") {
    return new Date(timestamp).toLocaleString();
  }

  return new Date(timestamp).toLocaleString();
};

/**
 * 获取文件类型名称
 * @param fileType 文件类型（字符串或包含 fileType 字段的对象）
 * @returns 格式化后的文件类型名称
 */
export const getFileTypeName = (
  fileType:
    | string
    | { fileType?: string; fileSuffix?: string }
    | null
    | undefined,
): string => {
  if (!fileType) return "-";

  // 如果传入的是对象，提取 fileType 或 fileSuffix
  if (typeof fileType === "object") {
    const type = fileType.fileType || fileType.fileSuffix;
    if (!type) return "-";

    // 处理文件夹类型
    if (type === "DIR" || type === "folder") return "文件夹";

    return typeof type === "string" ? type.toUpperCase() : "-";
  }

  // 处理字符串类型
  if (typeof fileType !== "string") return "-";
  if (fileType === "folder" || fileType === "DIR") return "文件夹";

  return fileType.toUpperCase();
};

/**
 * 获取文件扩展名
 * @param fileName 文件名
 * @returns 文件扩展名（小写）
 */
export const getFileExtension = (fileName: string): string => {
  const lastDot = fileName.lastIndexOf(".");
  return lastDot !== -1 ? fileName.substring(lastDot + 1).toLowerCase() : "";
};

/**
 * 判断是否为图片文件
 * @param fileSuffix 文件后缀
 * @returns 是否为图片
 */
export const isImageFile = (fileSuffix: string | null | undefined): boolean => {
  if (!fileSuffix) return false;
  const imageExtensions = ["jpg", "jpeg", "png", "gif", "bmp", "webp", "svg"];
  return imageExtensions.includes(fileSuffix.toLowerCase());
};

/**
 * 判断是否为视频文件
 * @param fileSuffix 文件后缀
 * @returns 是否为视频
 */
export const isVideoFile = (fileSuffix: string | null | undefined): boolean => {
  if (!fileSuffix) return false;
  const videoExtensions = ["mp4", "avi", "mov", "wmv", "flv", "mkv", "webm"];
  return videoExtensions.includes(fileSuffix.toLowerCase());
};

/**
 * 判断是否为音频文件
 * @param fileSuffix 文件后缀
 * @returns 是否为音频
 */
export const isAudioFile = (fileSuffix: string | null | undefined): boolean => {
  if (!fileSuffix) return false;
  const audioExtensions = ["mp3", "wav", "flac", "aac", "ogg", "wma", "m4a"];
  return audioExtensions.includes(fileSuffix.toLowerCase());
};

/**
 * 判断是否为文档文件
 * @param fileSuffix 文件后缀
 * @returns 是否为文档
 */
export const isDocumentFile = (
  fileSuffix: string | null | undefined,
): boolean => {
  if (!fileSuffix) return false;
  const docExtensions = [
    "pdf",
    "doc",
    "docx",
    "xls",
    "xlsx",
    "ppt",
    "pptx",
    "txt",
    "md",
  ];
  return docExtensions.includes(fileSuffix.toLowerCase());
};
