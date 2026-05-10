<template>
  <div class="home-view">
    <!-- 顶部欢迎区域 - Hero Section -->
    <div class="hero-section">
      <div class="hero-background">
        <div class="gradient-orb orb-1"></div>
        <div class="gradient-orb orb-2"></div>
        <div class="gradient-orb orb-3"></div>
      </div>
      
      <div class="hero-content">
        <div class="hero-text">
          <div class="badge">
            <svg class="badge-icon" width="16" height="16">
              <use href="#icon-ai"></use>
            </svg>
            <span>{{ t('app.aiPowered') }}</span>
          </div>
          <h1 class="title">
            <span class="title-highlight">AI Pan</span>
            {{ t('home.hero.title') }}
          </h1>
          <p class="subtitle">
            {{ t('home.hero.subtitle') }}
          </p>
          <div class="hero-buttons">
            <DSButton
              variant="primary"
              size="lg"
              @click="goToFiles"
              class="cta-button"
            >
              <svg class="btn-icon" width="20" height="20">
                <use href="#icon-rocket"></use>
              </svg>
              {{ t('home.hero.cta') }}
            </DSButton>
            <DSButton
              variant="outline"
              size="lg"
              @click="scrollToFeatures"
              class="learn-more-button"
            >
              <svg class="btn-icon" width="20" height="20">
                <use href="#icon-info"></use>
              </svg>
              {{ t('home.hero.learnMore') }}
            </DSButton>
          </div>
        </div>
        
        <div class="hero-visual">
          <div class="visual-card glass-card">
            <img src="@/assets/file-management.svg" alt="文件管理" />
          </div>
        </div>
      </div>
    </div>

    <!-- 特性展示区域 -->
    <div class="features-section" ref="featuresSection">
      <div class="section-header">
        <DSTag variant="info" size="sm">
          <svg class="tag-icon" width="14" height="14">
            <use href="#icon-ai"></use>
          </svg>
          {{ t('home.features.ai') }}
        </DSTag>
        <h2 class="section-title">{{ t('home.features.title') }}</h2>
        <p class="section-description">
          {{ t('home.features.subtitle') }}
        </p>
      </div>
      
      <div class="feature-grid">
        <div class="feature-card glass-card cursor-pointer">
          <div class="feature-header">
            <div class="feature-icon-wrapper">
              <svg class="feature-icon" width="32" height="32">
                <use href="#icon-brain"></use>
              </svg>
            </div>
            <h3 class="feature-title">{{ t('home.features.qa.title') }}</h3>
          </div>
          <p class="feature-description">
            {{ t('home.features.qa.desc') }}
          </p>
        </div>

        <div class="feature-card glass-card cursor-pointer">
          <div class="feature-header">
            <div class="feature-icon-wrapper">
              <svg class="feature-icon" width="32" height="32">
                <use href="#icon-document"></use>
              </svg>
            </div>
            <h3 class="feature-title">{{ t('home.features.doc.title') }}</h3>
          </div>
          <p class="feature-description">
            {{ t('home.features.doc.desc') }}
          </p>
        </div>

        <div class="feature-card glass-card cursor-pointer">
          <div class="feature-header">
            <div class="feature-icon-wrapper">
              <svg class="feature-icon" width="32" height="32">
                <use href="#icon-chat"></use>
              </svg>
            </div>
            <h3 class="feature-title">{{ t('home.features.chat.title') }}</h3>
          </div>
          <p class="feature-description">
            {{ t('home.features.chat.desc') }}
          </p>
        </div>

        <div class="feature-card glass-card cursor-pointer">
          <div class="feature-header">
            <div class="feature-icon-wrapper">
              <svg class="feature-icon" width="32" height="32">
                <use href="#icon-search"></use>
              </svg>
            </div>
            <h3 class="feature-title">{{ t('home.features.search.title') }}</h3>
          </div>
          <p class="feature-description">
            {{ t('home.features.search.desc') }}
          </p>
        </div>
      </div>
    </div>

    <!-- 统计数据展示 -->
    <div class="stats-section">
      <div class="stats-container">
        <div class="stat-card glass-card cursor-pointer">
          <div class="stat-icon">
            <svg width="28" height="28">
              <use href="#icon-user"></use>
            </svg>
          </div>
          <div class="stat-content">
            <div class="stat-number" ref="userCount">
              <span class="count">100</span>
              <span class="unit">万+</span>
            </div>
            <div class="stat-label">{{ t('home.stats.users') }}</div>
          </div>
        </div>

        <div class="stat-card glass-card cursor-pointer">
          <div class="stat-icon">
            <svg width="28" height="28">
              <use href="#icon-check-circle"></use>
            </svg>
          </div>
          <div class="stat-content">
            <div class="stat-number">
              <span class="count">99.9</span>
              <span class="unit">%</span>
            </div>
            <div class="stat-label">{{ t('home.stats.availability') }}</div>
          </div>
        </div>

        <div class="stat-card glass-card cursor-pointer">
          <div class="stat-icon">
            <svg width="28" height="28">
              <use href="#icon-files"></use>
            </svg>
          </div>
          <div class="stat-content">
            <div class="stat-number">
              <span class="count">50</span>
              <span class="unit">PB+</span>
            </div>
            <div class="stat-label">{{ t('home.stats.storage') }}</div>
          </div>
        </div>

        <div class="stat-card glass-card cursor-pointer">
          <div class="stat-icon">
            <svg width="28" height="28">
              <use href="#icon-clock"></use>
            </svg>
          </div>
          <div class="stat-content">
            <div class="stat-number">
              <span class="count">24/7</span>
              <span class="unit"></span>
            </div>
            <div class="stat-label">{{ t('home.stats.support') }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- CTA区域 -->
    <div class="cta-section">
      <div class="cta-content">
        <h2 class="cta-title">{{ t('home.cta.title') }}</h2>
        <p class="cta-description">
          {{ t('home.cta.desc') }}
        </p>
        <div class="cta-buttons">
          <DSButton
            variant="primary"
            size="lg"
            @click="goToRegister"
            class="cta-primary"
          >
            <svg class="btn-icon" width="20" height="20">
              <use href="#icon-user-add"></use>
            </svg>
            {{ t('home.cta.register') }}
          </DSButton>
          <DSButton
            variant="outline"
            size="lg"
            @click="goToLogin"
            class="cta-secondary"
          >
            <svg class="btn-icon" width="20" height="20">
              <use href="#icon-login"></use>
            </svg>
            {{ t('home.cta.login') }}
          </DSButton>
        </div>
      </div>
    </div>

    <!-- Footer 区域 -->
    <footer class="footer-section">
      <div class="footer-content">
        <!-- 顶部品牌区 -->
        <div class="footer-top">
          <div class="footer-brand">
            <div class="footer-logo">
              <svg class="logo-icon" width="36" height="36">
                <use href="#icon-cloud-ai"></use>
              </svg>
              <span class="logo-text">AI Pan</span>
            </div>
            <p class="footer-slogan">{{ t('home.footer.slogan') }}</p>
          </div>

          <div class="footer-nav">
            <router-link to="/about" class="footer-link cursor-pointer">{{ t('home.footer.about') }}</router-link>
            <router-link to="/terms" class="footer-link cursor-pointer">{{ t('home.footer.terms') }}</router-link>
            <router-link to="/privacy" class="footer-link cursor-pointer">{{ t('home.footer.privacy') }}</router-link>
            <router-link to="/about" class="footer-link cursor-pointer">{{ t('home.footer.contact') }}</router-link>
          </div>
        </div>
        
        <!-- 分隔线 -->
        <div class="footer-divider"></div>
        
        <!-- 底部信息区 -->
        <div class="footer-bottom">
          <div class="footer-copyright">
            <span>{{ t('home.footer.copyright') }}</span>
          </div>
          
          <div class="footer-icp">
            <a href="https://beian.miit.gov.cn/" target="_blank" rel="noopener noreferrer" class="icp-link cursor-pointer">
              京ICP备XXXXXXXX号-1
            </a>
            <span class="icp-divider">|</span>
            <a href="http://www.beian.gov.cn/portal/registerSystemInfo" target="_blank" rel="noopener noreferrer" class="icp-link cursor-pointer">
              京公网安备 XXXXXXXXXXXX号
            </a>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';
import {
  Cpu,
  Document,
  ChatDotRound,
  Search,
  Check,
  Rocket,
  InfoFilled,
  User,
  CircleCheck,
  Files,
  Timer,
  UserFilled
} from '@element-plus/icons-vue';
import DSButton from '@/components/design-system/DSButton.vue';
import DSTag from '@/components/design-system/DSTag.vue';

const router = useRouter();
const { t } = useI18n();

// 跳转到文件页面
const goToFiles = () => {
  router.push('/file');
};

// 跳转到注册页面
const goToRegister = () => {
  router.push('/user/register');
};

// 跳转到登录页面
const goToLogin = () => {
  router.push('/user/login');
};

// 滚动到特性区域
const scrollToFeatures = () => {
  const featuresSection = document.querySelector('.features-section');
  if (featuresSection) {
    featuresSection.scrollIntoView({ behavior: 'smooth' });
  }
};
</script>

<style scoped>
/* 引入字体 */
@import url('https://fonts.googleapis.com/css2?family=Poppins:wght@400;500;600;700&family=Open+Sans:wght@300;400;500;600;700&display=swap');

/* CSS变量定义 */
:root {
  --primary: #DB2777;
  --secondary: #F472B6;
  --cta: #CA8A04;
  --background: #FDF2F8;
  --text-primary: #831843;
  --text-secondary: #475569;
  --font-primary: 'Poppins', 'Open Sans', sans-serif;
  --font-secondary: 'Open Sans', sans-serif;
}

.home-view {
  min-height: 100vh;
  background: linear-gradient(180deg, #FDF2F8 0%, #FCE7F3 100%);
  overflow-x: hidden;
  font-family: var(--font-secondary);
}

/* ===== Hero Section ===== */
.hero-section {
  position: relative;
  min-height: 100vh;
  display: flex;
  align-items: center;
  padding: 120px var(--ds-spacing-lg);
  overflow: hidden;
  background: linear-gradient(135deg, #FDF2F8 0%, #FCE7F3 50%, #FBCFE8 100%);
}

.hero-background {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  overflow: hidden;
  z-index: 0;
}

.gradient-orb {
  position: absolute;
  border-radius: 50%;
  filter: blur(100px);
  opacity: 0.4;
  animation: float 20s infinite ease-in-out;
}

.orb-1 {
  width: 500px;
  height: 500px;
  background: linear-gradient(135deg, #DB2777 0%, #F472B6 50%, #CA8A04 100%);
  top: -150px;
  right: -100px;
  animation-delay: 0s;
}

.orb-2 {
  width: 400px;
  height: 400px;
  background: linear-gradient(135deg, #F472B6 0%, #CA8A04 50%, #FBCFE8 100%);
  bottom: -100px;
  left: -100px;
  animation-delay: 7s;
}

.orb-3 {
  width: 300px;
  height: 300px;
  background: linear-gradient(135deg, #CA8A04 0%, #DB2777 50%, #F472B6 100%);
  top: 40%;
  left: 30%;
  transform: translate(-50%, -50%);
  animation-delay: 14s;
}

.hero-content {
  position: relative;
  z-index: 1;
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: var(--ds-spacing-xxl);
  align-items: center;
  width: 100%;
}

.hero-text {
  animation: slideInLeft 1s ease-out;
}

.badge {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  margin-bottom: var(--ds-spacing-lg);
  padding: 8px 16px;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(10px);
  border-radius: 20px;
  border: 1px solid rgba(219, 39, 119, 0.2);
  transition: all 0.3s ease;
}

.badge-icon {
  color: var(--primary);
}

.badge:hover {
  background: rgba(255, 255, 255, 1);
  border-color: rgba(219, 39, 119, 0.4);
  transform: translateY(-2px);
}

.title {
  font-size: 3.5rem;
  font-weight: 700;
  color: var(--text-primary);
  line-height: 1.2;
  margin-bottom: var(--ds-spacing-lg);
  font-family: var(--font-primary);
}

.title-highlight {
  display: block;
  background: linear-gradient(135deg, #DB2777 0%, #CA8A04 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.subtitle {
  font-size: var(--ds-text-size-xl);
  color: var(--text-secondary);
  line-height: 1.6;
  margin-bottom: var(--ds-spacing-xl);
  font-family: var(--font-secondary);
}

.hero-features {
  display: flex;
  flex-wrap: wrap;
  gap: var(--ds-spacing-md);
  margin-bottom: var(--ds-spacing-xl);
}

.feature-tag {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-xs);
  padding: var(--ds-spacing-sm) var(--ds-spacing-md);
  background: white;
  border-radius: var(--ds-radius-full);
  font-size: var(--ds-text-size-sm);
  color: var(--ds-color-text-secondary);
  box-shadow: var(--ds-shadow-sm);
}

.feature-tag el-icon {
  color: var(--ds-color-success);
}

.hero-buttons {
  display: flex;
  gap: var(--ds-spacing-md);
}

.btn-icon {
  color: currentColor;
  transition: transform 0.3s ease;
}

.cta-button:hover .btn-icon,
.learn-more-button:hover .btn-icon {
  transform: translateX(2px);
}

.hero-visual {
  display: flex;
  justify-content: center;
  align-items: center;
  animation: slideInRight 1s ease-out;
}

/* 玻璃态卡片效果 */
.glass-card {
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
}

.visual-card {
  position: relative;
  width: 100%;
  max-width: 500px;
  padding: var(--ds-spacing-lg);
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: 20px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 8px 32px rgba(219, 39, 119, 0.1);
  overflow: hidden;
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.visual-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #DB2777 0%, #CA8A04 100%);
}

.visual-card:hover {
  transform: translateY(-8px) scale(1.02);
  box-shadow: 0 20px 48px rgba(219, 39, 119, 0.15);
}

.visual-card img {
  width: 100%;
  height: auto;
  display: block;
}

/* ===== Features Section ===== */
.features-section {
  padding: 100px var(--ds-spacing-lg);
  background: white;
}

.section-header {
  text-align: center;
  margin-bottom: var(--ds-spacing-xxl);
}

.section-title {
  font-size: var(--ds-text-size-xxl);
  font-weight: var(--ds-text-weight-bold);
  color: var(--ds-color-text-primary);
  margin: var(--ds-spacing-md) 0;
}

.section-description {
  font-size: var(--ds-text-size-lg);
  color: var(--ds-color-text-secondary);
  max-width: 600px;
  margin: 0 auto;
}

.feature-grid {
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: var(--ds-spacing-xl);
}

.feature-card {
  position: relative;
  padding: var(--ds-spacing-xl);
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: 20px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 8px 32px rgba(219, 39, 119, 0.05);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  overflow: hidden;
}

.feature-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: linear-gradient(90deg, #DB2777 0%, #CA8A04 100%);
  transform: scaleX(0);
  transition: transform 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.feature-card::after {
  content: '';
  position: absolute;
  inset: 0;
  background: linear-gradient(135deg, rgba(219, 39, 119, 0.03) 0%, rgba(202, 138, 4, 0.03) 100%);
  opacity: 0;
  transition: opacity 0.4s ease;
}

.feature-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(219, 39, 119, 0.15);
  border-color: rgba(219, 39, 119, 0.2);
}

.feature-card:hover::before {
  transform: scaleX(1);
}

.feature-card:hover::after {
  opacity: 1;
}

.feature-header {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-md);
  margin-bottom: var(--ds-spacing-md);
}

.feature-header .feature-title {
  display: flex;
  align-items: center;
}

.feature-icon-wrapper {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 64px;
  height: 64px;
  border-radius: 16px;
  background: linear-gradient(135deg, #DB2777 0%, #F472B6 100%);
  box-shadow: 0 8px 24px rgba(219, 39, 119, 0.3);
  flex-shrink: 0;
  transition: all 0.3s ease;
}

.feature-card:hover .feature-icon-wrapper {
  transform: scale(1.1) rotate(5deg);
  box-shadow: 0 12px 32px rgba(219, 39, 119, 0.4);
}

.feature-icon {
  color: white;
  transition: transform 0.3s ease;
}

.feature-card:hover .feature-icon {
  transform: scale(1.1);
}

.feature-title {
  font-size: 22px;
  font-weight: 700;
  color: var(--text-primary);
  margin: 0;
  font-family: var(--font-primary);
  background: linear-gradient(135deg, #831843 0%, #DB2777 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.feature-description {
  font-size: var(--ds-text-size-lg);
  color: var(--ds-color-text-secondary);
  line-height: 1.7;
}

.feature-list {
  list-style: none;
  padding: 0;
  margin: 0;
}

.feature-list li {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-xs);
  padding: var(--ds-spacing-xs) 0;
  font-size: var(--ds-text-size-sm);
  color: var(--ds-color-text-secondary);
}

.feature-list li el-icon {
  color: var(--ds-color-success);
  font-size: 16px;
}

/* ===== Stats Section ===== */
.stats-section {
  padding: 80px var(--ds-spacing-lg);
  background: linear-gradient(135deg, #DB2777 0%, #F472B6 50%, #CA8A04 100%);
  position: relative;
  overflow: hidden;
}

.stats-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: url("data:image/svg+xml,%3Csvg width='60' height='60' viewBox='0 0 60 60' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='none' fill-rule='evenodd'%3E%3Cg fill='%23ffffff' fill-opacity='0.05'%3E%3Cpath d='M36 34v-4h-2v4h-4v2h4v4h2v-4h4v-2h-4zm0-30V0h-2v4h-4v2h4v4h2V6h4V4h-4zM6 34v-4H4v4H0v2h4v4h2v-4h4v-2H6zM6 4V0H4v4H0v2h4v4h2V6h4V4H6z'/%3E%3C/g%3E%3C/g%3E%3C/svg%3E");
  opacity: 0.5;
}

.stats-container {
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: var(--ds-spacing-xl);
  position: relative;
  z-index: 1;
}

.stat-card {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-md);
  padding: var(--ds-spacing-xl);
  background: rgba(255, 255, 255, 0.15);
  backdrop-filter: blur(20px) saturate(180%);
  -webkit-backdrop-filter: blur(20px) saturate(180%);
  border-radius: 20px;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  overflow: hidden;
}

.stat-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 2px;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.5), transparent);
  transform: translateX(-100%);
  transition: transform 0.6s ease;
}

.stat-card:hover {
  transform: translateY(-6px);
  background: rgba(255, 255, 255, 0.2);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
}

.stat-card:hover::before {
  transform: translateX(100%);
}

.stat-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 60px;
  height: 60px;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.2) 0%, rgba(255, 255, 255, 0.1) 100%);
  border-radius: 16px;
  font-size: 28px;
  color: white;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
}

.stat-card:hover .stat-icon {
  transform: scale(1.1) rotate(5deg);
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.3) 0%, rgba(255, 255, 255, 0.2) 100%);
}

.stat-content {
  flex: 1;
}

.stat-number {
  display: flex;
  align-items: baseline;
  gap: var(--ds-spacing-xs);
  margin-bottom: var(--ds-spacing-xs);
}

.stat-number .count {
  font-size: var(--ds-text-size-xxl);
  font-weight: var(--ds-text-weight-bold);
  color: white;
}

.stat-number .unit {
  font-size: var(--ds-text-size-lg);
  color: rgba(255, 255, 255, 0.9);
}

.stat-label {
  font-size: var(--ds-text-size-sm);
  color: rgba(255, 255, 255, 0.8);
}

/* ===== CTA Section ===== */
.cta-section {
  padding: 100px var(--ds-spacing-lg);
  background: linear-gradient(180deg, #ffffff 0%, #FDF2F8 100%);
  position: relative;
}

.cta-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(219, 39, 119, 0.2), transparent);
}

.cta-content {
  max-width: 800px;
  margin: 0 auto;
  text-align: center;
  position: relative;
}

.cta-title {
  font-size: 3rem;
  font-weight: 700;
  font-family: var(--font-primary);
  background: linear-gradient(135deg, #831843 0%, #DB2777 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: var(--ds-spacing-md);
}

.cta-description {
  font-size: var(--ds-text-size-lg);
  color: var(--ds-color-text-secondary);
  margin-bottom: var(--ds-spacing-xl);
}

.cta-buttons {
  display: flex;
  justify-content: center;
  gap: var(--ds-spacing-md);
}

.cta-primary,
.cta-secondary {
  display: flex;
  align-items: center;
  gap: var(--ds-spacing-sm);
  padding: 16px 32px;
  border-radius: 12px;
  font-size: 16px;
  font-weight: 600;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.cta-primary {
  background: linear-gradient(135deg, #DB2777 0%, #CA8A04 100%);
  color: white;
  border: none;
  box-shadow: 0 4px 16px rgba(219, 39, 119, 0.4);
  transition: all 0.3s ease;
}

.cta-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 24px rgba(219, 39, 119, 0.5);
}

.cta-secondary {
  background: white;
  color: #DB2777;
  border: 2px solid rgba(219, 39, 119, 0.2);
  transition: all 0.3s ease;
}

.cta-secondary:hover {
  background: rgba(219, 39, 119, 0.05);
  border-color: rgba(219, 39, 119, 0.4);
  transform: translateY(-2px);
}

/* ===== Footer Section ===== */
.footer-section {
  padding: 60px var(--ds-spacing-lg) 30px;
  background: linear-gradient(180deg, #FDF2F8 0%, #FCE7F3 100%);
  position: relative;
}

.footer-section::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(219, 39, 119, 0.2), transparent);
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 30px;
}

/* Footer Top - 品牌和导航 */
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
  color: #DB2777;
  transition: transform 0.3s ease;
}

.footer-logo:hover .logo-icon {
  transform: scale(1.1) rotate(5deg);
}

.logo-text {
  font-size: 28px;
  font-weight: 700;
  font-family: var(--font-primary);
  background: linear-gradient(135deg, #DB2777 0%, #CA8A04 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.footer-slogan {
  font-size: 15px;
  color: #64748b;
  font-weight: 500;
  padding-left: 48px;
}

.footer-nav {
  display: flex;
  align-items: center;
  gap: 40px;
}

.footer-link {
  font-size: 15px;
  color: #475569;
  text-decoration: none;
  font-weight: 500;
  transition: all 0.3s ease;
  position: relative;
  padding: 8px 0;
}

.footer-link::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 0;
  height: 2px;
  background: linear-gradient(90deg, #DB2777 0%, #CA8A04 100%);
  transition: width 0.3s ease;
}

.footer-link:hover {
  color: #DB2777;
}

.footer-link:hover::after {
  width: 100%;
}

/* cursor-pointer utility class */
.cursor-pointer {
  cursor: pointer;
}

/* Footer Divider */
.footer-divider {
  width: 100%;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(102, 126, 234, 0.2), transparent);
}

/* Footer Bottom - 版权和备案 */
.footer-bottom {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 20px;
}

.footer-copyright {
  font-size: 14px;
  color: #94a3b8;
  font-weight: 500;
}

.footer-icp {
  display: flex;
  align-items: center;
  gap: 16px;
}

.icp-link {
  font-size: 13px;
  color: #64748b;
  text-decoration: none;
  transition: all 0.3s ease;
}

.icp-link:hover {
  color: #667eea;
}

.icp-divider {
  color: #cbd5e1;
}

/* ===== Animations ===== */
@keyframes float {
  0%, 100% {
    transform: translateY(0) scale(1) rotate(0deg);
  }
  25% {
    transform: translateY(-40px) scale(1.05) rotate(5deg);
  }
  50% {
    transform: translateY(20px) scale(0.95) rotate(-5deg);
  }
  75% {
    transform: translateY(-20px) scale(1.02) rotate(3deg);
  }
}

@keyframes slideInLeft {
  from {
    opacity: 0;
    transform: translateX(-50px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

@keyframes slideInRight {
  from {
    opacity: 0;
    transform: translateX(50px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

/* ===== Responsive Design ===== */
@media (max-width: 1200px) {
  .feature-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  
  .stats-container {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 968px) {
  .hero-content {
    grid-template-columns: 1fr;
    text-align: center;
  }
  
  .hero-text {
    order: 1;
  }
  
  .hero-visual {
    order: 0;
    margin-bottom: var(--ds-spacing-xl);
  }
  
  .hero-features {
    justify-content: center;
  }
  
  .hero-buttons {
    justify-content: center;
  }
  
  .title {
    font-size: 2.5rem;
  }
  
  .visual-card {
    max-width: 400px;
  }
}

@media (max-width: 768px) {
  .hero-section {
    padding: var(--ds-spacing-xl) var(--ds-spacing-md);
    min-height: auto;
    padding-top: var(--ds-spacing-xxl);
  }
  
  .hero-visual {
    display: none;
  }
  
  .title {
    font-size: 2rem;
  }
  
  .subtitle {
    font-size: var(--ds-text-size-lg);
  }
  
  .feature-grid {
    grid-template-columns: 1fr;
    gap: var(--ds-spacing-lg);
  }
  
  .feature-card {
    padding: var(--ds-spacing-lg);
  }
  
  .feature-header {
    gap: var(--ds-spacing-sm);
  }
  
  .feature-icon-wrapper {
    width: 48px;
    height: 48px;
  }
  
  .feature-icon {
    font-size: 24px;
  }
  
  .feature-title {
    font-size: 18px;
  }
  
  .stats-container {
    grid-template-columns: 1fr;
    gap: var(--ds-spacing-md);
  }
  
  .stat-card {
    padding: var(--ds-spacing-md);
    flex-direction: row;
    align-items: center;
  }
  
  .stat-icon {
    width: 44px;
    height: 44px;
    font-size: 20px;
  }
  
  .stat-number .count {
    font-size: 20px;
  }
  
  .stat-label {
    font-size: 13px;
  }
  
  .hero-buttons {
    flex-direction: column;
    align-items: stretch;
  }
  
  .cta-buttons {
    flex-direction: column;
    align-items: stretch;
  }
  
  .gradient-orb {
    display: none;
  }
  
  .section-header {
    margin-bottom: var(--ds-spacing-lg);
  }
  
  .section-title {
    font-size: var(--ds-text-size-xl);
  }
  
  .section-description {
    font-size: var(--ds-text-size-base);
  }
}

@media (max-width: 576px) {
  .title {
    font-size: 1.75rem;
  }
  
  .section-title {
    font-size: var(--ds-text-size-xl);
  }
  
  .stat-card {
    padding: var(--ds-spacing-md);
  }
  
  .stat-icon {
    width: 48px;
    height: 48px;
    font-size: 20px;
  }
  
  .stat-number .count {
    font-size: var(--ds-text-size-xl);
  }
}

/* Footer Responsive */
@media (max-width: 768px) {
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
</style>
