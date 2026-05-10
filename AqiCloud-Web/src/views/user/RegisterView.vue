<template>
  <div class="register-page">
    <!-- 浮动背景装饰 -->
    <div class="bg-decoration">
      <div class="bg-orb bg-orb-1"></div>
      <div class="bg-orb bg-orb-2"></div>
      <div class="bg-orb bg-orb-3"></div>
    </div>
    <!-- 注册表单区域 -->
    <div class="register-form-section">
      <div class="form-container">
        <div class="form-header">
          <h2 class="form-title">{{ t('register.title') }}</h2>
          <p class="form-subtitle">{{ t('register.subtitle') }}</p>
        </div>

        <div class="form-body">
          <!-- 头像上传 -->
          <div class="avatar-row">
            <div class="avatar-wrapper">
              <ImageUpload v-model="avatarUrl" />
            </div>
            <span class="avatar-hint">{{ t('register.avatarHint') }}</span>
          </div>

          <!-- 用户名 -->
          <div class="input-row">
            <label class="input-label">{{ t('register.username') }} <span class="required">*</span></label>
            <div class="input-wrapper">
              <el-icon class="input-prefix"><User /></el-icon>
              <input
                v-model="registerForm.phone"
                type="text"
                :placeholder="t('register.usernamePlaceholder')"
                class="custom-input"
                @focus="clearError('phone')"
              />
            </div>
            <span v-if="errors.phone" class="error-text">{{ errors.phone }}</span>
          </div>

          <!-- 密码 -->
          <div class="input-row">
            <label class="input-label">{{ t('register.password') }} <span class="required">*</span></label>
            <div class="input-wrapper">
              <el-icon class="input-prefix"><Lock /></el-icon>
              <input
                v-model="registerForm.password"
                type="password"
                :placeholder="t('register.passwordPlaceholder')"
                class="custom-input"
                @focus="clearError('password')"
              />
            </div>
            <span class="hint-text">{{ t('register.passwordHint') }}</span>
            <span v-if="errors.password" class="error-text">{{ errors.password }}</span>
          </div>

          <!-- 确认密码 -->
          <div class="input-row">
            <label class="input-label">{{ t('register.confirmPassword') }} <span class="required">*</span></label>
            <div class="input-wrapper">
              <el-icon class="input-prefix"><Lock /></el-icon>
              <input
                v-model="registerForm.confirmPassword"
                type="password"
                :placeholder="t('register.confirmPasswordPlaceholder')"
                class="custom-input"
                @focus="clearError('confirmPassword')"
              />
            </div>
            <span v-if="errors.confirmPassword" class="error-text">{{ errors.confirmPassword }}</span>
          </div>

          <!-- 注册按钮 -->
          <button
            class="register-btn"
            :class="{ 'loading': loading }"
            :disabled="loading"
            @click="handleRegister"
          >
            <span v-if="!loading">{{ t('register.submit') }}</span>
            <span v-else class="loading-spinner"></span>
          </button>

          <!-- 登录链接 -->
          <div class="login-link-row">
            <span>{{ t('register.hasAccount') }}</span>
            <router-link to="/user/login" class="login-link">{{ t('register.loginLink') }}</router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, watch } from "vue";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";
import { User, Lock, Check } from "@element-plus/icons-vue";
import { userRegister } from "@/api/user";
import ImageUpload from "@/components/common/ImageUpload.vue";
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const avatarUrl = ref("");
const router = useRouter();
const loading = ref(false);

// 错误信息
const errors = reactive({
  phone: "",
  password: "",
  confirmPassword: "",
});

const registerForm = reactive({
  phone: "",
  username: "",
  password: "",
  confirmPassword: "",
  avatarUrl: "",
} as API.UserRegisterRequest);

// 监听头像URL变化
watch(avatarUrl, (newVal) => {
  registerForm.avatarUrl = newVal;
}, { immediate: true });

// 从本地存储获取头像URL
if (localStorage.getItem("avatarUrl")) {
  avatarUrl.value = localStorage.getItem("avatarUrl")!;
}

// 清除错误
const clearError = (field: keyof typeof errors) => {
  errors[field] = "";
};

const handleRegister = async () => {
  // 清空错误信息
  Object.keys(errors).forEach(key => {
    errors[key as keyof typeof errors] = "";
  });

  // 验证表单
  if (!registerForm.phone) {
    errors.phone = t('register.usernameRequired');
    ElMessage.error(t('register.usernameRequired'));
    return;
  }

  if (!registerForm.password) {
    errors.password = t('register.passwordRequired');
    ElMessage.error(t('register.passwordRequired'));
    return;
  }

  if (registerForm.password.length < 6) {
    errors.password = t('register.passwordMinLength');
    ElMessage.error(t('register.passwordMinLength'));
    return;
  }

  if (!registerForm.confirmPassword) {
    errors.confirmPassword = t('register.confirmPasswordRequired');
    ElMessage.error(t('register.confirmPasswordRequired'));
    return;
  }

  if (registerForm.confirmPassword !== registerForm.password) {
    errors.confirmPassword = t('register.confirmPasswordMismatch');
    ElMessage.error(t('register.confirmPasswordMismatch'));
    return;
  }

  if (!registerForm.avatarUrl) {
    ElMessage.error(t('register.avatarRequired'));
    return;
  }

  try {
    loading.value = true;
    const res = await userRegister(registerForm);

    if (res.data.code === 0) {
      ElMessage.success(t('register.success'));
      router.push("/user/login");
      localStorage.removeItem("avatarUrl");
    } else {
      ElMessage.error(`${t('register.errorTitle')}：${res.data.msg}`);
    }
  } catch (error) {
    console.error(t('register.errorTitle'), error);
    ElMessage.error(t('register.errorMessage'));
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
.register-page {
  display: flex;
  min-height: 100vh;
  background: linear-gradient(135deg, #FDF2F8 0%, #F5F3FF 50%, #FCE7F3 100%);
  position: relative;
  overflow: hidden;
}

/* 注册表单区域 - 全屏居中 */
.register-form-section {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 48px 60px;
  position: relative;
  z-index: 1;
  box-sizing: border-box;
}

.form-container {
  width: 100%;
  max-width: 400px;
  background: rgba(255,255,255,0.65);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: 24px;
  padding: 40px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.08);
  border: 1px solid rgba(255,255,255,0.5) inset;
  animation: fadeInUp 0.6s ease-out;
}

.form-header {
  text-align: center;
  margin-bottom: 32px;
}

.form-title {
  font-size: 28px;
  font-weight: 700;
  background: linear-gradient(135deg, #DB2777 0%, #A855F7 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin: 0 0 8px 0;
}

.form-subtitle {
  font-size: 14px;
  color: #64748B;
  margin: 0;
}

.form-body {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

/* 头像上传 */
.avatar-row {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  margin-bottom: 8px;
}

.avatar-wrapper {
  width: 80px;
  height: 80px;
}

.avatar-wrapper :deep(.avatar-uploader) {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  border: 2px dashed #E2E8F0;
  background: #F8FAFC;
}

.avatar-wrapper :deep(.avatar-uploader:hover) {
  border-color: #A855F7;
  background: #F5F3FF;
}

.avatar-wrapper :deep(.avatar-uploader-icon) {
  font-size: 24px;
}

.avatar-hint {
  font-size: 12px;
  color: #94A3B8;
}

/* 输入框 */
.input-row {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.input-label {
  font-size: 14px;
  font-weight: 500;
  color: #1E1B4B;
}

.required {
  color: #EF4444;
}

.input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.input-prefix {
  position: absolute;
  left: 14px;
  font-size: 18px;
  color: #94A3B8;
  z-index: 1;
}

.custom-input {
  width: 100%;
  height: 44px;
  padding: 0 14px 0 42px;
  border: 1px solid #E2E8F0;
  border-radius: 10px;
  font-size: 14px;
  color: #1E1B4B;
  background: #FFFFFF;
  transition: all 0.2s ease;
  outline: none;
}

.custom-input:focus {
  border-color: #A855F7;
  box-shadow: 0 0 0 3px rgba(168, 85, 247, 0.1);
}

.custom-input::placeholder {
  color: #94A3B8;
}

.hint-text {
  font-size: 12px;
  color: #94A3B8;
}

.error-text {
  font-size: 12px;
  color: #EF4444;
}

/* 注册按钮 */
.register-btn {
  width: 100%;
  height: 48px;
  margin-top: 8px;
  background: linear-gradient(135deg, #DB2777 0%, #A855F7 100%);
  color: white;
  border: none;
  border-radius: 10px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.register-btn:hover:not(:disabled) {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(219, 39, 119, 0.4);
}

.register-btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.loading-spinner {
  width: 20px;
  height: 20px;
  border: 2px solid rgba(255,255,255,0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* 背景装饰 */
.bg-decoration {
  position: absolute;
  inset: 0;
  pointer-events: none;
  overflow: hidden;
}

.bg-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(80px);
}

.bg-orb-1 {
  width: 300px;
  height: 300px;
  background: rgba(244,114,182,0.2);
  top: 15%;
  right: 10%;
  animation: floatOrb 8s ease-in-out infinite;
}

.bg-orb-2 {
  width: 400px;
  height: 400px;
  background: rgba(192,132,252,0.2);
  bottom: 20%;
  left: 5%;
  animation: floatOrb 10s ease-in-out infinite 2s;
}

.bg-orb-3 {
  width: 200px;
  height: 200px;
  background: rgba(251,146,60,0.15);
  top: 50%;
  right: 30%;
  animation: floatOrb 7s ease-in-out infinite 4s;
}

@keyframes floatOrb {
  0%, 100% { transform: translate(0, 0) scale(1); }
  33% { transform: translate(30px, -20px) scale(1.05); }
  66% { transform: translate(-20px, 15px) scale(0.95); }
}

@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(30px); }
  to { opacity: 1; transform: translateY(0); }
}

/* 登录链接 */
.login-link-row {
  text-align: center;
  font-size: 14px;
  color: #64748B;
  margin-top: 8px;
}

.login-link {
  color: #A855F7;
  font-weight: 600;
  text-decoration: none;
  margin-left: 4px;
}

.login-link:hover {
  text-decoration: underline;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .register-page {
    flex-direction: column;
  }

  .register-form-section {
    flex: 1;
    padding: 40px 24px;
    background: #ffffff;
  }

  .form-container {
    width: 100%;
    max-width: 100%;
    padding: 40px 32px;
    background: white;
    border-radius: 16px;
    box-shadow: 0 4px 24px rgba(0, 0, 0, 0.06);
  }

  .form-header {
    margin-bottom: 32px;
  }

  .form-title {
    font-size: 26px;
  }

  .form-body {
    gap: 20px;
  }

  .custom-input {
    height: 48px;
  }

  .register-btn {
    height: 50px;
    margin-top: 8px;
  }
}

@media (max-width: 480px) {
  .register-form-section {
    padding: 24px 16px;
    min-height: 100vh;
  }

  .form-container {
    width: 100%;
    max-width: 100%;
    padding: 32px 20px;
    box-shadow: none;
    border-radius: 12px;
    background: transparent;
  }

  .form-title {
    font-size: 24px;
  }

  .form-subtitle {
    font-size: 13px;
  }

  .custom-input {
    height: 50px;
    font-size: 15px;
  }

  .register-btn {
    height: 50px;
    font-size: 16px;
  }
}
</style>
