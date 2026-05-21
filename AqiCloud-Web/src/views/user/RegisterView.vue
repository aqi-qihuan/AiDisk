<template>
  <div class="auth-hok">
    <!-- 光晕装饰 (HOK设计) -->
    <div class="hok-glow-orb pink"></div>
    <div class="hok-glow-orb gold"></div>
    <div class="hok-glow-orb purple"></div>

    <div class="auth-card-hok">
      <div class="auth-card-header-hok">
        <div class="auth-logo-hok">A</div>
        <h2 class="auth-title-hok">{{ t("register.title") }}</h2>
        <p class="auth-subtitle-hok">{{ t("register.subtitle") }}</p>
      </div>

      <div class="auth-body">
        <!-- 头像上传 -->
        <div class="avatar-row">
          <div class="avatar-wrapper">
            <ImageUpload v-model="avatarUrl" />
          </div>
          <span class="avatar-hint">{{ t("register.avatarHint") }}</span>
        </div>

        <!-- 用户名 -->
        <div class="auth-input-hok">
          <svg
            width="16"
            height="16"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            class="auth-input-icon-hok"
          >
            <path d="M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2" />
            <circle cx="12" cy="7" r="4" />
          </svg>
          <input
            v-model="registerForm.phone"
            type="text"
            :placeholder="t('register.usernamePlaceholder')"
            @focus="clearError('phone')"
          />
        </div>
        <span v-if="errors.phone" class="error-text">{{ errors.phone }}</span>

        <!-- 密码 -->
        <div class="auth-input-hok">
          <svg
            width="16"
            height="16"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            class="auth-input-icon-hok"
          >
            <rect x="3" y="11" width="18" height="11" rx="2" ry="2" />
            <path d="M7 11V7a5 5 0 0110 0v4" />
          </svg>
          <input
            v-model="registerForm.password"
            type="password"
            :placeholder="t('register.passwordPlaceholder')"
            @focus="clearError('password')"
          />
        </div>
        <span class="hint-text">{{ t("register.passwordHint") }}</span>
        <span v-if="errors.password" class="error-text">{{
          errors.password
        }}</span>

        <!-- 确认密码 -->
        <div class="auth-input-hok">
          <svg
            width="16"
            height="16"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            class="auth-input-icon-hok"
          >
            <rect x="3" y="11" width="18" height="11" rx="2" ry="2" />
            <path d="M7 11V7a5 5 0 0110 0v4" />
          </svg>
          <input
            v-model="registerForm.confirmPassword"
            type="password"
            :placeholder="t('register.confirmPasswordPlaceholder')"
            @focus="clearError('confirmPassword')"
          />
        </div>
        <span v-if="errors.confirmPassword" class="error-text">{{
          errors.confirmPassword
        }}</span>

        <!-- 注册按钮 -->
        <button
          class="auth-btn-hok"
          :disabled="loading"
          @click="handleRegister"
        >
          <svg
            width="20"
            height="20"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2.5"
          >
            <path
              d="M15 3h4a2 2 0 012 2v14a2 2 0 01-2 2h-4M10 17l5-5-5-5M13 12H3"
            />
          </svg>
          {{ loading ? t("common.loading") : t("register.submit") }}
        </button>

        <!-- 登录链接 -->
        <p class="auth-footer-text">
          {{ t("register.hasAccount") }}
          <router-link to="/user/login" class="auth-link"
            >{{ t("register.loginLink") }} →</router-link
          >
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";
import { useI18n } from "vue-i18n";
import ImageUpload from "@/components/common/ImageUpload.vue";
import DSButton from "@/components/design-system/DSButton.vue";
import { userRegister } from "@/api/user";

const { t } = useI18n();
const router = useRouter();
const loading = ref(false);
const avatarUrl = ref("");

const errors = reactive({
  phone: "",
  password: "",
  confirmPassword: "",
});

const registerForm = reactive({
  phone: "",
  password: "",
  confirmPassword: "",
});

const clearError = (field: keyof typeof errors) => {
  errors[field] = "";
};

const handleRegister = async () => {
  Object.keys(errors).forEach((key) => {
    errors[key as keyof typeof errors] = "";
  });

  if (!registerForm.phone) {
    errors.phone = t("register.usernameRequired");
    return;
  }

  if (!registerForm.password) {
    errors.password = t("register.passwordRequired");
    return;
  }

  if (registerForm.password.length < 6) {
    errors.password = t("register.passwordTooShort");
    return;
  }

  if (!registerForm.confirmPassword) {
    errors.confirmPassword = t("register.confirmPasswordRequired");
    return;
  }

  if (registerForm.password !== registerForm.confirmPassword) {
    errors.confirmPassword = t("register.passwordMismatch");
    return;
  }

  loading.value = true;
  try {
    const res = await userRegister({
      phone: registerForm.phone,
      password: registerForm.password,
      avatarUrl: avatarUrl.value,
    });
    if (res.data.code === 0) {
      ElMessage.success(t("register.success"));
      router.push("/user/login");
    } else {
      ElMessage.error(res.data.msg || t("register.error"));
    }
  } catch (error) {
    console.error("注册失败:", error);
    ElMessage.error(t("register.error"));
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
/* ===== HOK 暗色注册卡片 ===== */
.auth-hok {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100%;
  overflow: hidden;
}

.hok-glow-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(120px);
  pointer-events: none;
  opacity: 0.4;
}

.hok-glow-orb.pink {
  width: 400px;
  height: 400px;
  top: -100px;
  right: -50px;
  background: radial-gradient(circle, #d4a853, transparent);
}

.hok-glow-orb.gold {
  width: 300px;
  height: 300px;
  bottom: -50px;
  left: -50px;
  background: radial-gradient(circle, #d97706, transparent);
}

.hok-glow-orb.purple {
  width: 200px;
  height: 200px;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: radial-gradient(circle, #7c3aed, transparent);
  opacity: 0.15;
}

.auth-card-hok {
  position: relative;
  z-index: 1;
  width: 100%;
  max-width: 420px;
  padding: 40px 36px;
  background: rgba(26, 26, 36, 0.85);
  backdrop-filter: blur(24px) saturate(180%);
  -webkit-backdrop-filter: blur(24px) saturate(180%);
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-radius: 16px;
  box-shadow:
    0 8px 32px rgba(0, 0, 0, 0.4),
    0 2px 8px rgba(0, 0, 0, 0.2);
}

.auth-card-header-hok {
  text-align: center;
  margin-bottom: 28px;
}

.auth-logo-hok {
  width: 56px;
  height: 56px;
  margin: 0 auto 16px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-family: "Plus Jakarta Sans", "Fira Code", sans-serif;
  font-size: 28px;
  font-weight: 800;
  color: #0b0b10;
  background: linear-gradient(135deg, #d97706 0%, #f59e0b 50%, #fbbf24 100%);
  border-radius: 14px;
  box-shadow:
    0 4px 20px rgba(217, 119, 6, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.3);
}

.auth-title-hok {
  font-family: "Plus Jakarta Sans", "Fira Code", sans-serif;
  font-size: 24px;
  font-weight: 700;
  color: #f8fafc;
  margin: 0 0 8px;
  letter-spacing: -0.02em;
}

.auth-subtitle-hok {
  font-size: 14px;
  color: #94a3b8;
  margin: 0;
}

.auth-body {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

/* 头像 */
.avatar-row {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 8px;
  padding-bottom: 12px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
}

.avatar-wrapper {
  width: 72px;
  height: 72px;
  border-radius: 50%;
  overflow: hidden;
  border: 2px solid rgba(219, 39, 119, 0.3);
  transition: border-color 0.2s ease;
}

.avatar-wrapper:hover {
  border-color: #d4a853;
}

.avatar-hint {
  font-size: 12px;
  color: #64748b;
}

/* 输入框 - HOK 暗色玻璃风格 */
.auth-input-hok {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 8px 16px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-radius: 8px;
  transition: all 250ms cubic-bezier(0.4, 0, 0.2, 1);
  font-family:
    "Fira Sans",
    "Plus Jakarta Sans",
    -apple-system,
    sans-serif;
  font-size: 14px;
}

.auth-input-hok:focus-within {
  border-color: rgba(219, 39, 119, 0.3);
  box-shadow: 0 0 12px rgba(219, 39, 119, 0.08);
  background: rgba(255, 255, 255, 0.05);
}

.auth-input-icon-hok {
  color: #64748b;
  flex-shrink: 0;
}

.auth-input-hok input {
  border: none;
  outline: none;
  flex: 1;
  background: transparent;
  font-family: inherit;
  font-size: inherit;
  color: #f8fafc;
  height: 32px;
}

.auth-input-hok input::placeholder {
  color: #64748b;
}

.hint-text {
  font-size: 12px;
  color: #64748b;
  margin-top: -8px;
}

.error-text {
  font-size: 12px;
  color: #ef4444;
  margin-top: -8px;
}

/* 注册按钮 - 金色渐变 */
.auth-btn-hok {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  width: 100%;
  height: 48px;
  border: none;
  border-radius: 8px;
  background: linear-gradient(135deg, #d97706, #f59e0b, #fbbf24);
  color: #0b0b10;
  font-family: "Plus Jakarta Sans", "Fira Code", sans-serif;
  font-size: 16px;
  font-weight: 700;
  cursor: pointer;
  transition: all 250ms cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow:
    0 4px 20px rgba(217, 119, 6, 0.3),
    inset 0 1px 0 rgba(255, 255, 255, 0.3);
  letter-spacing: 0.02em;
}

.auth-btn-hok:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow:
    0 8px 32px rgba(217, 119, 6, 0.4),
    inset 0 1px 0 rgba(255, 255, 255, 0.3);
}

.auth-btn-hok:active:not(:disabled) {
  transform: translateY(0);
}

.auth-btn-hok:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.auth-footer-text {
  text-align: center;
  font-size: 13px;
  color: #64748b;
  margin: 0;
}

.auth-link {
  color: #fbbf24;
  font-weight: 600;
  text-decoration: none;
  transition: color 0.2s ease;
}

.auth-link:hover {
  color: #fde68a;
}
</style>
