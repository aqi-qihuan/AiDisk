<template>
  <div class="login-page">
    <!-- 浮动背景装饰 -->
    <div class="bg-decoration">
      <div class="bg-orb bg-orb-1"></div>
      <div class="bg-orb bg-orb-2"></div>
      <div class="bg-orb bg-orb-3"></div>
    </div>
    <!-- 登录表单区域 -->
    <div class="login-form-section">
      <div class="form-container">
        <div class="form-header">
          <h2 class="form-title">{{ t('login.title') }}</h2>
          <p class="form-subtitle">{{ t('login.subtitle') }}</p>
        </div>

        <div class="form-body">
          <!-- 用户名 -->
          <div class="input-row">
            <label class="input-label">{{ t('login.username') }} <span class="required">*</span></label>
            <div class="input-wrapper">
              <el-icon class="input-prefix"><User /></el-icon>
              <input
                v-model="loginForm.phone"
                type="text"
                :placeholder="t('login.usernamePlaceholder')"
                class="custom-input"
                @focus="clearError('phone')"
                @keyup.enter="handleSubmit"
              />
            </div>
            <span v-if="errors.phone" class="error-text">{{ errors.phone }}</span>
          </div>

          <!-- 密码 -->
          <div class="input-row">
            <label class="input-label">{{ t('login.password') }} <span class="required">*</span></label>
            <div class="input-wrapper">
              <el-icon class="input-prefix"><Lock /></el-icon>
              <input
                v-model="loginForm.password"
                type="password"
                :placeholder="t('login.passwordPlaceholder')"
                class="custom-input"
                @focus="clearError('password')"
                @keyup.enter="handleSubmit"
              />
            </div>
            <span v-if="errors.password" class="error-text">{{ errors.password }}</span>
          </div>

          <!-- 记住我 -->
          <div class="remember-row">
            <label class="checkbox-wrapper">
              <input type="checkbox" v-model="rememberMe" class="checkbox-input" />
              <span class="checkbox-custom"></span>
              <span class="checkbox-text">{{ t('login.rememberMe') }}</span>
            </label>
          </div>

          <!-- 登录按钮 -->
          <button
            class="login-btn"
            :class="{ 'loading': loading }"
            :disabled="loading"
            @click="handleSubmit"
          >
            <span v-if="!loading">{{ t('login.submit') }}</span>
            <span v-else class="loading-spinner"></span>
          </button>

          <!-- 注册链接 -->
          <div class="register-link-row">
            <span>{{ t('login.noAccount') }}</span>
            <router-link to="/user/register" class="register-link">{{ t('login.registerLink') }}</router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import { useLoginUserStore } from "@/store/user";
import { ElMessage } from "element-plus";
import { User, Lock } from "@element-plus/icons-vue";
import { useRouter } from "vue-router";
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const loading = ref(false);
const rememberMe = ref(false);
const router = useRouter();

const { login } = useLoginUserStore();

// 错误信息
const errors = reactive({
  phone: "",
  password: "",
});

const loginForm = reactive({
  phone: "",
  password: "",
} as API.UserLoginRequest);

// 清除错误
const clearError = (field: keyof typeof errors) => {
  errors[field] = "";
};

const handleSubmit = async () => {
  // 清空错误信息
  Object.keys(errors).forEach(key => {
    errors[key as keyof typeof errors] = "";
  });

  // 验证表单
  if (!loginForm.phone) {
    errors.phone = t('login.usernameRequired');
    ElMessage.error(t('login.usernameRequired'));
    return;
  }

  if (!loginForm.password) {
    errors.password = t('login.passwordRequired');
    ElMessage.error(t('login.passwordRequired'));
    return;
  }

  loading.value = true;
  try {
    await login(loginForm);
  } catch (error) {
    console.error(t('login.errorTitle'), error);
    ElMessage.error(t('login.errorMessage'));
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
.login-page {
  display: flex;
  min-height: 100vh;
  background: linear-gradient(135deg, #FDF2F8 0%, #F5F3FF 50%, #FCE7F3 100%);
  position: relative;
  overflow: hidden;
}

/* 登录表单区域 - 全屏居中 */
.login-form-section {
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
  max-width: 420px;
  background: rgba(255,255,255,0.65);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: 24px;
  padding: 48px 40px;
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

.error-text {
  font-size: 12px;
  color: #EF4444;
}

/* 记住我 */
.remember-row {
  display: flex;
  align-items: center;
}

.checkbox-wrapper {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  user-select: none;
}

.checkbox-input {
  display: none;
}

.checkbox-custom {
  width: 18px;
  height: 18px;
  border: 2px solid #E2E8F0;
  border-radius: 4px;
  position: relative;
  transition: all 0.2s ease;
}

.checkbox-input:checked + .checkbox-custom {
  background: #A855F7;
  border-color: #A855F7;
}

.checkbox-input:checked + .checkbox-custom::after {
  content: '';
  position: absolute;
  left: 5px;
  top: 2px;
  width: 4px;
  height: 8px;
  border: solid white;
  border-width: 0 2px 2px 0;
  transform: rotate(45deg);
}

.checkbox-text {
  font-size: 14px;
  color: #64748B;
}

/* 登录按钮 */
.login-btn {
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

.login-btn:hover:not(:disabled) {
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(219, 39, 119, 0.4);
}

.login-btn:disabled {
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

/* 注册链接 */
.register-link-row {
  text-align: center;
  font-size: 14px;
  color: #64748B;
  margin-top: 8px;
}

.register-link {
  color: #A855F7;
  font-weight: 600;
  text-decoration: none;
  margin-left: 4px;
}

.register-link:hover {
  text-decoration: underline;
}

/* 响应式设计 - 移动端优先 */
@media (max-width: 768px) {
  .login-page {
    flex-direction: column;
  }

  /* 移动端表单区域背景也改为白色 */
  .login-form-section {
    flex: 1;
    width: 100%;
    padding: 40px 24px;
    background: #ffffff;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .form-container {
    width: 100%;
    max-width: 420px;
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

  .form-subtitle {
    font-size: 14px;
  }

  .form-body {
    gap: 20px;
  }

  .custom-input {
    height: 48px;
  }

  .login-btn {
    height: 50px;
    margin-top: 8px;
  }
}

/* 超小屏幕进一步优化 */
@media (max-width: 480px) {
  .login-form-section {
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

  .login-btn {
    height: 50px;
    font-size: 16px;
  }
}
</style>
