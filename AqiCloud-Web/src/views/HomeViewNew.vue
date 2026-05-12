<template>
  <div class="home-view">
    <!-- Hero Section -->
    <div class="hero-section">
      <div class="hero-content">
        <div class="hero-text">
          <!-- Badge -->
          <div class="badge">
            <svg
              class="badge-icon"
              width="16"
              height="16"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
            >
              <path
                d="M9.663 17h4.674M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 1 1 7.072 0l-.548.547A3.374 3.374 0 0 0 14 18.469V19a2 2 0 1 1-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"
              />
            </svg>
            <span>AI 驱动</span>
          </div>

          <!-- Title -->
          <h1 class="title">
            <span class="title-highlight">AI Pan</span>
            智能云盘
          </h1>

          <!-- Subtitle -->
          <p class="subtitle">让 AI 帮你管理文件，更智能、更高效</p>

          <!-- Buttons -->
          <div class="hero-buttons">
            <DSButton
              variant="primary"
              size="lg"
              @click="goToFiles"
              class="cta-button"
            >
              <svg
                class="btn-icon"
                width="20"
                height="20"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <path d="M13 10V3L4 14h7v7l9-11h-7z" />
              </svg>
              立即体验
            </DSButton>
            <DSButton
              variant="secondary"
              size="lg"
              @click="scrollToFeatures"
              class="learn-more-button"
            >
              <svg
                class="btn-icon"
                width="20"
                height="20"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <circle cx="12" cy="12" r="10" />
                <path d="M12 16v-4M12 8h.01" />
              </svg>
              了解更多
            </DSButton>
          </div>
        </div>

        <!-- Visual Card -->
        <div class="hero-visual">
          <div class="visual-card">
            <div class="visual-header">
              <div class="visual-dots">
                <span class="dot dot-red"></span>
                <span class="dot dot-yellow"></span>
                <span class="dot dot-green"></span>
              </div>
              <span class="visual-title">文件管理系统</span>
            </div>
            <div class="visual-body">
              <div
                class="file-item"
                v-for="(file, index) in demoFiles"
                :key="index"
              >
                <svg
                  width="20"
                  height="20"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="#6B7280"
                  stroke-width="2"
                >
                  <path
                    d="M7 21h10a2 2 0 0 0 2-2V9.414a1 1 0 0 0-.293-.707l-5.414-5.414A1 1 0 0 0 12.586 3H7a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2z"
                  />
                </svg>
                <span class="file-name">{{ file.name }}</span>
                <span class="file-size">{{ file.size }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Features Section -->
    <div class="features-section" ref="featuresSection">
      <div class="section-header">
        <DSTag variant="info" size="sm">
          <svg
            class="tag-icon"
            width="14"
            height="14"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
          >
            <path
              d="M9.663 17h4.674M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 1 1 7.072 0l-.548.547A3.374 3.374 0 0 0 14 18.469V19a2 2 0 1 1-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"
            />
          </svg>
          AI 赋能
        </DSTag>
        <h2 class="section-title">核心功能</h2>
        <p class="section-description">
          融合 AI 技术，让文件管理变得智能、高效、安全
        </p>
      </div>

      <div class="feature-grid">
        <div
          class="feature-card"
          v-for="(feature, index) in features"
          :key="index"
        >
          <div class="feature-icon-wrapper">
            <svg
              class="feature-icon"
              width="32"
              height="32"
              viewBox="0 0 24 24"
              fill="none"
              :stroke="feature.color"
              stroke-width="2"
            >
              <path :d="feature.icon" />
            </svg>
          </div>
          <h3 class="feature-title">{{ feature.title }}</h3>
          <p class="feature-description">{{ feature.description }}</p>
        </div>
      </div>
    </div>

    <!-- Stats Section -->
    <div class="stats-section">
      <div class="stats-container">
        <div class="stat-card" v-for="(stat, index) in stats" :key="index">
          <div class="stat-icon">
            <svg
              width="28"
              height="28"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
            >
              <path :d="stat.icon" />
            </svg>
          </div>
          <div class="stat-content">
            <div class="stat-number">
              <span class="count">{{ stat.value }}</span>
              <span class="unit">{{ stat.unit }}</span>
            </div>
            <div class="stat-label">{{ stat.label }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- CTA Section -->
    <div class="cta-section">
      <div class="cta-content">
        <h2 class="cta-title">准备好开始了吗？</h2>
        <p class="cta-description">立即注册，免费体验所有 AI 功能</p>
        <div class="cta-buttons">
          <DSButton
            variant="primary"
            size="lg"
            @click="goToRegister"
            class="cta-primary"
          >
            <svg
              class="btn-icon"
              width="20"
              height="20"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
            >
              <path
                d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-7.372a12.004 12.004 0 0 1 7.072 0m0 0a12.004 12.004 0 0 1 7.072 0m-7.072 0A12 12 0 1 0 5.928 7.628m12.144 0a12 12 0 0 1-7.072 0"
              />
            </svg>
            免费注册
          </DSButton>
          <DSButton
            variant="secondary"
            size="lg"
            @click="goToLogin"
            class="cta-secondary"
          >
            <svg
              class="btn-icon"
              width="20"
              height="20"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
            >
              <path
                d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1"
              />
            </svg>
            登录账号
          </DSButton>
        </div>
      </div>
    </div>

    <!-- Footer Section -->
    <footer class="footer-section">
      <div class="footer-content">
        <div class="footer-top">
          <div class="footer-brand">
            <div class="footer-logo">
              <svg
                class="logo-icon"
                width="36"
                height="36"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
              >
                <path
                  d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z"
                />
              </svg>
              <span class="logo-text">AI Pan</span>
            </div>
            <p class="footer-slogan">AI 驱动的智能云盘</p>
          </div>

          <div class="footer-nav">
            <router-link to="/about" class="footer-link">关于我们</router-link>
            <router-link to="/terms" class="footer-link">使用条款</router-link>
            <router-link to="/privacy" class="footer-link"
              >隐私政策</router-link
            >
            <router-link to="/about" class="footer-link">联系我们</router-link>
          </div>
        </div>

        <div class="footer-divider"></div>

        <div class="footer-bottom">
          <div class="footer-copyright">
            <span>© 2026 AI Pan. All rights reserved.</span>
          </div>

          <div class="footer-icp">
            <a
              href="https://beian.miit.gov.cn/"
              target="_blank"
              rel="noopener noreferrer"
              class="icp-link"
            >
              京ICP备XXXXXXXX号-1
            </a>
            <span class="icp-divider">|</span>
            <a
              href="http://www.beian.gov.cn/"
              target="_blank"
              rel="noopener noreferrer"
              class="icp-link"
            >
              京公网安备 XXXXXXXXXXXX号
            </a>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useI18n } from "vue-i18n";
import DSButton from "@/components/design-system/DSButton.vue";
import DSTag from "@/components/design-system/DSTag.vue";

const router = useRouter();
const { t } = useI18n();

// Demo data
const demoFiles = ref([
  { name: "项目方案.pdf", size: "2.5 MB" },
  { name: "设计素材", size: "--" },
  { name: "会议纪要.docx", size: "856 KB" },
  { name: "数据分析.xlsx", size: "1.2 MB" },
]);

const features = ref([
  {
    title: "智能问答",
    description: "用自然语言查询文件，AI 秒速响应",
    icon: "M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z",
    color: "#3B82F6",
  },
  {
    title: "文档助手",
    description: "AI 自动整理和分析文档内容",
    icon: "M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z",
    color: "#10B981",
  },
  {
    title: "AI 聊天",
    description: "与 AI 对话，快速获取文件信息",
    icon: "M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z",
    color: "#F59E0B",
  },
  {
    title: "智能搜索",
    description: "语义搜索，快速定位目标文件",
    icon: "M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z",
    color: "#EF4444",
  },
]);

const stats = ref([
  {
    label: "活跃用户",
    value: "100",
    unit: "万+",
    icon: "M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z",
  },
  {
    label: "可用性",
    value: "99.9",
    unit: "%",
    icon: "M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01 8.618 3.04A12.02 12.02 0 013 13.69a12.088 12.088 0 01 1.254 4.06c.723 1.308 1.752 2.355 3.003 3.084",
  },
  {
    label: "存储容量",
    value: "50",
    unit: "PB+",
    icon: "M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z",
  },
  {
    label: "客户支持",
    value: "24/7",
    unit: "",
    icon: "M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z",
  },
]);

// Navigation methods
const goToFiles = () => {
  router.push("/file");
};

const goToRegister = () => {
  router.push("/user/register");
};

const goToLogin = () => {
  router.push("/user/login");
};

const scrollToFeatures = () => {
  const featuresSection = document.querySelector(".features-section");
  if (featuresSection) {
    featuresSection.scrollIntoView({ behavior: "smooth" });
  }
};
</script>

<style scoped>
/* Import Inter font */
@import url("https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap");

.home-view {
  min-height: 100vh;
  background: var(--color-gray-50, #f9fafb);
  font-family:
    "Inter",
    -apple-system,
    BlinkMacSystemFont,
    "Segoe UI",
    sans-serif;
}

/* ===== Hero Section ===== */
.hero-section {
  padding: 120px 48px 80px;
  background: white;
  min-height: 80vh;
  display: flex;
  align-items: center;
}

.hero-content {
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 64px;
  align-items: center;
  width: 100%;
}

.hero-text {
  animation: slideInLeft 0.6s ease-out;
}

.badge {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 6px 12px;
  background: var(--color-primary-light, #eff6ff);
  color: var(--color-primary, #3b82f6);
  border-radius: 9999px;
  font-size: 13px;
  font-weight: 500;
  margin-bottom: 24px;
  transition: background-color 0.15s ease;
}

.badge:hover {
  background: #dbeafe;
}

.badge-icon {
  color: var(--color-primary, #3b82f6);
}

.title {
  font-size: 48px;
  font-weight: 700;
  color: var(--color-gray-900, #111827);
  line-height: 1.2;
  margin-bottom: 24px;
  letter-spacing: -0.02em;
}

.title-highlight {
  display: block;
  color: var(--color-primary, #3b82f6);
}

.subtitle {
  font-size: 18px;
  color: var(--color-gray-600, #4b5563);
  line-height: 1.6;
  margin-bottom: 40px;
  max-width: 500px;
}

.hero-buttons {
  display: flex;
  gap: 16px;
  align-items: center;
}

.btn-icon {
  color: currentColor;
  transition: transform 0.15s ease;
}

.cta-button:hover .btn-icon,
.learn-more-button:hover .btn-icon {
  transform: translateX(2px);
}

/* Hero Visual */
.hero-visual {
  display: flex;
  justify-content: center;
  align-items: center;
  animation: slideInRight 0.6s ease-out;
}

.visual-card {
  width: 100%;
  max-width: 480px;
  background: white;
  border: 1px solid var(--color-gray-200, #e5e7eb);
  border-radius: 8px;
  overflow: hidden;
  transition: border-color 0.15s ease;
}

.visual-card:hover {
  border-color: var(--color-gray-300, #d1d5db);
}

.visual-header {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 16px;
  background: var(--color-gray-50, #f9fafb);
  border-bottom: 1px solid var(--color-gray-200, #e5e7eb);
}

.visual-dots {
  display: flex;
  gap: 6px;
}

.dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
}

.dot-red {
  background: #ef4444;
}
.dot-yellow {
  background: #f59e0b;
}
.dot-green {
  background: #10b981;
}

.visual-title {
  font-size: 13px;
  color: var(--color-gray-600, #4b5563);
  font-weight: 500;
}

.visual-body {
  padding: 16px;
}

.file-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px;
  border-radius: 6px;
  transition: background-color 0.15s ease;
}

.file-item:hover {
  background: var(--color-gray-50, #f9fafb);
}

.file-name {
  flex: 1;
  font-size: 14px;
  color: var(--color-gray-900, #111827);
}

.file-size {
  font-size: 13px;
  color: var(--color-gray-500, #6b7280);
}

/* ===== Features Section ===== */
.features-section {
  padding: 100px 48px;
  background: white;
}

.section-header {
  text-align: center;
  margin-bottom: 64px;
}

.section-title {
  font-size: 30px;
  font-weight: 700;
  color: var(--color-gray-900, #111827);
  margin: 16px 0;
  letter-spacing: -0.02em;
}

.section-description {
  font-size: 16px;
  color: var(--color-gray-600, #4b5563);
  max-width: 600px;
  margin: 0 auto;
  line-height: 1.6;
}

.feature-grid {
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
}

.feature-card {
  padding: 32px 24px;
  background: white;
  border: 1px solid var(--color-gray-200, #e5e7eb);
  border-radius: 8px;
  transition: border-color 0.15s ease;
}

.feature-card:hover {
  border-color: var(--color-gray-300, #d1d5db);
}

.feature-icon-wrapper {
  width: 48px;
  height: 48px;
  display: flex;
  align-items: center;
  justify-content: center;
  background: var(--color-gray-50, #f9fafb);
  border-radius: 8px;
  margin-bottom: 16px;
  transition: background-color 0.15s ease;
}

.feature-card:hover .feature-icon-wrapper {
  background: var(--color-gray-100, #f3f4f6);
}

.feature-icon {
  transition: transform 0.15s ease;
}

.feature-card:hover .feature-icon {
  transform: scale(1.05);
}

.feature-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--color-gray-900, #111827);
  margin: 0 0 8px 0;
}

.feature-description {
  font-size: 14px;
  color: var(--color-gray-600, #4b5563);
  line-height: 1.5;
  margin: 0;
}

/* ===== Stats Section ===== */
.stats-section {
  padding: 80px 48px;
  background: var(--color-primary, #3b82f6);
}

.stats-container {
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
}

.stat-card {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 24px;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 8px;
  transition: background-color 0.15s ease;
}

.stat-card:hover {
  background: rgba(255, 255, 255, 0.15);
}

.stat-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 48px;
  height: 48px;
  color: white;
  flex-shrink: 0;
}

.stat-content {
  flex: 1;
}

.stat-number {
  display: flex;
  align-items: baseline;
  gap: 4px;
  margin-bottom: 4px;
}

.stat-number .count {
  font-size: 30px;
  font-weight: 700;
  color: white;
  letter-spacing: -0.02em;
}

.stat-number .unit {
  font-size: 16px;
  color: rgba(255, 255, 255, 0.9);
}

.stat-label {
  font-size: 14px;
  color: rgba(255, 255, 255, 0.8);
}

/* ===== CTA Section ===== */
.cta-section {
  padding: 100px 48px;
  background: white;
}

.cta-content {
  max-width: 800px;
  margin: 0 auto;
  text-align: center;
}

.cta-title {
  font-size: 36px;
  font-weight: 700;
  color: var(--color-gray-900, #111827);
  margin-bottom: 16px;
  letter-spacing: -0.02em;
}

.cta-description {
  font-size: 16px;
  color: var(--color-gray-600, #4b5563);
  margin-bottom: 40px;
  line-height: 1.6;
}

.cta-buttons {
  display: flex;
  justify-content: center;
  gap: 16px;
}

.cta-primary,
.cta-secondary {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 24px;
  border-radius: 6px;
  font-size: 15px;
  font-weight: 500;
  transition: all 0.15s ease;
}

.cta-primary {
  background: var(--color-primary, #3b82f6);
  color: white;
  border: none;
}

.cta-primary:hover {
  background: var(--color-primary-hover, #2563eb);
}

.cta-secondary {
  background: white;
  color: var(--color-gray-700, #374151);
  border: 1px solid var(--color-gray-200, #e5e7eb);
}

.cta-secondary:hover {
  background: var(--color-gray-50, #f9fafb);
  border-color: var(--color-gray-300, #d1d5db);
}

/* ===== Footer Section ===== */
.footer-section {
  padding: 60px 48px 30px;
  background: var(--color-gray-50, #f9fafb);
  border-top: 1px solid var(--color-gray-200, #e5e7eb);
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 30px;
}

.footer-top {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 30px;
}

.footer-brand {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.footer-logo {
  display: flex;
  align-items: center;
  gap: 12px;
}

.logo-icon {
  color: var(--color-primary, #3b82f6);
  transition: transform 0.15s ease;
}

.footer-logo:hover .logo-icon {
  transform: scale(1.05);
}

.logo-text {
  font-size: 24px;
  font-weight: 700;
  color: var(--color-gray-900, #111827);
}

.footer-slogan {
  font-size: 14px;
  color: var(--color-gray-500, #6b7280);
  padding-left: 48px;
}

.footer-nav {
  display: flex;
  align-items: center;
  gap: 32px;
}

.footer-link {
  font-size: 14px;
  color: var(--color-gray-600, #4b5563);
  text-decoration: none;
  font-weight: 500;
  transition: color 0.15s ease;
  padding: 4px 0;
}

.footer-link:hover {
  color: var(--color-primary, #3b82f6);
}

.footer-divider {
  width: 100%;
  height: 1px;
  background: var(--color-gray-200, #e5e7eb);
}

.footer-bottom {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 20px;
}

.footer-copyright {
  font-size: 13px;
  color: var(--color-gray-500, #6b7280);
  font-weight: 500;
}

.footer-icp {
  display: flex;
  align-items: center;
  gap: 16px;
}

.icp-link {
  font-size: 13px;
  color: var(--color-gray-500, #6b7280);
  text-decoration: none;
  transition: color 0.15s ease;
}

.icp-link:hover {
  color: var(--color-primary, #3b82f6);
}

.icp-divider {
  color: var(--color-gray-300, #d1d5db);
}

/* ===== Animations ===== */
@keyframes slideInLeft {
  from {
    opacity: 0;
    transform: translateX(-20px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

@keyframes slideInRight {
  from {
    opacity: 0;
    transform: translateX(20px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

/* ===== Responsive Design ===== */
@media (max-width: 1024px) {
  .hero-content {
    grid-template-columns: 1fr;
    gap: 48px;
    text-align: center;
  }

  .hero-text {
    order: 1;
  }

  .hero-visual {
    order: 0;
  }

  .subtitle {
    margin-left: auto;
    margin-right: auto;
  }

  .hero-buttons {
    justify-content: center;
  }

  .feature-grid {
    grid-template-columns: repeat(2, 1fr);
  }

  .stats-container {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 768px) {
  .hero-section {
    padding: 80px 24px 60px;
    min-height: auto;
  }

  .title {
    font-size: 32px;
  }

  .subtitle {
    font-size: 16px;
  }

  .hero-visual {
    display: none;
  }

  .features-section {
    padding: 60px 24px;
  }

  .section-title {
    font-size: 24px;
  }

  .feature-grid {
    grid-template-columns: 1fr;
    gap: 16px;
  }

  .feature-card {
    padding: 24px 20px;
  }

  .stats-section {
    padding: 60px 24px;
  }

  .stats-container {
    grid-template-columns: 1fr;
    gap: 16px;
  }

  .stat-card {
    padding: 20px;
  }

  .cta-section {
    padding: 60px 24px;
  }

  .cta-title {
    font-size: 28px;
  }

  .cta-buttons {
    flex-direction: column;
    align-items: stretch;
  }

  .footer-section {
    padding: 40px 24px 20px;
  }

  .footer-top {
    flex-direction: column;
    text-align: center;
  }

  .footer-slogan {
    padding-left: 0;
  }

  .footer-nav {
    gap: 24px;
  }

  .footer-bottom {
    flex-direction: column;
    text-align: center;
  }
}

@media (max-width: 576px) {
  .title {
    font-size: 28px;
  }

  .hero-buttons {
    flex-direction: column;
    align-items: stretch;
  }

  .section-title {
    font-size: 20px;
  }

  .cta-title {
    font-size: 24px;
  }
}
</style>
