<template>
  <div id="user-layout">
    <!-- 背景光晕 -->
    <div class="auth-bg">
      <div class="orb orb-1"></div>
      <div class="orb orb-2"></div>
      <div class="orb orb-3"></div>
      <div class="grid-pattern"></div>
    </div>

    <!-- 居中玻璃卡片 -->
    <div class="auth-container">
      <div class="auth-card">
        <!-- Logo -->
        <div class="auth-logo">
          <div class="logo-icon-wrapper">
            <svg
              width="28"
              height="28"
              viewBox="0 0 24 24"
              fill="none"
              stroke="white"
              stroke-width="2"
            >
              <path
                d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z"
              />
            </svg>
          </div>
          <span class="auth-brand">小七云盘</span>
        </div>

        <!-- 表单内容 -->
        <div class="auth-form">
          <router-view v-slot="{ Component }">
            <transition name="page" mode="out-in">
              <component :is="Component" />
            </transition>
          </router-view>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped>
#user-layout {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #0b0b10;
  position: relative;
  overflow: hidden;
}

/* ===== 背景光晕 ===== */
.auth-bg {
  position: fixed;
  inset: 0;
  pointer-events: none;
  overflow: hidden;
}

.orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(120px);
  opacity: 0.25;
  animation: float 20s ease-in-out infinite;
}

.orb-1 {
  width: 500px;
  height: 500px;
  background: linear-gradient(135deg, #d4a853, #c9a96e);
  top: -200px;
  right: -100px;
  animation-delay: 0s;
}

.orb-2 {
  width: 400px;
  height: 400px;
  background: linear-gradient(135deg, #d97706, #fbbf24);
  bottom: -150px;
  left: -100px;
  animation-delay: -7s;
}

.orb-3 {
  width: 300px;
  height: 300px;
  background: linear-gradient(135deg, #d4a853, #d97706);
  top: 40%;
  left: 50%;
  transform: translate(-50%, -50%);
  animation-delay: -14s;
}

@keyframes float {
  0%,
  100% {
    transform: translate(0, 0) scale(1);
  }
  25% {
    transform: translate(30px, -30px) scale(1.05);
  }
  50% {
    transform: translate(-20px, 20px) scale(0.95);
  }
  75% {
    transform: translate(20px, 30px) scale(1.02);
  }
}

.grid-pattern {
  position: absolute;
  inset: 0;
  background-image:
    linear-gradient(rgba(255, 255, 255, 0.02) 1px, transparent 1px),
    linear-gradient(90deg, rgba(255, 255, 255, 0.02) 1px, transparent 1px);
  background-size: 50px 50px;
}

/* ===== 居中卡片 ===== */
.auth-container {
  position: relative;
  z-index: 1;
  width: 100%;
  max-width: 460px;
  padding: 24px;
}

.auth-card {
  width: 100%;
  background: rgba(26, 26, 36, 0.85);
  backdrop-filter: blur(40px) saturate(180%);
  -webkit-backdrop-filter: blur(40px) saturate(180%);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 24px;
  box-shadow:
    0 24px 80px rgba(0, 0, 0, 0.5),
    0 0 60px rgba(219, 39, 119, 0.05);
  padding: 40px;
  position: relative;
  overflow: hidden;
}

.auth-card::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: linear-gradient(90deg, #d4a853, #d97706, #fbbf24);
}

/* ===== Logo ===== */
.auth-logo {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 14px;
  margin-bottom: 32px;
}

.logo-icon-wrapper {
  width: 48px;
  height: 48px;
  background: linear-gradient(135deg, #d4a853, #c9a96e);
  border-radius: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 8px 24px rgba(219, 39, 119, 0.3);
  flex-shrink: 0;
}

.auth-brand {
  font-size: 26px;
  font-weight: 700;
  font-family: "Plus Jakarta Sans", sans-serif;
  background: linear-gradient(135deg, #d4a853 0%, #fbbf24 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  letter-spacing: -0.5px;
}

/* ===== 表单区域 ===== */
.auth-form {
  width: 100%;
}

/* ===== 响应式 ===== */
@media (max-width: 768px) {
  .auth-container {
    padding: 16px;
  }

  .auth-card {
    padding: 32px 24px;
    border-radius: 20px;
  }

  .auth-brand {
    font-size: 22px;
  }

  .logo-icon-wrapper {
    width: 44px;
    height: 44px;
  }
}

@media (max-width: 576px) {
  .auth-card {
    padding: 24px 20px;
    border-radius: 16px;
  }

  .auth-logo {
    margin-bottom: 24px;
  }

  .orb-1,
  .orb-2,
  .orb-3 {
    opacity: 0.15;
  }
}
</style>
