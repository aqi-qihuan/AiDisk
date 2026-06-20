<template>
  <div class="auth-hok">
    <!-- 光晕装饰 (HOK设计) -->
    <div class="hok-glow-orb pink"></div>
    <div class="hok-glow-orb gold"></div>
    <div class="hok-glow-orb purple"></div>

    <div class="auth-card-hok">
      <div class="auth-card-header-hok">
        <div class="auth-logo-hok">A</div>
        <h2 class="auth-title-hok">{{ t("login.title") }}</h2>
        <p class="auth-subtitle-hok">{{ t("login.subtitle") }}</p>
      </div>

      <div class="auth-body">
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
            v-model="loginForm.phone"
            type="text"
            :placeholder="t('login.usernamePlaceholder')"
            @focus="clearError('phone')"
            @keyup.enter="handleSubmit"
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
            v-model="loginForm.password"
            type="password"
            :placeholder="t('login.passwordPlaceholder')"
            @focus="clearError('password')"
            @keyup.enter="handleSubmit"
          />
        </div>
        <span v-if="errors.password" class="error-text">{{
          errors.password
        }}</span>

        <!-- 记住我 + 忘记密码 -->
        <div class="auth-options">
          <label class="auth-remember">
            <div class="checkbox-hok">
              <svg
                width="9"
                height="9"
                viewBox="0 0 24 24"
                fill="none"
                stroke="white"
                stroke-width="3"
              >
                <polyline points="20 6 9 17 4 12" />
              </svg>
            </div>
            <span>{{ t("login.rememberMe") }}</span>
          </label>
          <span class="auth-forgot">{{ t("auth.forgotPassword") }}</span>
        </div>

        <!-- 登录按钮 -->
        <button class="auth-btn-hok" :disabled="loading" @click="handleSubmit">
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
          {{ loading ? t("common.loading") : t("login.submit") }}
        </button>

        <!-- 注册链接 -->
        <p class="auth-footer-text">
          {{ t("login.noAccount") }}
          <router-link to="/user/register" class="auth-link"
            >{{ t("login.registerLink") }} →</router-link
          >
        </p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from "vue";
import { useLoginUserStore } from "@/store/user";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";
import { useI18n } from "vue-i18n";
import DSButton from "@/components/design-system/DSButton.vue";

const { t } = useI18n();
const loading = ref(false);
const router = useRouter();

const { login } = useLoginUserStore();

const errors = reactive({
  phone: "",
  password: "",
});

const loginForm = reactive({
  phone: "",
  password: "",
} as API.UserLoginRequest);

const clearError = (field: keyof typeof errors) => {
  errors[field] = "";
};

const handleSubmit = async () => {
  Object.keys(errors).forEach((key) => {
    errors[key as keyof typeof errors] = "";
  });

  if (!loginForm.phone) {
    errors.phone = t("login.usernameRequired");
    ElMessage.error(t("login.usernameRequired"));
    return;
  }

  if (!loginForm.password) {
    errors.password = t("login.passwordRequired");
    ElMessage.error(t("login.passwordRequired"));
    return;
  }

  loading.value = true;
  try {
    await login(loginForm);
  } catch (error) {
    console.error(t("login.errorTitle"), error);
    ElMessage.error(t("login.errorMessage"));
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
/* ===== HOK 暗色登录卡片 ===== */
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
  background: radial-gradient(circle, #db2777, transparent);
  opacity: 0.12;
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
  margin: 16px;
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

.error-text {
  font-size: 12px;
  color: #ef4444;
  margin-top: -8px;
}

/* 选项行 */
.auth-options {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.auth-remember {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  font-size: 13px;
  color: #94a3b8;
}

.checkbox-hok {
  width: 16px;
  height: 16px;
  border-radius: 4px;
  background: rgba(219, 39, 119, 0.3);
  display: flex;
  align-items: center;
  justify-content: center;
}

.auth-forgot {
  font-size: 13px;
  color: #fbbf24;
  cursor: pointer;
  transition: color 0.2s ease;
}

.auth-forgot:hover {
  color: #fde68a;
}

/* 登录按钮 - 金色渐变 */
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

/* ===== 移动端适配 - HOK 标准 ===== */
@media (max-width: 768px) {
  .auth-hok {
    padding: 0;
    align-items: flex-start;
  }

  .hok-glow-orb.pink {
    width: 250px;
    height: 250px;
    top: -60px;
    right: -80px;
  }

  .hok-glow-orb.gold {
    width: 200px;
    height: 200px;
    bottom: -30px;
    left: -80px;
  }

  .hok-glow-orb.purple {
    width: 120px;
    height: 120px;
  }

  .auth-card-hok {
    max-width: 100%;
    margin: 0;
    border-radius: 0;
    border-left: none;
    border-right: none;
    padding: 48px 24px 32px;
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  .auth-logo-hok {
    width: 48px;
    height: 48px;
    font-size: 24px;
    border-radius: 12px;
  }

  .auth-title-hok {
    font-size: 22px;
  }

  .auth-input-hok {
    padding: 10px 14px;
  }

  .auth-input-hok input {
    height: 36px;
    font-size: 16px; /* 防止 iOS 自动缩放 */
  }

  .auth-btn-hok {
    height: 52px; /* 触控友好最小 44px */
    font-size: 16px;
    border-radius: 12px;
  }

  .auth-options {
    flex-wrap: wrap;
    gap: 8px;
  }
}

@media (max-width: 480px) {
  .auth-card-hok {
    padding: 40px 18px 28px;
  }

  .auth-body {
    gap: 14px;
  }

  .auth-title-hok {
    font-size: 20px;
  }

  .auth-subtitle-hok {
    font-size: 13px;
  }

  .auth-remember,
  .auth-forgot,
  .auth-footer-text {
    font-size: 13px;
  }

  .auth-btn-hok {
    height: 50px;
  }
}
</style>
