<template>
  <el-upload
    class="avatar-uploader"
    :show-file-list="false"
    :before-upload="beforeAvatarUpload"
    :http-request="customUpload"
  >
    <img v-if="imageUrl" :src="imageUrl" class="avatar" :alt="'头像'" />
    <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
  </el-upload>
</template>

<script setup lang="ts">
import { uploadAvatar } from "@/api/user";
import { Plus } from "@element-plus/icons-vue";
import { ElMessage } from "element-plus";
import { ref, watch } from "vue";

/**
 * ImageUpload 组件 - 图片上传组件
 * 支持头像上传和预览功能
 */

interface ImageUploadProps {
  modelValue?: string;
}

interface UploadOptions {
  file: File;
  onSuccess: () => void;
  onError: (error: any) => void;
}

const props = defineProps<ImageUploadProps>();

const emit = defineEmits<{
  "update:modelValue": [value: string];
}>();

// 初始化时清理URL，移除不必要的前缀
const initialImageUrl =
  props.modelValue?.replace("http://localhost:8080/user/", "") || "";
const imageUrl = ref(initialImageUrl);

/**
 * 监听 modelValue 变化
 */
watch(
  () => props.modelValue,
  (newValue) => {
    imageUrl.value = newValue?.replace("http://localhost:8080/user/", "") || "";
  },
);

/**
 * 上传前的文件验证
 * @param file - 待上传的文件对象
 * @returns 是否通过验证
 */
const beforeAvatarUpload = (file: File): boolean => {
  const isJPG = file.type === "image/jpeg";
  const isPNG = file.type === "image/png";
  const isLt2M = file.size / 1024 / 1024 < 2;

  if (!isJPG && !isPNG) {
    ElMessage.error("上传头像图片只能是 JPG 或 PNG 格式!");
    return false;
  }
  if (!isLt2M) {
    ElMessage.error("上传头像图片大小不能超过 2MB!");
    return false;
  }
  return true;
};

/**
 * 自定义上传实现
 * @param options - 上传选项
 */
const customUpload = async (options: UploadOptions): Promise<void> => {
  try {
    const formData = new FormData();
    formData.append("file", options.file);
    const res = await uploadAvatar(formData);

    if (
      res &&
      res.data &&
      res.data.code === 0 &&
      typeof res.data.data === "string"
    ) {
      // 后端返回的已经是完整的URL，直接使用
      const avatarUrl = res.data.data;
      imageUrl.value = avatarUrl;
      localStorage.setItem("avatarUrl", avatarUrl);
      emit("update:modelValue", avatarUrl);
      ElMessage.success("头像上传成功");
      options.onSuccess();
    } else {
      throw new Error(res?.data?.msg || "上传失败");
    }
  } catch (error: any) {
    ElMessage.error("上传失败，请稍后重试");
    options.onError(error);
  }
};
</script>

<style scoped>
.avatar-uploader {
  text-align: center;
  border: 2px dashed var(--el-border-color);
  border-radius: 8px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: all 0.3s;
  aspect-ratio: 1 / 1;
  width: 100%;
  max-width: 178px;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #fafafa;
}

.avatar-uploader:hover {
  border-color: var(--el-color-primary);
  background-color: #f5f7fa;
}

.avatar-uploader-icon {
  font-size: 32px;
  color: #8c939d;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
}

.avatar {
  width: 100%;
  height: 100%;
  display: block;
  object-fit: cover;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .avatar-uploader {
    max-width: 150px;
  }

  .avatar-uploader-icon {
    font-size: 28px;
  }
}
</style>
