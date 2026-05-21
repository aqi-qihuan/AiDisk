<template>
  <div class="error-page not-found">
    <div class="error-bg">
      <div class="orb orb-1"></div>
      <div class="orb orb-2"></div>
    </div>
    <div class="error-content glass-card">
      <!-- 404 SVG Illustration -->
      <div class="error-illustration">
        <svg
          width="180"
          height="140"
          viewBox="0 0 180 140"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <circle
            cx="90"
            cy="70"
            r="60"
            stroke="rgba(219,39,119,0.2)"
            stroke-width="2"
            fill="none"
          />
          <circle
            cx="90"
            cy="70"
            r="40"
            stroke="rgba(251,191,36,0.15)"
            stroke-width="2"
            fill="none"
            stroke-dasharray="6 4"
          />
          <text
            x="90"
            y="82"
            text-anchor="middle"
            font-size="52"
            font-weight="800"
            font-family="'Plus Jakarta Sans', sans-serif"
            fill="url(#pinkGold)"
          >
            404
          </text>
          <defs>
            <linearGradient id="pinkGold" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stop-color="#DB2777" />
              <stop offset="100%" stop-color="#FBBF24" />
            </linearGradient>
          </defs>
        </svg>
      </div>

      <h1 class="error-code">404</h1>
      <h2 class="error-title">{{ $t("error.404.message") }}</h2>
      <p class="error-desc">{{ $t("error.404.description") }}</p>

      <div class="error-actions">
        <DSButton variant="primary" @click="$router.push('/')">
          <svg
            width="16"
            height="16"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
          >
            <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" />
            <polyline points="9 22 9 12 15 12 15 22" />
          </svg>
          {{ $t("error.backHome") }}
        </DSButton>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import DSButton from "@/components/design-system/DSButton.vue";
</script>

<style scoped>
.error-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #0b0b10;
  position: relative;
  overflow: hidden;
  padding: 24px;
}

.error-bg {
  position: absolute;
  inset: 0;
  pointer-events: none;
}

.orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(100px);
  opacity: 0.2;
}

.orb-1 {
  width: 400px;
  height: 400px;
  background: #d4a853;
  top: -150px;
  right: -100px;
  animation: float 20s ease-in-out infinite;
}

.orb-2 {
  width: 300px;
  height: 300px;
  background: #d97706;
  bottom: -100px;
  left: -80px;
  animation: float 25s ease-in-out infinite reverse;
}

@keyframes float {
  0%,
  100% {
    transform: translate(0, 0);
  }
  50% {
    transform: translate(30px, -30px);
  }
}

.error-content {
  position: relative;
  z-index: 1;
  text-align: center;
  padding: 48px;
  max-width: 480px;
  width: 100%;
  background: rgba(255, 255, 255, 0.04);
  backdrop-filter: blur(24px) saturate(180%);
  -webkit-backdrop-filter: blur(24px) saturate(180%);
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-radius: 24px;
  box-shadow: 0 24px 80px rgba(0, 0, 0, 0.4);
}

.error-illustration {
  margin-bottom: 16px;
}

.error-code {
  font-size: 0;
  margin: 0;
}

.error-title {
  font-size: 28px;
  font-weight: 700;
  font-family: "Plus Jakarta Sans", sans-serif;
  color: #f8fafc;
  margin: 0 0 12px;
  letter-spacing: -0.02em;
}

.error-desc {
  font-size: 16px;
  color: #94a3b8;
  margin: 0 0 32px;
  line-height: 1.6;
}

.error-actions {
  display: flex;
  justify-content: center;
  gap: 12px;
}

@media (max-width: 768px) {
  .error-content {
    padding: 32px 24px;
    border-radius: 20px;
  }

  .error-illustration svg {
    width: 140px;
    height: 110px;
  }

  .error-title {
    font-size: 22px;
  }

  .error-desc {
    font-size: 14px;
  }

  .error-actions {
    flex-direction: column;
  }

  .error-actions :deep(.ds-btn) {
    width: 100%;
  }
}

@media (max-width: 576px) {
  .error-content {
    padding: 24px 20px;
    border-radius: 16px;
  }

  .error-title {
    font-size: 20px;
  }
}
</style>
